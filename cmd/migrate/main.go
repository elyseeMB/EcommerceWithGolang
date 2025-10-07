package main

import (
	"EcommerceWithGolang/internal/infrastructure/database"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func main() {
	url := database.GetURLConnectionDB()

	d, err := iofs.New(database.Migrations, "migrations")
	if err != nil {
		log.Fatalf("cannot create iofs driver: %v", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, url)
	if err != nil {
		log.Fatalf("cannot create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("cannot run migrations: %v", err)
	}

	fmt.Println("Migrations applied successfully!")
}
