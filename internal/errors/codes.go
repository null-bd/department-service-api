package errors

const (
	// region request error codes
	ErrBadRequest ErrorCode = "DEPTAPI_101"

	// region service error codes
	ErrDeptExists ErrorCode = "DEPTAPI_201"
	ErrDeptActive ErrorCode = "DEPTAPI_202"

	ErrDeptInactive ErrorCode = "DEPTAPI_211"
	ErrBranchExists ErrorCode = "DEPTAPI_212"
	ErrBranchActive ErrorCode = "DEPTAPI_213"

	// region repository error codes
	ErrDatabaseConnection ErrorCode = "DEPTAPI_301"
	ErrDatabaseQuery      ErrorCode = "DEPTAPI_302"
	ErrCacheConnection    ErrorCode = "DEPTAPI_303"
	ErrDatabaseOperation  ErrorCode = "DEPTAPI_304"
	ErrDeptNotFound       ErrorCode = "DEPTAPI_305"

	ErrBranchNotFound ErrorCode = "DEPTAPI_311"
)
