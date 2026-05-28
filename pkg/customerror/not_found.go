package customerror

import (
	"fmt"

	"github.com/pkg/errors"
)

type notFound struct {
	TrackableError
}

type NotFoundError interface {
	error
	IsNotFoundError() bool
}

func (e *notFound) IsNotFoundError() bool { return true }

func NewNotFoundErrorf(format string, data ...interface{}) (err error) {
	err = errors.New(fmt.Sprintf(format, data...))
	return &notFound{TrackableError{err}}
}

func NewNotFoundError(message string) (err error) {
	err = errors.New(message)
	return &notFound{TrackableError{err}}
}
