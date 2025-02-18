package department

import (
	"context"

	"github.com/null-bd/logger"
	"github.com/stretchr/testify/mock"
)

type mockLogger struct {
	mock.Mock
}

func (m *mockLogger) Debug(msg string, fields logger.Fields) {
	m.Called(msg, fields)
}

func (m *mockLogger) Info(msg string, fields logger.Fields) {
	m.Called(msg, fields)
}

func (m *mockLogger) Warn(msg string, fields logger.Fields) {
	m.Called(msg, fields)
}

func (m *mockLogger) Error(msg string, fields logger.Fields) {
	m.Called(msg, fields)
}

func (m *mockLogger) Fatal(msg string, fields logger.Fields) {
	m.Called(msg, fields)
}

func (m *mockLogger) WithFields(fields logger.Fields) logger.Logger {
	m.Called(fields)
	return nil
}

// Mock Repository
type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) Create(ctx context.Context, dept *Department) (*Department, error) {
	args := m.Called(ctx, dept)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Department), args.Error(1)
}

func (m *mockRepository) GetByCode(ctx context.Context, code string) (*Department, error) {
	args := m.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Department), args.Error(1)
}

func (m *mockRepository) GetByID(ctx context.Context, id string) (*Department, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Department), args.Error(1)
}

func (m *mockRepository) List(ctx context.Context, branchId string, filter map[string]interface{}, page, limit int) ([]*Department, int, error) {
	args := m.Called(ctx, branchId, filter, page, limit)
	if args.Get(0) == nil {
		return []*Department{}, args.Int(1), args.Error(2)
	}
	return args.Get(0).([]*Department), args.Int(1), args.Error(2)
}

func (m *mockRepository) Update(ctx context.Context, dept *Department) error {
	args := m.Called(ctx, dept)
	return args.Error(0)
}

func (m *mockRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
