package worker

import (
	"context"
	"fmt"
	"votingapi/cfg"
	"votingapi/clients/pubsub"
)

func Run(ctx context.Context) {
	pubSubService := pubsub.NewPubSubClient()

	sub := pubSubService.Subscribe(ctx, cfg.VOTING_CHANNEL)

	ch := sub.Channel()

	for msg := range ch.Channel {
		fmt.Println(msg)
	}
}
