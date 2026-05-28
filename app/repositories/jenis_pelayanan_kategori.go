package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

// =====================
// INTERFACE
// =====================
type JenisPelayananRepository interface {
	GetByKategoriAndPeriode(kategoriID, periodeID int32) ([]models.JenisPelayanan, error)
}

// =====================
// IMPLEMENTATION
// =====================
type jenisPelayananRepository struct {
	db *gorm.DB
}

// =====================
// CONSTRUCTOR
// =====================
func NewJenisPelayananRepository(db *gorm.DB) JenisPelayananRepository {
	return &jenisPelayananRepository{db}
}

// =====================
// METHOD
// =====================
func (r *jenisPelayananRepository) GetByKategoriAndPeriode(kategoriID, periodeID int32) ([]models.JenisPelayanan, error) {
	var result []models.JenisPelayanan

	err := r.db.
		Table("jenis_pelayanan").
		Select("jenis_pelayanan.*").
		Joins("JOIN jenis_pelayanan_kategori jpk ON jpk.jenis_pelayanan_id = jenis_pelayanan.id").
		Where("jpk.kategori_umur_id = ? AND jpk.periode_id = ? AND jpk.is_active = ?", kategoriID, periodeID, true).
		Order("jpk.urutan ASC").
		Find(&result).Error

	return result, err
}