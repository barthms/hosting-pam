package controllers

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type kunjunganGiziController struct {
	usecase usecases.KunjunganGiziUseCase
}

func NewKunjunganGiziController(uc usecases.KunjunganGiziUseCase) *kunjunganGiziController {
	return &kunjunganGiziController{usecase: uc}
}

func (c *kunjunganGiziController) Create(ctx echo.Context) error {
	var req models.CreatePelayananGiziRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "format request tidak valid",
		})
	}

	// VALIDASI WAJIB
	if req.AnakID == 0 ||
		req.TenagaKesehatanID == 0 ||
		req.Tanggal == "" ||
		req.Lokasi == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "field wajib tidak boleh kosong",
		})
	}

	if err := c.usecase.Create(req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "kunjungan gizi berhasil dibuat",
	})
}
func (c *kunjunganGiziController) GetByAnakID(ctx echo.Context) error {
	anakIDStr := ctx.QueryParam("anak_id")

	// kalau kosong → ambil semua
	if anakIDStr == "" {
		data, err := c.usecase.GetAll()
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
		return ctx.JSON(http.StatusOK, data)
	}

	anakID, err := ParseInt(anakIDStr, "anak_id")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	data, err := c.usecase.GetByAnakID(anakID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, data)
}

func (c *kunjunganGiziController) GetByID(ctx echo.Context) error {
	id, err := ParseInt(ctx.Param("id"), "id")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	data, err := c.usecase.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.JSON(http.StatusNotFound, map[string]string{
				"error": "data tidak ditemukan",
			})
		}

		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, data)
}

func (c *kunjunganGiziController) Update(ctx echo.Context) error {
	id, err := ParseInt(ctx.Param("id"), "id")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	var req models.UpdatePelayananGiziRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "request tidak valid",
		})
	}

	if err := c.usecase.Update(id, req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "berhasil update kunjungan gizi",
	})
}

func (c *kunjunganGiziController) Delete(ctx echo.Context) error {
	id, err := ParseInt(ctx.Param("id"), "id")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	if err := c.usecase.Delete(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "berhasil hapus kunjungan gizi",
	})
}
