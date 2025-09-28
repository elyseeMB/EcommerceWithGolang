package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type AppConfig struct {
	DriverName     string
	DataSourceName string
}

func (config *AppConfig) InitDatabase() *sql.DB {
	db, err := sql.Open(config.DriverName, config.DataSourceName)
	if err != nil {
		panic(err)
	}

	return db
}
