package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type EdukasiTandaMelahirkanRepository interface {
	Create(data *models.EdukasiTandaMelahirkan) error
	FindAll() ([]models.EdukasiTandaMelahirkan, error)
	FindByID(id int32) (*models.EdukasiTandaMelahirkan, error)
	Update(data *models.EdukasiTandaMelahirkan) error
	Delete(id int32) error
}

type edukasiTandaMelahirkanRepository struct {
	db *gorm.DB
}

func NewEdukasiTandaMelahirkanRepository(db *gorm.DB) EdukasiTandaMelahirkanRepository {
	return &edukasiTandaMelahirkanRepository{db}
}

func (r *edukasiTandaMelahirkanRepository) Create(data *models.EdukasiTandaMelahirkan) error {
	return r.db.Create(data).Error
}

func (r *edukasiTandaMelahirkanRepository) FindAll() ([]models.EdukasiTandaMelahirkan, error) {
	var result []models.EdukasiTandaMelahirkan
	err := r.db.Find(&result).Error
	return result, err
}

func (r *edukasiTandaMelahirkanRepository) FindByID(id int32) (*models.EdukasiTandaMelahirkan, error) {
	var data models.EdukasiTandaMelahirkan
	err := r.db.First(&data, id).Error
	return &data, err
}

func (r *edukasiTandaMelahirkanRepository) Update(data *models.EdukasiTandaMelahirkan) error {
	return r.db.Save(data).Error
}

func (r *edukasiTandaMelahirkanRepository) Delete(id int32) error {
	return r.db.Delete(&models.EdukasiTandaMelahirkan{}, id).Error
}
