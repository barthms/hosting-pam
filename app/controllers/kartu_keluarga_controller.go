package controllers

// import (
// 	"net/http"
// 	"strconv"

// 	"monitoring-service/app/models"
// 	"monitoring-service/app/usecases"

// 	"github.com/labstack/echo/v4"
// )

// type KartuKeluargaController struct {
// 	usecase usecases.KartuKeluargaUsecase
// }

// func NewKartuKeluargaController(u usecases.KartuKeluargaUsecase) *KartuKeluargaController {
// 	return &KartuKeluargaController{usecase: u}
// }

// type createKartuKeluargaRequest struct {
// 	NoKartuKeluarga string `json:"no_kk"`
// }

// func (c *KartuKeluargaController) Create(ctx echo.Context) error {
// 	claims, _ := ctx.Get("auth_claims").(*models.AuthClaims)
// 	if claims == nil {
// 		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
// 	}
// 	var req createKartuKeluargaRequest
// 	if err := ctx.Bind(&req); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
// 	}

// 	// Validasi field wajib
// 	if req.NoKartuKeluarga == "" {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "no_kk wajib diisi"})
// 	}

// 	kk := &models.KartuKeluarga{
// 		NoKK: req.NoKartuKeluarga,
// 	}

// 	data, err := c.usecase.Create(kk)
// 	if err != nil {
// 		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
// 	}
// 	return ctx.JSON(http.StatusCreated, models.Response{StatusCode: http.StatusCreated, Data: data})
// }

// func (c *KartuKeluargaController) GetAll(ctx echo.Context) error {
// 	list, err := c.usecase.GetAll()
// 	if err != nil {
// 		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: list})
// }

// func (c *KartuKeluargaController) GetByID(ctx echo.Context) error {
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

// func (c *KartuKeluargaController) GetByNoKartuKeluarga(ctx echo.Context) error {
// 	noKK := ctx.Param("no_kk")
// 	data, err := c.usecase.GetByNoKartuKeluargaWithMembers(noKK)
// 	if err != nil {
// 		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: data})
// }

// func (c *KartuKeluargaController) Update(ctx echo.Context) error {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
// 	}
// 	var req createKartuKeluargaRequest
// 	if err := ctx.Bind(&req); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
// 	}
// 	existing, err := c.usecase.GetByID(int32(id))
// 	if err != nil {
// 		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: "Data tidak ditemukan"})
// 	}
// 	// Update field yang diisi
// 	if req.NoKartuKeluarga != "" {
// 		existing.NoKK = req.NoKartuKeluarga
// 	}
// 	if err := c.usecase.Update(existing); err != nil {
// 		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: existing})
// }

// func (c *KartuKeluargaController) Delete(ctx echo.Context) error {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
// 	}
// 	if err := c.usecase.Delete(int32(id)); err != nil {
// 		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Message: "deleted"})
// }
