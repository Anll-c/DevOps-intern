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

func processLogFile(nc *nats.Conn, subjectPrefix string, fileInfo os.FileInfo) int {
	metric := LogStructure{
		Filename: fileInfo.Name(),
		Size:     int(fileInfo.Size()),
	}

	jsonData, err := json.Marshal(metric)
	if err != nil {
		log.Printf("Error marshaling json for %s: %v", metric.Filename, err)
		return 0
	}

	subject := fmt.Sprintf("%s.%s", subjectPrefix, metric.Filename)
	err = nc.Publish(subject, jsonData)
	if err != nil {
		log.Printf("Error publishing message for file %s: %v", metric.Filename, err)
		return 0
	}

	log.Printf("Published message for file name:%s size:%d", metric.Filename, metric.Size)
	return metric.Size
}

func scanAndPublish(nc *nats.Conn, logPaths []string, subjectPrefix string) {
	totalCount := 0
	var totalSize float64 = 0.0

	for _, logPath := range logPaths {
		logPath = strings.TrimSpace(logPath)
		if logPath == "" {
			continue
		}

		fileInfo, err := os.Stat(logPath)
		if err != nil {
			log.Printf("Error accessing file %s: %v", logPath, err)
			continue
		}

		// Klasör ise içine bak
		if fileInfo.IsDir() {
			files, err := os.ReadDir(logPath)
			if err != nil {
				log.Printf("Error reading directory %s: %v", logPath, err)
				continue
			}

			for _, file := range files {
				if !file.IsDir() && strings.HasSuffix(file.Name(), ".log") {
					fullPath := filepath.Join(logPath, file.Name()) //filepath
					info, err := os.Stat(fullPath)
					if err == nil {
						size := processLogFile(nc, subjectPrefix, info)
						if size > 0 {
							totalCount++
							totalSize += float64(size)
						}
					}
				}
			}
		} else {
			// .log dosyası ise
			if strings.HasSuffix(fileInfo.Name(), ".log") {
				size := processLogFile(nc, subjectPrefix, fileInfo)
				if size > 0 {
					totalCount++
					totalSize += float64(size)
				}
			}
		}
		log.Printf("Total log files processed: %d, Total size: %.4f MB\n\n", totalCount, totalSize/1024/1024)
	}
}

func main() {
	fmt.Println("STARTING NATS LOG MONITORING SERVICE")
	natsURL := os.Getenv("NATS_URL")
	subjectPrefix := os.Getenv("NATS_SUBJECT_PREFIX")
	logPaths := strings.Split(os.Getenv("LOG_PATHS"), ",")

	natsConnection, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal("Nats Connection error", err)
	}
	defer natsConnection.Close()

	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		scanAndPublish(natsConnection, logPaths, subjectPrefix)
	}
}
