package controllers

import (
	"net/http"
	"strconv"

	"monitoring-service/app/constants"
	"monitoring-service/app/helpers"
	"monitoring-service/app/usecases"
	"monitoring-service/pkg/customerror"

	"github.com/labstack/echo/v4"
)

// ==================== POSYANDU MANAGEMENT ====================

// BidanCreatePosyandu - Bidan membuat posyandu baru
func (m *Main) BidanCreatePosyandu(c echo.Context) error {
	var req usecases.AdminCreatePosyanduRequest
	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid: " + err.Error()})
	}

	// Explicit validation
	if req.IDPuskesmas == 0 {
		return helpers.Response(c, http.StatusBadRequest, []string{"id_puskesmas wajib diisi dan tidak boleh 0"})
	}
	if req.Nama == "" {
		return helpers.Response(c, http.StatusBadRequest, []string{"nama wajib diisi"})
	}

	// TODO: Get posyandu_id dari bidan yang sedang login
	// Validate bahwa posyandu desa sesuai dengan desa bidan

	if err := m.usecases.AdminTenagaKesehatan.CreatePosyandu(&req); err != nil {
		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusCreated, []string{constants.SUCCESS_RESPONSE_MESSAGE}, map[string]string{"status": "created"}, nil)
}

// BidanListPosyandu - Bidan melihat list posyandu
func (m *Main) BidanListPosyandu(c echo.Context) error {
	search := c.QueryParam("search")

	data, err := m.usecases.AdminTenagaKesehatan.ListPosyandu(search)
	if err != nil {
		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

// BidanListPenduduk - Bidan melihat list penduduk untuk dropdown
func (m *Main) BidanListPenduduk(c echo.Context) error {
	search := c.QueryParam("search")

	data, err := m.usecases.AdminTenagaKesehatan.ListPenduduk(search)
	if err != nil {
		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

// BidanGetPosyanduDetail - Bidan melihat detail posyandu
func (m *Main) BidanGetPosyanduDetail(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"id tidak valid"})
	}

	data, err := m.usecases.AdminTenagaKesehatan.GetPosyanduDetail(int32(id))
	if err != nil {
		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

// BidanUpdatePosyandu - Bidan update posyandu
func (m *Main) BidanUpdatePosyandu(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"id tidak valid"})
	}

	var req usecases.AdminUpdatePosyanduRequest
	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
	}

	data, err := m.usecases.AdminTenagaKesehatan.UpdatePosyandu(int32(id), &req)
	if err != nil {
		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

// ==================== BIDAN MANAGEMENT ====================

// BidanCreateBidan - Bidan membuat bidan baru di posyandu mereka
func (m *Main) BidanCreateBidan(c echo.Context) error {
	var req usecases.AdminCreateBidanRequest
	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
	}

	// TODO: Get posyandu_id dari bidan yang sedang login
	// Validate bahwa posyandu_id sesuai dengan posyandu bidan

	data, err := m.usecases.AdminTenagaKesehatan.CreateBidan(&req)
	if err != nil {
		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusCreated, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

// BidanListBidan - Bidan melihat list bidan di posyandu mereka
func (m *Main) BidanListBidan(c echo.Context) error {
	desa := c.QueryParam("desa")

	data, err := m.usecases.AdminTenagaKesehatan.ListBidan(desa)
	if err != nil {
		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

// BidanGetBidanDetail - Bidan melihat detail bidan lain
func (m *Main) BidanGetBidanDetail(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"id bidan tidak valid"})
	}

	data, err := m.usecases.AdminTenagaKesehatan.GetBidanDetail(int32(id))
	if err != nil {
		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

// BidanUpdateBidan - Bidan update bidan di posyandu mereka
func (m *Main) BidanUpdateBidan(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"id bidan tidak valid"})
	}

	var req usecases.AdminUpdateBidanRequest
	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
	}

	data, err := m.usecases.AdminTenagaKesehatan.UpdateBidan(int32(id), &req)
	if err != nil {
		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

// BidanUpdateBidanStatus - Bidan update status bidan
func (m *Main) BidanUpdateBidanStatus(c echo.Context) error {
	idRaw, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"id bidan tidak valid"})
	}

	var req struct {
		Status string `json:"status"`
	}
	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
	}

	// TODO: Validate ownership - hanya bisa update status bidan di posyandu mereka

	if err := m.usecases.AdminTenagaKesehatan.SetStatusBidan(int32(idRaw), req.Status); err != nil {
		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, map[string]string{"status": "updated"}, nil)
}

// ==================== KADER MANAGEMENT ====================

// BidanCreateKader - Bidan membuat kader baru di posyandu mereka
func (m *Main) BidanCreateKader(c echo.Context) error {
	var req usecases.AdminCreateKaderRequest
	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
	}

	// TODO: Get posyandu_id dari bidan yang sedang login
	// Set posyandu_id ke kader jika belum ada

	data, err := m.usecases.AdminTenagaKesehatan.CreateKader(&req)
	if err != nil {
		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusCreated, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

// BidanListKader - Bidan melihat list kader di posyandu mereka
func (m *Main) BidanListKader(c echo.Context) error {
	// TODO: Get posyandu_id dari bidan yang sedang login
	// Filter kader hanya di posyandu mereka

	desa := c.QueryParam("desa")

	data, err := m.usecases.AdminTenagaKesehatan.ListKader(desa)
	if err != nil {
		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

// BidanGetKaderDetail - Bidan melihat detail kader
func (m *Main) BidanGetKaderDetail(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"id kader tidak valid"})
	}

	data, err := m.usecases.AdminTenagaKesehatan.GetKaderDetail(int32(id))
	if err != nil {
		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

// BidanUpdateKader - Bidan update kader di posyandu mereka
func (m *Main) BidanUpdateKader(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"id kader tidak valid"})
	}

	var req usecases.AdminUpdateKaderRequest
	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
	}

	data, err := m.usecases.AdminTenagaKesehatan.UpdateKader(int32(id), &req)
	if err != nil {
		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

// BidanUpdateKaderStatus - Bidan update status kader
func (m *Main) BidanUpdateKaderStatus(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"id kader tidak valid"})
	}

	var req struct {
		Status string `json:"status"`
	}
	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
	}

	if err := m.usecases.AdminTenagaKesehatan.SetStatusKader(int32(id), req.Status); err != nil {
		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, map[string]string{"status": "updated"}, nil)
}
