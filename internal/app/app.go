package app

import (
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/null-bd/department-service-api/config"
	"github.com/null-bd/department-service-api/internal/health"
	"github.com/null-bd/department-service-api/internal/rest"
	"github.com/null-bd/logger"
)

type Application struct {
	Handler *rest.Handler
	DB      *pgxpool.Pool
	Config  *config.Config
}

func NewApplication(logger logger.Logger, cfg *config.Config, db *pgxpool.Pool) *Application {
	// Initialize repositories
	healthRepo := health.NewHealthRepository(db, logger)

	// Initialize services
	healthSvc := health.NewHealthService(healthRepo, logger)

	// Initialize handler
	h := rest.NewHandler(healthSvc, logger)

	return &Application{
		Handler: h,
		DB:      db,
		Config:  cfg,
	}
}
