package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("[-] Failed to marshal JSON payload: %s", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)

}
func respondWithError(w http.ResponseWriter, code int, err string) {
	if code > 499 {
		log.Printf("[-]Responding with %d ERROR: %s", code, err)
	}
	type errStruct struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errStruct{
		Error: err,
	})
}
