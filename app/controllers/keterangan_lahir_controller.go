// package controllers

// import (
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"monitoring-service/app/models"
// 	"monitoring-service/app/usecases"

// 	"github.com/labstack/echo/v4"
// )

// type KeteranganLahirController struct {
// 	usecase usecases.KeteranganLahirUsecase
// }

// func NewKeteranganLahirController(u usecases.KeteranganLahirUsecase) *KeteranganLahirController {
// 	return &KeteranganLahirController{usecase: u}
// }

// type createKeteranganLahirRequest struct {
// 	IDIbuRelasi            int32  `json:"id_ibu_relasi"`
// 	NomorSurat             string `json:"nomor_surat"`
// 	HariLahir              string `json:"hari_lahir"`
// 	TanggalLahir           string `json:"tanggal_lahir"`
// 	PukulLahir             string `json:"pukul_lahir"`
// 	JenisKelamin           string `json:"jenis_kelamin"`
// 	JenisKelahiran         string `json:"jenis_kelahiran"`
// 	AnakKe                 int    `json:"anak_ke"`
// 	UsiaGestasiMinggu      int    `json:"usia_gestasi_minggu"`
// 	BeratLahirGram         int    `json:"berat_lahir_gram"`
// 	PanjangBadanCm         int    `json:"panjang_badan_cm"`
// 	LingkarKepalaCm        int    `json:"lingkar_kepala_cm"`
// 	LokasiPersalinan       string `json:"lokasi_persalinan"`
// 	AlamatLokasiPersalinan string `json:"alamat_lokasi_persalinan"`
// 	NamaBayiDiberiNama     string `json:"nama_bayi_diberi_nama"`
// 	NamaIbu                string `json:"nama_ibu"`
// 	UmurIbu                int    `json:"umur_ibu"`
// 	NIKIbu                 string `json:"nik_ibu"`
// 	NamaAyah               string `json:"nama_ayah"`
// 	NIKAyah                string `json:"nik_ayah"`
// 	PekerjaanOrangTua      string `json:"pekerjaan_orang_tua"`
// 	AlamatOrangTua         string `json:"alamat_orang_tua"`
// 	RWRTOrangTua           string `json:"rw_rt_orang_tua"`
// 	KecamatanOrangTua      string `json:"kecamatan_orang_tua"`
// 	KabKotaOrangTua        string `json:"kab_kota_orang_tua"`
// 	TanggalSurat           string `json:"tanggal_surat"`
// 	NamaSaksi1             string `json:"nama_saksi_1"`
// 	NamaSaksi2             string `json:"nama_saksi_2"`
// 	NamaPenolongKelahiran  string `json:"nama_penolong_kelahiran"`
// }

// type updateKeteranganLahirRequest struct {
// 	NomorSurat             string `json:"nomor_surat"`
// 	HariLahir              string `json:"hari_lahir"`
// 	TanggalLahir           string `json:"tanggal_lahir"`
// 	PukulLahir             string `json:"pukul_lahir"`
// 	JenisKelamin           string `json:"jenis_kelamin"`
// 	JenisKelahiran         string `json:"jenis_kelahiran"`
// 	AnakKe                 int    `json:"anak_ke"`
// 	UsiaGestasiMinggu      int    `json:"usia_gestasi_minggu"`
// 	BeratLahirGram         int    `json:"berat_lahir_gram"`
// 	PanjangBadanCm         int    `json:"panjang_badan_cm"`
// 	LingkarKepalaCm        int    `json:"lingkar_kepala_cm"`
// 	LokasiPersalinan       string `json:"lokasi_persalinan"`
// 	AlamatLokasiPersalinan string `json:"alamat_lokasi_persalinan"`
// 	NamaBayiDiberiNama     string `json:"nama_bayi_diberi_nama"`
// 	NamaIbu                string `json:"nama_ibu"`
// 	UmurIbu                int    `json:"umur_ibu"`
// 	NIKIbu                 string `json:"nik_ibu"`
// 	NamaAyah               string `json:"nama_ayah"`
// 	NIKAyah                string `json:"nik_ayah"`
// 	PekerjaanOrangTua      string `json:"pekerjaan_orang_tua"`
// 	AlamatOrangTua         string `json:"alamat_orang_tua"`
// 	RWRTOrangTua           string `json:"rw_rt_orang_tua"`
// 	KecamatanOrangTua      string `json:"kecamatan_orang_tua"`
// 	KabKotaOrangTua        string `json:"kab_kota_orang_tua"`
// 	TanggalSurat           string `json:"tanggal_surat"`
// 	NamaSaksi1             string `json:"nama_saksi_1"`
// 	NamaSaksi2             string `json:"nama_saksi_2"`
// 	NamaPenolongKelahiran  string `json:"nama_penolong_kelahiran"`
// }

