// app/controllers/laporan_ibu_controller.go
package controllers

import (
	"net/http"

	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type LaporanIbuController struct {
	usecase usecases.LaporanIbuUsecase
}

func NewLaporanIbuController(
	usecase usecases.LaporanIbuUsecase,
) *LaporanIbuController {
	return &LaporanIbuController{usecase}
}

func (c *LaporanIbuController) ExportExcel(
	ctx echo.Context,
) error {

	file, err := c.usecase.ExportExcelLaporanIbu()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return ctx.Attachment(file, file)
}
func (c *LaporanIbuController) Preview(
	ctx echo.Context,
) error {

	data, err := c.usecase.GetLaporanIbu()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data": data,
	})
}
