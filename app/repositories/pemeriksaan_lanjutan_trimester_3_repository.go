package repositories

import (
	"errors"
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type PemeriksaanLanjutanTrimester3Repository struct {
	db *gorm.DB
}

func NewPemeriksaanLanjutanTrimester3Repository(db *gorm.DB) *PemeriksaanLanjutanTrimester3Repository {
	return &PemeriksaanLanjutanTrimester3Repository{db: db}
}

func (r *PemeriksaanLanjutanTrimester3Repository) Create(p *models.PemeriksaanLanjutanTrimester3) error {
	return r.db.Create(p).Error
}

func (r *PemeriksaanLanjutanTrimester3Repository) FindByID(id int32) (*models.PemeriksaanLanjutanTrimester3, error) {
	var p models.PemeriksaanLanjutanTrimester3
	err := r.db.Preload("Kehamilan.Ibu.Kependudukan").First(&p, id).Error
	return &p, err
}

func (r *PemeriksaanLanjutanTrimester3Repository) FindByKehamilanID(kehamilanID int32) ([]models.PemeriksaanLanjutanTrimester3, error) {
	var list []models.PemeriksaanLanjutanTrimester3
	err := r.db.Where("kehamilan_id = ?", kehamilanID).Find(&list).Error
	return list, err
}

func (r *PemeriksaanLanjutanTrimester3Repository) Update(p *models.PemeriksaanLanjutanTrimester3) error {
	return r.db.Save(p).Error
}

func (r *PemeriksaanLanjutanTrimester3Repository) Delete(id int32) error {
	result := r.db.Delete(&models.PemeriksaanLanjutanTrimester3{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("data pemeriksaan lanjutan trimester 3 tidak ditemukan")
	}
	return nil
}
