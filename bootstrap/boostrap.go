package bootstrap

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/dig"
)

var container *dig.Container
var redisCli *redis.Client
var mongoCli *mongo.Client

func GetContainer() *dig.Container {
	if container == nil {
		container = dig.New()
	}
	return container
}

func Start() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Default().Printf("Error loading .env file: %s\n", err)
	}

	_ = GetContainer().Provide(func() *redis.Client {
		redisHost := os.Getenv("REDIS_HOST")
		fmt.Println("redisHost : " + redisHost)
		if redisCli == nil && redisHost != "" {
			redisPort := os.Getenv("REDIS_PORT")
			redisCli = redis.NewClient(&redis.Options{
				Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
			})

		}

		return redisCli
	})

	_ = GetContainer().Provide(func() *mongo.Client {
		mongoUrl := os.Getenv("MONGO_URL")
		fmt.Println("mongoUrl : " + mongoUrl)
		if mongoCli == nil && mongoUrl != "" {
			clientOptions := options.Client().ApplyURI(mongoUrl)
			var err error
			mongoCli, err = mongo.Connect(context.Background(), clientOptions)
			if err != nil {
				log.Fatal(err)
			}

		}

		return mongoCli
	})
}
