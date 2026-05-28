package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type EdukasiNifasRepository interface {
	Create(data *models.EdukasiNifas) error
	FindAll() ([]models.EdukasiNifas, error)
	FindByID(id int32) (*models.EdukasiNifas, error)
	Update(data *models.EdukasiNifas) error
	Delete(id int32) error
}

type edukasiNifasRepository struct {
	db *gorm.DB
}

func NewEdukasiNifasRepository(db *gorm.DB) EdukasiNifasRepository {
	return &edukasiNifasRepository{db}
}

func (r *edukasiNifasRepository) Create(data *models.EdukasiNifas) error {
	return r.db.Create(data).Error
}

func (r *edukasiNifasRepository) FindAll() ([]models.EdukasiNifas, error) {
	var result []models.EdukasiNifas
	err := r.db.Find(&result).Error
	return result, err
}

func (r *edukasiNifasRepository) FindByID(id int32) (*models.EdukasiNifas, error) {
	var data models.EdukasiNifas
	err := r.db.First(&data, id).Error
	return &data, err
}

func (r *edukasiNifasRepository) Update(data *models.EdukasiNifas) error {
	return r.db.Save(data).Error
}

func (r *edukasiNifasRepository) Delete(id int32) error {
	return r.db.Delete(&models.EdukasiNifas{}, id).Error
}
