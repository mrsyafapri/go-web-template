package helper

import (
	"encoding/json"
	"errors"
	apps "go-web-template/server/apps/api"
	"log"
	"net/http"
)

func HandleNotMethodAllowed(w http.ResponseWriter, method string) {
	handleResponseError(w, errors.New("method not allowed"), http.StatusMethodNotAllowed)
}

func HandleBadRequest(w http.ResponseWriter, err error) {
	handleResponseError(w, err, http.StatusBadRequest)
}

func HandleInternalServerError(w http.ResponseWriter, err error) {
	handleResponseError(w, err, http.StatusInternalServerError)
}

func HandleNotFound(w http.ResponseWriter, err error) {
	handleResponseError(w, err, http.StatusNotFound)
}

func handleResponseError(w http.ResponseWriter, err error, status int) {
	log.Println(err)
	var response apps.ResponseFail
	response.Status = status
	response.Message = err.Error()

	w.WriteHeader(response.Status)
	json.NewEncoder(w).Encode(response)
}

func HandleSuccess(w http.ResponseWriter, payload interface{}) {
	handleResponseSuccess(w, payload, http.StatusOK)
}

func HandleCreateSuccess(w http.ResponseWriter, payload interface{}) {
	handleResponseSuccess(w, payload, http.StatusCreated)
}

func handleResponseSuccess(w http.ResponseWriter, payload interface{}, status int) {
	var response apps.ResponseSuccess
	response.Status = status
	response.Data = payload

	w.WriteHeader(response.Status)
	json.NewEncoder(w).Encode(response)
}
