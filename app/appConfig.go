package app

import (
	"context"
	"discordbot/bot"
	"discordbot/datasource"
	"discordbot/model/config"
	"discordbot/service"
	"flag"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func Run() {
	defaultConfig()
	conf := loadConfigFile()
	checkNecessaryConfig(conf)
	bot := bot.New("Bot " + conf.Discordbot.Token)
	ctx, cancel := context.WithCancel(context.Background())
	wait := sync.WaitGroup{}

	datasource.ConnectDbs(ctx, &wait, conf)
	service.GetBackMessageService().RefreshAllToRedis()

	bot.ConnectDiscord()
	cancel()
	wait.Wait()
}

func loadFlag() config.Flag {
	flags := config.Flag{}
	flag.StringVar(&flags.Config, "config", "./config.yaml", "bot config")
	flag.Parse()

	return flags
}

func loadConfigFile() config.Config {
	flags := loadFlag()
	viper.SetConfigFile(flags.Config)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Read config error")
	}

	var c config.Config
	viper.Unmarshal(&c)

	return c
}

func checkNecessaryConfig(conf config.Config) {
	if conf.Datasource.Postgres.Uri == "" {
		log.Fatal().Msg("No DB uri???")
	}
	if conf.Datasource.Redis.Uri == "" {
		log.Fatal().Msg("No Redis uri???")
	}
}

func defaultConfig() {
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.With().Caller().Logger()
}
