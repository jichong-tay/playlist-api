package db

import (
	"fmt"

	"github.com/google/uuid"
)

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
