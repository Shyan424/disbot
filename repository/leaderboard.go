package repository

import (
	"context"
	"discordbot/datasource"

	"github.com/redis/go-redis/v9"
)

type LeaderboardRedisRepository interface {
	AddScore(key string, member string)
	AddScores(key string, member string, points int64)
	GetTopBoard(key string, top int64) []string
}

type LeaderboardRedis struct {
	*datasource.RedisClient
}

func GetLeaderboardRepository() LeaderboardRedisRepository {
	return &LeaderboardRedis{datasource.GetRedisClient()}
}

func (l *LeaderboardRedis) AddScore(key string, member string) {
	l.ZIncrBy(context.Background(), key, 1, member)
}

func (l *LeaderboardRedis) AddScores(key string, member string, points int64) {
	l.ZIncrBy(context.Background(), key, float64(points), member)
}

// ZRANGE {key} +inf (0 BYSCORE REV LIMIT 0 {top}
// score rev then (infinite > score > 0) get index (0 ~ top-1)
func (l *LeaderboardRedis) GetTopBoard(key string, top int64) []string {
	// go-redis 在 rev 時會自動把 start stop 互換
	arg := redis.ZRangeArgs{
		Key:     key,
		Start:   "1",
		Stop:    "+inf",
		ByScore: true,
		Rev:     true,
		Offset:  0,
		Count:   top,
	}

	return l.ZRangeArgs(context.Background(), arg).Val()
}
