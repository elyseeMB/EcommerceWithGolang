package persistence

import (
	"EcommerceWithGolang/internal/domain/entity"
	"EcommerceWithGolang/internal/infrastructure/database"
	"EcommerceWithGolang/internal/infrastructure/utils"
	"context"
	"database/sql"
	"fmt"
)

type TaxonomyRespository struct {
	client    *sql.DB
	tableName string
}

func NewTaxonomyRepository(ctx context.Context, cfg *database.AppConfig) *TaxonomyRespository {
	db := cfg.InitDatabase()

	query := `CREATE TABLE IF NOT EXISTS taxonomies (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL
	)`

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}

	return &TaxonomyRespository{
		client:    db,
		tableName: "taxonomies",
	}
}

func (r *TaxonomyRespository) Create(ctx context.Context, taxonomy *entity.Taxonomy) error {

	taxonomy.Id = utils.GenerateID()

	query := fmt.Sprintf("INSERT INTO %s (id, name) VALUES (?, ?)", r.tableName)

	fmt.Printf("value :%v", r.tableName)

	_, err := r.client.ExecContext(ctx, query, taxonomy.Id, taxonomy.Name)

	if err != nil {
		return fmt.Errorf("failed to insert: %w", err)
	}

	return nil

}
