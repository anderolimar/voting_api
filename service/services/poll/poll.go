package services

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"votingapi/cfg"
	"votingapi/clients/captcha"
	"votingapi/clients/pubsub"
	"votingapi/models"
	"votingapi/repositories"
)

type PollService interface {
	Poll(ctx context.Context) *models.PollResponse
	NewPoll(ctx context.Context, pollReq models.PollRequest) *models.Response
	Vote(ctx context.Context, voteRequest models.VoteRequest) *models.VoteResponse
}

func NewPollService() PollService {
	return &pollService{
		captchaClient: captcha.NewCaptchaClient(),
		pubSubClient:  pubsub.NewPubSubClient(),
		repo:          repositories.NewVotesRepository(),
	}
}

type pollService struct {
	captchaClient captcha.CaptchaClient
	pubSubClient  pubsub.PubSubClient
	repo          repositories.PollRepository
}

func (v pollService) Poll(ctx context.Context) *models.PollResponse {
	resp, err := v.repo.GetPoll(ctx)
	if err != nil {
		return &models.PollResponse{Response: models.InvalidBodyErrorReponse}
	}
	return &models.PollResponse{Poll: resp, Response: models.OKReponse}
}

func (v pollService) NewPoll(ctx context.Context, pollReq models.PollRequest) *models.Response {
	err := v.repo.AddPoll(ctx, &models.Poll{
		Title:   pollReq.Title,
		Options: pollReq.Options,
	}, time.Second*time.Duration(cfg.POLL_SEC_DURATION))
	if err != nil {
		return &models.InvalidBodyErrorReponse
	}
	return &models.Response{HttpStatusCode: http.StatusOK}
}

func (v pollService) Vote(ctx context.Context, voteRequest models.VoteRequest) *models.VoteResponse {
	if len(voteRequest.CaptchaInput) == 0 {
		return &models.VoteResponse{Response: models.Response{Code: "REQUIRED_CAPTCHA_INPUT", Message: "CAPTCHA answer is required", HttpStatusCode: http.StatusBadRequest}}
	}

	if v.captchaClient.ValidateCaptcha(voteRequest.CaptchaID, voteRequest.CaptchaInput) {
		value, err := json.Marshal(voteRequest)
		if err != nil {
			return &models.VoteResponse{Response: models.InvalidBodyErrorReponse}
		}

		v.pubSubClient.Publish(ctx, cfg.VOTING_CHANNEL, &pubsub.PubSubMessage{
			Payload: string(value),
		})

		err = v.repo.AddVote(ctx, voteRequest.PollID, voteRequest.Vote)
		if err != nil {
			ret := models.VoteResponse{Response: models.InternalServerErrorReponse}
			return &ret
		}

		pool, err := v.repo.GetParcial(ctx, voteRequest.PollID)
		if err != nil {
			ret := models.VoteResponse{Response: models.InternalServerErrorReponse}
			return &ret
		}

		return &models.VoteResponse{Poll: *pool, Response: models.OKReponse}
	}

	return &models.VoteResponse{Response: models.Response{Code: "INVALID_CAPTCHA", Message: "Invalid CAPTCHA answer", HttpStatusCode: http.StatusBadRequest}}
}