// func (c *KeteranganLahirController) Create(ctx echo.Context) error {
// 	claims, _ := ctx.Get("auth_claims").(*models.AuthClaims)
// 	if claims == nil {
// 		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
// 	}
// 	var req createKeteranganLahirRequest
// 	if err := ctx.Bind(&req); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
// 	}
// 	kl := &models.KeteranganLahir{
// 		IDIbuRelasi:            req.IDIbuRelasi,
// 		NomorSurat:             req.NomorSurat,
// 		HariLahir:              req.HariLahir,
// 		JenisKelamin:           req.JenisKelamin,
// 		JenisKelahiran:         req.JenisKelahiran,
// 		AnakKe:                 req.AnakKe,
// 		UsiaGestasiMinggu:      req.UsiaGestasiMinggu,
// 		BeratLahirGram:         req.BeratLahirGram,
// 		PanjangBadanCm:         req.PanjangBadanCm,
// 		LingkarKepalaCm:        req.LingkarKepalaCm,
// 		LokasiPersalinan:       req.LokasiPersalinan,
// 		AlamatLokasiPersalinan: req.AlamatLokasiPersalinan,
// 		NamaBayiDiberiNama:     req.NamaBayiDiberiNama,
// 		NamaIbu:                req.NamaIbu,
// 		UmurIbu:                req.UmurIbu,
// 		NIKIbu:                 req.NIKIbu,
// 		NamaAyah:               req.NamaAyah,
// 		NIKAyah:                req.NIKAyah,
// 		PekerjaanOrangTua:      req.PekerjaanOrangTua,
// 		AlamatOrangTua:         req.AlamatOrangTua,
// 		RWRTOrangTua:           req.RWRTOrangTua,
// 		KecamatanOrangTua:      req.KecamatanOrangTua,
// 		KabKotaOrangTua:        req.KabKotaOrangTua,
// 		NamaSaksi1:             req.NamaSaksi1,
// 		NamaSaksi2:             req.NamaSaksi2,
// 		NamaPenolongKelahiran:  req.NamaPenolongKelahiran,
// 	}
// 	if req.TanggalLahir != "" {
// 		if t, err := time.Parse("2006-01-02", req.TanggalLahir); err == nil {
// 			kl.TanggalLahir = &t
// 		}
// 	}
// 	if req.PukulLahir != "" {
// 		if t, err := time.Parse("15:04:05", req.PukulLahir); err == nil {
// 			kl.PukulLahir = &t
// 		}
// 	}
// 	if req.TanggalSurat != "" {
// 		if t, err := time.Parse("2006-01-02", req.TanggalSurat); err == nil {
// 			kl.TanggalSurat = &t
// 		}
// 	}
// 	if err := c.usecase.Create(kl); err != nil {
// 		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
// 	}
// 	return ctx.JSON(http.StatusCreated, models.Response{StatusCode: http.StatusCreated, Data: kl})
// }

// func (c *KeteranganLahirController) GetByID(ctx echo.Context) error {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
// 	}
// 	data, err := c.usecase.GetByID(int32(id))
// 	if err != nil {
// 		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: data})
// }

