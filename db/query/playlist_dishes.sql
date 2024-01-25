-- name: CreatePlaylist_Dish :one
INSERT INTO playlist_dishes (
  order_id,
  playlist_id,
  dish_id,
  dish_quantity
) 
VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetPlaylist_Dish :one
SELECT * FROM playlist_dishes
WHERE id = $1 LIMIT 1;

-- name: ListPlaylist_Dishes :many
SELECT * FROM playlist_dishes
WHERE id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdatePlaylist_Dish :one
UPDATE playlist_dishes
SET 
  order_id = $2,
  playlist_id = $3,
  dish_id = $4,
  dish_quantity = $5,
  added_at = $6
WHERE 
  id = $1
RETURNING *;

-- name: DeletePlaylist_Dish :exec
DELETE FROM playlist_dishes
WHERE id = $1;