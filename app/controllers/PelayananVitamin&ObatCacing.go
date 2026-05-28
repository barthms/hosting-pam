package controllers

import (
	"monitoring-service/app/models"
	"monitoring-service/app/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type KunjunganVitaminController struct {
	usecase usecases.KunjunganVitaminUseCase
}

func NewKunjunganVitaminController(u usecases.KunjunganVitaminUseCase) *KunjunganVitaminController {
	return &KunjunganVitaminController{usecase: u}
}
func success(ctx echo.Context, code int, message string, data interface{}) error {
	return ctx.JSON(code, map[string]interface{}{
		"message": message,
		"data":    data,
	})
}

func errorResponse(ctx echo.Context, code int, message string, err error) error {
	res := map[string]interface{}{
		"message": message,
	}

	if err != nil {
		res["error"] = err.Error()
	}

	return ctx.JSON(code, res)
}
func (c *KunjunganVitaminController) Create(ctx echo.Context) error {
	var req models.CreateKunjunganVitaminRequest

	if err := ctx.Bind(&req); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid request", err)
	}

	if err := c.usecase.Create(req); err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "failed to create kunjungan vitamin", err)
	}

	return success(ctx, http.StatusCreated, "success create kunjungan vitamin", nil)
}
func (c *KunjunganVitaminController) Update(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid id", err)
	}

	var req models.UpdateKunjunganVitaminRequest
	if err := ctx.Bind(&req); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid request", err)
	}

	if err := c.usecase.Update(int32(id), req); err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "failed to update", err)
	}

	return success(ctx, http.StatusOK, "success update", nil)
}
func (c *KunjunganVitaminController) GetByAnakID(ctx echo.Context) error {
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
func (c *KunjunganVitaminController) GetByID(ctx echo.Context) error {
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

	return success(ctx, http.StatusOK, "success", data)
}
func (c *KunjunganVitaminController) GetAll(ctx echo.Context) error {
	data, err := c.usecase.GetAll()
	if err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "failed to fetch data", err)
	}

	return success(ctx, http.StatusOK, "success", data)
}
func (c *KunjunganVitaminController) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid id", err)
	}

	if err := c.usecase.Delete(int32(id)); err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "failed to delete", err)
	}

	return success(ctx, http.StatusOK, "success delete", nil)
}