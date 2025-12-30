package configs

import (
	"fmt"
	"os"
)

type PostgresConfig struct {
	Driver           string
	ConnectionString string
}

func LoadPostgresConfig() (*PostgresConfig, error) {
	driver := os.Getenv("DB_DRIVER")
	source := os.Getenv("DB_SOURCE")

	if driver == "" || source == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	return &PostgresConfig{
		Driver:           driver,
		ConnectionString: source,
	}, nil
}
