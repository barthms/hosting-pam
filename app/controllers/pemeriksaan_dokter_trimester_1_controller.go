// // package controllers

// // import (
// // 	"net/http"
// // 	"strconv"
// // 	"time"

// // 	"monitoring-service/app/models"
// // 	"monitoring-service/app/usecases"

// // 	"github.com/labstack/echo/v4"
// // )

// // type PemeriksaanDokterTrimester1Controller struct {
// // 	usecase usecases.PemeriksaanDokterTrimester1Usecase
// // }

// // func NewPemeriksaanDokterTrimester1Controller(u usecases.PemeriksaanDokterTrimester1Usecase) *PemeriksaanDokterTrimester1Controller {
// // 	return &PemeriksaanDokterTrimester1Controller{usecase: u}
// // }

// // type createPemeriksaanDokterT1Request struct {
// // 	KehamilanID                 int32   `json:"kehamilan_id"`
// // 	NamaDokter                  string  `json:"nama_dokter"`
// // 	TanggalPeriksa              string  `json:"tanggal_periksa"`
// // 	KonsepAnamnesaPemeriksaan   string  `json:"konsep_anamnesa_pemeriksaan"`
// // 	FisikKonjungtiva            string  `json:"fisik_konjungtiva"`
// // 	FisikSklera                 string  `json:"fisik_sklera"`
// // 	FisikKulit                  string  `json:"fisik_kulit"`
// // 	FisikLeher                  string  `json:"fisik_leher"`
// // 	FisikGigiMulut              string  `json:"fisik_gigi_mulut"`
// // 	FisikTHT                    string  `json:"fisik_tht"`
// // 	FisikDadaJantung            string  `json:"fisik_dada_jantung"`
// // 	FisikDadaParu               string  `json:"fisik_dada_paru"`
// // 	FisikPerut                  string  `json:"fisik_perut"`
// // 	FisikTungkai                string  `json:"fisik_tungkai"`
// // 	HPHT                        string  `json:"hpht"`
// // 	KeteraturanHaid             string  `json:"keteraturan_haid"`
// // 	UmurHamilHPHTMinggu         int     `json:"umur_hamil_hpht_minggu"`
// // 	HPLBerdasarkanHPHT          string  `json:"hpl_berdasarkan_hpht"`
// // 	UmurHamilUSGMinggu          int     `json:"umur_hamil_usg_minggu"`
// // 	HPLBerdasarkanUSG           string  `json:"hpl_berdasarkan_usg"`
// // 	USGJumlahGS                 string  `json:"usg_jumlah_gs"`
// // 	USGDiameterGS_cm            float64 `json:"usg_diameter_gs_cm"`
// // 	USGDiameterGSMinggu         int     `json:"usg_diameter_gs_minggu"`
// // 	USGDiameterGSHari           int     `json:"usg_diameter_gs_hari"`
// // 	USGJumlahBayi               string  `json:"usg_jumlah_bayi"`
// // 	USGCRL_cm                   float64 `json:"usg_crl_cm"`
// // 	USGCRLMinggu                int     `json:"usg_crl_minggu"`
// // 	USGCRLHari                  int     `json:"usg_crl_hari"`
// // 	USGLetakProdukKehamilan     string  `json:"usg_letak_produk_kehamilan"`
// // 	USGPulsasiJantung           string  `json:"usg_pulsasi_jantung"`
// // 	USGKecurigaanTemuanAbnormal string  `json:"usg_kecurigaan_temuan_abnormal"`
// // 	USGKeteranganTemuanAbnormal string  `json:"usg_keterangan_temuan_abnormal"`
// // }

