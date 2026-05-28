package repositories

import (
	"errors"
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type PemeriksaanLaboratoriumJiwaRepository struct {
	db *gorm.DB
}

func NewPemeriksaanLaboratoriumJiwaRepository(db *gorm.DB) *PemeriksaanLaboratoriumJiwaRepository {
	return &PemeriksaanLaboratoriumJiwaRepository{db: db}
}

func (r *PemeriksaanLaboratoriumJiwaRepository) Create(p *models.PemeriksaanLaboratoriumJiwa) error {
	return r.db.Create(p).Error
}

func (r *PemeriksaanLaboratoriumJiwaRepository) FindByID(id int32) (*models.PemeriksaanLaboratoriumJiwa, error) {
	var p models.PemeriksaanLaboratoriumJiwa
	err := r.db.Preload("Kehamilan.Ibu.Kependudukan").First(&p, id).Error
	return &p, err
}

func (r *PemeriksaanLaboratoriumJiwaRepository) FindByKehamilanID(kehamilanID int32) ([]models.PemeriksaanLaboratoriumJiwa, error) {
	var list []models.PemeriksaanLaboratoriumJiwa
	err := r.db.Where("kehamilan_id = ?", kehamilanID).Find(&list).Error
	return list, err
}

func (r *PemeriksaanLaboratoriumJiwaRepository) FindByKehamilanIDAndTrimester(kehamilanID, trimester int32) (*models.PemeriksaanLaboratoriumJiwa, error) {
	var p models.PemeriksaanLaboratoriumJiwa
	err := r.db.Where("kehamilan_id = ? AND trimester = ?", kehamilanID, trimester).First(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *PemeriksaanLaboratoriumJiwaRepository) Update(p *models.PemeriksaanLaboratoriumJiwa) error {
	return r.db.Save(p).Error
}

func (r *PemeriksaanLaboratoriumJiwaRepository) Delete(id int32) error {
	result := r.db.Delete(&models.PemeriksaanLaboratoriumJiwa{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("data pemeriksaan laboratorium jiwa tidak ditemukan")
	}
	return nil
}
