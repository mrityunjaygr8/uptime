package server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/urfave/negroni"
)

func NewServer(logger *slog.Logger, db *pgx.Conn) http.Handler {
	mux := chi.NewMux()

	addRoutes(mux, logger, db)

	n := negroni.Classic()
	n.UseHandler(mux)

	return n
}

type ServiceError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Extra   any    `json:"extra"`
}
