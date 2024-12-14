package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/null-bd/microservice-name/internal/health"
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
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "unhealthy",
			"details": status,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"details": status,
	})
}
