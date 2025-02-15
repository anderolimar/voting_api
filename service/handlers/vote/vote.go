package vote

import (
	"encoding/json"
	"log"
	"net/http"
	"votingapi/clients/captcha"
	"votingapi/service/handlers/commom"
)

func RegisterRoutes(router *http.ServeMux) {

}

type VoteRequest struct {
	CaptchaID    string `json:"captchaID"`
	CaptchaInput string `json:"captchaInput"`
	Vote         string `json:"vote"`
}

func NewVoteHandler() *VoteHandler {
	return &VoteHandler{}
}

type VoteHandler struct {
	commom.CommonsHandler
	captchaClient captcha.CaptchaClient
}

func (v VoteHandler) RegisterRoutes(router *http.ServeMux) {
	http.HandleFunc("POST /vote", v.Vote)
}

func (v VoteHandler) Vote(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var param VoteRequest
	err := decoder.Decode(&param)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	var body map[string]interface{}

	if len(param.CaptchaInput) == 0 {
		http.Error(w, "CAPTCHA answer is required", http.StatusBadRequest)
		body = map[string]interface{}{"code": -1, "msg": "CAPTCHA answer is required"}
		v.SendJson(w, body, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	if v.captchaClient.ValidateCaptcha(param.CaptchaID, param.CaptchaInput) {
		body = map[string]interface{}{"code": 0, "msg": "success", "votes": 0}
	} else {
		body = map[string]interface{}{"code": -2, "msg": "Invalid CAPTCHA answer"}
	}
	v.SendJson(w, body, http.StatusOK)
}
