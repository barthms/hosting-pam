package repositories

import (
	"time"

	"monitoring-service/app/models"
	"monitoring-service/pkg/customerror"

	"gorm.io/gorm"
)

// PerawatanRepository defines the interface for perawatan data access
type PerawatanRepository interface {
	// Perawatan operations
	CreatePerawatan(perawatan *models.Perawatan) error
	GetPerawatanByID(id uint) (*models.Perawatan, error)
	GetPerawatanByAnakID(anakID int32) ([]models.Perawatan, error)
	GetPerawatanByAnakIDAndRentangUsia(anakID int32, rentangUsia string) ([]models.Perawatan, error)
	UpdatePerawatan(perawatan *models.Perawatan) error
	DeletePerawatan(id uint) error
	IsAnakExist(anakID int32) (bool, error)
	IsAnakOwnedByIbu(anakID int32, userID int32) (bool, error)
	GetPerawatanByAnakIDForIbu(anakID int32, userID int32) ([]models.Perawatan, error)
	GetPerawatanByAnakIDAndRentangUsiaForIbu(anakID int32, rentangUsia string, userID int32) ([]models.Perawatan, error)

	// KategoriCapaian operations
	GetAllKategoriCapaian() ([]models.KategoriCapaian, error)
	GetKategoriCapaianByRentangUsia(rentangUsia string) ([]models.KategoriCapaian, error)
	GetKategoriCapaianByID(id uint) (*models.KategoriCapaian, error)
}

// perawatanRepository is concrete implementation of PerawatanRepository
type perawatanRepository struct {
	db *gorm.DB
}

// NewPerawatanRepository creates and returns a new PerawatanRepository
func NewPerawatanRepository(db *gorm.DB) PerawatanRepository {
	return &perawatanRepository{db: db}
}

// ─────────────────────────────────────────────────────────
// PERAWATAN OPERATIONS
// ─────────────────────────────────────────────────────────

// CreatePerawatan creates a new perawatan record
func (r *perawatanRepository) CreatePerawatan(perawatan *models.Perawatan) error {
	if err := r.db.Create(perawatan).Error; err != nil {
		return customerror.NewInternalServiceError("gagal membuat data perawatan")
	}
	return nil
}

// GetPerawatanByID retrieves a single perawatan by ID
func (r *perawatanRepository) GetPerawatanByID(id uint) (*models.Perawatan, error) {
	var result models.Perawatan
	err := r.db.
		Preload("Anak").
		Preload("KategoriCapaian").
		Where("id = ? AND deleted_at IS NULL", id).
		First(&result).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, customerror.NewNotFoundError("data perawatan tidak ditemukan")
		}
		return nil, customerror.NewInternalServiceError("gagal mengambil data perawatan")
	}
	return &result, nil
}

// GetPerawatanByAnakID retrieves all perawatan records for a specific child
func (r *perawatanRepository) GetPerawatanByAnakID(anakID int32) ([]models.Perawatan, error) {
	var result []models.Perawatan
	if err := r.db.
		Preload("KategoriCapaian").
		Where("anak_id = ? AND deleted_at IS NULL", anakID).
		Order("tanggal_periksa DESC, created_at DESC").
		Find(&result).Error; err != nil {
		return nil, customerror.NewInternalServiceError("gagal mengambil data perawatan")
	}
	return result, nil
}

// GetPerawatanByAnakIDAndRentangUsia retrieves perawatan for a child with specific age range
func (r *perawatanRepository) GetPerawatanByAnakIDAndRentangUsia(anakID int32, rentangUsia string) ([]models.Perawatan, error) {
	var result []models.Perawatan
	if err := r.db.
		Joins("JOIN kategori_capaian ON kategori_capaian.id = perawatan.kategori_capaian_id").
		Preload("KategoriCapaian").
		Where("perawatan.anak_id = ? AND kategori_capaian.rentang_usia = ? AND perawatan.deleted_at IS NULL", anakID, rentangUsia).
		Order("perawatan.tanggal_periksa DESC, perawatan.created_at DESC").
		Find(&result).Error; err != nil {
		return nil, customerror.NewInternalServiceError("gagal mengambil data perawatan")
	}
	return result, nil
}

// UpdatePerawatan updates an existing perawatan record
func (r *perawatanRepository) UpdatePerawatan(perawatan *models.Perawatan) error {
	perawatan.UpdatedAt = time.Now()
	if err := r.db.Save(perawatan).Error; err != nil {
		return customerror.NewInternalServiceError("gagal mengubah data perawatan")
	}
	return nil
}

// DeletePerawatan soft deletes a perawatan record
func (r *perawatanRepository) DeletePerawatan(id uint) error {
	if err := r.db.Model(&models.Perawatan{}).
		Where("id = ?", id).
		Update("deleted_at", time.Now()).Error; err != nil {
		return customerror.NewInternalServiceError("gagal menghapus data perawatan")
	}
	return nil
}

// ─────────────────────────────────────────────────────────
// KATEGORI CAPAIAN OPERATIONS
// ─────────────────────────────────────────────────────────

