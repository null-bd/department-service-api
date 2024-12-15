package errors

import "fmt"

// ErrorCode represents a unique error code for the application
type ErrorCode string

// ErrorDetail contains specific information about an error field
type ErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// AppError represents a custom application error
type AppError struct {
	Code    ErrorCode     `json:"code"`
	Message string        `json:"message"`
	Details []ErrorDetail `json:"details,omitempty"`
	Err     error         `json:"-"` // Internal error, not exposed in JSON
}

func (e *AppError) Error() string {
	return fmt.Sprintf("error code: %s, message: %s", e.Code, e.Message)
}

// New creates a new AppError
func New(code ErrorCode, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// WithDetails adds error details to an existing AppError
func (e *AppError) WithDetails(details ...ErrorDetail) *AppError {
	e.Details = details
	return e
}
