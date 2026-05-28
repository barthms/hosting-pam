package controllers

import (
	"net/http"
	"regexp"
	"strconv"
	"time"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type JadwalLayananController struct {
	usecase usecases.JadwalLayananUsecase
}

func NewJadwalLayananController(u usecases.JadwalLayananUsecase) *JadwalLayananController {
	return &JadwalLayananController{u}
}

func timePtrToDBString(t *time.Time) *string {
	if t == nil {
		return nil
	}
	v := t.Format("15:04:05")
	return &v
}

func (c *JadwalLayananController) Create(ctx echo.Context) error {
	// Bind to intermediate struct to validate formats (tanggal as YYYY-MM-DD)
	var in struct {
		PosyanduID   *int32 `json:"posyandu_id"`
		Layanan      string `json:"layanan"`
		Tanggal      string `json:"tanggal"`
		WaktuMulai   string `json:"waktu_mulai"`
		WaktuSelesai string `json:"waktu_selesai"`
		Keterangan   string `json:"keterangan"`
	}

	if err := ctx.Bind(&in); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid_request", "details": err.Error()})
	}

	// validation
	var fieldErrors []map[string]string
	if in.Layanan == "" {
		fieldErrors = append(fieldErrors, map[string]string{"field": "layanan", "message": "is required"})
	}

	var tanggalPtr *time.Time
	if in.Tanggal == "" {
		fieldErrors = append(fieldErrors, map[string]string{"field": "tanggal", "message": "is required, format YYYY-MM-DD"})
	} else {
		if t, err := time.Parse("2006-01-02", in.Tanggal); err != nil {
			fieldErrors = append(fieldErrors, map[string]string{"field": "tanggal", "message": "invalid format, use YYYY-MM-DD"})
		} else {
			tanggalPtr = &t
		}
	}

	// waktu validation: waktu_mulai required (HH:MM), waktu_selesai optional (HH:MM)
	if in.WaktuMulai == "" {
		fieldErrors = append(fieldErrors, map[string]string{"field": "waktu_mulai", "message": "is required, format HH:MM"})
	} else {
		re := regexp.MustCompile(`^([01]\d|2[0-3]):[0-5]\d$`)
		if !re.MatchString(in.WaktuMulai) {
			fieldErrors = append(fieldErrors, map[string]string{"field": "waktu_mulai", "message": "invalid format, use HH:MM"})
		}
	}
	if in.WaktuSelesai != "" {
		re := regexp.MustCompile(`^([01]\d|2[0-3]):[0-5]\d$`)
		if !re.MatchString(in.WaktuSelesai) {
			fieldErrors = append(fieldErrors, map[string]string{"field": "waktu_selesai", "message": "invalid format, use HH:MM"})
		}
	}

	var waktuMulaiParsed *time.Time
	if in.WaktuMulai != "" {
		if t, err := time.Parse("15:04", in.WaktuMulai); err != nil {
			fieldErrors = append(fieldErrors, map[string]string{"field": "waktu_mulai", "message": "invalid format, use HH:MM"})
		} else {
			waktuMulaiParsed = &t
		}
	}

	var waktuSelesaiParsed *time.Time
	if in.WaktuSelesai != "" {
		if t, err := time.Parse("15:04", in.WaktuSelesai); err != nil {
			fieldErrors = append(fieldErrors, map[string]string{"field": "waktu_selesai", "message": "invalid format, use HH:MM"})
		} else {
			waktuSelesaiParsed = &t
		}
	}

	if waktuMulaiParsed != nil && waktuSelesaiParsed != nil && !waktuSelesaiParsed.After(*waktuMulaiParsed) {
		fieldErrors = append(fieldErrors, map[string]string{"field": "waktu_selesai", "message": "must be later than waktu_mulai"})
	}

	if len(fieldErrors) > 0 {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "validation", "fields": fieldErrors})
	}

	// build model
	model := models.JadwalLayanan{
		PosyanduID: in.PosyanduID,
		Layanan:    in.Layanan,
		Tanggal:    tanggalPtr,
		Keterangan: in.Keterangan,
	}

	// parse waktu_mulai and waktu_selesai into time.Time (time-only)
	model.WaktuMulai = timePtrToDBString(waktuMulaiParsed)
	model.WaktuSelesai = timePtrToDBString(waktuSelesaiParsed)

	if err := c.usecase.Create(&model); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "create_failed", "details": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, model)
}

