package controllers

import (
	"monitoring-service/app/models"
	"monitoring-service/app/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RiwayatKehamilanLaluController struct {
	usecase usecases.RiwayatKehamilanLaluUsecase
}

func NewRiwayatKehamilanLaluController(u usecases.RiwayatKehamilanLaluUsecase) *RiwayatKehamilanLaluController {
	return &RiwayatKehamilanLaluController{usecase: u}
}

func (c *RiwayatKehamilanLaluController) Create(ctx echo.Context) error {
	var req models.RiwayatKehamilanLalu
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: 400, Message: err.Error()})
	}
	if err := c.usecase.Create(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: 500, Message: err.Error()})
	}
	return ctx.JSON(http.StatusCreated, models.Response{StatusCode: 201, Data: req})
}

func (c *RiwayatKehamilanLaluController) GetByID(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: 400, Message: "invalid id"})
	}
	data, err := c.usecase.GetByID(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: 404, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: 200, Data: data})
}

func (c *RiwayatKehamilanLaluController) GetByEvaluasiID(ctx echo.Context) error {
	evaluasiID, err := strconv.ParseInt(ctx.QueryParam("evaluasi_id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: 400, Message: "evaluasi_id required"})
	}
	list, err := c.usecase.GetByEvaluasiID(int32(evaluasiID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: 500, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: 200, Data: list})
}

func (c *RiwayatKehamilanLaluController) Update(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: 400, Message: "invalid id"})
	}
	var req models.RiwayatKehamilanLalu
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: 400, Message: err.Error()})
	}
	req.IDRiwayat = int32(id)
	if err := c.usecase.Update(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: 500, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: 200, Data: req})
}

func (c *RiwayatKehamilanLaluController) Delete(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: 400, Message: "invalid id"})
	}
	if err := c.usecase.Delete(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: 500, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: 200, Message: "deleted"})
}
