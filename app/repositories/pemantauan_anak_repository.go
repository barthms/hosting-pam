package repositories

import (
	"monitoring-service/app/models"
	"gorm.io/gorm"
)

type PemantauanAnakRepository interface {
	Create(record *models.LembarPemantauan) error
	Update(record *models.LembarPemantauan) error
	Delete(id uint) error
	FindByID(id uint) (*models.LembarPemantauan, error)
	FindByChildAndRange(anakID uint, rentangID uint) ([]models.LembarPemantauan, error)
	GetByPeriode(anakID uint, rentangID uint, periode int) (*models.LembarPemantauan, error)
	GetRentangUsia() ([]models.RentangUsia, error)
	GetKategoriByRentang(rentangID uint) ([]models.KategoriTandaSakit, error)
	CreateKategori(data *models.KategoriTandaSakit) error
	UpdateKategori(data *models.KategoriTandaSakit) error
	DeleteKategori(id uint) error
	UpdateStatus(id uint, status string, pemeriksa string) error
}

type pemantauanAnakRepository struct {
	db *gorm.DB
}

func NewPemantauanAnakRepository(db *gorm.DB) PemantauanAnakRepository {
	return &pemantauanAnakRepository{db}
}

func (r *pemantauanAnakRepository) Create(record *models.LembarPemantauan) error {
	return r.db.Create(record).Error
}

func (r *pemantauanAnakRepository) Update(record *models.LembarPemantauan) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("lembar_pemantauan_id = ?", record.ID).Delete(&models.DetailPemantauan{}).Error; err != nil {
			return err
		}
		return tx.Save(record).Error
	})
}

func (r *pemantauanAnakRepository) Delete(id uint) error {
	return r.db.Delete(&models.LembarPemantauan{}, id).Error
}

func (r *pemantauanAnakRepository) FindByID(id uint) (*models.LembarPemantauan, error) {
	var record models.LembarPemantauan
	err := r.db.Preload("DetailGejala.KategoriTandaSakit").First(&record, id).Error
	return &record, err
}

func (r *pemantauanAnakRepository) FindByChildAndRange(anakID uint, rentangID uint) ([]models.LembarPemantauan, error) {
	var records []models.LembarPemantauan
	err := r.db.Preload("DetailGejala.KategoriTandaSakit").
		Where("anak_id = ? AND rentang_usia_id = ?", anakID, rentangID).
		Order("periode_waktu asc").Find(&records).Error
	return records, err
}

func (r *pemantauanAnakRepository) GetByPeriode(anakID uint, rentangID uint, periode int) (*models.LembarPemantauan, error) {
	var record models.LembarPemantauan
	err := r.db.Where("anak_id = ? AND rentang_usia_id = ? AND periode_waktu = ?", anakID, rentangID, periode).First(&record).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *pemantauanAnakRepository) GetRentangUsia() ([]models.RentangUsia, error) {
	var data []models.RentangUsia
	err := r.db.Find(&data).Error
	return data, err
}

func (r *pemantauanAnakRepository) GetKategoriByRentang(rentangID uint) ([]models.KategoriTandaSakit, error) {
	var data []models.KategoriTandaSakit
	err := r.db.Where("rentang_usia_id = ?", rentangID).Find(&data).Error
	return data, err
}

func (r *pemantauanAnakRepository) CreateKategori(data *models.KategoriTandaSakit) error {
	return r.db.Create(data).Error
}

func (r *pemantauanAnakRepository) UpdateKategori(data *models.KategoriTandaSakit) error {
	return r.db.Save(data).Error
}

func (r *pemantauanAnakRepository) DeleteKategori(id uint) error {
	return r.db.Delete(&models.KategoriTandaSakit{}, id).Error
}

func (r *pemantauanAnakRepository) UpdateStatus(id uint, status string, pemeriksa string) error {
	return r.db.Model(&models.LembarPemantauan{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status":         status,
		"nama_pemeriksa": pemeriksa,
	}).Error
}
