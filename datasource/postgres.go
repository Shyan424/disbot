package datasource

import (
	"discordbot/model/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func initDb(conf config.Config) SqlConnection {
	var conn SqlConnection
	dbx, err := sqlx.Open("pgx", conf.Datasource.Postgres.Uri)

	// dbx.SetMaxOpenConns(2)
	// dbx.SetConnMaxLifetime(3 * time.Minute)

	if err != nil {
		log.Fatal().Err(err).Msg(`DB Open FAIL`)
	}

	if dbx.Ping() != nil {
		log.Fatal().Err(err).Msg(`DB Connect FAIL`)
	}

	log.Info().Msg("DB connected")
	conn.DB = dbx

	return conn
}
