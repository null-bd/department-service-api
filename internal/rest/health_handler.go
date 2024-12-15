package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/null-bd/department-service-api/internal/health"
)

type Handler struct {
	healthSvc *health.HealthService
}

func NewHandler(healthSvc *health.HealthService) *Handler {
	return &Handler{
		healthSvc: healthSvc,
	}
}

func (h *Handler) HealthCheck(c *gin.Context) {
	status, err := h.healthSvc.CheckHealth()
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"details": status,
	})
}
