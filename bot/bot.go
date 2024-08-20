package bot

import (
	slashcommand "discordbot/bot/slashCommand"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

type Discordbot struct {
	Session *discordgo.Session
}

func New(token string) *Discordbot {
	bot := &Discordbot{}
	dg, err := discordgo.New(token)
	if err != nil {
		log.Fatal().Err(err).Msg("error creating Discord session")
	}

	bot.Session = dg

	return bot
}

func (d *Discordbot) ConnectDiscord() {
	// Register the messageCreate func as a callback for MessageCreate events.
	d.Session.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	d.Session.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err := d.Session.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	slashcommand.AddSlashCommand(d.Session)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	slashcommand.DeleteSlashCommand(d.Session)

	// Cleanly close down the Discord session.
	d.Session.Close()
}
