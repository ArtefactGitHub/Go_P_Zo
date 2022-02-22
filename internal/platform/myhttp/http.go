package myhttp

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/platform/myerror"
)

type ResponseBase struct {
	StatusCode int   `json:"statuscode"`
	Error      error `json:"error"`
}

func NewResponse(err error, statusCode int, description string) *ResponseBase {
	return &ResponseBase{
		StatusCode: statusCode,
		Error:      myerror.NewError(err, description)}
}

func Write(w http.ResponseWriter, response interface{}, statusCode int) {
	result, _ := json.MarshalIndent(response, "", "\t")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	w.Write(result)
	log.Println(string(result))
}

func WriteSuccessWithLocation(w http.ResponseWriter, response interface{}, statusCode int, location string) {
	result, _ := json.MarshalIndent(response, "", "\t")
	w.Header().Set("Location", location)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	w.Write(result)
	log.Println(string(result))
}

func WriteError(w http.ResponseWriter, err error, statusCode int, description string) {
	response := ResponseBase{
		StatusCode: statusCode,
		Error:      myerror.NewError(err, description)}
	result, _ := json.MarshalIndent(response, "", "\t")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	w.Write(result)
	log.Printf("WriteError() %v", err)
	log.Println(string(result))
}
