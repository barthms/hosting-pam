package controllers

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PengukuranLilAController struct {
	usecase usecases.PengukuranLilAUseCase
}

func NewPengukuranLilAController(u usecases.PengukuranLilAUseCase) *PengukuranLilAController {
	return &PengukuranLilAController{usecase: u}
}

func (c *PengukuranLilAController) Create(ctx echo.Context) error {
	var req models.CreatePengukuranLilARequest

	if err := ctx.Bind(&req); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid request", err)
	}

	if err := c.usecase.Create(req); err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "failed to create pengukuran lila", err)
	}

	return success(ctx, http.StatusCreated, "success create Pengukuran Lila", nil)
}

func (c *PengukuranLilAController) Update(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid id", err)
	}

	var req models.UpdatePengukuranLilARequest
	if err := ctx.Bind(&req); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid request", err)
	}

	if err := c.usecase.Update(int32(id), req); err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "failed to update pengukuran lila", err)
	}

	return success(ctx, http.StatusOK, "success update pengukuran lila", nil)
}
func (c *PengukuranLilAController) GetByAnakID(ctx echo.Context) error {
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

func (c *PengukuranLilAController) GetByAnakIDForIbu(ctx echo.Context) error {
	claims, ok := getAuthClaims(ctx)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{"statusCode": http.StatusUnauthorized, "message": "Unauthorized"})
	}

	param := ctx.Param("anak_id")
	if param == "" {
		param = ctx.QueryParam("anak_id")
	}
	if param == "" {
		return errorResponse(ctx, http.StatusBadRequest, "invalid anak_id", errors.New("anak_id wajib diisi"))
	}

	anakID, err := strconv.Atoi(param)
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid anak_id", err)
	}

	data, err := c.usecase.GetByAnakIDForIbu(int32(anakID), uint(claims.UserID))
	if err != nil {
		if err.Error() == "akses ditolak: anak tidak ditemukan atau bukan milik anda" {
			return errorResponse(ctx, http.StatusForbidden, "Forbidden", err)
		}
		return errorResponse(ctx, http.StatusInternalServerError, "failed to fetch data", err)
	}

	return success(ctx, http.StatusOK, "success get data", data)
}
func (c *PengukuranLilAController) GetByID(ctx echo.Context) error {
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
func (c *PengukuranLilAController) GetAll(ctx echo.Context) error {
	data, err := c.usecase.GetAll()
	if err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "failed to fetch data", err)
	}

	return success(ctx, http.StatusOK, "success get all data", data)
}

func (c *PengukuranLilAController) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid id", err)
	}

	if err := c.usecase.Delete(int32(id)); err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "failed to delete pengukuran lila", err)
	}

	return success(ctx, http.StatusOK, "success delete pengukuran lila", nil)
}