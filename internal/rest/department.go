package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/null-bd/department-service-api/internal/department"
	"github.com/null-bd/logger"
)

type IDepartmentHandler interface {
	// CreateDepartment(c *gin.Context)
	// GetDepartment(c *gin.Context)
	ListDepartments(c *gin.Context)
	// UpdateDepartment(c *gin.Context)
	// DeleteDepartment(c *gin.Context)
}

type departmentHandler struct {
	deptSvc department.IDepartmentService
	log     logger.Logger
}

func NewDepartmentHandler(deptSvc department.IDepartmentService, logger logger.Logger) IDepartmentHandler {
	return &departmentHandler{
		deptSvc: deptSvc,
		log:     logger,
	}
}

// func (h *departmentHandler) CreateDepartment(c *gin.Context) {
// 	h.log.Info("handler : CreateDepartment : begin", nil)

// 	var req CreateDepartmentRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		HandleError(c, errors.New(errors.ErrBadRequest, "invalid request body", err))
// 		return
// 	}

// 	dept := ToDepartment(&req)
// 	result, err := h.deptSvc.CreateDepartment(c.Request.Context(), dept)
// 	if err != nil {
// 		HandleError(c, err)
// 		return
// 	}

// 	c.JSON(http.StatusCreated, ToDepartmentResponse(result))
// 	h.log.Info("handler : CreateDepartment : exit", nil)
// }

// func (h *departmentHandler) GetDepartment(c *gin.Context) {
// 	h.log.Info("handler : GetDepartment : begin", nil)

// 	id := c.Param("id")
// 	if id == "" {
// 		HandleError(c, errors.New(errors.ErrBadRequest, "missing department id", nil))
// 		return
// 	}

// 	dept, err := h.deptSvc.GetDepartment(c.Request.Context(), id)
// 	if err != nil {
// 		HandleError(c, err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, ToDepartmentResponse(dept))
// 	h.log.Info("handler : GetDepartment : exit", nil)
// }

func (h *departmentHandler) ListDepartments(c *gin.Context) {
	h.log.Info("handler : ListDepartments : begin", nil)

	filter := make(map[string]interface{})
	if status := c.Query("status"); status != "" {
		filter["status"] = status
	}
	if deptType := c.Query("type"); deptType != "" {
		filter["type"] = deptType
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	departments, pagination, err := h.deptSvc.ListDepartments(c.Request.Context(), filter, page, limit)
	if err != nil {
		HandleError(c, err)
		return
	}

	// Convert domain objects to response DTOs
	responses := make([]*DepartmentResponse, len(departments))
	for i, dept := range departments {
		responses[i] = ToDepartmentResponse(dept)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       responses,
		"pagination": pagination,
	})
	h.log.Info("handler : ListDepartments : exit", nil)
}

// func (h *departmentHandler) UpdateDepartment(c *gin.Context) {
// 	h.log.Info("handler : UpdateDepartment : begin", nil)

// 	id := c.Param("id")
// 	if id == "" {
// 		HandleError(c, errors.New(errors.ErrBadRequest, "missing department id", nil))
// 		return
// 	}

// 	var req UpdateDepartmentRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		HandleError(c, errors.New(errors.ErrBadRequest, "invalid request body", err))
// 		return
// 	}

// 	// Create domain object from request
// 	dept := &department.Department{
// 		ID:          id,
// 		Name:        req.Name,
// 		Type:        req.Type,
// 		Description: req.Description,
// 		Status:      req.Status,
// 		ContactInfo: department.ContactInfo{
// 			Email:   req.ContactInfo.Email,
// 			Phone:   req.ContactInfo.Phone,
// 			Address: req.ContactInfo.Address,
// 		},
// 		Metadata: req.Metadata,
// 	}

// 	result, err := h.deptSvc.UpdateDepartment(c.Request.Context(), dept)
// 	if err != nil {
// 		HandleError(c, err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, ToDepartmentResponse(result))
// 	h.log.Info("handler : UpdateDepartment : exit", nil)
// }

// func (h *departmentHandler) DeleteDepartment(c *gin.Context) {
// 	h.log.Info("handler : DeleteDepartment : begin", nil)

// 	id := c.Param("id")
// 	if id == "" {
// 		HandleError(c, errors.New(errors.ErrBadRequest, "missing department id", nil))
// 		return
// 	}

// 	err := h.deptSvc.DeleteDepartment(c.Request.Context(), id)
// 	if err != nil {
// 		HandleError(c, err)
// 		return
// 	}

// 	c.Status(http.StatusNoContent)
// 	h.log.Info("handler : DeleteDepartment : exit", nil)
// }
