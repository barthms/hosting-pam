package controllers

import (
	"net/http"
	"strconv"
	"time"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type CatatanPelayananNifasController struct {
	usecase usecases.CatatanPelayananNifasUsecase
}

func NewCatatanPelayananNifasController(u usecases.CatatanPelayananNifasUsecase) *CatatanPelayananNifasController {
	return &CatatanPelayananNifasController{usecase: u}
}

// controllers/catatan_pelayanan_nifas_controller.go

type createCatatanNifasRequest struct {
	KehamilanID                     int32  `json:"kehamilan_id"`
	KunjunganKe                     string `json:"kunjungan_ke"`
	TanggalPeriksaStampParaf        string `json:"tanggal_periksa_stamp_paraf"`
	KeluhanPemeriksaanTindakanSaran string `json:"keluhan_pemeriksaan_tindakan_saran"`
	TanggalKembali                  string `json:"tanggal_kembali"`
	// JANGAN ADA ID DI SINI!
}
// controllers/catatan_pelayanan_nifas_controller.go

// controllers/catatan_pelayanan_nifas_controller.go

func (c *CatatanPelayananNifasController) Create(ctx echo.Context) error {
    claims, _ := ctx.Get("auth_claims").(*models.AuthClaims)
    if claims == nil {
        return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
    }
    
    var req createCatatanNifasRequest
    if err := ctx.Bind(&req); err != nil {
        return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
    }
    
    // Set default kunjungan_ke jika kosong
    kunjunganKe := req.KunjunganKe
    if kunjunganKe == "" {
        kunjunganKe = "KF1"
    }
    
    catatan := &models.CatatanPelayananNifas{
        KehamilanID:                     req.KehamilanID,
        KunjunganKe:                     kunjunganKe,
        KeluhanPemeriksaanTindakanSaran: req.KeluhanPemeriksaanTindakanSaran,
    }
    // JANGAN SET ID_CATATAN_NIFAS - biarkan auto increment
    
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
    
    if err := c.usecase.Create(catatan); err != nil {
        return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
    }
    return ctx.JSON(http.StatusCreated, models.Response{StatusCode: http.StatusCreated, Data: catatan})
}

func (c *CatatanPelayananNifasController) GetByID(ctx echo.Context) error {
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

func (c *CatatanPelayananNifasController) GetByKehamilanID(ctx echo.Context) error {
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

func (c *CatatanPelayananNifasController) Update(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	var req createCatatanNifasRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	existing, err := c.usecase.GetByID(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: "Data tidak ditemukan"})
	}
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
	if err := c.usecase.Update(existing); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: existing})
}

func (c *CatatanPelayananNifasController) GetMine(
	ctx echo.Context,
) error {

	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)

	if !ok || claims == nil {
		return ctx.JSON(
			http.StatusUnauthorized,
			models.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Unauthorized",
			},
		)
	}

	userID := claims.UserID

	data, err := c.usecase.GetMine(
		ctx.Request().Context(),
		userID,
	)

	if err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			models.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}

	return ctx.JSON(
		http.StatusOK,
		models.Response{
			StatusCode: http.StatusOK,
			Data:       data,
		},
	)
}

func (c *CatatanPelayananNifasController) Delete(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	if err := c.usecase.Delete(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Message: "deleted"})
}
