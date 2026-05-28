package controllers

import (
	"net/http"
	"strconv"
	"time"

	"monitoring-service/app/helpers"
	"monitoring-service/app/models"
	"monitoring-service/app/usecases"
	"monitoring-service/pkg/customerror"

	"github.com/labstack/echo/v4"
)

type KehamilanController struct {
	usecase usecases.KehamilanUsecase
}

func NewKehamilanController(u usecases.KehamilanUsecase) *KehamilanController {
	return &KehamilanController{usecase: u}
}

type createKehamilanRequest struct {
	IbuID              int32  `json:"ibu_id"`
	Gravida            int32  `json:"gravida,omitempty"`
	Paritas            int32  `json:"paritas,omitempty"`
	Abortus            int32  `json:"abortus,omitempty"`
	HPHT               string `json:"hpht,omitempty"`
	TaksiranPersalinan string `json:"taksiran_persalinan,omitempty"`
	// UKKehamilanSaatIni       int32   `json:"uk_kehamilan_saat_ini,omitempty"`
	JarakKehamilanSebelumnya int32 `json:"jarak_kehamilan_sebelumnya,omitempty"`
	// StatusKehamilan          string  `json:"status_kehamilan,omitempty"`
	BB_Awal float64 `json:"bb_awal,omitempty"`
	TB      float64 `json:"tb,omitempty"`
}

func (c *KehamilanController) Create(ctx echo.Context) error {
	var req createKehamilanRequest
	if err := ctx.Bind(&req); err != nil {
		return helpers.Response(ctx, http.StatusBadRequest, []string{"format request tidak valid"})
	}
	if req.IbuID == 0 {
		return helpers.Response(ctx, http.StatusBadRequest, []string{"ibu_id wajib diisi"})
	}
	kehamilan := &models.Kehamilan{
		IbuID: req.IbuID,
		// Gravida:                  req.Gravida,
		// Paritas:                  req.Paritas,
		// Abortus:                  req.Abortus,
		// UKKehamilanSaatIni:       req.UKKehamilanSaatIni,
		JarakKehamilanSebelumnya: req.JarakKehamilanSebelumnya,
		// StatusKehamilan:          req.StatusKehamilan,
		BB_Awal: req.BB_Awal,
		TB:      req.TB,
	}
	if req.HPHT != "" {
		if t, err := time.Parse("2006-01-02", req.HPHT); err == nil {
			kehamilan.HPHT = t
		}
	}
	if req.TaksiranPersalinan != "" {
		if t, err := time.Parse("2006-01-02", req.TaksiranPersalinan); err == nil {
			kehamilan.TaksiranPersalinan = t
		}
	}
	if err := c.usecase.Create(kehamilan); err != nil {
		statusCode := customerror.GetStatusCode(err)
		return helpers.Response(ctx, statusCode, []string{err.Error()})
	}
	return helpers.StandardResponse(ctx, http.StatusCreated, []string{"Kehamilan berhasil ditambahkan"}, kehamilan, nil)
}

func (c *KehamilanController) GetByID(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return helpers.Response(ctx, http.StatusBadRequest, []string{"id tidak valid"})
	}
	data, err := c.usecase.GetByID(int32(id))
	if err != nil {
		statusCode := customerror.GetStatusCode(err)
		return helpers.Response(ctx, statusCode, []string{err.Error()})
	}
	return helpers.StandardResponse(ctx, http.StatusOK, []string{"Data ditemukan"}, data, nil)
}

func (c *KehamilanController) GetByIbuID(ctx echo.Context) error {
	ibuID, err := strconv.ParseInt(ctx.QueryParam("ibu_id"), 10, 32)
	if err != nil || ibuID <= 0 {
		return helpers.Response(ctx, http.StatusBadRequest, []string{"ibu_id wajib diisi dan harus angka"})
	}
	list, err := c.usecase.GetByIbuID(int32(ibuID))
	if err != nil {
		statusCode := customerror.GetStatusCode(err)
		return helpers.Response(ctx, statusCode, []string{err.Error()})
	}
	return helpers.StandardResponse(ctx, http.StatusOK, []string{"Data ditemukan"}, list, nil)
}

func (c *KehamilanController) GetAll(ctx echo.Context) error {
	list, err := c.usecase.GetAll()
	if err != nil {
		statusCode := customerror.GetStatusCode(err)
		return helpers.Response(ctx, statusCode, []string{err.Error()})
	}
	return helpers.StandardResponse(ctx, http.StatusOK, []string{"Data ditemukan"}, list, nil)
}

