package internalerrors

import "errors"

var (
	ErrInternalError error = errors.New("internal server error")
)
