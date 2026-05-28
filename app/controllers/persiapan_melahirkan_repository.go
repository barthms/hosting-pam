package controllers

import (
	"net/http"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type PersiapanMelahirkanController struct {
	usecase usecases.PersiapanMelahirkanUsecase
}

func NewPersiapanMelahirkanController(
	u usecases.PersiapanMelahirkanUsecase,
) *PersiapanMelahirkanController {
	return &PersiapanMelahirkanController{usecase: u}
}

type savePersiapanMelahirkanRequest struct {
	PerkiraanPersalinan  bool `json:"perkiraan_persalinan"`
	PendampingPersalinan bool `json:"pendamping_persalinan"`
	DanaPersalinan       bool `json:"dana_persalinan"`
	StatusJKN            bool `json:"status_jkn"`
	FaskesPersalinan     bool `json:"faskes_persalinan"`
	PendonorDarah        bool `json:"pendonor_darah"`
	Transportasi         bool `json:"transportasi"`
	MetodeKB             bool `json:"metode_kb"`
	ProgramP4K           bool `json:"program_p4k"`
	DokumenPenting       bool `json:"dokumen_penting"`
}

func (c *PersiapanMelahirkanController) GetMine(ctx echo.Context) error {
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

func (c *PersiapanMelahirkanController) SaveMine(ctx echo.Context) error {
	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
	if !ok || claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "token tidak valid",
		})
	}

	var req savePersiapanMelahirkanRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "format request tidak valid",
		})
	}

	data, err := c.usecase.SaveMine(
		claims.UserID,
		models.PersiapanMelahirkan{
			PerkiraanPersalinan:  req.PerkiraanPersalinan,
			PendampingPersalinan: req.PendampingPersalinan,
			DanaPersalinan:       req.DanaPersalinan,
			StatusJKN:            req.StatusJKN,
			FaskesPersalinan:     req.FaskesPersalinan,
			PendonorDarah:        req.PendonorDarah,
			Transportasi:         req.Transportasi,
			MetodeKB:             req.MetodeKB,
			ProgramP4K:           req.ProgramP4K,
			DokumenPenting:       req.DokumenPenting,
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