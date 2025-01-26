package department

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/null-bd/department-service-api/internal/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDepartmentService_CreateDepartment(t *testing.T) {
	tests := []struct {
		name        string
		input       *Department
		setupMocks  func(*mockRepository, *mockLogger)
		checkResult func(*testing.T, *Department, error)
	}{
		{
			name: "Success - Create New Department",
			input: &Department{
				Name: "Test Department",
				Code: "TEST001",
				Type: "medical",
			},
			setupMocks: func(repo *mockRepository, logger *mockLogger) {
				// Expect both begin and exit logs for successful case
				logger.On("Info", "service : CreateDepartment : begin", mock.Anything).Return()
				logger.On("Info", "service : CreateDepartment : exit", mock.Anything).Return()

				repo.On("GetByCode", mock.Anything, "TEST001").Return(nil, nil)
				repo.On("Create", mock.Anything, mock.MatchedBy(func(dept *Department) bool {
					return dept.Code == "TEST001" && dept.Status == "inactive"
				})).Return(&Department{
					ID:        uuid.New().String(),
					Name:      "Test Department",
					Code:      "TEST001",
					Type:      "medical",
					Status:    "inactive",
					CreatedAt: "2024-01-01T00:00:00Z",
					UpdatedAt: "2024-01-01T00:00:00Z",
				}, nil)
			},
			checkResult: func(t *testing.T, result *Department, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, "TEST001", result.Code)
				assert.Equal(t, "inactive", result.Status)
				assert.NotEmpty(t, result.ID)
			},
		},
		{
			name: "Error - Department Already Exists",
			input: &Department{
				Name: "Test Department",
				Code: "TEST001",
				Type: "medical",
			},
			setupMocks: func(repo *mockRepository, logger *mockLogger) {
				// Only expect begin log for error case
				logger.On("Info", "service : CreateDepartment : begin", mock.Anything).Return()

				existingDept := &Department{
					ID:   uuid.New().String(),
					Code: "TEST001",
				}
				repo.On("GetByCode", mock.Anything, "TEST001").Return(existingDept, nil)
			},
			checkResult: func(t *testing.T, result *Department, err error) {
				assert.Nil(t, result)
				assert.Error(t, err)
				appErr, ok := err.(*errors.AppError)
				assert.True(t, ok)
				assert.Equal(t, errors.ErrDeptExists, appErr.Code)
			},
		},
		{
			name: "Error - Repository Error on GetByCode",
			input: &Department{
				Name: "Test Department",
				Code: "TEST001",
				Type: "medical",
			},
			setupMocks: func(repo *mockRepository, logger *mockLogger) {
				// Only expect begin log for error case
				logger.On("Info", "service : CreateDepartment : begin", mock.Anything).Return()

				repo.On("GetByCode", mock.Anything, "TEST001").
					Return(nil, errors.New(errors.ErrDatabaseOperation, "database error", nil))
			},
			checkResult: func(t *testing.T, result *Department, err error) {
				assert.Nil(t, result)
				assert.Error(t, err)
				appErr, ok := err.(*errors.AppError)
				assert.True(t, ok)
				assert.Equal(t, errors.ErrDatabaseOperation, appErr.Code)
			},
		},
		{
			name: "Error - Repository Error on Create",
			input: &Department{
				Name: "Test Department",
				Code: "TEST001",
				Type: "medical",
			},
			setupMocks: func(repo *mockRepository, logger *mockLogger) {
				// Only expect begin log for error case
				logger.On("Info", "service : CreateDepartment : begin", mock.Anything).Return()

				repo.On("GetByCode", mock.Anything, "TEST001").Return(nil, nil)
				repo.On("Create", mock.Anything, mock.MatchedBy(func(dept *Department) bool {
					return dept.Code == "TEST001" && dept.Status == "inactive"
				})).Return(nil, errors.New(errors.ErrDatabaseOperation, "database error", nil))
			},
			checkResult: func(t *testing.T, result *Department, err error) {
				assert.Nil(t, result)
				assert.Error(t, err)
				appErr, ok := err.(*errors.AppError)
				assert.True(t, ok)
				assert.Equal(t, errors.ErrDatabaseOperation, appErr.Code)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mocks
			repo := new(mockRepository)
			logger := new(mockLogger)

			// Setup mocks
			tt.setupMocks(repo, logger)

			// Create service instance
			service := NewDepartmentService(repo, logger)

			// Execute test
			result, err := service.CreateDepartment(context.Background(), tt.input)

			// Check results
			tt.checkResult(t, result, err)

			// Verify mock expectations
			repo.AssertExpectations(t)
			logger.AssertExpectations(t)
		})
	}
}

