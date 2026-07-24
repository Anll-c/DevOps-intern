package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/nats-io/nats.go"
)

type TotalLogStructure struct {
	TotalCount int `json:"total_count"`
	TotalSize  int `json:"total_size"`
}

func publishLogFiles(nc *nats.Conn, logPaths []string, subjectPrefix string) {
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

						totalCount++
						totalSize += size

					}
				}
			}
		} else {
			// .log dosyası ise
			if strings.HasSuffix(fileInfo.Name(), ".log") {
				size := processLogFile(nc, subjectPrefix, fileInfo)

				totalCount++
				totalSize += size

			}
		}
		log.Printf("Total log files processed: %d, Total size: %.4f MB\n\n", totalCount, totalSize/1024/1024)
	}
}
