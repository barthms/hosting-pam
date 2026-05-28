package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type EdukasiSetelahMelahirkanRepository interface {
	Create(data *models.EdukasiSetelahMelahirkan) error
	FindAll() ([]models.EdukasiSetelahMelahirkan, error)
	FindByID(id int32) (*models.EdukasiSetelahMelahirkan, error)
	Update(data *models.EdukasiSetelahMelahirkan) error
	Delete(id int32) error
}

type edukasiSetelahMelahirkanRepository struct {
	db *gorm.DB
}

func NewEdukasiSetelahMelahirkanRepository(db *gorm.DB) EdukasiSetelahMelahirkanRepository {
	return &edukasiSetelahMelahirkanRepository{db}
}

func (r *edukasiSetelahMelahirkanRepository) Create(data *models.EdukasiSetelahMelahirkan) error {
	return r.db.Create(data).Error
}

func (r *edukasiSetelahMelahirkanRepository) FindAll() ([]models.EdukasiSetelahMelahirkan, error) {
	var result []models.EdukasiSetelahMelahirkan
	err := r.db.Find(&result).Error
	return result, err
}

func (r *edukasiSetelahMelahirkanRepository) FindByID(id int32) (*models.EdukasiSetelahMelahirkan, error) {
	var data models.EdukasiSetelahMelahirkan
	err := r.db.First(&data, id).Error
	return &data, err
}

func (r *edukasiSetelahMelahirkanRepository) Update(data *models.EdukasiSetelahMelahirkan) error {
	return r.db.Save(data).Error
}

func (r *edukasiSetelahMelahirkanRepository) Delete(id int32) error {
	return r.db.Delete(&models.EdukasiSetelahMelahirkan{}, id).Error
}
