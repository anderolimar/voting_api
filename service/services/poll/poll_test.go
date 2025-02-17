package services

import (
	"context"
	"net/http"
	"testing"
	"votingapi/mocks"
	"votingapi/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestVoteService(t *testing.T) {
	t.Run("Testing Vote() method : success response", func(t *testing.T) {
		voteReq := models.VoteRequest{
			CaptchaID:    "123",
			CaptchaInput: "54321",
			PollID:       "q2w3e4r5",
			Vote:         1,
		}

		expectedPoll := models.Poll{
			ID: voteReq.PollID,
			Options: []models.VoteOption{
				{Index: 1, Quantity: 1},
				{Index: 2, Quantity: 2},
			},
		}

		ctrl := gomock.NewController(t)

		captchaCli := mocks.NewMockCaptchaClient(ctrl)
		captchaCli.EXPECT().ValidateCaptcha(gomock.Any(), gomock.Any()).Return(true)

		pubsubCli := mocks.NewMockPubSubClient(ctrl)
		pubsubCli.EXPECT().Publish(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		repo := mocks.NewMockPollRepository(ctrl)
		repo.EXPECT().AddVote(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		repo.EXPECT().GetParcial(gomock.Any(), gomock.Any()).Return(&expectedPoll, nil)

		voteSvc := pollService{
			captchaClient: captchaCli,
			pubSubClient:  pubsubCli,
			repo:          repo,
		}

		resp := voteSvc.Vote(context.TODO(), voteReq)
		assert.Equal(t, http.StatusOK, resp.HttpStatusCode)
		assert.EqualValues(t, expectedPoll, resp.Poll)
	})

	t.Run("Testing Vote() method : empty captcha", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		captchaCli := mocks.NewMockCaptchaClient(ctrl)

		pubsubCli := mocks.NewMockPubSubClient(ctrl)

		voteSvc := pollService{
			captchaClient: captchaCli,
			pubSubClient:  pubsubCli,
		}

		voteReq := models.VoteRequest{
			CaptchaID: "123",
			PollID:    "q2w3e4r5",
			Vote:      1,
		}

		resp := voteSvc.Vote(context.TODO(), voteReq)
		assert.Equal(t, http.StatusBadRequest, resp.HttpStatusCode)
		assert.Equal(t, "REQUIRED_CAPTCHA_INPUT", resp.Code)
		assert.Equal(t, "CAPTCHA answer is required", resp.Message)

	})

	t.Run("Testing Vote() method : invalid captcha answer", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		captchaCli := mocks.NewMockCaptchaClient(ctrl)
		captchaCli.EXPECT().ValidateCaptcha(gomock.Any(), gomock.Any()).Return(false)

		pubsubCli := mocks.NewMockPubSubClient(ctrl)

		voteSvc := pollService{
			captchaClient: captchaCli,
			pubSubClient:  pubsubCli,
		}

		voteReq := models.VoteRequest{
			CaptchaID:    "123",
			CaptchaInput: "54321",
			PollID:       "q2w3e4r5",
			Vote:         1,
		}

		resp := voteSvc.Vote(context.TODO(), voteReq)
		assert.Equal(t, http.StatusBadRequest, resp.HttpStatusCode)
		assert.Equal(t, "INVALID_CAPTCHA", resp.Code)
		assert.Equal(t, "Invalid CAPTCHA answer", resp.Message)

	})
}

func TestPollService(t *testing.T) {
	t.Run("Testing Poll() method : success response", func(t *testing.T) {
		expectedPoll := &models.Poll{
			ID: "q2w3e4r5",
			Options: []models.VoteOption{
				{Index: 1, Quantity: 1},
				{Index: 2, Quantity: 2},
			},
		}

		ctrl := gomock.NewController(t)

		captchaCli := mocks.NewMockCaptchaClient(ctrl)

		pubsubCli := mocks.NewMockPubSubClient(ctrl)

		repo := mocks.NewMockPollRepository(ctrl)
		repo.EXPECT().GetPoll(gomock.Any()).Return(expectedPoll, nil)

		voteSvc := pollService{
			captchaClient: captchaCli,
			pubSubClient:  pubsubCli,
			repo:          repo,
		}

		resp := voteSvc.Poll(context.TODO())
		assert.Equal(t, http.StatusOK, resp.HttpStatusCode)
		assert.EqualValues(t, expectedPoll, resp.Poll)
	})
}
