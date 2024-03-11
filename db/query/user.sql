-- name: CreateUser :one
INSERT INTO users (
  username,
  email,
  password_hash,
  address,
  uuid
) 
VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
WHERE id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateUser :one
UPDATE users
SET 
  username = $2,
  email = $3,
  password_hash = $4,
  address = $5,
  uuid = $6
WHERE 
  id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;