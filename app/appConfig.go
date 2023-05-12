package app

import (
	"context"
	"discordbot/bot"
	"discordbot/datasource/mongosource"
	"log"
	"sync"

	"github.com/spf13/viper"
)

func Run() {
	loadConfig()
	bot := getBot()
	ctx, cancel := context.WithCancel(context.Background())
	wait := sync.WaitGroup{}

	bot.ConnectDiscord()
	if viper.GetString("datasource.mongodb.uri") != "" {
		wait.Add(1)
		go mongosource.RunWithMongo(ctx, &wait)
	}

	bot.WaitForClose()
	cancel()
	wait.Wait()
}

func loadConfig() {
	log.SetFlags(log.Lshortfile)
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func getBot() *bot.Discordbot {
	return bot.New("Bot " + viper.GetString("discordbot.token"))
}