// // func (c *PemeriksaanDokterTrimester1Controller) Create(ctx echo.Context) error {
// // 	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
// // 	if !ok || claims == nil {
// // 		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
// // 	}
// // 	var req createPemeriksaanDokterT1Request
// // 	if err := ctx.Bind(&req); err != nil {
// // 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
// // 	}
// // 	p := &models.PemeriksaanDokterTrimester1{
// // 		KehamilanID:                 req.KehamilanID,
// // 		NamaDokter:                  req.NamaDokter,
// // 		KonsepAnamnesaPemeriksaan:   req.KonsepAnamnesaPemeriksaan,
// // 		FisikKonjungtiva:            req.FisikKonjungtiva,
// // 		FisikSklera:                 req.FisikSklera,
// // 		FisikKulit:                  req.FisikKulit,
// // 		FisikLeher:                  req.FisikLeher,
// // 		FisikGigiMulut:              req.FisikGigiMulut,
// // 		FisikTHT:                    req.FisikTHT,
// // 		FisikDadaJantung:            req.FisikDadaJantung,
// // 		FisikDadaParu:               req.FisikDadaParu,
// // 		FisikPerut:                  req.FisikPerut,
// // 		FisikTungkai:                req.FisikTungkai,
// // 		KeteraturanHaid:             req.KeteraturanHaid,
// // 		UmurHamilHPHTMinggu:         &req.UmurHamilHPHTMinggu,
// // 		UmurHamilUSGMinggu:          &req.UmurHamilUSGMinggu,
// // 		USGJumlahGS:                 req.USGJumlahGS,
// // 		USGDiameterGS_cm:            &req.USGDiameterGS_cm,
// // 		USGDiameterGSMinggu:         &req.USGDiameterGSMinggu,
// // 		USGDiameterGSHari:           &req.USGDiameterGSHari,
// // 		USGJumlahBayi:               req.USGJumlahBayi,
// // 		USGCRL_cm:                   &req.USGCRL_cm,
// // 		USGCRLMinggu:                &req.USGCRLMinggu,
// // 		USGCRLHari:                  &req.USGCRLHari,
// // 		USGLetakProdukKehamilan:     req.USGLetakProdukKehamilan,
// // 		USGPulsasiJantung:           req.USGPulsasiJantung,
// // 		USGKecurigaanTemuanAbnormal: req.USGKecurigaanTemuanAbnormal,
// // 		USGKeteranganTemuanAbnormal: req.USGKeteranganTemuanAbnormal,
// // 	}
// // 	if req.TanggalPeriksa != "" {
// // 		if t, err := time.Parse("2006-01-02", req.TanggalPeriksa); err == nil {
// // 			p.TanggalPeriksa = &t
// // 		}
// // 	}
// // 	if req.HPHT != "" {
// // 		if t, err := time.Parse("2006-01-02", req.HPHT); err == nil {
// // 			p.HPHT = &t
// // 		}
// // 	}
// // 	if req.HPLBerdasarkanHPHT != "" {
// // 		if t, err := time.Parse("2006-01-02", req.HPLBerdasarkanHPHT); err == nil {
// // 			p.HPLBerdasarkanHPHT = &t
// // 		}
// // 	}
// // 	if req.HPLBerdasarkanUSG != "" {
// // 		if t, err := time.Parse("2006-01-02", req.HPLBerdasarkanUSG); err == nil {
// // 			p.HPLBerdasarkanUSG = &t
// // 		}
// // 	}
// // 	if err := c.usecase.Create(p); err != nil {
// // 		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
// // 	}
// // 	return ctx.JSON(http.StatusCreated, models.Response{StatusCode: http.StatusCreated, Data: p})
// // }

// // func (c *PemeriksaanDokterTrimester1Controller) GetByID(ctx echo.Context) error {
// // 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// // 	if err != nil {
// // 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
// // 	}
// // 	data, err := c.usecase.GetByID(int32(id))
// // 	if err != nil {
// // 		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: err.Error()})
// // 	}
// // 	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: data})
// // }

// // func (c *PemeriksaanDokterTrimester1Controller) GetByKehamilanID(ctx echo.Context) error {
// // 	kehamilanID, err := strconv.ParseInt(ctx.QueryParam("kehamilan_id"), 10, 32)
// // 	if err != nil {
// // 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "kehamilan_id required"})
// // 	}
// // 	list, err := c.usecase.GetByKehamilanID(int32(kehamilanID))
// // 	if err != nil {
// // 		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
// // 	}
// // 	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: list})
// // }

