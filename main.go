package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/olebedev/config"
)

var token string

func init() {
	b, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("error reading config.yaml,", err)
		return
	}
	contents := string(b)
	cfg, err := config.ParseYaml(contents)
	if err != nil {
		fmt.Println("error parsing config.yaml,", err)
		return
	}
	token, err = cfg.String("token")
	if err != nil {
		fmt.Println("error finding token in config.yaml,", err)
		return
	}
}

func main() {

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	discord.AddHandler(messageCreate)

	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

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
	if err != nil {
		fmt.Println("error sending Discord message,", err)
	}
}
