package controllers

import (
	"net/http"
	"strconv"
	"time"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type RujukanController struct {
	usecase usecases.RujukanUsecase
}

func NewRujukanController(u usecases.RujukanUsecase) *RujukanController {
	return &RujukanController{usecase: u}
}

type createRujukanRequest struct {
	KehamilanID                              int32  `json:"kehamilan_id"`
	RujukanResumePemeriksaanTatalaksana      string `json:"rujukan_resume_pemeriksaan_tatalaksana"`
	RujukanDiagnosisAkhir                    string `json:"rujukan_diagnosis_akhir"`
	RujukanAlasanDirujukKeFKRTL              string `json:"rujukan_alasan_dirujuk_ke_fkrtl"`
	RujukanBalikTanggal                      string `json:"rujukan_balik_tanggal"`
	RujukanBalikDiagnosisAkhir               string `json:"rujukan_balik_diagnosis_akhir"`
	RujukanBalikResumePemeriksaanTatalaksana string `json:"rujukan_balik_resume_pemeriksaan_tatalaksana"`
	AnjuranRekomendasiTempatMelahirkan       string `json:"anjuran_rekomendasi_tempat_melahirkan"`
}

func (c *RujukanController) Create(ctx echo.Context) error {
	claims, _ := ctx.Get("auth_claims").(*models.AuthClaims)
	if claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
	}
	var req createRujukanRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	r := &models.Rujukan{
		KehamilanID:                              req.KehamilanID,
		RujukanResumePemeriksaanTatalaksana:      req.RujukanResumePemeriksaanTatalaksana,
		RujukanDiagnosisAkhir:                    req.RujukanDiagnosisAkhir,
		RujukanAlasanDirujukKeFKRTL:              req.RujukanAlasanDirujukKeFKRTL,
		RujukanBalikDiagnosisAkhir:               req.RujukanBalikDiagnosisAkhir,
		RujukanBalikResumePemeriksaanTatalaksana: req.RujukanBalikResumePemeriksaanTatalaksana,
		AnjuranRekomendasiTempatMelahirkan:       req.AnjuranRekomendasiTempatMelahirkan,
	}
	if req.RujukanBalikTanggal != "" {
		if t, err := time.Parse("2006-01-02", req.RujukanBalikTanggal); err == nil {
			r.RujukanBalikTanggal = &t
		}
	}
	if err := c.usecase.Create(r); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusCreated, models.Response{StatusCode: http.StatusCreated, Data: r})
}

// GetByID, GetByKehamilanID, Update, Delete (sama seperti pola sebelumnya)
func (c *RujukanController) GetByID(ctx echo.Context) error {
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

func (c *RujukanController) GetByKehamilanID(ctx echo.Context) error {
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

func (c *RujukanController) Update(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	var req createRujukanRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	existing, err := c.usecase.GetByID(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: "Data tidak ditemukan"})
	}
	// Update semua field
	existing.RujukanResumePemeriksaanTatalaksana = req.RujukanResumePemeriksaanTatalaksana
	existing.RujukanDiagnosisAkhir = req.RujukanDiagnosisAkhir
	existing.RujukanAlasanDirujukKeFKRTL = req.RujukanAlasanDirujukKeFKRTL
	existing.RujukanBalikDiagnosisAkhir = req.RujukanBalikDiagnosisAkhir
	existing.RujukanBalikResumePemeriksaanTatalaksana = req.RujukanBalikResumePemeriksaanTatalaksana
	existing.AnjuranRekomendasiTempatMelahirkan = req.AnjuranRekomendasiTempatMelahirkan

	// Update tanggal rujukan balik jika ada
	if req.RujukanBalikTanggal != "" {
		if t, err := time.Parse("2006-01-02", req.RujukanBalikTanggal); err == nil {
			existing.RujukanBalikTanggal = &t
		}
	} else {
		existing.RujukanBalikTanggal = nil
	}

	if err := c.usecase.Update(existing); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: existing})
}

func (c *RujukanController) Delete(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	if err := c.usecase.Delete(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Message: "deleted"})
}

//.MODUL IBU
func (c *RujukanController) GetByIDForOrangtua(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid id",
		})
	}

	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
	if !ok || claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "token tidak valid",
		})
	}

	data, err := c.usecase.GetByIDForOrangtua(int32(id), claims.UserID)
	if err != nil {
		return ctx.JSON(http.StatusForbidden, models.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Data:       data,
	})
}