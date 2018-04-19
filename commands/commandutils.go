package commands

import (
	"strings"

	"../logger"
	"../yamlutils"
	"github.com/bwmarrin/discordgo"
)

// Command interface stores functions for commands
type Command interface {
	execute(*Context, []string)
	description() string
	category() string
	numArgs() (int, int)
	names() []string
	onlyOwner() bool
	usage() []string
}

// Context stores the context of a command
type Context struct {
	Guild   *discordgo.Guild
	Message *discordgo.MessageCreate
	Author  *discordgo.User
	Session *discordgo.Session
	Content string
}

// Send messages in discord into the same channel as the command
func (ctx *Context) send(s string) (msg *discordgo.Message, err error) {
	msg, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, s)
	logger.Warning(err, "Could not send message")
	return
}

func (ctx *Context) sendEmbed(em *discordgo.MessageEmbed) (msg *discordgo.Message, err error) {
	msg, err = ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID, em)
	logger.Warning(err, "Could not send embed")
	return
}

// Reply is the same as Send, but appends a mention to the user who did the command
func (ctx *Context) reply(s string) (msg *discordgo.Message, err error) {
	msg, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, ctx.Author.Mention()+" "+s)
	logger.Warning(err, "Could not send message")
	return
}

func (ctx *Context) wrongNumArgs(cmd string) (msg *discordgo.Message, err error) {
	msg, err = ctx.reply("Incorrect number of arguments for command `" + cmd + "`. Try `!help " + cmd + "`.")
	return
}

func (ctx *Context) wrongPerms(cmd string) (msg *discordgo.Message, err error) {
	msg, err = ctx.reply("the command `" + cmd + "` can only be used by the owner of this bot!")
	return
}

var commands map[string]Command
var categories map[string][]Command

// AddCommands binds each Command to a message
func AddCommands(c ...Command) {
	for _, command := range c {
		AddCommand(command)
	}
}

// AddCommand binds a Command to a message
func AddCommand(c Command) {
	if commands == nil {
		commands = make(map[string]Command)
	}

	if categories == nil {
		categories = make(map[string][]Command)
	}

	for _, name := range c.names() {
		commands[name] = c
	}

	categories[c.category()] = append(categories[c.category()], c)
}

// Handle calls a command if the message is a command
func Handle(msg *discordgo.MessageCreate, session *discordgo.Session) {
	content := msg.Content
	channel, _ := session.Channel(msg.ChannelID)
	isCmd := false
	fullCmd := ""

	if channel.GuildID == "" {
		fullCmd = content
		isCmd = true
	} else {
		prefix, ok := getMsgPrefix(msg, session)
		if ok {
			fullCmd = content[len(prefix):]
			isCmd = true
		}
	}

	if isCmd {
		splitCmd := strings.Fields(fullCmd)
		cmdName := strings.ToLower(splitCmd[0])
		args := splitCmd[1:]

		cmd := commands[cmdName]
		if cmd != nil {
			if cmd.onlyOwner() == true && msg.Author.ID != yamlutils.GetOwnerID() {
				session.ChannelMessageSend(msg.ChannelID, msg.Author.Mention()+", the command `"+cmdName+"` can only be used by the owner of this bot!")
				return
			}
			min, max := cmd.numArgs()

			// if it has a wrong amount of arguments, exit
			if (len(args) < min && min != -1) || (len(args) > max && max != -1) {
				session.ChannelMessageSend(msg.ChannelID, msg.Author.Mention()+" Incorrect number of arguments for command `"+cmdName+"`. Try `!help "+cmdName+"`.")
				return
			}

			go session.ChannelTyping(msg.ChannelID)
			go cmd.execute(createContext(msg, session, fullCmd), args)
		}
	}
}

func createContext(msg *discordgo.MessageCreate, session *discordgo.Session, content string) (ctx *Context) {
	ctx = new(Context)

	channel, _ := session.Channel(msg.ChannelID)
	guild, _ := session.Guild(channel.GuildID)

	ctx.Guild = guild
	ctx.Message = msg
	ctx.Author = msg.Author
	ctx.Session = session
	ctx.Content = content

	return
}

// GetCategories gives the category data required to make a help command
func GetCategories() map[string][]Command {
	return categories
}

// GetCommands gives the command data required to make a help command
func GetCommands() map[string]Command {
	return commands
}

func getMsgPrefix(msg *discordgo.MessageCreate, session *discordgo.Session) (string, bool) {
	channel, err := session.Channel(msg.ChannelID)
	logger.Warning(err, "Couldn't get channel")

	prefix, ok := yamlutils.GetPrefix(channel.GuildID)

	if ok {
		if strings.HasPrefix(msg.Content, prefix) {
			return prefix, true
		}
	}

	mention := session.State.User.Mention()

	if strings.HasPrefix(msg.Content, mention) {
		return mention, true
	}

	return "", false
}