func (c *JadwalLayananController) GetAll(ctx echo.Context) error {
	// Query params: posyandu_id, from, to, upcoming (bool), limit
	posyanduParam := ctx.QueryParam("posyandu_id")
	fromParam := ctx.QueryParam("from")
	toParam := ctx.QueryParam("to")
	upcomingParam := ctx.QueryParam("upcoming")
	limitParam := ctx.QueryParam("limit")

	// upcoming handling
	if upcomingParam == "true" {
		limit := 5
		if limitParam != "" {
			if v, err := strconv.Atoi(limitParam); err == nil && v > 0 {
				limit = v
			}
		}
		data, err := c.usecase.GetUpcoming(limit)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
		return ctx.JSON(http.StatusOK, data)
	}

	// date range handling
	var posyanduID *int32
	if posyanduParam != "" {
		if v, err := strconv.Atoi(posyanduParam); err == nil {
			tmp := int32(v)
			posyanduID = &tmp
		}
	}

	var fromTime *time.Time
	var toTime *time.Time
	if fromParam != "" {
		if t, err := time.Parse("2006-01-02", fromParam); err == nil {
			fromTime = &t
		}
	}
	if toParam != "" {
		if t, err := time.Parse("2006-01-02", toParam); err == nil {
			toTime = &t
		}
	}

	// If any filter provided, use date-range query; otherwise return all
	if posyanduID != nil || fromTime != nil || toTime != nil {
		data, err := c.usecase.GetByDateRange(posyanduID, fromTime, toTime)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
		return ctx.JSON(http.StatusOK, data)
	}

	data, err := c.usecase.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, data)
}

func (c *JadwalLayananController) GetByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}

	data, err := c.usecase.GetByID(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Data not found"})
	}

	return ctx.JSON(http.StatusOK, data)
}

func (c *JadwalLayananController) Update(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}
	var in struct {
		PosyanduID   *int32 `json:"posyandu_id"`
		Layanan      string `json:"layanan"`
		Tanggal      string `json:"tanggal"`
		WaktuMulai   string `json:"waktu_mulai"`
		WaktuSelesai string `json:"waktu_selesai"`
		Keterangan   string `json:"keterangan"`
	}

	if err := ctx.Bind(&in); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid_request", "details": err.Error()})
	}

	var fieldErrors []map[string]string
	var tanggalPtr *time.Time
	if in.Tanggal != "" {
		if t, err := time.Parse("2006-01-02", in.Tanggal); err != nil {
			fieldErrors = append(fieldErrors, map[string]string{"field": "tanggal", "message": "invalid format, use YYYY-MM-DD"})
		} else {
			tanggalPtr = &t
		}
	}

	if in.WaktuMulai != "" {
		re := regexp.MustCompile(`^([01]\d|2[0-3]):[0-5]\d$`)
		if !re.MatchString(in.WaktuMulai) {
			fieldErrors = append(fieldErrors, map[string]string{"field": "waktu_mulai", "message": "invalid format, use HH:MM"})
		}
	}
	if in.WaktuSelesai != "" {
		re := regexp.MustCompile(`^([01]\d|2[0-3]):[0-5]\d$`)
		if !re.MatchString(in.WaktuSelesai) {
			fieldErrors = append(fieldErrors, map[string]string{"field": "waktu_selesai", "message": "invalid format, use HH:MM"})
		}
	}

	var waktuMulaiParsed *time.Time
	if in.WaktuMulai != "" {
		if t, err := time.Parse("15:04", in.WaktuMulai); err != nil {
			fieldErrors = append(fieldErrors, map[string]string{"field": "waktu_mulai", "message": "invalid format, use HH:MM"})
		} else {
			waktuMulaiParsed = &t
		}
	}

	var waktuSelesaiParsed *time.Time
	if in.WaktuSelesai != "" {
		if t, err := time.Parse("15:04", in.WaktuSelesai); err != nil {
			fieldErrors = append(fieldErrors, map[string]string{"field": "waktu_selesai", "message": "invalid format, use HH:MM"})
		} else {
			waktuSelesaiParsed = &t
		}
	}

	if waktuMulaiParsed != nil && waktuSelesaiParsed != nil && !waktuSelesaiParsed.After(*waktuMulaiParsed) {
		fieldErrors = append(fieldErrors, map[string]string{"field": "waktu_selesai", "message": "must be later than waktu_mulai"})
	}

	if len(fieldErrors) > 0 {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "validation", "fields": fieldErrors})
	}

	// build partial model for update
	model := models.JadwalLayanan{}
	if in.PosyanduID != nil {
		model.PosyanduID = in.PosyanduID
	}
	if in.Layanan != "" {
		model.Layanan = in.Layanan
	}
	if tanggalPtr != nil {
		model.Tanggal = tanggalPtr
	}
	if waktuMulaiParsed != nil {
		model.WaktuMulai = timePtrToDBString(waktuMulaiParsed)
	}
	if waktuSelesaiParsed != nil {
		model.WaktuSelesai = timePtrToDBString(waktuSelesaiParsed)
	}
	if in.Keterangan != "" {
		model.Keterangan = in.Keterangan
	}
	// kapasitas removed

	if err := c.usecase.Update(int32(id), &model); err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "updated successfully"})
}

func (c *JadwalLayananController) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}

	if err := c.usecase.Delete(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "deleted successfully"})
}
