-- name: CreateUser_Playlist :one
INSERT INTO user_playlists (
  user_id,
  playlist_id,
  delivery_day,
  delivery_time,
  status
) 
VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetUser_Playlist :one
SELECT * FROM user_playlists
WHERE id = $1 LIMIT 1;

-- name: ListUser_Playlists :many
SELECT * FROM user_playlists
WHERE user_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateUser_Playlist :one
UPDATE user_playlists
SET 
  user_id = $2,
  playlist_id = $3,
  delivery_day = $4,
  delivery_time = $5,
  status = $6
WHERE 
  id = $1
RETURNING *;

-- name: DeleteUser_Playlist :exec
DELETE FROM user_playlists
WHERE id = $1;