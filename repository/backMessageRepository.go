package repository

import (
	"discordbot/datasource/sqlsource"
	"discordbot/model/vo"
)

type BackMessageRepository interface {
	Insert(backMessages []vo.BackMessageVo) bool
	FindByKeyAndGuildId(key string, guildId string) []vo.BackMessageVo
	FindAll() []vo.BackMessageVo
	FindByGuildId(guildId string) []vo.BackMessageVo
	DeleteById(id string) bool
	DeleteByIdAndKeyAndGuildId(id string, key string, guildId string) bool
	DeleteByKeyAndValue(key string, value string) bool
}

func GetBackMessageRepository() BackMessageRepository {
	var connection BackMessageSqlConnection
	connection.Connection = sqlsource.GetDatasource()

	return &connection
}
