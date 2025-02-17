package models

type CaptchaResponse struct {
	Response
	Id          string `json:"id"`
	Base64Image string `json:"base64Img"`
}
