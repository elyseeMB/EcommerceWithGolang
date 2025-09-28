package api

import (
	"EcommerceWithGolang/internal/application"
	"EcommerceWithGolang/internal/domain/entity"
	"context"

	"github.com/gin-gonic/gin"
)

type TaxonomyHandler struct {
	taxonomyApp application.TaxonomyAppInterface
}

func NewTaxonomyHandler(app application.TaxonomyAppInterface) *TaxonomyHandler {
	return &TaxonomyHandler{
		taxonomyApp: app,
	}
}

func (h *TaxonomyHandler) Create(c *gin.Context) {

	h.taxonomyApp.Create(context.TODO(), &entity.Taxonomy{
		Name: "taxonomy_1",
	})

	c.JSON(200, gin.H{"message": "taxonomy created"})
}
