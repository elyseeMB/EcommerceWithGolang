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

	// Handlers
	taxonomyHandler := api.NewTaxonomyHandler(taxonomyApp)

	taxonomyRouterGroup := g.Group("/taxonomy")

	taxonomyRouterGroup.GET("/", taxonomyHandler.Create)
}