// func (c *KeteranganLahirController) GetByIbuID(ctx echo.Context) error {
// 	ibuID, err := strconv.ParseInt(ctx.QueryParam("ibu_id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "ibu_id required"})
// 	}
// 	list, err := c.usecase.GetByIbuID(int32(ibuID))
// 	if err != nil {
// 		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: list})
// }

// func (c *KeteranganLahirController) Update(ctx echo.Context) error {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
// 	}
// 	var req updateKeteranganLahirRequest
// 	if err := ctx.Bind(&req); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
// 	}
// 	existing, err := c.usecase.GetByID(int32(id))
// 	if err != nil {
// 		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: "Data tidak ditemukan"})
// 	}
// 	// Update field yang diisi
// 	if req.NomorSurat != "" {
// 		existing.NomorSurat = req.NomorSurat
// 	}
// 	if req.HariLahir != "" {
// 		existing.HariLahir = req.HariLahir
// 	}
// 	if req.TanggalLahir != "" {
// 		if t, err := time.Parse("2006-01-02", req.TanggalLahir); err == nil {
// 			existing.TanggalLahir = &t
// 		}
// 	}
// 	if req.PukulLahir != "" {
// 		if t, err := time.Parse("15:04:05", req.PukulLahir); err == nil {
// 			existing.PukulLahir = &t
// 		}
// 	}
// 	if req.JenisKelamin != "" {
// 		existing.JenisKelamin = req.JenisKelamin
// 	}
// 	if req.JenisKelahiran != "" {
// 		existing.JenisKelahiran = req.JenisKelahiran
// 	}
// 	if req.AnakKe != 0 {
// 		existing.AnakKe = req.AnakKe
// 	}
// 	if req.UsiaGestasiMinggu != 0 {
// 		existing.UsiaGestasiMinggu = req.UsiaGestasiMinggu
// 	}
// 	if req.BeratLahirGram != 0 {
// 		existing.BeratLahirGram = req.BeratLahirGram
// 	}
// 	if req.PanjangBadanCm != 0 {
// 		existing.PanjangBadanCm = req.PanjangBadanCm
// 	}
// 	if req.LingkarKepalaCm != 0 {
// 		existing.LingkarKepalaCm = req.LingkarKepalaCm
// 	}
// 	if req.LokasiPersalinan != "" {
// 		existing.LokasiPersalinan = req.LokasiPersalinan
// 	}
// 	if req.AlamatLokasiPersalinan != "" {
// 		existing.AlamatLokasiPersalinan = req.AlamatLokasiPersalinan
// 	}
// 	if req.NamaBayiDiberiNama != "" {
// 		existing.NamaBayiDiberiNama = req.NamaBayiDiberiNama
// 	}
// 	if req.NamaIbu != "" {
// 		existing.NamaIbu = req.NamaIbu
// 	}
// 	if req.UmurIbu != 0 {
// 		existing.UmurIbu = req.UmurIbu
// 	}
// 	if req.NIKIbu != "" {
// 		existing.NIKIbu = req.NIKIbu
// 	}
// 	if req.NamaAyah != "" {
// 		existing.NamaAyah = req.NamaAyah
// 	}
// 	if req.NIKAyah != "" {
// 		existing.NIKAyah = req.NIKAyah
// 	}
// 	if req.PekerjaanOrangTua != "" {
// 		existing.PekerjaanOrangTua = req.PekerjaanOrangTua
// 	}
// 	if req.AlamatOrangTua != "" {
// 		existing.AlamatOrangTua = req.AlamatOrangTua
// 	}
// 	if req.RWRTOrangTua != "" {
// 		existing.RWRTOrangTua = req.RWRTOrangTua
// 	}
// 	if req.KecamatanOrangTua != "" {
// 		existing.KecamatanOrangTua = req.KecamatanOrangTua
// 	}
// 	if req.KabKotaOrangTua != "" {
// 		existing.KabKotaOrangTua = req.KabKotaOrangTua
// 	}
// 	if req.TanggalSurat != "" {
// 		if t, err := time.Parse("2006-01-02", req.TanggalSurat); err == nil {
// 			existing.TanggalSurat = &t
// 		}
// 	}
// 	if req.NamaSaksi1 != "" {
// 		existing.NamaSaksi1 = req.NamaSaksi1
// 	}
// 	if req.NamaSaksi2 != "" {
// 		existing.NamaSaksi2 = req.NamaSaksi2
// 	}
// 	if req.NamaPenolongKelahiran != "" {
// 		existing.NamaPenolongKelahiran = req.NamaPenolongKelahiran
// 	}
// 	if err := c.usecase.Update(existing); err != nil {
// 		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: existing})
// }

