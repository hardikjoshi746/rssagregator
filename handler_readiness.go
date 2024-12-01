package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request){
	respondWithJson(w, 200, struct{}{}) // the respondWithJson function is called with a status code of 200 and an empty object
}