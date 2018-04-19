package commands

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
)

// Info gives information about the bot
type Info struct{}

func (cmd *Info) execute(ctx *Context, args []string) {
	guilds, _ := ctx.Session.UserGuilds(100, "", "")

	embed := discordgo.MessageEmbed{
		Title:       "About Marvin",
		Color:       0xffffff,
		Description: "Marvin is an all-purpose Discord bot written in Golang using the discordgo library.",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{Name: "Creator", Value: "Random17#0608", Inline: true},
			&discordgo.MessageEmbedField{Name: "Servers", Value: strconv.Itoa(len(guilds)), Inline: true},
			&discordgo.MessageEmbedField{Name: "Dev Server", Value: "http://discord.gg/yvqTyGt", Inline: true},
			&discordgo.MessageEmbedField{Name: "Github Repo", Value: "http://github.com/kevinzwang/marvin-go", Inline: true},
		},
	}

	ctx.sendEmbed(&embed)
}

func (cmd *Info) description() string { return "gives information about this bot" }

func (cmd *Info) category() string { return "misc" }

func (cmd *Info) numArgs() (int, int) { return 0, 0 }

func (cmd *Info) names() []string { return []string{"info", "about"} }

func (cmd *Info) onlyOwner() bool { return false }

func (cmd *Info) usage() []string { return []string{""} }
