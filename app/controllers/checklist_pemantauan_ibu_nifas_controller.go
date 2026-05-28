package controllers

import (
	"net/http"
	"strconv"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type ChecklistPemantauanIbuNifasController struct {
	usecase *usecases.ChecklistPemantauanIbuNifasUsecase
}

func NewChecklistPemantauanIbuNifasController(usecase *usecases.ChecklistPemantauanIbuNifasUsecase) *ChecklistPemantauanIbuNifasController {
	return &ChecklistPemantauanIbuNifasController{
		usecase: usecase,
	}
}

func (ctrl *ChecklistPemantauanIbuNifasController) GetMine(c echo.Context) error {
	claims, ok := c.Get("auth_claims").(*models.AuthClaims)
	if !ok || claims == nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "claims token tidak ditemukan",
		})
	}

	hariNifasParam := c.QueryParam("hari_nifas")
	if hariNifasParam == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "hari_nifas wajib diisi",
		})
	}

	hariNifas64, err := strconv.ParseInt(hariNifasParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "hari_nifas tidak valid",
		})
	}

	hariNifas := int32(hariNifas64)
	if hariNifas < 1 || hariNifas > 42 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "hari_nifas harus antara 1 sampai 42",
		})
	}

	data, err := ctrl.usecase.GetByUserIDAndHariNifas(claims.UserID, hariNifas)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Checklist pemantauan ibu nifas tidak ditemukan",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Get checklist pemantauan ibu nifas berhasil",
		"data":    data,
	})
}

func (ctrl *ChecklistPemantauanIbuNifasController) GetFilledDays(c echo.Context) error {
	claims, ok := c.Get("auth_claims").(*models.AuthClaims)
	if !ok || claims == nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "claims token tidak ditemukan",
		})
	}

	days, err := ctrl.usecase.GetFilledDaysByUserID(claims.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Gagal mengambil daftar hari nifas",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Daftar hari nifas berhasil diambil",
		"data":    days,
	})
}

func (ctrl *ChecklistPemantauanIbuNifasController) SaveMine(c echo.Context) error {
	type Request struct {
		KehamilanID int32 `json:"kehamilan_id"`
		HariNifas int32 `json:"hari_nifas"`

		// NIFAS A
		PemeriksaanNifas     bool `json:"pemeriksaan_nifas"`
		KonsumsiVitaminA     bool `json:"konsumsi_vitamin_a"`
		PemenuhanGizi        bool `json:"pemenuhan_gizi"`
		DemamLebih38         bool `json:"demam_lebih_38"`
		SakitKepala          bool `json:"sakit_kepala"`
		PandanganKabur       bool `json:"pandangan_kabur"`
		NyeriUluHati         bool `json:"nyeri_ulu_hati"`
		MasalahKesehatanJiwa bool `json:"masalah_kesehatan_jiwa"`

		// NIFAS B
		JantungBerdebar    bool `json:"jantung_berdebar"`
		CairanJalanLahir   bool `json:"cairan_jalan_lahir"`
		NapasPendek        bool `json:"napas_pendek"`
		PayudaraBermasalah bool `json:"payudara_bermasalah"`
		GangguanBAK        bool `json:"gangguan_bak"`
		KelaminBermasalah  bool `json:"kelamin_bermasalah"`
		DarahNifasBerbau   bool `json:"darah_nifas_berbau"`
		PendarahanBerat    bool `json:"pendarahan_berat"`
		Keputihan          bool `json:"keputihan"`
	}

	req := new(Request)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Request tidak valid",
			"error":   err.Error(),
		})
	}

	if req.KehamilanID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "kehamilan_id wajib diisi",
		})
	}

	data := &models.ChecklistPemantauanIbuNifas{
		KehamilanID: req.KehamilanID,
		HariNifas:   req.HariNifas,

		PemeriksaanNifas:     req.PemeriksaanNifas,
		KonsumsiVitaminA:     req.KonsumsiVitaminA,
		PemenuhanGizi:        req.PemenuhanGizi,
		DemamLebih38:         req.DemamLebih38,
		SakitKepala:          req.SakitKepala,
		PandanganKabur:       req.PandanganKabur,
		NyeriUluHati:         req.NyeriUluHati,
		MasalahKesehatanJiwa: req.MasalahKesehatanJiwa,

		JantungBerdebar:    req.JantungBerdebar,
		CairanJalanLahir:   req.CairanJalanLahir,
		NapasPendek:        req.NapasPendek,
		PayudaraBermasalah: req.PayudaraBermasalah,
		GangguanBAK:        req.GangguanBAK,
		KelaminBermasalah:  req.KelaminBermasalah,
		DarahNifasBerbau:   req.DarahNifasBerbau,
		PendarahanBerat:    req.PendarahanBerat,
		Keputihan:          req.Keputihan,
	}

	if err := ctrl.usecase.Save(data); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Gagal menyimpan checklist pemantauan ibu nifas",
			"error":   err.Error(),
		})
	}

	if req.HariNifas < 1 || req.HariNifas > 42 {
	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "hari_nifas harus antara 1 sampai 42",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Checklist pemantauan ibu nifas berhasil disimpan",
		"data":    data,
	})
}