package controllers

import (
	"errors"
	"net/http"
	"strings"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type LembarPemantauanController struct {
	usecase usecases.LembarPemantauanUsecase
}

func NewLembarPemantauanController(uc usecases.LembarPemantauanUsecase) *LembarPemantauanController {
	return &LembarPemantauanController{usecase: uc}
}

// func getAuthClaims(ctx echo.Context) (*models.AuthClaims, bool) {
// 	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
// 	if !ok || claims == nil {
// 		return nil, false
// 	}
// 	return claims, true
// }

func isLembarPemantauanForbidden(err error) bool {
	return errors.Is(err, usecases.ErrLembarPemantauanForbidden)
}

func (c *LembarPemantauanController) Create(ctx echo.Context) error {
	var req models.LembarPemantauanRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Format request tidak valid",
		})
	}

	lembar, err := c.usecase.Create(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, models.Response{
		StatusCode: http.StatusCreated,
		Message:    "Lembar pemantauan berhasil dibuat",
		Data:       lembar,
	})
}

func (c *LembarPemantauanController) CreateForIbu(ctx echo.Context) error {
	claims, ok := getAuthClaims(ctx)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
	}

	var req models.LembarPemantauanRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "Format request tidak valid"})
	}

	// Cast ke uint sesuai spesifikasi
	lembar, err := c.usecase.CreateForIbu(req, uint(claims.UserID))
	if err != nil {
		if isLembarPemantauanForbidden(err) {
			return ctx.JSON(http.StatusForbidden, models.Response{StatusCode: http.StatusForbidden, Message: "Anda tidak memiliki akses ke data anak ini"})
		}
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, models.Response{StatusCode: http.StatusCreated, Message: "Lembar pemantauan berhasil dibuat", Data: lembar})
}

func (c *LembarPemantauanController) GetRentangUsiaForIbu(ctx echo.Context) error {
	if _, ok := getAuthClaims(ctx); !ok {
		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
	}

	rentang, err := c.usecase.GetRentangUsia()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: rentang})
}

func (c *LembarPemantauanController) GetKategoriByRentangUsiaForIbu(ctx echo.Context) error {
	if _, ok := getAuthClaims(ctx); !ok {
		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
	}

	rentangUsiaID := strings.TrimSpace(ctx.QueryParam("rentang_usia_id"))
	if rentangUsiaID == "" {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "rentang_usia_id wajib diisi"})
	}

	kategori, err := c.usecase.GetKategoriTandaSakitByRentangUsiaID(rentangUsiaID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: kategori})
}

func (c *LembarPemantauanController) GetByID(ctx echo.Context) error {
	id := ctx.Param("id")

	lembar, err := c.usecase.GetByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Data:       lembar,
	})
}

func (c *LembarPemantauanController) GetByAnakID(ctx echo.Context) error {
	anakID := ctx.QueryParam("anak_id")

	if anakID == "" {
		lembars, err := c.usecase.GetAll()
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, models.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			})
		}
		return ctx.JSON(http.StatusOK, models.Response{
			StatusCode: http.StatusOK,
			Data:       lembars,
		})
	}

	lembars, err := c.usecase.GetByAnakID(anakID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Data:       lembars,
	})
}

func (c *LembarPemantauanController) GetByAnakIDForIbu(ctx echo.Context) error {
	claims, ok := getAuthClaims(ctx)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
	}

	anakID := strings.TrimSpace(ctx.QueryParam("anak_id"))
	if anakID == "" {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "anak_id wajib diisi"})
	}

	lembars, err := c.usecase.GetByAnakIDForIbu(anakID, uint(claims.UserID))
	if err != nil {
		if isLembarPemantauanForbidden(err) {
			return ctx.JSON(http.StatusForbidden, models.Response{StatusCode: http.StatusForbidden, Message: "Anda tidak memiliki akses ke data anak ini"})
		}
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: lembars})
}

func (c *LembarPemantauanController) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	var req models.LembarPemantauanRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Format request tidak valid",
		})
	}

	lembar, err := c.usecase.Update(id, req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Lembar pemantauan berhasil diupdate",
		Data:       lembar,
	})
}

// Verify endpoint untuk Tenaga Kesehatan
func (c *LembarPemantauanController) Verify(ctx echo.Context) error {
	id := ctx.Param("id")

	var req models.LembarPemantauanVerifikasiRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Format request tidak valid",
		})
	}

	lembar, err := c.usecase.Verify(id, req)
	if err != nil {
		statusCode := http.StatusBadRequest
		if strings.Contains(err.Error(), "tidak ditemukan") {
			statusCode = http.StatusNotFound
		}
		return ctx.JSON(statusCode, models.Response{
			StatusCode: statusCode,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Lembar pemantauan berhasil diverifikasi",
		Data:       lembar,
	})
}

func (c *LembarPemantauanController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")

	err := c.usecase.Delete(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Lembar pemantauan berhasil dihapus",
	})
}
