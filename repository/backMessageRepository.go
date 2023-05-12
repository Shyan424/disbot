package repository

import (
	"discordbot/datasource/mongosource"
	"discordbot/model/vo"
)

type BackMessageRepository interface {
	Insert(backMessages []vo.BackMessageVo) bool
	UpdateByKey(key string, values ...string) bool
	FindByKey(key string) *vo.BackMessageVo
	FindByKeyAndUpdate(key string, values ...string) (*vo.BackMessageVo, error)
}

func GetBackMessageRepository() BackMessageRepository {
	backMessage := mongosource.GetCollection("backmessage")
	return &BackMessageMongoConnection{backMessage}
}
