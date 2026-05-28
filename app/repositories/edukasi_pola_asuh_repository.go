package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type EdukasiPolaAsuhRepository interface {
	Create(data *models.EdukasiPolaAsuh) error
	FindAll() ([]models.EdukasiPolaAsuh, error)
	FindByID(id int32) (*models.EdukasiPolaAsuh, error)
	Update(data *models.EdukasiPolaAsuh) error
	Delete(id int32) error
}

type edukasiPolaAsuhRepository struct {
	db *gorm.DB
}

func NewEdukasiPolaAsuhRepository(db *gorm.DB) EdukasiPolaAsuhRepository {
	return &edukasiPolaAsuhRepository{db}
}

func (r *edukasiPolaAsuhRepository) Create(data *models.EdukasiPolaAsuh) error {
	return r.db.Create(data).Error
}

func (r *edukasiPolaAsuhRepository) FindAll() ([]models.EdukasiPolaAsuh, error) {
	var result []models.EdukasiPolaAsuh
	err := r.db.Find(&result).Error
	return result, err
}

func (r *edukasiPolaAsuhRepository) FindByID(id int32) (*models.EdukasiPolaAsuh, error) {
	var data models.EdukasiPolaAsuh
	err := r.db.First(&data, id).Error
	return &data, err
}

func (r *edukasiPolaAsuhRepository) Update(data *models.EdukasiPolaAsuh) error {
	return r.db.Save(data).Error
}

func (r *edukasiPolaAsuhRepository) Delete(id int32) error {
	return r.db.Delete(&models.EdukasiPolaAsuh{}, id).Error
}
