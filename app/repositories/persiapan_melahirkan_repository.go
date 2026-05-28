package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type PersiapanMelahirkanRepository struct {
	db *gorm.DB
}

func NewPersiapanMelahirkanRepository(db *gorm.DB) *PersiapanMelahirkanRepository {
	return &PersiapanMelahirkanRepository{db: db}
}

func (r *PersiapanMelahirkanRepository) FindActiveKehamilanByUserID(userID int32) (*models.Kehamilan, error) {
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

func (r *PersiapanMelahirkanRepository) FindByKehamilanID(kehamilanID int32) (*models.PersiapanMelahirkan, error) {
	var data models.PersiapanMelahirkan

	err := r.db.
		Where("kehamilan_id = ?", kehamilanID).
		First(&data).Error

	return &data, err
}

func (r *PersiapanMelahirkanRepository) Upsert(data *models.PersiapanMelahirkan) error {
	var existing models.PersiapanMelahirkan

	err := r.db.
		Where("kehamilan_id = ?", data.KehamilanID).
		First(&existing).Error

	if err == nil {
		existing.PerkiraanPersalinan = data.PerkiraanPersalinan
		existing.PendampingPersalinan = data.PendampingPersalinan
		existing.DanaPersalinan = data.DanaPersalinan
		existing.StatusJKN = data.StatusJKN
		existing.FaskesPersalinan = data.FaskesPersalinan
		existing.PendonorDarah = data.PendonorDarah
		existing.Transportasi = data.Transportasi
		existing.MetodeKB = data.MetodeKB
		existing.ProgramP4K = data.ProgramP4K
		existing.DokumenPenting = data.DokumenPenting

		return r.db.Save(&existing).Error
	}

	if err == gorm.ErrRecordNotFound {
		return r.db.Create(data).Error
	}

	return err
}