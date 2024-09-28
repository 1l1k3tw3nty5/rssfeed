package main

import "net/http"

func handlerAlive(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct{}{})
}
