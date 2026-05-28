package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type PemantauanIbuHamilRepository struct {
	db *gorm.DB
}

func NewPemantauanIbuHamilRepository(db *gorm.DB) *PemantauanIbuHamilRepository {
	return &PemantauanIbuHamilRepository{db: db}
}

func (r *PemantauanIbuHamilRepository) FindActiveKehamilanByUserID(userID int32) (*models.Kehamilan, error) {
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

func (r *PemantauanIbuHamilRepository) FindByKehamilanID(kehamilanID int32) ([]models.PemantauanIbuHamil, error) {
	var list []models.PemantauanIbuHamil

	err := r.db.
		Where("kehamilan_id = ?", kehamilanID).
		Order("minggu_kehamilan ASC").
		Find(&list).Error

	return list, err
}

func (r *PemantauanIbuHamilRepository) Upsert(data *models.PemantauanIbuHamil) error {
	var existing models.PemantauanIbuHamil

	err := r.db.
		Where("kehamilan_id = ? AND minggu_kehamilan = ?",
			data.KehamilanID,
			data.MingguKehamilan,
		).
		First(&existing).Error

	if err == nil {
		existing.DemamLebih2Hari = data.DemamLebih2Hari
		existing.SakitKepala = data.SakitKepala
		existing.CemasBerlebih = data.CemasBerlebih
		existing.ResikoTB = data.ResikoTB
		existing.GerakanBayiKurang = data.GerakanBayiKurang
		existing.NyeriPerut = data.NyeriPerut
		existing.CairanJalanLahir = data.CairanJalanLahir
		existing.MasalahKemaluan = data.MasalahKemaluan
		existing.DiareBerulang = data.DiareBerulang

		return r.db.Save(&existing).Error
	}

	if err == gorm.ErrRecordNotFound {
		return r.db.Create(data).Error
	}

	return err
}