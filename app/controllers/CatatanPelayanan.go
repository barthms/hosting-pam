package controllers

import (
	"monitoring-service/app/models"
	"monitoring-service/app/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CatatanPelayananController struct {
	usecase usecases.CatatanPelayananUseCase
}

func NewCatatanPelayananController(u usecases.CatatanPelayananUseCase) *CatatanPelayananController {
	return &CatatanPelayananController{usecase: u}
}

func (c *CatatanPelayananController) Create(ctx echo.Context) error {
	var req models.CreateCatatanPelayananRequest

	if err := ctx.Bind(&req); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid request", err)
	}

	if err := c.usecase.Create(req); err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "failed to create catatan pelayanan", err)
	}

	return success(ctx, http.StatusCreated, "success create Catatan Pelayanan", nil)
}

func (c *CatatanPelayananController) Update(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid id", err)
	}

	var req models.UpdateCatatanPelayananRequest
	if err := ctx.Bind(&req); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid request", err)
	}

	if err := c.usecase.Update(int32(id), req); err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "failed to update catatan pelayanan", err)
	}

	return success(ctx, http.StatusOK, "success update catatan pelayanan", nil)
}
func (c *CatatanPelayananController) GetByAnakID(ctx echo.Context) error {
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
func (c *CatatanPelayananController) GetByID(ctx echo.Context) error {
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
func (c *CatatanPelayananController) GetAll(ctx echo.Context) error {
	data, err := c.usecase.GetAll()
	if err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "failed to fetch data", err)
	}

	return success(ctx, http.StatusOK, "success get all data", data)
}

func (c *CatatanPelayananController) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid id", err)
	}

	if err := c.usecase.Delete(int32(id)); err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "failed to delete catatan pelayanan", err)
	}

	return success(ctx, http.StatusOK, "success delete catatan pelayanan", nil)
}