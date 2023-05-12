package bot

import (
	"discordbot/service"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type Discordbot struct {
	session *discordgo.Session
}

func New(token string) *Discordbot {
	bot := &Discordbot{}
	dg, err := discordgo.New(token)
	if err != nil {
		log.Fatal("error creating Discord session,", err)
	}

	bot.session = dg

	return bot
}

func (d *Discordbot) ConnectDiscord() {
	// Create a new Discord session using the provided bot token.
	backMessageService = service.GetBackMessageService()

	// Register the messageCreate func as a callback for MessageCreate events.
	d.session.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	d.session.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err := d.session.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

}

func (d *Discordbot) WaitForClose() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	d.session.Close()
}
