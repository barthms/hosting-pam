package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type EdukasiIMDRepository interface {
	Create(data *models.EdukasiIMD) error
	FindAll() ([]models.EdukasiIMD, error)
	FindByID(id int32) (*models.EdukasiIMD, error)
	Update(data *models.EdukasiIMD) error
	Delete(id int32) error
}

type edukasiIMDRepository struct {
	db *gorm.DB
}

func NewEdukasiIMDRepository(db *gorm.DB) EdukasiIMDRepository {
	return &edukasiIMDRepository{db}
}

func (r *edukasiIMDRepository) Create(data *models.EdukasiIMD) error {
	return r.db.Create(data).Error
}

func (r *edukasiIMDRepository) FindAll() ([]models.EdukasiIMD, error) {
	var result []models.EdukasiIMD
	err := r.db.Find(&result).Error
	return result, err
}

func (r *edukasiIMDRepository) FindByID(id int32) (*models.EdukasiIMD, error) {
	var data models.EdukasiIMD
	err := r.db.First(&data, id).Error
	return &data, err
}

func (r *edukasiIMDRepository) Update(data *models.EdukasiIMD) error {
	return r.db.Save(data).Error
}

func (r *edukasiIMDRepository) Delete(id int32) error {
	return r.db.Delete(&models.EdukasiIMD{}, id).Error
}
