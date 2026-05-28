package controllers

// import (
// 	"monitoring-service/app/constants"
// 	"monitoring-service/app/helpers"
// 	"monitoring-service/app/models"
// 	"monitoring-service/pkg/customerror"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// // GET
// func (m *Main) GetMasterStandar(c echo.Context) error {
// 	parameter := c.QueryParam("parameter")
// 	jenisKelamin := c.QueryParam("jenis_kelamin")

// 	data, err := m.usecases.GetMasterStandar(parameter, jenisKelamin)
// 	if err != nil {
// 		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
// 	}

// 	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
// }

// // CREATE
// func (m *Main) CreateMasterStandar(c echo.Context) error {
// 	var req models.CreateMasterStandarRequest
// 	if err := c.Bind(&req); err != nil {
// 		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
// 	}

// 	if err := m.usecases.CreateMasterStandar(&req); err != nil {
// 		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
// 	}

// 	return helpers.StandardResponse(c, http.StatusCreated, []string{constants.SUCCESS_RESPONSE_MESSAGE}, map[string]string{
// 		"message": "master standar antropometri berhasil ditambahkan",
// 	}, nil)
// }
