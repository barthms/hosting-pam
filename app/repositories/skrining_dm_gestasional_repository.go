package repositories

import (
	"errors"
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type SkriningDMGestasionalRepository struct {
	db *gorm.DB
}

func NewSkriningDMGestasionalRepository(db *gorm.DB) *SkriningDMGestasionalRepository {
	return &SkriningDMGestasionalRepository{db: db}
}

func (r *SkriningDMGestasionalRepository) Create(s *models.SkriningDMGestasional) error {
	return r.db.Create(s).Error
}

func (r *SkriningDMGestasionalRepository) FindByID(id int32) (*models.SkriningDMGestasional, error) {
	var s models.SkriningDMGestasional
	err := r.db.Preload("Kehamilan.Ibu.Kependudukan").First(&s, id).Error
	return &s, err
}

func (r *SkriningDMGestasionalRepository) FindByKehamilanID(kehamilanID int32) ([]models.SkriningDMGestasional, error) {
	var list []models.SkriningDMGestasional
	err := r.db.Where("kehamilan_id = ?", kehamilanID).Find(&list).Error
	return list, err
}

func (r *SkriningDMGestasionalRepository) Update(s *models.SkriningDMGestasional) error {
	return r.db.Save(s).Error
}

func (r *SkriningDMGestasionalRepository) Delete(id int32) error {
	result := r.db.Delete(&models.SkriningDMGestasional{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("data skrining DM gestasional tidak ditemukan")
	}
	return nil
}