// // func (c *PemeriksaanDokterTrimester1Controller) Update(ctx echo.Context) error {
// // 	// Update mirip dengan create, saya singkat untuk menghemat tempat.
// // 	// Pola: binding, ambil existing, update field yang dikirim, lalu panggil usecase.Update.
// // 	// Sama seperti sebelumnya.
// // 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// // 	if err != nil {
// // 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
// // 	}
// // 	var req createPemeriksaanDokterT1Request
// // 	if err := ctx.Bind(&req); err != nil {
// // 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
// // 	}
// // 	existing, err := c.usecase.GetByID(int32(id))
// // 	if err != nil {
// // 		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: "Data tidak ditemukan"})
// // 	}
// // 	// Update field yang tidak zero value (mirip dengan sebelumnya)
// // 	// Saya asumsikan semua field bisa diupdate.
// // 	if req.NamaDokter != "" {
// // 		existing.NamaDokter = req.NamaDokter
// // 	}
// // 	if req.TanggalPeriksa != "" {
// // 		if t, err := time.Parse("2006-01-02", req.TanggalPeriksa); err == nil {
// // 			existing.TanggalPeriksa = &t
// // 		}
// // 	}
// // 	// ... update semua field lainnya
// // 	if err := c.usecase.Update(existing); err != nil {
// // 		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
// // 	}
// // 	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: existing})
// // }

// // func (c *PemeriksaanDokterTrimester1Controller) Delete(ctx echo.Context) error {
// // 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// // 	if err != nil {
// // 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
// // 	}
// // 	if err := c.usecase.Delete(int32(id)); err != nil {
// // 		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
// // 	}
// // 	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Message: "deleted"})
// // }

// // // MODUL IBU
// // func (c *PemeriksaanDokterTrimester1Controller) GetMine(ctx echo.Context) error {
// // 	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
// // 	if !ok || claims == nil {
// // 		return ctx.JSON(http.StatusUnauthorized, models.Response{
// // 			StatusCode: http.StatusUnauthorized,
// // 			Message:    "token tidak valid",
// // 		})
// // 	}

// // 	data, err := c.usecase.GetMine(claims.UserID)
// // 	if err != nil {
// // 		return ctx.JSON(http.StatusNotFound, models.Response{
// // 			StatusCode: http.StatusNotFound,
// // 			Message:    "data pemeriksaan dokter trimester 1 tidak ditemukan",
// // 		})
// // 	}

// // 	return ctx.JSON(http.StatusOK, models.Response{
// // 		StatusCode: http.StatusOK,
// // 		Data:       data,
// // 	})
// // }

// package controllers

// import (
// 	"monitoring-service/app/models"
// 	"monitoring-service/app/usecases"
// 	"monitoring-service/pkg/customerror"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/labstack/echo/v4"
// )

// type PemeriksaanDokterTrimester1Controller struct {
// 	usecase usecases.PemeriksaanDokterTrimester1Usecase
// }

// func NewPemeriksaanDokterTrimester1Controller(u usecases.PemeriksaanDokterTrimester1Usecase) *PemeriksaanDokterTrimester1Controller {
// 	return &PemeriksaanDokterTrimester1Controller{usecase: u}
// }

// func (c *PemeriksaanDokterTrimester1Controller) Create(ctx echo.Context) error {
// 	var req usecases.PemeriksaanDokterTrimester1Request
// 	if err := ctx.Bind(&req); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "format request tidak valid"})
// 	}
// 	if err := c.usecase.Create(&req); err != nil {
// 		return ctx.JSON(customerror.GetStatusCode(err), map[string]string{"message": err.Error()})
// 	}
// 	return ctx.JSON(http.StatusCreated, map[string]interface{}{
// 		"status_code": http.StatusCreated,
// 		"message":     "Data pemeriksaan dokter trimester 1 berhasil disimpan",
// 	})
// }

// func (c *PemeriksaanDokterTrimester1Controller) Update(ctx echo.Context) error {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
// 	}
// 	var req usecases.PemeriksaanDokterTrimester1Request
// 	if err := ctx.Bind(&req); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "format request tidak valid"})
// 	}
// 	if err := c.usecase.Update(int32(id), &req); err != nil {
// 		return ctx.JSON(customerror.GetStatusCode(err), map[string]string{"message": err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, map[string]interface{}{
// 		"status_code": http.StatusOK,
// 		"message":     "Data berhasil diperbarui",
// 	})
// }

