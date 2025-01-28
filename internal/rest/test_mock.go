package rest

import (
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
