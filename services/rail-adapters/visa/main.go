// Rail Adapter — Visa
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	port := getEnv("PORT", "8091")

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	log.Printf("[rail-adapter-visa] starting on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("[rail-adapter-visa] server failed: %v", err)
	}
}

// healthHandler returns a JSON health response.
// Used by Docker Compose, Kubernetes liveness probes, and docker compose up checks.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"service": "rail-adapter-visa",
	})
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
