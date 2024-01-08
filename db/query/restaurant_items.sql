-- name: CreateRestaurant_Item :one
INSERT INTO restaurant_items (
  restaurant_id,
  is_available,
  name,
  description,
  price,
  cuisine,
  image_url
) 
VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetRestaurant_Item :one
SELECT * FROM restaurant_items
WHERE id = $1 LIMIT 1;

-- name: ListRestaurant_Items :many
SELECT * FROM restaurant_items
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateRestaurant_Item :one
UPDATE restaurant_items
SET 
  restaurant_id = $2,
  is_available = $3,
  name = $4,
  description = $5,
  price = $6,
  cuisine = $7,
  image_url = $8
WHERE 
  id = $1
RETURNING *;

-- name: DeleteRestaurant_Item :exec
DELETE FROM restaurant_items
WHERE id = $1;