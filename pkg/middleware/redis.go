package securitty

import (
	"os"

	"github.com/go-redis/redis/v7"
)

var Client *redis.Client

func Init() {
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}

	Client = redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	_, err := Client.Ping().Result()
	if err != nil {
		panic(err)
	}
}
