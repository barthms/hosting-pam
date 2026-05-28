package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type JadwalHarianMPASIRepository interface {
	Create(data *models.JadwalHarianMPASI) error
	GetAll() ([]models.JadwalHarianMPASI, error)
	GetByID(id int32) (*models.JadwalHarianMPASI, error)
	GetByAgeRange(bulanMin, bulanMax int) ([]models.JadwalHarianMPASI, error)
	Update(id int32, data *models.JadwalHarianMPASI) error
	Delete(id int32) error
}

type jadwalHarianMPASIRepository struct {
	db *gorm.DB
}

func NewJadwalHarianMPASIRepository(db *gorm.DB) JadwalHarianMPASIRepository {
	return &jadwalHarianMPASIRepository{db}
}

func (r *jadwalHarianMPASIRepository) Create(data *models.JadwalHarianMPASI) error {
	return r.db.Create(data).Error
}

func (r *jadwalHarianMPASIRepository) GetAll() ([]models.JadwalHarianMPASI, error) {
	var data []models.JadwalHarianMPASI
	err := r.db.Order("bulan_min asc, bulan_max asc, waktu asc").Find(&data).Error
	return data, err
}

func (r *jadwalHarianMPASIRepository) GetByID(id int32) (*models.JadwalHarianMPASI, error) {
	var data models.JadwalHarianMPASI
	err := r.db.First(&data, id).Error
	return &data, err
}

func (r *jadwalHarianMPASIRepository) GetByAgeRange(bulanMin, bulanMax int) ([]models.JadwalHarianMPASI, error) {
	var data []models.JadwalHarianMPASI
	err := r.db.Where("bulan_min = ? AND bulan_max = ?", bulanMin, bulanMax).Order("waktu asc").Find(&data).Error
	return data, err
}

func (r *jadwalHarianMPASIRepository) Update(id int32, data *models.JadwalHarianMPASI) error {
	return r.db.Model(&models.JadwalHarianMPASI{}).Where("id = ?", id).Updates(data).Error
}

func (r *jadwalHarianMPASIRepository) Delete(id int32) error {
	return r.db.Delete(&models.JadwalHarianMPASI{}, id).Error
}
