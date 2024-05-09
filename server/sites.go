package server

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/mrityunjaygr8/uptime/db/models"
)

type SiteResponse struct {
	ID         int       `json:"id"`
	URL        string    `json:"url"`
	Created_at time.Time `json:"created_at"`
}

func (s *SiteResponse) fromModelSite(m models.Site) {
	s.Created_at = m.CreatedAt.Time
	s.ID = int(m.ID)
	s.URL = m.Url
}

type SiteListResponse struct {
	Data []SiteResponse `json:"data"`
}

func (s *SiteListResponse) fromModelSites(m []models.Site) {
	for _, site := range m {
		var tmp SiteResponse
		tmp.fromModelSite(site)
		s.Data = append(s.Data, tmp)
	}
}

func handleSiteCreate(logger *slog.Logger, db *pgx.Conn) http.HandlerFunc {
	queries := models.New(db)
	type SiteCreateParams struct {
		URL string `json:"url"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params, err := decode[SiteCreateParams](r)
		if err != nil {
			_ = encode(w, r, http.StatusInternalServerError, ServiceError{
				Message: err.Error(),
				Code:    3,
				Extra:   nil,
			})
		}

		site, err := queries.CreateSite(r.Context(), params.URL)
		if err != nil {
			_ = encode(w, r, http.StatusBadRequest, ServiceError{
				Message: err.Error(),
				Code:    4,
				Extra:   nil,
			})
		}

		_ = encode(w, r, http.StatusCreated, SiteResponse{
			ID:         int(site.ID),
			URL:        site.Url,
			Created_at: site.CreatedAt.Time,
		})
	})
}

func handleSiteList(logger *slog.Logger, db *pgx.Conn) http.HandlerFunc {
	queries := models.New(db)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sites, err := queries.ListSites(r.Context())
		if err != nil {
			_ = encode(w, r, http.StatusInternalServerError, ServiceError{
				Message: err.Error(),
				Code:    3,
				Extra:   nil,
			})
		}

		var sitesResponse SiteListResponse
		sitesResponse.fromModelSites(sites)
		logger.Info("testing stugg", "sites", sitesResponse)
		_ = encode(w, r, http.StatusCreated, &sitesResponse)
	})
}
