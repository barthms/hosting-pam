package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"monitoring-service/app/constants"
	"monitoring-service/app/helpers"
	"monitoring-service/app/models"
	"monitoring-service/pkg/customerror"
)

// ─────────────────────────────────────────────────────────
// PERAWATAN ENDPOINTS
// ─────────────────────────────────────────────────────────

// CreatePerawatan handles POST /perawatan to create a new perawatan record
// @Summary Create a new perawatan
// @Description Create a new perawatan milestone/achievement record for a child
// @Tags Perawatan
// @Accept json
// @Produce json
// @Param request body models.CreatePerawatanRequest true "Perawatan request"
// @Router /perawatan [post]
func (m *Main) CreatePerawatan(c echo.Context) error {
	var req models.CreatePerawatanRequest
	if err := c.Bind(&req); err != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
	}

	// Get auth claims for access control
	var data *models.Perawatan
	var usecaseErr error

	claims, ok := c.Get("auth_claims").(*models.AuthClaims)
	if ok && claims != nil && claims.Role == "ibu" {
		// For ibu, use access control method
		data, usecaseErr = m.usecases.Perawatan.CreatePerawatanForIbu(req, claims.UserID)
	} else {
		// For nakes, use regular method
		data, usecaseErr = m.usecases.Perawatan.CreatePerawatan(req)
	}

	if usecaseErr != nil {
		return helpers.Response(c, customerror.GetStatusCode(usecaseErr), []string{usecaseErr.Error()})
	}

	return helpers.StandardResponse(
		c,
		http.StatusCreated,
		[]string{constants.SUCCESS_RESPONSE_MESSAGE},
		data,
		nil,
	)
}

// GetPerawatanByID handles GET /perawatan/:id to retrieve a single perawatan
// @Summary Get perawatan by ID
// @Description Retrieve a specific perawatan record by ID
// @Tags Perawatan
// @Produce json
// @Param id path int true "Perawatan ID"
// @Router /perawatan/:id [get]
func (m *Main) GetPerawatanByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		return helpers.Response(c, http.StatusBadRequest, []string{"id tidak valid"})
	}

	data, usecaseErr := m.usecases.Perawatan.GetPerawatanByID(uint(id))
	if usecaseErr != nil {
		return helpers.Response(c, customerror.GetStatusCode(usecaseErr), []string{usecaseErr.Error()})
	}

	return helpers.StandardResponse(
		c,
		http.StatusOK,
		[]string{constants.SUCCESS_RESPONSE_MESSAGE},
		data,
		nil,
	)
}

// GetPerawatanByAnakID handles GET /perawatan/anak/:anak_id to retrieve all perawatan for a child
// @Summary Get all perawatan for a child
// @Description Retrieve all perawatan records for a specific child
// @Tags Perawatan
// @Produce json
// @Param anak_id path int true "Child ID"
// @Router /perawatan/anak/:anak_id [get]
func (m *Main) GetPerawatanByAnakID(c echo.Context) error {
	anakID, err := strconv.Atoi(c.Param("anak_id"))
	if err != nil || anakID <= 0 {
		return helpers.Response(c, http.StatusBadRequest, []string{"anak_id tidak valid"})
	}

	// Get auth claims for access control
	var data interface{}
	var usecaseErr error

	claims, ok := c.Get("auth_claims").(*models.AuthClaims)
	if ok && claims != nil && claims.Role == "ibu" {
		// For ibu, use access control method
		data, usecaseErr = m.usecases.Perawatan.GetPerawatanByAnakIDForIbu(int32(anakID), claims.UserID)
	} else {
		// For nakes, use regular method
		data, usecaseErr = m.usecases.Perawatan.GetPerawatanByAnakID(int32(anakID))
	}

	if usecaseErr != nil {
		return helpers.Response(c, customerror.GetStatusCode(usecaseErr), []string{usecaseErr.Error()})
	}

	return helpers.StandardResponse(
		c,
		http.StatusOK,
		[]string{constants.SUCCESS_RESPONSE_MESSAGE},
		data,
		nil,
	)
}

// GetPerawatanByAnakIDAndRentangUsia handles GET /perawatan/anak/:anak_id/rentang-usia/:rentang_usia
// @Summary Get perawatan for a child by age range
// @Description Retrieve perawatan records for a specific child within a specific age range
// @Tags Perawatan
// @Produce json
// @Param anak_id path int true "Child ID"
// @Param rentang_usia path string true "Age range (e.g., 0-12 Bulan, 1-2 Tahun)"
// @Router /perawatan/anak/:anak_id/rentang-usia/:rentang_usia [get]
func (m *Main) GetPerawatanByAnakIDAndRentangUsia(c echo.Context) error {
	anakID, err := strconv.Atoi(c.Param("anak_id"))
	if err != nil || anakID <= 0 {
		return helpers.Response(c, http.StatusBadRequest, []string{"anak_id tidak valid"})
	}

	rentangUsia := c.Param("rentang_usia")
	if rentangUsia == "" {
		return helpers.Response(c, http.StatusBadRequest, []string{"rentang_usia tidak boleh kosong"})
	}

	// Get auth claims for access control
	var data interface{}
	var usecaseErr error

	claims, ok := c.Get("auth_claims").(*models.AuthClaims)
	if ok && claims != nil && claims.Role == "ibu" {
		// For ibu, use access control method
		data, usecaseErr = m.usecases.Perawatan.GetPerawatanByAnakIDAndRentangUsiaForIbu(int32(anakID), rentangUsia, claims.UserID)
	} else {
		// For nakes, use regular method
		data, usecaseErr = m.usecases.Perawatan.GetPerawatanByAnakIDAndRentangUsia(int32(anakID), rentangUsia)
	}

	if usecaseErr != nil {
		return helpers.Response(c, customerror.GetStatusCode(usecaseErr), []string{usecaseErr.Error()})
	}

	return helpers.StandardResponse(
		c,
		http.StatusOK,
		[]string{constants.SUCCESS_RESPONSE_MESSAGE},
		data,
		nil,
	)
}

