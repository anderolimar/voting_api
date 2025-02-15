package pubsub

import "context"

type PubSubMessage struct {
	Payload string `json:"payload"`
}

type PubSubChannel struct {
	Channel     chan *PubSubMessage
	Disconected bool
}

type PubSubService interface {
	Subscribe(ctx context.Context, channel string) ChannelContainer
	Publish(ctx context.Context, channel string, message *PubSubMessage) error
}

type ChannelContainer interface {
	Channel() *PubSubChannel
	Close() error
}
