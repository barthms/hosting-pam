package controllers

import (
	"net/http"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type AbsensiKelasIbuBalitaController struct {
	usecase usecases.AbsensiKelasIbuBalitaUsecase
}

func NewAbsensiKelasIbuBalitaController(
	u usecases.AbsensiKelasIbuBalitaUsecase,
) *AbsensiKelasIbuBalitaController {
	return &AbsensiKelasIbuBalitaController{usecase: u}
}

type saveAbsensiKelasIbuBalitaRequest struct {
	Tanggal      string `json:"tanggal"`
	NamaKader    string `json:"nama_kader"`
	TanggalParaf string `json:"tanggal_paraf"`
}

func (c *AbsensiKelasIbuBalitaController) GetMine(ctx echo.Context) error {
	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
	if !ok || claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "token tidak valid",
		})
	}

	data, err := c.usecase.GetMine(claims.UserID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Data:       data,
	})
}

func (c *AbsensiKelasIbuBalitaController) SaveMine(ctx echo.Context) error {
	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
	if !ok || claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "token tidak valid",
		})
	}

	var req saveAbsensiKelasIbuBalitaRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "format request tidak valid",
		})
	}

	tanggal, err := parseOptionalDate(req.Tanggal)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "format tanggal harus YYYY-MM-DD",
		})
	}

	tanggalParaf, err := parseOptionalDate(req.TanggalParaf)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "format tanggal_paraf harus YYYY-MM-DD",
		})
	}

	data, err := c.usecase.SaveMine(
		claims.UserID,
		models.AbsensiKelasIbuBalita{
			Tanggal:      tanggal,
			NamaKader:    req.NamaKader,
			TanggalParaf: tanggalParaf,
		},
	)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Data:       data,
	})
}


