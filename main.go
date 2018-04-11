package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/kevinzwang/marvin-go/errhandler"
	"github.com/olebedev/config"
)

var token string

func main() {
	b, err := ioutil.ReadFile("config.yaml")
	errhandler.HandleFatal(err, "Error reading config.yaml")

	contents := string(b)
	cfg, err := config.ParseYaml(contents)
	errhandler.HandleFatal(err, "Error parsing config.yaml")

	token, err = cfg.String("token")
	errhandler.HandleFatal(err, "Error finding token in config.yaml")

	discord, err := discordgo.New("Bot " + token)
	errhandler.HandleFatal(err, "Error creating Discord session")

	discord.AddHandler(messageCreate)
	discord.AddHandler(connect)

	err = discord.Open()
	errhandler.HandleFatal(err, "Error opening connection")

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

	if message.Content == "ping" {
		sendMsg(session, message.ChannelID, "pong")
	}
}

func sendMsg(session *discordgo.Session, channelID string, message string) {
	_, err := session.ChannelMessageSend(channelID, message)
	errhandler.Handle(err, "Error sending Discord message")
}

func connect(session *discordgo.Session, _ *discordgo.Connect) {
	fmt.Println("=============")
	// fmt.Println("Name: " + session.User("@me"))
}
