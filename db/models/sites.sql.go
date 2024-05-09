// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: sites.sql

package models

import (
	"context"
)

const listSites = `-- name: ListSites :many
SELECT id, url FROM SITES
ORDER BY id
`

func (q *Queries) ListSites(ctx context.Context) ([]Site, error) {
	rows, err := q.db.Query(ctx, listSites)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Site
	for rows.Next() {
		var i Site
		if err := rows.Scan(&i.ID, &i.Url); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}