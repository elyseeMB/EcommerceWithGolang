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

func NewTaxonomyApp(r repository.TaxonomyRespositoryI) *taxonomyApp {
	return &taxonomyApp{
		taxonomyRepo: r,
	}
}

type TaxonomyAppInterface interface {
	List(context.Context) ([]*entity.Taxonomy, error)
	Create(context.Context, *entity.Taxonomy) error
}

func (app *taxonomyApp) List(ctx context.Context) ([]*entity.Taxonomy, error) {
	return app.taxonomyRepo.List(ctx)
}

func (app *taxonomyApp) Create(ctx context.Context, taxonomy *entity.Taxonomy) error {
	return app.taxonomyRepo.Create(ctx, taxonomy)
}
