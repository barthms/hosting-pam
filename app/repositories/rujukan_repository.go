package repositories

import (
	"errors"
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type RujukanRepository struct {
	db *gorm.DB
}

func NewRujukanRepository(db *gorm.DB) *RujukanRepository {
	return &RujukanRepository{db: db}
}

func (r *RujukanRepository) Create(rj *models.Rujukan) error {
	return r.db.Create(rj).Error
}

func (r *RujukanRepository) FindByID(id int32) (*models.Rujukan, error) {
	var rj models.Rujukan
	err := r.db.Preload("Kehamilan.Ibu.Kependudukan").First(&rj, id).Error
	return &rj, err
}

func (r *RujukanRepository) FindByKehamilanID(kehamilanID int32) ([]models.Rujukan, error) {
	var list []models.Rujukan
	err := r.db.Where("kehamilan_id = ?", kehamilanID).Find(&list).Error
	return list, err
}

func (r *RujukanRepository) Update(rj *models.Rujukan) error {
	return r.db.Save(rj).Error
	return r.db.Model(&models.Rujukan{}).Where("id = ?", rj.ID).Updates(rj).Error
}

func (r *RujukanRepository) Delete(id int32) error {
	result := r.db.Delete(&models.Rujukan{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("data rujukan tidak ditemukan")
	}
	return nil
}

// MODUL IBU
func (r *RujukanRepository) IsOwnedByUser(rujukanID int32, userID int32) (bool, error) {
	var count int64

	err := r.db.
		Table("rujukan rj").
		Joins("JOIN kehamilan k ON k.id = rj.kehamilan_id").
		Joins("JOIN ibu i ON i.id = k.ibu_id").
		Joins("JOIN penduduk p ON p.id = i.penduduk_id").
		Joins("JOIN pengguna u ON u.penduduk_id = p.id").
		Where("rj.id = ? AND u.id = ?", rujukanID, userID).
		Count(&count).Error

	return count > 0, err
}