package controllers

import (
	"net/http"
	"strconv"
	"time"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type CatatanPelayananTrimester1Controller struct {
	usecase usecases.CatatanPelayananTrimester1Usecase
}

func NewCatatanPelayananTrimester1Controller(u usecases.CatatanPelayananTrimester1Usecase) *CatatanPelayananTrimester1Controller {
	return &CatatanPelayananTrimester1Controller{usecase: u}
}

type createCatatanT1Request struct {
	KehamilanID                     int32  `json:"kehamilan_id"`
	TanggalPeriksaStampParaf        string `json:"tanggal_periksa_stamp_paraf"`
	KeluhanPemeriksaanTindakanSaran string `json:"keluhan_pemeriksaan_tindakan_saran"`
	TanggalKembali                  string `json:"tanggal_kembali"`
}

// Create menambah catatan pelayanan trimester 1 baru dan mengembalikan data yang baru dibuat dengan ID dari database
func (c *CatatanPelayananTrimester1Controller) Create(ctx echo.Context) error {
	claims, _ := ctx.Get("auth_claims").(*models.AuthClaims)
	if claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
	}
	var req createCatatanT1Request
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}

	catatan := &models.CatatanPelayananTrimester1{
		KehamilanID:                     req.KehamilanID,
		KeluhanPemeriksaanTindakanSaran: req.KeluhanPemeriksaanTindakanSaran,
	}

	if req.TanggalPeriksaStampParaf != "" {
		if t, err := time.Parse("2006-01-02", req.TanggalPeriksaStampParaf); err == nil {
			catatan.TanggalPeriksaStampParaf = &t
		}
	}

	if req.TanggalKembali != "" {
		if t, err := time.Parse("2006-01-02", req.TanggalKembali); err == nil {
			catatan.TanggalKembali = &t
		}
	}

	// Usecase Create sekarang mengembalikan data yang sudah tersimpan di database dengan ID
	createdCatatan, err := c.usecase.Create(catatan)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, models.Response{StatusCode: http.StatusCreated, Data: createdCatatan})
}

// GetByID mengambil data catatan berdasarkan ID
func (c *CatatanPelayananTrimester1Controller) GetByID(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}

	data, err := c.usecase.GetByID(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: data})
}

// GetByKehamilanID mengambil semua catatan untuk satu kehamilan
func (c *CatatanPelayananTrimester1Controller) GetByKehamilanID(ctx echo.Context) error {
	kehamilanID, err := strconv.ParseInt(ctx.QueryParam("kehamilan_id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "kehamilan_id required"})
	}

	list, err := c.usecase.GetByKehamilanID(int32(kehamilanID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: list})
}

// Update memperbarui catatan pelayanan dan mengembalikan data yang diperbarui dari database
func (c *CatatanPelayananTrimester1Controller) Update(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}

	var req createCatatanT1Request
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}

	// Ambil data yang ada terlebih dahulu
	existing, err := c.usecase.GetByID(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: "Data tidak ditemukan"})
	}

	// Update field yang dikirim
	if req.KeluhanPemeriksaanTindakanSaran != "" {
		existing.KeluhanPemeriksaanTindakanSaran = req.KeluhanPemeriksaanTindakanSaran
	}

	if req.TanggalPeriksaStampParaf != "" {
		if t, err := time.Parse("2006-01-02", req.TanggalPeriksaStampParaf); err == nil {
			existing.TanggalPeriksaStampParaf = &t
		}
	}

	if req.TanggalKembali != "" {
		if t, err := time.Parse("2006-01-02", req.TanggalKembali); err == nil {
			existing.TanggalKembali = &t
		}
	}

	// Usecase Update sekarang mengembalikan data yang sudah diperbarui di database
	updatedCatatan, err := c.usecase.Update(existing)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: updatedCatatan})
}

// Delete menghapus catatan pelayanan
func (c *CatatanPelayananTrimester1Controller) Delete(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}

	if err := c.usecase.Delete(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Message: "Data berhasil dihapus"})
}
