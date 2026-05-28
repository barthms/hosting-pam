package repositories

import (
	"errors"
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type CatatanPelayananTrimester2Repository struct {
	db *gorm.DB
}

func NewCatatanPelayananTrimester2Repository(db *gorm.DB) *CatatanPelayananTrimester2Repository {
	return &CatatanPelayananTrimester2Repository{db: db}
}

func (r *CatatanPelayananTrimester2Repository) Create(c *models.CatatanPelayananTrimester2) error {
	return r.db.Create(c).Error
}

func (r *CatatanPelayananTrimester2Repository) FindByID(id int32) (*models.CatatanPelayananTrimester2, error) {
	var c models.CatatanPelayananTrimester2
	err := r.db.Preload("Kehamilan.Ibu.Kependudukan").First(&c, id).Error
	return &c, err
}

func (r *CatatanPelayananTrimester2Repository) FindByKehamilanID(kehamilanID int32) ([]models.CatatanPelayananTrimester2, error) {
	var list []models.CatatanPelayananTrimester2
	err := r.db.Where("kehamilan_id = ?", kehamilanID).Order("tanggal_periksa_stamp_paraf DESC").Find(&list).Error
	return list, err
}

func (r *CatatanPelayananTrimester2Repository) Update(c *models.CatatanPelayananTrimester2) error {
	return r.db.Save(c).Error
}

func (r *CatatanPelayananTrimester2Repository) Delete(id int32) error {
	result := r.db.Delete(&models.CatatanPelayananTrimester2{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("data catatan pelayanan trimester 2 tidak ditemukan")
	}
	return nil
}
