
-- name: ListPlaylist_DishesByPlaylistID :many
SELECT * FROM playlist_dishes
WHERE playlist_id = $1
ORDER BY id;

-- name: ListRestaurantNameByDishID :one
SELECT restaurants.name
FROM dishes
JOIN restaurants ON dishes.restaurant_id = restaurants.id
WHERE dishes.id = $1;

-- name: ListPlaylistPublicAndCategory :many
SELECT *
FROM playlists
WHERE is_public = true
LIMIT $1
OFFSET $2;

