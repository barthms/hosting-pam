package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type KeluhanAnakRepository interface {
	Create(data *models.KeluhanAnak) error
	Update(data *models.KeluhanAnak) error
	Delete(id uint) error
	FindAllByAnakID(anakID uint) ([]models.KeluhanAnak, error)
	FindByID(id uint) (*models.KeluhanAnak, error)
	IsAnakMilikIbu(userID, anakID uint) (bool, error)
}

type keluhanAnakRepository struct {
	db *gorm.DB
}

func NewKeluhanAnakRepository(db *gorm.DB) KeluhanAnakRepository {
	return &keluhanAnakRepository{db}
}

func (r *keluhanAnakRepository) Create(data *models.KeluhanAnak) error {
	return r.db.Create(data).Error
}

func (r *keluhanAnakRepository) Update(data *models.KeluhanAnak) error {
	return r.db.Save(data).Error
}

func (r *keluhanAnakRepository) Delete(id uint) error {
	return r.db.Delete(&models.KeluhanAnak{}, id).Error
}

func (r *keluhanAnakRepository) FindAllByAnakID(anakID uint) ([]models.KeluhanAnak, error) {
	var data []models.KeluhanAnak
	err := r.db.Where("anak_id = ?", anakID).Where("deleted_at IS NULL").Order("tanggal desc").Find(&data).Error
	return data, err
}

func (r *keluhanAnakRepository) FindByID(id uint) (*models.KeluhanAnak, error) {
	var data models.KeluhanAnak
	err := r.db.Where("deleted_at IS NULL").First(&data, id).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *keluhanAnakRepository) IsAnakMilikIbu(userID, anakID uint) (bool, error) {
	var count int64
	err := r.db.Table("anak a").
		Joins("JOIN kehamilan k ON k.id = a.kehamilan_id").
		Joins("JOIN ibu i ON i.id = k.ibu_id").
		Joins("JOIN penduduk ki ON ki.id = i.penduduk_id").
		Joins("JOIN pengguna p ON p.penduduk_id = ki.id").
		Where("a.id = ?", anakID).
		Where("p.id = ?", userID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
