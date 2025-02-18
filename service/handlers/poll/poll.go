package vote

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"votingapi/models"
	"votingapi/service/handlers/commom"
	services "votingapi/service/services/poll"
)

func NewPollHandler() *PollHandler {
	return &PollHandler{
		service: services.NewPollService(),
	}
}

type PollHandler struct {
	commom.CommonsHandler
	service services.PollService
}

func (v PollHandler) RegisterRoutes(router *http.ServeMux) {

	router.HandleFunc("GET /poll", v.Poll)
	router.HandleFunc("POST /poll", v.NewPoll)
	router.HandleFunc("POST /vote", v.Vote)
	router.HandleFunc("GET /poll/{id}", v.PollSummary)

}

func (v PollHandler) Poll(w http.ResponseWriter, r *http.Request) {
	response := v.service.Poll(context.Background())
	v.SendJson(w, response, response.HttpStatusCode)
}

func (v PollHandler) NewPoll(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var param models.PollRequest
	err := decoder.Decode(&param)
	if err != nil {
		log.Println(err)
		v.SendJson(w, models.InvalidBodyErrorReponse, models.InvalidBodyErrorReponse.HttpStatusCode)
	}
	defer r.Body.Close()

	response := v.service.NewPoll(context.Background(), param)
	v.SendJson(w, response, response.HttpStatusCode)
}

func (v PollHandler) Vote(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var param models.VoteRequest
	err := decoder.Decode(&param)
	if err != nil {
		log.Println(err)
		v.SendJson(w, models.InvalidBodyErrorReponse, models.InvalidBodyErrorReponse.HttpStatusCode)
	}
	defer r.Body.Close()

	response := v.service.Vote(context.Background(), param)
	v.SendJson(w, response, response.HttpStatusCode)
}

func (v PollHandler) PollSummary(w http.ResponseWriter, r *http.Request) {
	pollID := r.PathValue("id")
	response := v.service.PollSummary(context.Background(), pollID)
	v.SendJson(w, response, response.HttpStatusCode)
}
