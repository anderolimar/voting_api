package captcha

import (
	"net/http"
	"votingapi/service/handlers/commom"
	"votingapi/service/services/captcha"
)

func NewVoteHandler() *CaptchaHandler {
	return &CaptchaHandler{
		captchaService: captcha.NewCatpchaService(),
	}
}

type CaptchaHandler struct {
	commom.CommonsHandler
	captchaService captcha.CaptchaService
}

func (c CaptchaHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /captcha", c.Captcha)
}

func (c CaptchaHandler) Captcha(w http.ResponseWriter, r *http.Request) {
	resp := c.captchaService.GenerateCaptcha()
	c.SendJson(w, resp, resp.HttpStatusCode)
}
