package redis

import (
	"context"
	"fiber_test/pkg/logs"
	"github.com/redis/go-redis/v9"
	"time"
)

var ctx = context.Background()

type Service struct {
	DB *redis.Client
}

func InitService() Service {
	return Service{
		DB: GetDBClient(),
	}
}

func GetData(key string) any {
	result, err := InitService().DB.Get(ctx, key).Result()
	if err != nil {
		logs.Error(err)

		return nil
	}

	return result
}

func SetData(key string, value any, expireSecond time.Duration) {
	err := InitService().DB.Set(ctx, key, value, expireSecond*time.Second).Err()

	if err != nil {
		logs.Error(err)
	}
}
