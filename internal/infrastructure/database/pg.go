package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var ConnInstance *pgx.Conn

func Connection() *pgx.Conn {
	url := GetURLConnectionDB()
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	ConnInstance = conn

	return ConnInstance
}

func GetURLConnectionDB() string {
	cfg := Load()
	return fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", cfg.DBUser, cfg.DBPassword, cfg.DBName)
}
