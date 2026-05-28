package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetRole mengambil role pengguna dari Echo context.
func GetRole(c echo.Context) string {
	role, _ := c.Get("role").(string)
	return role
}

func TenagaKesehatan() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			role, _ := c.Get("role").(string)

			if role == "" {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"status_code": http.StatusUnauthorized,
					"message":     "role tidak ditemukan",
				})
			}

			if role != "Bidan" && role != "Dokter" && role != "Tenaga-kesehatan" {
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"status_code": http.StatusForbidden,
					"message":     "Anda Tidak Memiliki Akses",
				})
			}

			return next(c)
		}
	}
}

func PemantauanLembarAccess() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			role, _ := c.Get("role").(string)

			if role == "" {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"status_code": http.StatusUnauthorized,
					"message":     "role tidak ditemukan",
				})
			}

			if role != "Bidan" && role != "Dokter" && role != "Tenaga-kesehatan" && role != "Kader" {
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"status_code": http.StatusForbidden,
					"message":     "Anda Tidak Memiliki Akses",
				})
			}

			return next(c)
		}
	}
}

func AdminOnly() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			role, _ := c.Get("role").(string)

			if role == "" {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"status_code": http.StatusUnauthorized,
					"message":     "role tidak ditemukan",
				})
			}

			if role != "Admin" {
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"status_code": http.StatusForbidden,
					"message":     "Anda Tidak Memiliki Akses",
				})
			}

			return next(c)
		}
	}
}

func IbuOnly() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			role, _ := c.Get("role").(string)

			if role == "" {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"status_code": http.StatusUnauthorized,
					"message":     "role tidak ditemukan",
				})
			}

			if role != "Ibu" {
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"status_code": http.StatusForbidden,
					"message":     "Anda Tidak Memiliki Akses",
				})
			}

			return next(c)
		}
	}
}

func BidanOnly() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			role, _ := c.Get("role").(string)

			if role == "" {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"status_code": http.StatusUnauthorized,
					"message":     "role tidak ditemukan",
				})
			}

			if role != "Bidan" {
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"status_code": http.StatusForbidden,
					"message":     "Anda Tidak Memiliki Akses",
				})
			}

			return next(c)
		}
	}
}
func Kader() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			role, _ := c.Get("role").(string)

			if role == "" {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"status_code": http.StatusUnauthorized,
					"message":     "role tidak ditemukan",
				})
			}

			if role != "Kader" {
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"status_code": http.StatusForbidden,
					"message":     "Anda Tidak Memiliki Akses",
				})
			}

			return next(c)
		}
	}
}
