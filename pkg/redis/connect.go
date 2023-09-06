package redis

import (
	"fiber_test/configs"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func ConnectDB(cnf configs.Config) {
	addr := fmt.Sprintf(
		"%v:%v",
		cnf.RedisHost,
		cnf.RedisPort,
	)

	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cnf.RedisPassword,
		DB:       0,
	})
}

func GetDBClient() *redis.Client {
	return rdb
}
