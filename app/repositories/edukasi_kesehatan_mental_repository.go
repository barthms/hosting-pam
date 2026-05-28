package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type EdukasiKesehatanMentalRepository interface {
	Create(data *models.EdukasiKesehatanMental) error
	FindAll() ([]models.EdukasiKesehatanMental, error)
	FindByID(id int32) (*models.EdukasiKesehatanMental, error)
	Update(data *models.EdukasiKesehatanMental) error
	Delete(id int32) error
}

type edukasiKesehatanMentalRepository struct {
	db *gorm.DB
}

func NewEdukasiKesehatanMentalRepository(db *gorm.DB) EdukasiKesehatanMentalRepository {
	return &edukasiKesehatanMentalRepository{db}
}

func (r *edukasiKesehatanMentalRepository) Create(data *models.EdukasiKesehatanMental) error {
	return r.db.Create(data).Error
}

func (r *edukasiKesehatanMentalRepository) FindAll() ([]models.EdukasiKesehatanMental, error) {
	var result []models.EdukasiKesehatanMental
	err := r.db.Find(&result).Error
	return result, err
}

func (r *edukasiKesehatanMentalRepository) FindByID(id int32) (*models.EdukasiKesehatanMental, error) {
	var data models.EdukasiKesehatanMental
	err := r.db.First(&data, id).Error
	return &data, err
}

func (r *edukasiKesehatanMentalRepository) Update(data *models.EdukasiKesehatanMental) error {
	return r.db.Save(data).Error
}

func (r *edukasiKesehatanMentalRepository) Delete(id int32) error {
	return r.db.Delete(&models.EdukasiKesehatanMental{}, id).Error
}
