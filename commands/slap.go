package commands

// Slap gives a random slap pic
type Slap struct{}

func (cmd *Slap) execute(ctx *Context, args []string) {
	ctx.send(nekoslife("slap"))
}

func (cmd *Slap) description() string { return "gives a random slap pic" }

func (cmd *Slap) category() string { return "weeb" }

func (cmd *Slap) numArgs() (int, int) { return 0, 0 }

func (cmd *Slap) names() []string { return []string{"slap"} }

func (cmd *Slap) onlyOwner() bool { return false }

func (cmd *Slap) usage() []string { return []string{""} }
