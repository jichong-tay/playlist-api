
-- name: ListPlaylist_DishesByPlaylistID :many
SELECT * FROM playlist_dishes
WHERE playlist_id = $1
ORDER BY id;

-- name: ListRestaurantNameByDishID :one
SELECT r.name AS restaurant_name
FROM dishes d
JOIN restaurants r ON d.restaurant_id = r.id
WHERE d.id = $1;