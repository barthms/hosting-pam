package controllers

import (
	"net/http"

	"monitoring-service/app/constants"
	"monitoring-service/app/helpers"
	"monitoring-service/app/models"

	"github.com/labstack/echo/v4"
)

func (m *Main) Me(c echo.Context) error {
	claims, ok := c.Get("auth_claims").(*models.AuthClaims)
	if !ok || claims == nil {
		return helpers.Response(c, http.StatusUnauthorized, []string{"token tidak valid"})
	}

	response := models.MeResponse{
		UserID:        claims.UserID,
		Email:         claims.Email,
		PhoneNumber:   claims.PhoneNumber,
		Role:          claims.Role,
		TargetApp:     claims.TargetApp,
		RedirectRoute: claims.RedirectRoute,
	}

	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, response, nil)
}
