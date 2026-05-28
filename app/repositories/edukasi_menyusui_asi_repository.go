package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type EdukasiMenyusuiASIRepository interface {
	Create(data *models.EdukasiMenyusuiASI) error
	FindAll() ([]models.EdukasiMenyusuiASI, error)
	FindByID(id int32) (*models.EdukasiMenyusuiASI, error)
	Update(data *models.EdukasiMenyusuiASI) error
	Delete(id int32) error
}

type edukasiMenyusuiASIRepository struct {
	db *gorm.DB
}

func NewEdukasiMenyusuiASIRepository(db *gorm.DB) EdukasiMenyusuiASIRepository {
	return &edukasiMenyusuiASIRepository{db}
}

func (r *edukasiMenyusuiASIRepository) Create(data *models.EdukasiMenyusuiASI) error {
	return r.db.Create(data).Error
}

func (r *edukasiMenyusuiASIRepository) FindAll() ([]models.EdukasiMenyusuiASI, error) {
	var result []models.EdukasiMenyusuiASI
	err := r.db.Find(&result).Error
	return result, err
}

func (r *edukasiMenyusuiASIRepository) FindByID(id int32) (*models.EdukasiMenyusuiASI, error) {
	var data models.EdukasiMenyusuiASI
	err := r.db.First(&data, id).Error
	return &data, err
}

func (r *edukasiMenyusuiASIRepository) Update(data *models.EdukasiMenyusuiASI) error {
	return r.db.Save(data).Error
}

func (r *edukasiMenyusuiASIRepository) Delete(id int32) error {
	return r.db.Delete(&models.EdukasiMenyusuiASI{}, id).Error
}
