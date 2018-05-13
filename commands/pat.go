package commands

// Pat gives a random pat pic
type Pat struct{}

func (cmd *Pat) execute(ctx *Context, args []string) {
	ctx.send(nekoslife("pat"))
}

func (cmd *Pat) description() string { return "gives a random pat pic" }

func (cmd *Pat) category() string { return "weeb" }

func (cmd *Pat) numArgs() (int, int) { return 0, 0 }

func (cmd *Pat) names() []string { return []string{"pat"} }

func (cmd *Pat) onlyOwner() bool { return false }

func (cmd *Pat) usage() []string { return []string{""} }
