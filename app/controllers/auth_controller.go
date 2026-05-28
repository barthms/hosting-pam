package controllers

import (
	"net/http"
	"strconv"

	"monitoring-service/app/constants"
	"monitoring-service/app/helpers"
	"monitoring-service/app/models"
	"monitoring-service/app/usecases"
	"monitoring-service/pkg/customerror"

	"github.com/labstack/echo/v4"
)

// func (m *Main) RegisterOrangTua(c echo.Context) error {
// 	var req usecases.RegisterOrangTuaRequest
// 	if err := c.Bind(&req); err != nil {
// 		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
// 	}
// 	if err := m.usecases.RegisterOrangTua.Register(&req); err != nil {
// 		statusCode := customerror.GetStatusCode(err)
// 		return helpers.Response(c, statusCode, []string{err.Error()})
// 	}
// 	return helpers.StandardResponse(c, http.StatusCreated, []string{constants.SUCCESS_RESPONSE_MESSAGE}, map[string]string{
// 		"message": "registrasi orang tua berhasil",
// 	}, nil)
// }

func (m *Main) Register(c echo.Context) error {
	var req models.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
	}

	if err := m.usecases.Register(&req); err != nil {
		statusCode := customerror.GetStatusCode(err)
		return helpers.Response(c, statusCode, []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusCreated, []string{constants.SUCCESS_RESPONSE_MESSAGE}, map[string]string{
		"message": "registrasi berhasil",
	}, nil)
}

func (m *Main) Login(c echo.Context) error {
	var req models.LoginRequest
	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
	}

	data, err := m.usecases.Login(&req)
	if err != nil {
		statusCode := customerror.GetStatusCode(err)
		return helpers.Response(c, statusCode, []string{err.Error()})
	}
	if req.FcmToken != "" {

		tokenReq := &models.TokenRequest{
			PenggunaID: uint(data.UserID),
			FcmToken:   req.FcmToken,
		}
		_ = m.usecases.SaveFCMToken(tokenReq)
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

func (m *Main) AdminCreateKartuKeluarga(c echo.Context) error {
	var req usecases.AdminCreateKartuKeluargaRequest
	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
	}

	data, err := m.usecases.AdminAkunKeluarga.CreateKartuKeluarga(&req)
	if err != nil {
		statusCode := customerror.GetStatusCode(err)
		return helpers.Response(c, statusCode, []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusCreated, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

func (m *Main) AdminListKartuKeluarga(c echo.Context) error {
	search := c.QueryParam("search")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	sortBy := c.QueryParam("sort_by")
	sortDir := c.QueryParam("sort_dir")

	data, err := m.usecases.AdminAkunKeluarga.ListKartuKeluarga(search, page, limit, sortBy, sortDir)
	if err != nil {
		statusCode := customerror.GetStatusCode(err)
		return helpers.Response(c, statusCode, []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

func (m *Main) AdminDetailKartuKeluarga(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("kartu_keluarga_id"), 10, 64)
	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"kartu_keluarga_id tidak valid"})
	}

	data, detailErr := m.usecases.AdminAkunKeluarga.DetailKartuKeluarga(id)
	if detailErr != nil {
		statusCode := customerror.GetStatusCode(detailErr)
		return helpers.Response(c, statusCode, []string{detailErr.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

func (m *Main) AdminUpdateKartuKeluarga(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("kartu_keluarga_id"), 10, 64)
	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"kartu_keluarga_id tidak valid"})
	}

	var req usecases.AdminUpdateKartuKeluargaRequest
	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
	}

	data, updateErr := m.usecases.AdminAkunKeluarga.UpdateKartuKeluarga(id, &req)
	if updateErr != nil {
		statusCode := customerror.GetStatusCode(updateErr)
		return helpers.Response(c, statusCode, []string{updateErr.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

func (m *Main) AdminUpdateAnggotaKeluarga(c echo.Context) error {
	kkID, err := strconv.ParseInt(c.Param("kartu_keluarga_id"), 10, 64)
	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"kartu_keluarga_id tidak valid"})
	}

	pendudukID64, err := strconv.ParseInt(c.Param("penduduk_id"), 10, 32)
	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"penduduk_id tidak valid"})
	}

	var req usecases.AdminAnggotaKeluargaRequest
	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
	}

	data, updateErr := m.usecases.AdminAkunKeluarga.UpdateAnggotaKeluarga(kkID, int32(pendudukID64), &req)
	if updateErr != nil {
		statusCode := customerror.GetStatusCode(updateErr)
		return helpers.Response(c, statusCode, []string{updateErr.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

func (m *Main) AdminAddAnggotaKeluarga(c echo.Context) error {
	kkID, err := strconv.ParseInt(c.Param("kartu_keluarga_id"), 10, 64)
	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"kartu_keluarga_id tidak valid"})
	}

	var req usecases.AdminAnggotaKeluargaRequest
	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
	}

	data, addErr := m.usecases.AdminAkunKeluarga.AddAnggotaKeluarga(kkID, &req)
	if addErr != nil {
		statusCode := customerror.GetStatusCode(addErr)
		return helpers.Response(c, statusCode, []string{addErr.Error()})
	}

	return helpers.StandardResponse(c, http.StatusCreated, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

func (m *Main) AdminDeleteAnggotaKeluarga(c echo.Context) error {
	kkID, err := strconv.ParseInt(c.Param("kartu_keluarga_id"), 10, 64)
	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"kartu_keluarga_id tidak valid"})
	}

	pendudukID64, err := strconv.ParseInt(c.Param("penduduk_id"), 10, 32)
	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"penduduk_id tidak valid"})
	}

	if deleteErr := m.usecases.AdminAkunKeluarga.DeleteAnggotaKeluarga(kkID, int32(pendudukID64)); deleteErr != nil {
		statusCode := customerror.GetStatusCode(deleteErr)
		return helpers.Response(c, statusCode, []string{deleteErr.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, map[string]bool{"deleted": true}, nil)
}

func (m *Main) AdminDeleteKartuKeluarga(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("kartu_keluarga_id"), 10, 64)
	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"kartu_keluarga_id tidak valid"})
	}

	if deleteErr := m.usecases.AdminAkunKeluarga.DeleteKartuKeluarga(id); deleteErr != nil {
		statusCode := customerror.GetStatusCode(deleteErr)
		return helpers.Response(c, statusCode, []string{deleteErr.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, map[string]bool{"deleted": true}, nil)
}
