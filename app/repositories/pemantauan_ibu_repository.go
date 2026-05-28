package repositories

import (
	"monitoring-service/app/models"
	"gorm.io/gorm"
)

type PemantauanIbuRepository interface {
	Create(record *models.LembarPemantauanIbu) error
	Update(record *models.LembarPemantauanIbu) error
	Delete(id int32) error
	FindByID(id int32) (*models.LembarPemantauanIbu, error)
	FindByIbu(ibuID int32) ([]models.LembarPemantauanIbu, error)
	GetByPeriode(ibuID int32, periode int) (*models.LembarPemantauanIbu, error)
	GetKategori() ([]models.KategoriPemantauanIbu, error)
}

type pemantauanIbuRepository struct {
	db *gorm.DB
}

func NewPemantauanIbuRepository(db *gorm.DB) PemantauanIbuRepository {
	return &pemantauanIbuRepository{db}
}

func (r *pemantauanIbuRepository) Create(record *models.LembarPemantauanIbu) error {
	return r.db.Create(record).Error
}

func (r *pemantauanIbuRepository) Update(record *models.LembarPemantauanIbu) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("lembar_pemantauan_ibu_id = ?", record.ID).Delete(&models.DetailPemantauanIbu{}).Error; err != nil {
			return err
		}
		return tx.Save(record).Error
	})
}

func (r *pemantauanIbuRepository) Delete(id int32) error {
	return r.db.Delete(&models.LembarPemantauanIbu{}, id).Error
}

func (r *pemantauanIbuRepository) FindByID(id int32) (*models.LembarPemantauanIbu, error) {
	var record models.LembarPemantauanIbu
	err := r.db.Preload("DetailGejala.KategoriPemantauan").First(&record, id).Error
	return &record, err
}

func (r *pemantauanIbuRepository) FindByIbu(ibuID int32) ([]models.LembarPemantauanIbu, error) {
	var records []models.LembarPemantauanIbu
	err := r.db.Preload("DetailGejala.KategoriPemantauan").
		Where("ibu_id = ?", ibuID).
		Order("periode_waktu asc").Find(&records).Error
	return records, err
}

func (r *pemantauanIbuRepository) GetByPeriode(ibuID int32, periode int) (*models.LembarPemantauanIbu, error) {
	var record models.LembarPemantauanIbu
	err := r.db.Where("ibu_id = ? AND periode_waktu = ?", ibuID, periode).First(&record).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *pemantauanIbuRepository) GetKategori() ([]models.KategoriPemantauanIbu, error) {
	var data []models.KategoriPemantauanIbu
	err := r.db.Find(&data).Error
	return data, err
}
