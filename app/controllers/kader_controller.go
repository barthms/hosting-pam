package controllers

import (
	"net/http"
	"strconv"

	"monitoring-service/app/constants"
	"monitoring-service/app/helpers"
	"monitoring-service/app/middlewares"
	"monitoring-service/app/models"
	"monitoring-service/app/usecases"
	"monitoring-service/pkg/customerror"

	"github.com/labstack/echo/v4"
)

// KaderController - Controller untuk mengelola kader profile
type KaderController struct {
	usecase usecases.KaderUsecase
}

func NewKaderController(usecase usecases.KaderUsecase) *KaderController {
	return &KaderController{usecase: usecase}
}

// ListMyKader - Bidan mendapatkan list kader di posyandu mereka
func (c *KaderController) ListMyKader(ctx echo.Context) error {
	// Ambil data bidan dari context
	claims, _ := ctx.Get("auth_claims").(*models.AuthClaims)
	if claims == nil {
		return helpers.Response(ctx, http.StatusUnauthorized, []string{"unauthorized"})
	}

	// Ambil search keyword dari query param
	searchKeyword := ctx.QueryParam("search")

	// TODO: implementasi fetch posyandu_id dari database berdasarkan user ID
	// Untuk sekarang, kita return error jika belum punya posyandu

	kaders, err := c.usecase.GetMyKaderList(nil, searchKeyword)
	if err != nil {
		return helpers.Response(ctx, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(ctx, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, kaders, nil)
}

// GetKaderDetail - Mendapatkan detail profile kader
func (c *KaderController) GetKaderDetail(ctx echo.Context) error {
	idRaw, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return helpers.Response(ctx, http.StatusBadRequest, []string{"id kader tidak valid"})
	}

	kader, err := c.usecase.GetKaderDetail(int32(idRaw))
	if err != nil {
		return helpers.Response(ctx, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(ctx, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, kader, nil)
}

// UpdateKaderProfile - Bidan mengupdate profile kader mereka
func (c *KaderController) UpdateKaderProfile(ctx echo.Context) error {
	// Check role - hanya bidan yang bisa
	role := middlewares.GetRole(ctx)
	if role != "Bidan" {
		return helpers.Response(ctx, http.StatusForbidden, []string{"hanya bidan yang dapat mengupdate kader"})
	}

	idRaw, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return helpers.Response(ctx, http.StatusBadRequest, []string{"id kader tidak valid"})
	}

	var req usecases.UpdateKaderProfileRequest
	if err := ctx.Bind(&req); err != nil {
		return helpers.Response(ctx, http.StatusBadRequest, []string{"format request tidak valid: " + err.Error()})
	}

	kader, err := c.usecase.UpdateMyKaderProfile(int32(idRaw), &req)
	if err != nil {
		return helpers.Response(ctx, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(ctx, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, kader, nil)
}

// AdminGetAllKader - Admin mendapatkan semua kader
func (c *KaderController) AdminGetAllKader(ctx echo.Context) error {
	desa := ctx.QueryParam("desa")

	kaders, err := c.usecase.GetAllKader(desa)
	if err != nil {
		return helpers.Response(ctx, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(ctx, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, kaders, nil)
}

// AdminCreateKader - Admin membuat kader baru
func (c *KaderController) AdminCreateKader(ctx echo.Context) error {
	var req usecases.CreateKaderRequest
	if err := ctx.Bind(&req); err != nil {
		return helpers.Response(ctx, http.StatusBadRequest, []string{"format request tidak valid: " + err.Error()})
	}

	kader, err := c.usecase.CreateKader(&req)
	if err != nil {
		return helpers.Response(ctx, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(ctx, http.StatusCreated, []string{constants.SUCCESS_RESPONSE_MESSAGE}, kader, nil)
}

// AdminDeleteKader - Admin menghapus kader
func (c *KaderController) AdminDeleteKader(ctx echo.Context) error {
	idRaw, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return helpers.Response(ctx, http.StatusBadRequest, []string{"id kader tidak valid"})
	}

	if err := c.usecase.DeleteKader(int32(idRaw)); err != nil {
		return helpers.Response(ctx, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(ctx, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, map[string]bool{"deleted": true}, nil)
}

// ==================== ROUTE HANDLERS (dengan nama sesuai routes.go) ====================

// AdminTambahKader - Admin membuat kader baru (wrapper untuk AdminCreateKader)
func (c *KaderController) AdminTambahKader(ctx echo.Context) error {
	return c.AdminCreateKader(ctx)
}

// AdminListKader - Admin mendapatkan semua kader (wrapper untuk AdminGetAllKader)
func (c *KaderController) AdminListKader(ctx echo.Context) error {
	return c.AdminGetAllKader(ctx)
}

// AdminUpdateKader - Admin mengupdate data kader
func (c *KaderController) AdminUpdateKader(ctx echo.Context) error {
	idRaw, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return helpers.Response(ctx, http.StatusBadRequest, []string{"id kader tidak valid"})
	}

	var req usecases.UpdateKaderRequest
	if err := ctx.Bind(&req); err != nil {
		return helpers.Response(ctx, http.StatusBadRequest, []string{"format request tidak valid: " + err.Error()})
	}

	kader, err := c.usecase.UpdateKader(int32(idRaw), &req)
	if err != nil {
		return helpers.Response(ctx, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(ctx, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, kader, nil)
}

// AdminUpdateStatusKader - Admin mengupdate status kader
func (c *KaderController) AdminUpdateStatusKader(ctx echo.Context) error {
	idRaw, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return helpers.Response(ctx, http.StatusBadRequest, []string{"id kader tidak valid"})
	}

	var req struct {
		Status string `json:"status"`
	}
	if err := ctx.Bind(&req); err != nil {
		return helpers.Response(ctx, http.StatusBadRequest, []string{"format request tidak valid: " + err.Error()})
	}

	kader, err := c.usecase.UpdateKaderStatus(int32(idRaw), req.Status)
	if err != nil {
		return helpers.Response(ctx, customerror.GetStatusCode(err), []string{err.Error()})
	}

	return helpers.StandardResponse(ctx, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, kader, nil)
}