// func (c *PemeriksaanDokterTrimester1Controller) GetByID(ctx echo.Context) error {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
// 	}
// 	data, err := c.usecase.GetByID(int32(id))
// 	if err != nil {
// 		return ctx.JSON(customerror.GetStatusCode(err), map[string]string{"message": err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, data)
// }

// func (c *PemeriksaanDokterTrimester1Controller) GetByKehamilanID(ctx echo.Context) error {
// 	kehamilanID, err := strconv.ParseInt(ctx.QueryParam("kehamilan_id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "kehamilan_id required"})
// 	}
// 	list, err := c.usecase.GetByKehamilanID(int32(kehamilanID))
// 	if err != nil {
// 		return ctx.JSON(customerror.GetStatusCode(err), map[string]string{"message": err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: list})
// }

// func (c *PemeriksaanDokterTrimester1Controller) Update(ctx echo.Context) error {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
// 	}
// 	var req createPemeriksaanDokterT1Request
// 	if err := ctx.Bind(&req); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
// 	}
// 	existing, err := c.usecase.GetByID(int32(id))
// 	if err != nil {
// 		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: "Data tidak ditemukan"})
// 	}
// 	if req.NamaDokter != "" {
// 		existing.NamaDokter = req.NamaDokter
// 	}
// 	if req.TanggalPeriksa != "" {
// 		if t, err := time.Parse("2006-01-02", req.TanggalPeriksa); err == nil {
// 			existing.TanggalPeriksa = &t
// 		}
// 	}
// 	if err := c.usecase.Update(existing); err != nil {
// 		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: existing})
// }

// func (c *PemeriksaanDokterTrimester1Controller) Delete(ctx echo.Context) error {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
// 	}
// 	if err := c.usecase.Delete(int32(id)); err != nil {
// 		return ctx.JSON(customerror.GetStatusCode(err), map[string]string{"message": err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, map[string]string{"message": "Data berhasil dihapus"})
// }

// // MODUL IBU - single (existing)
// func (c *PemeriksaanDokterTrimester1Controller) GetMine(ctx echo.Context) error {
// 	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
// 	if !ok || claims == nil {
// 		return ctx.JSON(http.StatusUnauthorized, models.Response{
// 			StatusCode: http.StatusUnauthorized,
// 			Message:    "token tidak valid",
// 		})
// 	}

// 	data, err := c.usecase.GetMine(claims.UserID)
// 	if err != nil {
// 		return ctx.JSON(http.StatusNotFound, models.Response{
// 			StatusCode: http.StatusNotFound,
// 			Message:    "data pemeriksaan dokter trimester 1 tidak ditemukan",
// 		})
// 	}

// 	return ctx.JSON(http.StatusOK, models.Response{
// 		StatusCode: http.StatusOK,
// 		Data:       data,
// 	})
// }

// // MODUL IBU - list (BARU)
// func (c *PemeriksaanDokterTrimester1Controller) GetAllMine(ctx echo.Context) error {
// 	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
// 	if !ok || claims == nil {
// 		return ctx.JSON(http.StatusUnauthorized, models.Response{
// 			StatusCode: http.StatusUnauthorized,
// 			Message:    "token tidak valid",
// 		})
// 	}

// 	list, err := c.usecase.GetAllMine(claims.UserID)
// 	if err != nil {
// 		return ctx.JSON(http.StatusInternalServerError, models.Response{
// 			StatusCode: http.StatusInternalServerError,
// 			Message:    err.Error(),
// 		})
// 	}

// 	return ctx.JSON(http.StatusOK, models.Response{
// 		StatusCode: http.StatusOK,
// 		Data:       list,
// 	})
// }

// package controllers

// import (
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"monitoring-service/app/models"
// 	"monitoring-service/app/usecases"

// 	"github.com/labstack/echo/v4"
// )

// type PemeriksaanDokterTrimester1Controller struct {
// 	usecase usecases.PemeriksaanDokterTrimester1Usecase
// }

