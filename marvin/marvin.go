package marvin

import "github.com/bwmarrin/discordgo"

var session *discordgo.Session

func Init(s *discordgo.Session) {
	session = s
}

func Session() *discordgo.Session {
	return session
}
