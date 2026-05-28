package controllers

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PemeriksaanGigiController struct {
	usecase usecases.PemeriksaanGigiUseCase
}

func NewPemeriksaanGigiController(u usecases.PemeriksaanGigiUseCase) *PemeriksaanGigiController {
	return &PemeriksaanGigiController{usecase: u}
}

func (c *PemeriksaanGigiController) Create(ctx echo.Context) error {
	var req models.CreatePemeriksaanGigiRequest

	if err := ctx.Bind(&req); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid request", err)
	}

	if err := c.usecase.Create(req); err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "failed to create pemeriksaaan Gigi", err)
	}

	return success(ctx, http.StatusCreated, "success create Pemeriksaan Gigi", nil)
}

func (c *PemeriksaanGigiController) Update(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid id", err)
	}

	var req models.UpdatePemeriksaanGigiRequest
	if err := ctx.Bind(&req); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid request", err)
	}

	if err := c.usecase.Update(int32(id), req); err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "failed to update pemeriksaan gigi", err)
	}

	return success(ctx, http.StatusOK, "success update pemeriksaan gigi", nil)
}
func (c *PemeriksaanGigiController) GetByAnakID(ctx echo.Context) error {
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

func (c *PemeriksaanGigiController) GetByAnakIDForIbu(ctx echo.Context) error {
	claims, ok := getAuthClaims(ctx)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{"statusCode": http.StatusUnauthorized, "message": "Unauthorized"})
	}

	param := ctx.QueryParam("anak_id")
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

func (c *PemeriksaanGigiController) GetByID(ctx echo.Context) error {
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
func (c *PemeriksaanGigiController) GetAll(ctx echo.Context) error {
	data, err := c.usecase.GetAll()
	if err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "failed to fetch data", err)
	}

	return success(ctx, http.StatusOK, "success get all data", data)
}

func (c *PemeriksaanGigiController) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid id", err)
	}

	if err := c.usecase.Delete(int32(id)); err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "failed to delete pemeriksaan gigi", err)
	}

	return success(ctx, http.StatusOK, "success delete pemeriksaan gigi", nil)
}