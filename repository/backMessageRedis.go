package repository

import (
	"context"
	"discordbot/datasource"
	"discordbot/model/vo"
)

type BackMessageRedisRepositry interface {
	DeleteByKeyAndValue(guildId string, key string, value string) error
	FindByKeyAndGuildId(guildId string, key string) ([]string, error)
	FindRandomValue(guildId string, key string) string
	InsertByGuildIdAndKey(guildId string, key string, values ...string)
	Insert(backMessages []vo.BackMessageVo)
}

type BackMessageRedis struct {
	*datasource.RedisClient
}

func GetBackMessageRedis() BackMessageRedisRepositry {
	return &BackMessageRedis{datasource.GetRedisClient()}
}

func (r *BackMessageRedis) DeleteByKeyAndValue(guildId string, key string, value string) error {
	return r.SRem(context.Background(), guildIdAndKeyToSetsKey(guildId, key), value).Err()
}

func (r *BackMessageRedis) FindByKeyAndGuildId(guildId string, key string) ([]string, error) {
	vos := []string{}
	err := r.SMembers(context.Background(), guildIdAndKeyToSetsKey(guildId, key)).ScanSlice(vos)

	return vos, err
}

func (r *BackMessageRedis) FindRandomValue(guildId string, key string) string {
	return r.SRandMember(context.Background(), guildIdAndKeyToSetsKey(guildId, key)).Val()
}

func (r *BackMessageRedis) InsertByGuildIdAndKey(guildId string, key string, values ...string) {
	r.SAdd(context.Background(), guildIdAndKeyToSetsKey(guildId, key), values)
}

func (r *BackMessageRedis) Insert(backMessages []vo.BackMessageVo) {
	backMessageKeyValues := splitByGuildIdAndKey(backMessages)
	for guildId, vs := range backMessageKeyValues {
		for key, vos := range vs {
			r.SAdd(context.Background(), guildIdAndKeyToSetsKey(guildId, key), vos)
		}
	}
}

func splitByGuildIdAndKey(backMessages []vo.BackMessageVo) map[string]map[string][]string {
	guildV := map[string]map[string][]string{}
	for _, v := range backMessages {
		if keyV, ok := guildV[v.GuildId]; ok {
			guildV[v.GuildId][v.Key] = append(keyV[v.Key], v.Value)
		} else {
			guildV[v.GuildId] = map[string][]string{v.Key: {v.Value}}
		}
	}

	return guildV
}

func guildIdAndKeyToSetsKey(guildId string, key string) string {
	return guildId + ":" + key
}
