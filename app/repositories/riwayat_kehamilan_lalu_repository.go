package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type RiwayatKehamilanLaluRepository struct {
	db *gorm.DB
}

func NewRiwayatKehamilanLaluRepository(db *gorm.DB) *RiwayatKehamilanLaluRepository {
	return &RiwayatKehamilanLaluRepository{db: db}
}

func (r *RiwayatKehamilanLaluRepository) Create(rk *models.RiwayatKehamilanLalu) error {
	return r.db.Create(rk).Error
}

func (r *RiwayatKehamilanLaluRepository) FindByID(id int32) (*models.RiwayatKehamilanLalu, error) {
	var rk models.RiwayatKehamilanLalu
	err := r.db.First(&rk, id).Error
	return &rk, err
}

func (r *RiwayatKehamilanLaluRepository) FindByEvaluasiID(evaluasiID int32) ([]models.RiwayatKehamilanLalu, error) {
	var list []models.RiwayatKehamilanLalu
	err := r.db.Where("id_evaluasi = ?", evaluasiID).Find(&list).Error
	return list, err
}

func (r *RiwayatKehamilanLaluRepository) Update(rk *models.RiwayatKehamilanLalu) error {
	return r.db.Save(rk).Error
}

func (r *RiwayatKehamilanLaluRepository) Delete(id int32) error {
	return r.db.Delete(&models.RiwayatKehamilanLalu{}, id).Error
}
