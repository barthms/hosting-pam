package customerror

import (
	"fmt"

	"github.com/pkg/errors"
)

type badRequest struct {
	TrackableError
}

type BadRequestError interface {
	error
	IsBadRequestError() bool
}

func (e *badRequest) IsBadRequestError() bool { return true }

func NewBadRequestErrorf(format string, data ...interface{}) (err error) {
	err = errors.New(fmt.Sprintf(format, data...))
	return &badRequest{TrackableError{err}}
}

func NewBadRequestError(message string) (err error) {
	err = errors.New(message)
	return &badRequest{TrackableError{err}}
}
