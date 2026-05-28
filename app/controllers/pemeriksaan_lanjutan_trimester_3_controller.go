package controllers

import (
	"net/http"
	"strconv"
	"time"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type PemeriksaanLanjutanTrimester3Controller struct {
	usecase usecases.PemeriksaanLanjutanTrimester3Usecase
}

func NewPemeriksaanLanjutanTrimester3Controller(u usecases.PemeriksaanLanjutanTrimester3Usecase) *PemeriksaanLanjutanTrimester3Controller {
	return &PemeriksaanLanjutanTrimester3Controller{usecase: u}
}

type createLanjutanT3Request struct {
	KehamilanID                           int32   `json:"kehamilan_id"`
	HasilUSGCatatan                       string  `json:"hasil_usg_catatan"`
	TanggalLab                            string  `json:"tanggal_lab"`
	LabHemoglobinHasil                    float64 `json:"lab_hemoglobin_hasil"`
	LabHemoglobinRencana                  string  `json:"lab_hemoglobin_rencana_tindak_lanjut"`
	LabProteinUrinHasil                   int     `json:"lab_protein_urin_hasil"`
	LabProteinUrinRencana                 string  `json:"lab_protein_urin_rencana_tindak_lanjut"`
	LabUrinReduksiHasil                   string  `json:"lab_urin_reduksi_hasil"`
	LabUrinReduksiRencana                 string  `json:"lab_urin_reduksi_rencana_tindak_lanjut"`
	TanggalSkriningJiwa                   string  `json:"tanggal_skrining_jiwa"`
	SkriningJiwaHasil                     string  `json:"skrining_jiwa_hasil"`
	SkriningJiwaTindakLanjut              string  `json:"skrining_jiwa_tindak_lanjut"`
	SkriningJiwaPerluRujukan              string  `json:"skrining_jiwa_perlu_rujukan"`
	RencanaKonsultasiGizi                 bool    `json:"rencana_konsultasi_gizi"`
	RencanaKonsultasiKebidanan            bool    `json:"rencana_konsultasi_kebidanan"`
	RencanaKonsultasiAnak                 bool    `json:"rencana_konsultasi_anak"`
	RencanaKonsultasiPenyakitDalam        bool    `json:"rencana_konsultasi_penyakit_dalam"`
	RencanaKonsultasiNeurologi            bool    `json:"rencana_konsultasi_neurologi"`
	RencanaKonsultasiTHT                  bool    `json:"rencana_konsultasi_tht"`
	RencanaKonsultasiPsikiatri            bool    `json:"rencana_konsultasi_psikiatri"`
	RencanaKonsultasiLainLain             string  `json:"rencana_konsultasi_lain_lain"`
	RencanaProsesMelahirkan               string  `json:"rencana_proses_melahirkan"`
	RencanaKontrasepsiAKDR                bool    `json:"rencana_kontrasepsi_akdr"`
	RencanaKontrasepsiPil                 bool    `json:"rencana_kontrasepsi_pil"`
	RencanaKontrasepsiSuntik              bool    `json:"rencana_kontrasepsi_suntik"`
	RencanaKontrasepsiSteril              bool    `json:"rencana_kontrasepsi_steril"`
	RencanaKontrasepsiMAL                 bool    `json:"rencana_kontrasepsi_mal"`
	RencanaKontrasepsiImplan              bool    `json:"rencana_kontrasepsi_implan"`
	RencanaKontrasepsiBelumMemilih        bool    `json:"rencana_kontrasepsi_belum_memilih"`
	KebutuhanKonseling                    string  `json:"kebutuhan_konseling"`
	Penjelasan                            string  `json:"penjelasan"`
	KesimpulanRekomendasiTempatMelahirkan string  `json:"kesimpulan_rekomendasi_tempat_melahirkan"`
}

func (c *PemeriksaanLanjutanTrimester3Controller) Create(ctx echo.Context) error {
	claims, _ := ctx.Get("auth_claims").(*models.AuthClaims)
	if claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
	}
	var req createLanjutanT3Request
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	p := &models.PemeriksaanLanjutanTrimester3{
		KehamilanID:                           req.KehamilanID,
		HasilUSGCatatan:                       req.HasilUSGCatatan,
		LabHemoglobinHasil:                    &req.LabHemoglobinHasil,
		LabHemoglobinRencana:                  req.LabHemoglobinRencana,
		LabProteinUrinHasil:                   &req.LabProteinUrinHasil,
		LabProteinUrinRencana:                 req.LabProteinUrinRencana,
		LabUrinReduksiHasil:                   req.LabUrinReduksiHasil,
		LabUrinReduksiRencana:                 req.LabUrinReduksiRencana,
		SkriningJiwaHasil:                     req.SkriningJiwaHasil,
		SkriningJiwaTindakLanjut:              req.SkriningJiwaTindakLanjut,
		SkriningJiwaPerluRujukan:              req.SkriningJiwaPerluRujukan,
		RencanaKonsultasiGizi:                 req.RencanaKonsultasiGizi,
		RencanaKonsultasiKebidanan:            req.RencanaKonsultasiKebidanan,
		RencanaKonsultasiAnak:                 req.RencanaKonsultasiAnak,
		RencanaKonsultasiPenyakitDalam:        req.RencanaKonsultasiPenyakitDalam,
		RencanaKonsultasiNeurologi:            req.RencanaKonsultasiNeurologi,
		RencanaKonsultasiTHT:                  req.RencanaKonsultasiTHT,
		RencanaKonsultasiPsikiatri:            req.RencanaKonsultasiPsikiatri,
		RencanaKonsultasiLainLain:             req.RencanaKonsultasiLainLain,
		RencanaProsesMelahirkan:               req.RencanaProsesMelahirkan,
		RencanaKontrasepsiAKDR:                req.RencanaKontrasepsiAKDR,
		RencanaKontrasepsiPil:                 req.RencanaKontrasepsiPil,
		RencanaKontrasepsiSuntik:              req.RencanaKontrasepsiSuntik,
		RencanaKontrasepsiSteril:              req.RencanaKontrasepsiSteril,
		RencanaKontrasepsiMAL:                 req.RencanaKontrasepsiMAL,
		RencanaKontrasepsiImplan:              req.RencanaKontrasepsiImplan,
		RencanaKontrasepsiBelumMemilih:        req.RencanaKontrasepsiBelumMemilih,
		KebutuhanKonseling:                    req.KebutuhanKonseling,
		Penjelasan:                            req.Penjelasan,
		KesimpulanRekomendasiTempatMelahirkan: req.KesimpulanRekomendasiTempatMelahirkan,
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

// GetByID, GetByKehamilanID, Update, Delete (sama seperti pola sebelumnya)
func (c *PemeriksaanLanjutanTrimester3Controller) GetByID(ctx echo.Context) error {
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

func (c *PemeriksaanLanjutanTrimester3Controller) GetByKehamilanID(ctx echo.Context) error {
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

func (c *PemeriksaanLanjutanTrimester3Controller) Update(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	var req createLanjutanT3Request
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	existing, err := c.usecase.GetByID(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: "Data tidak ditemukan"})
	}
	// Update field (contoh)
	if req.HasilUSGCatatan != "" {
		existing.HasilUSGCatatan = req.HasilUSGCatatan
	}
	if req.LabHemoglobinHasil != 0 {
		existing.LabHemoglobinHasil = &req.LabHemoglobinHasil
	}
	// ... update semua field lainnya
	if err := c.usecase.Update(existing); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: existing})
}

func (c *PemeriksaanLanjutanTrimester3Controller) Delete(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	if err := c.usecase.Delete(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Message: "deleted"})
}
