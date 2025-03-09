package error

import "errors"

var (
	ErrInternalServerError = errors.New("Internal server error")
	ErrSqlError            = errors.New("Database server failed to execute query")
	ErrToManyRequests      = errors.New("Too many requests")
	ErrUnauthorized        = errors.New("Unautorizhed")
	ErrInvalidToken        = errors.New("Invalid token")
	ErrForBidden           = errors.New("Forbidden")
)

var GeneralErrors = []error{
	ErrInternalServerError,
	ErrSqlError,
	ErrToManyRequests,
	ErrUnauthorized,
	ErrInvalidToken,
	ErrForBidden,
}
