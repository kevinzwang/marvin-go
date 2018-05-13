package commands

// Hug gives a random hug pic
type Hug struct{}

func (cmd *Hug) execute(ctx *Context, args []string) {
	ctx.send(nekoslife("hug"))
}

func (cmd *Hug) description() string { return "gives a random hug pic" }

func (cmd *Hug) category() string { return "weeb" }

func (cmd *Hug) numArgs() (int, int) { return 0, 0 }

func (cmd *Hug) names() []string { return []string{"hug"} }

func (cmd *Hug) onlyOwner() bool { return false }

func (cmd *Hug) usage() []string { return []string{""} }
