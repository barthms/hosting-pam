package helpers

import (
	"monitoring-service/app/models"

	"github.com/labstack/echo/v4"
)

func ResponseWrapper(c echo.Context, statusCode int, response interface{}) error {
	return c.JSON(statusCode, response)
}

func StandardResponse(c echo.Context, statusCode int, message interface{}, data interface{}, pagination *models.Pagination) error {
	switch {
	case pagination == nil:
		return ResponseWrapper(c, statusCode, models.Response{
			StatusCode: statusCode,
			Message:    message,
			Data:       data,
		})
	default:
		return ResponseWrapper(c, statusCode, models.ResponseWithPaginate{
			StatusCode: statusCode,
			Message:    message,
			Data:       data,
			Pagination: pagination,
		})
	}
}

func Response(c echo.Context, statusCode int, message []string) error {
	return ResponseWrapper(c, statusCode, models.BasicResponse{
		StatusCode: statusCode,
		Message:    message,
	})
}
