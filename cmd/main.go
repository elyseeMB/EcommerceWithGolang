package main

import (
	"EcommerceWithGolang/internal/infrastructure/database"
	"EcommerceWithGolang/internal/interfaces"
	"EcommerceWithGolang/pkg/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load configuration
	cfg := loadConfig()

	router := setupRouter(*cfg)
	server := setupServer("8080", router)

	if err := startServer(server); err != nil {
		panic(err)
	}

}

func setupRouter(cfg database.AppConfig) *gin.Engine {
	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"response": "ok",
		})
	})

	apiRouter := router.Group("/api/v1")
	interfaces.RegisterRoutes(apiRouter, cfg)

	return router
}

func setupServer(port string, router *gin.Engine) *http.Server {
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	return server

}

func startServer(server *http.Server) error {
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err

	}
	return nil
}

func loadConfig() *database.AppConfig {
	return &database.AppConfig{
		DriverName:     "sqlite",
		DataSourceName: constants.SQLITE_DB_NAME + ".db",
	}
}
