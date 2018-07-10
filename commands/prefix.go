package commands

import "../yamlutils"

// Prefix allows you to view and modify the prefix
type Prefix struct{}

func (cmd *Prefix) execute(ctx *Context, args []string) {
	if ctx.Guild == nil {
		ctx.reply("You don't need prefixes in DMs!")
		return
	}

	if len(args) == 0 {
		prefix := yamlutils.GetPrefix(ctx.Message.ChannelID)
		ctx.reply("Current server prefix: `" + prefix + "`")
		return
	}

	if args[0] == "-g" && len(args) == 1 {
		prefix := yamlutils.GetPrefix("global")
		ctx.reply("Current global prefix: `" + prefix + "`")
		return
	}
	ctx.wrongNumArgs("prefix")
}

func (cmd *Prefix) description() string {
	return "replies with the server prefix, or the global prefix if called with `-g`."
}

func (cmd *Prefix) category() string { return "misc" }

func (cmd *Prefix) numArgs() (int, int) { return 0, 1 }

func (cmd *Prefix) names() []string { return []string{"prefix"} }

func (cmd *Prefix) onlyOwner() bool { return false }

func (cmd *Prefix) usage() []string { return []string{"", "-g"} }