// func NewPemeriksaanDokterTrimester1Controller(u usecases.PemeriksaanDokterTrimester1Usecase) *PemeriksaanDokterTrimester1Controller {
// 	return &PemeriksaanDokterTrimester1Controller{usecase: u}
// }

// type createPemeriksaanDokterT1Request struct {
// 	KehamilanID                 int32   `json:"kehamilan_id"`
// 	NamaDokter                  string  `json:"nama_dokter"`
// 	TanggalPeriksa              string  `json:"tanggal_periksa"`
// 	KonsepAnamnesaPemeriksaan   string  `json:"konsep_anamnesa_pemeriksaan"`
// 	FisikKonjungtiva            string  `json:"fisik_konjungtiva"`
// 	FisikSklera                 string  `json:"fisik_sklera"`
// 	FisikKulit                  string  `json:"fisik_kulit"`
// 	FisikLeher                  string  `json:"fisik_leher"`
// 	FisikGigiMulut              string  `json:"fisik_gigi_mulut"`
// 	FisikTHT                    string  `json:"fisik_tht"`
// 	FisikDadaJantung            string  `json:"fisik_dada_jantung"`
// 	FisikDadaParu               string  `json:"fisik_dada_paru"`
// 	FisikPerut                  string  `json:"fisik_perut"`
// 	FisikTungkai                string  `json:"fisik_tungkai"`
// 	HPHT                        string  `json:"hpht"`
// 	KeteraturanHaid             string  `json:"keteraturan_haid"`
// 	UmurHamilHPHTMinggu         int     `json:"umur_hamil_hpht_minggu"`
// 	HPLBerdasarkanHPHT          string  `json:"hpl_berdasarkan_hpht"`
// 	UmurHamilUSGMinggu          int     `json:"umur_hamil_usg_minggu"`
// 	HPLBerdasarkanUSG           string  `json:"hpl_berdasarkan_usg"`
// 	USGJumlahGS                 string  `json:"usg_jumlah_gs"`
// 	USGDiameterGS_cm            float64 `json:"usg_diameter_gs_cm"`
// 	USGDiameterGSMinggu         int     `json:"usg_diameter_gs_minggu"`
// 	USGDiameterGSHari           int     `json:"usg_diameter_gs_hari"`
// 	USGJumlahBayi               string  `json:"usg_jumlah_bayi"`
// 	USGCRL_cm                   float64 `json:"usg_crl_cm"`
// 	USGCRLMinggu                int     `json:"usg_crl_minggu"`
// 	USGCRLHari                  int     `json:"usg_crl_hari"`
// 	USGLetakProdukKehamilan     string  `json:"usg_letak_produk_kehamilan"`
// 	USGPulsasiJantung           string  `json:"usg_pulsasi_jantung"`
// 	USGKecurigaanTemuanAbnormal string  `json:"usg_kecurigaan_temuan_abnormal"`
// 	USGKeteranganTemuanAbnormal string  `json:"usg_keterangan_temuan_abnormal"`
// }

