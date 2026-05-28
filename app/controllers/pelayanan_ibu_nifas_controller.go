package controllers

import (
	"net/http"
	"strconv"
	"time"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"
	
	"github.com/labstack/echo/v4"
)

type PelayananIbuNifasController struct {
	usecase usecases.PelayananIbuNifasUsecase
}

func NewPelayananIbuNifasController(u usecases.PelayananIbuNifasUsecase) *PelayananIbuNifasController {
	return &PelayananIbuNifasController{usecase: u}
}

type createPelayananNifasRequest struct {
	KehamilanID                         int32   `json:"kehamilan_id"`
	KunjunganKe                         string  `json:"kunjungan_ke"`
	TanggalPeriksa                      string  `json:"tanggal_periksa"`
	TandaVitalTekananDarah              string  `json:"tanda_vital_tekanan_darah"`
	TandaVitalSuhuTubuh                 float64 `json:"tanda_vital_suhu_tubuh"`
	PelayananInvolusiUteri              string  `json:"pelayanan_involusi_uteri"`
	PelayananCairanPervaginam           string  `json:"pelayanan_cairan_pervaginam"`
	PelayananPeriksaJalanLahir          string  `json:"pelayanan_periksa_jalan_lahir"`
	PelayananPeriksaPayudara            string  `json:"pelayanan_periksa_payudara"`
	PelayananASIEkslusif 				string 	`gorm:"column:pelayanan_asi_eksklusif;type:varchar(50)" json:"pelayanan_asi_eksklusif"`	
	PemberianKapsulVitaminA             bool    `json:"pemberian_kapsul_vitamin_a"`
	PemberianTabletTambahDarahJumlah    int     `json:"pemberian_tablet_tambah_darah_jumlah"`
	PelayananSkriningDepresiNifas       string  `json:"pelayanan_skrining_depresi_nifas"`
	PelayananKontrasepsiPascaPersalinan string  `json:"pelayanan_kontrasepsi_pasca_persalinan"`
	PelayananPenangananRisikoMalaria    string  `json:"pelayanan_penanganan_risiko_malaria"`
	KomplikasiNifas                     string  `json:"komplikasi_nifas"`
	TindakanSaran                       string  `json:"tindakan_saran"`
	NamaPemeriksaParaf                  string  `json:"nama_pemeriksa_paraf"`
	TanggalKembali                      string  `json:"tanggal_kembali"`
}

func (c *PelayananIbuNifasController) Create(ctx echo.Context) error {
	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
	if !ok || claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
	}
	var req createPelayananNifasRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	p := &models.PelayananIbuNifas{
		KehamilanID:                         req.KehamilanID,
		KunjunganKe:                         req.KunjunganKe,
		TandaVitalTekananDarah:              req.TandaVitalTekananDarah,
		TandaVitalSuhuTubuh:                 &req.TandaVitalSuhuTubuh,
		PelayananInvolusiUteri:              req.PelayananInvolusiUteri,
		PelayananCairanPervaginam:           req.PelayananCairanPervaginam,
		PelayananPeriksaJalanLahir:          req.PelayananPeriksaJalanLahir,
		PelayananPeriksaPayudara:            req.PelayananPeriksaPayudara,
		PelayananASIEkslusif:                req.PelayananASIEkslusif,
		PemberianKapsulVitaminA:             req.PemberianKapsulVitaminA,
		PemberianTabletTambahDarahJumlah:    &req.PemberianTabletTambahDarahJumlah,
		PelayananSkriningDepresiNifas:       req.PelayananSkriningDepresiNifas,
		PelayananKontrasepsiPascaPersalinan: req.PelayananKontrasepsiPascaPersalinan,
		PelayananPenangananRisikoMalaria:    req.PelayananPenangananRisikoMalaria,
		KomplikasiNifas:                     req.KomplikasiNifas,
		TindakanSaran:                       req.TindakanSaran,
		NamaPemeriksaParaf:                  req.NamaPemeriksaParaf,
	}
	if req.TanggalPeriksa != "" {
		if t, err := time.Parse("2006-01-02", req.TanggalPeriksa); err == nil {
			p.TanggalPeriksa = &t
		}
	}
	if req.TanggalKembali != "" {
		if t, err := time.Parse("2006-01-02", req.TanggalKembali); err == nil {
			p.TanggalKembali = &t
		}
	}
	if err := c.usecase.Create(p); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusCreated, models.Response{StatusCode: http.StatusCreated, Data: p})
}

func (c *PelayananIbuNifasController) GetByID(ctx echo.Context) error {
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

func (c *PelayananIbuNifasController) GetByKehamilanID(ctx echo.Context) error {
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

func (c *PelayananIbuNifasController) Update(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	var req createPelayananNifasRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	existing, err := c.usecase.GetByID(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: "Data tidak ditemukan"})
	}
	if req.KunjunganKe != "" {
		existing.KunjunganKe = req.KunjunganKe
	}
	if req.TanggalPeriksa != "" {
		if t, err := time.Parse("2006-01-02", req.TanggalPeriksa); err == nil {
			existing.TanggalPeriksa = &t
		}
	}
	if req.TandaVitalTekananDarah != "" {
		existing.TandaVitalTekananDarah = req.TandaVitalTekananDarah
	}
	if req.TandaVitalSuhuTubuh != 0 {
		existing.TandaVitalSuhuTubuh = &req.TandaVitalSuhuTubuh
	}
	if req.PelayananInvolusiUteri != "" {
		existing.PelayananInvolusiUteri = req.PelayananInvolusiUteri
	}
	if req.PelayananCairanPervaginam != "" {
		existing.PelayananCairanPervaginam = req.PelayananCairanPervaginam
	}
	if req.PelayananPeriksaJalanLahir != "" {
		existing.PelayananPeriksaJalanLahir = req.PelayananPeriksaJalanLahir
	}
	if req.PelayananPeriksaPayudara != "" {
		existing.PelayananPeriksaPayudara = req.PelayananPeriksaPayudara
	}
	if req.PelayananASIEkslusif != "" {
		existing.PelayananASIEkslusif = req.PelayananASIEkslusif
	}
	existing.PemberianKapsulVitaminA = req.PemberianKapsulVitaminA
	if req.PemberianTabletTambahDarahJumlah != 0 {
		existing.PemberianTabletTambahDarahJumlah = &req.PemberianTabletTambahDarahJumlah
	}
	if req.PelayananSkriningDepresiNifas != "" {
		existing.PelayananSkriningDepresiNifas = req.PelayananSkriningDepresiNifas
	}
	if req.PelayananKontrasepsiPascaPersalinan != "" {
		existing.PelayananKontrasepsiPascaPersalinan = req.PelayananKontrasepsiPascaPersalinan
	}
	if req.PelayananPenangananRisikoMalaria != "" {
		existing.PelayananPenangananRisikoMalaria = req.PelayananPenangananRisikoMalaria
	}
	if req.KomplikasiNifas != "" {
		existing.KomplikasiNifas = req.KomplikasiNifas
	}
	if req.TindakanSaran != "" {
		existing.TindakanSaran = req.TindakanSaran
	}
	if req.NamaPemeriksaParaf != "" {
		existing.NamaPemeriksaParaf = req.NamaPemeriksaParaf
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

func (c *PelayananIbuNifasController) GetMine(
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

func (c *PelayananIbuNifasController) Delete(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	if err := c.usecase.Delete(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Message: "deleted"})
}
