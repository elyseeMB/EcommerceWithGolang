package interfaces

import (
	"EcommerceWithGolang/internal/application"
	"EcommerceWithGolang/internal/infrastructure/database"
	"EcommerceWithGolang/internal/infrastructure/persistence"
	"EcommerceWithGolang/internal/interfaces/api"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(g *gin.RouterGroup, cfg database.AppConfig) {
	repositories, err := persistence.NewRepositories(cfg)

	if err != nil {
		panic(err)
	}

	// Application
	taxonomyApp := application.NewTaxonomyApp(repositories.Taxonomy)
	itemApp := application.NewItemApp(repositories.Item)

	// Handlers
	taxonomyHandler := api.NewTaxonomyHandler(taxonomyApp)
	itemHandler := api.NewItemHandler(itemApp)

	// Taxonomy routes
	taxonomyRouterGroup := g.Group("/taxonomy")
	taxonomyRouterGroup.GET("/list", taxonomyHandler.List)
	taxonomyRouterGroup.GET("/", taxonomyHandler.Create)
	taxonomyRouterGroup.GET("/:id", taxonomyHandler.GetById)

	// Items
	itemRouterGroup := g.Group("/item")
	itemRouterGroup.GET("/list", itemHandler.Create)

	// Promotion routes

}
