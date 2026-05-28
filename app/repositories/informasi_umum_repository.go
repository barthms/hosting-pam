package repositories

import (
	"errors"

	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type InformasiUmumRepository struct {
	db *gorm.DB
}

func NewInformasiUmumRepository(db *gorm.DB) *InformasiUmumRepository {
	return &InformasiUmumRepository{db: db}
}

func (r *InformasiUmumRepository) Create(data *models.InformasiUmum) error {
	return r.db.Create(data).Error
}

func (r *InformasiUmumRepository) FindByID(id int32) (*models.InformasiUmum, error) {
	var data models.InformasiUmum
	if err := r.db.First(&data, id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *InformasiUmumRepository) FindAll() ([]models.InformasiUmum, error) {
	var list []models.InformasiUmum
	err := r.db.Order("created_at DESC").Find(&list).Error
	return list, err
}

func (r *InformasiUmumRepository) Update(data *models.InformasiUmum) error {
	return r.db.Save(data).Error
}

func (r *InformasiUmumRepository) Delete(id int32) error {
	result := r.db.Delete(&models.InformasiUmum{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("data informasi umum tidak ditemukan")
	}
	return nil
}
