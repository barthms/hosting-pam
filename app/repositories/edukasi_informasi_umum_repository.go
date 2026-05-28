package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type EdukasiInformasiUmumRepository interface {
	Create(data *models.EdukasiInformasiUmum) error
	FindAll() ([]models.EdukasiInformasiUmum, error)
	FindByID(id int32) (*models.EdukasiInformasiUmum, error)
	Update(data *models.EdukasiInformasiUmum) error
	Delete(id int32) error
}

type edukasiInformasiUmumRepository struct {
	db *gorm.DB
}

func NewEdukasiInformasiUmumRepository(db *gorm.DB) EdukasiInformasiUmumRepository {
	return &edukasiInformasiUmumRepository{db}
}

func (r *edukasiInformasiUmumRepository) Create(data *models.EdukasiInformasiUmum) error {
	return r.db.Create(data).Error
}

func (r *edukasiInformasiUmumRepository) FindAll() ([]models.EdukasiInformasiUmum, error) {
	var result []models.EdukasiInformasiUmum
	err := r.db.Find(&result).Error
	return result, err
}

func (r *edukasiInformasiUmumRepository) FindByID(id int32) (*models.EdukasiInformasiUmum, error) {
	var data models.EdukasiInformasiUmum
	err := r.db.First(&data, id).Error
	return &data, err
}

func (r *edukasiInformasiUmumRepository) Update(data *models.EdukasiInformasiUmum) error {
	return r.db.Save(data).Error
}

func (r *edukasiInformasiUmumRepository) Delete(id int32) error {
	return r.db.Delete(&models.EdukasiInformasiUmum{}, id).Error
}
