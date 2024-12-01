package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("responding with 5xx error: ", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJson(w, code, errorResponse{
		Error : msg,
	})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) { // this function is used to send a JSON response back to the client
	data, err := json.Marshal(payload) // the payload is marshalled into a JSON string
	if err != nil {
		log.Printf("failed to marshal because of %v", err)
		w.WriteHeader(code) // if there is an error, the status code is set to 500
		return
	}
	w.Header().Add("Content-Type", "application/json") // the Content-Type header is set to application/json
	w.WriteHeader(code) // the status code is set to 200
	w.Write(data) // the JSON string is written to the response
}