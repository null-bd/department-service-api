package department

import (
	"context"

	"github.com/null-bd/logger"
)

type IDepartmentService interface {
	//GetDepartment(ctx context.Context, id string) (*Department, error)
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
