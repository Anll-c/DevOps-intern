package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func connectNATS(natsURL string) (*nats.Conn, error) {
	nc, err := nats.Connect(
		natsURL,
		nats.RetryOnFailedConnect(true),
		nats.MaxReconnects(-1),
		nats.ReconnectWait(2*time.Second),

		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			log.Printf("NATS connection lost: %v", err)
		}),

		nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Printf("NATS reconnected: %s", nc.ConnectedUrl())
		}),

		nats.ClosedHandler(func(nc *nats.Conn) {
			log.Println("NATS connection closed permanently")
		}),
	)

	return nc, err
}
