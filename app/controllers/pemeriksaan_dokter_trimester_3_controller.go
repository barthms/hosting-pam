// package controllers

// import (
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"monitoring-service/app/models"
// 	"monitoring-service/app/usecases"

// 	"github.com/labstack/echo/v4"
// )

// type PemeriksaanDokterTrimester3Controller struct {
// 	usecase usecases.PemeriksaanDokterTrimester3Usecase
// }

// func NewPemeriksaanDokterTrimester3Controller(u usecases.PemeriksaanDokterTrimester3Usecase) *PemeriksaanDokterTrimester3Controller {
// 	return &PemeriksaanDokterTrimester3Controller{usecase: u}
// }

// type createPemeriksaanDokterT3Request struct {
// 	KehamilanID                              int32   `json:"kehamilan_id"`
// 	NamaDokter                               string  `json:"nama_dokter"`
// 	TanggalPeriksa                           string  `json:"tanggal_periksa"`
// 	KonsepAnamnesaPemeriksaan                string  `json:"konsep_anamnesa_pemeriksaan"`
// 	FisikKonjungtiva                         string  `json:"fisik_konjungtiva"`
// 	FisikSklera                              string  `json:"fisik_sklera"`
// 	FisikKulit                               string  `json:"fisik_kulit"`
// 	FisikLeher                               string  `json:"fisik_leher"`
// 	FisikGigiMulut                           string  `json:"fisik_gigi_mulut"`
// 	FisikTHT                                 string  `json:"fisik_tht"`
// 	FisikDadaJantung                         string  `json:"fisik_dada_jantung"`
// 	FisikDadaParu                            string  `json:"fisik_dada_paru"`
// 	FisikPerut                               string  `json:"fisik_perut"`
// 	FisikTungkai                             string  `json:"fisik_tungkai"`
// 	USGTrimester3Dilakukan                   string  `json:"usg_trimester_3_dilakukan"`
// 	UKBerdasarkanUSGTrimester1Minggu         int     `json:"uk_berdasarkan_usg_trimester_1_minggu"`
// 	UKBerdasarkanHPHTMinggu                  int     `json:"uk_berdasarkan_hpht_minggu"`
// 	UKBerdasarkanBiometriUSGTrimester3Minggu int     `json:"uk_berdasarkan_biometri_usg_trimester_3_minggu"`
// 	SelisihUK3MingguAtauLebih                string  `json:"selisih_uk_3_minggu_atau_lebih"`
// 	USGJumlahBayi                            string  `json:"usg_jumlah_bayi"`
// 	USGLetakBayi                             string  `json:"usg_letak_bayi"`
// 	USGPresentasiBayi                        string  `json:"usg_presentasi_bayi"`
// 	USGKeadaanBayi                           string  `json:"usg_keadaan_bayi"`
// 	USGDJNilai                               int     `json:"usg_djj_nilai"`
// 	USGDJJStatus                             string  `json:"usg_djj_status"`
// 	USGLokasiPlasenta                        string  `json:"usg_lokasi_plasenta"`
// 	USGCairanKetubanSDPCm                    float64 `json:"usg_cairan_ketuban_sdp_cm"`
// 	USGCairanKetubanStatus                   string  `json:"usg_cairan_ketuban_status"`
// 	BiometriBPDCm                            float64 `json:"biometri_bpd_cm"`
// 	BiometriBPDMinggu                        int     `json:"biometri_bpd_minggu"`
// 	BiometriHCCm                             float64 `json:"biometri_hc_cm"`
// 	BiometriHCMinggu                         int     `json:"biometri_hc_minggu"`
// 	BiometriACCm                             float64 `json:"biometri_ac_cm"`
// 	BiometriACMinggu                         int     `json:"biometri_ac_minggu"`
// 	BiometriFLCm                             float64 `json:"biometri_fl_cm"`
// 	BiometriFLMinggu                         int     `json:"biometri_fl_minggu"`
// 	BiometriEFWTBJGram                       int     `json:"biometri_efw_tbj_gram"`
// 	BiometriEFWTBJMinggu                     int     `json:"biometri_efw_tbj_minggu"`
// 	USGKecurigaanTemuanAbnormal              string  `json:"usg_kecurigaan_temuan_abnormal"`
// 	USGKeteranganTemuanAbnormal              string  `json:"usg_keterangan_temuan_abnormal"`
// }

