package database

import "embed"

var (
	//go:embed migrations/*.sql
	Migrations embed.FS
)
