package api

import (
	"EcommerceWithGolang/internal/application"
	"EcommerceWithGolang/internal/domain/entity"
	"context"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	itemApp application.ItemAppInterface
}

func NewItemHandler(itemApp application.ItemAppInterface) *ItemHandler {
	return &ItemHandler{
		itemApp: itemApp,
	}
}

func (h *ItemHandler) Create(c *gin.Context) {

	if err := h.itemApp.Create(context.TODO(), &entity.Item{
		Taxonomy: 1,
		SellerID: 1,
		CardID:   1,
		Price:    45,
		ItemType: 1,
		Quantity: 4,
	}); err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{"message": "item created"})

}

func (h *ItemHandler) ListByCartId(c *gin.Context) {

	c.JSON(404, gin.H{"message": "ok"})
}
