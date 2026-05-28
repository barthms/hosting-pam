package controllers

import (
	"net/http"
	"strconv"

	"monitoring-service/app/helpers"
	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

// AnakController menangani endpoint CRUD data anak.
type AnakController struct {
	anakUC *usecases.AnakUseCase
}

func NewAnakController(anakUC *usecases.AnakUseCase) *AnakController {
	return &AnakController{anakUC: anakUC}
}

// Create godoc
// @Summary      Tambah anak baru
// @Tags         anak
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body  body      models.CreateAnakRequest  true  "Data anak"
// @Success      201   {object}  models.Response
// @Failure      400   {object}  models.Response
// @Router       /anak [post]

func (h *AnakController) AdminList(c echo.Context) error {
	kehamilanParam := c.QueryParam("kehamilan_id")

	var kehamilanID int32 = 0

	if kehamilanParam != "" {
		id64, err := strconv.ParseInt(kehamilanParam, 10, 32)
		if err != nil {
			return helpers.StandardResponse(c, http.StatusBadRequest, "kehamilan_id tidak valid", nil, nil)
		}
		kehamilanID = int32(id64)
	}

	list, err := h.anakUC.AdminListAnak(kehamilanID)
	if err != nil {
		return helpers.StandardResponse(c, http.StatusInternalServerError, err.Error(), nil, nil)
	}
	return helpers.StandardResponse(c, http.StatusOK, "Berhasil", list, nil)
}

func (h *AnakController) Create(c echo.Context) error {

	var req models.CreateAnakRequest
	if err := c.Bind(&req); err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, "request tidak valid: "+err.Error(), nil, nil)
	}

	anak, err := h.anakUC.CreateAnak(req)
	if err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, err.Error(), nil, nil)
	}

	return helpers.StandardResponse(c, http.StatusCreated, "data anak berhasil ditambahkan", anak, nil)
}

// CreateDenganPenduduk: create anak + auto-create kependudukan
func (h *AnakController) CreateDenganPenduduk(c echo.Context) error {

	var req models.CreateAnakDenganPendudukRequest
	if err := c.Bind(&req); err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, "request tidak valid: "+err.Error(), nil, nil)
	}

	anak, err := h.anakUC.CreateAnakDenganPenduduk(req)
	if err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, err.Error(), nil, nil)
	}

	return helpers.StandardResponse(c, http.StatusCreated, "data anak berhasil ditambahkan", anak, nil)
}

// Detail godoc
// @Summary      Detail anak berdasarkan ID
// @Tags         anak
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      string  true  "Anak ID"
// @Success      200  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Router       /anak/{id} [get]
func (h *AnakController) Detail(c echo.Context) error {
	id64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, "id tidak valid", nil, nil)
	}

	anak, err := h.anakUC.GetAnak(int32(id64))
	if err != nil {
		return helpers.StandardResponse(c, http.StatusNotFound, err.Error(), nil, nil)
	}

	return helpers.StandardResponse(c, http.StatusOK, "berhasil", anak, nil)
}

// Update godoc
// @Summary      Update data anak
// @Tags         anak
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id    path      string                   true  "Anak ID"
// @Param        body  body      models.UpdateAnakRequest true  "Data anak yang diperbarui"
// @Success      200   {object}  models.Response
// @Failure      404   {object}  models.Response
// @Router       /anak/{id} [put]
func (h *AnakController) Update(c echo.Context) error {
	idParam := c.Param("id")

	id64, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, "id tidak valid", nil, nil)
	}

	var req models.UpdateAnakRequest
	if err := c.Bind(&req); err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, "request tidak valid: "+err.Error(), nil, nil)
	}

	anak, err := h.anakUC.UpdateAnak(int32(id64), req)
	if err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, err.Error(), nil, nil)
	}

	return helpers.StandardResponse(c, http.StatusOK, "data anak berhasil diperbarui", anak, nil)
}

// Delete godoc
// @Summary      Hapus data anak
// @Tags         anak
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      string  true  "Anak ID"
// @Success      200  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Router       /anak/{id} [delete]
func (h *AnakController) Delete(c echo.Context) error {
	idParam := c.Param("id")

	id64, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, "id tidak valid", nil, nil)
	}

	if err := h.anakUC.DeleteAnak(int32(id64)); err != nil {
		return helpers.StandardResponse(c, http.StatusNotFound, err.Error(), nil, nil)
	}

	return helpers.StandardResponse(c, http.StatusOK, "data anak berhasil dihapus", nil, nil)
}
