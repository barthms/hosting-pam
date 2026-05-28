package repositories

import (
	"monitoring-service/app/models"
	"time"

	"gorm.io/gorm"
)

type KunjunganImunisasiRepository interface {
	Create(k *models.KehadiranImunisasi) error
	GetByAnakID(anakID int32) ([]models.KehadiranImunisasi, error)
	GetByID(id int32) (*models.KehadiranImunisasi, error)
	GetAll() ([]models.KehadiranImunisasi, error)
	Update(id int32, req models.UpdateKunjunganImunisasiRequest, now time.Time) error
	Delete(id int32) error
}
type kunjunganImunisasiRepository struct {
	db *gorm.DB
}

func NewKunjunganImunisasiRepository(db *gorm.DB) KunjunganImunisasiRepository {
	return &kunjunganImunisasiRepository{db: db}
}
func (r *kunjunganImunisasiRepository) withTx(fn func(tx *gorm.DB) error) error {
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
func (r *kunjunganImunisasiRepository) Create(k *models.KehadiranImunisasi) error {
	return r.withTx(func(tx *gorm.DB) error {
		return tx.Create(k).Error
	})
}
func (r *kunjunganImunisasiRepository) GetByAnakID(anakID int32) ([]models.KehadiranImunisasi, error) {
	var data []models.KehadiranImunisasi

	err := r.db.
		Preload("Detail").
		Where("anak_id = ?", anakID).
		Find(&data).Error

	return data, err
}
func (r *kunjunganImunisasiRepository) GetByID(id int32) (*models.KehadiranImunisasi, error) {
	var data models.KehadiranImunisasi

	err := r.db.
		Preload("Detail").
		First(&data, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // lebih aman untuk service layer
		}
		return nil, err
	}

	return &data, nil
}
func (r *kunjunganImunisasiRepository) GetAll() ([]models.KehadiranImunisasi, error) {
	var data []models.KehadiranImunisasi

	err := r.db.
		Preload("Detail").
		Order("created_at DESC").
		Find(&data).Error

	return data, err
}
func (r *kunjunganImunisasiRepository) Update(id int32, req models.UpdateKunjunganImunisasiRequest, now time.Time) error {
	return r.withTx(func(tx *gorm.DB) error {

		// dynamic update (hindari overwrite kosong)
		updates := map[string]interface{}{
			"updated_at": now,
		}

		if req.AnakID != 0 {
			updates["anak_id"] = req.AnakID
		}

		if err := tx.Model(&models.KehadiranImunisasi{}).
			Where("id = ?", id).
			Updates(updates).Error; err != nil {
			return err
		}

		// delete old detail
		if err := tx.Where("kunjungan_imunisasi_id = ?", id).
			Delete(&models.DetailPelayananImunisasi{}).Error; err != nil {
			return err
		}

		// batch insert (lebih cepat)
		var details []models.DetailPelayananImunisasi
		for _, d := range req.Detail {
			details = append(details, models.DetailPelayananImunisasi{
				KunjunganImunisasiID: id,
				JenisPelayananID:     d.JenisPelayananID,
				Keterangan:           d.Keterangan,
				CreatedAt:            now,
				UpdatedAt:            now,
			})
		}

		if len(details) > 0 {
			if err := tx.Create(&details).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
func (r *kunjunganImunisasiRepository) Delete(id int32) error {
	return r.withTx(func(tx *gorm.DB) error {

		if err := tx.Where("kunjungan_imunisasi_id = ?", id).
			Delete(&models.DetailPelayananImunisasi{}).Error; err != nil {
			return err
		}

		if err := tx.Delete(&models.KehadiranImunisasi{}, id).Error; err != nil {
			return err
		}

		return nil
	})
}
