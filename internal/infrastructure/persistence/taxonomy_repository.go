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

func (r *TaxonomyRespository) List(ctx context.Context) ([]*entity.Taxonomy, error) {
	taxonomies := []*entity.Taxonomy{}

	query := fmt.Sprintf("SELECT * FROM %s", r.tableName)

	rows, err := r.client.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error getting columns: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		t := &entity.Taxonomy{}
		if err := rows.Scan(&t.Id, &t.Name); err != nil {
			return nil, fmt.Errorf("error scanning row %w", err)
		}
		taxonomies = append(taxonomies, t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error : %w", err)
	}

	return taxonomies, nil

}

func (r *TaxonomyRespository) Create(ctx context.Context, taxonomy *entity.Taxonomy) error {

	taxonomy.Id = utils.GenerateID()

	query := fmt.Sprintf("INSERT INTO %s (id, name) VALUES (?, ?)", r.tableName)

	_, err := r.client.ExecContext(ctx, query, taxonomy.Id, taxonomy.Name)

	if err != nil {
		return fmt.Errorf("failed to insert: %w", err)
	}

	return nil
}

func (r *TaxonomyRespository) GetById(ctx context.Context, id int) (*entity.Taxonomy, error) {

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", r.tableName)

	taxonomy := &entity.Taxonomy{}

	rows := r.client.QueryRowContext(ctx, query, id)

	err := rows.Scan(&taxonomy.Id, &taxonomy.Name)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("Error %w", err)
	}
	return taxonomy, nil

}
