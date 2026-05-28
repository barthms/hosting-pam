package controllers

// import (
// 	"monitoring-service/app/constants"
// 	"monitoring-service/app/helpers"
// 	"monitoring-service/pkg/customerror"
// 	"net/http"
// 	"strings"

// 	"github.com/labstack/echo/v4"
// )

// type getAnakSearchRequest struct {
// 	Nama    string `json:"nama"`
// 	NamaIbu string `json:"nama_ibu"`
// 	NoKK    string `json:"no_kk"`
// }

// func (m *Main) GetAllAnak(c echo.Context) error {

// 	data, err := m.usecases.GetAllAnak()
// 	if err != nil {
// 		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
// 	}

// 	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
// }

// func (m *Main) GetAnakById(c echo.Context) error {
// 	claims, ok := getAuthClaims(c)
// 	if !ok {
// 		return helpers.Response(c, http.StatusUnauthorized, []string{"token tidak valid"})
// 	}

// 	if !isTenagaKesehatanRole(claims.Role) && !isOrangtuaRole(claims.Role) {
// 		return helpers.Response(c, http.StatusForbidden, []string{"role anda tidak memiliki akses ke fitur ini"})
// 	}

// 	data, err := m.usecases.GetAnakById(c.Param("anak_id"))
// 	if err != nil {
// 		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
// 	}

// 	return helpers.StandardResponse(c, http.StatusOK, []string{constants.SUCCESS_RESPONSE_MESSAGE}, data, nil)
// }

// func (m *Main) GetAnak(c echo.Context) error {
// 	claims, ok := getAuthClaims(c)
// 	if !ok {
// 		return helpers.Response(c, http.StatusUnauthorized, []string{"token tidak valid"})
// 	}

// 	if !isTenagaKesehatanRole(claims.Role) {
// 		return helpers.Response(c, http.StatusForbidden, []string{"hanya tenaga-kesehatan yang dapat menambah catatan pertumbuhan"})
// 	}

// 	nama := strings.TrimSpace(c.QueryParam("nama"))
// 	namaIbu := strings.TrimSpace(c.QueryParam("nama_ibu"))
// 	noKK := strings.TrimSpace(c.QueryParam("no_kk"))

// 	if c.Request().ContentLength > 0 {
// 		var req getAnakSearchRequest
// 		if err := c.Bind(&req); err != nil {
// 			return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
// 		}

// 		if nama == "" {
// 			nama = strings.TrimSpace(req.Nama)
// 		}
// 		if namaIbu == "" {
// 			namaIbu = strings.TrimSpace(req.NamaIbu)
// 		}
// 		if noKK == "" {
// 			noKK = strings.TrimSpace(req.NoKK)
// 		}
// 	}

// 	data, meta, err := m.usecases.GetAnak(nama, namaIbu, noKK)
// 	if err != nil {
// 		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
// 	}

// 	return helpers.StandardResponse(
// 		c,
// 		http.StatusOK,
// 		[]string{constants.SUCCESS_RESPONSE_MESSAGE},
// 		data,
// 		meta,
// 	)
// }

// func (m *Main) CreateAnak(c echo.Context) error {
// 	var req models.AnakRequest
// 	if err := c.Bind(&req); err != nil {
// 		return helpers.Response(c, http.StatusBadRequest, []string{"format request tidak valid"})
// 	}

// 	if err := m.usecases.CreateAnak(&req); err != nil {
// 		return helpers.Response(c, customerror.GetStatusCode(err), []string{err.Error()})
// 	}

// 	return helpers.StandardResponse(c, http.StatusCreated, []string{constants.SUCCESS_RESPONSE_MESSAGE}, map[string]string{
// 		"message": "anak berhasil ditambahkan",
// 	}, nil)
// }
