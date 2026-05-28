package controllers

import (
	"net/http"
	"strconv"

	"monitoring-service/app/helpers"
	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

// SkriningPemantauanController menangani endpoint CRUD skrining pemantauan tanda bahaya
type SkriningPemantauanController struct {
	skriningPemantauanUC usecases.SkriningPemantauanUsecase
}

func NewSkriningPemantauanController(skriningPemantauanUC usecases.SkriningPemantauanUsecase) *SkriningPemantauanController {
	return &SkriningPemantauanController{skriningPemantauanUC: skriningPemantauanUC}
}

// GetAll godoc
// @Summary      Dapatkan semua skrining pemantauan
// @Tags         skrining-pemantauan
// @Security     BearerAuth
// @Produce      json
// @Success      200  {object}  models.Response
// @Failure      500  {object}  models.Response
// @Router       /skrining-pemantauan [get]
func (h *SkriningPemantauanController) GetAll(c echo.Context) error {
	list, err := h.skriningPemantauanUC.GetAll()
	if err != nil {
		return helpers.StandardResponse(c, http.StatusInternalServerError, err.Error(), nil, nil)
	}
	return helpers.StandardResponse(c, http.StatusOK, "Berhasil", list, nil)
}

// Create godoc
// @Summary      Tambah skrining pemantauan baru
// @Tags         skrining-pemantauan
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body  body      models.SkriningPemantauan  true  "Data skrining pemantauan"
// @Success      201   {object}  models.Response
// @Failure      400   {object}  models.Response
// @Router       /skrining-pemantauan [post]
func (h *SkriningPemantauanController) Create(c echo.Context) error {
	var req models.SkriningPemantauan
	if err := c.Bind(&req); err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, "request tidak valid: "+err.Error(), nil, nil)
	}

	if err := h.skriningPemantauanUC.Create(&req); err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, err.Error(), nil, nil)
	}

	return helpers.StandardResponse(c, http.StatusCreated, "data skrining pemantauan berhasil ditambahkan", req, nil)
}

// Detail godoc
// @Summary      Detail skrining pemantauan berdasarkan ID
// @Tags         skrining-pemantauan
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      string  true  "Skrining Pemantauan ID"
// @Success      200  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Router       /skrining-pemantauan/{id} [get]
func (h *SkriningPemantauanController) Detail(c echo.Context) error {
	id64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, "id tidak valid", nil, nil)
	}

	skrining, err := h.skriningPemantauanUC.GetByID(int32(id64))
	if err != nil {
		return helpers.StandardResponse(c, http.StatusNotFound, err.Error(), nil, nil)
	}

	return helpers.StandardResponse(c, http.StatusOK, "berhasil", skrining, nil)
}

// GetByAnakID godoc
// @Summary      Dapatkan skrining pemantauan berdasarkan Anak ID
// @Tags         skrining-pemantauan
// @Security     BearerAuth
// @Produce      json
// @Param        anak_id   path      string  true  "Anak ID"
// @Success      200       {object}  models.Response
// @Failure      404       {object}  models.Response
// @Router       /skrining-pemantauan/anak/{anak_id} [get]
func (h *SkriningPemantauanController) GetByAnakID(c echo.Context) error {
	anakID64, err := strconv.ParseInt(c.Param("anak_id"), 10, 32)
	if err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, "anak_id tidak valid", nil, nil)
	}

	list, err := h.skriningPemantauanUC.GetByAnakID(int32(anakID64))
	if err != nil {
		return helpers.StandardResponse(c, http.StatusNotFound, err.Error(), nil, nil)
	}

	return helpers.StandardResponse(c, http.StatusOK, "berhasil", list, nil)
}

// Update godoc
// @Summary      Update data skrining pemantauan
// @Tags         skrining-pemantauan
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id    path      string                   true  "Skrining Pemantauan ID"
// @Param        body  body      models.SkriningPemantauan true  "Data skrining pemantauan yang diperbarui"
// @Success      200   {object}  models.Response
// @Failure      404   {object}  models.Response
// @Router       /skrining-pemantauan/{id} [put]
func (h *SkriningPemantauanController) Update(c echo.Context) error {
	idParam := c.Param("id")

	id64, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, "id tidak valid", nil, nil)
	}

	var req models.SkriningPemantauan
	if err := c.Bind(&req); err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, "request tidak valid: "+err.Error(), nil, nil)
	}

	req.ID = int32(id64)

	if err := h.skriningPemantauanUC.Update(&req); err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, err.Error(), nil, nil)
	}

	return helpers.StandardResponse(c, http.StatusOK, "data skrining pemantauan berhasil diperbarui", req, nil)
}

// Delete godoc
// @Summary      Hapus data skrining pemantauan
// @Tags         skrining-pemantauan
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      string  true  "Skrining Pemantauan ID"
// @Success      200  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Router       /skrining-pemantauan/{id} [delete]
func (h *SkriningPemantauanController) Delete(c echo.Context) error {
	idParam := c.Param("id")

	id64, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, "id tidak valid", nil, nil)
	}

	if err := h.skriningPemantauanUC.Delete(int32(id64)); err != nil {
		return helpers.StandardResponse(c, http.StatusNotFound, err.Error(), nil, nil)
	}

	return helpers.StandardResponse(c, http.StatusOK, "data skrining pemantauan berhasil dihapus", nil, nil)
}
