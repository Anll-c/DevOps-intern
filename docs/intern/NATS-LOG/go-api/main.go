package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"github.com/nats-io/nats.go/v2"
)

type LogMetric struct{
	Filename string `json:"filename"`
	Size int `json:"size"`
}

func main() {
	Nats_URL := os.Getenv("NATS_URL")
	Subject_Prefx := os.Getenv("NATS_SUBJECT_PREFIX")
	Log_Paths := strings.Split(os.Getenv("LOG_PATHS"), ",")

	nc, err := nats.Connect(Nats_URL)
	if err != nil { 
		log.Fatal("Nats Connection error",err)
	}
	defer nc.Close()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	
	for range ticker.C{
		for _, log_path := range Log_Paths{
			file_info, err := os.Stat(log_path)
			if err == nil {
				metric := LogMetric{
					Filename: file_info.Name(),
					Size: int(file_info.Size()),
				}
				jsonData,_ := json.Marshal(metric)
				subject := fmt.Sprintf("%s.%s", Subject_Prefx, metric.Filename)

				if err := nc.Publish(subject, jsonData); err != nil {
					log.Println("publish error:", err)
				}
			}
		}
	}
}