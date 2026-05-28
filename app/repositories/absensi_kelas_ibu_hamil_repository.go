package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type AbsensiKelasIbuHamilRepository struct {
	db *gorm.DB
}

func NewAbsensiKelasIbuHamilRepository(db *gorm.DB) *AbsensiKelasIbuHamilRepository {
	return &AbsensiKelasIbuHamilRepository{db: db}
}

func (r *AbsensiKelasIbuHamilRepository) FindActiveKehamilanByUserID(userID int32) (*models.Kehamilan, error) {
	var kehamilan models.Kehamilan

	err := r.db.
		Table("kehamilan AS k").
		Select("k.*").
		Joins("JOIN ibu AS i ON i.id = k.ibu_id").
		Joins("JOIN penduduk AS p ON p.id = i.penduduk_id").
		Joins("JOIN pengguna AS u ON u.penduduk_id = p.id").
		Where("u.id = ?", userID).
		Where("k.status_kehamilan IN ?", []string{"aktif", "TRIMESTER 1", "TRIMESTER 2", "TRIMESTER 3"}).
		Order("k.created_at DESC").
		First(&kehamilan).Error

	return &kehamilan, err
}

func (r *AbsensiKelasIbuHamilRepository) FindByKehamilanID(
	kehamilanID int32,
) ([]models.AbsensiKelasIbuHamil, error) {
	var list []models.AbsensiKelasIbuHamil

	err := r.db.
		Where("kehamilan_id = ?", kehamilanID).
		Order("pertemuan_ke ASC").
		Find(&list).Error

	return list, err
}

func (r *AbsensiKelasIbuHamilRepository) Upsert(
	data *models.AbsensiKelasIbuHamil,
) error {
	var existing models.AbsensiKelasIbuHamil

	err := r.db.
		Where("kehamilan_id = ? AND pertemuan_ke = ?",
			data.KehamilanID,
			data.PertemuanKe,
		).
		First(&existing).Error

	if err == nil {
		existing.Tanggal = data.Tanggal
		existing.NamaKader = data.NamaKader
		existing.TanggalParaf = data.TanggalParaf

		return r.db.Save(&existing).Error
	}

	if err == gorm.ErrRecordNotFound {
		return r.db.Create(data).Error
	}

	return err
}