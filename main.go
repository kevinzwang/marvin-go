package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"./commands"
	"./errors"
	"github.com/bwmarrin/discordgo"
	"github.com/olebedev/config"
)

var token string
var prefix string
var ownerID string

func main() {
	b, err := ioutil.ReadFile("config.yaml")
	errors.Fatal(err, "Error reading config.yaml")

	contents := string(b)
	cfg, err := config.ParseYaml(contents)
	errors.Fatal(err, "Error parsing config.yaml")

	token, err = cfg.String("token")
	errors.Fatal(err, "Error finding token in config.yaml")

	prefix, err = cfg.String("prefix")
	errors.Fatal(err, "Error finding prefix in config.yaml")

	ownerID, err = cfg.String("owner")
	errors.Fatal(err, "Error finding owner in config.yaml")

	discord, err := discordgo.New("Bot " + token)
	errors.Fatal(err, "Error creating Discord session")

	discord.AddHandler(messageCreate)
	discord.AddHandler(connect)

	commands.Init(&prefix, &ownerID)

	commands.AddCommands(&commands.Ping{}, &commands.Quit{}, &commands.Help{})

	err = discord.Open()
	errors.Fatal(err, "Error opening connection")

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()

}

func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	commands.Handle(message, session)
}

func sendMsg(session *discordgo.Session, channelID string, message string) {
	_, err := session.ChannelMessageSend(channelID, message)
	errors.Warning(err, "Error sending Discord message")
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
}
