package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type KesehatanLingkunganRepository interface {
	// Kategori & Indikator (Master)
	CreateKategori(kategori *models.KategoriLingkungan) error
	GetAllKategori() ([]models.KategoriLingkungan, error)
	GetKategoriByID(id int32) (*models.KategoriLingkungan, error)
	UpdateKategori(kategori *models.KategoriLingkungan) error
	DeleteKategori(id int32) error

	CreateIndikator(indikator *models.IndikatorLingkungan) error
	UpdateIndikator(indikator *models.IndikatorLingkungan) error
	DeleteIndikator(id int32) error

	// Lembar & Detail (Transaksi)
	SaveLembar(lembar *models.LembarLingkungan) error
	GetLembarHistory(ibuID int32) ([]models.LembarLingkungan, error)
	GetLembarByID(id int32) (*models.LembarLingkungan, error)
	DeleteLembar(id int32) error
}

type kesehatanLingkunganRepository struct {
	db *gorm.DB
}

func NewKesehatanLingkunganRepository(db *gorm.DB) KesehatanLingkunganRepository {
	return &kesehatanLingkunganRepository{db}
}

func (r *kesehatanLingkunganRepository) CreateKategori(kategori *models.KategoriLingkungan) error {
	return r.db.Create(kategori).Error
}

func (r *kesehatanLingkunganRepository) GetAllKategori() ([]models.KategoriLingkungan, error) {
	var kategori []models.KategoriLingkungan
	err := r.db.Preload("Indikator").Find(&kategori).Error
	return kategori, err
}

func (r *kesehatanLingkunganRepository) GetKategoriByID(id int32) (*models.KategoriLingkungan, error) {
	var kategori models.KategoriLingkungan
	err := r.db.Preload("Indikator").First(&kategori, id).Error
	return &kategori, err
}

func (r *kesehatanLingkunganRepository) UpdateKategori(kategori *models.KategoriLingkungan) error {
	return r.db.Save(kategori).Error
}

func (r *kesehatanLingkunganRepository) DeleteKategori(id int32) error {
	return r.db.Delete(&models.KategoriLingkungan{}, id).Error
}

func (r *kesehatanLingkunganRepository) CreateIndikator(indikator *models.IndikatorLingkungan) error {
	return r.db.Create(indikator).Error
}

func (r *kesehatanLingkunganRepository) UpdateIndikator(indikator *models.IndikatorLingkungan) error {
	return r.db.Save(indikator).Error
}

func (r *kesehatanLingkunganRepository) DeleteIndikator(id int32) error {
	return r.db.Delete(&models.IndikatorLingkungan{}, id).Error
}

func (r *kesehatanLingkunganRepository) SaveLembar(lembar *models.LembarLingkungan) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(lembar).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *kesehatanLingkunganRepository) GetLembarHistory(ibuID int32) ([]models.LembarLingkungan, error) {
	var history []models.LembarLingkungan
	query := r.db.Preload("DetailJawaban.Indikator").Order("tanggal_periksa DESC")
	if ibuID > 0 {
		query = query.Where("ibu_id = ?", ibuID)
	}
	err := query.Find(&history).Error
	return history, err
}

func (r *kesehatanLingkunganRepository) GetLembarByID(id int32) (*models.LembarLingkungan, error) {
	var lembar models.LembarLingkungan
	err := r.db.Preload("DetailJawaban.Indikator").Preload("Ibu").First(&lembar, id).Error
	return &lembar, err
}

func (r *kesehatanLingkunganRepository) DeleteLembar(id int32) error {
	return r.db.Delete(&models.LembarLingkungan{}, id).Error
}
