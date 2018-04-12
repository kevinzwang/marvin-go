package commands

import (
	"../errors"
)

// Help gives info on all the commands
type Help struct{}

func (cmd *Help) execute(ctx *Context, args []string) {
	if len(args) == 0 {
		if ctx.Guild != nil {
			ctx.reply(":mailbox_with_mail:")
		}
		msg := "These are the commands for Marvin, the Chaos server admin bot.\n\nFor more info about a command, do `help <command>`.\n\n"
		categories = GetCategories()
		for cat, cmds := range categories {
			msg += "__" + cat + "__\n"

			for _, c := range cmds {
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
				msg += "** - " + c.description() + "\n"
			}

			msg += "\n"
		}

		usrChannel, err := ctx.Session.UserChannelCreate(ctx.Author.ID)
		if err != nil {
			errors.Warning("Could not create user channel")
		}
		_, err = ctx.Session.ChannelMessageSend(usrChannel.ID, msg)
		if err != nil {
			errors.Warning("Could not send message to DM")
		}

	} else if len(args) == 1 {
		commands := GetCommands()
		cmdName := args[0]
		cmdToHelp := commands[cmdName]

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
		msg += "**\n" + cmdToHelp.description() + "\n\n__Usage__\n"

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