// func (c *PemeriksaanDokterTrimester1Controller) Create(ctx echo.Context) error {
// 	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
// 	if !ok || claims == nil {
// 		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
// 	}
// 	var req createPemeriksaanDokterT1Request
// 	if err := ctx.Bind(&req); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
// 	}
// 	p := &models.PemeriksaanDokterTrimester1{
// 		KehamilanID:                 req.KehamilanID,
// 		NamaDokter:                  req.NamaDokter,
// 		KonsepAnamnesaPemeriksaan:   req.KonsepAnamnesaPemeriksaan,
// 		FisikKonjungtiva:            req.FisikKonjungtiva,
// 		FisikSklera:                 req.FisikSklera,
// 		FisikKulit:                  req.FisikKulit,
// 		FisikLeher:                  req.FisikLeher,
// 		FisikGigiMulut:              req.FisikGigiMulut,
// 		FisikTHT:                    req.FisikTHT,
// 		FisikDadaJantung:            req.FisikDadaJantung,
// 		FisikDadaParu:               req.FisikDadaParu,
// 		FisikPerut:                  req.FisikPerut,
// 		FisikTungkai:                req.FisikTungkai,
// 		KeteraturanHaid:             req.KeteraturanHaid,
// 		UmurHamilHPHTMinggu:         &req.UmurHamilHPHTMinggu,
// 		UmurHamilUSGMinggu:          &req.UmurHamilUSGMinggu,
// 		USGJumlahGS:                 req.USGJumlahGS,
// 		USGDiameterGS_cm:            &req.USGDiameterGS_cm,
// 		USGDiameterGSMinggu:         &req.USGDiameterGSMinggu,
// 		USGDiameterGSHari:           &req.USGDiameterGSHari,
// 		USGJumlahBayi:               req.USGJumlahBayi,
// 		USGCRL_cm:                   &req.USGCRL_cm,
// 		USGCRLMinggu:                &req.USGCRLMinggu,
// 		USGCRLHari:                  &req.USGCRLHari,
// 		USGLetakProdukKehamilan:     req.USGLetakProdukKehamilan,
// 		USGPulsasiJantung:           req.USGPulsasiJantung,
// 		USGKecurigaanTemuanAbnormal: req.USGKecurigaanTemuanAbnormal,
// 		USGKeteranganTemuanAbnormal: req.USGKeteranganTemuanAbnormal,
// 	}
// 	if req.TanggalPeriksa != "" {
// 		if t, err := time.Parse("2006-01-02", req.TanggalPeriksa); err == nil {
// 			p.TanggalPeriksa = &t
// 		}
// 	}
// 	if req.HPHT != "" {
// 		if t, err := time.Parse("2006-01-02", req.HPHT); err == nil {
// 			p.HPHT = &t
// 		}
// 	}
// 	if req.HPLBerdasarkanHPHT != "" {
// 		if t, err := time.Parse("2006-01-02", req.HPLBerdasarkanHPHT); err == nil {
// 			p.HPLBerdasarkanHPHT = &t
// 		}
// 	}
// 	if req.HPLBerdasarkanUSG != "" {
// 		if t, err := time.Parse("2006-01-02", req.HPLBerdasarkanUSG); err == nil {
// 			p.HPLBerdasarkanUSG = &t
// 		}
// 	}
// 	if err := c.usecase.Create(p); err != nil {
// 		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
// 	}
// 	return ctx.JSON(http.StatusCreated, models.Response{StatusCode: http.StatusCreated, Data: p})
// }

// func (c *PemeriksaanDokterTrimester1Controller) GetByID(ctx echo.Context) error {
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

// func (c *PemeriksaanDokterTrimester1Controller) GetByKehamilanID(ctx echo.Context) error {
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

// func (c *PemeriksaanDokterTrimester1Controller) Update(ctx echo.Context) error {
// 	// Update mirip dengan create, saya singkat untuk menghemat tempat.
// 	// Pola: binding, ambil existing, update field yang dikirim, lalu panggil usecase.Update.
// 	// Sama seperti sebelumnya.
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
// 	}
// 	var req createPemeriksaanDokterT1Request
// 	if err := ctx.Bind(&req); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
// 	}
// 	existing, err := c.usecase.GetByID(int32(id))
// 	if err != nil {
// 		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: "Data tidak ditemukan"})
// 	}
// 	// Update field yang tidak zero value (mirip dengan sebelumnya)
// 	// Saya asumsikan semua field bisa diupdate.
// 	if req.NamaDokter != "" {
// 		existing.NamaDokter = req.NamaDokter
// 	}
// 	if req.TanggalPeriksa != "" {
// 		if t, err := time.Parse("2006-01-02", req.TanggalPeriksa); err == nil {
// 			existing.TanggalPeriksa = &t
// 		}
// 	}
// 	// ... update semua field lainnya
// 	if err := c.usecase.Update(existing); err != nil {
// 		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: existing})
// }

// func (c *PemeriksaanDokterTrimester1Controller) Delete(ctx echo.Context) error {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
// 	}
// 	if err := c.usecase.Delete(int32(id)); err != nil {
// 		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Message: "deleted"})
// }

