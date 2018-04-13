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
	errors.Fatal(err, "Could not read config.yaml", nil)

	contents := string(b)
	cfg, err := config.ParseYaml(contents)
	errors.Fatal(err, "Could not parse config.yaml", nil)

	token, err = cfg.String("token")
	errors.Fatal(err, "Could not find token in config.yaml", nil)

	prefix, err = cfg.String("prefix")
	errors.Fatal(err, "Could not find prefix in config.yaml", nil)

	ownerID, err = cfg.String("owner")
	errors.Warning(err, "Could not find owner in config.yaml")

	discord, err := discordgo.New("Bot " + token)
	errors.Fatal(err, "Could not create Discord session", nil)

	discord.AddHandler(messageCreate)
	discord.AddHandler(connect)

	commands.Init(&prefix, &ownerID)

	commands.AddCommands(&commands.Ping{}, &commands.Quit{}, &commands.Help{}, &commands.Info{},
		&commands.Cat{}, &commands.Anime{})

	err = discord.Open()
	errors.Fatal(err, "Could not open connection to Discord", discord)

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
