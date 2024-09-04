package service

import (
	"discordbot/repository"
	"strings"
	"time"
)

type LeaderboardService interface {
	AddScore(guildId string, member string)
	GetTopLeader(guildId string, top int64) []string
}

type Leaderboard struct {
	r repository.LeaderboardRedisRepository
}

func GetLeaderboardService() LeaderboardService {
	return &Leaderboard{repository.GetLeaderboardRepository()}
}

const PREFIX string = "leaderboard"

func (l *Leaderboard) AddScore(guildId string, member string) {
	newkey := boaedKey(guildId, time.Now())
	l.r.AddScore(newkey, member)
}

func (l *Leaderboard) GetTopLeader(guildId string, top int64) []string {
	newkey := boaedKey(guildId, time.Now())

	return l.r.GetTopBoard(newkey, top)
}

func boaedKey(guildId string, date time.Time) string {
	sb := []string{PREFIX, date.Format("2006"), guildId}

	return strings.Join(sb, ":")
}