// GetAllKategoriCapaian retrieves all kategori capaian records
func (r *perawatanRepository) GetAllKategoriCapaian() ([]models.KategoriCapaian, error) {
	var result []models.KategoriCapaian
	if err := r.db.
		Where("deleted_at IS NULL").
		Order("rentang_usia ASC, id ASC").
		Find(&result).Error; err != nil {
		return nil, customerror.NewInternalServiceError("gagal mengambil kategori capaian")
	}
	return result, nil
}

// GetKategoriCapaianByRentangUsia retrieves kategori capaian for a specific age range
func (r *perawatanRepository) GetKategoriCapaianByRentangUsia(rentangUsia string) ([]models.KategoriCapaian, error) {
	var result []models.KategoriCapaian
	if err := r.db.
		Where("rentang_usia = ? AND deleted_at IS NULL", rentangUsia).
		Order("id ASC").
		Find(&result).Error; err != nil {
		return nil, customerror.NewInternalServiceError("gagal mengambil kategori capaian")
	}
	return result, nil
}

// GetKategoriCapaianByID retrieves a specific kategori capaian by ID
func (r *perawatanRepository) GetKategoriCapaianByID(id uint) (*models.KategoriCapaian, error) {
	var result models.KategoriCapaian
	err := r.db.
		Where("id = ? AND deleted_at IS NULL", id).
		First(&result).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, customerror.NewNotFoundError("kategori capaian tidak ditemukan")
		}
		return nil, customerror.NewInternalServiceError("gagal mengambil kategori capaian")
	}
	return &result, nil
}

// ─────────────────────────────────────────────────────────
// ACCESS CONTROL OPERATIONS
// ─────────────────────────────────────────────────────────

// IsAnakExist checks if an anak exists
func (r *perawatanRepository) IsAnakExist(anakID int32) (bool, error) {
	var count int64
	err := r.db.
		Model(&models.Anak{}).
		Where("id = ? AND deleted_at IS NULL", anakID).
		Count(&count).Error

	if err != nil {
		return false, customerror.NewInternalServiceError("gagal memverifikasi anak")
	}

	return count > 0, nil
}

// IsAnakOwnedByIbu checks if a child belongs to a specific ibu (user)
func (r *perawatanRepository) IsAnakOwnedByIbu(anakID int32, userID int32) (bool, error) {
	var count int64
	err := r.db.
		Model(&models.Anak{}).
		Joins("JOIN kehamilan k ON k.id = anak.kehamilan_id").
		Joins("JOIN ibu i ON i.id = k.ibu_id").
		Joins("JOIN penduduk p_ibu ON p_ibu.id = i.penduduk_id").
		Joins("JOIN pengguna u ON u.penduduk_id = p_ibu.id").
		Where("anak.id = ? AND u.id = ? AND anak.deleted_at IS NULL", anakID, userID).
		Count(&count).Error

	if err != nil {
		return false, customerror.NewInternalServiceError("gagal memverifikasi ownership anak")
	}

	return count > 0, nil
}

// GetPerawatanByAnakIDForIbu retrieves perawatan for a child with ownership check
func (r *perawatanRepository) GetPerawatanByAnakIDForIbu(anakID int32, userID int32) ([]models.Perawatan, error) {
	// First, verify ownership
	owned, err := r.IsAnakOwnedByIbu(anakID, userID)
	if err != nil {
		return nil, err
	}

	if !owned {
		return nil, customerror.NewForbiddenError("Anda tidak memiliki akses ke anak ini")
	}

	// If owned, get perawatan
	var result []models.Perawatan
	if err := r.db.
		Preload("KategoriCapaian").
		Where("anak_id = ? AND deleted_at IS NULL", anakID).
		Order("tanggal_periksa DESC, created_at DESC").
		Find(&result).Error; err != nil {
		return nil, customerror.NewInternalServiceError("gagal mengambil data perawatan")
	}
	return result, nil
}

// GetPerawatanByAnakIDAndRentangUsiaForIbu retrieves perawatan with ownership check and age range filter
func (r *perawatanRepository) GetPerawatanByAnakIDAndRentangUsiaForIbu(anakID int32, rentangUsia string, userID int32) ([]models.Perawatan, error) {
	// First, verify ownership
	owned, err := r.IsAnakOwnedByIbu(anakID, userID)
	if err != nil {
		return nil, err
	}

	if !owned {
		return nil, customerror.NewForbiddenError("Anda tidak memiliki akses ke anak ini")
	}

	// If owned, get perawatan
	var result []models.Perawatan
	if err := r.db.
		Joins("JOIN kategori_capaian ON kategori_capaian.id = perawatan.kategori_capaian_id").
		Preload("KategoriCapaian").
		Where("perawatan.anak_id = ? AND kategori_capaian.rentang_usia = ? AND perawatan.deleted_at IS NULL", anakID, rentangUsia).
		Order("perawatan.tanggal_periksa DESC, perawatan.created_at DESC").
		Find(&result).Error; err != nil {
		return nil, customerror.NewInternalServiceError("gagal mengambil data perawatan")
	}
	return result, nil
}
