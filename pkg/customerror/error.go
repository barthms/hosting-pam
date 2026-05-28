package customerror

import (
	"net/http"

	"github.com/pkg/errors"
)

type TrackableError struct {
	errors error
}

func (e TrackableError) Error() string { return e.errors.Error() }
func (e TrackableError) Cause() error  { return e.errors }

func New(msg string) error {
	return errors.New(msg)
}

func GetStatusCode(err error) int {
	if _, ok := err.(NotFoundError); ok {
		return http.StatusNotFound
	} else if _, ok := err.(BadRequestError); ok {
		return http.StatusBadRequest
	} else if _, ok := err.(ConflictError); ok {
		return http.StatusConflict
	} else if _, ok := err.(ForbiddenError); ok {
		return http.StatusForbidden
	}
	return http.StatusInternalServerError

}
