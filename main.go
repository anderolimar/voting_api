package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mojocn/base64Captcha"
)

var _captcha *base64Captcha.Captcha

var _votes map[string]int = map[string]int{
	"opcao1": 0,
	"opcao2": 0,
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/captcha", captchaHandler)
	http.HandleFunc("/vote", voteHandler)
	fmt.Println("Servidor iniciado na porta 8080")
	http.ListenAndServe(":8080", nil)
}

func captchaHandler(w http.ResponseWriter, r *http.Request) {
	id, cap, err := generateCaptcha()

	body := map[string]interface{}{"captchaID": id, "base64Img": cap}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)
}

type VoteRequest struct {
	CaptchaID    string `json:"captchaID"`
	CaptchaInput string `json:"captchaInput"`
	Vote         string `json:"vote"`
}

func voteHandler(w http.ResponseWriter, r *http.Request) {

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
		sendJson(w, body)
		return
	}

	captcha := getCaptcha()
	w.WriteHeader(http.StatusOK)
	if captcha.Verify(param.CaptchaID, param.CaptchaInput, true) {
		if _, ok := _votes[param.Vote]; ok {
			_votes[param.Vote]++
		}
		body = map[string]interface{}{"code": 0, "msg": "success", "votes": _votes}
	} else {
		body = map[string]interface{}{"code": -2, "msg": "Invalid CAPTCHA answer"}
	}
	sendJson(w, body)
}

func sendJson(w http.ResponseWriter, body map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)
}

func generateCaptcha() (string, string, error) {
	captcha := getCaptcha()
	id, b64s, _, err := captcha.Generate()
	return id, b64s, err
}

func getCaptcha() *base64Captcha.Captcha {
	if _captcha == nil {
		driver := base64Captcha.NewDriverDigit(100, 240, 4, 0.7, 80)
		_captcha = base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	}
	return _captcha
}
