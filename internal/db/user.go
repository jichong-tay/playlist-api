package db

import (
	"fmt"
)

// CreateUser creates a new user in the database.
func (c *Client) CreateUser(name, email, hashedPassword, address string) error {
	if _, err := c.db.Exec(
		fmt.Sprintf(
			"INSERT INTO users (username, email, password_hash, address) VALUES ('%s', '%s', '%s', '%s')",
			name,
			email,
			hashedPassword,
			address,
		),
	); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
