package server

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
)

func addRoutes(mux *chi.Mux, logger *slog.Logger) {
	mux.Handle("/api/v1/ping-pong", handlePingPong(logger))
	mux.Handle("/api/v1/ping/*", handlePing(logger))
}
