package bootstrap

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog"
	"go.uber.org/dig"
)

var container *dig.Container
var redisCli *redis.Client

func GetContainer() *dig.Container {
	if container == nil {
		container = dig.New()
	}
	return container
}

func Start() {
	_ = GetContainer().Provide(func() zerolog.Logger { return zerolog.New(nil).With().Timestamp().Logger() })

	_ = GetContainer().Provide(func() *redis.Client {
		redisHost := os.Getenv("REDIS_HOST")
		if redisCli == nil && redisHost != "" {
			redisPort := os.Getenv("REDIS_PORT")
			redisCli = redis.NewClient(&redis.Options{
				Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
			})
			return redisCli
		}

		return nil
	})
}
