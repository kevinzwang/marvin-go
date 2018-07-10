package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"./commands"
	"./logger"
	"./marvin"
	"./yamlutils"
	"github.com/bwmarrin/discordgo"
)

var token string

var reconnect = false

func init() {
	logger.Info("SYSTEM START")
	readConfigs()
}

func main() {
	discord, err := discordgo.New("Bot " + token)
	logger.Fatal(err, "Could not create Discord session")

	marvin.Init(discord)

	discord.AddHandler(messageCreate)
	discord.AddHandler(connect)

	commands.AddCommands(
		&commands.Ping{},
		&commands.Quit{},
		&commands.Help{},
		&commands.Info{},
		&commands.Cat{},
		&commands.Anime{},
		&commands.Invite{},
		&commands.Prefix{},
		&commands.SetPrefix{},
		&commands.Servers{},
		&commands.Avatar{},
		&commands.If{},
		&commands.Rules{},
		&commands.Neko{},
		&commands.Foxgirl{},
		&commands.Hug{},
		&commands.Pat{},
		&commands.Slap{},
		&commands.Spotify{},
		&commands.Ignore{},
	)

	err = discord.Open()
	logger.Fatal(err, "Could not open connection to Discord")

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	self, _ := discord.User("@me")
	discord.UpdateStatus(0, "@"+self.Username+"#"+self.Discriminator+"  help")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()

}

func readConfigs() {
	token = yamlutils.GetToken()
	yamlutils.GetOwnerID()
}

func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	ignore := commands.Handle(message, session)

	currChan, _ := session.Channel(message.ChannelID)
	if !ignore && currChan.Type != discordgo.ChannelTypeDM {
		neutralizeSpotifyLink(session, message)
	}
}

func connect(session *discordgo.Session, _ *discordgo.Connect) {
	if !reconnect {
		fmt.Println("=============")
		usr, _ := session.User("@me")
		fmt.Printf("Name: %v#%v\n", usr.Username, usr.Discriminator)
		prefix := yamlutils.GetPrefix("global")
		if prefix == session.State.User.Mention()+" " {
			fmt.Printf("Global Prefix: none\n")
		} else {
			fmt.Printf("Global Prefix: \"%v\"\n", prefix)
		}
		guilds, _ := session.UserGuilds(100, "", "")
		fmt.Printf("Servers: %v\n", len(guilds))
		fmt.Println("=============")
		logger.Info("CONNECTED TO DISCORD")
		reconnect = true
	} else {
		logger.Info("RECONNECTEDCONNECTED TO DISCORD")
	}

}
