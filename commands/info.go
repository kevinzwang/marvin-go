package commands

// Info gives information about the bot
type Info struct{}

func (cmd *Info) execute(ctx *Context, args []string) {
	ctx.send("Marvin is the official administrative bot of the Chaos Discord server, written in Golang (<http://discord.gg/mXnshyA>)." +
		"\nIf you are interested in the development of this bot, you are invited to join the official Chaos Dev server: <http://discord.gg/yvqTyGt>" +
		"\nYou can also find all of the bot's source code here: https://github.com/kevinzwang/marvin-go")
}

func (cmd *Info) description() string { return "gives information about this bot" }

func (cmd *Info) category() string { return "misc" }

func (cmd *Info) numArgs() (int, int) { return 0, 0 }

func (cmd *Info) names() []string { return []string{"info"} }

func (cmd *Info) onlyOwner() bool { return false }

func (cmd *Info) usage() []string { return []string{""} }
