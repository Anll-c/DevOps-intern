package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/nats-io/nats.go"
)

type TotalLogStructure struct {
	TotalCount int     `json:"total_count"`
	TotalSize  float64 `json:"total_size"`
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
			slog.Error("Error accessing log path",
				"path", logPath,
				"error", err,
			)
			continue
		}

		// Klasör ise içine bak
		if fileInfo.IsDir() {
			files, err := os.ReadDir(logPath)
			if err != nil {
				slog.Error("Error reading directory",
					"path", logPath,
					"error", err,
				)
				continue
			}

			for _, file := range files {
				if !file.IsDir() && strings.HasSuffix(file.Name(), ".log") {
					fullPath := filepath.Join(logPath, file.Name())
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
		slog.Info("Total log files processed", "count", totalCount, "size", fmt.Sprintf("%.4f MB", totalSize/1024/1024))
	}

	publishSummary(nc, subjectPrefix, totalCount, totalSize)
	if err := nc.Flush(); err != nil {
		slog.Error("Failed to flush NATS connection", "error", err)
	}

}

func publishSummary(nc *nats.Conn, subjectPrefix string, totalCount int, totalSize float64) {
	if nc.Status() != nats.CONNECTED {
		slog.Warn("Skipping summary publish: NATS not connected", "status", nc.Status())
		return
	}

	summary := TotalLogStructure{
		TotalCount: totalCount,
		TotalSize:  totalSize,
	}

	jsonData, err := json.Marshal(summary)
	if err != nil {
		slog.Error("Error marshaling summary json", "error", err)
		return
	}

	subject := fmt.Sprintf("%s.summary", subjectPrefix)
	if err := nc.Publish(subject, jsonData); err != nil {
		slog.Error("Error publishing summary", "error", err)
		return
	}

	slog.Info("Published summary", "total_count", totalCount, "total_size", fmt.Sprintf("%.4f MB", totalSize/1024/1024))
	nc.Flush()
}
