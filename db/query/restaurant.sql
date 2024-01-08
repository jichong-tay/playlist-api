-- name: CreateRestaurant :one
INSERT INTO restaurants (
  name,
  description,
  location,
  cuisine,
  image_url
) 
VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetRestaurant :one
SELECT * FROM restaurants
WHERE id = $1 LIMIT 1;

-- name: ListRestaurants :many
SELECT * FROM restaurants
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateRestaurant :one
UPDATE restaurants
SET 
  name = $2,
  description = $3,
  location = $4,
  cuisine = $5,
  image_url = $6
WHERE 
  id = $1
RETURNING *;

-- name: DeleteRestaurant :exec
DELETE FROM restaurants
WHERE id = $1;