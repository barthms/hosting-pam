package controllers

import (
	"monitoring-service/app/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type JenisPelayananController struct {
	usecase usecases.JenisPelayananUsecase
}

func NewJenisPelayananController(u usecases.JenisPelayananUsecase) *JenisPelayananController {
	return &JenisPelayananController{usecase: u}
}

func (c *JenisPelayananController) GetJenisPelayanan(ctx echo.Context) error {
	kategoriParam := ctx.QueryParam("kategori_umur_id")
	periodeParam := ctx.QueryParam("periode_id")

	if kategoriParam == "" || periodeParam == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "kategori_umur_id dan periode_id wajib diisi",
		})
	}

	kategoriID, err := strconv.Atoi(kategoriParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "kategori_umur_id harus angka",
		})
	}

	periodeID, err := strconv.Atoi(periodeParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "periode_id harus angka",
		})
	}

	data, err := c.usecase.GetJenisPelayanan(int32(kategoriID), int32(periodeID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    data,
	})
}