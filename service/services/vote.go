package services

import (
	"context"
	"net/http"
	"votingapi/cfg"
	"votingapi/clients/captcha"
	"votingapi/clients/pubsub"
)

type VoteRequest struct {
	CaptchaID    string `json:"captchaID"`
	CaptchaInput string `json:"captchaInput"`
	Vote         string `json:"vote"`
}

type VoteResponse struct {
	Code           int    `json:"code"`
	Message        string `json:"message"`
	HttpStatusCode int    `json:"statusCode"`
}

type VoteService interface {
	Vote(ctx context.Context, voteRequest VoteRequest) VoteResponse
}

func NewVoteService() VoteService {
	return &voteService{}
}

type voteService struct {
	captchaClient captcha.CaptchaClient
	pubSubClient  pubsub.PubSubClient
}

func (v voteService) Vote(ctx context.Context, voteRequest VoteRequest) VoteResponse {
	if len(voteRequest.CaptchaInput) == 0 {
		return VoteResponse{Code: -1, Message: "CAPTCHA answer is required", HttpStatusCode: http.StatusBadRequest}
	}

	if v.captchaClient.ValidateCaptcha(voteRequest.CaptchaID, voteRequest.CaptchaInput) {
		v.pubSubClient.Publish(ctx, cfg.VOTING_CHANNEL, &pubsub.PubSubMessage{
			Payload: voteRequest.Vote,
		})
		return VoteResponse{Code: 0, Message: "Success", HttpStatusCode: http.StatusOK}
	}

	return VoteResponse{Code: -1, Message: "Invalid CAPTCHA answer", HttpStatusCode: http.StatusBadRequest}
}
