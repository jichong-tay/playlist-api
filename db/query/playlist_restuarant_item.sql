-- name: CreatePlaylist_Restaurant_Item :one
INSERT INTO playlist_restaurant_items (
  playlist_id,
  restaurant_item_id,
  restaurant_item_quantity,
  added_at
) 
VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetPlaylist_Restaurant_Item :one
SELECT * FROM playlist_restaurant_items
WHERE id = $1 LIMIT 1;

-- name: ListPlaylist_Restaurant_Items :many
SELECT * FROM playlist_restaurant_items
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdatePlaylist_Restaurant_Item :one
UPDATE playlist_restaurant_items
SET 
  playlist_id = $2,
  restaurant_item_id = $3,
  restaurant_item_quantity = $4,
  added_at = $5
WHERE 
  id = $1
RETURNING *;

-- name: DeletePlaylist_Restaurant_Item :exec
DELETE FROM playlist_restaurant_items
WHERE id = $1;