package commands

import (
	"strconv"
)

// Ping returns "pong!"" when the ping command is run
type Ping struct{}

func (cmd *Ping) execute(ctx *Context, args []string) {
	msg, err := ctx.reply("pong!")
	if err != nil {
		return
	}
	t1, _ := ctx.Message.Timestamp.Parse()
	t2, _ := msg.Timestamp.Parse()
	elapsed := t2.Sub(t1)
	ctx.Session.ChannelMessageEdit(msg.ChannelID, msg.ID, msg.Content+"\nLatency: "+strconv.FormatFloat(elapsed.Seconds()*1000, 'f', 0, 64)+"ms")
}

func (cmd *Ping) description() string { return "replies \"pong!\" when pinged" }

func (cmd *Ping) category() string { return "misc" }

func (cmd *Ping) numArgs() (int, int) { return 0, 0 }

func (cmd *Ping) names() []string { return []string{"ping"} }

func (cmd *Ping) onlyOwner() bool { return false }

func (cmd *Ping) usage() []string { return []string{""} }
