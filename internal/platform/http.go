package platform

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ArtefactGitHub/Go_P_Zo/internal/myerror"
)

type ResponseBase struct {
	StatusCode int   `json:"statuscode"`
	Error      error `json:"error"`
}

func WriteSuccess(w http.ResponseWriter, response interface{}, statusCode int) {
	result, _ := json.MarshalIndent(response, "", "\t")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	w.Write(result)
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
	log.Println(err)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	w.Write(result)
	log.Println(err)
}
