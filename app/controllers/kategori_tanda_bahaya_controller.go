package controllers

import (
	"net/http"
	"strconv"

	"monitoring-service/app/helpers"
	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

// KategoriTandaBahayaController menangani endpoint CRUD kategori tanda bahaya
type KategoriTandaBahayaController struct {
	kategoriTandaBahayaUC usecases.KategoriTandaBahayaUsecase
}

func NewKategoriTandaBahayaController(kategoriTandaBahayaUC usecases.KategoriTandaBahayaUsecase) *KategoriTandaBahayaController {
	return &KategoriTandaBahayaController{kategoriTandaBahayaUC: kategoriTandaBahayaUC}
}

// GetAll godoc
// @Summary      Dapatkan semua kategori tanda bahaya
// @Tags         kategori-tanda-bahaya
// @Security     BearerAuth
// @Produce      json
// @Success      200  {object}  models.Response
// @Failure      500  {object}  models.Response
// @Router       /kategori-tanda-bahaya [get]
func (h *KategoriTandaBahayaController) GetAll(c echo.Context) error {
	list, err := h.kategoriTandaBahayaUC.GetAll()
	if err != nil {
		return helpers.StandardResponse(c, http.StatusInternalServerError, err.Error(), nil, nil)
	}
	return helpers.StandardResponse(c, http.StatusOK, "Berhasil", list, nil)
}

// Create godoc
// @Summary      Tambah kategori tanda bahaya baru
// @Tags         kategori-tanda-bahaya
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body  body      models.KategoriTandaBahaya  true  "Data kategori tanda bahaya"
// @Success      201   {object}  models.Response
// @Failure      400   {object}  models.Response
// @Router       /kategori-tanda-bahaya [post]
func (h *KategoriTandaBahayaController) Create(c echo.Context) error {
	var req models.KategoriTandaBahaya
	if err := c.Bind(&req); err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, "request tidak valid: "+err.Error(), nil, nil)
	}

	if err := h.kategoriTandaBahayaUC.Create(&req); err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, err.Error(), nil, nil)
	}

	return helpers.StandardResponse(c, http.StatusCreated, "data kategori tanda bahaya berhasil ditambahkan", req, nil)
}

// Detail godoc
// @Summary      Detail kategori tanda bahaya berdasarkan ID
// @Tags         kategori-tanda-bahaya
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      string  true  "Kategori Tanda Bahaya ID"
// @Success      200  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Router       /kategori-tanda-bahaya/{id} [get]
func (h *KategoriTandaBahayaController) Detail(c echo.Context) error {
	id64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, "id tidak valid", nil, nil)
	}

	kategori, err := h.kategoriTandaBahayaUC.GetByID(int32(id64))
	if err != nil {
		return helpers.StandardResponse(c, http.StatusNotFound, err.Error(), nil, nil)
	}

	return helpers.StandardResponse(c, http.StatusOK, "berhasil", kategori, nil)
}

// GetByTipeAndKategoriUsia godoc
// @Summary      Dapatkan kategori tanda bahaya berdasarkan tipe dan kategori usia
// @Tags         kategori-tanda-bahaya
// @Security     BearerAuth
// @Produce      json
// @Param        tipe             query   string  true  "Tipe Lembar (A, B, Umum)"
// @Param        kategori_usia    query   string  true  "Kategori Usia"
// @Success      200              {object}  models.Response
// @Failure      400              {object}  models.Response
// @Router       /kategori-tanda-bahaya/filter [get]
func (h *KategoriTandaBahayaController) GetByTipeAndKategoriUsia(c echo.Context) error {
	tipe := c.QueryParam("tipe")
	kategoriUsia := c.QueryParam("kategori_usia")

	if tipe == "" || kategoriUsia == "" {
		return helpers.StandardResponse(c, http.StatusBadRequest, "tipe dan kategori_usia wajib diisi", nil, nil)
	}

	list, err := h.kategoriTandaBahayaUC.GetByTipeAndKategoriUsia(tipe, kategoriUsia)
	if err != nil {
		return helpers.StandardResponse(c, http.StatusInternalServerError, err.Error(), nil, nil)
	}

	return helpers.StandardResponse(c, http.StatusOK, "berhasil", list, nil)
}

// Update godoc
// @Summary      Update data kategori tanda bahaya
// @Tags         kategori-tanda-bahaya
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id    path      string                     true  "Kategori Tanda Bahaya ID"
// @Param        body  body      models.KategoriTandaBahaya true  "Data kategori tanda bahaya yang diperbarui"
// @Success      200   {object}  models.Response
// @Failure      404   {object}  models.Response
// @Router       /kategori-tanda-bahaya/{id} [put]
func (h *KategoriTandaBahayaController) Update(c echo.Context) error {
	idParam := c.Param("id")

	id64, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, "id tidak valid", nil, nil)
	}

	var req models.KategoriTandaBahaya
	if err := c.Bind(&req); err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, "request tidak valid: "+err.Error(), nil, nil)
	}

	req.ID = int32(id64)

	if err := h.kategoriTandaBahayaUC.Update(&req); err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, err.Error(), nil, nil)
	}

	return helpers.StandardResponse(c, http.StatusOK, "data kategori tanda bahaya berhasil diperbarui", req, nil)
}

// Delete godoc
// @Summary      Hapus data kategori tanda bahaya
// @Tags         kategori-tanda-bahaya
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      string  true  "Kategori Tanda Bahaya ID"
// @Success      200  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Router       /kategori-tanda-bahaya/{id} [delete]
func (h *KategoriTandaBahayaController) Delete(c echo.Context) error {
	idParam := c.Param("id")

	id64, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		return helpers.StandardResponse(c, http.StatusBadRequest, "id tidak valid", nil, nil)
	}

	if err := h.kategoriTandaBahayaUC.Delete(int32(id64)); err != nil {
		return helpers.StandardResponse(c, http.StatusNotFound, err.Error(), nil, nil)
	}

	return helpers.StandardResponse(c, http.StatusOK, "data kategori tanda bahaya berhasil dihapus", nil, nil)
}
