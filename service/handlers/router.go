package handlers

import (
	"net/http"
	"votingapi/service/handlers/captcha"
	vote "votingapi/service/handlers/poll"
)

var pollHandler *vote.PollHandler
var captchaHandler *captcha.CaptchaHandler

func RegisterRoutes(router *http.ServeMux) {

	pollHandler = vote.NewPollHandler()

	captchaHandler = captcha.NewVoteHandler()

	router.Handle("/", http.FileServer(http.Dir("./service/static")))
	captchaHandler.RegisterRoutes(router)
	pollHandler.RegisterRoutes(router)

}
