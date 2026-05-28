package controllers

import (
	"net/http"
	"strconv"

	"monitoring-service/app/helpers"
	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type InformasiUmumController struct {
	informasiUmumUC usecases.InformasiUmumUsecase
}

func NewInformasiUmumController(informasiUmumUC usecases.InformasiUmumUsecase) *InformasiUmumController {
	return &InformasiUmumController{informasiUmumUC: informasiUmumUC}
}

// GetAll godoc
// @Summary      Dapatkan semua informasi umum
// @Tags         informasi-umum
// @Produce      json
// @Success      200  {object}  models.Response
// @Failure      500  {object}  models.Response
// @Router       /informasi-umum [get]
func (h *InformasiUmumController) GetAll(c echo.Context) error {
	list, err := h.informasiUmumUC.GetAll()
	if err != nil {
		return helpers.StandardResponse(c, http.StatusInternalServerError, err.Error(), nil, nil)
	}
	return helpers.StandardResponse(c, http.StatusOK, "berhasil", list, nil)
}

// Create godoc
// @Summary      Tambah informasi umum baru
// @Tags         informasi-umum
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body  body      models.InformasiUmum  true  "Data informasi umum"
// @Success      201   {object}  models.Response
// @Failure      400   {object}  models.Response
// @Router       /informasi-umum [post]
func (h *InformasiUmumController) Create(c echo.Context) error {
	var req models.InformasiUmum
	if err := c.Bind(&req); err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, "request tidak valid: "+err.Error(), nil, nil)
	}

	if err := h.informasiUmumUC.Create(&req); err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, err.Error(), nil, nil)
	}

	return helpers.StandardResponse(c, http.StatusCreated, "data informasi umum berhasil ditambahkan", req, nil)
}

// Detail godoc
// @Summary      Detail informasi umum berdasarkan ID
// @Tags         informasi-umum
// @Produce      json
// @Param        id   path      string  true  "Informasi Umum ID"
// @Success      200  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Router       /informasi-umum/{id} [get]
func (h *InformasiUmumController) Detail(c echo.Context) error {
	id64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, "id tidak valid", nil, nil)
	}

	data, err := h.informasiUmumUC.GetByID(int32(id64))
	if err != nil {
		return helpers.StandardResponse(c, http.StatusNotFound, err.Error(), nil, nil)
	}

	return helpers.StandardResponse(c, http.StatusOK, "berhasil", data, nil)
}

// Update godoc
// @Summary      Update informasi umum
// @Tags         informasi-umum
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id    path      string                 true  "Informasi Umum ID"
// @Param        body  body      models.InformasiUmum   true  "Data informasi umum yang diperbarui"
// @Success      200   {object}  models.Response
// @Failure      400   {object}  models.Response
// @Router       /informasi-umum/{id} [put]
func (h *InformasiUmumController) Update(c echo.Context) error {
	id64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, "id tidak valid", nil, nil)
	}

	var req models.InformasiUmum
	if err := c.Bind(&req); err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, "request tidak valid: "+err.Error(), nil, nil)
	}

	req.ID = int32(id64)

	if err := h.informasiUmumUC.Update(&req); err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, err.Error(), nil, nil)
	}

	return helpers.StandardResponse(c, http.StatusOK, "data informasi umum berhasil diperbarui", req, nil)
}

// Delete godoc
// @Summary      Hapus informasi umum
// @Tags         informasi-umum
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      string  true  "Informasi Umum ID"
// @Success      200  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Router       /informasi-umum/{id} [delete]
func (h *InformasiUmumController) Delete(c echo.Context) error {
	id64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, "id tidak valid", nil, nil)
	}

	if err := h.informasiUmumUC.Delete(int32(id64)); err != nil {
		return helpers.StandardResponse(c, http.StatusNotFound, err.Error(), nil, nil)
	}

	return helpers.StandardResponse(c, http.StatusOK, "data informasi umum berhasil dihapus", nil, nil)
}
