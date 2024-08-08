-- name: CreateSession :one
INSERT INTO session (id, expires_at, user_id)
VALUES($1, $2, $3)
RETURNING id;
