package controllers

import (
	"net/http"

	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type KategoriUmurController struct {
	usecase usecases.KategoriUmurUsecase
}

func NewKategoriUmurController(u usecases.KategoriUmurUsecase) *KategoriUmurController {
	return &KategoriUmurController{u}
}

func (c *KategoriUmurController) GetAll(ctx echo.Context) error {
	data, err := c.usecase.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"data": data})
}