// // MODUL IBU
// func (c *PemeriksaanDokterTrimester1Controller) GetMine(ctx echo.Context) error {
// 	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
// 	if !ok || claims == nil {
// 		return ctx.JSON(http.StatusUnauthorized, models.Response{
// 			StatusCode: http.StatusUnauthorized,
// 			Message:    "token tidak valid",
// 		})
// 	}

// 	data, err := c.usecase.GetMine(claims.UserID)
// 	if err != nil {
// 		return ctx.JSON(http.StatusNotFound, models.Response{
// 			StatusCode: http.StatusNotFound,
// 			Message:    "data pemeriksaan dokter trimester 1 tidak ditemukan",
// 		})
// 	}

// 	return ctx.JSON(http.StatusOK, models.Response{
// 		StatusCode: http.StatusOK,
// 		Data:       data,
// 	})
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

type PemeriksaanDokterTrimester1Controller struct {
	usecase usecases.PemeriksaanDokterTrimester1Usecase
}

func NewPemeriksaanDokterTrimester1Controller(u usecases.PemeriksaanDokterTrimester1Usecase) *PemeriksaanDokterTrimester1Controller {
	return &PemeriksaanDokterTrimester1Controller{usecase: u}
}

type createPemeriksaanDokterT1Request struct {
	KehamilanID                 int32   `json:"kehamilan_id"`
	NamaDokter                  string  `json:"nama_dokter"`
	TanggalPeriksa              string  `json:"tanggal_periksa"`
	KonsepAnamnesaPemeriksaan   string  `json:"konsep_anamnesa_pemeriksaan"`
	FisikKonjungtiva            string  `json:"fisik_konjungtiva"`
	FisikSklera                 string  `json:"fisik_sklera"`
	FisikKulit                  string  `json:"fisik_kulit"`
	FisikLeher                  string  `json:"fisik_leher"`
	FisikGigiMulut              string  `json:"fisik_gigi_mulut"`
	FisikTHT                    string  `json:"fisik_tht"`
	FisikDadaJantung            string  `json:"fisik_dada_jantung"`
	FisikDadaParu               string  `json:"fisik_dada_paru"`
	FisikPerut                  string  `json:"fisik_perut"`
	FisikTungkai                string  `json:"fisik_tungkai"`
	HPHT                        string  `json:"hpht"`
	KeteraturanHaid             string  `json:"keteraturan_haid"`
	UmurHamilHPHTMinggu         int     `json:"umur_hamil_hpht_minggu"`
	HPLBerdasarkanHPHT          string  `json:"hpl_berdasarkan_hpht"`
	UmurHamilUSGMinggu          int     `json:"umur_hamil_usg_minggu"`
	HPLBerdasarkanUSG           string  `json:"hpl_berdasarkan_usg"`
	USGJumlahGS                 string  `json:"usg_jumlah_gs"`
	USGDiameterGS_cm            float64 `json:"usg_diameter_gs_cm"`
	USGDiameterGSMinggu         int     `json:"usg_diameter_gs_minggu"`
	USGDiameterGSHari           int     `json:"usg_diameter_gs_hari"`
	USGJumlahBayi               string  `json:"usg_jumlah_bayi"`
	USGCRL_cm                   float64 `json:"usg_crl_cm"`
	USGCRLMinggu                int     `json:"usg_crl_minggu"`
	USGCRLHari                  int     `json:"usg_crl_hari"`
	USGLetakProdukKehamilan     string  `json:"usg_letak_produk_kehamilan"`
	USGPulsasiJantung           string  `json:"usg_pulsasi_jantung"`
	USGKecurigaanTemuanAbnormal string  `json:"usg_kecurigaan_temuan_abnormal"`
	USGKeteranganTemuanAbnormal string  `json:"usg_keterangan_temuan_abnormal"`
}

