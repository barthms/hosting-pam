package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type KategoriUmurRepository interface {
	FindAll() ([]models.KategoriUmur, error)
}

type kategoriUmurRepository struct {
	db *gorm.DB
}

func NewKategoriUmurRepository(db *gorm.DB) KategoriUmurRepository {
	return &kategoriUmurRepository{db}
}

func (r *kategoriUmurRepository) FindAll() ([]models.KategoriUmur, error) {
	var result []models.KategoriUmur
	err := r.db.Order("min_value asc, max_value asc").Find(&result).Error
	return result, err
}