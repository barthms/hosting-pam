package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type AturanPorsiMPASIRepository interface {
	Create(data *models.AturanPorsiMPASI) error
	GetAll() ([]models.AturanPorsiMPASI, error)
	GetByID(id int32) (*models.AturanPorsiMPASI, error)
	GetByAgeRange(bulanMin, bulanMax int) ([]models.AturanPorsiMPASI, error)
	Update(id int32, data *models.AturanPorsiMPASI) error
	Delete(id int32) error
}

type aturanPorsiMPASIRepository struct {
	db *gorm.DB
}

func NewAturanPorsiMPASIRepository(db *gorm.DB) AturanPorsiMPASIRepository {
	return &aturanPorsiMPASIRepository{db}
}

func (r *aturanPorsiMPASIRepository) Create(data *models.AturanPorsiMPASI) error {
	return r.db.Create(data).Error
}

func (r *aturanPorsiMPASIRepository) GetAll() ([]models.AturanPorsiMPASI, error) {
	var data []models.AturanPorsiMPASI
	err := r.db.Order("bulan_min asc, bulan_max asc").Find(&data).Error
	return data, err
}

func (r *aturanPorsiMPASIRepository) GetByID(id int32) (*models.AturanPorsiMPASI, error) {
	var data models.AturanPorsiMPASI
	err := r.db.First(&data, id).Error
	return &data, err
}

func (r *aturanPorsiMPASIRepository) GetByAgeRange(bulanMin, bulanMax int) ([]models.AturanPorsiMPASI, error) {
	var data []models.AturanPorsiMPASI
	err := r.db.Where("bulan_min = ? AND bulan_max = ?", bulanMin, bulanMax).Find(&data).Error
	return data, err
}

func (r *aturanPorsiMPASIRepository) Update(id int32, data *models.AturanPorsiMPASI) error {
	return r.db.Model(&models.AturanPorsiMPASI{}).Where("id = ?", id).Updates(data).Error
}

func (r *aturanPorsiMPASIRepository) Delete(id int32) error {
	return r.db.Delete(&models.AturanPorsiMPASI{}, id).Error
}
