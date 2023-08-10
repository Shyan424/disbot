package sqlsource

import (
	"context"
	"sync"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var connection = new(Connection)

func ConnectPostSql(context context.Context, wait *sync.WaitGroup) {
	initDb()
	<-context.Done()
	defer func() {
		clossDb()
		wait.Done()
	}()
}

func GetDatasource() *Connection {
	return connection
}

func initDb() {
	var once sync.Once
	once.Do(func() {
		dbx, err := sqlx.Open("pgx", viper.GetString("datasource.postgres.uri"))

		// dbx.SetMaxOpenConns(2)
		// dbx.SetConnMaxLifetime(3 * time.Minute)

		if err != nil {
			log.Fatal().Err(err).Msg(`DB Open FAIL`)
		}

		if dbx.Ping() != nil {
			log.Fatal().Err(err).Msg(`DB Connect FAIL`)
		}

		connection.DB = dbx
		log.Info().Msg("DB connected")
	})
}

func clossDb() {
	connection.Close()
}

type Connection struct {
	*sqlx.DB
}
