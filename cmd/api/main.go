package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/alutec-inventory-be/internal/adapters/datasources"
	"github.com/tapiaw38/alutec-inventory-be/internal/adapters/datasources/repositories"
	"github.com/tapiaw38/alutec-inventory-be/internal/adapters/web"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/appcontext"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/config"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/database"
	"github.com/tapiaw38/alutec-inventory-be/internal/usecases"
)

func main() {
	loadConfig()

	cfg := config.GetConfigService()
	gin.SetMode(cfg.ServerConfig.GinMode)

	db, err := database.GetSQLClient()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	ds := datasources.CreateDatasources(db)
	reposFactory := repositories.NewFactory(ds)
	repos := reposFactory()
	factory := appcontext.NewFactory(repos)
	uc := usecases.NewUsecases(factory)

	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.ServerConfig.FrontendURL, "http://localhost:9000", "http://localhost:9001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           2 * time.Hour,
	}))

	app.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": cfg.ServerConfig.AppName})
	})

	web.RegisterRoutes(app, uc)

	port := cfg.ServerConfig.Port
	log.Printf("%s running on port %s", cfg.ServerConfig.AppName, port)
	if err := app.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
