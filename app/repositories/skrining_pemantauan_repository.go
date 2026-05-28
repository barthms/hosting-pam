package repositories

import (
	"errors"
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type SkriningPemantauanRepository struct {
	db *gorm.DB
}

func NewSkriningPemantauanRepository(db *gorm.DB) *SkriningPemantauanRepository {
	return &SkriningPemantauanRepository{db: db}
}

func (r *SkriningPemantauanRepository) Create(s *models.SkriningPemantauan) error {
	return r.db.Create(s).Error
}

func (r *SkriningPemantauanRepository) FindByID(id int32) (*models.SkriningPemantauan, error) {
	var s models.SkriningPemantauan
	err := r.db.Preload("Anak").Preload("KategoriTandaBahaya").First(&s, id).Error
	return &s, err
}

func (r *SkriningPemantauanRepository) FindByAnakID(anakID int32) ([]models.SkriningPemantauan, error) {
	var list []models.SkriningPemantauan
	err := r.db.Where("anak_id = ?", anakID).Preload("KategoriTandaBahaya").Find(&list).Error
	return list, err
}

func (r *SkriningPemantauanRepository) FindAll() ([]models.SkriningPemantauan, error) {
	var list []models.SkriningPemantauan
	err := r.db.Preload("Anak").Preload("KategoriTandaBahaya").Find(&list).Error
	return list, err
}

func (r *SkriningPemantauanRepository) Update(s *models.SkriningPemantauan) error {
	return r.db.Save(s).Error
}

func (r *SkriningPemantauanRepository) Delete(id int32) error {
	result := r.db.Delete(&models.SkriningPemantauan{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("data skrining pemantauan tidak ditemukan")
	}
	return nil
}
