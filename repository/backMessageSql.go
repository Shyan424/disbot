package repository

import (
	"discordbot/datasource/sqlsource"
	"discordbot/model/vo"

	"github.com/rs/zerolog/log"
)

type BackMessageSqlConnection struct {
	*sqlsource.Connection
}

func (db *BackMessageSqlConnection) Insert(backMessages []vo.BackMessageVo) error {
	insertSql := `INSERT INTO backmessage(key, value, guildid) VALUES(:key, :value, :guildid)`

	tx, _ := db.Beginx()
	_, err := tx.NamedExec(insertSql, backMessages)

	if err != nil {
		tx.Rollback()
		log.Err(err).Msg(`BackMessage Insert ERROR`)
		return err
	}

	return tx.Commit()
}

func (db *BackMessageSqlConnection) FindByKeyAndGuildId(key string, guildId string) ([]vo.BackMessageVo, error) {
	stmt, _ := db.PrepareNamed(`SELECT * FROM backmessage WHERE key=:key AND GUILDID=:guildid`)
	backMessage := vo.BackMessageVo{Key: key, GuildId: guildId}
	backMessageVoSlice := []vo.BackMessageVo{}
	err := stmt.Select(&backMessageVoSlice, backMessage)

	if err != nil {
		log.Err(err).Msg(`BackMessage FindByKey ERROR`)
	}

	return backMessageVoSlice, err
}

func (db *BackMessageSqlConnection) FindAll() []vo.BackMessageVo {
	selectSql := `SELECT * FROM backmessage`
	allValue := []vo.BackMessageVo{}

	db.Select(&allValue, selectSql)

	return allValue
}

func (db *BackMessageSqlConnection) FindByGuildId(guildId string) ([]vo.BackMessageVo, error) {
	stmt, _ := db.PrepareNamed(`SELECT DISTINCT KEY FROM backmessage WHERE GUILDID=:guildid`)
	arg := vo.BackMessageVo{GuildId: guildId}
	valuse := []vo.BackMessageVo{}
	err := stmt.Select(&valuse, arg)

	if err != nil {
		log.Err(err).Msg(`find backmessage table error`)
	}

	return valuse, err
}

func (db *BackMessageSqlConnection) DeleteById(id string) error {
	sql := `DELETE FROM backmessage WHERE id=:id`
	backMessage := vo.BackMessageVo{Id: id}
	_, err := db.NamedExec(sql, backMessage)

	if err != nil {
		log.Err(err).Msgf("Delete %s Error", id)
	}

	return err
}

func (db *BackMessageSqlConnection) DeleteByIdAndKeyAndGuildId(id string, key string, guildId string) error {
	sql := `DELETE FROM backmessage WHERE ID=:id AND KEY=:key AND GUILDID=:guildid`
	arg := vo.BackMessageVo{Id: id, Key: key, GuildId: guildId}
	_, err := db.NamedExec(sql, arg)

	if err != nil {
		log.Err(err).Msgf("Delete Error")
	}

	return err
}

func (db *BackMessageSqlConnection) DeleteByKeyAndValue(key string, value string) error {
	sql := `DELETE FROM backmessage WHERE key=:key AND value=:value`
	backMessage := vo.BackMessageVo{Key: key, Value: value}
	_, err := db.NamedExec(sql, backMessage)

	if err != nil {
		log.Err(err).Msgf(`Delete %v Error`, key)
	}

	return err
}
