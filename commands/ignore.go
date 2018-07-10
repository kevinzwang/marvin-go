package commands

// Ignore is a command to disable actions for a message.
type Ignore struct{}

func (cmd *Ignore) execute(ctx *Context, args []string) {}

func (cmd *Ignore) description() string {
	return "Disables any further actions performed by Marvin for the message."
}

func (cmd *Ignore) category() string { return "actions" }

func (cmd *Ignore) numArgs() (int, int) { return -1, -1 }

func (cmd *Ignore) names() []string { return []string{"ignore", "i"} }

func (cmd *Ignore) onlyOwner() bool { return false }

func (cmd *Ignore) usage() []string {
	return []string{"all of the blood on the walls", "https://open.spotify.com/track/7GhIk7Il098yCjg4BQjzvb?si=1dzqM0rOSuiCRViPNix9Tg"}
}
