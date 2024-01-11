package db

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
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

// CreateUser creates a new user in the database.
func (c *Client) CreateUser(name, email, password, address string) error {
	if _, err := c.db.Exec(
		fmt.Sprintf(
			"INSERT INTO users (id, name, email, password, address) VALUES ('%s', '%s', '%s', '%s', '%s')",
			uuid.NewString(),
			name,
			email,
			password,
			address,
		),
	); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
