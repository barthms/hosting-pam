package repositories

import (
	"errors"
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type SkriningPreeklampsiaRepository struct {
	db *gorm.DB
}

func NewSkriningPreeklampsiaRepository(db *gorm.DB) *SkriningPreeklampsiaRepository {
	return &SkriningPreeklampsiaRepository{db: db}
}

func (r *SkriningPreeklampsiaRepository) Create(s *models.SkriningPreeklampsia) error {
	return r.db.Create(s).Error
}

func (r *SkriningPreeklampsiaRepository) FindByID(id int32) (*models.SkriningPreeklampsia, error) {
	var s models.SkriningPreeklampsia
	err := r.db.Preload("Kehamilan.Ibu.Kependudukan").First(&s, id).Error
	return &s, err
}

func (r *SkriningPreeklampsiaRepository) FindByKehamilanID(kehamilanID int32) ([]models.SkriningPreeklampsia, error) {
	var list []models.SkriningPreeklampsia
	err := r.db.Where("kehamilan_id = ?", kehamilanID).Find(&list).Error
	return list, err
}

func (r *SkriningPreeklampsiaRepository) Update(s *models.SkriningPreeklampsia) error {
	return r.db.Save(s).Error
}

func (r *SkriningPreeklampsiaRepository) Delete(id int32) error {
	result := r.db.Delete(&models.SkriningPreeklampsia{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("data skrining preeklampsia tidak ditemukan")
	}
	return nil
}

// MODUL IBU
func (r *SkriningPreeklampsiaRepository) IsOwnedByUser(skriningID int32, userID int32) (bool, error) {
	var count int64

	err := r.db.
		Table("skrining_preeklampsia s").
		Joins("JOIN kehamilan k ON k.id = s.kehamilan_id").
		Joins("JOIN ibu i ON i.id = k.ibu_id").
		Joins("JOIN penduduk p ON p.id = i.penduduk_id").
		Joins("JOIN pengguna u ON u.penduduk_id = p.id").
		Where("s.id = ? AND u.id = ?", skriningID, userID).
		Count(&count).Error

	return count > 0, err
}

// // MODUL IBU (INTERNAL BACKUP ONLY)
// func (r *SkriningPreeklampsiaRepository) FindMineByUserID(userID int32) (*models.SkriningPreeklampsia, error) {
// 	var data models.SkriningPreeklampsia

// 	err := r.db.
// 		Table("skrining_preeklampsia s").
// 		Select("s.*").
// 		Joins("JOIN kehamilan k ON k.id = s.kehamilan_id").
// 		Joins("JOIN ibu i ON i.id = k.ibu_id").
// 		Joins("JOIN penduduk p ON p.id = i.penduduk_id").
// 		Joins("JOIN pengguna u ON u.id = p.id").
// 		Where("u.id = ?", userID).
// 		Order("s.id DESC").
// 		First(&data).Error

// 	return &data, err
// }

// MODUL IBU (SUPABASE UTAMA)
func (r *SkriningPreeklampsiaRepository) FindMineByUserID(userID int32) (*models.SkriningPreeklampsia, error) {
	var data models.SkriningPreeklampsia

	err := r.db.
		Table("skrining_preeklampsia s").
		Select("s.*").
		Joins("JOIN kehamilan k ON k.id = s.kehamilan_id").
		Joins("JOIN ibu i ON i.id = k.ibu_id").
		Joins("JOIN penduduk p ON p.id = i.penduduk_id").
		Joins("JOIN pengguna u ON u.penduduk_id = p.id").
		Where("u.id = ?", userID).
		Order("s.id DESC").
		First(&data).Error

	return &data, err
}