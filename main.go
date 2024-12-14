package main

import (
	"log"
	"os"

	"github.com/null-bd/department-service-api/config"
	"github.com/null-bd/department-service-api/config/database"

	// "github.com/null-bd/department-service-api/config/router"
	"github.com/gin-gonic/gin"
	"github.com/null-bd/department-service-api/internal/app"
)

func main() {
	// Get environment
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	// Load configuration
	cfg, err := config.LoadConfig(env)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database connection
	db, err := database.NewPostgresConnection(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize application
	application := app.NewApplication(db.Pool, cfg)

	handler := application.Handler

	// Initialize router with auth middleware
	// router, err := router.NewRouter(cfg, application.Handler)
	// if err != nil {
	// 	log.Fatalf("Failed to initialize router: %v", err)
	// }

	// router, err := gin.r
	// Start server
	router := gin.New()
	router.GET("/health", handler.HealthCheck)

	log.Printf("Starting server on port %d in %s mode", cfg.App.Port, cfg.App.Env)
	if err := router.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
