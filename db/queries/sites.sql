-- name: ListSites :many
SELECT * FROM sites
ORDER BY id;

-- name: CreateSite :one
INSERT INTO sites (
  url
) VALUES (
  $1
)
RETURNING *;
