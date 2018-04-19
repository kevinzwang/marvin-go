package commands

import (
	"github.com/bwmarrin/discordgo"
)

// Servers lists the servers the bot is on
type Servers struct{}

func (cmd *Servers) execute(ctx *Context, args []string) {
	guilds := ctx.Session.State.Guilds
	for _, g := range guilds {
		name := g.Name
		blankInv := discordgo.Invite{MaxAge: 30, MaxUses: 1, Temporary: true}
		invite, _ := ctx.Session.ChannelInviteCreate(ctx.Message.ChannelID, blankInv)
		ctx.send(name + ": discord.gg/" + invite.Code)
	}
}

func (cmd *Servers) description() string { return "lists the servers that the bot is on" }

func (cmd *Servers) category() string { return "misc" }

func (cmd *Servers) numArgs() (int, int) { return 0, 0 }

func (cmd *Servers) names() []string { return []string{"servers", "guilds"} }

func (cmd *Servers) onlyOwner() bool { return true }

func (cmd *Servers) usage() []string { return []string{""} }
