package commands

import "../yamlutils"

// SetPrefix sets the server or global prefix
type SetPrefix struct{}

func (cmd *SetPrefix) execute(ctx *Context, args []string) {
	if ctx.Guild == nil {
		ctx.reply("You don't need prefixes in DMs!")
		return
	}

	if args[0] == "-g" {
		if len(args) != 2 {
			ctx.wrongNumArgs("setprefix")
			return
		}
		prefix := args[1]

		ok := yamlutils.SetPrefix("global", prefix)
		if ok {
			ctx.send("Success in setting global prefix to `" + prefix + "`")
		}

	} else {
		if len(args) != 1 {
			ctx.wrongNumArgs("setprefix")
			return
		}
		server := ctx.Guild.ID
		prefix := args[0]

		ok := yamlutils.SetPrefix(server, prefix)
		if ok {
			ctx.send("Success in setting server prefix to `" + prefix + "`")
		}
	}
}

func (cmd *SetPrefix) description() string {
	return "Sets the server or global (with `-g` flag) prefix."
}

func (cmd *SetPrefix) category() string { return "admin" }

func (cmd *SetPrefix) numArgs() (int, int) { return 1, 2 }

func (cmd *SetPrefix) names() []string { return []string{"setprefix", "prefixset"} }

func (cmd *SetPrefix) onlyOwner() bool { return true }

func (cmd *SetPrefix) usage() []string { return []string{"<new prefix>", "-g <new prefix>"} }