// UpdatePerawatan handles PUT /perawatan/:id to update an existing perawatan
// @Summary Update perawatan
// @Description Update an existing perawatan record
// @Tags Perawatan
// @Accept json
// @Produce json
// @Param id path int true "Perawatan ID"
// @Param request body models.UpdatePerawatanRequest true "Perawatan update request"
// @Router /perawatan/:id [put]
func (m *Main) UpdatePerawatan(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		return helpers.Response(c, http.StatusBadRequest, []string{"id tidak valid"})
	}

	var req models.UpdatePerawatanRequest
	if bindErr := c.Bind(&req); bindErr != nil {
		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
	}

	// Get auth claims for access control
	var data *models.Perawatan
	var usecaseErr error

	claims, ok := c.Get("auth_claims").(*models.AuthClaims)
	if ok && claims != nil && claims.Role == "ibu" {
		// For ibu, use access control method
		data, usecaseErr = m.usecases.Perawatan.UpdatePerawatanForIbu(uint(id), req, claims.UserID)
	} else {
		// For nakes, use regular method
		data, usecaseErr = m.usecases.Perawatan.UpdatePerawatan(uint(id), req)
	}

	if usecaseErr != nil {
		return helpers.Response(c, customerror.GetStatusCode(usecaseErr), []string{usecaseErr.Error()})
	}

	return helpers.StandardResponse(
		c,
		http.StatusOK,
		[]string{constants.SUCCESS_RESPONSE_MESSAGE},
		data,
		nil,
	)
}

// DeletePerawatan handles DELETE /perawatan/:id to delete a perawatan
// @Summary Delete perawatan
// @Description Delete (soft delete) a perawatan record
// @Tags Perawatan
// @Produce json
// @Param id path int true "Perawatan ID"
// @Router /perawatan/:id [delete]
func (m *Main) DeletePerawatan(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		return helpers.Response(c, http.StatusBadRequest, []string{"id tidak valid"})
	}

	// Get auth claims for access control
	var usecaseErr error

	claims, ok := c.Get("auth_claims").(*models.AuthClaims)
	if ok && claims != nil && claims.Role == "ibu" {
		// For ibu, use access control method
		usecaseErr = m.usecases.Perawatan.DeletePerawatanForIbu(uint(id), claims.UserID)
	} else {
		// For nakes, use regular method
		usecaseErr = m.usecases.Perawatan.DeletePerawatan(uint(id))
	}

	if usecaseErr != nil {
		return helpers.Response(c, customerror.GetStatusCode(usecaseErr), []string{usecaseErr.Error()})
	}

	return helpers.StandardResponse(
		c,
		http.StatusOK,
		[]string{constants.SUCCESS_RESPONSE_MESSAGE},
		map[string]string{"message": "perawatan berhasil dihapus"},
		nil,
	)
}

// ─────────────────────────────────────────────────────────
// KATEGORI CAPAIAN ENDPOINTS
// ─────────────────────────────────────────────────────────

// GetAllKategoriCapaian handles GET /kategori-capaian to retrieve all milestone categories
// @Summary Get all milestone categories
// @Description Retrieve all kategori capaian (milestone categories) sorted by age range
// @Tags KategoriCapaian
// @Produce json
// @Router /kategori-capaian [get]
func (m *Main) GetAllKategoriCapaian(c echo.Context) error {
	data, usecaseErr := m.usecases.Perawatan.GetAllKategoriCapaian()
	if usecaseErr != nil {
		return helpers.Response(c, customerror.GetStatusCode(usecaseErr), []string{usecaseErr.Error()})
	}

	return helpers.StandardResponse(
		c,
		http.StatusOK,
		[]string{constants.SUCCESS_RESPONSE_MESSAGE},
		data,
		nil,
	)
}

// GetKategoriCapaianByRentangUsia handles GET /kategori-capaian/rentang-usia/:rentang_usia
// @Summary Get milestone categories by age range
// @Description Retrieve kategori capaian for a specific age range
// @Tags KategoriCapaian
// @Produce json
// @Param rentang_usia path string true "Age range (e.g., 0-12 Bulan, 1-2 Tahun)"
// @Router /kategori-capaian/rentang-usia/:rentang_usia [get]
func (m *Main) GetKategoriCapaianByRentangUsia(c echo.Context) error {
	rentangUsia := c.Param("rentang_usia")
	if rentangUsia == "" {
		return helpers.Response(c, http.StatusBadRequest, []string{"rentang_usia tidak boleh kosong"})
	}

	data, usecaseErr := m.usecases.Perawatan.GetKategoriCapaianByRentangUsia(rentangUsia)
	if usecaseErr != nil {
		return helpers.Response(c, customerror.GetStatusCode(usecaseErr), []string{usecaseErr.Error()})
	}

	return helpers.StandardResponse(
		c,
		http.StatusOK,
		[]string{constants.SUCCESS_RESPONSE_MESSAGE},
		data,
		nil,
	)
}
