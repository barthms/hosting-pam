package controllers

import (
	"net/http"
	"strconv"
	"time"

	"monitoring-service/app/constants"
	"monitoring-service/app/helpers"
	"monitoring-service/app/models"

	"github.com/labstack/echo/v4"
)

func (m *Main) GetJadwalImunisasi(
	c echo.Context,
) error {

	claimsValue :=
		c.Get("auth_claims")

	claims, ok :=
		claimsValue.(*models.AuthClaims)

	if !ok {

		return helpers.Response(
			c,
			http.StatusUnauthorized,
			[]string{
				"user tidak valid",
			},
		)
	}

	data, err :=
		m.usecases.
			GetJadwalImunisasi(
				int32(claims.UserID),
			)

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

func (m *Main) GetJadwalImunisasiByAnakID(c echo.Context) error {

	// =========================
	// 1. Ambil auth claims
	// =========================
	claimsValue := c.Get("auth_claims")

	claims, ok := claimsValue.(*models.AuthClaims)
	if !ok {
		return helpers.Response(
			c,
			http.StatusUnauthorized,
			[]string{"user tidak valid"},
		)
	}

	// =========================
	// 2. Ambil anakId dari param
	// =========================
	anakIDParam := c.Param("anak_id")

	anakID, err := strconv.Atoi(anakIDParam)
	if err != nil || anakID <= 0 {
		return helpers.Response(
			c,
			http.StatusBadRequest,
			[]string{"anak_id tidak valid"},
		)
	}

	// =========================
	// 3. Call usecase
	// =========================
	data, err := m.usecases.GetJadwalImunisasiByAnakID(
		int32(claims.UserID),
		int32(anakID),
	)

	if err != nil {
		return helpers.Response(
			c,
			http.StatusInternalServerError,
			[]string{err.Error()},
		)
	}

	// =========================
	// 4. Return response sukses
	// =========================
	return helpers.StandardResponse(
		c,
		http.StatusOK,
		[]string{constants.SUCCESS_RESPONSE_MESSAGE},
		data,
		nil,
	)
}	

func (m *Main) UpdateTanggalEstimasi(c echo.Context) error {

	idParam := c.Param("id")

	jadwalID, err := strconv.Atoi(idParam)
	if err != nil || jadwalID <= 0 {
		return helpers.Response(c, 400, []string{"id tidak valid"})
	}

	var req models.UpdateTanggalEstimasiRequest

	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, 400, []string{"invalid request"})
	}

	parsedDate, err := time.Parse("2006-01-02", req.TanggalEstimasi)
	if err != nil {
		return helpers.Response(c, 400, []string{"format tanggal harus YYYY-MM-DD"})
	}

	// 🔥 AMBIL USER ID
	claims := c.Get("auth_claims").(*models.AuthClaims)
	userID := claims.UserID

	err = m.usecases.UpdateTanggalEstimasi(
		userID,
		uint(jadwalID),
		parsedDate,
	)

	if err != nil {
		return helpers.Response(c, 500, []string{err.Error()})
	}

	return helpers.StandardResponse(
		c,
		200,
		[]string{"success"},
		nil,
		nil,
	)
}

func (m *Main) GetJadwalByID(c echo.Context) error {

	claimsValue := c.Get("auth_claims")

	claims, ok := claimsValue.(*models.AuthClaims)
	if !ok || claims == nil {
		return c.JSON(401, map[string]string{
			"message": "unauthorized",
		})
	}

	jadwalID, err := strconv.Atoi(c.Param("id"))
	if err != nil || jadwalID <= 0 {
		return c.JSON(400, map[string]string{
			"message": "id tidak valid",
		})
	}

	data, err := m.usecases.GetJadwalImunisasiByJadwalID(
		int32(claims.UserID),
		uint(jadwalID),
	)

	if err != nil {
		return c.JSON(500, map[string]string{
			"message": err.Error(),
		})
	}

	if data == nil {
		return c.JSON(404, map[string]string{
			"message": "data tidak ditemukan",
		})
	}

	return c.JSON(200, data)
}