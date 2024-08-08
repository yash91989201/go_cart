-- name: InsertUser :one
INSERT INTO "user" (id, created_at, name, email, password)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUsers :many
SELECT * FROM "user";

-- name: GetUserById :one
SELECT * FROM "user" WHERE id=$1;

-- name: GetUserByEmail :one
SELECT * FROM "user" WHERE email=$1;

-- name: GetUserWithSession :one
SELECT * FROM "user" INNER JOIN session ON "user".id = session.user_id;
