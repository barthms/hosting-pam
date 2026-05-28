package repositories

import (
	"monitoring-service/app/models"
	"time"

	"gorm.io/gorm"
)

type PemeriksaanGigiRepository interface {
	Create(k *models.PeriksaGigi) error
	GetByAnakID(anakID int32) ([]models.PeriksaGigi, error)
	GetByID(id int32) (*models.PeriksaGigi, error)
	GetAll() ([]models.PeriksaGigi, error)
	Update(id int32, req models.UpdatePemeriksaanGigiRequest, now time.Time) error
	Delete(id int32) error
	IsAnakMilikIbu(userID uint, anakID int32) (bool, error)
}
type pemeriksaanGigiRepository struct {
	db *gorm.DB
}


func NewPemeriksaanGigiRepository(db *gorm.DB) PemeriksaanGigiRepository {
	return &pemeriksaanGigiRepository{db: db}
}
func (r *pemeriksaanGigiRepository) withTx(fn func(tx *gorm.DB) error) error {
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
func (r *pemeriksaanGigiRepository) Create(k *models.PeriksaGigi) error {
	return r.withTx(func(tx *gorm.DB) error {
		return tx.Create(k).Error
	})
}
func (r *pemeriksaanGigiRepository) GetByAnakID(anakID int32) ([]models.PeriksaGigi, error) {
	var data []models.PeriksaGigi

	err := r.db.
		Where("anak_id = ?", anakID).
		Order("tanggal DESC").
		Find(&data).Error

	return data, err
}
func (r *pemeriksaanGigiRepository) GetByID(id int32) (*models.PeriksaGigi, error) {
	var data models.PeriksaGigi

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
func (r *pemeriksaanGigiRepository) GetAll() ([]models.PeriksaGigi, error) {
	var data []models.PeriksaGigi

	err := r.db.
		Order("created_at DESC").
		Find(&data).Error

	return data, err
}
func (r *pemeriksaanGigiRepository) Update(id int32, req models.UpdatePemeriksaanGigiRequest, now time.Time) error {
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
		if req.Jumlahgigi != 0 {
			updates["jumlahgigi"] = req.Jumlahgigi
		}
		if req.GigiBerlubang != 0 {
			updates["gigi_berlubang"] = req.GigiBerlubang
		}
		if req.StatusPlak != "" {
			updates["status_plak"] = req.StatusPlak
		}
		if req.ResikoGigiBerlubang != "" {
			updates["resiko_gigi_berlubang"] = req.ResikoGigiBerlubang
		}

		return tx.Model(&models.PeriksaGigi{}).
			Where("id = ?", id).
			Updates(updates).Error
	})
}
func (r *pemeriksaanGigiRepository) Delete(id int32) error {
	return r.withTx(func(tx *gorm.DB) error {
		return tx.Delete(&models.PeriksaGigi{}, id).Error
	})
}

func (r *pemeriksaanGigiRepository) IsAnakMilikIbu(userID uint, anakID int32) (bool, error) {
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