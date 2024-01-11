package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// Client is a wrapper around the database connection.
type Client struct {
	db *sql.DB
}

// Init creates a new database connection.
func Init(dsn string) (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Client{
		db,
	}, nil
}

// Close closes the database connection.
func (c *Client) Close() error {
	return c.db.Close()
}
