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
var prefix string
var ownerID string

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

	commands.Init(&prefix, &ownerID)

	commands.AddCommands(&commands.Ping{}, &commands.Quit{}, &commands.Help{}, &commands.Info{},
		&commands.Cat{}, &commands.Anime{})

	err = discord.Open()
	logger.Fatal(err, "Could not open connection to Discord")

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()

}

func readConfigs() {
	token = yamlutils.GetToken()
	prefix = yamlutils.GetPrefix()
	ownerID = yamlutils.GetOwnerID()
}

func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	commands.Handle(message, session)
}

func connect(session *discordgo.Session, _ *discordgo.Connect) {
	fmt.Println("=============")
	usr, _ := session.User("@me")
	fmt.Printf("Name: %v#%v\n", usr.Username, usr.Discriminator)
	owner, _ := session.User(ownerID)
	fmt.Printf("Owner: %v#%v\n", owner.Username, owner.Discriminator)
	guilds, _ := session.UserGuilds(100, "", "")
	fmt.Printf("Servers: %v\n", len(guilds))
	fmt.Printf("Prefix: %v\n", prefix)
	fmt.Println("=============")
	logger.Info("CONNECTED TO DISCORD")
}
