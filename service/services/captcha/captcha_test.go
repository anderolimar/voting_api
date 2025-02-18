package captcha

import (
	"net/http"
	"testing"
	"votingapi/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCaptchaService(t *testing.T) {
	t.Run("Testing GenerateCaptcha() method : success response", func(t *testing.T) {
		expectedId := "captchaID"
		expectedBase64 := "base64Image"

		ctrl := gomock.NewController(t)

		captchaCli := mocks.NewMockCaptchaClient(ctrl)
		captchaCli.EXPECT().GenerateCaptcha().Return(expectedId, expectedBase64, nil)

		captchaSvc := captchaService{
			captchaClient: captchaCli,
		}

		resp := captchaSvc.GenerateCaptcha()
		assert.Equal(t, http.StatusOK, resp.HttpStatusCode)
		assert.EqualValues(t, expectedId, resp.Id)
		assert.EqualValues(t, expectedBase64, resp.Base64Image)
	})
}
