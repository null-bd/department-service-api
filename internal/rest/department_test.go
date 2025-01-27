package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/null-bd/department-service-api/internal/department"
	"github.com/null-bd/department-service-api/internal/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockDeptSvc struct {
	mock.Mock
}

func (m *mockDeptSvc) GetDepartment(ctx context.Context, id string) (*department.Department, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*department.Department), args.Error(1)
}

func (m *mockDeptSvc) ListDepartments(ctx context.Context, branchId string, filter map[string]interface{}, page, limit int) ([]*department.Department, *department.Pagination, error) {
	args := m.Called(ctx, branchId, filter, page, limit)
	return args.Get(0).([]*department.Department), args.Get(1).(*department.Pagination), args.Error(2)
}

func (m *mockDeptSvc) CreateDepartment(ctx context.Context, dept *department.Department) (*department.Department, error) {
	args := m.Called(ctx, dept)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*department.Department), args.Error(1)
}

func setupTest(t *testing.T) (*gin.Engine, *mockDeptSvc, *mockLogger) {
	t.Log("Setting up test")
	gin.SetMode(gin.TestMode)

	mockDeptSvc := new(mockDeptSvc)
	mockLog := new(mockLogger)

	handler := NewDepartmentHandler(mockDeptSvc, mockLog)

	router := gin.New()
	router.POST("/departments", handler.CreateDepartment)
	router.GET("/departments", handler.ListDepartments)
	router.GET("/departments/:deptId", handler.GetDepartment)

	return router, mockDeptSvc, mockLog
}
func TestGetDepartment(t *testing.T) {

	router, mockDeptSvc, mockLog := setupTest(t)

	tests := []struct {
		name           string
		id             string
		setupMocks     func()
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "Success",
			id:   "test-id-1",
			setupMocks: func() {
				mockLog.On("Info", "handler : GetDepartment : begin", mock.Anything).Return()
				mockLog.On("Info", "handler : GetDepartment : exit", mock.Anything).Return()

				mockDeptSvc.On("GetDepartment", mock.Anything, "test-id-1").Return(
					&department.Department{
						ID:     "test-id-1",
						Name:   "Test Department",
						Code:   "TEST001",
						Status: "active",
					}, nil)
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var response DepartmentResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, "test-id-1", response.ID)
				assert.Equal(t, "Test Department", response.Name)
			},
		},
		{
			name: "Not Found",
			id:   "non-existent-id",
			setupMocks: func() {
				mockLog.On("Info", "handler : GetDepartment : begin", mock.Anything).Return()
				mockLog.On("Info", "handler : GetDepartment : exit", mock.Anything).Return()

				mockDeptSvc.On("GetDepartment", mock.Anything, "non-existent-id").Return(
					nil, errors.New(errors.ErrDeptNotFound, "department not found", nil))
			},
			expectedStatus: http.StatusNotFound,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {

				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, string(errors.ErrDeptNotFound), response["code"])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			req, _ := http.NewRequest("GET", "/departments/"+tt.id, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			tt.checkResponse(t, w)

			mockDeptSvc.AssertExpectations(t)
			mockLog.AssertExpectations(t)
		})
	}
}

func TestCreateDepartment(t *testing.T) {
	router, mockDeptSvc, mockLog := setupTest(t)

	mockLog.On("Info", "handler : CreateDepartment : begin", mock.Anything).Return()
	mockLog.On("Info", "handler : CreateDepartment : exit", mock.Anything).Return()

	inputDTO := CreateDepartmentRequest{
		Name: "Test Department",
		Code: "TEST001",
		Type: "medical",
	}

	expectedDept := &department.Department{
		ID:     "test-id-1",
		Name:   "Test Department",
		Code:   "TEST001",
		Type:   "medical",
		Status: "inactive",
	}

	mockDeptSvc.On("CreateDepartment", mock.Anything, mock.MatchedBy(func(dept *department.Department) bool {
		return dept.Name == inputDTO.Name && dept.Code == inputDTO.Code
	})).Return(expectedDept, nil)

	body, _ := json.Marshal(inputDTO)
	req, _ := http.NewRequest("POST", "/departments", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response DepartmentResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedDept.ID, response.ID)
	assert.Equal(t, inputDTO.Name, response.Name)

	mockDeptSvc.AssertExpectations(t)
	mockLog.AssertExpectations(t)
}

func TestListDepartments(t *testing.T) {
	router, mockDeptSvc, mockLog := setupTest(t)

	mockLog.On("Info", "handler : ListDepartments : begin", mock.Anything).Return()
	mockLog.On("Info", "handler : ListDepartments : exit", mock.Anything).Return()

	depts := []*department.Department{
		{
			ID:     "test-id-1",
			Name:   "Test Department 1",
			Code:   "TEST001",
			Status: "active",
		},
		{
			ID:     "test-id-2",
			Name:   "Test Department 2",
			Code:   "TEST002",
			Status: "inactive",
		},
	}

	pagination := &department.Pagination{
		Total: 2,
		Page:  1,
		Pages: 1,
	}

	expectedFilter := map[string]interface{}{"status": "active"}
	mockDeptSvc.On("ListDepartments", mock.Anything, expectedFilter, 1, 20).Return(depts, pagination, nil)

	req, _ := http.NewRequest("GET", "/Departments?status=active", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	data := response["data"].([]interface{})
	assert.Equal(t, 2, len(data))

	mockDeptSvc.AssertExpectations(t)
	mockLog.AssertExpectations(t)
}
