package controllers

import (
	"net/http"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type PemantauanIbuHamilController struct {
	usecase usecases.PemantauanIbuHamilUsecase
}

func NewPemantauanIbuHamilController(u usecases.PemantauanIbuHamilUsecase) *PemantauanIbuHamilController {
	return &PemantauanIbuHamilController{usecase: u}
}

type savePemantauanIbuHamilRequest struct {
	MingguKehamilan   int32 `json:"minggu_kehamilan"`
	DemamLebih2Hari   bool  `json:"demam_lebih_2_hari"`
	SakitKepala       bool  `json:"sakit_kepala"`
	CemasBerlebih     bool  `json:"cemas_berlebih"`
	ResikoTB          bool  `json:"resiko_tb"`
	GerakanBayiKurang bool  `json:"gerakan_bayi_kurang"`
	NyeriPerut        bool  `json:"nyeri_perut"`
	CairanJalanLahir  bool  `json:"cairan_jalan_lahir"`
	MasalahKemaluan   bool  `json:"masalah_kemaluan"`
	DiareBerulang     bool  `json:"diare_berulang"`
}

func (c *PemantauanIbuHamilController) GetMine(ctx echo.Context) error {
	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
	if !ok || claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "token tidak valid",
		})
	}

	data, err := c.usecase.GetMine(claims.UserID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Data:       data,
	})
}

func (c *PemantauanIbuHamilController) SaveMine(ctx echo.Context) error {
	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
	if !ok || claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "token tidak valid",
		})
	}

	var req savePemantauanIbuHamilRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "format request tidak valid",
		})
	}

	data, err := c.usecase.SaveMine(
		claims.UserID,
		models.PemantauanIbuHamil{
			MingguKehamilan:   req.MingguKehamilan,
			DemamLebih2Hari:   req.DemamLebih2Hari,
			SakitKepala:       req.SakitKepala,
			CemasBerlebih:     req.CemasBerlebih,
			ResikoTB:          req.ResikoTB,
			GerakanBayiKurang: req.GerakanBayiKurang,
			NyeriPerut:        req.NyeriPerut,
			CairanJalanLahir:  req.CairanJalanLahir,
			MasalahKemaluan:   req.MasalahKemaluan,
			DiareBerulang:     req.DiareBerulang,
		},
	)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Data:       data,
	})
}