package controllers

import (
	"monitoring-service/app/models"
	"monitoring-service/app/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PemantauanAnakController struct {
	useCase usecases.PemantauanAnakUseCase
}

func NewPemantauanAnakController(useCase usecases.PemantauanAnakUseCase) *PemantauanAnakController {
	return &PemantauanAnakController{useCase}
}

func (c *PemantauanAnakController) Save(ctx echo.Context) error {
	var req models.LembarPemantauanRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request"})
	}

	if err := req.Validate(); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	if err := c.useCase.SavePemantauan(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "Data pemantauan berhasil disimpan"})
}

func (c *PemantauanAnakController) GetHistory(ctx echo.Context) error {
	anakID, _ := strconv.ParseUint(ctx.QueryParam("anak_id"), 10, 32)
	rentangID, _ := strconv.ParseUint(ctx.QueryParam("rentang_usia_id"), 10, 32)

	data, err := c.useCase.GetHistory(uint(anakID), uint(rentangID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"data": data})
}

func (c *PemantauanAnakController) GetRentangUsia(ctx echo.Context) error {
	data, err := c.useCase.GetRentangUsia()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"data": data})
}

func (c *PemantauanAnakController) GetKategori(ctx echo.Context) error {
	rentangID, _ := strconv.ParseUint(ctx.Param("rentang_id"), 10, 32)

	data, err := c.useCase.GetKategoriByRentang(uint(rentangID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"data": data})
}

func (c *PemantauanAnakController) Delete(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err := c.useCase.DeletePemantauan(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "Data berhasil dihapus"})
}

func (c *PemantauanAnakController) CreateKategori(ctx echo.Context) error {
	var data models.KategoriTandaSakit
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request"})
	}

	if err := c.useCase.CreateKategori(&data); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, echo.Map{"data": data})
}

func (c *PemantauanAnakController) UpdateKategori(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var data models.KategoriTandaSakit
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request"})
	}

	if err := c.useCase.UpdateKategori(uint(id), &data); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"data": data})
}

func (c *PemantauanAnakController) DeleteKategori(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err := c.useCase.DeleteKategori(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "Indikator berhasil dihapus"})
}

func (c *PemantauanAnakController) Verifikasi(ctx echo.Context) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var req models.LembarPemantauanVerifikasiRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request"})
	}

	if err := req.Validate(); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	if err := c.useCase.VerifikasiPemantauan(uint(id), &req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "Verifikasi pemantauan berhasil disimpan"})
}
