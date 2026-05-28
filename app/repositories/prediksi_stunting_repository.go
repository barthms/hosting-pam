package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type PrediksiStuntingRepository interface {
	GetMeasurementDataByAnakID(anakID int32) (*models.MeasurementDataForPrediction, error)
	GetLatestMeasurementByAnakID(anakID int32) (*models.MeasurementDataForPrediction, error)
	SavePrediction(prediction *models.PrediksiStunting) error
	GetPredictionByAnakID(anakID int32) ([]models.PrediksiStunting, error)
	GetLatestPredictionByAnakID(anakID int32) (*models.PrediksiStunting, error)
}

type prediksiStuntingRepository struct {
	db *gorm.DB
}

func NewPrediksiStuntingRepository(db *gorm.DB) PrediksiStuntingRepository {
	return &prediksiStuntingRepository{db: db}
}

// GetMeasurementDataByAnakID - ambil semua data pengukuran anak (latest)
func (r *prediksiStuntingRepository) GetLatestMeasurementByAnakID(anakID int32) (*models.MeasurementDataForPrediction, error) {
	var measurement models.MeasurementDataForPrediction

	// Join catatan_pertumbuhan dengan pengukuran_lila
	err := r.db.
		Table("catatan_pertumbuhan cp").
		Select(`
			cp.anak_id,
			cp.berat_badan,
			cp.tinggi_badan,
			cp.lingkar_kepala,
			cp.usia_ukur_bulan,
			cp.tgl_ukur,
			COALESCE(pl.hasil_lila, 0) as hasil_lila,
			cp.status_tb_u,
			cp.z_score_tb_u
		`).
		Joins("LEFT JOIN pengukuran_lila pl ON cp.anak_id = pl.anak_id AND DATE(cp.tgl_ukur) = DATE(pl.tanggal)").
		Where("cp.anak_id = ? AND cp.deleted_at IS NULL", anakID).
		Order("cp.tgl_ukur DESC").
		Limit(1).
		Scan(&measurement).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &measurement, nil
}

// GetMeasurementDataByAnakID - ambil riwayat lengkap data pengukuran (untuk training/analisis)
func (r *prediksiStuntingRepository) GetMeasurementDataByAnakID(anakID int32) (*models.MeasurementDataForPrediction, error) {
	return r.GetLatestMeasurementByAnakID(anakID)
}

// SavePrediction - simpan hasil prediksi ke database
func (r *prediksiStuntingRepository) SavePrediction(prediction *models.PrediksiStunting) error {
	return r.db.Create(prediction).Error
}

// GetPredictionByAnakID - ambil riwayat prediksi stunting
func (r *prediksiStuntingRepository) GetPredictionByAnakID(anakID int32) ([]models.PrediksiStunting, error) {
	var predictions []models.PrediksiStunting

	err := r.db.
		Where("anak_id = ? AND deleted_at IS NULL", anakID).
		Order("created_at DESC").
		Find(&predictions).Error

	return predictions, err
}

// GetLatestPredictionByAnakID - ambil prediksi terbaru
func (r *prediksiStuntingRepository) GetLatestPredictionByAnakID(anakID int32) (*models.PrediksiStunting, error) {
	var prediction models.PrediksiStunting

	err := r.db.
		Where("anak_id = ? AND deleted_at IS NULL", anakID).
		Order("created_at DESC").
		Limit(1).
		First(&prediction).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &prediction, nil
}
