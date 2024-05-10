package server

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/jackc/pgx/v5"
	"github.com/mrityunjaygr8/uptime/db/models"
)

type SiteResponse struct {
	ID         int       `json:"id"`
	URL        string    `json:"url"`
	Created_at time.Time `json:"created_at"`
}

func (s *SiteResponse) fromModelSite(m models.ListSitesRow) {
	s.Created_at = m.CreatedAt.Time
	s.ID = int(m.ID)
	s.URL = m.Url
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
	type SiteListQP struct {
		Page     int `schema:"page,default:1"`
		PageSize int `schema:"page_size,default:20"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var decoder = schema.NewDecoder()
		var qp SiteListQP
		err := decoder.Decode(&qp, r.URL.Query())
		if err != nil {
			serr, ok := err.(*schema.ConversionError)
			if ok {
				logger.Info("not conversion error")
			} else {
				logger.Info("is conversion error")
				logger.Info("key", "key", serr)
			}
			// logger.Info("is multi error", "error", serr.Error())
			_ = encode(w, r, http.StatusInternalServerError, ServiceError{
				Message: err.Error(),
				Code:    3,
				Extra:   nil,
			})
			return
		}
		filters := models.ListSitesParams{PageOffset: int32((qp.Page - 1) * qp.PageSize), PageSize: int32(qp.PageSize)}
		sites, err := queries.ListSites(r.Context(), filters)
		if err != nil {
			_ = encode(w, r, http.StatusInternalServerError, ServiceError{
				Message: err.Error(),
				Code:    3,
				Extra:   nil,
			})
			return
		}

		var sitesResponse PaginatedResponse[SiteResponse]
		for _, site := range sites {
			var tmpSite SiteResponse
			tmpSite.fromModelSite(site)
			sitesResponse.Data = append(sitesResponse.Data, tmpSite)
		}
		sitesResponse.Page = 1
		sitesResponse.Count = len(sites)
		sitesResponse.Pages = 1
		sitesResponse.PageSize = 20
		logger.Info("testing stugg", "sites", sitesResponse)
		_ = encode(w, r, http.StatusCreated, &sitesResponse)
	})
}
