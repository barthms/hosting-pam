package customerror

import "fmt"

// ForbiddenError represents a 403 Forbidden error
type ForbiddenError struct {
	message string
}

// Error returns the error message
func (e ForbiddenError) Error() string {
	return e.message
}

// NewForbiddenError creates a new ForbiddenError
func NewForbiddenError(message string) error {
	return ForbiddenError{message: message}
}

// NewForbiddenErrorf creates a new ForbiddenError with formatted message
func NewForbiddenErrorf(format string, args ...interface{}) error {
	return ForbiddenError{message: fmt.Sprintf(format, args...)}
}
