package server

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func addRoutes(mux *chi.Mux, logger *slog.Logger, db *pgx.Conn) {
	mux.Get("/api/v1/ping-pong", handlePingPong(logger))
	mux.Get("/api/v1/ping/*", handlePing(logger))
	mux.Post("/api/v1/sites/", handleSiteCreate(logger, db))
	mux.Get("/api/v1/sites/", handleSiteList(logger, db))
}
