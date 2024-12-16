package rest

import (
	"github.com/null-bd/department-service-api/internal/health"
	"github.com/null-bd/logger"
)

func NewHandler(healthSvc health.IHealthService, logger logger.Logger) IHealthHandler {
	return &healthHandler{
		healthSvc: healthSvc,
		log:       logger,
	}
}
