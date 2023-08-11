package app

import (
	"context"
	"discordbot/bot"
	"discordbot/datasource/sqlsource"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func Run() {
	defaultConfig()
	loadConfigFile()
	bot := bot.New("Bot " + viper.GetString("discordbot.token"))
	ctx, cancel := context.WithCancel(context.Background())
	wait := sync.WaitGroup{}

	if viper.GetString("datasource.postgres.uri") == "" {
		log.Fatal().Msg("No DB uri???")
	}

	wait.Add(1)
	go sqlsource.ConnectPostSql(ctx, &wait)

	bot.ConnectDiscord()
	cancel()
	wait.Wait()
}

func loadConfigFile() {
	// logrus.SetReportCaller(true)
	// log.SetFlags(log.Lshortfile)
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Read config error")
	}
}

func defaultConfig() {
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.With().Caller().Logger()
}
