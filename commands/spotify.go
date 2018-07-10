package commands

import (
	"../yamlutils"
)

// Spotify is a command wrapper for the neutralizeSpotifyLink() action
type Spotify struct{}

func (cmd *Spotify) execute(ctx *Context, args []string) {
	ctx.send("This command doesn't do anything itself, it's just a help page for the Spotify action." +
		"\n\nSome Discord clients crash when the Spotify widget is displayed, so this action wraps the links in an embed to prevent that from happening." +
		"\n\nIf you would like to disable this action, prepend `" + yamlutils.GetPrefix(ctx.Guild.ID) + "ignore`")
}

func (cmd *Spotify) description() string {

	return "This command doesn't do anything itself, it's just a help page for the Spotify action." +
		"\n\nSome Discord clients crash when the Spotify widget is displayed, so this action wraps the links in an embed to prevent that from happening." +
		"\n\nIf you would like to disable this action, prepend `" + yamlutils.GetPrefix("global") + "ignore`"
}

func (cmd *Spotify) category() string { return "actions" }

func (cmd *Spotify) numArgs() (int, int) { return -1, -1 }

func (cmd *Spotify) names() []string { return []string{"spotify"} }

func (cmd *Spotify) onlyOwner() bool { return false }

func (cmd *Spotify) usage() []string { return []string{} }
