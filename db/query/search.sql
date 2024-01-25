-- name: CreateSearch :one
INSERT INTO searches (
  user_id,
  keyword
) 
VALUES (
  $1, $2
)
RETURNING *;

-- name: GetSearch :one
SELECT * FROM searches
WHERE id = $1 LIMIT 1;

-- name: ListSearches :many
SELECT * FROM searches
WHERE id = $1
ORDER BY id DESC
LIMIT $2
OFFSET $3;

-- -- name: UpdateSearch :one
-- UPDATE searches
-- SET 
--   user_id = $2,
--   keyword = $3
-- WHERE 
--   id = $1
-- RETURNING *;

-- -- name: DeleteSearch :exec
-- DELETE FROM searches
-- WHERE id = $1;