-- name: GetUserByID :one
SELECT *
FROM users u
WHERE u.id = $1;

-- name: FindManyUsers :many
SELECT id, name, email, password, created_at, updated_at
FROM users;
