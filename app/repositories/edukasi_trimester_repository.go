package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type EdukasiTrimesterRepository interface {
	Create(data *models.EdukasiTrimester) error
	FindAll() ([]models.EdukasiTrimester, error)
	FindByID(id int32) (*models.EdukasiTrimester, error)
	Update(data *models.EdukasiTrimester) error
	Delete(id int32) error
}

type edukasiTrimesterRepository struct {
	db *gorm.DB
}

func NewEdukasiTrimesterRepository(db *gorm.DB) EdukasiTrimesterRepository {
	return &edukasiTrimesterRepository{db}
}

func (r *edukasiTrimesterRepository) Create(data *models.EdukasiTrimester) error {
	return r.db.Create(data).Error
}

func (r *edukasiTrimesterRepository) FindAll() ([]models.EdukasiTrimester, error) {
	var result []models.EdukasiTrimester
	err := r.db.Find(&result).Error
	return result, err
}

func (r *edukasiTrimesterRepository) FindByID(id int32) (*models.EdukasiTrimester, error) {
	var data models.EdukasiTrimester
	err := r.db.First(&data, id).Error
	return &data, err
}

func (r *edukasiTrimesterRepository) Update(data *models.EdukasiTrimester) error {
	return r.db.Save(data).Error
}

func (r *edukasiTrimesterRepository) Delete(id int32) error {
	return r.db.Delete(&models.EdukasiTrimester{}, id).Error
}
