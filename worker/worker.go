package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"votingapi/bootstrap"
	"votingapi/cfg"
	"votingapi/clients/pubsub"
	"votingapi/models"
	"votingapi/repositories"
)

func Run(ctx context.Context) {
	bootstrap.Start()

	pubSubService := pubsub.NewPubSubClient()
	repository := repositories.NewVotesRepository()

	sub := pubSubService.Subscribe(ctx, cfg.VOTING_CHANNEL)

	ch := sub.Channel()

	fmt.Println("Worker Channel Start")
	for msg := range ch.Channel {
		fmt.Println(msg)
		var vote models.VoteRequest
		json.Unmarshal([]byte(msg.Payload), &vote)

		vote.Retry++
		if vote.Retry > cfg.MAX_RETRIES {
			fmt.Printf("MAX RETRIES REACHED : %s", msg)
			continue
		}

		if err := repository.UpdateVote(ctx, vote.PollID, vote.Vote); err != nil {
			bytes, err := json.Marshal(vote)
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = pubSubService.Publish(ctx, cfg.VOTING_CHANNEL, &pubsub.PubSubMessage{Payload: string(bytes)})
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	fmt.Println("Worker End")
}
