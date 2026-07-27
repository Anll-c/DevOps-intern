package main

import (
	"log/slog"
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
			slog.Error("NATS disconnected", "error", err)
		}),

		nats.ReconnectHandler(func(nc *nats.Conn) {
			slog.Info("NATS reconnected", "url", nc.ConnectedUrl())
		}),

		nats.ClosedHandler(func(nc *nats.Conn) {
			slog.Info("NATS connection closed permanently")
		}),
	)

	return nc, err
}
