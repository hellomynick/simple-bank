package api

import (
	"fmt"
	"os"
)

var DbDriver = os.Getenv("DB_DRIVER")
var DbSource = os.Getenv("DB_SOURCE")

type DatabaseConfig struct {
	Driver           string
	ConnectionString string
}

func LoadDatabaseConfig() (*DatabaseConfig, error) {
	if DbDriver == "" {
		return nil, fmt.Errorf("can not load db driver")
	}

	return &DatabaseConfig{
		Driver:           DbDriver,
		ConnectionString: DbSource,
	}, nil
}
