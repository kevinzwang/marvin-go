package commands

import (
	"os"
)

// Quit exits the bot.
type Quit struct{}

func (cmd *Quit) execute(ctx *Context, args []string) {
	ctx.send("Bye!")
	ctx.Session.Close()
	os.Exit(0)
}

func (cmd *Quit) description() string { return "says bye then quits the program" }

func (cmd *Quit) category() string { return "admin" }

func (cmd *Quit) numArgs() (int, int) { return 0, 0 }

func (cmd *Quit) names() []string { return []string{"quit", "exit"} }

func (cmd *Quit) onlyOwner() bool { return true }

func (cmd *Quit) usage() []string { return []string{""} }
