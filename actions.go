package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func neutralizeSpotifyLink(session *discordgo.Session, message *discordgo.MessageCreate) {
	if strings.Contains(message.Content, "open.spotify.com") {
		session.ChannelMessageDelete(message.ChannelID, message.ID)
		em := discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name:    message.Author.String(),
				IconURL: message.Author.AvatarURL(""),
			},
			Color:       0x1db954,
			Description: message.Content,
			Footer: &discordgo.MessageEmbedFooter{
				IconURL: "https://cdn.freebiesupply.com/logos/large/2x/spotify-2-logo-png-transparent.png",
				Text:    "Please escape your Spotify links with `<` and `>` because the Spotify embed crashes some clients.",
			},
		}
		session.ChannelMessageSendEmbed(message.ChannelID, &em)
	}
}
