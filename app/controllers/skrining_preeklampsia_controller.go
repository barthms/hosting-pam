package controllers

import (
	"net/http"
	"strconv"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type SkriningPreeklampsiaController struct {
	usecase usecases.SkriningPreeklampsiaUsecase
}

func NewSkriningPreeklampsiaController(u usecases.SkriningPreeklampsiaUsecase) *SkriningPreeklampsiaController {
	return &SkriningPreeklampsiaController{usecase: u}
}

type createSkriningPERequest struct {
	KehamilanID                                  int32  `json:"kehamilan_id"`
	AnamnesisMultiparaPasanganBaruSedang         bool   `json:"anamnesis_multipara_pasangan_baru_sedang"`
	AnamnesisTeknologiReproduksiBerbantuSedang   bool   `json:"anamnesis_teknologi_reproduksi_berbantu_sedang"`
	AnamnesisUmurDiatas35TahunSedang             bool   `json:"anamnesis_umur_diatas_35_tahun_sedang"`
	AnamnesisNuliparaSedang                      bool   `json:"anamnesis_nulipara_sedang"`
	AnamnesisJarakKehamilanDiatas10TahunSedang   bool   `json:"anamnesis_jarak_kehamilan_diatas_10_tahun_sedang"`
	AnamnesisRiwayatPreeklampsiaKeluargaSedang   bool   `json:"anamnesis_riwayat_preeklampsia_keluarga_sedang"`
	AnamnesisObesitasIMTDiatas30Sedang           bool   `json:"anamnesis_obesitas_imt_diatas_30_sedang"`
	AnamnesisRiwayatPreeklampsiaSebelumnyaTinggi bool   `json:"anamnesis_riwayat_preeklampsia_sebelumnya_tinggi"`
	AnamnesisKehamilanMultipelTinggi             bool   `json:"anamnesis_kehamilan_multipel_tinggi"`
	AnamnesisDiabetesDalamKehamilanTinggi        bool   `json:"anamnesis_diabetes_dalam_kehamilan_tinggi"`
	AnamnesisHipertensiKronikTinggi              bool   `json:"anamnesis_hipertensi_kronik_tinggi"`
	AnamnesisPenyakitGinjalTinggi                bool   `json:"anamnesis_penyakit_ginjal_tinggi"`
	AnamnesisPenyakitAutoimunSLETinggi           bool   `json:"anamnesis_penyakit_autoimun_sle_tinggi"`
	AnamnesisAntiPhospholipidSyndromeTinggi      bool   `json:"anamnesis_anti_phospholipid_syndrome_tinggi"`
	FisikMAPDiatas90mmHg                         bool   `json:"fisik_map_diatas_90_mmhg"`
	FisikProteinuriaUrinCelup                    bool   `json:"fisik_proteinuria_urin_celup"`
	KesimpulanSkriningPreeklampsia               string `json:"kesimpulan_skrining_preeklampsia"`
}

func (c *SkriningPreeklampsiaController) Create(ctx echo.Context) error {
	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
	if !ok || claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
	}
	var req createSkriningPERequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	s := &models.SkriningPreeklampsia{
		KehamilanID:                                  req.KehamilanID,
		AnamnesisMultiparaPasanganBaruSedang:         req.AnamnesisMultiparaPasanganBaruSedang,
		AnamnesisTeknologiReproduksiBerbantuSedang:   req.AnamnesisTeknologiReproduksiBerbantuSedang,
		AnamnesisUmurDiatas35TahunSedang:             req.AnamnesisUmurDiatas35TahunSedang,
		AnamnesisNuliparaSedang:                      req.AnamnesisNuliparaSedang,
		AnamnesisJarakKehamilanDiatas10TahunSedang:   req.AnamnesisJarakKehamilanDiatas10TahunSedang,
		AnamnesisRiwayatPreeklampsiaKeluargaSedang:   req.AnamnesisRiwayatPreeklampsiaKeluargaSedang,
		AnamnesisObesitasIMTDiatas30Sedang:           req.AnamnesisObesitasIMTDiatas30Sedang,
		AnamnesisRiwayatPreeklampsiaSebelumnyaTinggi: req.AnamnesisRiwayatPreeklampsiaSebelumnyaTinggi,
		AnamnesisKehamilanMultipelTinggi:             req.AnamnesisKehamilanMultipelTinggi,
		AnamnesisDiabetesDalamKehamilanTinggi:        req.AnamnesisDiabetesDalamKehamilanTinggi,
		AnamnesisHipertensiKronikTinggi:              req.AnamnesisHipertensiKronikTinggi,
		AnamnesisPenyakitGinjalTinggi:                req.AnamnesisPenyakitGinjalTinggi,
		AnamnesisPenyakitAutoimunSLETinggi:           req.AnamnesisPenyakitAutoimunSLETinggi,
		AnamnesisAntiPhospholipidSyndromeTinggi:      req.AnamnesisAntiPhospholipidSyndromeTinggi,
		FisikMAPDiatas90mmHg:                         req.FisikMAPDiatas90mmHg,
		FisikProteinuriaUrinCelup:                    req.FisikProteinuriaUrinCelup,
		KesimpulanSkriningPreeklampsia:               req.KesimpulanSkriningPreeklampsia,
	}
	if err := c.usecase.Create(s); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusCreated, models.Response{StatusCode: http.StatusCreated, Data: s})
}

