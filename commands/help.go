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
		msg := ""
		categories = GetCategories()
		for cat, cmds := range categories {
			msg += "__" + cat + "__\n"

			for _, c := range cmds {
				msg += "**"
				names := c.names()
				for _, n := range names {
					msg += n + ", "
				}
				msg = msg[:len(msg)-2]
				msg += "** - " + c.description() + "\n"
			}

			msg += "\n"
		}

		usrChannel, err := ctx.Session.UserChannelCreate(ctx.Author.ID)
		errors.Warning(err, "Error creating user channel")
		_, err = ctx.Session.ChannelMessageSend(usrChannel.ID, msg)
		errors.Warning(err, "Error sending DM")

	} else if len(args) == 1 {
		commands := GetCommands()
		cmdName := args[0]
		cmdToHelp := commands[cmdName]

		names := cmdToHelp.names()

		msg := "**"
		msg += cmdName
		if len(names) > 1 {
			msg += " (aka "
			for _, n := range names {
				if n != cmdName {
					msg += n + ", "
				}
			}
			msg = msg[:len(msg)-2]
			msg += ")"
		}
		msg += "**\n" + cmdToHelp.description()

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
