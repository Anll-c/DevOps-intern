package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func sendJSONResponse(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		slog.Error("Failed to write JSON response", "error", err)
	}
}

func sendJSONError(w http.ResponseWriter, statusCode int, message string) {
	sendJSONResponse(w, statusCode, map[string]string{"error": message})
}
