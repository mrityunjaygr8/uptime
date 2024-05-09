package server

import (
	"log/slog"
	"net/http"
)

func handlePingPong(logger *slog.Logger) http.Handler {
	type response struct {
		Response string `json:"response"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = encode(w, r, http.StatusOK, response{Response: "pong"})
	})
}
