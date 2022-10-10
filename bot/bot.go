package bot

import (
	"discordbot/service"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

var session *discordgo.Session

func ConnectDiscord() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + viper.GetString("discordbot.token"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	session = dg
	backMessageService = service.GetBackMessageService()

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

}

func CloseDiscord() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	session.Close()
}
