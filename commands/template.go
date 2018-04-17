package commands

// Template is a command template
type Template struct{}

func (cmd *Template) execute(ctx *Context, args []string) {}

func (cmd *Template) description() string { return "" }

func (cmd *Template) category() string { return "misc" }

func (cmd *Template) numArgs() (int, int) { return -1, -1 }

func (cmd *Template) names() []string { return []string{"temp"} }

func (cmd *Template) onlyOwner() bool { return false }

func (cmd *Template) usage() []string { return []string{"", "<stuff>"} }
