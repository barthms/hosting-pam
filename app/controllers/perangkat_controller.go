package controllers

import (
	"net/http"

	"monitoring-service/app/helpers"
	"monitoring-service/app/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (m *Main) CreatePerangkat(c echo.Context) error {
	var req models.CreatePerangkatRequest

	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
	}

	if err := m.usecases.CreatePerangkat(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"Format ID Perangkat tidak valid, harus berupa angka"})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Data perangkat berhasil ditambahkan!",
	})
}

func (m *Main) DeletePerangkat(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return helpers.Response(c, http.StatusBadRequest, []string{"ID Perangkat harus dicantumkan"})
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"Format ID Perangkat tidak valid, harus berupa angka"})
	}

	idUint := uint(idInt)
	if err := m.usecases.DeletePerangkat(idUint); err != nil {
		statusCode := http.StatusInternalServerError
		return helpers.Response(c, statusCode, []string{err.Error()})
	}
	return helpers.StandardResponse(c, http.StatusOK, []string{"berhasil menghapus data perangkat"}, nil, nil)
}

func (m *Main) UpdatePerangkat(c echo.Context) error {
	claims, ok := c.Get("auth_claims").(*models.AuthClaims)
	if !ok {
		return helpers.Response(c, http.StatusUnauthorized, []string{"token tidak valid"})
	}

	penggunaID := uint(claims.UserID)

	var req models.UpdatePerangkatRequest
	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"Format request tidak valid"})
	}

	if err := m.usecases.UpdatePerangkat(penggunaID, &req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{"Berhasil memperbarui data perangkat"}, nil, nil)
}


func (m *Main) GetAllPerangkat(c echo.Context) error {
	data, err := m.usecases.GetAllPerangkat()
	if err != nil {
		return helpers.Response(c, http.StatusInternalServerError, []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{"Berhasil mendapatkan data perangkat"}, data, nil)
}

func (m *Main) SaveToken(c echo.Context) error {
	var req models.TokenRequest
	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"Format request tidak valid"})
	}

	if err := m.usecases.SaveFCMToken(&req); err != nil {
		return helpers.Response(c, http.StatusInternalServerError, []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{"Token berhasil diproses"}, nil, nil)
}