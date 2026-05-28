package controllers

import (
	"monitoring-service/app/models"
	"monitoring-service/app/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PemantauanIbuController struct {
	usecase usecases.PemantauanIbuUseCase
}

func NewPemantauanIbuController(u usecases.PemantauanIbuUseCase) *PemantauanIbuController {
	return &PemantauanIbuController{u}
}

func (c *PemantauanIbuController) Save(ctx echo.Context) error {
	var req models.LembarPemantauanIbuRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request body"})
	}

	if err := c.usecase.SavePemantauan(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Data pemantauan berhasil disimpan"})
}

func (c *PemantauanIbuController) GetByIbu(ctx echo.Context) error {
	ibuIDStr := ctx.QueryParam("ibu_id")
	ibuID, _ := strconv.Atoi(ibuIDStr)

	data, err := c.usecase.GetAllByIbu(int32(ibuID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"data": data})
}

func (c *PemantauanIbuController) GetKategori(ctx echo.Context) error {
	data, err := c.usecase.GetKategori()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"data": data})
}

func (c *PemantauanIbuController) Delete(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := c.usecase.DeletePemantauan(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Data berhasil dihapus"})
}
