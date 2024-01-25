// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: user.sql

package db

import (
	"context"

	null "gopkg.in/guregu/null.v4"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username,
  email,
  password_hash,
  address
) 
VALUES (
  $1, $2, $3, $4
)
RETURNING id, username, email, password_hash, address
`

type CreateUserParams struct {
	Username     string      `json:"username"`
	Email        string      `json:"email"`
	PasswordHash string      `json:"password_hash"`
	Address      null.String `json:"address"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.PasswordHash,
		arg.Address,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.Address,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, username, email, password_hash, address FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.Address,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, username, email, password_hash, address FROM users
WHERE id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListUsersParams struct {
	ID     int64 `json:"id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.ID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.PasswordHash,
			&i.Address,
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

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET 
  username = $2,
  email = $3,
  password_hash = $4,
  address = $5
WHERE 
  id = $1
RETURNING id, username, email, password_hash, address
`

type UpdateUserParams struct {
	ID           int64       `json:"id"`
	Username     string      `json:"username"`
	Email        string      `json:"email"`
	PasswordHash string      `json:"password_hash"`
	Address      null.String `json:"address"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.Username,
		arg.Email,
		arg.PasswordHash,
		arg.Address,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.Address,
	)
	return i, err
}
