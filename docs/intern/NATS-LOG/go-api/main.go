package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("STARTING NATS LOG MONITORING SERVICE")
	natsURL := os.Getenv("NATS_URL")
	subjectPrefix := os.Getenv("NATS_SUBJECT_PREFIX")
	logPaths := strings.Split(os.Getenv("LOG_PATHS"), ",")

	natsConnection, err := connectNATS(natsURL)
	if err != nil {
		log.Fatal("Failed to connect to NATS:", err)
	}

	defer natsConnection.Close()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	isRunning := false

	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		if isRunning {
			fmt.Fprintln(w, "Log monitoring is already running.")
			return
		}
		fmt.Fprintln(w, "Log monitoring started.")
		go func() {
			isRunning = true
			ticker.Reset(5 * time.Second)
			for range ticker.C {
				publishLogFiles(natsConnection, logPaths, subjectPrefix)
			}
		}()
	})

	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		if isRunning {
			ticker.Stop()
			isRunning = false
			fmt.Fprintln(w, "Log monitoring stopped.")
			log.Print(".log counting stopped")
		} else {
			fmt.Fprintln(w, "Log monitoring is not running.")
		}

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
