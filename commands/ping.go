package commands

// Ping returns "pong!"" when the ping command is run
type Ping struct{}

func (cmd *Ping) execute(ctx *Context, args []string) { ctx.reply("pong!") }

func (cmd *Ping) description() string { return "replies \"pong!\" when pinged" }

func (cmd *Ping) category() string { return "misc" }

func (cmd *Ping) numArgs() (int, int) { return 0, 0 }

func (cmd *Ping) names() []string { return []string{"ping"} }

func (cmd *Ping) onlyOwner() bool { return false }

func (cmd *Ping) usage() []string { return []string{""} }
