package controllers

import (
	"net/http"
	"strconv"
	"time"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type KependudukanController struct {
	usecase usecases.KependudukanUsecase
}

func NewKependudukanController(u usecases.KependudukanUsecase) *KependudukanController {
	return &KependudukanController{usecase: u}
}

type createKependudukanRequest struct {
	KartuKeluargaID    *int64 `json:"kartu_keluarga_id"`
	NIK                string `json:"nik"`
	Dusun              string `json:"dusun"`
	Kecamatan          string `json:"kecamatan"`
	Desa               string `json:"desa"`
	NamaLengkap        string `json:"nama_lengkap"`
	GolonganDarah      string `json:"golongan_darah"`
	JenisKelamin       string `json:"jenis_kelamin"`
	TempatLahir        string `json:"tempat_lahir"`
	TanggalLahir       string `json:"tanggal_lahir"`
	Pekerjaan          string `json:"pekerjaan"`
	PendidikanTerakhir string `json:"pendidikan_terakhir"`
	Agama              string `json:"agama"`
	StatusPerkawinan   string `json:"status_perkawinan"`
	BacaHuruf          string `json:"baca_huruf"`
	KedudukanKeluarga  string `json:"kedudukan_keluarga"`
	AsalPenduduk       string `json:"asal_penduduk"`
	TujuanPindah       string `json:"tujuan_pindah"`
	TempatMeninggal    string `json:"tempat_meninggal"`
	Keterangan         string `json:"keterangan"`
}

func (c *KependudukanController) Create(ctx echo.Context) error {
	claims, _ := ctx.Get("auth_claims").(*models.AuthClaims)
	if claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
	}
	var req createKependudukanRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}

	// Validasi field wajib
	if req.NamaLengkap == "" || req.TanggalLahir == "" {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "nama_lengkap dan tanggal_lahir wajib diisi"})
	}

	// Parse tanggal lahir
	tanggalLahir, err := time.Parse("2006-01-02", req.TanggalLahir)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "format tanggal_lahir harus YYYY-MM-DD"})
	}

	// Convert NIK string to *string (nil jika kosong)
	var nikPtr *string
	if req.NIK != "" {
		nikPtr = &req.NIK
	}

	k := &models.Kependudukan{
		KartuKeluargaID:    req.KartuKeluargaID,
		NIK:                nikPtr,
		Dusun:              req.Dusun,
		Kecamatan:          req.Kecamatan,
		Desa:               req.Desa,
		NamaLengkap:        req.NamaLengkap,
		GolonganDarah:      req.GolonganDarah,
		JenisKelamin:       req.JenisKelamin,
		TempatLahir:        req.TempatLahir,
		TanggalLahir:       tanggalLahir,
		Pekerjaan:          req.Pekerjaan,
		PendidikanTerakhir: req.PendidikanTerakhir,
		Agama:              req.Agama,
		StatusPerkawinan:   req.StatusPerkawinan,
		BacaHuruf:          req.BacaHuruf,
		KedudukanKeluarga:  req.KedudukanKeluarga,
		AsalPenduduk:       req.AsalPenduduk,
		TujuanPindah:       req.TujuanPindah,
		TempatMeninggal:    req.TempatMeninggal,
		Keterangan:         req.Keterangan,
	}

	data, err := c.usecase.Create(k)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusCreated, models.Response{StatusCode: http.StatusCreated, Data: data})
}

func (c *KependudukanController) GetAll(ctx echo.Context) error {
	list, err := c.usecase.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: list})
}

func (c *KependudukanController) GetByID(ctx echo.Context) error {
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

func (c *KependudukanController) Update(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	var req createKependudukanRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	existing, err := c.usecase.GetByID(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: "Data tidak ditemukan"})
	}
	// Update field yang diisi
	if req.KartuKeluargaID != nil {
		existing.KartuKeluargaID = req.KartuKeluargaID
	}
	if req.NIK != "" {
		existing.NIK = &req.NIK
	} else if req.NIK == "" {
		// Jika NIK kosong dalam request, set ke nil
		existing.NIK = nil
	}
	if req.Dusun != "" {
		existing.Dusun = req.Dusun
	}
	if req.Kecamatan != "" {
		existing.Kecamatan = req.Kecamatan
	}
	if req.Desa != "" {
		existing.Desa = req.Desa
	}
	if req.NamaLengkap != "" {
		existing.NamaLengkap = req.NamaLengkap
	}
	if req.GolonganDarah != "" {
		existing.GolonganDarah = req.GolonganDarah
	}
	if req.JenisKelamin != "" {
		existing.JenisKelamin = req.JenisKelamin
	}
	if req.TempatLahir != "" {
		existing.TempatLahir = req.TempatLahir
	}
	if req.TanggalLahir != "" {
		if t, err := time.Parse("2006-01-02", req.TanggalLahir); err == nil {
			existing.TanggalLahir = t
		}
	}
	if req.Pekerjaan != "" {
		existing.Pekerjaan = req.Pekerjaan
	}
	if req.PendidikanTerakhir != "" {
		existing.PendidikanTerakhir = req.PendidikanTerakhir
	}
	if req.Agama != "" {
		existing.Agama = req.Agama
	}
	if req.StatusPerkawinan != "" {
		existing.StatusPerkawinan = req.StatusPerkawinan
	}
	if req.BacaHuruf != "" {
		existing.BacaHuruf = req.BacaHuruf
	}
	if req.KedudukanKeluarga != "" {
		existing.KedudukanKeluarga = req.KedudukanKeluarga
	}
	if req.AsalPenduduk != "" {
		existing.AsalPenduduk = req.AsalPenduduk
	}
	if req.TujuanPindah != "" {
		existing.TujuanPindah = req.TujuanPindah
	}
	if req.TempatMeninggal != "" {
		existing.TempatMeninggal = req.TempatMeninggal
	}
	if req.Keterangan != "" {
		existing.Keterangan = req.Keterangan
	}
	if err := c.usecase.Update(existing); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: existing})
}

func (c *KependudukanController) Delete(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	if err := c.usecase.Delete(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Message: "deleted"})
}

func (c *KependudukanController) GetRekapPerDusun(ctx echo.Context) error {
	kecamatan := ctx.QueryParam("kecamatan")
	desa := ctx.QueryParam("desa")

	data, err := c.usecase.GetRekapPerDusun(kecamatan, desa)
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
