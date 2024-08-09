package datasource

import (
	"context"
	"discordbot/model/config"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

var sqlConn SqlConnection
var rdb RedisClient

// Connect Sql DB
// Connect redis
//
// Close connection when context is cancel
func ConnectDbs(ctx context.Context, wg *sync.WaitGroup, conf config.Config) {
	sqlConn = initDb(conf)
	rdb = initRedis(conf)

	wg.Add(1)
	go func() {
		<-ctx.Done()
		sqlConn.Close()
		rdb.Close()
		wg.Done()
	}()
}

func GetDatasource() *SqlConnection {
	return &sqlConn
}

func GetRedisClient() *RedisClient {
	return &rdb
}

type SqlConnection struct {
	*sqlx.DB
}

type RedisClient struct {
	*redis.Client
}