// func (c *PemeriksaanDokterTrimester3Controller) Create(ctx echo.Context) error {
// 	claims, _ := ctx.Get("auth_claims").(*models.AuthClaims)
// 	if claims == nil {
// 		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
// 	}
// 	var req createPemeriksaanDokterT3Request
// 	if err := ctx.Bind(&req); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
// 	}
// 	p := &models.PemeriksaanDokterTrimester3{
// 		KehamilanID:                              req.KehamilanID,
// 		NamaDokter:                               req.NamaDokter,
// 		KonsepAnamnesaPemeriksaan:                req.KonsepAnamnesaPemeriksaan,
// 		FisikKonjungtiva:                         req.FisikKonjungtiva,
// 		FisikSklera:                              req.FisikSklera,
// 		FisikKulit:                               req.FisikKulit,
// 		FisikLeher:                               req.FisikLeher,
// 		FisikGigiMulut:                           req.FisikGigiMulut,
// 		FisikTHT:                                 req.FisikTHT,
// 		FisikDadaJantung:                         req.FisikDadaJantung,
// 		FisikDadaParu:                            req.FisikDadaParu,
// 		FisikPerut:                               req.FisikPerut,
// 		FisikTungkai:                             req.FisikTungkai,
// 		USGTrimester3Dilakukan:                   req.USGTrimester3Dilakukan,
// 		UKBerdasarkanUSGTrimester1Minggu:         &req.UKBerdasarkanUSGTrimester1Minggu,
// 		UKBerdasarkanHPHTMinggu:                  &req.UKBerdasarkanHPHTMinggu,
// 		UKBerdasarkanBiometriUSGTrimester3Minggu: &req.UKBerdasarkanBiometriUSGTrimester3Minggu,
// 		SelisihUK3MingguAtauLebih:                req.SelisihUK3MingguAtauLebih,
// 		USGJumlahBayi:                            req.USGJumlahBayi,
// 		USGLetakBayi:                             req.USGLetakBayi,
// 		USGPresentasiBayi:                        req.USGPresentasiBayi,
// 		USGKeadaanBayi:                           req.USGKeadaanBayi,
// 		USGDJNilai:                               &req.USGDJNilai,
// 		USGDJJStatus:                             req.USGDJJStatus,
// 		USGLokasiPlasenta:                        req.USGLokasiPlasenta,
// 		USGCairanKetubanSDPCm:                    &req.USGCairanKetubanSDPCm,
// 		USGCairanKetubanStatus:                   req.USGCairanKetubanStatus,
// 		BiometriBPDCm:                            &req.BiometriBPDCm,
// 		BiometriBPDMinggu:                        &req.BiometriBPDMinggu,
// 		BiometriHCCm:                             &req.BiometriHCCm,
// 		BiometriHCMinggu:                         &req.BiometriHCMinggu,
// 		BiometriACCm:                             &req.BiometriACCm,
// 		BiometriACMinggu:                         &req.BiometriACMinggu,
// 		BiometriFLCm:                             &req.BiometriFLCm,
// 		BiometriFLMinggu:                         &req.BiometriFLMinggu,
// 		BiometriEFWTBJGram:                       &req.BiometriEFWTBJGram,
// 		BiometriEFWTBJMinggu:                     &req.BiometriEFWTBJMinggu,
// 		USGKecurigaanTemuanAbnormal:              req.USGKecurigaanTemuanAbnormal,
// 		USGKeteranganTemuanAbnormal:              req.USGKeteranganTemuanAbnormal,
// 	}
// 	if req.TanggalPeriksa != "" {
// 		if t, err := time.Parse("2006-01-02", req.TanggalPeriksa); err == nil {
// 			p.TanggalPeriksa = &t
// 		}
// 	}
// 	if err := c.usecase.Create(p); err != nil {
// 		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
// 	}
// 	return ctx.JSON(http.StatusCreated, models.Response{StatusCode: http.StatusCreated, Data: p})
// }

// // GetByID, GetByKehamilanID, Update, Delete mirip seperti sebelumnya (saya tulis ringkas)
// func (c *PemeriksaanDokterTrimester3Controller) GetByID(ctx echo.Context) error {
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

// func (c *PemeriksaanDokterTrimester3Controller) GetByKehamilanID(ctx echo.Context) error {
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

