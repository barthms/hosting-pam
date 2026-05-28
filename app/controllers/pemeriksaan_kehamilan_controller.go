package controllers

import (
	"net/http"
	"strconv"
	"time"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type PemeriksaanKehamilanController struct {
	usecase usecases.PemeriksaanKehamilanUsecase
}

func NewPemeriksaanKehamilanController(u usecases.PemeriksaanKehamilanUsecase) *PemeriksaanKehamilanController {
	return &PemeriksaanKehamilanController{usecase: u}
}

// ================= REQUEST =================

type createPemeriksaanKehamilanRequest struct {
	KehamilanID            int32    `json:"kehamilan_id"`
	KunjunganKe            int32    `json:"kunjungan_ke"`
	MingguKehamilan        int32    `json:"minggu_kehamilan"`
	TanggalPeriksa         string   `json:"tanggal_periksa"`
	TempatPeriksa          string   `json:"tempat_periksa"`
	BeratBadan             *float64 `json:"berat_badan"`
	TinggiBadan            *float64 `json:"tinggi_badan"`
	LingkarLenganAtas      *float64 `json:"lingkar_lengan_atas"`
	Sistole                int32    `json:"sistole"`
	Diastole               int32    `json:"diastole"`
	TinggiRahim            *float64 `json:"tinggi_rahim"`
	DenyutJantungJanin     int32    `json:"denyut_jantung_janin"`
	StatusImunisasiTetanus string   `json:"status_imunisasi_tetanus"`
	TabletTambahDarah      *int32   `json:"tablet_tambah_darah"`
	TesLabHb               *float64 `json:"tes_lab_hb"`
	TesLabProteinUrine     string   `json:"tes_lab_protein_urine"`
	TesLabGulaDarah        *int32   `json:"tes_lab_gula_darah"`
	USG                    string   `json:"usg"`
	TripelEliminasi        string   `json:"tripel_eliminasi"`
	TataLaksanaKasus       string   `json:"tata_laksana_kasus"`
	LetakDenyutJantungBayi string   `json:"letak_denyut_jantung_bayi"`
	Konseling              string   `json:"konseling"`
	SkriningDokter         string   `json:"skrining_dokter"`
	TesGolonganDarah       string   `json:"tes_golongan_darah"`
	Trimester              string   `json:"trimester"`
}

type updatePemeriksaanKehamilanRequest struct {
	KunjunganKe        *int32   `json:"kunjungan_ke"`
	MingguKehamilan    *int32   `json:"minggu_kehamilan"`
	TanggalPeriksa     *string  `json:"tanggal_periksa"`
	TempatPeriksa      *string  `json:"tempat_periksa"`
	BeratBadan         *float64 `json:"berat_badan"`
	TinggiBadan        *float64 `json:"tinggi_badan"`
	LingkarLenganAtas  *float64 `json:"lingkar_lengan_atas"`
	Sistole            *int32   `json:"sistole"`
	Diastole           *int32   `json:"diastole"`
	TinggiRahim        *float64 `json:"tinggi_rahim"`
	DenyutJantungJanin *int32   `json:"denyut_jantung_janin"`
	TesLabHb           *float64 `json:"tes_lab_hb"`
	TesLabProteinUrine *string  `json:"tes_lab_protein_urine"`
	TesLabGulaDarah    *int32   `json:"tes_lab_gula_darah"`
	USG                *string  `json:"usg"`
	TripelEliminasi    *string  `json:"tripel_eliminasi"`
}

// ================= CREATE =================

func (c *PemeriksaanKehamilanController) Create(ctx echo.Context) error {
	var req createPemeriksaanKehamilanRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	p := &models.PemeriksaanKehamilan{
		KehamilanID:            req.KehamilanID,
		KunjunganKe:            req.KunjunganKe,
		MingguKehamilan:        req.MingguKehamilan,
		TempatPeriksa:          req.TempatPeriksa,
		BeratBadan:             req.BeratBadan,
		TinggiBadan:            req.TinggiBadan,
		LingkarLenganAtas:      req.LingkarLenganAtas,
		Sistole:                req.Sistole,
		Diastole:               req.Diastole,
		TinggiRahim:            req.TinggiRahim,
		DenyutJantungJanin:     req.DenyutJantungJanin,
		StatusImunisasiTetanus: req.StatusImunisasiTetanus,
		TabletTambahDarah:      req.TabletTambahDarah,
		TesLabHb:               req.TesLabHb,
		TesLabProteinUrine:     req.TesLabProteinUrine,
		TesLabGulaDarah:        req.TesLabGulaDarah,
		USG:                    req.USG,
		TripelEliminasi:        req.TripelEliminasi,
		TataLaksanaKasus:       req.TataLaksanaKasus,
		LetakDenyutJantungBayi: req.LetakDenyutJantungBayi,
		Konseling:              req.Konseling,
		SkriningDokter:         req.SkriningDokter,
		TesGolonganDarah:       req.TesGolonganDarah,
		Trimester:              req.Trimester,
	}

	if req.TanggalPeriksa != "" {
		if t, err := time.Parse("2006-01-02", req.TanggalPeriksa); err == nil {
			p.TanggalPeriksa = &t
		}
	}

	if err := c.usecase.Create(p); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, models.Response{
		StatusCode: http.StatusCreated,
		Data:       p,
	})
}

