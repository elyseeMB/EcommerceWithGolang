package repository

import (
	"EcommerceWithGolang/internal/domain/entity"
	"context"
)

type ItemRepositoryI interface {
	Create(context.Context, *entity.Item) error
}
