-- name: ListSites :many
SELECT *, count(*) OVER() AS full_count
FROM sites
ORDER BY id
LIMIT sqlc.arg(page_size)::int OFFSET sqlc.arg(page_offset)::int;

-- name: CreateSite :one
INSERT INTO sites (
  url
) VALUES (
  $1
)
RETURNING *;
