package repositories

import (
	"monitoring-service/app/models"
	"time"

	"gorm.io/gorm"
)

type PengukuranLilaRepository interface {
	Create(k *models.PengukuranLila) error
	GetByAnakID(anakID int32) ([]models.PengukuranLila, error)
	GetByID(id int32) (*models.PengukuranLila, error)
	GetAll() ([]models.PengukuranLila, error)
	Update(id int32, req models.UpdatePengukuranLilARequest, now time.Time) error
	Delete(id int32) error
	IsAnakMilikIbu(userID uint, anakID int32) (bool, error)
}
type pengukuranLilaRepository struct {
	db *gorm.DB
}


func NewPengukuranLilaRepository(db *gorm.DB) PengukuranLilaRepository {
	return &pengukuranLilaRepository{db: db}
}
func (r *pengukuranLilaRepository) withTx(fn func(tx *gorm.DB) error) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
func (r *pengukuranLilaRepository) Create(k *models.PengukuranLila) error {
	return r.withTx(func(tx *gorm.DB) error {
		return tx.Create(k).Error
	})
}
func (r *pengukuranLilaRepository) GetByAnakID(anakID int32) ([]models.PengukuranLila, error) {
	var data []models.PengukuranLila

	err := r.db.
		Where("anak_id = ?", anakID).
		Order("tanggal DESC").
		Find(&data).Error

	return data, err
}
func (r *pengukuranLilaRepository) GetByID(id int32) (*models.PengukuranLila, error) {
	var data models.PengukuranLila

	err := r.db.
		First(&data, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // lebih aman untuk service layer
		}
		return nil, err
	}

	return &data, nil
}
func (r *pengukuranLilaRepository) GetAll() ([]models.PengukuranLila, error) {
	var data []models.PengukuranLila

	err := r.db.
		Order("created_at DESC").
		Find(&data).Error

	return data, err
}
func (r *pengukuranLilaRepository) Update(id int32, req models.UpdatePengukuranLilARequest, now time.Time) error {
	return r.withTx(func(tx *gorm.DB) error {

		updates := map[string]interface{}{
			"updated_at": now,
		}

		if req.AnakID != 0 {
			updates["anak_id"] = req.AnakID
		}
		if req.Bulanke != 0 {
			updates["bulanke"] = req.Bulanke
		}
		if req.Tanggal != "" {
			updates["tanggal"] = req.Tanggal
		}
		if req.HasilLila != 0 {
			updates["hasil_lila"] = req.HasilLila
		}
		if req.KategoriRisiko != "" {
			updates["kategori_risiko"] = req.KategoriRisiko
		}


		return tx.Model(&models.PengukuranLila{}).
			Where("id = ?", id).
			Updates(updates).Error
	})
}
func (r *pengukuranLilaRepository) Delete(id int32) error {
	return r.withTx(func(tx *gorm.DB) error {
		return tx.Delete(&models.PengukuranLila{}, id).Error
	})
}

func (r *pengukuranLilaRepository) IsAnakMilikIbu(userID uint, anakID int32) (bool, error) {
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
		return false, err
	}
	return count > 0, nil
}