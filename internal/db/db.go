package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Client is a wrapper around the database connection.
type Client struct {
	db *sql.DB
}

// Init creates a new database connection.
func Init(connectionString string) (*Client, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	return &Client{
		db,
	}, nil
}
