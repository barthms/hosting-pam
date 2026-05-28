package repositories

import (
	"monitoring-service/app/models"
	"gorm.io/gorm"
)

type EdukasiPerawatanAnakRepository interface {
	Create(data *models.EdukasiPerawatanAnak) error
	Update(data *models.EdukasiPerawatanAnak) error
	Delete(id uint) error
	FindAll() ([]models.EdukasiPerawatanAnak, error)
	FindByID(id uint) (*models.EdukasiPerawatanAnak, error)
}

type edukasiPerawatanAnakRepository struct {
	db *gorm.DB
}

func NewEdukasiPerawatanAnakRepository(db *gorm.DB) EdukasiPerawatanAnakRepository {
	return &edukasiPerawatanAnakRepository{db}
}

func (r *edukasiPerawatanAnakRepository) Create(data *models.EdukasiPerawatanAnak) error {
	return r.db.Create(data).Error
}

func (r *edukasiPerawatanAnakRepository) Update(data *models.EdukasiPerawatanAnak) error {
	return r.db.Save(data).Error
}

func (r *edukasiPerawatanAnakRepository) Delete(id uint) error {
	return r.db.Delete(&models.EdukasiPerawatanAnak{}, id).Error
}

func (r *edukasiPerawatanAnakRepository) FindAll() ([]models.EdukasiPerawatanAnak, error) {
	var data []models.EdukasiPerawatanAnak
	err := r.db.Order("id desc").Find(&data).Error
	return data, err
}

func (r *edukasiPerawatanAnakRepository) FindByID(id uint) (*models.EdukasiPerawatanAnak, error) {
	var data models.EdukasiPerawatanAnak
	err := r.db.First(&data, id).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
