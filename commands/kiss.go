package commands

// Kiss gives a random kiss pic
type Kiss struct{}

func (cmd *Kiss) execute(ctx *Context, args []string) {
	ctx.send(nekoslife("kiss"))
}

func (cmd *Kiss) description() string { return "gives a random kiss pic" }

func (cmd *Kiss) category() string { return "weeb" }

func (cmd *Kiss) numArgs() (int, int) { return 0, 0 }

func (cmd *Kiss) names() []string { return []string{"kiss"} }

func (cmd *Kiss) onlyOwner() bool { return false }

func (cmd *Kiss) usage() []string { return []string{""} }
