package controllers

import (
	"net/http"
	"strconv"

	"monitoring-service/app/constants"
	"monitoring-service/app/helpers"
	"monitoring-service/app/models"

	"github.com/labstack/echo/v4"
)

func (m *Main) GetKunjunganImunisasiByID(
	c echo.Context,
) error {

	idParam := c.Param("id")

	kunjunganID, err :=
		strconv.Atoi(idParam)

	if err != nil || kunjunganID <= 0 {

		return helpers.Response(
			c,
			http.StatusBadRequest,
			[]string{
				"id tidak valid",
			},
		)
	}

	data, err :=
		m.usecases.
			GetKunjunganImunisasiByID(
				uint(kunjunganID),
			)

	if err != nil {

		return helpers.Response(
			c,
			http.StatusInternalServerError,
			[]string{
				err.Error(),
			},
		)
	}

	if data == nil {

		return helpers.Response(
			c,
			http.StatusNotFound,
			[]string{
				"data tidak ditemukan",
			},
		)
	}

	return helpers.StandardResponse(
		c,
		http.StatusOK,
		[]string{
			constants.SUCCESS_RESPONSE_MESSAGE,
		},
		data,
		nil,
	)
}

func (m *Main) GetAllKunjunganImunisasi(
	c echo.Context,
) error {

	data, err :=
		m.usecases.
			GetAllKunjunganImunisasi()

	if err != nil {

		return helpers.Response(
			c,
			http.StatusInternalServerError,
			[]string{
				err.Error(),
			},
		)
	}

	return helpers.StandardResponse(
		c,
		http.StatusOK,
		[]string{
			constants.SUCCESS_RESPONSE_MESSAGE,
		},
		data,
		nil,
	)
}

func (m *Main) UpdateStatusKunjungan(
	c echo.Context,
) error {

	idParam := c.Param("id")

	kunjunganID, err :=
		strconv.Atoi(idParam)

	if err != nil ||
		kunjunganID <= 0 {

		return helpers.Response(
			c,
			http.StatusBadRequest,
			[]string{
				"id tidak valid",
			},
		)
	}

	var req models.UpdateStatusKunjunganRequest

	if err := c.Bind(&req); err != nil {

		return helpers.Response(
			c,
			http.StatusBadRequest,
			[]string{
				"invalid request",
			},
		)
	}

	if req.StatusKunjunganID == 0 {

		return helpers.Response(
			c,
			http.StatusBadRequest,
			[]string{
				"status_kunjungan_id wajib diisi",
			},
		)
	}

	err =
		m.usecases.
			UpdateStatusKunjungan(
				uint(kunjunganID),
				req.StatusKunjunganID,
			)

	if err != nil {

		return helpers.Response(
			c,
			http.StatusInternalServerError,
			[]string{
				err.Error(),
			},
		)
	}

	return helpers.StandardResponse(
		c,
		http.StatusOK,
		[]string{
			"success",
		},
		nil,
		nil,
	)
}

func (m *Main) UpdateTanggalKunjungan(
	c echo.Context,
) error {

	idParam := c.Param("id")

	kunjunganID, err :=
		strconv.Atoi(idParam)

	if err != nil ||
		kunjunganID <= 0 {

		return helpers.Response(
			c,
			http.StatusBadRequest,
			[]string{
				"id tidak valid",
			},
		)
	}

	var req models.UpdateTanggalKunjunganRequest

	if err := c.Bind(&req); err != nil {

		return helpers.Response(
			c,
			http.StatusBadRequest,
			[]string{
				"invalid request",
			},
		)
	}

	if req.TanggalKunjungan == "" {

		return helpers.Response(
			c,
			http.StatusBadRequest,
			[]string{
				"tanggal_kunjungan wajib diisi",
			},
		)
	}

	err =
		m.usecases.
			UpdateTanggalKunjungan(
				uint(kunjunganID),
				req.TanggalKunjungan,
			)

	if err != nil {

		return helpers.Response(
			c,
			http.StatusInternalServerError,
			[]string{
				err.Error(),
			},
		)
	}

	return helpers.StandardResponse(
		c,
		http.StatusOK,
		[]string{
			"success",
		},
		nil,
		nil,
	)
}

func (m *Main) GetKunjunganImunisasiByStatus(
	c echo.Context,
) error {

	idParam :=
		c.Param("status_id")

	statusID, err :=
		strconv.Atoi(idParam)

	if err != nil ||
		statusID <= 0 {

		return helpers.Response(
			c,
			http.StatusBadRequest,
			[]string{
				"status id tidak valid",
			},
		)
	}

	data, err :=
		m.usecases.
			GetKunjunganImunisasiByStatus(
				uint(statusID),
			)

	if err != nil {

		return helpers.Response(
			c,
			http.StatusInternalServerError,
			[]string{
				err.Error(),
			},
		)
	}

	return helpers.StandardResponse(
		c,
		http.StatusOK,
		[]string{
			constants.SUCCESS_RESPONSE_MESSAGE,
		},
		data,
		nil,
	)
}