// func (c *KeteranganLahirController) Delete(ctx echo.Context) error {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
// 	}
// 	if err := c.usecase.Delete(int32(id)); err != nil {
// 		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Message: "deleted"})
// }


package controllers

import (
	"net/http"
	"strconv"
	"time"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type KeteranganLahirController struct {
	usecase usecases.KeteranganLahirUsecase
}

func NewKeteranganLahirController(u usecases.KeteranganLahirUsecase) *KeteranganLahirController {
	return &KeteranganLahirController{usecase: u}
}

type createKeteranganLahirRequest struct {
	IDIbuRelasi            int32  `json:"id_ibu_relasi"`
	NomorSurat             string `json:"nomor_surat"`
	HariLahir              string `json:"hari_lahir"`
	TanggalLahir           string `json:"tanggal_lahir"`
	PukulLahir             string `json:"pukul_lahir"`
	JenisKelamin           string `json:"jenis_kelamin"`
	JenisKelahiran         string `json:"jenis_kelahiran"`
	AnakKe                 int    `json:"anak_ke"`
	UsiaGestasiMinggu      int    `json:"usia_gestasi_minggu"`
	BeratLahirGram         int    `json:"berat_lahir_gram"`
	PanjangBadanCm         int    `json:"panjang_badan_cm"`
	LingkarKepalaCm        int    `json:"lingkar_kepala_cm"`
	LokasiPersalinan       string `json:"lokasi_persalinan"`
	AlamatLokasiPersalinan string `json:"alamat_lokasi_persalinan"`
	NamaBayiDiberiNama     string `json:"nama_bayi_diberi_nama"`
	NamaIbu                string `json:"nama_ibu"`
	UmurIbu                int    `json:"umur_ibu"`
	NIKIbu                 string `json:"nik_ibu"`
	NamaAyah               string `json:"nama_ayah"`
	NIKAyah                string `json:"nik_ayah"`
	PekerjaanOrangTua      string `json:"pekerjaan_orang_tua"`
	AlamatOrangTua         string `json:"alamat_orang_tua"`
	RWRTOrangTua           string `json:"rw_rt_orang_tua"`
	KecamatanOrangTua      string `json:"kecamatan_orang_tua"`
	KabKotaOrangTua        string `json:"kab_kota_orang_tua"`
	TanggalSurat           string `json:"tanggal_surat"`
	NamaSaksi1             string `json:"nama_saksi_1"`
	NamaSaksi2             string `json:"nama_saksi_2"`
	NamaPenolongKelahiran  string `json:"nama_penolong_kelahiran"`
}

type updateKeteranganLahirRequest struct {
	NomorSurat             string `json:"nomor_surat"`
	HariLahir              string `json:"hari_lahir"`
	TanggalLahir           string `json:"tanggal_lahir"`
	PukulLahir             string `json:"pukul_lahir"`
	JenisKelamin           string `json:"jenis_kelamin"`
	JenisKelahiran         string `json:"jenis_kelahiran"`
	AnakKe                 int    `json:"anak_ke"`
	UsiaGestasiMinggu      int    `json:"usia_gestasi_minggu"`
	BeratLahirGram         int    `json:"berat_lahir_gram"`
	PanjangBadanCm         int    `json:"panjang_badan_cm"`
	LingkarKepalaCm        int    `json:"lingkar_kepala_cm"`
	LokasiPersalinan       string `json:"lokasi_persalinan"`
	AlamatLokasiPersalinan string `json:"alamat_lokasi_persalinan"`
	NamaBayiDiberiNama     string `json:"nama_bayi_diberi_nama"`
	NamaIbu                string `json:"nama_ibu"`
	UmurIbu                int    `json:"umur_ibu"`
	NIKIbu                 string `json:"nik_ibu"`
	NamaAyah               string `json:"nama_ayah"`
	NIKAyah                string `json:"nik_ayah"`
	PekerjaanOrangTua      string `json:"pekerjaan_orang_tua"`
	AlamatOrangTua         string `json:"alamat_orang_tua"`
	RWRTOrangTua           string `json:"rw_rt_orang_tua"`
	KecamatanOrangTua      string `json:"kecamatan_orang_tua"`
	KabKotaOrangTua        string `json:"kab_kota_orang_tua"`
	TanggalSurat           string `json:"tanggal_surat"`
	NamaSaksi1             string `json:"nama_saksi_1"`
	NamaSaksi2             string `json:"nama_saksi_2"`
	NamaPenolongKelahiran  string `json:"nama_penolong_kelahiran"`
}

func (c *KeteranganLahirController) Create(ctx echo.Context) error {
	claims, _ := ctx.Get("auth_claims").(*models.AuthClaims)
	if claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
	}
	var req createKeteranganLahirRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	kl := &models.KeteranganLahir{
		IDIbuRelasi:            req.IDIbuRelasi,
		NomorSurat:             req.NomorSurat,
		HariLahir:              req.HariLahir,
		JenisKelamin:           req.JenisKelamin,
		JenisKelahiran:         req.JenisKelahiran,
		AnakKe:                 req.AnakKe,
		UsiaGestasiMinggu:      req.UsiaGestasiMinggu,
		BeratLahirGram:         req.BeratLahirGram,
		PanjangBadanCm:         req.PanjangBadanCm,
		LingkarKepalaCm:        req.LingkarKepalaCm,
		LokasiPersalinan:       req.LokasiPersalinan,
		AlamatLokasiPersalinan: req.AlamatLokasiPersalinan,
		NamaBayiDiberiNama:     req.NamaBayiDiberiNama,
		NamaIbu:                req.NamaIbu,
		UmurIbu:                req.UmurIbu,
		NIKIbu:                 req.NIKIbu,
		NamaAyah:               req.NamaAyah,
		NIKAyah:                req.NIKAyah,
		PekerjaanOrangTua:      req.PekerjaanOrangTua,
		AlamatOrangTua:         req.AlamatOrangTua,
		RWRTOrangTua:           req.RWRTOrangTua,
		KecamatanOrangTua:      req.KecamatanOrangTua,
		KabKotaOrangTua:        req.KabKotaOrangTua,
		NamaSaksi1:             req.NamaSaksi1,
		NamaSaksi2:             req.NamaSaksi2,
		NamaPenolongKelahiran:  req.NamaPenolongKelahiran,
	}
	if req.TanggalLahir != "" {
		if t, err := time.Parse("2006-01-02", req.TanggalLahir); err == nil {
			kl.TanggalLahir = &t
		}
	}
	if req.PukulLahir != "" {
		if t, err := time.Parse("15:04:05", req.PukulLahir); err == nil {
			kl.PukulLahir = &t
		}
	}
	if req.TanggalSurat != "" {
		if t, err := time.Parse("2006-01-02", req.TanggalSurat); err == nil {
			kl.TanggalSurat = &t
		}
	}
	if err := c.usecase.Create(kl); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusCreated, models.Response{StatusCode: http.StatusCreated, Data: kl})
}

