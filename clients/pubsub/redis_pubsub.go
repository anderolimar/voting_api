package pubsub

import (
	"context"
	"log"
	"votingapi/bootstrap"

	"github.com/go-redis/redis/v8"
)

func NewPubSubClient() PubSubClient {
	var client *redis.Client

	err := bootstrap.GetContainer().Invoke(func(c *redis.Client) { client = c })
	if err != nil || client == nil {
		log.Default().Fatalln(err)
	}

	return &redisPubSubClient{
		redisClient: client,
	}
}

type redisPubSubClient struct {
	redisClient *redis.Client
}

func (c *redisPubSubClient) Subscribe(ctx context.Context, channel string) ChannelContainer {
	pubsub := c.redisClient.Subscribe(ctx, channel)
	return &redisPubSub{pubsub: pubsub}
}

func (c *redisPubSubClient) Publish(ctx context.Context, channel string, message *PubSubMessage) error {
	err := c.redisClient.Publish(ctx, channel, message.Payload)
	if err != nil {
		return err.Err()
	}
	return nil
}

type redisPubSub struct {
	pubsub *redis.PubSub
}

func (c *redisPubSub) Channel() *PubSubChannel {
	redisch := c.pubsub.Channel()

	wschannel := make(chan *PubSubMessage)

	go func() {
		for msg := range redisch {
			wschannel <- &PubSubMessage{Payload: msg.Payload}
		}
	}()
	return &PubSubChannel{Channel: wschannel}
}

func (c *redisPubSub) Close() error {
	return c.pubsub.Close()
}
