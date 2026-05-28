package customerror

import (
	"fmt"

	"github.com/pkg/errors"
)

type internalService struct {
	TrackableError
}

type InternalServiceError interface {
	error
	IsInternalError() bool
}

func (e *internalService) IsInternalServiceError() bool { return true }

func NewInternalServiceErrorf(format string, data ...interface{}) (err error) {
	err = errors.New(fmt.Sprintf(format, data...))
	return &internalService{TrackableError{err}}
}

func NewInternalServiceError(message string) (err error) {
	err = errors.New(message)
	return &internalService{TrackableError{err}}
}
