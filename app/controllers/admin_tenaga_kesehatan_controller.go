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

func (m *Main) AdminTambahBidan(c echo.Context) error {
	var req usecases.AdminCreateBidanRequest
	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
	}

	data, err := m.usecases.AdminTenagaKesehatan.CreateBidan(&req)
	if err != nil {
		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusCreated, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

func (m *Main) AdminListBidan(c echo.Context) error {
	desa := c.QueryParam("desa")

	data, err := m.usecases.AdminTenagaKesehatan.ListBidan(desa)
	if err != nil {
		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

func (m *Main) AdminUpdateBidan(c echo.Context) error {
	idRaw, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"id bidan tidak valid"})
	}

	var req usecases.AdminUpdateBidanRequest
	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
	}

	data, updateErr := m.usecases.AdminTenagaKesehatan.UpdateBidan(int32(idRaw), &req)
	if updateErr != nil {
		return helpers.Response(c, customerror.GetStatusCode(updateErr), []string{updateErr.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

func (m *Main) AdminUpdateStatusBidan(c echo.Context) error {
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

	if err := m.usecases.AdminTenagaKesehatan.SetStatusBidan(int32(idRaw), req.Status); err != nil {
		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, map[string]bool{"updated": true}, nil)
}

func (m *Main) AdminTambahKader(c echo.Context) error {
	var req usecases.AdminCreateKaderRequest
	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
	}

	data, err := m.usecases.AdminTenagaKesehatan.CreateKader(&req)
	if err != nil {
		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusCreated, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

func (m *Main) AdminListKader(c echo.Context) error {
	desa := c.QueryParam("desa")

	data, err := m.usecases.AdminTenagaKesehatan.ListKader(desa)
	if err != nil {
		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

func (m *Main) AdminListEligiblePenduduk(c echo.Context) error {
	role := c.QueryParam("role")
	if role == "" {
		return helpers.Response(c, http.StatusBadRequest, []string{"role wajib diisi"})
	}

	data, err := m.usecases.AdminTenagaKesehatan.ListEligiblePenduduk(
		role,
		c.QueryParam("search"),
		c.QueryParam("kecamatan"),
		c.QueryParam("desa"),
	)
	if err != nil {
		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

func (m *Main) AdminTambahPosyandu(c echo.Context) error {
	var req usecases.AdminCreatePosyanduRequest
	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
	}

	if err := m.usecases.AdminTenagaKesehatan.CreatePosyandu(&req); err != nil {
		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusCreated, []string{constants.SUCCESS_RESPONSE_MESSAGE}, map[string]string{"nama": req.Nama}, nil)
}

func (m *Main) AdminListPosyandu(c echo.Context) error {
	data, err := m.usecases.AdminTenagaKesehatan.ListPosyandu(c.QueryParam("search"))
	if err != nil {
		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

func (m *Main) AdminUpdateKader(c echo.Context) error {
	idRaw, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"id kader tidak valid"})
	}

	var req usecases.AdminUpdateKaderRequest
	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
	}

	data, updateErr := m.usecases.AdminTenagaKesehatan.UpdateKader(int32(idRaw), &req)
	if updateErr != nil {
		return helpers.Response(c, customerror.GetStatusCode(updateErr), []string{updateErr.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
}

func (m *Main) AdminUpdateStatusKader(c echo.Context) error {
	idRaw, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"id kader tidak valid"})
	}

	var req struct {
		Status string `json:"status"`
	}
	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
	}

	if err := m.usecases.AdminTenagaKesehatan.SetStatusKader(int32(idRaw), req.Status); err != nil {
		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, map[string]bool{"updated": true}, nil)
}

// AdminCreateAkunBidan dan AdminCreateAkunKader belum diimplementasi di usecase
