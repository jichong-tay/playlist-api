package config

import "os"

// Config is the configuration for the application.
type Config struct {
	DSN string
}

// Init initializes the configuration for the application.
func Init() Config {
	dsn, hasDSN := os.LookupEnv("DSN")
	if !hasDSN {
		dsn = "postgres://root:secret@localhost:5432/foodpanda-playlist?sslmode=disable"
	}

	return Config{
		DSN: dsn,
	}
}
