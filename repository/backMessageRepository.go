package repository

import (
	"discordbot/datasource/sqlsource"
	"discordbot/model/vo"
)

type BackMessageRepository interface {
	Insert(backMessages []vo.BackMessageVo) error
	FindByKeyAndGuildId(key string, guildId string) ([]vo.BackMessageVo, error)
	FindAll() []vo.BackMessageVo
	FindByGuildId(guildId string) ([]vo.BackMessageVo, error)
	DeleteById(id string) error
	DeleteByIdAndKeyAndGuildId(id string, key string, guildId string) error
	DeleteByKeyAndValue(key string, value string) error
}

func GetBackMessageRepository() BackMessageRepository {
	var connection BackMessageSqlConnection
	connection.Connection = sqlsource.GetDatasource()

	return &connection
}
