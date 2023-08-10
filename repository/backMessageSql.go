package repository

import (
	"discordbot/datasource/sqlsource"
	"discordbot/model/vo"

	"github.com/rs/zerolog/log"
)

type BackMessageSqlConnection struct {
	*sqlsource.Connection
}

func (db *BackMessageSqlConnection) Insert(backMessages []vo.BackMessageVo) bool {
	insertSql := `INSERT INTO backmessage(key, value, guildid) VALUES($1, $2, $3)`

	tx, _ := db.Begin()
	for _, v := range backMessages {
		_, err := tx.Exec(insertSql, v.Key, v.Value, v.GuildId)

		if err != nil {
			tx.Rollback()
			log.Err(err).Msg(`BackMessage Insert ERROR`)
			return false
		}
	}
	tx.Commit()

	return true
}

func (db *BackMessageSqlConnection) FindByKeyAndGuildId(key string, guildId string) []vo.BackMessageVo {
	selectSql := `SELECT * FROM backmessage WHERE key=:key AND GUILDID=:guildid`
	backMessage := vo.BackMessageVo{Key: key, GuildId: guildId}
	backMessageVoSlice := []vo.BackMessageVo{}

	nstmt, _ := db.PrepareNamed(selectSql)
	err := nstmt.Select(&backMessageVoSlice, backMessage)

	if err != nil {
		log.Err(err).Msg(`BackMessage FindByKey ERROR`)
	}

	return backMessageVoSlice
}

func (db *BackMessageSqlConnection) FindAll() []vo.BackMessageVo {
	selectSql := `SELECT * FROM backmessage`
	allValue := []vo.BackMessageVo{}

	db.Select(&allValue, selectSql)

	return allValue
}

func (db *BackMessageSqlConnection) FindByGuildId(guildId string) []vo.BackMessageVo {
	selectSql := `SELECT DISTINCT KEY FROM backmessage WHERE GUILDID=:guildid`
	arg := vo.BackMessageVo{GuildId: guildId}
	valuse := []vo.BackMessageVo{}

	stmt, err := db.PrepareNamed(selectSql)

	stmt.Select(&valuse, arg)

	if err != nil {
		log.Err(err).Msg(`find backmessage table error`)
	}

	return valuse
}

func (db *BackMessageSqlConnection) DeleteById(id string) bool {
	sql := `DELETE FROM backmessage WHERE id=:id`
	backMessage := vo.BackMessageVo{Id: id}
	_, err := db.NamedExec(sql, backMessage)

	return err == nil
}

func (db *BackMessageSqlConnection) DeleteByIdAndKeyAndGuildId(id string, key string, guildId string) bool {
	sql := `DELETE FROM backmessage WHERE ID=:id AND KEY=:key AND GUILDID=:guildid`
	arg := vo.BackMessageVo{Id: id, Key: key, GuildId: guildId}
	_, err := db.NamedExec(sql, arg)

	return err == nil
}

func (db *BackMessageSqlConnection) DeleteByKeyAndValue(key string, value string) bool {
	sql := `DELETE FROM backmessage WHERE key=:key AND value=:value`
	backMessage := vo.BackMessageVo{Key: key, Value: value}
	result, err := db.NamedExec(sql, backMessage)

	if err != nil {
		log.Err(err).Msgf(`Delete %v Error`, key)
	}

	row, err := result.RowsAffected()

	return err == nil && row > 0
}
