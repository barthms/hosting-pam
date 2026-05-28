package repositories

import (
	"errors"
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type KategoriTandaBahayaRepository struct {
	db *gorm.DB
}

func NewKategoriTandaBahayaRepository(db *gorm.DB) *KategoriTandaBahayaRepository {
	return &KategoriTandaBahayaRepository{db: db}
}

func (r *KategoriTandaBahayaRepository) Create(k *models.KategoriTandaBahaya) error {
	return r.db.Create(k).Error
}

func (r *KategoriTandaBahayaRepository) FindByID(id int32) (*models.KategoriTandaBahaya, error) {
	var k models.KategoriTandaBahaya
	err := r.db.First(&k, id).Error
	return &k, err
}

func (r *KategoriTandaBahayaRepository) FindAll() ([]models.KategoriTandaBahaya, error) {
	var list []models.KategoriTandaBahaya
	err := r.db.Find(&list).Error
	return list, err
}

func (r *KategoriTandaBahayaRepository) FindByTipeAndKategoriUsia(tipe, kategoriUsia string) ([]models.KategoriTandaBahaya, error) {
	var list []models.KategoriTandaBahaya
	err := r.db.Where("tipe_lembar = ? AND kategori_usia = ?", tipe, kategoriUsia).Find(&list).Error
	return list, err
}

func (r *KategoriTandaBahayaRepository) Update(k *models.KategoriTandaBahaya) error {
	return r.db.Save(k).Error
}

func (r *KategoriTandaBahayaRepository) Delete(id int32) error {
	result := r.db.Delete(&models.KategoriTandaBahaya{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("data kategori tanda bahaya tidak ditemukan")
	}
	return nil
}
