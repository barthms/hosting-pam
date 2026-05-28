package customerror

import (
	"fmt"

	"github.com/pkg/errors"
)

type conflict struct {
	TrackableError
}

type ConflictError interface {
	error
	IsConflictError() bool
}

func (e *conflict) IsConflictError() bool { return true }

func NewConflictErrorf(format string, data ...interface{}) (err error) {
	err = errors.New(fmt.Sprintf(format, data...))
	return &conflict{TrackableError{err}}
}

func NewConflictError(message string) (err error) {
	err = errors.New(message)
	return &conflict{TrackableError{err}}
}
