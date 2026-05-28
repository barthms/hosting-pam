package controllers

import (
	"errors"
	"net/http"
	"strings"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type WarnaTinjaController struct {
	usecase usecases.WarnaTinjaUsecase
}

func NewWarnaTinjaController(uc usecases.WarnaTinjaUsecase) *WarnaTinjaController {
	return &WarnaTinjaController{usecase: uc}
}

func isWarnaTinjaForbidden(err error) bool {
	return errors.Is(err, usecases.ErrWarnaTinjaForbidden)
}

func (c *WarnaTinjaController) GetByAnakIDForIbu(ctx echo.Context) error {
	claims, ok := getAuthClaims(ctx)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
	}

	anakID := strings.TrimSpace(ctx.QueryParam("anak_id"))
	if anakID == "" {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "anak_id wajib diisi"})
	}

	data, err := c.usecase.GetByAnakIDForIbu(anakID, uint(claims.UserID))
	if err != nil {
		if isWarnaTinjaForbidden(err) {
			return ctx.JSON(http.StatusForbidden, models.Response{StatusCode: http.StatusForbidden, Message: "Anda tidak memiliki akses ke data anak ini"})
		}
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: data})
}

func (c *WarnaTinjaController) SaveForIbu(ctx echo.Context) error {
	claims, ok := getAuthClaims(ctx)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
	}

	var req models.WarnaTinjaSaveRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "format request tidak valid"})
	}

	data, err := c.usecase.SaveForIbu(req, uint(claims.UserID))
	if err != nil {
		if isWarnaTinjaForbidden(err) {
			return ctx.JSON(http.StatusForbidden, models.Response{StatusCode: http.StatusForbidden, Message: "Anda tidak memiliki akses ke data anak ini"})
		}
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Data warna tinja berhasil disimpan",
		Data:       data,
	})
}
