package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
)

type TotalLogStructure struct {
	TotalCount int `json:"total_count"`
	TotalSize  int `json:"total_size"`
}
type LogStructure struct {
	Filename string `json:"filename"`
	Size     int    `json:"size"`
}

func main() {
	fmt.Println("STARTING NATS LOG MONITORING SERVICE")
	Nats_URL := os.Getenv("NATS_URL")
	Subject_Prefx := os.Getenv("NATS_SUBJECT_PREFIX")
	Log_Paths := strings.Split(os.Getenv("LOG_PATHS"), ",")

	nats_connection, err := nats.Connect(Nats_URL)
	if err != nil {
		log.Fatal("Nats Connection error", err)
	}
	defer nats_connection.Close()

	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		totalCount := 0
		var totalSize float64 = 0.0
		for _, log_path := range Log_Paths {

			log_path = strings.TrimSpace(log_path)

			if log_path == "" {
				continue
			}

			file_info, err := os.Stat(log_path)
			if err != nil {
				log.Printf("Error accessing file %s: %v", log_path, err)
				continue
			}

			if file_info.IsDir() {
				files, err := os.ReadDir(log_path)
				if err != nil {
					log.Printf("Error reading directory %s: %v", log_path, err)
					continue
				}
				for _, file := range files {
					if !file.IsDir() && strings.HasSuffix(file.Name(), ".log") {
						fullPath := filepath.Join(log_path, file.Name())
						info, err := os.Stat(fullPath)
						if err == nil {
							metric := LogStructure{
								Filename: info.Name(),
								Size:     int(info.Size()),
							}
							jsonData, _ := json.Marshal(metric)
							subject := fmt.Sprintf("%s.%s", Subject_Prefx, metric.Filename)

							err := nats_connection.Publish(subject, jsonData)
							if err != nil {
								log.Printf("Error publishing message for file %s: %v", metric.Filename, err)
							} else {
								log.Printf("Published message for file %s", metric.Filename)
								totalCount++
								totalSize += float64(metric.Size)
							}
						}
					}
				}
			} else {
				if strings.HasSuffix(file_info.Name(), ".log") {
					metric := LogStructure{
						Filename: file_info.Name(),
						Size:     int(file_info.Size()),
					}
					jsonData, _ := json.Marshal(metric)
					subject := fmt.Sprintf("%s.%s", Subject_Prefx, metric.Filename)

					err := nats_connection.Publish(subject, jsonData)
					if err != nil {
						log.Printf("Error publishing message for file %s: %v", metric.Filename, err)
					} else {
						log.Printf("Published message for file %s", metric.Filename)
						totalCount++
						totalSize += float64(metric.Size)
					}
				}
			}
			log.Printf("Total log files processed: %d, Total size: %.4f MB", totalCount, totalSize/1024/1024)
		}
	}
}
