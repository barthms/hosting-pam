package controllers

import (
	"net/http"
	"time"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type AbsensiKelasIbuHamilController struct {
	usecase usecases.AbsensiKelasIbuHamilUsecase
}

func NewAbsensiKelasIbuHamilController(
	u usecases.AbsensiKelasIbuHamilUsecase,
) *AbsensiKelasIbuHamilController {
	return &AbsensiKelasIbuHamilController{usecase: u}
}

type saveAbsensiKelasIbuHamilRequest struct {
	PertemuanKe  int32  `json:"pertemuan_ke"`
	Tanggal      string `json:"tanggal"`
	NamaKader    string `json:"nama_kader"`
	TanggalParaf string `json:"tanggal_paraf"`
}

func parseOptionalDate(value string) (*time.Time, error) {
	if value == "" {
		return nil, nil
	}

	parsed, err := time.Parse("2006-01-02", value)
	if err != nil {
		return nil, err
	}

	return &parsed, nil
}

func (c *AbsensiKelasIbuHamilController) GetMine(ctx echo.Context) error {
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

func (c *AbsensiKelasIbuHamilController) SaveMine(ctx echo.Context) error {
	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
	if !ok || claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "token tidak valid",
		})
	}

	var req saveAbsensiKelasIbuHamilRequest
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
		models.AbsensiKelasIbuHamil{
			PertemuanKe:  req.PertemuanKe,
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
