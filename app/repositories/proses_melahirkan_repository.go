package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type ProsesMelahirkanRepository struct {
	db *gorm.DB
}

func NewProsesMelahirkanRepository(db *gorm.DB) *ProsesMelahirkanRepository {
	return &ProsesMelahirkanRepository{db: db}
}

func (r *ProsesMelahirkanRepository) FindActiveKehamilanByUserID(userID int32) (*models.Kehamilan, error) {
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

func (r *ProsesMelahirkanRepository) FindByKehamilanID(kehamilanID int32) (*models.ProsesMelahirkan, error) {
	var data models.ProsesMelahirkan
	err := r.db.Where("kehamilan_id = ?", kehamilanID).First(&data).Error
	return &data, err
}

func (r *ProsesMelahirkanRepository) Upsert(data *models.ProsesMelahirkan) error {
	var existing models.ProsesMelahirkan

	err := r.db.Where("kehamilan_id = ?", data.KehamilanID).First(&existing).Error

	if err == nil {
		existing.TandaPersalinan = data.TandaPersalinan
		existing.ProsesMelahirkan = data.ProsesMelahirkan
		existing.HakIbuPendamping = data.HakIbuPendamping
		existing.HakIbuPosisiMelahirkan = data.HakIbuPosisiMelahirkan
		existing.Mulas = data.Mulas
		existing.TeknikMengurangiNyeri = data.TeknikMengurangiNyeri
		existing.IMDKontakKulit = data.IMDKontakKulit

		return r.db.Save(&existing).Error
	}

	if err == gorm.ErrRecordNotFound {
		return r.db.Create(data).Error
	}

	return err
}