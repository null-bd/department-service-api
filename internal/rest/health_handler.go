package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/null-bd/department-service-api/internal/health"
	"github.com/null-bd/logger"
)

type Handler struct {
	healthSvc *health.HealthService
	log       logger.Logger
}

func NewHandler(healthSvc *health.HealthService, logger logger.Logger) *Handler {
	return &Handler{
		log:       logger,
		healthSvc: healthSvc,
	}
}

func (h *Handler) HealthCheck(c *gin.Context) {
	h.log.Info("handler : HealthCheck : begin", nil)
	status, err := h.healthSvc.CheckHealth()
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"details": status,
	})
	h.log.Info("handler : HealthCheck : exit", nil)
}
