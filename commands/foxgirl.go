package commands

// Foxgirl gives a random foxgirl pic
type Foxgirl struct{}

func (cmd *Foxgirl) execute(ctx *Context, args []string) {
	ctx.send(nekoslife("fox_girl"))
}

func (cmd *Foxgirl) description() string { return "gives a random foxgirl pic" }

func (cmd *Foxgirl) category() string { return "weeb" }

func (cmd *Foxgirl) numArgs() (int, int) { return 0, 0 }

func (cmd *Foxgirl) names() []string { return []string{"foxgirl"} }

func (cmd *Foxgirl) onlyOwner() bool { return false }

func (cmd *Foxgirl) usage() []string { return []string{""} }
