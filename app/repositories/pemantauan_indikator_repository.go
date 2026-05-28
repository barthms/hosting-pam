package repositories

import (
	"errors"
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type PemantauanIndikatorRepository struct {
	db *gorm.DB
}

func NewPemantauanIndikatorRepository(db *gorm.DB) *PemantauanIndikatorRepository {
	return &PemantauanIndikatorRepository{db: db}
}

func (r *PemantauanIndikatorRepository) Create(data *models.PemantauanIndikator) error {
	return r.db.Create(data).Error
}

func (r *PemantauanIndikatorRepository) GetAll(kategoriUsia, q string) ([]models.PemantauanIndikator, error) {
	var list []models.PemantauanIndikator
	query := r.db.Model(&models.PemantauanIndikator{})

	if kategoriUsia != "" {
		query = query.Where("kategori_usia = ?", kategoriUsia)
	}
	if q != "" {
		query = query.Where("LOWER(deskripsi) LIKE LOWER(?)", "%"+q+"%")
	}

	err := query.Order("id ASC").Find(&list).Error
	return list, err
}

func (r *PemantauanIndikatorRepository) FindByID(id int32) (*models.PemantauanIndikator, error) {
	var data models.PemantauanIndikator
	err := r.db.First(&data, id).Error
	return &data, err
}

func (r *PemantauanIndikatorRepository) Update(data *models.PemantauanIndikator) error {
	return r.db.Save(data).Error
}

func (r *PemantauanIndikatorRepository) Delete(id int32) error {
	result := r.db.Delete(&models.PemantauanIndikator{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("indikator pemantauan tidak ditemukan")
	}
	return nil
}
