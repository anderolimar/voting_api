package captcha

import (
	"encoding/json"
	"net/http"
	"votingapi/clients/captcha"
	"votingapi/service/handlers/commom"
)

func NewVoteHandler() *CaptchaHandler {
	return &CaptchaHandler{}
}

type CaptchaHandler struct {
	commom.CommonsHandler
	captchaClient captcha.CaptchaClient
}

func (c CaptchaHandler) RegisterRoutes(router *http.ServeMux) {
	http.HandleFunc("GET /captcha", c.Captcha)
}

func (c CaptchaHandler) Captcha(w http.ResponseWriter, r *http.Request) {
	id, cap, err := c.captchaClient.GenerateCaptcha()

	body := map[string]interface{}{"captchaID": id, "base64Img": cap}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)
}
