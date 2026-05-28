package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type ResepMPASIRepository interface {
	Create(data *models.ResepMPASI) error
	GetAll() ([]models.ResepMPASI, error)
	GetByID(id int32) (*models.ResepMPASI, error)
	GetByAgeRange(bulanMin, bulanMax int) ([]models.ResepMPASI, error)
	GetByType(tipe string) ([]models.ResepMPASI, error)
	Update(id int32, data *models.ResepMPASI) error
	Delete(id int32) error
}

type resepMPASIRepository struct {
	db *gorm.DB
}

func NewResepMPASIRepository(db *gorm.DB) ResepMPASIRepository {
	return &resepMPASIRepository{db}
}

func (r *resepMPASIRepository) Create(data *models.ResepMPASI) error {
	return r.db.Create(data).Error
}

func (r *resepMPASIRepository) GetAll() ([]models.ResepMPASI, error) {
	var data []models.ResepMPASI
	err := r.db.Order("bulan_min asc, bulan_max asc, judul asc").Find(&data).Error
	return data, err
}

func (r *resepMPASIRepository) GetByID(id int32) (*models.ResepMPASI, error) {
	var data models.ResepMPASI
	err := r.db.First(&data, id).Error
	return &data, err
}

func (r *resepMPASIRepository) GetByAgeRange(bulanMin, bulanMax int) ([]models.ResepMPASI, error) {
	var data []models.ResepMPASI
	err := r.db.Where("bulan_min = ? AND bulan_max = ?", bulanMin, bulanMax).Order("judul asc").Find(&data).Error
	return data, err
}

func (r *resepMPASIRepository) GetByType(tipe string) ([]models.ResepMPASI, error) {
	var data []models.ResepMPASI
	err := r.db.Where("tipe = ?", tipe).Order("bulan_min asc, judul asc").Find(&data).Error
	return data, err
}

func (r *resepMPASIRepository) Update(id int32, data *models.ResepMPASI) error {
	return r.db.Model(&models.ResepMPASI{}).Where("id = ?", id).Updates(data).Error
}

func (r *resepMPASIRepository) Delete(id int32) error {
	return r.db.Delete(&models.ResepMPASI{}, id).Error
}
