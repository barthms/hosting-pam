package routes

import (
	"monitoring-service/app/controllers"
	"github.com/labstack/echo/v4"
)

// SetupPrediksiStuntingRoutes - Register routes untuk prediksi stunting
func SetupPrediksiStuntingRoutes(e *echo.Echo, c *controllers.PrediksiStuntingController) {
	g := e.Group("/api/v1")
	
	// Measurement data endpoints
	g.GET("/anak/:anak_id/measurement-data", c.GetMeasurementData)
	
	// Stunting prediction endpoints
	g.POST("/prediksi-stunting", c.PredictStunting)
	g.GET("/anak/:anak_id/prediksi-stunting", c.GetPredictionHistory)
	g.GET("/anak/:anak_id/prediksi-stunting/latest", c.GetLatestPrediction)
}
