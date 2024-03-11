// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: adhoc.sql

package db

import (
	"context"

	null "gopkg.in/guregu/null.v4"
)

const getUserPlaylistByPlaylistID = `-- name: GetUserPlaylistByPlaylistID :one
SELECT id, user_id, playlist_id, delivery_day, delivery_time, status FROM user_playlists
WHERE playlist_id = $1 LIMIT 1
`

func (q *Queries) GetUserPlaylistByPlaylistID(ctx context.Context, playlistID int64) (UserPlaylist, error) {
	row := q.db.QueryRowContext(ctx, getUserPlaylistByPlaylistID, playlistID)
	var i UserPlaylist
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.PlaylistID,
		&i.DeliveryDay,
		&i.DeliveryTime,
		&i.Status,
	)
	return i, err
}

const listPlaylistByCategory = `-- name: ListPlaylistByCategory :many
SELECT id, name, description, image_url, is_public, delivery_day, category, created_at, added_at
FROM
    playlists
WHERE
    LOWER(category) = LOWER($1)
`

func (q *Queries) ListPlaylistByCategory(ctx context.Context, lower string) ([]Playlist, error) {
	rows, err := q.db.QueryContext(ctx, listPlaylistByCategory, lower)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Playlist{}
	for rows.Next() {
		var i Playlist
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.ImageUrl,
			&i.IsPublic,
			&i.DeliveryDay,
			&i.Category,
			&i.CreatedAt,
			&i.AddedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPlaylistPublicAndCategory = `-- name: ListPlaylistPublicAndCategory :many
SELECT id, name, description, image_url, is_public, delivery_day, category, created_at, added_at
FROM playlists
WHERE is_public = true
LIMIT $1
OFFSET $2
`

type ListPlaylistPublicAndCategoryParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPlaylistPublicAndCategory(ctx context.Context, arg ListPlaylistPublicAndCategoryParams) ([]Playlist, error) {
	rows, err := q.db.QueryContext(ctx, listPlaylistPublicAndCategory, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Playlist{}
	for rows.Next() {
		var i Playlist
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.ImageUrl,
			&i.IsPublic,
			&i.DeliveryDay,
			&i.Category,
			&i.CreatedAt,
			&i.AddedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPlaylistPublicAndCategoryAll = `-- name: ListPlaylistPublicAndCategoryAll :many
SELECT id, name, description, image_url, is_public, delivery_day, category, created_at, added_at
FROM playlists
WHERE is_public = true
`

func (q *Queries) ListPlaylistPublicAndCategoryAll(ctx context.Context) ([]Playlist, error) {
	rows, err := q.db.QueryContext(ctx, listPlaylistPublicAndCategoryAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Playlist{}
	for rows.Next() {
		var i Playlist
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.ImageUrl,
			&i.IsPublic,
			&i.DeliveryDay,
			&i.Category,
			&i.CreatedAt,
			&i.AddedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPlaylist_DishesByPlaylistID = `-- name: ListPlaylist_DishesByPlaylistID :many
SELECT id, order_id, playlist_id, dish_id, dish_quantity, created_at, added_at 
FROM playlist_dishes
WHERE playlist_id = $1
ORDER BY id
`

func (q *Queries) ListPlaylist_DishesByPlaylistID(ctx context.Context, playlistID int64) ([]PlaylistDish, error) {
	rows, err := q.db.QueryContext(ctx, listPlaylist_DishesByPlaylistID, playlistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []PlaylistDish{}
	for rows.Next() {
		var i PlaylistDish
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.PlaylistID,
			&i.DishID,
			&i.DishQuantity,
			&i.CreatedAt,
			&i.AddedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPlaylistsByUserID = `-- name: ListPlaylistsByUserID :many
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
OFFSET $3
`

type ListPlaylistsByUserIDParams struct {
	UserID int64 `json:"user_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPlaylistsByUserID(ctx context.Context, arg ListPlaylistsByUserIDParams) ([]Playlist, error) {
	rows, err := q.db.QueryContext(ctx, listPlaylistsByUserID, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Playlist{}
	for rows.Next() {
		var i Playlist
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.ImageUrl,
			&i.IsPublic,
			&i.DeliveryDay,
			&i.Category,
			&i.CreatedAt,
			&i.AddedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPlaylistsByUserIDAll = `-- name: ListPlaylistsByUserIDAll :many
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
`

func (q *Queries) ListPlaylistsByUserIDAll(ctx context.Context, userID int64) ([]Playlist, error) {
	rows, err := q.db.QueryContext(ctx, listPlaylistsByUserIDAll, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Playlist{}
	for rows.Next() {
		var i Playlist
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.ImageUrl,
			&i.IsPublic,
			&i.DeliveryDay,
			&i.Category,
			&i.CreatedAt,
			&i.AddedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRestaurantNameByDishID = `-- name: ListRestaurantNameByDishID :one
SELECT restaurants.name
FROM dishes
JOIN restaurants ON dishes.restaurant_id = restaurants.id
WHERE dishes.id = $1
`

func (q *Queries) ListRestaurantNameByDishID(ctx context.Context, id int64) (string, error) {
	row := q.db.QueryRowContext(ctx, listRestaurantNameByDishID, id)
	var name string
	err := row.Scan(&name)
	return name, err
}

const listStatusByPlaylistID = `-- name: ListStatusByPlaylistID :one
SELECT status
FROM
   user_playlists
WHERE
    user_playlists.playlist_id = $1
`

func (q *Queries) ListStatusByPlaylistID(ctx context.Context, playlistID int64) (null.String, error) {
	row := q.db.QueryRowContext(ctx, listStatusByPlaylistID, playlistID)
	var status null.String
	err := row.Scan(&status)
	return status, err
}

const updateStatusForUser_Playlist = `-- name: UpdateStatusForUser_Playlist :many
UPDATE user_playlists
SET 
  status = $3
WHERE 
  user_id = $1 AND playlist_id = $2
RETURNING id, user_id, playlist_id, delivery_day, delivery_time, status
`

type UpdateStatusForUser_PlaylistParams struct {
	UserID     int64       `json:"user_id"`
	PlaylistID int64       `json:"playlist_id"`
	Status     null.String `json:"status"`
}

func (q *Queries) UpdateStatusForUser_Playlist(ctx context.Context, arg UpdateStatusForUser_PlaylistParams) ([]UserPlaylist, error) {
	rows, err := q.db.QueryContext(ctx, updateStatusForUser_Playlist, arg.UserID, arg.PlaylistID, arg.Status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserPlaylist{}
	for rows.Next() {
		var i UserPlaylist
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.PlaylistID,
			&i.DeliveryDay,
			&i.DeliveryTime,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
