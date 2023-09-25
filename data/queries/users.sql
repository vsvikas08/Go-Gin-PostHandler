-- name: CreateUser :one
INSERT INTO users(email, id)
VALUES ($1, $2)
RETURNING *;