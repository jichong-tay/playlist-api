-- name: CreateUser :one
INSERT INTO users (
  username,
  email,
  password_hash,
  address
) 
VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET 
  username = $2,
  email = $3,
  password_hash = $4,
  address = $5
WHERE 
  id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;