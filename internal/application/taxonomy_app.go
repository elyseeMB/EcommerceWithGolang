package application

import (
	"EcommerceWithGolang/internal/domain/entity"
	"EcommerceWithGolang/internal/domain/repository"
	"context"
)

var _ TaxonomyAppInterface = &taxonomyApp{}

type taxonomyApp struct {
	taxonomyRepo repository.TaxonomyRespositoryI
}

type TaxonomyAppInterface interface {
	Create(context.Context, *entity.Taxonomy) error
}

func NewTaxonomyApp(r repository.TaxonomyRespositoryI) *taxonomyApp {
	return &taxonomyApp{
		taxonomyRepo: r,
	}
}

func (app *taxonomyApp) Create(ctx context.Context, taxonomy *entity.Taxonomy) error {
	return app.taxonomyRepo.Create(ctx, taxonomy)
}
