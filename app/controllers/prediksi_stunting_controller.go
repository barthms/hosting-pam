package controllers

import (
	"monitoring-service/app/models"
	"monitoring-service/app/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PrediksiStuntingController struct {
	usecase usecases.PrediksiStuntingUsecase
}

func NewPrediksiStuntingController(usecase usecases.PrediksiStuntingUsecase) *PrediksiStuntingController {
	return &PrediksiStuntingController{usecase: usecase}
}

// GetMeasurementData - ambil data pengukuran terakhir anak (BB, TB, LILA, Lingkar Kepala)
// GET /api/v1/anak/:anak_id/measurement-data
func (c *PrediksiStuntingController) GetMeasurementData(ctx echo.Context) error {
	anakIDStr := ctx.Param("anak_id")
	anakID, err := strconv.ParseInt(anakIDStr, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid anak_id",
		})
	}

	data, err := c.usecase.GetMeasurementDataByAnakID(ctx.Request().Context(), int32(anakID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	if data == nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error": "no measurement data found",
		})
	}

	return ctx.JSON(http.StatusOK, data)
}

// PredictStunting - prediksi stunting berdasarkan input pengukuran
// POST /api/v1/prediksi-stunting
// Body:
// {
//   "anak_id": 1,
//   "berat_badan": 12.5,
//   "tinggi_badan": 85.0,
//   "lingkar_kepala": 46.5,
//   "hasil_lila": 14.5,
//   "usia_ukur_bulan": 24,
//   "jenis_kelamin": "Laki-laki"
// }
func (c *PrediksiStuntingController) PredictStunting(ctx echo.Context) error {
	req := new(models.PrediksiStuntingRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request payload",
		})
	}

	// Validate input
	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	// Call usecase
	result, err := c.usecase.PredictStunting(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, result)
}

// GetPredictionHistory - ambil riwayat prediksi stunting
// GET /api/v1/anak/:anak_id/prediksi-stunting
func (c *PrediksiStuntingController) GetPredictionHistory(ctx echo.Context) error {
	anakIDStr := ctx.Param("anak_id")
	anakID, err := strconv.ParseInt(anakIDStr, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid anak_id",
		})
	}

	predictions, err := c.usecase.GetPredictionHistory(ctx.Request().Context(), int32(anakID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, predictions)
}

// GetLatestPrediction - ambil prediksi stunting terbaru
// GET /api/v1/anak/:anak_id/prediksi-stunting/latest
func (c *PrediksiStuntingController) GetLatestPrediction(ctx echo.Context) error {
	anakIDStr := ctx.Param("anak_id")
	anakID, err := strconv.ParseInt(anakIDStr, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid anak_id",
		})
	}

	prediction, err := c.usecase.GetLatestPrediction(ctx.Request().Context(), int32(anakID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	if prediction == nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error": "no prediction found",
		})
	}

	return ctx.JSON(http.StatusOK, prediction)
}