func (c *SkriningPreeklampsiaController) GetByID(ctx echo.Context) error {
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

func (c *SkriningPreeklampsiaController) GetByKehamilanID(ctx echo.Context) error {
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

func (c *SkriningPreeklampsiaController) Update(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	var req createSkriningPERequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	existing, err := c.usecase.GetByID(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: "Data tidak ditemukan"})
	}
	// Update boolean fields (langsung assign)
	existing.AnamnesisMultiparaPasanganBaruSedang = req.AnamnesisMultiparaPasanganBaruSedang
	existing.AnamnesisTeknologiReproduksiBerbantuSedang = req.AnamnesisTeknologiReproduksiBerbantuSedang
	existing.AnamnesisUmurDiatas35TahunSedang = req.AnamnesisUmurDiatas35TahunSedang
	existing.AnamnesisNuliparaSedang = req.AnamnesisNuliparaSedang
	existing.AnamnesisJarakKehamilanDiatas10TahunSedang = req.AnamnesisJarakKehamilanDiatas10TahunSedang
	existing.AnamnesisRiwayatPreeklampsiaKeluargaSedang = req.AnamnesisRiwayatPreeklampsiaKeluargaSedang
	existing.AnamnesisObesitasIMTDiatas30Sedang = req.AnamnesisObesitasIMTDiatas30Sedang
	existing.AnamnesisRiwayatPreeklampsiaSebelumnyaTinggi = req.AnamnesisRiwayatPreeklampsiaSebelumnyaTinggi
	existing.AnamnesisKehamilanMultipelTinggi = req.AnamnesisKehamilanMultipelTinggi
	existing.AnamnesisDiabetesDalamKehamilanTinggi = req.AnamnesisDiabetesDalamKehamilanTinggi
	existing.AnamnesisHipertensiKronikTinggi = req.AnamnesisHipertensiKronikTinggi
	existing.AnamnesisPenyakitGinjalTinggi = req.AnamnesisPenyakitGinjalTinggi
	existing.AnamnesisPenyakitAutoimunSLETinggi = req.AnamnesisPenyakitAutoimunSLETinggi
	existing.AnamnesisAntiPhospholipidSyndromeTinggi = req.AnamnesisAntiPhospholipidSyndromeTinggi
	existing.FisikMAPDiatas90mmHg = req.FisikMAPDiatas90mmHg
	existing.FisikProteinuriaUrinCelup = req.FisikProteinuriaUrinCelup
	if req.KesimpulanSkriningPreeklampsia != "" {
		existing.KesimpulanSkriningPreeklampsia = req.KesimpulanSkriningPreeklampsia
	}
	if err := c.usecase.Update(existing); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: existing})
}

func (c *SkriningPreeklampsiaController) Delete(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	if err := c.usecase.Delete(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Message: "deleted"})
}


// MODUL IBU
func (c *SkriningPreeklampsiaController) GetByIDForOrangtua(ctx echo.Context) error {
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

func (c *SkriningPreeklampsiaController) GetMine(ctx echo.Context) error {
	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
	if !ok || claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "token tidak valid",
		})
	}

	data, err := c.usecase.GetMine(claims.UserID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{
			StatusCode: http.StatusNotFound,
			Message:    "data skrining preeklampsia tidak ditemukan",
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Data:       data,
	})
}