// package controllers

// import (
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"monitoring-service/app/models"
// 	"monitoring-service/app/usecases"

// 	"github.com/labstack/echo/v4"
// )

// type RiwayatProsesMelahirkanController struct {
// 	usecase usecases.RiwayatProsesMelahirkanUsecase
// }

// func NewRiwayatProsesMelahirkanController(u usecases.RiwayatProsesMelahirkanUsecase) *RiwayatProsesMelahirkanController {
// 	return &RiwayatProsesMelahirkanController{usecase: u}
// }

// type createRiwayatProsesRequest struct {
// 	KehamilanID                   int32  `json:"kehamilan_id"`
// 	GGravida                      int    `json:"g_gravida"`
// 	PPartus                       int    `json:"p_partus"`
// 	AAbortus                      int    `json:"a_abortus"`
// 	HariMelahirkan                string `json:"hari_melahirkan"`
// 	TanggalMelahirkan             string `json:"tanggal_melahirkan"`
// 	PukulMelahirkan               string `json:"pukul_melahirkan"`
// 	CaraMelahirkanSpontan         bool   `json:"cara_melahirkan_spontan"`
// 	CaraMelahirkanSungsang        bool   `json:"cara_melahirkan_sungsang"`
// 	TindakanEkstraksiVakum        bool   `json:"tindakan_ekstraksi_vakum"`
// 	TindakanEkstraksiForsep       bool   `json:"tindakan_ekstraksi_forsep"`
// 	TindakanSC                    bool   `json:"tindakan_sc"`
// 	PenolongDokterSpesialis       bool   `json:"penolong_dokter_spesialis"`
// 	PenolongDokter                bool   `json:"penolong_dokter"`
// 	PenolongBidan                 bool   `json:"penolong_bidan"`
// 	TaksiranMelahirkan            string `json:"taksiran_melahirkan"`
// 	FasyankesTempatMelahirkan     string `json:"fasyankes_tempat_melahirkan"`
// 	RujukanKeterangan             string `json:"rujukan_keterangan"`
// 	InisiasiMenyusuDiniKeterangan string `json:"inisiasi_menyusu_dini_keterangan"`
// 	CapKakiBayiImageURL           []byte `json:"cap_kaki_bayi_image_url"`
// }

// func (c *RiwayatProsesMelahirkanController) Create(ctx echo.Context) error {
// 	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
// 	if !ok || claims == nil {
// 		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
// 	}
// 	var req createRiwayatProsesRequest
// 	if err := ctx.Bind(&req); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
// 	}
// 	r := &models.RiwayatProsesMelahirkan{
// 		KehamilanID:                   req.KehamilanID,
// 		GGravida:                      req.GGravida,
// 		PPartus:                       req.PPartus,
// 		AAbortus:                      req.AAbortus,
// 		HariMelahirkan:                req.HariMelahirkan,
// 		CaraMelahirkanSpontan:         req.CaraMelahirkanSpontan,
// 		CaraMelahirkanSungsang:        req.CaraMelahirkanSungsang,
// 		TindakanEkstraksiVakum:        req.TindakanEkstraksiVakum,
// 		TindakanEkstraksiForsep:       req.TindakanEkstraksiForsep,
// 		TindakanSC:                    req.TindakanSC,
// 		PenolongDokterSpesialis:       req.PenolongDokterSpesialis,
// 		PenolongDokter:                req.PenolongDokter,
// 		PenolongBidan:                 req.PenolongBidan,
// 		TaksiranMelahirkan:            req.TaksiranMelahirkan,
// 		FasyankesTempatMelahirkan:     req.FasyankesTempatMelahirkan,
// 		RujukanKeterangan:             req.RujukanKeterangan,
// 		InisiasiMenyusuDiniKeterangan: req.InisiasiMenyusuDiniKeterangan,
// 		CapKakiBayiImageURL:           req.CapKakiBayiImageURL,
// 	}
// 	if req.TanggalMelahirkan != "" {
// 		if t, err := time.Parse("2006-01-02", req.TanggalMelahirkan); err == nil {
// 			r.TanggalMelahirkan = &t
// 		}
// 	}
// 	if req.PukulMelahirkan != "" {
// 		if t, err := time.Parse("15:04:05", req.PukulMelahirkan); err == nil {
// 			r.PukulMelahirkan = &t
// 		}
// 	}
// 	if err := c.usecase.Create(r); err != nil {
// 		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
// 	}
// 	return ctx.JSON(http.StatusCreated, models.Response{StatusCode: http.StatusCreated, Data: r})
// }

// func (c *RiwayatProsesMelahirkanController) GetByID(ctx echo.Context) error {
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

// func (c *RiwayatProsesMelahirkanController) GetByKehamilanID(ctx echo.Context) error {
// 	kehamilanID, err := strconv.ParseInt(ctx.QueryParam("kehamilan_id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "kehamilan_id required"})
// 	}
// 	list, err := c.usecase.GetByKehamilanID(int32(kehamilanID))
// 	if err != nil {
// 		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: list})
// }

