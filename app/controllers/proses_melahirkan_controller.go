package controllers

import (
	"net/http"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type ProsesMelahirkanController struct {
	usecase usecases.ProsesMelahirkanUsecase
}

func NewProsesMelahirkanController(
	u usecases.ProsesMelahirkanUsecase,
) *ProsesMelahirkanController {
	return &ProsesMelahirkanController{usecase: u}
}

type saveProsesMelahirkanRequest struct {
	TandaPersalinan        bool `json:"tanda_persalinan"`
	ProsesMelahirkan       bool `json:"proses_melahirkan"`
	HakIbuPendamping       bool `json:"hak_ibu_pendamping"`
	HakIbuPosisiMelahirkan bool `json:"hak_ibu_posisi_melahirkan"`
	Mulas                  bool `json:"mulas"`
	TeknikMengurangiNyeri  bool `json:"teknik_mengurangi_nyeri"`
	IMDKontakKulit         bool `json:"imd_kontak_kulit"`
}

func (c *ProsesMelahirkanController) GetMine(ctx echo.Context) error {
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

func (c *ProsesMelahirkanController) SaveMine(ctx echo.Context) error {
	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
	if !ok || claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "token tidak valid",
		})
	}

	var req saveProsesMelahirkanRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "format request tidak valid",
		})
	}

	data, err := c.usecase.SaveMine(
		claims.UserID,
		models.ProsesMelahirkan{
			TandaPersalinan:        req.TandaPersalinan,
			ProsesMelahirkan:       req.ProsesMelahirkan,
			HakIbuPendamping:       req.HakIbuPendamping,
			HakIbuPosisiMelahirkan: req.HakIbuPosisiMelahirkan,
			Mulas:                  req.Mulas,
			TeknikMengurangiNyeri:  req.TeknikMengurangiNyeri,
			IMDKontakKulit:         req.IMDKontakKulit,
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
