package repositories

import (
	"monitoring-service/app/models"
	"time"

	"gorm.io/gorm"
)

type KunjunganVitaminRepository interface {
	Create(k *models.KunjunganVitamin) error
	GetByAnakID(anakID int32) ([]models.KunjunganVitamin, error)
	GetByID(id int32) (*models.KunjunganVitamin, error)
	GetAll() ([]models.KunjunganVitamin, error)
	Update(id int32, req models.UpdateKunjunganVitaminRequest, now time.Time) error
	Delete(id int32) error
}
type kunjunganVitaminRepository struct {
	db *gorm.DB
}

func NewKunjunganVitaminRepository(db *gorm.DB) KunjunganVitaminRepository {
	return &kunjunganVitaminRepository{db: db}
}
func (r *kunjunganVitaminRepository) withTx(fn func(tx *gorm.DB) error) error {
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
func (r *kunjunganVitaminRepository) Create(k *models.KunjunganVitamin) error {
	return r.withTx(func(tx *gorm.DB) error {
		return tx.Create(k).Error
	})
}
func (r *kunjunganVitaminRepository) GetByAnakID(anakID int32) ([]models.KunjunganVitamin, error) {
	var data []models.KunjunganVitamin

	err := r.db.
		Preload("Detail").
		Where("anak_id = ?", anakID).
		Order("tanggal DESC").
		Find(&data).Error

	return data, err
}
func (r *kunjunganVitaminRepository) GetByID(id int32) (*models.KunjunganVitamin, error) {
	var data models.KunjunganVitamin

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
func (r *kunjunganVitaminRepository) GetAll() ([]models.KunjunganVitamin, error) {
	var data []models.KunjunganVitamin

	err := r.db.
		Preload("Detail").
		Order("created_at DESC").
		Find(&data).Error

	return data, err
}
func (r *kunjunganVitaminRepository) Update(id int32, req models.UpdateKunjunganVitaminRequest, now time.Time) error {
	return r.withTx(func(tx *gorm.DB) error {

		// dynamic update (hindari overwrite kosong)
		updates := map[string]interface{}{
			"updated_at": now,
		}

		if req.AnakID != 0 {
			updates["anak_id"] = req.AnakID
		}
		if req.Tanggal != "" {
			updates["tanggal"] = req.Tanggal
		}

		if err := tx.Model(&models.KunjunganVitamin{}).
			Where("id = ?", id).
			Updates(updates).Error; err != nil {
			return err
		}

		// delete old detail
		if err := tx.Where("kunjungan_vitamin_id = ?", id).
			Delete(&models.DetailPelayananVitamin{}).Error; err != nil {
			return err
		}

		// batch insert (lebih cepat)
		var details []models.DetailPelayananVitamin
		for _, d := range req.Detail {
			details = append(details, models.DetailPelayananVitamin{
				KunjunganVitaminID: id,
				JenisPelayananID:   d.JenisPelayananID,
				Keterangan:         d.Keterangan,
				CreatedAt:          now,
				UpdatedAt:          now,
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
func (r *kunjunganVitaminRepository) Delete(id int32) error {
	return r.withTx(func(tx *gorm.DB) error {

		if err := tx.Where("kunjungan_vitamin_id = ?", id).
			Delete(&models.DetailPelayananVitamin{}).Error; err != nil {
			return err
		}

		if err := tx.Delete(&models.KunjunganVitamin{}, id).Error; err != nil {
			return err
		}

		return nil
	})
}
