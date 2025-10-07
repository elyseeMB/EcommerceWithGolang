package persistence

import (
	"EcommerceWithGolang/internal/domain/entity"
	"EcommerceWithGolang/internal/domain/repository"
	"EcommerceWithGolang/internal/infrastructure/database"
	"EcommerceWithGolang/internal/infrastructure/utils"
	"context"
	"fmt"
)

type ItemRespository struct {
	client    repository.DBExecutor
	tableName string
}

var _ repository.ItemRepositoryI = &ItemRespository{}

func NewItemRepository(ctx context.Context, cfg *database.AppConfig) *ItemRespository {

	db := cfg.InitDatabase()

	query := `CREATE TABLE IF NOT EXISTS items (
		id TEXT PRIMARY KEY,
		taxonomy_id TEXT,
		seller_id TEXT,
		card_id TEXT,
		price TEXT,
		quantity TEXT,
		item_type TEXT
	)`

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}

	return &ItemRespository{
		client:    db,
		tableName: "items",
	}
}

func (r *ItemRespository) Create(ctx context.Context, item *entity.Item) error {

	item.Id = utils.GenerateID()

	query := fmt.Sprintf("INSERT INTO %s (id, taxonomy_id, seller_id, card_id, price, quantity, item_type) VALUES (?, ?, ?, ?, ?, ?, ?)", r.tableName)

	_, err := r.client.ExecContext(ctx, query, item.Id, item.Taxonomy, item.CardID, item.SellerID, item.Price, item.ItemType, item.Quantity)

	if err != nil {
		return fmt.Errorf("failed to insert: %w", err)
	}

	return nil
}
