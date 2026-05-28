package middlewares

import (
	"net/http"
	"strings"

	"monitoring-service/app/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTAuth(jwtSecret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"status_code": http.StatusUnauthorized,
					"message":     []string{"authorization header wajib diisi"},
				})
			}

			tokenValue := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
			if tokenValue == "" || tokenValue == authHeader {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"status_code": http.StatusUnauthorized,
					"message":     []string{"format authorization harus Bearer <token>"},
				})
			}

			parsedToken, err := jwt.ParseWithClaims(tokenValue, &models.AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtSecret), nil
			})
			if err != nil || !parsedToken.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"status_code": http.StatusUnauthorized,
					"message":     []string{"token tidak valid atau sudah expired"},
				})
			}

			claims, ok := parsedToken.Claims.(*models.AuthClaims)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"status_code": http.StatusUnauthorized,
					"message":     []string{"claims token tidak valid"},
				})
			}

			c.Set("auth_claims", claims)
			c.Set("role", claims.Role)
			return next(c)
		}
	}
}
