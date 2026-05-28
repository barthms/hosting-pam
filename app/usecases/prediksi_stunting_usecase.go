package usecases

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
	"monitoring-service/pkg/customerror"
	"net/http"
)

type PrediksiStuntingUsecase interface {
	GetMeasurementDataByAnakID(ctx context.Context, anakID int32) (*models.MeasurementDataForPrediction, error)
	PredictStunting(ctx context.Context, req *models.PrediksiStuntingRequest) (*models.PrediksiStunting, error)
	GetPredictionHistory(ctx context.Context, anakID int32) ([]models.PrediksiStunting, error)
	GetLatestPrediction(ctx context.Context, anakID int32) (*models.PrediksiStunting, error)
}

type prediksiStuntingUsecase struct {
	repo                repositories.PrediksiStuntingRepository
	pythonServiceURL    string
}

func NewPrediksiStuntingUsecase(
	repo repositories.PrediksiStuntingRepository,
	pythonServiceURL string,
) PrediksiStuntingUsecase {
	return &prediksiStuntingUsecase{
		repo:             repo,
		pythonServiceURL: pythonServiceURL,
	}
}

// GetMeasurementDataByAnakID - ambil data pengukuran terakhir anak
func (u *prediksiStuntingUsecase) GetMeasurementDataByAnakID(ctx context.Context, anakID int32) (*models.MeasurementDataForPrediction, error) {
	data, err := u.repo.GetMeasurementDataByAnakID(anakID)
	if err != nil {
		return nil, customerror.NewInternalServiceError("gagal mengambil data pengukuran")
	}
	return data, nil
}

// PredictStunting - prediksi stunting dengan memanggil Python service
func (u *prediksiStuntingUsecase) PredictStunting(ctx context.Context, req *models.PrediksiStuntingRequest) (*models.PrediksiStunting, error) {
	// Panggil Python service untuk prediksi
	prediksi, err := u.callPythonPredictionService(ctx, req)
	if err != nil {
		return nil, err
	}

	// Simpan hasil prediksi ke database
	result := &models.PrediksiStunting{
		AnakID:          req.AnakID,
		BeratBadan:      req.BeratBadan,
		TinggiBadan:     req.TinggiBadan,
		LingkarKepala:   req.LingkarKepala,
		HasilLila:       req.HasilLila,
		UsiaUkurBulan:   req.UsiaUkurBulan,
		RiskPercentage:  prediksi.StuntingRisk,
		Classification:  prediksi.Classification,
		Confidence:      prediksi.Confidence,
		ZScoreTBU:       prediksi.ZScoreTBU,
		StatusTBU:       prediksi.StatusTBU,
		Rekomendasi:     prediksi.Rekomendasi,
	}

	if err := u.repo.SavePrediction(result); err != nil {
		return nil, customerror.NewInternalServiceError("gagal menyimpan hasil prediksi")
	}

	return result, nil
}

// callPythonPredictionService - panggil Python ML service
func (u *prediksiStuntingUsecase) callPythonPredictionService(ctx context.Context, req *models.PrediksiStuntingRequest) (*models.PrediksiResponse, error) {
	// Siapkan payload
	payload := map[string]interface{}{
		"bb":            req.BeratBadan,
		"tb":            req.TinggiBadan,
		"lila":          req.HasilLila,
		"lingkar_kepala": req.LingkarKepala,
		"umur":          req.UsiaUkurBulan,
		"jenis_kelamin": req.JenisKelamin,
	}

	// Convert ke JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, customerror.NewInternalServiceError("gagal serialize payload")
	}

	// Request ke Python service
	httpReq, err := http.NewRequestWithContext(ctx, "POST", 
		fmt.Sprintf("%s/predict", u.pythonServiceURL), 
		io.NopCloser(io.Reader(bytes.NewBuffer(payloadBytes))))
	if err != nil {
		return nil, customerror.NewInternalServiceError("gagal membuat request ke Python service")
	}
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, customerror.NewInternalServiceError("gagal koneksi ke Python service")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, customerror.NewInternalServiceError("Python service error")
	}

	// Parse response
	var prediction models.PrediksiResponse
	if err := json.NewDecoder(resp.Body).Decode(&prediction); err != nil {
		return nil, customerror.NewInternalServiceError("gagal parse response dari Python service")
	}

	return &prediction, nil
}

// GetPredictionHistory - ambil riwayat prediksi
func (u *prediksiStuntingUsecase) GetPredictionHistory(ctx context.Context, anakID int32) ([]models.PrediksiStunting, error) {
	predictions, err := u.repo.GetPredictionByAnakID(anakID)
	if err != nil {
		return nil, customerror.NewInternalServiceError("gagal mengambil riwayat prediksi")
	}
	return predictions, nil
}

// GetLatestPrediction - ambil prediksi terbaru
func (u *prediksiStuntingUsecase) GetLatestPrediction(ctx context.Context, anakID int32) (*models.PrediksiStunting, error) {
	prediction, err := u.repo.GetLatestPredictionByAnakID(anakID)
	if err != nil {
		return nil, customerror.NewInternalServiceError("gagal mengambil prediksi terbaru")
	}
	return prediction, nil
}
