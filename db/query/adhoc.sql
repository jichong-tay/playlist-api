
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

-- name: GetUserPlaylistByPlaylistID :one
SELECT * FROM user_playlists
WHERE playlist_id = $1 LIMIT 1;

-- name: UpdateStatusForUser_Playlist :many
UPDATE user_playlists
SET 
  status = $3
WHERE 
  user_id = $1 AND playlist_id = $2
RETURNING *;

-- name: SearchDishes :many
SELECT d.id AS dish_id,
       d.name AS dish_name,
       d.description AS dish_description,
       d.price AS dish_price,
       d.image_url AS dish_imageURL,
       r.name AS restaurant_name,
       r.id AS restaurant_id
FROM dishes d
JOIN restaurants r ON d.restaurant_id = r.id
WHERE d.name ILIKE '%'||$1||'%' OR d.description ILIKE '%'||$1||'%';

-- name: ListSearchesByUserID :many
SELECT *
FROM (
    SELECT DISTINCT ON (keyword) *
    FROM searches
    WHERE user_id = $1
    ORDER BY keyword, id DESC
) AS subquery
ORDER BY id DESC
LIMIT $2
OFFSET $3;

-- name: DeleteSearchByKeyword :many
DELETE FROM searches
WHERE keyword = $1 AND user_id = $2
RETURNING *;

-- name: DeletePlaylistDishes :many
DELETE FROM playlist_dishes
WHERE playlist_id = $1
RETURNING *;

-- name: UpdateUser_PlaylistDelivery :one
UPDATE user_playlists
SET 
  delivery_day = $3,
  delivery_time = $4,
  status = $5
WHERE 
  user_id = $1 AND playlist_id = $2
RETURNING *;

-- name: ListDishesByCuisine :many
SELECT * FROM dishes
WHERE cuisine ILIKE '%'||$1||'%' OR description ILIKE '%'||$1||'%';