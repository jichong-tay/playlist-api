
-- name: ListPlaylist_DishesByPlaylistID :many
SELECT * 
FROM playlist_dishes
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

-- name: ListPlaylistPublicAndCategoryAll :many
SELECT *
FROM playlists
WHERE is_public = true;

-- name: ListPlaylistsByUserID :many
SELECT
    playlists.id,
    playlists.name,
    playlists.description,
    playlists.image_url,
    playlists.is_public,
    playlists.delivery_day,
    playlists.category,
    playlists.created_at,
    playlists.added_at
FROM
    playlists
JOIN
    user_playlists ON playlists.id = user_playlists.playlist_id
WHERE
    user_playlists.user_id = $1
LIMIT $2
OFFSET $3;

-- name: ListPlaylistsByUserIDAll :many
SELECT
    playlists.id,
    playlists.name,
    playlists.description,
    playlists.image_url,
    playlists.is_public,
    playlists.delivery_day,
    playlists.category,
    playlists.created_at,
    playlists.added_at
FROM
    playlists
JOIN
    user_playlists ON playlists.id = user_playlists.playlist_id
WHERE
    user_playlists.user_id = $1;

-- name: ListPlaylistByCategory :many
SELECT *
FROM
    playlists
WHERE
    LOWER(category) = LOWER($1);

-- name: ListStatusByPlaylistID :one
SELECT status
FROM
   user_playlists
WHERE
    user_playlists.playlist_id = $1;