package vote

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"votingapi/service/handlers/commom"
	"votingapi/service/services"
)

func NewVoteHandler() *VoteHandler {
	return &VoteHandler{}
}

type VoteHandler struct {
	commom.CommonsHandler
	service services.VoteService
}

func (v VoteHandler) RegisterRoutes(router *http.ServeMux) {
	http.HandleFunc("POST /vote", v.Vote)
}

func (v VoteHandler) Vote(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var param services.VoteRequest
	err := decoder.Decode(&param)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	response := v.service.Vote(context.Background(), param)
	v.SendJson(w, response, response.HttpStatusCode)
}