// func (c *PemeriksaanDokterTrimester3Controller) Update(ctx echo.Context) error {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
// 	}
// 	var req createPemeriksaanDokterT3Request
// 	if err := ctx.Bind(&req); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
// 	}
// 	existing, err := c.usecase.GetByID(int32(id))
// 	if err != nil {
// 		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: "Data tidak ditemukan"})
// 	}
// 	// Update field non-zero (contoh sebagian)
// 	if req.NamaDokter != "" {
// 		existing.NamaDokter = req.NamaDokter
// 	}
// 	if req.TanggalPeriksa != "" {
// 		if t, err := time.Parse("2006-01-02", req.TanggalPeriksa); err == nil {
// 			existing.TanggalPeriksa = &t
// 		}
// 	}
// 	// ... update semua field yang diperlukan
// 	if err := c.usecase.Update(existing); err != nil {
// 		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: existing})
// }

// func (c *PemeriksaanDokterTrimester3Controller) Delete(ctx echo.Context) error {
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
// func (c *PemeriksaanDokterTrimester3Controller) GetMine(ctx echo.Context) error {
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
// 			Message:    "data pemeriksaan dokter trimester 3 tidak ditemukan",
// 		})
// 	}

// 	return ctx.JSON(http.StatusOK, models.Response{
// 		StatusCode: http.StatusOK,
// 		Data:       data,
// 	})
// }


// package controllers

// import (
// 	"monitoring-service/app/usecases"
// 	"monitoring-service/pkg/customerror"
// 	"net/http"
// 	"strconv"

// 	"github.com/labstack/echo/v4"
// )

// type PemeriksaanDokterTrimester3Controller struct {
// 	usecase usecases.PemeriksaanDokterTrimester3Usecase
// }

// func NewPemeriksaanDokterTrimester3Controller(u usecases.PemeriksaanDokterTrimester3Usecase) *PemeriksaanDokterTrimester3Controller {
// 	return &PemeriksaanDokterTrimester3Controller{usecase: u}
// }

// func (c *PemeriksaanDokterTrimester3Controller) Create(ctx echo.Context) error {
// 	var req usecases.PemeriksaanDokterTrimester3Request
// 	if err := ctx.Bind(&req); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "format request tidak valid"})
// 	}
// 	if err := c.usecase.Create(&req); err != nil {
// 		return ctx.JSON(customerror.GetStatusCode(err), map[string]string{"message": err.Error()})
// 	}
// 	return ctx.JSON(http.StatusCreated, map[string]interface{}{
// 		"status_code": http.StatusCreated,
// 		"message":     "Data pemeriksaan dokter trimester 3 berhasil disimpan",
// 	})
// }

// func (c *PemeriksaanDokterTrimester3Controller) Update(ctx echo.Context) error {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
// 	}
// 	var req usecases.PemeriksaanDokterTrimester3Request
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

// func (c *PemeriksaanDokterTrimester3Controller) GetByID(ctx echo.Context) error {
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

// func (c *PemeriksaanDokterTrimester3Controller) GetByKehamilanID(ctx echo.Context) error {
// 	kehamilanID, err := strconv.ParseInt(ctx.QueryParam("kehamilan_id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "kehamilan_id required"})
// 	}
// 	list, err := c.usecase.GetByKehamilanID(int32(kehamilanID))
// 	if err != nil {
// 		return ctx.JSON(customerror.GetStatusCode(err), map[string]string{"message": err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, list)
// }

// func (c *PemeriksaanDokterTrimester3Controller) Delete(ctx echo.Context) error {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
// 	}
// 	if err := c.usecase.Delete(int32(id)); err != nil {
// 		return ctx.JSON(customerror.GetStatusCode(err), map[string]string{"message": err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, map[string]string{"message": "Data berhasil dihapus"})
// }

// // MODUL IBU
// func (c *PemeriksaanDokterTrimester3Controller) GetMine(ctx echo.Context) error {
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
// 			Message:    "data pemeriksaan dokter trimester 3 tidak ditemukan",
// 		})
// 	}

// 	return ctx.JSON(http.StatusOK, models.Response{
// 		StatusCode: http.StatusOK,
// 		Data:       data,
// 	})
// }
// // MODUL IBU - list (BARU)
// func (c *PemeriksaanDokterTrimester3Controller) GetAllMine(ctx echo.Context) error {
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
//     "monitoring-service/app/models"
//     "monitoring-service/app/usecases"
//     "monitoring-service/pkg/customerror"
//     "net/http"
//     "strconv"

