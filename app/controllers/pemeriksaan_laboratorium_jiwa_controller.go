package controllers

import (
	"net/http"
	"strconv"
	"time"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type PemeriksaanLaboratoriumJiwaController struct {
	usecase usecases.PemeriksaanLaboratoriumJiwaUsecase
}

func NewPemeriksaanLaboratoriumJiwaController(u usecases.PemeriksaanLaboratoriumJiwaUsecase) *PemeriksaanLaboratoriumJiwaController {
	return &PemeriksaanLaboratoriumJiwaController{usecase: u}
}

type createLabJiwaRequest struct {
	KehamilanID                      int32   `json:"kehamilan_id"`
	TanggalLab                       string  `json:"tanggal_lab"`
	LabHemoglobinHasil               float64 `json:"lab_hemoglobin_hasil"`
	LabHemoglobinRencanaTindakLanjut string  `json:"lab_hemoglobin_rencana_tindak_lanjut"`
	LabGolonganDarahRhesusHasil      string  `json:"lab_golongan_darah_rhesus_hasil"`
	LabGolonganDarahRhesusRencana    string  `json:"lab_golongan_darah_rhesus_rencana_tindak_lanjut"`
	LabGulaDarahSewaktuHasil         int     `json:"lab_gula_darah_sewaktu_hasil"`
	LabGulaDarahSewaktuRencana       string  `json:"lab_gula_darah_sewaktu_rencana_tindak_lanjut"`
	LabHIVHasil                      string  `json:"lab_hiv_hasil"`
	LabHIVRencana                    string  `json:"lab_hiv_rencana_tindak_lanjut"`
	LabSifilisHasil                  string  `json:"lab_sifilis_hasil"`
	LabSifilisRencana                string  `json:"lab_sifilis_rencana_tindak_lanjut"`
	LabHepatitisBHasil               string  `json:"lab_hepatitis_b_hasil"`
	LabHepatitisBRencana             string  `json:"lab_hepatitis_b_rencana_tindak_lanjut"`
	TanggalSkriningJiwa              string  `json:"tanggal_skrining_jiwa"`
	SkriningJiwaHasil                string  `json:"skrining_jiwa_hasil"`
	SkriningJiwaTindakLanjut         string  `json:"skrining_jiwa_tindak_lanjut"`
	SkriningJiwaPerluRujukan         string  `json:"skrining_jiwa_perlu_rujukan"`
	Kesimpulan                       string  `json:"kesimpulan"`
	Rekomendasi                      string  `json:"rekomendasi"`
}

func (c *PemeriksaanLaboratoriumJiwaController) Create(ctx echo.Context) error {
	claims, _ := ctx.Get("auth_claims").(*models.AuthClaims)
	if claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
	}
	var req createLabJiwaRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	p := &models.PemeriksaanLaboratoriumJiwa{
		KehamilanID:                      req.KehamilanID,
		LabHemoglobinHasil:               &req.LabHemoglobinHasil,
		LabHemoglobinRencanaTindakLanjut: req.LabHemoglobinRencanaTindakLanjut,
		LabGolonganDarahRhesusHasil:      req.LabGolonganDarahRhesusHasil,
		LabGolonganDarahRhesusRencana:    req.LabGolonganDarahRhesusRencana,
		LabGulaDarahSewaktuHasil:         &req.LabGulaDarahSewaktuHasil,
		LabGulaDarahSewaktuRencana:       req.LabGulaDarahSewaktuRencana,
		LabHIVHasil:                      req.LabHIVHasil,
		LabHIVRencana:                    req.LabHIVRencana,
		LabSifilisHasil:                  req.LabSifilisHasil,
		LabSifilisRencana:                req.LabSifilisRencana,
		LabHepatitisBHasil:               req.LabHepatitisBHasil,
		LabHepatitisBRencana:             req.LabHepatitisBRencana,
		SkriningJiwaHasil:                req.SkriningJiwaHasil,
		SkriningJiwaTindakLanjut:         req.SkriningJiwaTindakLanjut,
		SkriningJiwaPerluRujukan:         req.SkriningJiwaPerluRujukan,
		Kesimpulan:                       req.Kesimpulan,
		Rekomendasi:                      req.Rekomendasi,
	}
	if req.TanggalLab != "" {
		if t, err := time.Parse("2006-01-02", req.TanggalLab); err == nil {
			p.TanggalLab = &t
		}
	}
	if req.TanggalSkriningJiwa != "" {
		if t, err := time.Parse("2006-01-02", req.TanggalSkriningJiwa); err == nil {
			p.TanggalSkriningJiwa = &t
		}
	}
	if err := c.usecase.Create(p); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusCreated, models.Response{StatusCode: http.StatusCreated, Data: p})
}

// GetByID, GetByKehamilanID, Update, Delete sama seperti pola sebelumnya
func (c *PemeriksaanLaboratoriumJiwaController) GetByID(ctx echo.Context) error {
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

func (c *PemeriksaanLaboratoriumJiwaController) GetByKehamilanID(ctx echo.Context) error {
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

func (c *PemeriksaanLaboratoriumJiwaController) Update(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	var req createLabJiwaRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	existing, err := c.usecase.GetByID(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: "Data tidak ditemukan"})
	}
	if req.LabHemoglobinHasil != 0 {
		existing.LabHemoglobinHasil = &req.LabHemoglobinHasil
	}
	if req.LabHemoglobinRencanaTindakLanjut != "" {
		existing.LabHemoglobinRencanaTindakLanjut = req.LabHemoglobinRencanaTindakLanjut
	}
	// ... update semua field lainnya
	if err := c.usecase.Update(existing); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: existing})
}

func (c *PemeriksaanLaboratoriumJiwaController) Delete(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	if err := c.usecase.Delete(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Message: "deleted"})
}
