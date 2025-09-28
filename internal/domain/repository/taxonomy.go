package repository

import (
	"EcommerceWithGolang/internal/domain/entity"
	"context"
)

type TaxonomyRespositoryI interface {
	List(context.Context) ([]*entity.Taxonomy, error)
	Create(context.Context, *entity.Taxonomy) error
	// GetById(id int) (*entity.Taxonomy, error)
}