// func (c *RiwayatProsesMelahirkanController) Update(ctx echo.Context) error {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
// 	}
// 	var req createRiwayatProsesRequest
// 	if err := ctx.Bind(&req); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
// 	}
// 	existing, err := c.usecase.GetByID(int32(id))
// 	if err != nil {
// 		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: "Data tidak ditemukan"})
// 	}
// 	if req.HariMelahirkan != "" {
// 		existing.HariMelahirkan = req.HariMelahirkan
// 	}
// 	if req.TanggalMelahirkan != "" {
// 		if t, err := time.Parse("2006-01-02", req.TanggalMelahirkan); err == nil {
// 			existing.TanggalMelahirkan = &t
// 		}
// 	}
// 	if req.PukulMelahirkan != "" {
// 		if t, err := time.Parse("15:04:05", req.PukulMelahirkan); err == nil {
// 			existing.PukulMelahirkan = &t
// 		}
// 	}
// 	if req.GGravida != 0 {
// 		existing.GGravida = req.GGravida
// 	}
// 	if req.PPartus != 0 {
// 		existing.PPartus = req.PPartus
// 	}
// 	if req.AAbortus != 0 {
// 		existing.AAbortus = req.AAbortus
// 	}
// 	existing.CaraMelahirkanSpontan = req.CaraMelahirkanSpontan
// 	existing.CaraMelahirkanSungsang = req.CaraMelahirkanSungsang
// 	existing.TindakanEkstraksiVakum = req.TindakanEkstraksiVakum
// 	existing.TindakanEkstraksiForsep = req.TindakanEkstraksiForsep
// 	existing.TindakanSC = req.TindakanSC
// 	existing.PenolongDokterSpesialis = req.PenolongDokterSpesialis
// 	existing.PenolongDokter = req.PenolongDokter
// 	existing.PenolongBidan = req.PenolongBidan
// 	if req.TaksiranMelahirkan != "" {
// 		existing.TaksiranMelahirkan = req.TaksiranMelahirkan
// 	}
// 	if req.FasyankesTempatMelahirkan != "" {
// 		existing.FasyankesTempatMelahirkan = req.FasyankesTempatMelahirkan
// 	}
// 	if req.RujukanKeterangan != "" {
// 		existing.RujukanKeterangan = req.RujukanKeterangan
// 	}
// 	if req.InisiasiMenyusuDiniKeterangan != "" {
// 		existing.InisiasiMenyusuDiniKeterangan = req.InisiasiMenyusuDiniKeterangan
// 	}
// 	if req.CapKakiBayiImageURL != nil {
// 		existing.CapKakiBayiImageURL = req.CapKakiBayiImageURL
// 	}
// 	if err := c.usecase.Update(existing); err != nil {
// 		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: existing})
// }

// func (c *RiwayatProsesMelahirkanController) Delete(ctx echo.Context) error {
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

type RiwayatProsesMelahirkanController struct {
	usecase usecases.RiwayatProsesMelahirkanUsecase
}

func NewRiwayatProsesMelahirkanController(u usecases.RiwayatProsesMelahirkanUsecase) *RiwayatProsesMelahirkanController {
	return &RiwayatProsesMelahirkanController{usecase: u}
}

type createRiwayatProsesRequest struct {
	KehamilanID                   int32  `json:"kehamilan_id"`
	GGravida                      int    `json:"g_gravida"`
	PPartus                       int    `json:"p_partus"`
	AAbortus                      int    `json:"a_abortus"`
	HariMelahirkan                string `json:"hari_melahirkan"`
	TanggalMelahirkan             string `json:"tanggal_melahirkan"`
	PukulMelahirkan               string `json:"pukul_melahirkan"`
	CaraMelahirkanSpontan         bool   `json:"cara_melahirkan_spontan"`
	CaraMelahirkanSungsang        bool   `json:"cara_melahirkan_sungsang"`
	TindakanEkstraksiVakum        bool   `json:"tindakan_ekstraksi_vakum"`
	TindakanEkstraksiForsep       bool   `json:"tindakan_ekstraksi_forsep"`
	TindakanSC                    bool   `json:"tindakan_sc"`
	PenolongDokterSpesialis       bool   `json:"penolong_dokter_spesialis"`
	PenolongDokter                bool   `json:"penolong_dokter"`
	PenolongBidan                 bool   `json:"penolong_bidan"`
	TaksiranMelahirkan            string `json:"taksiran_melahirkan"`
	FasyankesTempatMelahirkan     string `json:"fasyankes_tempat_melahirkan"`
	RujukanKeterangan             string `json:"rujukan_keterangan"`
	InisiasiMenyusuDiniKeterangan string `json:"inisiasi_menyusu_dini_keterangan"`
	CapKakiBayiImageURL           []byte `json:"cap_kaki_bayi_image_url"`
}

