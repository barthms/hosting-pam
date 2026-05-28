package controllers

import (
	"net/http"
	"strconv"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type EdukasiMPASIController struct {
	usecase usecases.EdukasiMPASIUsecase
}

func NewEdukasiMPASIController(u usecases.EdukasiMPASIUsecase) *EdukasiMPASIController {
	return &EdukasiMPASIController{u}
}

// ==================== MATERI ====================
func (c *EdukasiMPASIController) CreateMateri(ctx echo.Context) error {
	var input models.MateriMPASI
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := c.usecase.CreateMateri(&input); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusCreated, input)
}

func (c *EdukasiMPASIController) GetMateriAll(ctx echo.Context) error {
	data, err := c.usecase.GetMateriAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, data)
}

func (c *EdukasiMPASIController) GetMateriByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}
	data, err := c.usecase.GetMateriByID(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Data not found"})
	}
	return ctx.JSON(http.StatusOK, data)
}

func (c *EdukasiMPASIController) UpdateMateri(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}
	var input models.MateriMPASI
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := c.usecase.UpdateMateri(int32(id), &input); err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "updated successfully"})
}

func (c *EdukasiMPASIController) DeleteMateri(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}
	if err := c.usecase.DeleteMateri(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "deleted successfully"})
}

func (c *EdukasiMPASIController) GetMateriByBulan(ctx echo.Context) error {
	bulan, err := strconv.Atoi(ctx.Param("bulan"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid bulan"})
	}
	data, err := c.usecase.GetMateriByBulan(bulan)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, data)
}

// ==================== ATURAN PORSI ====================
func (c *EdukasiMPASIController) CreateAturanPorsi(ctx echo.Context) error {
	var input models.AturanPorsiMPASI
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := c.usecase.CreateAturanPorsi(&input); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusCreated, input)
}

func (c *EdukasiMPASIController) GetAturanPorsiAll(ctx echo.Context) error {
	data, err := c.usecase.GetAturanPorsiAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, data)
}

func (c *EdukasiMPASIController) GetAturanPorsiByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}
	data, err := c.usecase.GetAturanPorsiByID(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Data not found"})
	}
	return ctx.JSON(http.StatusOK, data)
}

func (c *EdukasiMPASIController) UpdateAturanPorsi(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}
	var input models.AturanPorsiMPASI
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := c.usecase.UpdateAturanPorsi(int32(id), &input); err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "updated successfully"})
}

func (c *EdukasiMPASIController) DeleteAturanPorsi(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}
	if err := c.usecase.DeleteAturanPorsi(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "deleted successfully"})
}

// ==================== JADWAL ====================
func (c *EdukasiMPASIController) CreateJadwal(ctx echo.Context) error {
	var input models.JadwalHarianMPASI
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := c.usecase.CreateJadwal(&input); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusCreated, input)
}

func (c *EdukasiMPASIController) GetJadwalAll(ctx echo.Context) error {
	data, err := c.usecase.GetJadwalAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, data)
}

func (c *EdukasiMPASIController) GetJadwalByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}
	data, err := c.usecase.GetJadwalByID(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Data not found"})
	}
	return ctx.JSON(http.StatusOK, data)
}

func (c *EdukasiMPASIController) UpdateJadwal(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}
	var input models.JadwalHarianMPASI
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := c.usecase.UpdateJadwal(int32(id), &input); err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "updated successfully"})
}

func (c *EdukasiMPASIController) DeleteJadwal(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}
	if err := c.usecase.DeleteJadwal(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "deleted successfully"})
}

// ==================== RESEP ====================
func (c *EdukasiMPASIController) CreateResep(ctx echo.Context) error {
	var input models.ResepMPASI
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := c.usecase.CreateResep(&input); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusCreated, input)
}

func (c *EdukasiMPASIController) GetResepAll(ctx echo.Context) error {
	data, err := c.usecase.GetResepAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, data)
}

func (c *EdukasiMPASIController) GetResepByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}
	data, err := c.usecase.GetResepByID(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Data not found"})
	}
	return ctx.JSON(http.StatusOK, data)
}

func (c *EdukasiMPASIController) UpdateResep(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}
	var input models.ResepMPASI
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := c.usecase.UpdateResep(int32(id), &input); err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "updated successfully"})
}

func (c *EdukasiMPASIController) DeleteResep(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}
	if err := c.usecase.DeleteResep(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "deleted successfully"})
}

func (c *EdukasiMPASIController) GetResepByBulan(ctx echo.Context) error {
	bulan, err := strconv.Atoi(ctx.Param("bulan"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid bulan"})
	}
	data, err := c.usecase.GetResepByBulan(bulan)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, data)
}

// ==================== AGGREGATE UNTUK IBU: Porsi & Jadwal ====================
func (c *EdukasiMPASIController) GetPorsiJadwalByBulan(ctx echo.Context) error {
	bulan, err := strconv.Atoi(ctx.Param("bulan"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid bulan"})
	}

	aturanPorsi, _ := c.usecase.GetAturanPorsiByBulan(bulan)
	jadwalHarian, _ := c.usecase.GetJadwalByBulan(bulan)

	return ctx.JSON(http.StatusOK, echo.Map{
		"aturan_porsi":  aturanPorsi,
		"jadwal_harian": jadwalHarian,
	})
}
