package server

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func handlePing(logger *slog.Logger) http.HandlerFunc {
	type response struct {
		Up bool `json:"up"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := chi.URLParam(r, "*")
		logger.Info("recieved ping request", "url", url)
		if !strings.HasPrefix(url, "http:") && !strings.HasPrefix(url, "https:") {
			url = "https://" + url
		}

		logger.Info("getting final url", "url", url)
		// Make an HTTP request to check if it's up.
		req, err := http.NewRequest("GET", url, nil)
		logger.Info("Request created", "url", url)
		if err != nil {
			_ = encode(w, r, http.StatusBadRequest, ServiceError{
				Message: err.Error(),
				Code:    1,
				Extra:   nil,
			})
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			_ = encode(w, r, http.StatusBadRequest, ServiceError{
				Message: err.Error(),
				Code:    2,
				Extra:   nil,
			})
		}
		logger.Info("request sent and recieved", "url", url)
		resp.Body.Close()

		// 2xx and 3xx status codes are considered up
		up := resp.StatusCode < 400
		logger.Info("finished", "url", url)
		_ = encode(w, r, http.StatusOK, response{
			Up: up,
		})

	})
}
