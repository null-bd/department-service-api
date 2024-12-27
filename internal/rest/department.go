package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/null-bd/department-service-api/internal/department"
	"github.com/null-bd/department-service-api/internal/errors"
	"github.com/null-bd/logger"
)

type IDepartmentHandler interface {
	// GetDepartment(c *gin.Context)
	ListDepartments(c *gin.Context)
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

func (h *departmentHandler) ListDepartments(c *gin.Context) {
	h.log.Info("handler : ListDepartments : begin", nil)

	branchId := c.Query("branchId")
	if branchId == "" {
		HandleError(c, errors.New(errors.ErrBadRequest, "Missing BranchId", nil))
		return
	}

	filter := make(map[string]interface{})
	if status := c.Query("status"); status != "" {
		filter["status"] = status
	}
	if deptType := c.Query("type"); deptType != "" {
		filter["type"] = deptType
	}
	if deptSpeciality := c.Query("speciality"); deptSpeciality != "" {
		filter["speciality"] = deptSpeciality
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	departments, pagination, err := h.deptSvc.ListDepartments(c.Request.Context(), branchId, filter, page, limit)
	if err != nil {
		HandleError(c, err)
		return
	}

	// Convert domain objects to response DTOs
	responses := make([]*ListDepartmentResponse, len(departments))
	for i, dept := range departments {
		responses[i] = ToDepartmentResponse(dept)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       responses,
		"pagination": pagination,
	})
	h.log.Info("handler : ListDepartments : exit", nil)
}

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
