package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type LogTTDMMSRepository struct {
	db *gorm.DB
}

func NewLogTTDMMSRepository(db *gorm.DB) *LogTTDMMSRepository {
	return &LogTTDMMSRepository{db: db}
}

func (r *LogTTDMMSRepository) FindActiveKehamilanByUserID(userID int32) (*models.Kehamilan, error) {
	var kehamilan models.Kehamilan

	err := r.db.
		Table("kehamilan k").
		Select("k.*").
		Joins("JOIN ibu i ON i.id = k.ibu_id").
		Joins("JOIN penduduk p ON p.id = i.penduduk_id").
		Joins("JOIN pengguna u ON u.penduduk_id = p.id").
		Where("u.id = ?", userID).
		Where("k.status_kehamilan IN ?", []string{"aktif", "TRIMESTER 1", "TRIMESTER 2", "TRIMESTER 3"}).
		Order("k.created_at DESC").
		First(&kehamilan).Error

	return &kehamilan, err
}

func (r *LogTTDMMSRepository) FindByKehamilanID(kehamilanID int32) ([]models.LogTTDMMS, error) {
	var list []models.LogTTDMMS

	err := r.db.
		Where("kehamilan_id = ?", kehamilanID).
		Order("bulan_ke ASC, hari_ke ASC").
		Find(&list).Error

	return list, err
}

func (r *LogTTDMMSRepository) Upsert(log *models.LogTTDMMS) error {
	var existing models.LogTTDMMS

	err := r.db.
		Where("kehamilan_id = ? AND bulan_ke = ? AND hari_ke = ?",
			log.KehamilanID,
			log.BulanKe,
			log.HariKe,
		).
		First(&existing).Error

	if err == nil {
		existing.SudahDiminum = log.SudahDiminum
		return r.db.Save(&existing).Error
	}

	if err == gorm.ErrRecordNotFound {
		return r.db.Create(log).Error
	}

	return err
}