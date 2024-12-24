package app

import (
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/null-bd/department-service-api/config"
	"github.com/null-bd/department-service-api/internal/department"
	"github.com/null-bd/department-service-api/internal/health"
	"github.com/null-bd/department-service-api/internal/rest"
	"github.com/null-bd/logger"
)

type Application struct {
	HealthHandler rest.IHealthHandler
	DeptHandler   rest.IDepartmentHandler
	DB            *pgxpool.Pool
	Config        *config.Config
}

func NewApplication(logger logger.Logger, cfg *config.Config, db *pgxpool.Pool) *Application {
	// Initialize repositories
	healthRepo := health.NewHealthRepository(db, logger)
	deptRepo := department.NewDepartmentRepository(db, logger)

	// Initialize services
	healthSvc := health.NewHealthService(healthRepo, logger)
	deptSvc := department.NewDepartmentService(deptRepo, logger)

	// Initialize handler
	h := rest.NewHealthHandler(healthSvc, logger)
	d := rest.NewDepartmentHandler(deptSvc, logger)

	return &Application{
		HealthHandler: h,
		DeptHandler:   d,
		DB:            db,
		Config:        cfg,
	}
}
