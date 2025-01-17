package department

import (
	"context"
	stderr "errors"

	"github.com/google/uuid"
	"github.com/null-bd/department-service-api/internal/errors"
	"github.com/null-bd/logger"
)

type IDepartmentService interface {
	CreateDepartment(ctx context.Context, dept *Department) (*Department, error)
	GetDepartment(ctx context.Context, id string) (*Department, error)
	ListDepartments(ctx context.Context, branchId string, filter map[string]interface{}, page, limit int) ([]*Department, *Pagination, error)
}

type departmentService struct {
	repo IDepartmentRepository
	log  logger.Logger
}

func NewDepartmentService(repo IDepartmentRepository, logger logger.Logger) IDepartmentService {
	return &departmentService{
		repo: repo,
		log:  logger,
	}
}

func (s *departmentService) CreateDepartment(ctx context.Context, dept *Department) (*Department, error) {
	s.log.Info("service : CreateDepartment : begin", nil)

	// Check if organization exists
	existingdept, err := s.repo.GetByID(ctx, dept.ID)
	if err != nil {
		return nil, err
	}
	if existingdept != nil {
		return nil, &errors.AppError{
			Code:    errors.ErrDeptExists,
			Message: "department with this code already exists",
			Err:     stderr.New("department with this code already exists"),
		}
	}

	// Set required fields
	dept.ID = uuid.New().String()
	dept.Status = "inactive"

	// Create organization
	createdDept, err := s.repo.Create(ctx, dept)
	if err != nil {
		return nil, err
	}

	s.log.Info("service : CreateDepartment : exit", nil)
	return createdDept, nil
}

func (s *departmentService) GetDepartment(ctx context.Context, id string) (*Department, error) {
	s.log.Info("service : GetDepatment : begin", logger.Fields{"id": id})

	dept, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	s.log.Info("service : GetDepatment : exit", nil)
	return dept, nil
}

func (s *departmentService) ListDepartments(ctx context.Context, branchId string, filter map[string]interface{}, page, limit int) ([]*Department, *Pagination, error) {
	s.log.Info("service : ListDepartments : begin", logger.Fields{"branchId": branchId})

	departments, total, err := s.repo.List(ctx, branchId, filter, page, limit)
	if err != nil {
		return nil, nil, err
	}

	pages := (total + limit - 1) / limit
	pagination := &Pagination{
		Total: total,
		Page:  page,
		Pages: pages,
	}

	s.log.Info("service : ListDepartments : exit", nil)
	return departments, pagination, nil
}
