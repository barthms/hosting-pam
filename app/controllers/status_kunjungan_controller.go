package controllers

import (
	"net/http"
	// "strconv"

	"monitoring-service/app/constants"
	"monitoring-service/app/helpers"
	// "monitoring-service/app/models"

	"github.com/labstack/echo/v4"
)

func (m *Main) GetJumlahKunjunganByStatus(
	c echo.Context,
) error {

	data, err :=
		m.usecases.
			GetJumlahKunjunganByStatus()

	if err != nil {

		return helpers.Response(
			c,
			http.StatusInternalServerError,
			[]string{
				err.Error(),
			},
		)
	}

	return helpers.StandardResponse(
		c,
		http.StatusOK,
		[]string{
			constants.SUCCESS_RESPONSE_MESSAGE,
		},
		data,
		nil,
	)
}