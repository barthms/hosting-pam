package controllers

import (
	"errors"
	"fmt"
	"monitoring-service/app/models"
	"monitoring-service/app/usecases"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PelayananKesehatanAnakController struct {
	usecase usecases.PelayananKesehatanAnakUseCase
}

func NewPelayananKesehatanAnakController(uc usecases.PelayananKesehatanAnakUseCase) *PelayananKesehatanAnakController {
	return &PelayananKesehatanAnakController{usecase: uc}
}

// Create
func (c *PelayananKesehatanAnakController) CreatePelayananHandler(ctx echo.Context) error {
	var req models.CreatePelayananKesehatanAnakRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request format",
		})
	}

	//  DEBUG (hapus nanti kalau sudah aman)
	fmt.Println("DEBUG REQUEST:", req)

	//  VALIDASI WAJIB (INI KUNCI)
	if req.AnakID == 0 ||
		req.PeriodeID == 0 ||
		req.KategoriUmurID == 0 ||
		req.Tanggal == "" ||
		req.Lokasi == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "field wajib tidak boleh kosong atau 0",
		})
	}

	if len(req.DetailPelayanan) == 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "detail_pelayanan wajib diisi",
		})
	}

	if err := c.usecase.Create(req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Pelayanan kesehatan anak berhasil dibuat",
	})
}

// parse int 32
func ParseInt(param string, fieldName string) (int32, error) {
	val, err := strconv.Atoi(param)
	if err != nil {
		return 0, fmt.Errorf("%s harus berupa angka", fieldName)
	}
	return int32(val), nil
}

// GetByAnakID
func (c *PelayananKesehatanAnakController) GetByAnakID(ctx echo.Context) error {
	anakIDStr := ctx.QueryParam("anak_id")

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

// // get by id kunjungan
func (c *PelayananKesehatanAnakController) GetByID(ctx echo.Context) error {
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
				"error": "Data tidak ditemukan",
			})
		}

		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, data)
}

// Update
func (c *PelayananKesehatanAnakController) Update(ctx echo.Context) error {
	id, err := ParseInt(ctx.Param("id"), "id")

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	var req models.UpdatePelayananKesehatanAnakRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Request tidak valid",
		})
	}

	if err := c.usecase.Update(id, req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Berhasil update data",
	})
}

func (c *PelayananKesehatanAnakController) Delete(ctx echo.Context) error {
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
		"message": "Berhasil hapus data",
	})
}
