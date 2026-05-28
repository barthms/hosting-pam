package repositories

import (
	"monitoring-service/app/models"
	"time"

	"gorm.io/gorm"
)

type ImunisasiRepository interface {
	Create(imunisasi *models.Imunisasi) error
	FindAllByAnakID(anakID int) ([]models.Imunisasi, error)
	FindByID(id uint) (*models.Imunisasi, error)
	Update(imunisasi *models.Imunisasi) error
	Delete(id uint) error
}

type imunisasiRepository struct {
	db *gorm.DB
}

func NewImunisasiRepository(db *gorm.DB) ImunisasiRepository {
	return &imunisasiRepository{db}
}

func (r *imunisasiRepository) Create(imunisasi *models.Imunisasi) error {
	return r.db.Create(imunisasi).Error
}

func (r *imunisasiRepository) FindAllByAnakID(anakID int) ([]models.Imunisasi, error) {
	var imunisasis []models.Imunisasi
	// Menggunakan Preload untuk mengambil data Master Vaksinnya juga
	err := r.db.Where("anak_id = ? AND is_deleted IS NULL", anakID).
		Preload("Imunisasi").
		Find(&imunisasis).Error
	return imunisasis, err
}

func (r *imunisasiRepository) FindByID(id uint) (*models.Imunisasi, error) {
	var imunisasi models.Imunisasi
	err := r.db.Where("id = ? AND is_deleted IS NULL", id).
		Preload("Imunisasi").
		Preload("Anak.Kependudukan"). // Mengambil data anak beserta data kependudukannya
		First(&imunisasi).Error
	return &imunisasi, err
}

func (r *imunisasiRepository) Update(imunisasi *models.Imunisasi) error {
	return r.db.Save(imunisasi).Error
}

func (r *imunisasiRepository) Delete(id uint) error {
	// Karena di model Anda Isdeleted menggunakan time.Time, kita lakukan soft delete manual
	return r.db.Model(&models.Imunisasi{}).Where("id = ?", id).Update("is_deleted", time.Now()).Error
}
