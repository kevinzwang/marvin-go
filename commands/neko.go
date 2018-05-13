package commands

// Neko gives a random catgirl pic
type Neko struct{}

func (cmd *Neko) execute(ctx *Context, args []string) {
	ctx.send(nekoslife("neko"))
}

func (cmd *Neko) description() string { return "gives a random catgirl pic" }

func (cmd *Neko) category() string { return "weeb" }

func (cmd *Neko) numArgs() (int, int) { return 0, 0 }

func (cmd *Neko) names() []string { return []string{"neko", "catgirl"} }

func (cmd *Neko) onlyOwner() bool { return false }

func (cmd *Neko) usage() []string { return []string{""} }
