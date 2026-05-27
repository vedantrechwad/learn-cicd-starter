package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string, logErr error) {
	if logErr != nil {
		log.Println("error occurred") // #nosec G706
	}
	if code > 499 {
		log.Println("5XX error occurred") // #nosec G706
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Println("failed to marshal JSON response") // #nosec G706
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	if _, err := w.Write(dat); err != nil {
	log.Printf("error writing JSON response: %v", err)
	}
}
