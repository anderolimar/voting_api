package models

import "net/http"

type Response struct {
	Code           string `json:"code,omitempty"`
	Message        string `json:"message,omitempty"`
	HttpStatusCode int    `json:"statusCode,omitempty"`
}

var OKReponse = Response{HttpStatusCode: http.StatusOK}

var InternalServerErrorReponse = Response{Code: "INTERNAL_SERVER_ERROR", Message: "Internal Error", HttpStatusCode: http.StatusInternalServerError}

var InvalidBodyErrorReponse = Response{Code: "INVALID_BODY", Message: "Invalid Body Error", HttpStatusCode: http.StatusBadRequest}

var InvalidCaptchaErrorReponse = Response{Code: "INVALID_CAPTCHA", Message: "Invalid CAPTCHA answer", HttpStatusCode: http.StatusBadRequest}
