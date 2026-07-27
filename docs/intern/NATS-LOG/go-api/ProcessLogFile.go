package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"github.com/nats-io/nats.go"
)

type LogStructure struct {
	Filename string  `json:"filename"`
	Size     float64 `json:"size"`
}

func processLogFile(nc *nats.Conn, subjectPrefix string, fileInfo os.FileInfo) float64 {
	metric := LogStructure{
		Filename: fileInfo.Name(),
		Size:     float64(fileInfo.Size()),
	}

	jsonData, err := json.Marshal(metric)
	if err != nil {
		slog.Error("Failed to marshal log structure to JSON",
			"filename", metric.Filename,
			"error", err,
		)
		return 0
	}

	subject := fmt.Sprintf("%s.%s", subjectPrefix, metric.Filename)

	if nc.IsConnected() {
		err = nc.Publish(subject, jsonData)
		if err != nil {
			slog.Error("Failed to publish message for file",
				"filename", metric.Filename,
				"error", err,
			)
			return 0
		}
	}
	slog.Info("Published message for file",
		"filename", metric.Filename,
		"size", fmt.Sprintf("%.3f KB", metric.Size/1024),
	)
	return metric.Size

}
