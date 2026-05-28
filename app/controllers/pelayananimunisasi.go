package controllers

import (
	"monitoring-service/app/models"
	"monitoring-service/app/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type KunjunganImunisasiController struct {
	usecase usecases.KunjunganImunisasiUseCase
}

func NewKunjunganImunisasiController(u usecases.KunjunganImunisasiUseCase) *KunjunganImunisasiController {
	return &KunjunganImunisasiController{usecase: u}
}

func (c *KunjunganImunisasiController) Create(ctx echo.Context) error {
	var req models.CreateKunjunganImunisasiRequest

	if err := ctx.Bind(&req); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid request", err)
	}

	if err := c.usecase.Create(req); err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "failed to create kunjungan imunisasi", err)
	}

	return success(ctx, http.StatusCreated, "success create kunjungan imunisasi", nil)
}

func (c *KunjunganImunisasiController) Update(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid id", err)
	}

	var req models.UpdateKunjunganImunisasiRequest
	if err := ctx.Bind(&req); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid request", err)
	}

	if err := c.usecase.Update(int32(id), req); err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "failed to update kunjungan imunisasi", err)
	}

	return success(ctx, http.StatusOK, "success update kunjungan imunisasi", nil)
}

func (c *KunjunganImunisasiController) GetByAnakID(ctx echo.Context) error {
	param := ctx.QueryParam("anak_id")

	// kalau tidak ada anak_id → ambil semua
	if param == "" {
		data, err := c.usecase.GetAll()
		if err != nil {
			return errorResponse(ctx, http.StatusInternalServerError, "failed to fetch data", err)
		}
		return success(ctx, http.StatusOK, "success get all data", data)
	}

	anakID, err := strconv.Atoi(param)
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid anak_id", err)
	}

	data, err := c.usecase.GetByAnakID(int32(anakID))
	if err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "failed to fetch data", err)
	}

	return success(ctx, http.StatusOK, "success get data", data)
}

func (c *KunjunganImunisasiController) GetByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid id", err)
	}

	data, err := c.usecase.GetByID(int32(id))
	if err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "failed to fetch data", err)
	}

	if data == nil {
		return errorResponse(ctx, http.StatusNotFound, "data not found", nil)
	}

	return success(ctx, http.StatusOK, "success get data", data)
}

func (c *KunjunganImunisasiController) GetAll(ctx echo.Context) error {
	data, err := c.usecase.GetAll()
	if err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "failed to fetch data", err)
	}

	return success(ctx, http.StatusOK, "success get all data", data)
}

func (c *KunjunganImunisasiController) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid id", err)
	}

	if err := c.usecase.Delete(int32(id)); err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "failed to delete kunjungan imunisasi", err)
	}

	return success(ctx, http.StatusOK, "success delete kunjungan imunisasi", nil)
}