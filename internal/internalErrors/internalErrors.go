package internalerrors

import "errors"

var (
	ErrInternalError    error = errors.New("internal server error")
	ErrResourceNotFound error = errors.New("Resource not found")
)
