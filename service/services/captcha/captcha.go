package captcha

import (
	"votingapi/clients/captcha"
	"votingapi/models"
)

type CaptchaService interface {
	GenerateCaptcha() *models.CaptchaResponse
}

func NewCatpchaService() CaptchaService {
	return &captchaService{
		captchaClient: captcha.NewCaptchaClient(),
	}
}

type captchaService struct {
	captchaClient captcha.CaptchaClient
}

func (c captchaService) GenerateCaptcha() *models.CaptchaResponse {
	id, base64Image, err := c.captchaClient.GenerateCaptcha()
	if err != nil {
		return &models.CaptchaResponse{Response: models.InternalServerErrorReponse}
	}

	return &models.CaptchaResponse{
		Id:          id,
		Base64Image: base64Image,
		Response:    models.OKReponse,
	}
}
