package commands

import (
	"strings"

	"../logger"
)

// Help gives info on all the commands
type Help struct{}

func (cmd *Help) execute(ctx *Context, args []string) {
	if len(args) == 0 {
		if ctx.Guild != nil {
			ctx.reply(":mailbox_with_mail:")
		}
		msg := "These are the commands for Marvin, the Chaos server admin bot.\nFor more info about a command, do `help <command>`.\n\n(Note: you don't need to use prefixes in this DM.)\n\n"
		categories = GetCategories()
		for cat, cmds := range categories {
			msg += "__" + cat + "__\n"

			for _, c := range cmds {
				if c.onlyOwner() {
					msg += "\\*"
				}
				msg += "**"
				names := c.names()
				mainName := names[0]
				msg += mainName
				altNames := names[1:]
				if len(altNames) > 0 {
					msg += " (aka "
					for _, n := range altNames {
						msg += n + ", "
					}
					msg = msg[:len(msg)-2]
					msg += ")"
				}
				desc := c.description()
				if nlIndex := strings.Index(desc, "\n"); nlIndex != -1 {
					desc = desc[:nlIndex] + "..."
				}
				if len(desc) > 100 {
					desc = desc[:97] + "..."
				}
				msg += "** - " + desc + "\n"
			}

			msg += "\n"
		}

		msg += "`*` = bot owner only"

		usrChannel, err := ctx.Session.UserChannelCreate(ctx.Author.ID)
		logger.Warning(err, "Could not create user channel")

		_, err = ctx.Session.ChannelMessageSend(usrChannel.ID, msg)
		logger.Warning(err, "Could not send message to DM")

	} else if len(args) == 1 {
		commands := GetCommands()
		cmdName := args[0]
		cmdToHelp, ok := commands[cmdName]

		if !ok {
			ctx.reply("No such command `" + cmdName + "`.")
			return
		}

		names := cmdToHelp.names()

		mainName := names[0]
		altNames := names[1:]

		msg := "**"
		msg += mainName
		if len(altNames) > 0 {
			msg += " (aka "
			for _, n := range altNames {
				msg += n + ", "
			}
			msg = msg[:len(msg)-2]
			msg += ")"
		}
		msg += "**\n" + cmdToHelp.description()

		if cmdToHelp.onlyOwner() {
			msg += "\n\n(Note: Only the bot owner can use this command.)"
		}

		msg += "\n\n__Usage__\n"

		usage := cmdToHelp.usage()
		for _, u := range usage {
			msg += "\n`" + mainName + " " + u + "`"
		}

		ctx.send(msg)
	}
}

func (cmd *Help) description() string {
	return "gives info on all the commands"
}

func (cmd *Help) category() string { return "misc" }

func (cmd *Help) numArgs() (int, int) { return 0, 1 }

func (cmd *Help) names() []string { return []string{"help", "halp"} }

func (cmd *Help) onlyOwner() bool { return false }

func (cmd *Help) usage() []string { return []string{"", "<command>"} }
