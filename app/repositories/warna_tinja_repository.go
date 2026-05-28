package repositories

import (
	"monitoring-service/app/models"
	"monitoring-service/pkg/customerror"
	"time"

	"gorm.io/gorm"
)

type WarnaTinjaRepository interface {
	FindByAnakID(anakID uint) ([]models.WarnaTinjaAnak, error)
	Upsert(data *models.WarnaTinjaAnak) error
	IsAnakMilikIbu(userID uint, anakID uint) (bool, error)
}

type warnaTinjaRepository struct {
	db *gorm.DB
}

func NewWarnaTinjaRepository(db *gorm.DB) WarnaTinjaRepository {
	return &warnaTinjaRepository{db: db}
}

func (r *warnaTinjaRepository) FindByAnakID(anakID uint) ([]models.WarnaTinjaAnak, error) {
	var list []models.WarnaTinjaAnak
	err := r.db.
		Where("anak_id = ? AND deleted_at IS NULL", anakID).
		Order(`CASE periode_key
			WHEN '2_minggu' THEN 1
			WHEN '1_bulan' THEN 2
			WHEN '2_4_bulan' THEN 3
			ELSE 99
		END`).
		Find(&list).Error
	if err != nil {
		return nil, customerror.NewInternalServiceError("gagal mengambil data warna tinja")
	}
	return list, nil
}

func (r *warnaTinjaRepository) Upsert(data *models.WarnaTinjaAnak) error {
	var existing models.WarnaTinjaAnak
	err := r.db.
		Where("anak_id = ? AND periode_key = ? AND deleted_at IS NULL", data.AnakID, data.PeriodeKey).
		First(&existing).Error

	if err == nil {
		existing.TanggalCatat = data.TanggalCatat
		existing.NomorWarna = data.NomorWarna
		existing.PeriodeLabel = data.PeriodeLabel
		existing.UpdatedAt = time.Now()
		return r.db.Save(&existing).Error
	}

	if err == gorm.ErrRecordNotFound {
		return r.db.Create(data).Error
	}

	return customerror.NewInternalServiceError("gagal menyimpan data warna tinja")
}

func (r *warnaTinjaRepository) IsAnakMilikIbu(userID uint, anakID uint) (bool, error) {
	var count int64
	err := r.db.Table("anak a").
		Joins("JOIN kehamilan k ON k.id = a.kehamilan_id").
		Joins("JOIN ibu i ON i.id = k.ibu_id").
		Joins("JOIN penduduk ki ON ki.id = i.penduduk_id").
		Joins("JOIN pengguna p ON p.penduduk_id = ki.id").
		Where("a.id = ?", anakID).
		Where("p.id = ?", userID).
		Count(&count).Error
	if err != nil {
		return false, customerror.NewInternalServiceError("gagal memverifikasi kepemilikan data anak")
	}
	return count > 0, nil
}
