package datasource

import (
	"context"
	"discordbot/model/config"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func initRedis(conf config.Config) RedisClient {
	var redisClient RedisClient
	option, err := redis.ParseURL(conf.Datasource.Redis.Uri)
	if err != nil {
		log.Fatal().Err(err).Msg("Redis parse URL error")
	}

	redisClient.Client = redis.NewClient(option)

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatal().Err(err).Msg("Redis connect error")
	}

	log.Info().Msg("Redis connected")

	return redisClient
}
