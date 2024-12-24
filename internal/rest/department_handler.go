package rest

import (
	"net/http"
	"strconv"

	"github.com/null-bd/department-service-api/internal/department"

	"github.com/gin-gonic/gin"
)

type DepartmentHandler struct {
	service department.DepartmentService
}

func NewDepartmentHandler(service department.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{service: service}
}

func (h *DepartmentHandler) ListDepartments(c *gin.Context) {
	filters := map[string]string{
		"branchId":  c.Query("branchId"),
		"type":      c.Query("type"),
		"status":    c.Query("status"),
		"specialty": c.Query("specialty"),
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	response, err := h.service.ListDepartments(filters, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch departments"})
		return
	}
	c.JSON(http.StatusOK, response)
}
