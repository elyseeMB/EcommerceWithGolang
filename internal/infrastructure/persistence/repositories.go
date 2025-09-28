package persistence

import (
	"EcommerceWithGolang/internal/domain/repository"
	"EcommerceWithGolang/internal/infrastructure/database"
	"context"
)

type Repositories struct {
	Taxonomy repository.TaxonomyRespositoryI
}

func NewRepositories(cfg database.AppConfig) (*Repositories, error) {
	return &Repositories{
		Taxonomy: NewTaxonomyRepository(context.TODO(), &cfg),
	}, nil
}
