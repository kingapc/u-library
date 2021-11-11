package securitty

import (
	"crypto/tls"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/rpinedafocus/u-library/internal/app/u-library/utils"
)

var Client *redis.Client

func Init() {
	//Local configuration
	// dsn := os.Getenv("REDIS_DSN")
	// if len(dsn) == 0 {
	// 	dsn = "localhost:6379"
	// }

	opt, err := redis.ParseURL(utils.GoDotEnvVariable("REDIS_HOST"))

	fmt.Print(opt.Addr + "\n")
	fmt.Print(opt.Password + "\n")
	fmt.Print(opt.DB)
	fmt.Print("\n")

	if err != nil {
		fmt.Print("error1")
		panic(err.Error())
	}

	Client = redis.NewClient(&redis.Options{
		Addr:      opt.Addr,
		Password:  opt.Password,
		DB:        opt.DB,
		TLSConfig: &tls.Config{},
	})

	_, err = Client.Ping().Result()
	if err != nil {
		fmt.Print("error2")
		panic(err.Error())
	}
}
