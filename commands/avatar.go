package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Avatar gives the avatar image for the caller or the user specified
type Avatar struct{}

func (cmd *Avatar) execute(ctx *Context, args []string) {
	usrID, ok := cmd.getID(ctx, args)

	if ok {
		usr, _ := ctx.Session.User(usrID)
		embed := discordgo.MessageEmbed{
			Title: "Avatar for " + usr.Username + "#" + usr.Discriminator,
			Image: &discordgo.MessageEmbedImage{
				URL: usr.AvatarURL("256"),
			},
		}

		ctx.sendEmbed(&embed)
	} else {
		ctx.reply("No such user " + usrID)
	}
}

func (cmd *Avatar) matchAny(content string, usr *discordgo.User, nick string) (string, bool) {
	if content == usr.Username || content == usr.String() || content == nick || content == usr.ID {
		return usr.ID, true
	}
	return "", false
}

func (cmd *Avatar) getID(ctx *Context, args []string) (string, bool) {
	if len(args) == 0 {
		return ctx.Message.Author.ID, true
	}
	// check if it's a mention
	mentions := ctx.Message.Mentions
	if len(mentions) == 1 {
		return mentions[0].ID, true
	}

	content := strings.TrimSpace(ctx.Content)
	if ctx.Guild == nil {
		self, _ := ctx.Session.User("@me")
		if id, match := cmd.matchAny(content, self, ""); match {
			return id, true
		}
		if id, match := cmd.matchAny(content, ctx.Author, ""); match {
			return id, true
		}
	} else {
		// check if it's a username, nick, or ID
		users := ctx.Guild.Members
		for _, u := range users {
			if id, match := cmd.matchAny(content, u.User, u.Nick); match {
				return id, true
			}
		}
	}

	return content, false
}

func (cmd *Avatar) description() string {
	return "Gives the avatar of the caller or the user specified."
}

func (cmd *Avatar) category() string { return "misc" }

func (cmd *Avatar) numArgs() (int, int) { return 0, -1 }

func (cmd *Avatar) names() []string { return []string{"avatar"} }

func (cmd *Avatar) onlyOwner() bool { return false }

func (cmd *Avatar) usage() []string { return []string{"", "<username, server nick, or ID>"} }
