package controllers

import (
	"net/http"
	"strconv"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type KesehatanLingkunganController struct {
	usecase usecases.KesehatanLingkunganUsecase
}

func NewKesehatanLingkunganController(u usecases.KesehatanLingkunganUsecase) *KesehatanLingkunganController {
	return &KesehatanLingkunganController{usecase: u}
}

// Master Data: Kategori & Indikator
func (c *KesehatanLingkunganController) CreateKategori(ctx echo.Context) error {
	var req models.KategoriLingkunganRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "request tidak valid"})
	}

	if err := c.usecase.CreateKategori(req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, models.Response{StatusCode: http.StatusCreated, Message: "kategori & indikator berhasil dibuat"})
}

func (c *KesehatanLingkunganController) AddIndikator(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var req models.IndikatorLingkunganRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "request tidak valid"})
	}

	if err := c.usecase.AddIndikator(int32(id), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, models.Response{StatusCode: http.StatusCreated, Message: "indikator berhasil ditambahkan"})
}

func (c *KesehatanLingkunganController) DeleteIndikator(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.usecase.DeleteIndikator(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Message: "indikator berhasil dihapus"})
}

func (c *KesehatanLingkunganController) GetAllKategori(ctx echo.Context) error {
	data, err := c.usecase.GetAllKategori()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: data})
}

func (c *KesehatanLingkunganController) DeleteKategori(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.usecase.DeleteKategori(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Message: "kategori berhasil dihapus"})
}

// Transaksi: Lembar & Detail
func (c *KesehatanLingkunganController) SubmitLembar(ctx echo.Context) error {
	var req models.LembarLingkunganRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "request tidak valid"})
	}

	if err := c.usecase.SubmitLembar(req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, models.Response{StatusCode: http.StatusCreated, Message: "data berhasil disimpan"})
}

func (c *KesehatanLingkunganController) GetHistory(ctx echo.Context) error {
	ibuIDStr := ctx.QueryParam("ibu_id")
	var ibuID int
	if ibuIDStr != "" {
		ibuID, _ = strconv.Atoi(ibuIDStr)
	}

	data, err := c.usecase.GetHistory(int32(ibuID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: data})
}

func (c *KesehatanLingkunganController) GetDetail(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := c.usecase.GetDetail(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: data})
}

func (c *KesehatanLingkunganController) DeleteLembar(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.usecase.DeleteLembar(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Message: "data berhasil dihapus"})
}
