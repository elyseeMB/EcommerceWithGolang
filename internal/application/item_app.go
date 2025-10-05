package application

import (
	"EcommerceWithGolang/internal/domain/entity"
	"EcommerceWithGolang/internal/domain/repository"
	"context"
)

var _ ItemAppInterface = &ItemApp{}

type ItemAppInterface interface {
	Create(context.Context, *entity.Item) error
}

type ItemApp struct {
	itemRepo repository.ItemRepositoryI
}

func NewItemApp(itemRepo repository.ItemRepositoryI) *ItemApp {
	return &ItemApp{
		itemRepo: itemRepo,
	}
}

func (app *ItemApp) Create(ctx context.Context, item *entity.Item) error {
	return app.itemRepo.Create(ctx, item)
}
