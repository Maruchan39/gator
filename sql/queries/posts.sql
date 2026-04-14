-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1,
    $2,
    $3,
    sqlc.narg('title'),
    $4,
    sqlc.narg('description'),
    sqlc.narg('published_at'),
    $5
)
RETURNING *;

-- name: GetPosts :many
SELECT * FROM posts
ORDER BY published_at ASC
LIMIT $1;