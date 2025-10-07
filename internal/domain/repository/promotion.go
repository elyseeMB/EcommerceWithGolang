package repository

import (
	"EcommerceWithGolang/internal/domain/entity"
	"context"
)

type PromotionRepositoryI interface {
	Create(context.Context, *entity.Promotion) error
}