func TestDepartmentService_GetDepartment(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		setupMocks  func(*mockRepository, *mockLogger)
		checkResult func(*testing.T, *Department, error)
	}{
		{
			name: "Success - Get Department",
			id:   "test-id-1",
			setupMocks: func(repo *mockRepository, logger *mockLogger) {
				logger.On("Info", "service : GetDepartment : begin", mock.Anything).Return()
				logger.On("Info", "service : GetDepartment : exit", mock.Anything).Return()

				repo.On("GetByID", mock.Anything, "test-id-1").Return(&Department{
					ID:     "test-id-1",
					Name:   "Test Department",
					Code:   "TEST001",
					Status: "active",
				}, nil)
			},
			checkResult: func(t *testing.T, result *Department, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, "test-id-1", result.ID)
			},
		},
		{
			name: "Error - Department Not Found",
			id:   "non-existent-id",
			setupMocks: func(repo *mockRepository, logger *mockLogger) {
				logger.On("Info", "service : GetDepartment : begin", mock.Anything).Return()

				repo.On("GetByID", mock.Anything, "non-existent-id").
					Return(nil, errors.New(errors.ErrDeptNotFound, "Department not found", nil))
			},
			checkResult: func(t *testing.T, result *Department, err error) {
				assert.Nil(t, result)
				assert.Error(t, err)
				appErr, ok := err.(*errors.AppError)
				assert.True(t, ok)
				assert.Equal(t, errors.ErrDeptNotFound, appErr.Code)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(mockRepository)
			logger := new(mockLogger)
			tt.setupMocks(repo, logger)

			service := NewDepartmentService(repo, logger)
			result, err := service.GetDepartment(context.Background(), tt.id)

			tt.checkResult(t, result, err)
			repo.AssertExpectations(t)
			logger.AssertExpectations(t)
		})
	}
}
func TestDepartmentService_ListDepartments(t *testing.T) {
	tests := []struct {
		name        string
		branchId    string
		setupMocks  func(*mockRepository, *mockLogger)
		checkResult func(*testing.T, []*Department, *Pagination, error)
	}{
		{
			name:     "Success - List Department",
			branchId: "test-branch-id-1",
			setupMocks: func(repo *mockRepository, logger *mockLogger) {
				logger.On("Info", "service : ListDepartments : begin", mock.Anything).Return()
				logger.On("Info", "service : ListDepartments : exit", mock.Anything).Return()

				repo.On("List", mock.Anything, "test-branch-id-1", mock.Anything, 1, 10).Return([]*Department{
					{
						ID:     "test-id-1",
						Name:   "Test Department1",
						Code:   "TEST001",
						Status: "active",
						Type:   "medical",
					},
					{
						ID:     "test-id-2",
						Name:   "Test Department2",
						Code:   "TEST002",
						Status: "active",
						Type:   "medical",
					},
				}, 2, nil)
			},
			checkResult: func(t *testing.T, result []*Department, pagination *Pagination, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.NotNil(t, pagination)
				assert.Equal(t, 2, pagination.Total)
			},
		},
		{
			name:     "Error - Department Not Found",
			branchId: "non-existent-branch-id",
			setupMocks: func(repo *mockRepository, logger *mockLogger) {
				logger.On("Info", "service : ListDepartments : begin", mock.Anything).Return()

				repo.On("List", mock.Anything, "non-existent-branch-id", mock.Anything, 1, 10).
					Return(nil, 0, errors.New(errors.ErrDeptNotFound, "Department not found", nil))
			},
			checkResult: func(t *testing.T, result []*Department, pagination *Pagination, err error) {
				assert.Nil(t, result)
				assert.Error(t, err)
				assert.Equal(t, "Department not found", err.Error())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(mockRepository)
			logger := new(mockLogger)
			tt.setupMocks(repo, logger)

			filter := map[string]interface{}{
				"status": "active",
				"type":   "medical",
			}
			service := NewDepartmentService(repo, logger)
			result, pagination, err := service.ListDepartments(context.Background(), tt.branchId, filter, 1, 10)

			tt.checkResult(t, result, pagination, err)

			repo.AssertExpectations(t)
			logger.AssertExpectations(t)
		})
	}
}
