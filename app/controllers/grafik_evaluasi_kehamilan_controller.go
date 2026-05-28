package controllers

import (
	"net/http"
	"strconv"
	"time"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type GrafikEvaluasiKehamilanController struct {
	usecase usecases.GrafikEvaluasiKehamilanUsecase
}

func NewGrafikEvaluasiKehamilanController(u usecases.GrafikEvaluasiKehamilanUsecase) *GrafikEvaluasiKehamilanController {
	return &GrafikEvaluasiKehamilanController{usecase: u}
}

type createGrafikEvaluasiRequest struct {
	KehamilanID             *int32   `json:"kehamilan_id"`
	TanggalBulanTahun       *string  `json:"tanggal_bulan_tahun"`
	UsiaGestasiMinggu       *int     `json:"usia_gestasi_minggu"`
	TinggiFundusUteriCm     *float64 `json:"tinggi_fundus_uteri_cm"`
	DenyutJantungBayiXMenit *int     `json:"denyut_jantung_bayi_x_menit"`
	TekananDarahSistole     *int     `json:"tekanan_darah_sistole"`
	TekananDarahDiastole    *int     `json:"tekanan_darah_diastole"`
	NadiPerMenit            *int     `json:"nadi_per_menit"`
	GerakanBayi             *string  `json:"gerakan_bayi"`
	UrinProtein             *string  `json:"urin_protein"`
	UrinReduksi             *string  `json:"urin_reduksi"`
	Hemoglobin              *float64 `json:"hemoglobin"`
	TabletTambahDarah       *int     `json:"tablet_tambah_darah"`
	Kalsium                 *string  `json:"kalsium"`
	Aspirin                 *string  `json:"aspirin"`
}

// ====================== CREATE ======================
func (c *GrafikEvaluasiKehamilanController) Create(ctx echo.Context) error {

	if ctx.Get("auth_claims") == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Unauthorized",
		})
	}

	var req createGrafikEvaluasiRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	if req.KehamilanID == nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "kehamilan_id wajib diisi",
		})
	}

	g := &models.GrafikEvaluasiKehamilan{
		KehamilanID:             *req.KehamilanID,
		UsiaGestasiMinggu:       req.UsiaGestasiMinggu,
		TinggiFundusUteriCm:     req.TinggiFundusUteriCm,
		DenyutJantungBayiXMenit: req.DenyutJantungBayiXMenit,
		TekananDarahSistole:     req.TekananDarahSistole,
		TekananDarahDiastole:    req.TekananDarahDiastole,
		NadiPerMenit:            req.NadiPerMenit,
		GerakanBayi:             req.GerakanBayi,
		UrinProtein:             req.UrinProtein,
		UrinReduksi:             req.UrinReduksi,
		Hemoglobin:              req.Hemoglobin,
		TabletTambahDarah:       req.TabletTambahDarah,
		Kalsium:                 req.Kalsium,
		Aspirin:                 req.Aspirin,
	}

	// parse tanggal (optional)
	if req.TanggalBulanTahun != nil {
		t, err := time.Parse("2006-01-02", *req.TanggalBulanTahun)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, models.Response{
				StatusCode: http.StatusBadRequest,
				Message:    "format tanggal harus YYYY-MM-DD",
			})
		}
		g.TanggalBulanTahun = &t
	}

	if err := c.usecase.Create(g); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, models.Response{
		StatusCode: http.StatusCreated,
		Data:       g,
	})
}

// ====================== GET BY ID ======================
func (c *GrafikEvaluasiKehamilanController) GetByID(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.Param("id"))
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

// ====================== GET BY KEHAMILAN ======================
func (c *GrafikEvaluasiKehamilanController) GetByKehamilanID(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.QueryParam("kehamilan_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "kehamilan_id required",
		})
	}

	list, err := c.usecase.GetByKehamilanID(int32(id))
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

// ====================== UPDATE ======================
func (c *GrafikEvaluasiKehamilanController) Update(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid id",
		})
	}

	var req createGrafikEvaluasiRequest
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

	// update fields (safe pointer)
	if req.UsiaGestasiMinggu != nil {
		existing.UsiaGestasiMinggu = req.UsiaGestasiMinggu
	}
	if req.TinggiFundusUteriCm != nil {
		existing.TinggiFundusUteriCm = req.TinggiFundusUteriCm
	}
	if req.DenyutJantungBayiXMenit != nil {
		existing.DenyutJantungBayiXMenit = req.DenyutJantungBayiXMenit
	}
	if req.TekananDarahSistole != nil {
		existing.TekananDarahSistole = req.TekananDarahSistole
	}
	if req.TekananDarahDiastole != nil {
		existing.TekananDarahDiastole = req.TekananDarahDiastole
	}
	if req.NadiPerMenit != nil {
		existing.NadiPerMenit = req.NadiPerMenit
	}

	if req.GerakanBayi != nil {
		existing.GerakanBayi = req.GerakanBayi
	}
	if req.UrinProtein != nil {
		existing.UrinProtein = req.UrinProtein
	}
	if req.UrinReduksi != nil {
		existing.UrinReduksi = req.UrinReduksi
	}
	if req.Hemoglobin != nil {
		existing.Hemoglobin = req.Hemoglobin
	}
	if req.TabletTambahDarah != nil {
		existing.TabletTambahDarah = req.TabletTambahDarah
	}
	if req.Kalsium != nil {
		existing.Kalsium = req.Kalsium
	}
	if req.Aspirin != nil {
		existing.Aspirin = req.Aspirin
	}

	if req.TanggalBulanTahun != nil {
		t, err := time.Parse("2006-01-02", *req.TanggalBulanTahun)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, models.Response{
				StatusCode: http.StatusBadRequest,
				Message:    "format tanggal salah",
			})
		}
		existing.TanggalBulanTahun = &t
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

// ====================== DELETE ======================
func (c *GrafikEvaluasiKehamilanController) Delete(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid id",
		})
	}

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

// ====================== GET GRAFIK ======================
func (c *GrafikEvaluasiKehamilanController) GetGrafik(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.QueryParam("kehamilan_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "kehamilan_id required",
		})
	}

	data, err := c.usecase.GetGrafik(int32(id))
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

//
// ====================== MODUL IBU ====================
//

func (c *GrafikEvaluasiKehamilanController) GetMine(ctx echo.Context) error {
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

func (c *GrafikEvaluasiKehamilanController) GetByIDForOrangtua(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
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

func (c *GrafikEvaluasiKehamilanController) GetGrafikForOrangtua(ctx echo.Context) error {
	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
	if !ok || claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "token tidak valid",
		})
	}

	data, err := c.usecase.GetGrafikForOrangtua(claims.UserID)
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

//
// ====================== MODUL IBU ======================
//

func (c *GrafikEvaluasiKehamilanController) GetGrafikOnTheFlyForOrangtua(ctx echo.Context) error {
	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
	if !ok || claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "token tidak valid",
		})
	}

	data, err := c.usecase.GetGrafikOnTheFlyForOrangtua(claims.UserID)
	if err != nil {
		// Handle error spesifik untuk kehamilan tidak ditemukan
		if err.Error() == "kehamilan aktif tidak ditemukan" {
			return ctx.JSON(http.StatusNotFound, models.Response{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
			})
		}

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
