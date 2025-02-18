package captcha

import (
	"os"

	"github.com/mojocn/base64Captcha"
)

var _captcha *base64Captcha.Captcha

//go:generate mockgen -source=./catcha.go -destination=./../../mocks/mock_catcha.go -typed
type CaptchaClient interface {
	GenerateCaptcha() (captchaID string, base64Image string, err error)
	ValidateCaptcha(captchaID string, captchaValue string) (match bool)
}

func NewCaptchaClient() CaptchaClient {
	return &captchaClient{}
}

type captchaClient struct{}

func (c captchaClient) GenerateCaptcha() (captchaID string, base64Image string, err error) {
	captcha := c.getCaptcha()
	captchaID, base64Image, _, err = captcha.Generate()
	return
}

func (c captchaClient) ValidateCaptcha(captchaID string, captchaValue string) bool {
	ignoreCaptcha := os.Getenv("IGNORE_CAPTCHA") == "true"
	if ignoreCaptcha {
		return true
	}
	return _captcha.Verify(captchaID, captchaValue, true)
}

func (c captchaClient) getCaptcha() *base64Captcha.Captcha {
	if _captcha == nil {
		driver := base64Captcha.NewDriverDigit(100, 240, 4, 0.7, 80)
		_captcha = base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	}
	return _captcha
}
