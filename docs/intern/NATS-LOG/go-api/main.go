package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
)

type LogStructure struct {
	Filename string `json:"filename"`
	Size     int    `json:"size"`
}

func main() {
	Nats_URL := os.Getenv("NATS_URL")
	Subject_Prefx := os.Getenv("NATS_SUBJECT_PREFIX")
	Log_Paths := strings.Split(os.Getenv("LOG_PATHS"), ",")

	nats_connection, err := nats.Connect(Nats_URL)
	if err != nil {
		log.Fatal("Nats Connection error", err)
	}
	defer nats_connection.Close()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		for _, log_path := range Log_Paths {
			file_info, err := os.Stat(log_path)
			if err == nil {
				metric := LogStructure{
					Filename: file_info.Name(),
					Size:     int(file_info.Size()),
				}
				jsonData, _ := json.Marshal(metric)
				subject := fmt.Sprintf("%s.%s", Subject_Prefx, metric.Filename)

				nats_connection.Publish(subject, jsonData)
			}
		}
	}
}
