package commands

// Invite replies with the invite for this bot
type Invite struct{}

func (cmd *Invite) execute(ctx *Context, args []string) {
	usr, _ := ctx.Session.User("@me")
	ctx.send("https://discordapp.com/oauth2/authorize?client_id=" + usr.ID + "&scope=bot")
}

func (cmd *Invite) description() string { return "replies with the invite for this bot" }

func (cmd *Invite) category() string { return "misc" }

func (cmd *Invite) numArgs() (int, int) { return 0, 0 }

func (cmd *Invite) names() []string { return []string{"invite"} }

func (cmd *Invite) onlyOwner() bool { return false }

func (cmd *Invite) usage() []string { return []string{""} }
