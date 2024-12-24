package department

import (
	"context"

	"github.com/null-bd/logger"
)

type IDepartmentService interface {
	// CreateDepartment(ctx context.Context, dept *Department) (*Department, error)
	//GetDepartment(ctx context.Context, id string) (*Department, error)
	ListDepartments(ctx context.Context, filter map[string]interface{}, page, limit int) ([]*Department, *Pagination, error)
	// UpdateDepartment(ctx context.Context, dept *Department) (*Department, error)
	// DeleteDepartment(ctx context.Context, id string) error
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

// func (s *departmentService) CreateDepartment(ctx context.Context, dept *Department) (*Department, error) {
// 	s.log.Info("service : CreateDepartment : begin", nil)

// 	// Check if department exists
// 	existingDept, err := s.repo.GetByCode(ctx, dept.Code)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if existingDept != nil {
// 		return nil, &errors.AppError{
// 			Code:    errors.ErrDeptExists,
// 			Message: "department with this code already exists",
// 			Err:     stderr.New("department with this code already exists"),
// 		}
// 	}

// 	// Set required fields
// 	dept.ID = uuid.New().String()
// 	dept.Status = "inactive"

// 	// Create department
// 	createdDept, err := s.repo.Create(ctx, dept)
// 	if err != nil {
// 		return nil, err
// 	}

// 	s.log.Info("service : CreateDepartment : exit", nil)
// 	return createdDept, nil
// }

// func (s *departmentService) GetDepartment(ctx context.Context, id string) (*Department, error) {
// 	s.log.Info("service : GetDepartment : begin", logger.Fields{"id": id})

// 	dept, err := s.repo.GetByID(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	s.log.Info("service : GetDepartment : exit", nil)
// 	return dept, nil
// }

func (s *departmentService) ListDepartments(ctx context.Context, filter map[string]interface{}, page, limit int) ([]*Department, *Pagination, error) {
	s.log.Info("service : ListDepartments : begin", nil)

	departments, total, err := s.repo.List(ctx, filter, page, limit)
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

// func (s *departmentService) UpdateDepartment(ctx context.Context, dept *Department) (*Department, error) {
// 	s.log.Info("service : UpdateDepartment : begin", nil)

// 	// Verify department exists
// 	existing, err := s.repo.GetByID(ctx, dept.ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Prevent modification of immutable fields
// 	dept.Code = existing.Code
// 	dept.CreatedAt = existing.CreatedAt

// 	// Update department
// 	err = s.repo.Update(ctx, dept)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Get updated department
// 	updated, err := s.repo.GetByID(ctx, dept.ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	s.log.Info("service : UpdateDepartment : exit", nil)
// 	return updated, nil
// }

// func (s *departmentService) DeleteDepartment(ctx context.Context, id string) error {
// 	s.log.Info("service : DeleteDepartment : begin", nil)

// 	// Verify department exists and check if it can be deleted
// 	dept, err := s.repo.GetByID(ctx, id)
// 	if err != nil {
// 		return err
// 	}

// 	if dept.Status != "inactive" {
// 		return &errors.AppError{
// 			Code:    errors.ErrDeptActive,
// 			Message: "cannot delete active department",
// 			Err:     stderr.New("department must be inactive before deletion"),
// 		}
// 	}

// 	err = s.repo.Delete(ctx, id)
// 	if err != nil {
// 		return err
// 	}

// 	s.log.Info("service : DeleteDepartment : exit", nil)
// 	return nil
// }
