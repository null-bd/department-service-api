package errors

// Health service specific error constructors
func NewDatabaseConnectionError(err error) *AppError {
	return New(
		ErrDatabaseConnection,
		"Failed to connect to database",
		err,
	)
}

func NewDatabaseQueryError(err error) *AppError {
	return New(
		ErrDatabaseQuery,
		"Database query failed",
		err,
	)
}