//     "github.com/labstack/echo/v4"
// )

// type PemeriksaanDokterTrimester3Controller struct {
//     usecase usecases.PemeriksaanDokterTrimester3Usecase
// }

// func NewPemeriksaanDokterTrimester3Controller(u usecases.PemeriksaanDokterTrimester3Usecase) *PemeriksaanDokterTrimester3Controller {
//     return &PemeriksaanDokterTrimester3Controller{usecase: u}
// }

// func (c *PemeriksaanDokterTrimester3Controller) Create(ctx echo.Context) error {
//     var req usecases.PemeriksaanDokterTrimester3Request
//     if err := ctx.Bind(&req); err != nil {
//         return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "format request tidak valid"})
//     }
//     if err := c.usecase.Create(&req); err != nil {
//         return ctx.JSON(customerror.GetStatusCode(err), map[string]string{"message": err.Error()})
//     }
//     return ctx.JSON(http.StatusCreated, map[string]interface{}{
//         "status_code": http.StatusCreated,
//         "message":     "Data pemeriksaan dokter trimester 3 berhasil disimpan",
//     })
// }

// func (c *PemeriksaanDokterTrimester3Controller) Update(ctx echo.Context) error {
//     id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
//     if err != nil {
//         return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
//     }
//     var req usecases.PemeriksaanDokterTrimester3Request
//     if err := ctx.Bind(&req); err != nil {
//         return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "format request tidak valid"})
//     }
//     if err := c.usecase.Update(int32(id), &req); err != nil {
//         return ctx.JSON(customerror.GetStatusCode(err), map[string]string{"message": err.Error()})
//     }
//     return ctx.JSON(http.StatusOK, map[string]interface{}{
//         "status_code": http.StatusOK,
//         "message":     "Data berhasil diperbarui",
//     })
// }

// func (c *PemeriksaanDokterTrimester3Controller) GetByID(ctx echo.Context) error {
//     id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
//     if err != nil {
//         return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
//     }
//     data, err := c.usecase.GetByID(int32(id))
//     if err != nil {
//         return ctx.JSON(customerror.GetStatusCode(err), map[string]string{"message": err.Error()})
//     }
//     return ctx.JSON(http.StatusOK, data)
// }

// func (c *PemeriksaanDokterTrimester3Controller) GetByKehamilanID(ctx echo.Context) error {
//     kehamilanID, err := strconv.ParseInt(ctx.QueryParam("kehamilan_id"), 10, 32)
//     if err != nil {
//         return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "kehamilan_id required"})
//     }
//     list, err := c.usecase.GetByKehamilanID(int32(kehamilanID))
//     if err != nil {
//         return ctx.JSON(customerror.GetStatusCode(err), map[string]string{"message": err.Error()})
//     }
//     return ctx.JSON(http.StatusOK, list)
// }

// func (c *PemeriksaanDokterTrimester3Controller) Delete(ctx echo.Context) error {
//     id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
//     if err != nil {
//         return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
//     }
//     if err := c.usecase.Delete(int32(id)); err != nil {
//         return ctx.JSON(customerror.GetStatusCode(err), map[string]string{"message": err.Error()})
//     }
//     return ctx.JSON(http.StatusOK, map[string]string{"message": "Data berhasil dihapus"})
// }

// // MODUL IBU
// func (c *PemeriksaanDokterTrimester3Controller) GetMine(ctx echo.Context) error {
//     claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
//     if !ok || claims == nil {
//         return ctx.JSON(http.StatusUnauthorized, models.Response{
//             StatusCode: http.StatusUnauthorized,
//             Message:    "token tidak valid",
//         })
//     }

//     data, err := c.usecase.GetMine(claims.UserID)
//     if err != nil {
//         return ctx.JSON(http.StatusNotFound, models.Response{
//             StatusCode: http.StatusNotFound,
//             Message:    "data pemeriksaan dokter trimester 3 tidak ditemukan",
//         })
//     }

//     return ctx.JSON(http.StatusOK, models.Response{
//         StatusCode: http.StatusOK,
//         Data:       data,
//     })
// }

