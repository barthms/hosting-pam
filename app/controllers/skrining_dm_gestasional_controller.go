package controllers

import (
	"net/http"
	"strconv"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type SkriningDMGestasionalController struct {
	usecase usecases.SkriningDMGestasionalUsecase
}

func NewSkriningDMGestasionalController(u usecases.SkriningDMGestasionalUsecase) *SkriningDMGestasionalController {
	return &SkriningDMGestasionalController{usecase: u}
}

type createSkriningDMRequest struct {
	KehamilanID                      int32  `json:"kehamilan_id"`
	GulaDarahPuasaHasil              string `json:"gula_darah_puasa_hasil"`
	GulaDarahPuasaRencana            string `json:"gula_darah_puasa_rencana_tindak_lanjut"`
	GulaDarah2JamPostPrandialHasil   string `json:"gula_darah_2_jam_post_prandial_hasil"`
	GulaDarah2JamPostPrandialRencana string `json:"gula_darah_2_jam_post_prandial_rencana_tindak_lanjut"`
}

func (c *SkriningDMGestasionalController) Create(ctx echo.Context) error {
	claims, _ := ctx.Get("auth_claims").(*models.AuthClaims)
	if claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
	}
	var req createSkriningDMRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	s := &models.SkriningDMGestasional{
		KehamilanID:                      req.KehamilanID,
		GulaDarahPuasaHasil:              req.GulaDarahPuasaHasil,
		GulaDarahPuasaRencana:            req.GulaDarahPuasaRencana,
		GulaDarah2JamPostPrandialHasil:   req.GulaDarah2JamPostPrandialHasil,
		GulaDarah2JamPostPrandialRencana: req.GulaDarah2JamPostPrandialRencana,
	}
	if err := c.usecase.Create(s); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusCreated, models.Response{StatusCode: http.StatusCreated, Data: s})
}

// GetByID, GetByKehamilanID, Update, Delete (sama seperti pola sebelumnya)
func (c *SkriningDMGestasionalController) GetByID(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	data, err := c.usecase.GetByID(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: data})
}

func (c *SkriningDMGestasionalController) GetByKehamilanID(ctx echo.Context) error {
	kehamilanID, err := strconv.ParseInt(ctx.QueryParam("kehamilan_id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "kehamilan_id required"})
	}
	list, err := c.usecase.GetByKehamilanID(int32(kehamilanID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: list})
}

func (c *SkriningDMGestasionalController) Update(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	var req createSkriningDMRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	existing, err := c.usecase.GetByID(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: "Data tidak ditemukan"})
	}
	if req.GulaDarahPuasaHasil != "" {
		existing.GulaDarahPuasaHasil = req.GulaDarahPuasaHasil
	}
	if req.GulaDarahPuasaRencana != "" {
		existing.GulaDarahPuasaRencana = req.GulaDarahPuasaRencana
	}
	if req.GulaDarah2JamPostPrandialHasil != "" {
		existing.GulaDarah2JamPostPrandialHasil = req.GulaDarah2JamPostPrandialHasil
	}
	if req.GulaDarah2JamPostPrandialRencana != "" {
		existing.GulaDarah2JamPostPrandialRencana = req.GulaDarah2JamPostPrandialRencana
	}
	if err := c.usecase.Update(existing); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: existing})
}

func (c *SkriningDMGestasionalController) Delete(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "invalid id"})
	}
	if err := c.usecase.Delete(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Message: "deleted"})
}
