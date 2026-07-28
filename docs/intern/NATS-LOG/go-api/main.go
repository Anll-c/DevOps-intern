package main

import (
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	Logger()
	slog.Info("Starting NATS Log Publisher API")
	natsURL := os.Getenv("NATS_URL")
	subjectPrefix := os.Getenv("NATS_SUBJECT_PREFIX")
	logPaths := strings.Split(os.Getenv("LOG_PATHS"), ",")
	goPort := os.Getenv("GO_API_PORT")
	if goPort == "" {
		slog.Warn("GO_API_PORT environment variable is not set. Using default port 8080.")
		goPort = "8080"
	}

	natsConnection, err := connectNATS(natsURL)
	if err != nil {
		slog.Error("There was an error connecting to NATS",
			"url", natsURL,
			"error", err,
		)
		defer natsConnection.Close()
	}

	if natsConnection != nil {
		defer natsConnection.Close()
	}

	defer natsConnection.Close()

	ticker := time.NewTicker(5 * time.Second)
	ticker.Stop()
	defer ticker.Stop()
	isRunning := false

	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		if isRunning {
			slog.Warn("Log monitoring is already running.")
			sendJSONResponse(w, http.StatusConflict, map[string]string{
				"status":  "already_running",
				"message": "Log monitoring is already running.",
			})
			return
		}

		go func() {
			isRunning = true
			ticker.Reset(5 * time.Second)
			for range ticker.C {
				publishLogFiles(natsConnection, logPaths, subjectPrefix)
			}
		}()

		slog.Info("Log monitoring started")
		sendJSONResponse(w, http.StatusOK, map[string]string{
			"status":  "started",
			"message": "Log monitoring started.",
		})

	})

	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		if isRunning {
			ticker.Stop()
			isRunning = false
			slog.Info("Log monitoring stopped")
			sendJSONResponse(w, http.StatusOK, map[string]string{
				"status":  "stopped",
				"message": "Log monitoring stopped.",
			})

		} else {
			slog.Warn("Log monitoring is already stopped")
			sendJSONResponse(w, http.StatusConflict, map[string]string{
				"status":  "already_stopped",
				"message": "Log monitoring is not running.",
			})
			return

		}

	})
	slog.Info("HTTP server starting", "port", goPort)
	slog.Error("HTTP server error", "error", http.ListenAndServe(":"+goPort, nil))
}

//err loglarına bak nats stop ile, log türlerine bak ne eklenebilir düşün.
