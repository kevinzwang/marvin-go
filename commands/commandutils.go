package commands

import (
	"strings"

	"../logger"
	"github.com/bwmarrin/discordgo"
)

var prefix *string
var ownerID *string

// Init should be called by main()
func Init(p *string, o *string) {
	prefix = p
	ownerID = o
}

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
}

// Send messages in discord into the same channel as the command
func (ctx *Context) send(s string) (msg *discordgo.Message, err error) {
	msg, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, s)
	logger.Warning(err, "Could not send message")
	return
}

// Reply is the same as Send, but appends a mention to the user who did the command
func (ctx *Context) reply(s string) (msg *discordgo.Message, err error) {
	msg, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, ctx.Author.Mention()+" "+s)
	logger.Warning(err, "Could not send message")
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

	if strings.HasPrefix(content, *prefix) {
		if channel.GuildID == "" {
			session.ChannelMessageSend(msg.ChannelID, "You don't need to include prefixes in DMs")
		}

		fullCmd = content[len(*prefix):]
		isCmd = true

	} else if channel.GuildID == "" {
		fullCmd = content
		isCmd = true
	}

	if isCmd {
		splitCmd := strings.Fields(fullCmd)
		cmdName := strings.ToLower(splitCmd[0])
		args := splitCmd[1:]

		cmd := commands[cmdName]
		if cmd != nil {
			if cmd.onlyOwner() == true && msg.Author.ID != *ownerID {
				session.ChannelMessageSend(msg.ChannelID, msg.Author.Mention()+", the command `"+cmdName+"` can only be used by the owner of this bot!")
				return
			}
			min, max := cmd.numArgs()

			// if it has a wrong amount of arguments, exit
			if (len(args) < min && min != -1) || (len(args) > max && max != -1) {
				session.ChannelMessageSend(msg.ChannelID, "Incorrect number of arguments for command `"+cmdName+"`")
				return
			}

			session.ChannelTyping(msg.ChannelID)
			go cmd.execute(createContext(msg, session), args)
		}
	}
}

func createContext(msg *discordgo.MessageCreate, session *discordgo.Session) (ctx *Context) {
	ctx = new(Context)

	channel, _ := session.Channel(msg.ChannelID)
	guild, _ := session.Guild(channel.GuildID)

	ctx.Guild = guild
	ctx.Message = msg
	ctx.Author = msg.Author
	ctx.Session = session

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