func (c *PemeriksaanDokterTrimester1Controller) Create(ctx echo.Context) error {
	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
	if !ok || claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
	}
	var req createPemeriksaanDokterT1Request
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	p := &models.PemeriksaanDokterTrimester1{
		KehamilanID:                 req.KehamilanID,
		NamaDokter:                  req.NamaDokter,
		KonsepAnamnesaPemeriksaan:   req.KonsepAnamnesaPemeriksaan,
		FisikKonjungtiva:            req.FisikKonjungtiva,
		FisikSklera:                 req.FisikSklera,
		FisikKulit:                  req.FisikKulit,
		FisikLeher:                  req.FisikLeher,
		FisikGigiMulut:              req.FisikGigiMulut,
		FisikTHT:                    req.FisikTHT,
		FisikDadaJantung:            req.FisikDadaJantung,
		FisikDadaParu:               req.FisikDadaParu,
		FisikPerut:                  req.FisikPerut,
		FisikTungkai:                req.FisikTungkai,
		KeteraturanHaid:             req.KeteraturanHaid,
		UmurHamilHPHTMinggu:         &req.UmurHamilHPHTMinggu,
		UmurHamilUSGMinggu:          &req.UmurHamilUSGMinggu,
		USGJumlahGS:                 req.USGJumlahGS,
		USGDiameterGS_cm:            &req.USGDiameterGS_cm,
		USGDiameterGSMinggu:         &req.USGDiameterGSMinggu,
		USGDiameterGSHari:           &req.USGDiameterGSHari,
		USGJumlahBayi:               req.USGJumlahBayi,
		USGCRL_cm:                   &req.USGCRL_cm,
		USGCRLMinggu:                &req.USGCRLMinggu,
		USGCRLHari:                  &req.USGCRLHari,
		USGLetakProdukKehamilan:     req.USGLetakProdukKehamilan,
		USGPulsasiJantung:           req.USGPulsasiJantung,
		USGKecurigaanTemuanAbnormal: req.USGKecurigaanTemuanAbnormal,
		USGKeteranganTemuanAbnormal: req.USGKeteranganTemuanAbnormal,
	}
	if req.TanggalPeriksa != "" {
		if t, err := time.Parse("2006-01-02", req.TanggalPeriksa); err == nil {
			p.TanggalPeriksa = &t
		}
	}
	if req.HPHT != "" {
		if t, err := time.Parse("2006-01-02", req.HPHT); err == nil {
			p.HPHT = &t
		}
	}
	if req.HPLBerdasarkanHPHT != "" {
		if t, err := time.Parse("2006-01-02", req.HPLBerdasarkanHPHT); err == nil {
			p.HPLBerdasarkanHPHT = &t
		}
	}
	if req.HPLBerdasarkanUSG != "" {
		if t, err := time.Parse("2006-01-02", req.HPLBerdasarkanUSG); err == nil {
			p.HPLBerdasarkanUSG = &t
		}
	}
	if err := c.usecase.Create(p); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusCreated, models.Response{StatusCode: http.StatusCreated, Data: p})
}

func (c *PemeriksaanDokterTrimester1Controller) GetByID(ctx echo.Context) error {
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

func (c *PemeriksaanDokterTrimester1Controller) GetByKehamilanID(ctx echo.Context) error {
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

func (c *PemeriksaanDokterTrimester1Controller) Update(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	var req createPemeriksaanDokterT1Request
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	existing, err := c.usecase.GetByID(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: "Data tidak ditemukan"})
	}
	if req.NamaDokter != "" {
		existing.NamaDokter = req.NamaDokter
	}
	if req.TanggalPeriksa != "" {
		if t, err := time.Parse("2006-01-02", req.TanggalPeriksa); err == nil {
			existing.TanggalPeriksa = &t
		}
	}
	if err := c.usecase.Update(int32(id),existing); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: existing})
}

func (c *PemeriksaanDokterTrimester1Controller) Delete(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	if err := c.usecase.Delete(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Message: "deleted"})
}

// MODUL IBU - single (existing)
func (c *PemeriksaanDokterTrimester1Controller) GetMine(ctx echo.Context) error {
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
			Message:    "data pemeriksaan dokter trimester 1 tidak ditemukan",
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Data:       data,
	})
}

// MODUL IBU - list (BARU)
func (c *PemeriksaanDokterTrimester1Controller) GetAllMine(ctx echo.Context) error {
	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
	if !ok || claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "token tidak valid",
		})
	}

	list, err := c.usecase.GetAllMine(claims.UserID)
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
