package repositories

import (
	"monitoring-service/app/models"
	"time"

	"gorm.io/gorm"
)

type CatatanPelayananRepository interface {
	Create(c *models.CatatanPelayanan) error
	GetByAnakID(anakID int32) ([]models.CatatanPelayanan, error)
	GetByID(id int32) (*models.CatatanPelayanan, error)
	GetAll() ([]models.CatatanPelayanan, error)
	Update(id int32, req models.UpdateCatatanPelayananRequest, now time.Time) error
	Delete(id int32) error
}
type catatanpelayananRepository struct {
	db *gorm.DB
}


func NewCatatanPelayananRepository(db *gorm.DB) CatatanPelayananRepository {
	return &catatanpelayananRepository{db: db}
}
func (r *catatanpelayananRepository) withTx(fn func(tx *gorm.DB) error) error {
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
func (r *catatanpelayananRepository) Create(c *models.CatatanPelayanan) error {
	return r.withTx(func(tx *gorm.DB) error {
		return tx.Create(c).Error
	})
}
func (r *catatanpelayananRepository) GetByAnakID(anakID int32) ([]models.CatatanPelayanan, error) {
	var data []models.CatatanPelayanan

	err := r.db.
		Where("anak_id = ?", anakID).
		Order("tanggal_periksa DESC").
		Find(&data).Error

	return data, err
}
func (r *catatanpelayananRepository) GetByID(id int32) (*models.CatatanPelayanan, error) {
	var data models.CatatanPelayanan

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
func (r *catatanpelayananRepository) GetAll() ([]models.CatatanPelayanan, error) {
	var data []models.CatatanPelayanan

	err := r.db.
		Order("created_at DESC").
		Find(&data).Error

	return data, err
}
func (r *catatanpelayananRepository) Update(id int32, req models.UpdateCatatanPelayananRequest, now time.Time) error {
	return r.withTx(func(tx *gorm.DB) error {

		updates := map[string]interface{}{
			"updated_at": now,
		}

		if req.AnakID != 0 {
			updates["anak_id"] = req.AnakID
		}
		if req.TenagaKesehatanID != 0 {
			updates["tenaga_kesehatan_id"] = req.TenagaKesehatanID
		}
		if !req.TanggalPeriksa.IsZero() {
			updates["tanggal_periksa"] = req.TanggalPeriksa
		}
		if !req.TanggalKembali.IsZero() {
			updates["tanggal_kembali"] = req.TanggalKembali
		}
		if req.CatatanPelayanan != "" {
			updates["catatan_pelayanan"] = req.CatatanPelayanan
		}


		return tx.Model(&models.CatatanPelayanan{}).
			Where("id = ?", id).
			Updates(updates).Error
	})
}
func (r *catatanpelayananRepository) Delete(id int32) error {
	return r.withTx(func(tx *gorm.DB) error {
		return tx.Delete(&models.CatatanPelayanan{}, id).Error
	})
}