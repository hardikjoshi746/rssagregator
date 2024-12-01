package main

import "net/http"

func handlerError(w http.ResponseWriter, r *http.Request){
	respondWithError(w, 400, "somethin went wrong") // the respondWithError function is called with a status code of 500 and an error message
}