// FX Rate — live foreign exchange rate simulation
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	port := getEnv("PORT", "8084")

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	log.Printf("[fx-rate] starting on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("[fx-rate] server failed: %v", err)
	}
}

// healthHandler returns a JSON health response.
// Used by Docker Compose, Kubernetes liveness probes, and docker compose up checks.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"service": "fx-rate",
	})
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
