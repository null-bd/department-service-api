package app

import (
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/null-bd/microservice-name/config"
	"github.com/null-bd/microservice-name/internal/health"
	"github.com/null-bd/microservice-name/internal/rest"
)

type Application struct {
	Handler *rest.Handler
	DB      *pgxpool.Pool
	Config  *config.Config
}

func NewApplication(db *pgxpool.Pool, cfg *config.Config) *Application {
	// Initialize repositories with pgx pool
	healthRepo := health.NewHealthRepository(db)

	// Initialize services
	healthSvc := health.NewHealthService(healthRepo)

	// Initialize handler
	h := rest.NewHandler(healthSvc)

	return &Application{
		Handler: h,
		DB:      db,
		Config:  cfg,
	}
}