func (c *KehamilanController) Update(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return helpers.Response(ctx, http.StatusBadRequest, []string{"id tidak valid"})
	}
	var req createKehamilanRequest
	if err := ctx.Bind(&req); err != nil {
		return helpers.Response(ctx, http.StatusBadRequest, []string{"format request tidak valid"})
	}
	existing, err := c.usecase.GetByID(int32(id))
	if err != nil {
		return helpers.Response(ctx, http.StatusNotFound, []string{"Kehamilan tidak ditemukan"})
	}
	// Update hanya field yang dikirim (bukan zero value)
	if req.IbuID != 0 {
		existing.IbuID = req.IbuID
	}
	// if req.Gravida != 0 {
	// 	existing.Gravida = req.Gravida
	// }
	// if req.Paritas != 0 {
	// 	existing.Paritas = req.Paritas
	// }
	// if req.Abortus != 0 {
	// 	existing.Abortus = req.Abortus
	// }
	if req.HPHT != "" {
		if t, err := time.Parse("2006-01-02", req.HPHT); err == nil {
			existing.HPHT = t
		}
	}
	if req.TaksiranPersalinan != "" {
		if t, err := time.Parse("2006-01-02", req.TaksiranPersalinan); err == nil {
			existing.TaksiranPersalinan = t
		}
	}
	// if req.UKKehamilanSaatIni != 0 {
	// 	existing.UKKehamilanSaatIni = req.UKKehamilanSaatIni
	// }
	if req.JarakKehamilanSebelumnya != 0 {
		existing.JarakKehamilanSebelumnya = req.JarakKehamilanSebelumnya
	}
	// if req.StatusKehamilan != "" {
	// 	existing.StatusKehamilan = req.StatusKehamilan
	// }
	if req.BB_Awal > 0 {
		existing.BB_Awal = req.BB_Awal
	}
	if req.TB > 0 {
		existing.TB = req.TB
	}
	if err := c.usecase.Update(existing); err != nil {
		statusCode := customerror.GetStatusCode(err)
		return helpers.Response(ctx, statusCode, []string{err.Error()})
	}
	return helpers.StandardResponse(ctx, http.StatusOK, []string{"Kehamilan berhasil diperbarui"}, existing, nil)
}

func (c *KehamilanController) Delete(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return helpers.Response(ctx, http.StatusBadRequest, []string{"id tidak valid"})
	}
	if err := c.usecase.Delete(int32(id)); err != nil {
		statusCode := customerror.GetStatusCode(err)
		return helpers.Response(ctx, statusCode, []string{err.Error()})
	}
	return helpers.StandardResponse(ctx, http.StatusOK, []string{"Kehamilan berhasil dihapus"}, nil, nil)
}

// func (c *KehamilanController) GetDashboard(ctx echo.Context) error {
// 	data, err := c.usecase.GetDashboardIbuHamil()
// 	if err != nil {
// 		statusCode := customerror.GetStatusCode(err)
// 		return helpers.Response(ctx, statusCode, []string{err.Error()})
// 	}

// 	return helpers.StandardResponse(ctx, http.StatusOK, []string{"Dashboard ibu hamil"}, data, nil)
// }

func (c *KehamilanController) GetDashboard(ctx echo.Context) error {
    ibuID, err := strconv.ParseInt(ctx.QueryParam("ibu_id"), 10, 32)
    if err != nil || ibuID <= 0 {
        return helpers.Response(ctx, http.StatusBadRequest, []string{"ibu_id wajib diisi dan harus angka"})
    }

    data, err := c.usecase.GetDashboardIbuHamil(int32(ibuID))
    if err != nil {
        statusCode := customerror.GetStatusCode(err)
        return helpers.Response(ctx, statusCode, []string{err.Error()})
    }

    return helpers.StandardResponse(ctx, http.StatusOK, []string{"Dashboard ibu hamil"}, data, nil)
}

func (c *KehamilanController) CheckActiveByIbuID(ctx echo.Context) error {
	ibuIDParam := ctx.Param("ibu_id")

	ibuID, err := strconv.ParseInt(ibuIDParam, 10, 32)
	if err != nil {
		return helpers.Response(ctx, http.StatusBadRequest, []string{"ID ibu tidak valid"})
	}

	isActive, err := c.usecase.ExistsActiveByIbuID(int32(ibuID))
	if err != nil {
		statusCode := customerror.GetStatusCode(err)
		return helpers.Response(ctx, statusCode, []string{"Gagal melakukan pengecekan"})
	}

	message := "Ibu siap didaftarkan"
	if isActive {
		message = "Ibu ini masih dalam masa kehamilan, silakan lanjutkan data yang sudah ada"
	}

	return helpers.StandardResponse(ctx, http.StatusOK, []string{message}, map[string]bool{
		"is_active": isActive,
	}, nil)
}

// MODUL IBU
// func (c *KehamilanController) GetAktifForOrangtua(ctx echo.Context) error {
// 	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
// 	if !ok || claims == nil {
// 		return ctx.JSON(http.StatusUnauthorized, models.Response{
// 			StatusCode: http.StatusUnauthorized,
// 			Message:    "Unauthorized",
// 		})
// 	}

// 	data, err := c.usecase.GetAktifByPhoneNumber(claims.PhoneNumber)
// 	if err != nil {
// 		return ctx.JSON(http.StatusNotFound, models.Response{
// 			StatusCode: http.StatusNotFound,
// 			Message:    "Data kehamilan aktif tidak ditemukan",
// 		})
// 	}

// 	return ctx.JSON(http.StatusOK, models.Response{
// 		StatusCode: http.StatusOK,
// 		Data:       data,
// 	})
// }

func (c *KehamilanController) GetAktifForOrangtua(ctx echo.Context) error {
	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
	if !ok || claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Unauthorized",
		})
	}

	data, err := c.usecase.GetAktifByUserID(claims.UserID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{
			StatusCode: http.StatusNotFound,
			Message:    "Data kehamilan aktif tidak ditemukan",
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Data:       data,
	})
}
