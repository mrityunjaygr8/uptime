// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: sites.sql

package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSite = `-- name: CreateSite :one
INSERT INTO sites (
  url
) VALUES (
  $1
)
RETURNING id, url, created_at
`

func (q *Queries) CreateSite(ctx context.Context, url string) (Site, error) {
	row := q.db.QueryRow(ctx, createSite, url)
	var i Site
	err := row.Scan(&i.ID, &i.Url, &i.CreatedAt)
	return i, err
}

const listSites = `-- name: ListSites :many
SELECT id, url, created_at, count(*) OVER() AS full_count
FROM sites
ORDER BY id
LIMIT $2::int OFFSET $1::int
`

type ListSitesParams struct {
	PageOffset int32
	PageSize   int32
}

type ListSitesRow struct {
	ID        int64
	Url       string
	CreatedAt pgtype.Timestamptz
	FullCount int64
}

func (q *Queries) ListSites(ctx context.Context, arg ListSitesParams) ([]ListSitesRow, error) {
	rows, err := q.db.Query(ctx, listSites, arg.PageOffset, arg.PageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListSitesRow
	for rows.Next() {
		var i ListSitesRow
		if err := rows.Scan(
			&i.ID,
			&i.Url,
			&i.CreatedAt,
			&i.FullCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
