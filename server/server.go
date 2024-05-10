package server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

func NewServer(logger *slog.Logger, db *pgx.Conn) http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	addRoutes(mux, logger, db)

	return mux
}

type ServiceError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Extra   any    `json:"extra"`
}

type PaginatedResponse[T any] struct {
	Data     []T `json:"data"`
	Page     int `json:"page"`
	Count    int `json:"count"`
	Pages    int `json:"pages"`
	PageSize int `json:"page_size"`
}
