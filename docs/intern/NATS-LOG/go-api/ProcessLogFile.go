package main

import (
	"encoding/json"
	"fmt"
	"log"
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
		log.Printf("Error marshaling json for %s: %v", metric.Filename, err)
		return 0
	}

	subject := fmt.Sprintf("%s.%s", subjectPrefix, metric.Filename)

	if nc.IsConnected() {
		err = nc.Publish(subject, jsonData)
		if err != nil {
			log.Printf("Error publishing message for file %s: %v", metric.Filename, err)
			return 0
		}
	}
	if nc.Status() != nats.CONNECTED {
		log.Printf("Skipping publish for %s: NATS not connected (status: %v)", metric.Filename, nc.Status())
		return 0
	} else {
		log.Printf("Published message for file name:%s size:%.3f KB", metric.Filename, metric.Size/1024)
		return metric.Size
	}

}
