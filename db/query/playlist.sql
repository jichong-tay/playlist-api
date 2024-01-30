-- name: CreatePlaylist :one
INSERT INTO playlists (
  name,
  description,
  image_url,
  is_public,
  delivery_day,
  category
) 
VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetPlaylist :one
SELECT * FROM playlists
WHERE id = $1 LIMIT 1;

-- name: ListPlaylists :many
SELECT * FROM playlists
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdatePlaylist :one
UPDATE playlists
SET 
  name = $2,
  description = $3,
  image_url = $4,
  is_public = $5,
  delivery_day = $6,
  category = $7,
  added_at = $8
WHERE 
  id = $1
RETURNING *;

-- name: DeletePlaylist :exec
DELETE FROM playlists
WHERE id = $1;