package controllers

import (
	"net/http"
	"strconv"
	"time"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type CatatanPelayananTrimester2Controller struct {
	usecase usecases.CatatanPelayananTrimester2Usecase
}

func NewCatatanPelayananTrimester2Controller(u usecases.CatatanPelayananTrimester2Usecase) *CatatanPelayananTrimester2Controller {
	return &CatatanPelayananTrimester2Controller{usecase: u}
}

type createCatatanT2Request struct {
	KehamilanID                     int32  `json:"kehamilan_id"`
	TanggalPeriksaStampParaf        string `json:"tanggal_periksa_stamp_paraf"`
	KeluhanPemeriksaanTindakanSaran string `json:"keluhan_pemeriksaan_tindakan_saran"`
	TanggalKembali                  string `json:"tanggal_kembali"`
}

func (c *CatatanPelayananTrimester2Controller) Create(ctx echo.Context) error {
	claims, _ := ctx.Get("auth_claims").(*models.AuthClaims)
	if claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
	}
	var req createCatatanT2Request
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	catatan := &models.CatatanPelayananTrimester2{
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
	if err := c.usecase.Create(catatan); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusCreated, models.Response{StatusCode: http.StatusCreated, Data: catatan})
}

func (c *CatatanPelayananTrimester2Controller) GetByID(ctx echo.Context) error {
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

func (c *CatatanPelayananTrimester2Controller) GetByKehamilanID(ctx echo.Context) error {
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

func (c *CatatanPelayananTrimester2Controller) Update(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	var req createCatatanT1Request
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

func (c *CatatanPelayananTrimester2Controller) Delete(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	if err := c.usecase.Delete(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Message: "deleted"})
}
