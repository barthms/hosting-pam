package controllers

import (
	"monitoring-service/app/models"
	"strings"

	"github.com/labstack/echo/v4"
)

func getAuthClaims(c echo.Context) (*models.AuthClaims, bool) {
	claims, ok := c.Get("auth_claims").(*models.AuthClaims)
	if !ok || claims == nil {
		return nil, false
	}
	return claims, true
}

func isTenagaKesehatanRole(role string) bool {
	v := strings.TrimSpace(strings.ToLower(role))
	return v == "tenaga-kesehatan" || v == "tenaga kesehatan"
}

func isOrangtuaRole(role string) bool {
	v := strings.TrimSpace(strings.ToLower(role))
	return v == "orangtua" || v == "orang tua" || v == "ibu"
}
