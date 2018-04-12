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
	if err != nil {
		errors.Fatal.Fatal("Could not read config.yaml")
	}

	contents := string(b)
	cfg, err := config.ParseYaml(contents)
	if err != nil {
		errors.Fatal.Fatal("Could not parse config.yaml")
	}

	token, err = cfg.String("token")
	if err != nil {
		errors.Fatal.Fatal("Could not find token in config.yaml")
	}

	prefix, err = cfg.String("prefix")
	if err != nil {
		errors.Fatal.Fatal("Could not find prefix in config.yaml")
	}

	ownerID, err = cfg.String("owner")
	if err != nil {
		errors.Warning.Fatal("Could not find owner in config.yaml")
	}

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		errors.Fatal.Fatal("Could not create Discord session")
	}

	discord.AddHandler(messageCreate)
	discord.AddHandler(connect)

	commands.Init(&prefix, &ownerID)

	commands.AddCommands(&commands.Ping{}, &commands.Quit{}, &commands.Help{}, &commands.Info{}, &commands.Cat{})

	err = discord.Open()
	if err != nil {
		errors.Fatal.Fatal("Could not open connection to Discord")
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
}