// ================= GET BY ID =================

func (c *PemeriksaanKehamilanController) GetByID(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid id",
		})
	}

	data, err := c.usecase.GetByID(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Data:       data,
	})
}

// ================= GET BY KEHAMILAN =================

func (c *PemeriksaanKehamilanController) GetByKehamilanID(ctx echo.Context) error {
	kehamilanID, err := strconv.ParseInt(ctx.QueryParam("kehamilan_id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "kehamilan_id required",
		})
	}

	list, err := c.usecase.GetByKehamilanID(int32(kehamilanID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Data:       list,
	})
}

// ================= UPDATE =================

func (c *PemeriksaanKehamilanController) Update(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid id",
		})
	}

	var req updatePemeriksaanKehamilanRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	existing, err := c.usecase.GetByID(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{
			StatusCode: http.StatusNotFound,
			Message:    "data tidak ditemukan",
		})
	}

	if req.KunjunganKe != nil {
		existing.KunjunganKe = *req.KunjunganKe
	}
	if req.MingguKehamilan != nil {
		existing.MingguKehamilan = *req.MingguKehamilan
	}
	if req.TanggalPeriksa != nil {
		if t, err := time.Parse("2006-01-02", *req.TanggalPeriksa); err == nil {
			existing.TanggalPeriksa = &t
		}
	}
	if req.TempatPeriksa != nil {
		existing.TempatPeriksa = *req.TempatPeriksa
	}
	if req.BeratBadan != nil {
		existing.BeratBadan = req.BeratBadan
	}
	if req.TinggiBadan != nil {
		existing.TinggiBadan = req.TinggiBadan
	}
	if req.LingkarLenganAtas != nil {
		existing.LingkarLenganAtas = req.LingkarLenganAtas
	}
	if req.Sistole != nil {
		existing.Sistole = *req.Sistole
	}
	if req.Diastole != nil {
		existing.Diastole = *req.Diastole
	}
	if req.TinggiRahim != nil {
		existing.TinggiRahim = req.TinggiRahim
	}
	if req.DenyutJantungJanin != nil {
		existing.DenyutJantungJanin = *req.DenyutJantungJanin
	}
	if req.TesLabHb != nil {
		existing.TesLabHb = req.TesLabHb
	}
	if req.TesLabProteinUrine != nil {
		existing.TesLabProteinUrine = *req.TesLabProteinUrine
	}
	if req.TesLabGulaDarah != nil {
		existing.TesLabGulaDarah = req.TesLabGulaDarah
	}
	if req.USG != nil {
		existing.USG = *req.USG
	}
	if req.TripelEliminasi != nil {
		existing.TripelEliminasi = *req.TripelEliminasi
	}

	if err := c.usecase.Update(existing); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Data:       existing,
	})
}

// ================= DELETE =================

func (c *PemeriksaanKehamilanController) Delete(ctx echo.Context) error {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	if err := c.usecase.Delete(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "deleted",
	})
}
func (c *PemeriksaanKehamilanController) GetGrafikANC(ctx echo.Context) error {

	kehamilanID, err := strconv.ParseInt(ctx.QueryParam("kehamilan_id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "kehamilan_id required",
		})
	}

	data, err := c.usecase.GetGrafikANC(int32(kehamilanID))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{
			StatusCode: http.StatusNotFound,
			Message:    "data tidak ditemukan",
		})
	}
	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Data:       data,
	})
}

// MODUL IBU
func (c *PemeriksaanKehamilanController) GetByIDForOrangtua(ctx echo.Context) error {
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

// MODUL IBU
func (c *PemeriksaanKehamilanController) GetMine(ctx echo.Context) error {
	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
	if !ok || claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "token tidak valid",
		})
	}

	data, err := c.usecase.GetMine(claims.UserID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Data:       data,
	})
}
