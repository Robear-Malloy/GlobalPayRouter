package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.HandleFunc("/route", routeHandler).Methods("POST")

	log.Println("GlobalRoute Pay starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "routing",
		"message": "payment route calculated",
	})
}