// // MODUL IBU - list (BARU)
// func (c *PemeriksaanDokterTrimester3Controller) GetAllMine(ctx echo.Context) error {
//     claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
//     if !ok || claims == nil {
//         return ctx.JSON(http.StatusUnauthorized, models.Response{
//             StatusCode: http.StatusUnauthorized,
//             Message:    "token tidak valid",
//         })
//     }

//     list, err := c.usecase.GetAllMine(claims.UserID)
//     if err != nil {
//         return ctx.JSON(http.StatusInternalServerError, models.Response{
//             StatusCode: http.StatusInternalServerError,
//             Message:    err.Error(),
//         })
//     }

//     return ctx.JSON(http.StatusOK, models.Response{
//         StatusCode: http.StatusOK,
//         Data:       list,
//     })
// }


package controllers

import (
    "monitoring-service/app/models"
    "monitoring-service/app/usecases"
    "monitoring-service/pkg/customerror"
    "net/http"
    "strconv"

    "github.com/labstack/echo/v4"
)

type PemeriksaanDokterTrimester3Controller struct {
    usecase usecases.PemeriksaanDokterTrimester3Usecase
}

func NewPemeriksaanDokterTrimester3Controller(u usecases.PemeriksaanDokterTrimester3Usecase) *PemeriksaanDokterTrimester3Controller {
    return &PemeriksaanDokterTrimester3Controller{usecase: u}
}

func (c *PemeriksaanDokterTrimester3Controller) Create(ctx echo.Context) error {
    var req models.PemeriksaanDokterTrimester3 // DIUBAH: Bind langsung ke models
    if err := ctx.Bind(&req); err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "format request tidak valid"})
    }
    if err := c.usecase.Create(&req); err != nil {
        return ctx.JSON(customerror.GetStatusCode(err), map[string]string{"message": err.Error()})
    }
    return ctx.JSON(http.StatusCreated, map[string]interface{}{
        "status_code": http.StatusCreated,
        "message":     "Data pemeriksaan dokter trimester 3 berhasil disimpan",
    })
}

func (c *PemeriksaanDokterTrimester3Controller) Update(ctx echo.Context) error {
    id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
    }
    var req models.PemeriksaanDokterTrimester3 // DIUBAH
    if err := ctx.Bind(&req); err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "format request tidak valid"})
    }
    if err := c.usecase.Update(int32(id), &req); err != nil {
        return ctx.JSON(customerror.GetStatusCode(err), map[string]string{"message": err.Error()})
    }
    return ctx.JSON(http.StatusOK, map[string]interface{}{
        "status_code": http.StatusOK,
        "message":     "Data berhasil diperbarui",
    })
}

func (c *PemeriksaanDokterTrimester3Controller) GetByID(ctx echo.Context) error {
    id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
    }
    data, err := c.usecase.GetByID(int32(id))
    if err != nil {
        return ctx.JSON(customerror.GetStatusCode(err), map[string]string{"message": err.Error()})
    }
    return ctx.JSON(http.StatusOK, data)
}

func (c *PemeriksaanDokterTrimester3Controller) GetByKehamilanID(ctx echo.Context) error {
    kehamilanID, err := strconv.ParseInt(ctx.QueryParam("kehamilan_id"), 10, 32)
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "kehamilan_id required"})
    }
    list, err := c.usecase.GetByKehamilanID(int32(kehamilanID))
    if err != nil {
        return ctx.JSON(customerror.GetStatusCode(err), map[string]string{"message": err.Error()})
    }
    return ctx.JSON(http.StatusOK, list)
}

func (c *PemeriksaanDokterTrimester3Controller) Delete(ctx echo.Context) error {
    id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
    }
    if err := c.usecase.Delete(int32(id)); err != nil {
        return ctx.JSON(customerror.GetStatusCode(err), map[string]string{"message": err.Error()})
    }
    return ctx.JSON(http.StatusOK, map[string]string{"message": "Data berhasil dihapus"})
}

// MODUL IBU
func (c *PemeriksaanDokterTrimester3Controller) GetMine(ctx echo.Context) error {
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
            Message:    "data pemeriksaan dokter trimester 3 tidak ditemukan",
        })
    }

    return ctx.JSON(http.StatusOK, models.Response{
        StatusCode: http.StatusOK,
        Data:       data,
    })
}

// MODUL IBU - list (BARU)
func (c *PemeriksaanDokterTrimester3Controller) GetAllMine(ctx echo.Context) error {
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