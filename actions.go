package main

import (
	"strings"

	"./yamlutils"
	"github.com/bwmarrin/discordgo"
)

func neutralizeSpotifyLink(session *discordgo.Session, message *discordgo.MessageCreate) {
	if strings.Contains(message.Content, "open.spotify.com") {
		session.ChannelMessageDelete(message.ChannelID, message.ID)

		prefix := yamlutils.GetPrefix(message.ChannelID)

		em := discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name:    message.Author.String() + " said:",
				IconURL: message.Author.AvatarURL(""),
			},
			Color:       0x1db954,
			Description: message.Content,
			Footer: &discordgo.MessageEmbedFooter{
				IconURL: "https://cdn.freebiesupply.com/logos/large/2x/spotify-2-logo-png-transparent.png",
				Text:    "`" + prefix + "help spotify` for more info about this action.",
			},
		}
		session.ChannelMessageSendEmbed(message.ChannelID, &em)
	}
}
