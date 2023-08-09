package service

import (
	"discordbot/model/vo"
	"discordbot/repository"
	"math/rand"
	"time"
)

type BackMessageService interface {
	InsertMessage(key string, value string, guildId string) bool
	GetAllValue() []vo.BackMessageVo
	GetAllValueByKeyAndGuild(key string, guildId string) []vo.BackMessageVo
	GetAllKey() map[string][]string
	GetAllKeyByGuildId(guildId string) []string
	GetRandomValue(key string, guildId string) string
	DeleteMessageById(id string) bool
	DeleteMessageByIdAndKeyAndGuildId(id string, key string, guildId string) bool
	DeleteMessage(key string, value string) bool
}

type BackMessageConnection struct {
	repository.BackMessageRepository
}

func GetBackMessageService() BackMessageService {
	modelConn := repository.GetBackMessageRepository()
	return &BackMessageConnection{modelConn}
}

func (conn *BackMessageConnection) InsertMessage(key string, value string, guildId string) bool {
	return conn.Insert([]vo.BackMessageVo{{Key: key, Value: value, GuildId: guildId}})
}

func (conn *BackMessageConnection) GetAllValue() []vo.BackMessageVo {
	return conn.FindAll()
}

func (conn *BackMessageConnection) GetAllValueByKeyAndGuild(key string, guildId string) []vo.BackMessageVo {
	return conn.FindByKeyAndGuildId(key, guildId)
}

func (conn *BackMessageConnection) GetAllKey() map[string][]string {
	allBackMessageKey := map[string][]string{}
	allKey := conn.FindAll()

	for _, v := range allKey {
		allBackMessageKey[v.GuildId] = append(allBackMessageKey[v.GuildId], v.Key)
	}

	return allBackMessageKey
}

func (conn *BackMessageConnection) GetAllKeyByGuildId(guildId string) []string {
	backMessages := conn.FindByGuildId(guildId)
	keys := make([]string, len(backMessages))

	for i, v := range backMessages {
		keys[i] = v.Key
	}

	return keys
}

func (conn *BackMessageConnection) GetRandomValue(key string, guildId string) string {
	value := ""
	bm := conn.FindByKeyAndGuildId(key, guildId)
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

func (conn *BackMessageConnection) DeleteMessageById(id string) bool {
	return conn.DeleteById(id)
}

func (conn *BackMessageConnection) DeleteMessageByIdAndKeyAndGuildId(id string, key string, guildId string) bool {
	return conn.DeleteByIdAndKeyAndGuildId(id, key, guildId)
}

func (conn *BackMessageConnection) DeleteMessage(key string, value string) bool {
	return conn.DeleteByKeyAndValue(key, value)
}
