package securitty

import (
	"crypto/tls"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/rpinedafocus/u-library/pkg/utils"
)

var Client *redis.Client

func Init() {
	//Local configuration
	// dsn := os.Getenv("REDIS_DSN")
	// if len(dsn) == 0 {
	// 	dsn = "localhost:6379"
	// }

	opt, err := redis.ParseURL(utils.GoDotEnvVariable("REDIS_HOST"))
	if err != nil {
		panic(err.Error())
	}

	Client = redis.NewClient(&redis.Options{
		Addr:        opt.Addr,
		Password:    opt.Password,
		DB:          opt.DB,
		IdleTimeout: 5 * time.Minute,
		MaxRetries:  2,
		TLSConfig:   &tls.Config{},
	})

	_, err = Client.Ping().Result()
	if err != nil {
		panic(err.Error())
	}
}
