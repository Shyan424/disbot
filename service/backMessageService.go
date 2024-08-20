package service

import (
	"discordbot/model/vo"
	"discordbot/repository"
	"math/rand"
	"time"

	"github.com/rs/zerolog/log"
)

type BackMessageService interface {
	RefreshAllToRedis()
	InsertMessage(key string, value string, guildId string) bool
	GetAllValueByKeyAndGuild(key string, guildId string) []string
	GetAllKeyByGuildId(guildId string) []string
	GetRandomValue(key string, guildId string) string
	DeleteMessage(guildId string, key string, value string) bool
}

type BackMessageConnection struct {
	s repository.BackMessageRepository
	r repository.BackMessageRedisRepositry
}

func GetBackMessageService() BackMessageService {
	modelConn := repository.GetBackMessageRepository()
	rdb := repository.GetBackMessageRedis()

	return &BackMessageConnection{s: modelConn, r: rdb}
}

func (conn *BackMessageConnection) RefreshAllToRedis() {
	// datasource.GetRedisClient().FlushDB(context.Background())
	conn.r.Insert(conn.s.FindAll())
}

func (conn *BackMessageConnection) InsertMessage(key string, value string, guildId string) bool {
	conn.r.InsertByGuildIdAndKey(guildId, key, value)
	return conn.s.Insert([]vo.BackMessageVo{{Key: key, Value: value, GuildId: guildId}}) == nil
}

func (conn *BackMessageConnection) GetAllValueByKeyAndGuild(key string, guildId string) []string {
	values, _ := conn.r.FindByKeyAndGuildId(key, guildId)
	if values != nil {
		return values
	}

	backMessages, err := conn.s.FindByKeyAndGuildId(key, guildId)
	if err != nil {
		log.Err(err).Msgf("FindByKeyAndGuildId %s Error", key)
	}
	for _, v := range backMessages {
		values = append(values, v.Value)
	}

	conn.r.InsertByGuildIdAndKey(guildId, key, values...)

	return values
}

func (conn *BackMessageConnection) GetAllKeyByGuildId(guildId string) []string {
	allKey := "AllKey:" + guildId
	keys, _ := conn.r.FindByKeyAndGuildId(guildId, allKey)
	if len(keys) != 0 {
		return keys
	}

	backMessages, err := conn.s.FindByGuildId(guildId)
	if err != nil {
		return keys
	}
	for _, v := range backMessages {
		keys = append(keys, v.Key)
	}
	conn.r.InsertByGuildIdAndKey(guildId, allKey, keys...)

	return keys
}

func (conn *BackMessageConnection) GetRandomValue(key string, guildId string) string {
	value := conn.r.FindRandomValue(guildId, key)
	if value != "" {
		return value
	}

	bm, _ := conn.s.FindByKeyAndGuildId(key, guildId)
	bmLen := len(bm)

	if bmLen != 0 {
		value = bm[random(bmLen)].Value
	}

	return value
}

func random(maxIndex int) int {
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	return r.Intn(maxIndex)
}

func (conn *BackMessageConnection) DeleteMessage(guildId string, key string, value string) bool {
	conn.r.DeleteByKeyAndValue(guildId, key, value)
	err := conn.s.DeleteByKeyAndValue(guildId, key, value)
	if err != nil {
		log.Err(err).Msgf("DeleteByKeyAndValue(%s, %s) Error", key, value)
	}

	return err == nil
}
