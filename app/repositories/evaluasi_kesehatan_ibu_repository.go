package repositories

import (
	"errors"
	"monitoring-service/app/models"

	"gorm.io/gorm"
	"log"
)

type EvaluasiKesehatanIbuRepository struct {
	db *gorm.DB
}

func NewEvaluasiKesehatanIbuRepository(db *gorm.DB) *EvaluasiKesehatanIbuRepository {
	return &EvaluasiKesehatanIbuRepository{db: db}
}

func (r *EvaluasiKesehatanIbuRepository) Create(e *models.EvaluasiKesehatanIbu) error {
	return r.db.Create(e).Error
}

func (r *EvaluasiKesehatanIbuRepository) FindByID(id int32) (*models.EvaluasiKesehatanIbu, error) {
	var e models.EvaluasiKesehatanIbu
	err := r.db.Preload("Kehamilan.Ibu.Kependudukan").First(&e, id).Error
	return &e, err
}

func (r *EvaluasiKesehatanIbuRepository) FindByKehamilanID(kehamilanID int32) ([]models.EvaluasiKesehatanIbu, error) {
	var list []models.EvaluasiKesehatanIbu
	err := r.db.Where("kehamilan_id = ?", kehamilanID).Find(&list).Error
	return list, err
}

func (r *EvaluasiKesehatanIbuRepository) Update(e *models.EvaluasiKesehatanIbu) error {
	return r.db.Save(e).Error
}

func (r *EvaluasiKesehatanIbuRepository) Delete(id int32) error {
	result := r.db.Delete(&models.EvaluasiKesehatanIbu{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("data evaluasi kesehatan ibu tidak ditemukan")
	}
	return nil
}

// MODUL IBU
func (r *EvaluasiKesehatanIbuRepository) IsOwnedByUser(evaluasiID int32, userID int32) (bool, error) {
	var count int64

	err := r.db.
		Table("evaluasi_kesehatan_ibu e").
		Joins("JOIN kehamilan k ON k.id = e.kehamilan_id").
		Joins("JOIN ibu i ON i.id = k.ibu_id").
		Joins("JOIN penduduk p ON p.id = i.penduduk_id").
		Joins("JOIN pengguna u ON u.penduduk_id = p.id").
		Where("e.id = ? AND u.id = ?", evaluasiID, userID).
		Count(&count).Error

	return count > 0, err
}

// MODUL IBU
// func (r *EvaluasiKesehatanIbuRepository) FindMineByUserID(userID int32) (*models.EvaluasiKesehatanIbu, error) {
// 	var e models.EvaluasiKesehatanIbu

// 	err := r.db.
// 		Table("evaluasi_kesehatan_ibu e").
// 		Joins("JOIN kehamilan k ON k.id = e.id").
// 		Joins("JOIN ibu i ON i.id = k.ibu_id").
// 		Joins("JOIN penduduk p ON p.id = i.id").
// 		Joins("JOIN pengguna u ON u.id = p.id").
// 		Where("u.id = ?", userID).
// 		Order("e.created_at DESC").
// 		First(&e).Error

// 	return &e, err
// }

// MODUL IBU (BACKUP ONLY)
// func (r *EvaluasiKesehatanIbuRepository) FindMineByUserID(userID int32) (*models.EvaluasiKesehatanIbu, error) {
// 	var e models.EvaluasiKesehatanIbu

// 	err := r.db.
// 		Table("evaluasi_kesehatan_ibu e").
// 		Select("e.*").
// 		Joins("JOIN kehamilan k ON k.id = e.kehamilan_id").
// 		Joins("JOIN ibu i ON i.id = k.ibu_id").
// 		Joins("JOIN penduduk p ON p.id = i.penduduk_id").
// 		Joins("JOIN pengguna u ON u.id = p.id"). // sementara, sesuai data dummy kamu
// 		Where("u.id = ?", userID).
// 		Order("e.created_at DESC").
// 		First(&e).Error

// 	return &e, err
// }



func (r *EvaluasiKesehatanIbuRepository) FindMineByUserID(userID int32) (*models.EvaluasiKesehatanIbu, error) {
	var e models.EvaluasiKesehatanIbu

	log.Println("DEBUG userID:", userID)

	err := r.db.Debug().
		Table("evaluasi_kesehatan_ibu e").
		Select("e.*").
		Joins("JOIN kehamilan k ON k.id = e.kehamilan_id").
		Joins("JOIN ibu i ON i.id = k.ibu_id").
		Joins("JOIN penduduk p ON p.id = i.penduduk_id").
		Joins("JOIN pengguna u ON u.id = p.id").
		Where("u.id = ?", userID).
		Order("e.created_at DESC").
		First(&e).Error

	log.Println("DEBUG err:", err)

	return &e, err
}