package api

import (
	"EcommerceWithGolang/internal/application"
	"EcommerceWithGolang/internal/domain/entity"
	"context"
	"strconv"

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

func (h *TaxonomyHandler) List(c *gin.Context) {
	taxonomies, err := h.taxonomyApp.List(context.TODO())
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": taxonomies,
	})

}

func (h *TaxonomyHandler) Create(c *gin.Context) {

	if err := h.taxonomyApp.Create(context.TODO(), &entity.Taxonomy{
		Name: "taxonomy_1",
	}); err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{"message": "taxonomy created"})
}

func (h *TaxonomyHandler) GetById(c *gin.Context) {

	paramId := c.Param("id")

	if paramId == "" {
		c.JSON(400, gin.H{"message": "Category Id is required"})
		return
	}

	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID formt"})
		return
	}

	taxonomy, err := h.taxonomyApp.GetById(context.TODO(), id)

	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": taxonomy,
	})

}