// GetMine — dipakai oleh ibu (ibuk route group).
// Mengambil data riwayat proses melahirkan berdasarkan kehamilan aktif user yang sedang login.
func (c *RiwayatProsesMelahirkanController) GetMine(ctx echo.Context) error {
	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
	if !ok || claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "token tidak valid",
		})
	}

	data, err := c.usecase.GetMineByUserID(claims.UserID)
	if err != nil {
		// Kembalikan 404 dengan data kosong supaya Flutter bisa handle gracefully
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

func (c *RiwayatProsesMelahirkanController) Create(ctx echo.Context) error {
	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
	if !ok || claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
	}
	var req createRiwayatProsesRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	r := &models.RiwayatProsesMelahirkan{
		KehamilanID:                   req.KehamilanID,
		GGravida:                      req.GGravida,
		PPartus:                       req.PPartus,
		AAbortus:                      req.AAbortus,
		HariMelahirkan:                req.HariMelahirkan,
		CaraMelahirkanSpontan:         req.CaraMelahirkanSpontan,
		CaraMelahirkanSungsang:        req.CaraMelahirkanSungsang,
		TindakanEkstraksiVakum:        req.TindakanEkstraksiVakum,
		TindakanEkstraksiForsep:       req.TindakanEkstraksiForsep,
		TindakanSC:                    req.TindakanSC,
		PenolongDokterSpesialis:       req.PenolongDokterSpesialis,
		PenolongDokter:                req.PenolongDokter,
		PenolongBidan:                 req.PenolongBidan,
		TaksiranMelahirkan:            req.TaksiranMelahirkan,
		FasyankesTempatMelahirkan:     req.FasyankesTempatMelahirkan,
		RujukanKeterangan:             req.RujukanKeterangan,
		InisiasiMenyusuDiniKeterangan: req.InisiasiMenyusuDiniKeterangan,
		CapKakiBayiImageURL:           req.CapKakiBayiImageURL,
	}
	if req.TanggalMelahirkan != "" {
		if t, err := time.Parse("2006-01-02", req.TanggalMelahirkan); err == nil {
			r.TanggalMelahirkan = &t
		}
	}
	if req.PukulMelahirkan != "" {
		if t, err := time.Parse("15:04:05", req.PukulMelahirkan); err == nil {
			r.PukulMelahirkan = &t
		}
	}
	if err := c.usecase.Create(r); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusCreated, models.Response{StatusCode: http.StatusCreated, Data: r})
}

func (c *RiwayatProsesMelahirkanController) GetByID(ctx echo.Context) error {
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

func (c *RiwayatProsesMelahirkanController) GetByKehamilanID(ctx echo.Context) error {
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

func (c *RiwayatProsesMelahirkanController) Update(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	var req createRiwayatProsesRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	existing, err := c.usecase.GetByID(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: "Data tidak ditemukan"})
	}
	if req.HariMelahirkan != "" {
		existing.HariMelahirkan = req.HariMelahirkan
	}
	if req.TanggalMelahirkan != "" {
		if t, err := time.Parse("2006-01-02", req.TanggalMelahirkan); err == nil {
			existing.TanggalMelahirkan = &t
		}
	}
	if req.PukulMelahirkan != "" {
		if t, err := time.Parse("15:04:05", req.PukulMelahirkan); err == nil {
			existing.PukulMelahirkan = &t
		}
	}
	if req.GGravida != 0 {
		existing.GGravida = req.GGravida
	}
	if req.PPartus != 0 {
		existing.PPartus = req.PPartus
	}
	if req.AAbortus != 0 {
		existing.AAbortus = req.AAbortus
	}
	existing.CaraMelahirkanSpontan = req.CaraMelahirkanSpontan
	existing.CaraMelahirkanSungsang = req.CaraMelahirkanSungsang
	existing.TindakanEkstraksiVakum = req.TindakanEkstraksiVakum
	existing.TindakanEkstraksiForsep = req.TindakanEkstraksiForsep
	existing.TindakanSC = req.TindakanSC
	existing.PenolongDokterSpesialis = req.PenolongDokterSpesialis
	existing.PenolongDokter = req.PenolongDokter
	existing.PenolongBidan = req.PenolongBidan
	if req.TaksiranMelahirkan != "" {
		existing.TaksiranMelahirkan = req.TaksiranMelahirkan
	}
	if req.FasyankesTempatMelahirkan != "" {
		existing.FasyankesTempatMelahirkan = req.FasyankesTempatMelahirkan
	}
	if req.RujukanKeterangan != "" {
		existing.RujukanKeterangan = req.RujukanKeterangan
	}
	if req.InisiasiMenyusuDiniKeterangan != "" {
		existing.InisiasiMenyusuDiniKeterangan = req.InisiasiMenyusuDiniKeterangan
	}
	if req.CapKakiBayiImageURL != nil {
		existing.CapKakiBayiImageURL = req.CapKakiBayiImageURL
	}
	if err := c.usecase.Update(existing); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: existing})
}

func (c *RiwayatProsesMelahirkanController) Delete(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	if err := c.usecase.Delete(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Message: "deleted"})
}