package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
	"votingapi/bootstrap"
	"votingapi/cfg"
	"votingapi/models"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var LOCK_KEY = "poll-lock"

type PollDoc struct {
	ID      primitive.ObjectID  `json:"id" bson:"id"`
	Title   string              `json:"title" bson:"title"`
	Options []models.VoteOption `json:"options" bson:"options"`
}

func (p PollDoc) IncrementVote(index int) {
	for idx := range p.Options {
		if p.Options[idx].Index == index {
			p.Options[idx].Quantity++
			break
		}
	}
}

type PollRepository interface {
	GetPoll(ctx context.Context) (*models.Poll, error)
	GetParcial(ctx context.Context, pollID string) (*models.Poll, error)
	AddPoll(ctx context.Context, poll *models.Poll, duration time.Duration) error
	AddVote(ctx context.Context, pollID string, voteIndex int) error
	UpdateVote(ctx context.Context, pollID string, voteIndex int) error
}

func NewVotesRepository() PollRepository {
	var redisclient *redis.Client

	err := bootstrap.GetContainer().Invoke(func(c *redis.Client) { redisclient = c })
	if err != nil || redisclient == nil {
		log.Default().Fatalln(err)
	}
	var mongoclient *mongo.Client

	err = bootstrap.GetContainer().Invoke(func(c *mongo.Client) { mongoclient = c })
	if err != nil || mongoclient == nil {
		log.Default().Fatalln(err)
	}

	return &pollRepository{
		mongoclient: mongoclient,
		redisclient: redisclient,
	}
}

type pollRepository struct {
	mongoclient *mongo.Client
	redisclient *redis.Client
}

func (v pollRepository) GetPoll(ctx context.Context) (*models.Poll, error) {
	// TODO : Get poll by ID
	res := v.redisclient.Get(ctx, cfg.CURRENT_POLL)
	if res == nil {
		return nil, nil
	}

	bytes, err := res.Bytes()
	if err != nil {
		return nil, err
	}
	var poll models.Poll
	err = json.Unmarshal(bytes, &poll)
	if err != nil {
		return nil, err
	}

	return &poll, nil
}

func (v pollRepository) AddPoll(ctx context.Context, poll *models.Poll, duration time.Duration) error {
	db := v.mongoclient.Database(cfg.MONGO_DATABASE)
	col := db.Collection(cfg.MONGO_COLLECTION)

	pollDoc := PollDoc{
		ID:      primitive.NewObjectID(),
		Title:   poll.Title,
		Options: poll.Options,
	}

	_, err := col.InsertOne(ctx, pollDoc)
	if err != nil {
		fmt.Printf("Error to AddPoll: %v\n", err)
		return err
	}

	poll.ID = pollDoc.ID.Hex()

	pipe := v.redisclient.Pipeline()
	// TODO : Add poll by ID
	pipe.Set(ctx, cfg.CURRENT_POLL, poll, duration)

	for idx := range poll.Options {
		pipe.Set(ctx, fmt.Sprintf("%d_%s", idx, pollDoc.ID.Hex()), 0, duration)
	}
	if _, err = pipe.Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (v pollRepository) AddVote(ctx context.Context, pollID string, voteIndex int) error {
	pipe := v.redisclient.Pipeline()
	pipe.Incr(ctx, fmt.Sprintf("%d_%s", voteIndex, pollID))
	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (v pollRepository) GetParcial(ctx context.Context, pollID string) (*models.Poll, error) {
	poll, err := v.GetPoll(ctx)
	if err != nil {
		return nil, err
	}

	for idx, opt := range poll.Options {
		res := v.redisclient.Get(ctx, fmt.Sprintf("%d_%s", opt.Index, pollID))
		err := res.Err()
		if err != nil {
			if err == redis.Nil {
				continue
			}
			return nil, err
		}

		var quantity = 0
		quantityStr := res.Val()
		quantity, err = strconv.Atoi(quantityStr)
		if err != nil {
			return nil, err
		}
		poll.Options[idx].Quantity = quantity
	}

	return poll, nil
}

func (v pollRepository) UpdateVote(ctx context.Context, pollID string, voteIndex int) error {
	fmt.Printf("UpdateVote : poll : %s | vote : %d", pollID, voteIndex)

	db := v.mongoclient.Database(cfg.MONGO_DATABASE)
	col := db.Collection(cfg.MONGO_COLLECTION)

	id, err := primitive.ObjectIDFromHex(pollID)
	if err != nil {
		return err
	}
	filter := bson.M{"id": id}

	v.acquireLock(ctx, LOCK_KEY, time.Second*10)
	defer v.releaseLock(ctx, LOCK_KEY)

	var result PollDoc
	err = col.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No document found")
		}
		return err
	}

	result.IncrementVote(voteIndex)

	update := bson.M{"$set": result}
	_, err = col.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (v pollRepository) acquireLock(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	// Use the SET command to try to acquire the lock
	result, err := v.redisclient.SetNX(ctx, key, "lock", expiration).Result()
	if err != nil {
		return false, err
	}
	return result, nil
}

func (v pollRepository) releaseLock(ctx context.Context, key string) error {
	// Use the DEL command to release the lock
	_, err := v.redisclient.Del(ctx, key).Result()
	return err
}
