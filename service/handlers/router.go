package handlers

import (
	"net/http"
	"votingapi/service/handlers/captcha"
	"votingapi/service/handlers/vote"
)

var voteHandler = vote.NewVoteHandler()
var captchaHandler = captcha.NewVoteHandler()

func RegisterRoutes(router *http.ServeMux) {
	router.Handle("GET /", http.FileServer(http.Dir("../static")))
	captchaHandler.RegisterRoutes(router)
	voteHandler.RegisterRoutes(router)

}