func (c *KeteranganLahirController) GetByID(ctx echo.Context) error {
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

func (c *KeteranganLahirController) GetByIbuID(ctx echo.Context) error {
	ibuID, err := strconv.ParseInt(ctx.QueryParam("ibu_id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "ibu_id required"})
	}
	list, err := c.usecase.GetByIbuID(int32(ibuID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: list})
}

// GetMine mengembalikan keterangan lahir milik ibu yang sedang login.
// Dipakai oleh endpoint GET /modul-ibu/keterangan-lahir/me
func (c *KeteranganLahirController) GetMine(ctx echo.Context) error {
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
			Message:    "data keterangan lahir belum tersedia",
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Data:       data,
	})
}

func (c *KeteranganLahirController) Update(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	var req updateKeteranganLahirRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	existing, err := c.usecase.GetByID(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: "Data tidak ditemukan"})
	}
	// Update field yang diisi
	if req.NomorSurat != "" {
		existing.NomorSurat = req.NomorSurat
	}
	if req.HariLahir != "" {
		existing.HariLahir = req.HariLahir
	}
	if req.TanggalLahir != "" {
		if t, err := time.Parse("2006-01-02", req.TanggalLahir); err == nil {
			existing.TanggalLahir = &t
		}
	}
	if req.PukulLahir != "" {
		if t, err := time.Parse("15:04:05", req.PukulLahir); err == nil {
			existing.PukulLahir = &t
		}
	}
	if req.JenisKelamin != "" {
		existing.JenisKelamin = req.JenisKelamin
	}
	if req.JenisKelahiran != "" {
		existing.JenisKelahiran = req.JenisKelahiran
	}
	if req.AnakKe != 0 {
		existing.AnakKe = req.AnakKe
	}
	if req.UsiaGestasiMinggu != 0 {
		existing.UsiaGestasiMinggu = req.UsiaGestasiMinggu
	}
	if req.BeratLahirGram != 0 {
		existing.BeratLahirGram = req.BeratLahirGram
	}
	if req.PanjangBadanCm != 0 {
		existing.PanjangBadanCm = req.PanjangBadanCm
	}
	if req.LingkarKepalaCm != 0 {
		existing.LingkarKepalaCm = req.LingkarKepalaCm
	}
	if req.LokasiPersalinan != "" {
		existing.LokasiPersalinan = req.LokasiPersalinan
	}
	if req.AlamatLokasiPersalinan != "" {
		existing.AlamatLokasiPersalinan = req.AlamatLokasiPersalinan
	}
	if req.NamaBayiDiberiNama != "" {
		existing.NamaBayiDiberiNama = req.NamaBayiDiberiNama
	}
	if req.NamaIbu != "" {
		existing.NamaIbu = req.NamaIbu
	}
	if req.UmurIbu != 0 {
		existing.UmurIbu = req.UmurIbu
	}
	if req.NIKIbu != "" {
		existing.NIKIbu = req.NIKIbu
	}
	if req.NamaAyah != "" {
		existing.NamaAyah = req.NamaAyah
	}
	if req.NIKAyah != "" {
		existing.NIKAyah = req.NIKAyah
	}
	if req.PekerjaanOrangTua != "" {
		existing.PekerjaanOrangTua = req.PekerjaanOrangTua
	}
	if req.AlamatOrangTua != "" {
		existing.AlamatOrangTua = req.AlamatOrangTua
	}
	if req.RWRTOrangTua != "" {
		existing.RWRTOrangTua = req.RWRTOrangTua
	}
	if req.KecamatanOrangTua != "" {
		existing.KecamatanOrangTua = req.KecamatanOrangTua
	}
	if req.KabKotaOrangTua != "" {
		existing.KabKotaOrangTua = req.KabKotaOrangTua
	}
	if req.TanggalSurat != "" {
		if t, err := time.Parse("2006-01-02", req.TanggalSurat); err == nil {
			existing.TanggalSurat = &t
		}
	}
	if req.NamaSaksi1 != "" {
		existing.NamaSaksi1 = req.NamaSaksi1
	}
	if req.NamaSaksi2 != "" {
		existing.NamaSaksi2 = req.NamaSaksi2
	}
	if req.NamaPenolongKelahiran != "" {
		existing.NamaPenolongKelahiran = req.NamaPenolongKelahiran
	}
	if err := c.usecase.Update(existing); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: existing})
}

func (c *KeteranganLahirController) Delete(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	if err := c.usecase.Delete(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Message: "deleted"})
}