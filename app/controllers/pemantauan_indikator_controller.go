package controllers

import (
	"net/http"
	"strconv"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type PemantauanIndikatorController struct {
	usecase usecases.PemantauanIndikatorUsecase
}

func NewPemantauanIndikatorController(u usecases.PemantauanIndikatorUsecase) *PemantauanIndikatorController {
	return &PemantauanIndikatorController{usecase: u}
}

type pemantauanIndikatorRequest struct {
	KategoriUsia string `json:"kategori_usia"`
	Deskripsi    string `json:"deskripsi"`
}

func (c *PemantauanIndikatorController) GetAll(ctx echo.Context) error {
	list, err := c.usecase.GetAll(ctx.QueryParam("kategori_usia"), ctx.QueryParam("q"))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: list})
}

func (c *PemantauanIndikatorController) Create(ctx echo.Context) error {
	var req pemantauanIndikatorRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}

	payload := &models.PemantauanIndikator{
		KategoriUsia: req.KategoriUsia,
		Deskripsi:    req.Deskripsi,
	}

	data, err := c.usecase.Create(payload)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	return ctx.JSON(http.StatusCreated, models.Response{StatusCode: http.StatusCreated, Data: data})
}

func (c *PemantauanIndikatorController) Update(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}

	var req pemantauanIndikatorRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := c.usecase.Update(int32(id), req.KategoriUsia, req.Deskripsi)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: data})
}

func (c *PemantauanIndikatorController) Delete(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}

	if err := c.usecase.Delete(int32(id)); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Message: "deleted"})
}
