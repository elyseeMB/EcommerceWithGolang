package persistence

import (
	"EcommerceWithGolang/internal/domain/repository"
	"EcommerceWithGolang/internal/infrastructure/database"
	"context"
)

type Repositories struct {
	Taxonomy repository.TaxonomyRespositoryI
	Item     repository.ItemRepositoryI
}

func NewRepositories(cfg database.AppConfig) (*Repositories, error) {
	return &Repositories{
		Taxonomy: NewTaxonomyRepository(context.TODO(), &cfg),
		Item:     NewItemRepository(context.TODO(), &cfg),
	}, nil
}
