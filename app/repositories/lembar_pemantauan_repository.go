package repositories

import (
	"monitoring-service/app/models"
	"monitoring-service/pkg/customerror"
	"strings"
	"time"

	"gorm.io/gorm"
)

type LembarPemantauanRepository interface {
	Create(lembarPemantauan *models.LembarPemantauan) error
	FindByID(id uint) (*models.LembarPemantauan, error)
	FindByAnakID(anakID uint) ([]models.LembarPemantauan, error)
	FindAll() ([]models.LembarPemantauan, error)
	FindRentangUsiaByID(id uint) (*models.RentangUsia, error)
	FindRentangUsia() ([]models.RentangUsia, error)
	FindKategoriTandaSakitByRentangUsiaID(rentangUsiaID uint) ([]models.KategoriTandaSakit, error)
	IsAnakMilikIbu(userID, anakID uint) (bool, error)
	Update(lembarPemantauan *models.LembarPemantauan) error
	Verify(id uint, status string, namaPemeriksa string) error
	Delete(id uint) error
}

type lembarPemantauanRepository struct {
	db *gorm.DB
}

func NewLembarPemantauanRepository(db *gorm.DB) LembarPemantauanRepository {
	return &lembarPemantauanRepository{db: db}
}

func (r *lembarPemantauanRepository) Create(lembarPemantauan *models.LembarPemantauan) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(lembarPemantauan).Error; err != nil {
			// Menangkap error Unique Constraint PostgreSQL (SQLSTATE 23505)
			if strings.Contains(err.Error(), "23505") || strings.Contains(err.Error(), "duplicate key value") {
				return customerror.NewConflictError("Data pemantauan untuk periode ini sudah pernah diisi")
			}
			return customerror.NewInternalServiceError("gagal membuat lembar pemantauan")
		}
		return nil
	})
}

func (r *lembarPemantauanRepository) FindRentangUsiaByID(id uint) (*models.RentangUsia, error) {
	var ru models.RentangUsia
	if err := r.db.First(&ru, id).Error; err != nil {
		return nil, customerror.NewNotFoundError("rentang usia tidak ditemukan")
	}
	return &ru, nil
}

func (r *lembarPemantauanRepository) FindByID(id uint) (*models.LembarPemantauan, error) {
	var lembar models.LembarPemantauan
	err := r.db.
		Where("id = ? AND deleted_at IS NULL", id).
		Preload("Anak").
		Preload("Anak.Penduduk").
		Preload("Anak.Kehamilan").
		Preload("Anak.Kehamilan.Ibu").
		Preload("Anak.Kehamilan.Ibu.Kependudukan").
		Preload("RentangUsia").
		Preload("DetailGejala", func(db *gorm.DB) *gorm.DB {
			return db.Where("deleted_at IS NULL")
		}).
		Preload("DetailGejala.KategoriTandaSakit").
		First(&lembar).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, customerror.NewNotFoundError("lembar pemantauan tidak ditemukan")
		}
		return nil, customerror.NewInternalServiceError("gagal mengambil data lembar pemantauan")
	}
	return &lembar, nil
}

func (r *lembarPemantauanRepository) FindByAnakID(anakID uint) ([]models.LembarPemantauan, error) {
	var lembars []models.LembarPemantauan
	err := r.db.
		Where("anak_id = ? AND deleted_at IS NULL", anakID).
		Preload("Anak").
		Preload("Anak.Penduduk").
		Preload("Anak.Kehamilan").
		Preload("Anak.Kehamilan.Ibu").
		Preload("Anak.Kehamilan.Ibu.Kependudukan").
		Preload("RentangUsia").
		Preload("DetailGejala", func(db *gorm.DB) *gorm.DB {
			return db.Where("deleted_at IS NULL")
		}).
		Preload("DetailGejala.KategoriTandaSakit").
		Order("tanggal_periksa DESC").
		Find(&lembars).Error

	if err != nil {
		return nil, customerror.NewInternalServiceError("gagal mengambil data lembar pemantauan anak")
	}
	return lembars, nil
}

func (r *lembarPemantauanRepository) IsAnakMilikIbu(userID, anakID uint) (bool, error) {
	var count int64
	err := r.db.Table("anak a").
		Joins("JOIN kehamilan k ON k.id = a.kehamilan_id").
		Joins("JOIN ibu i ON i.id = k.ibu_id").
		Joins("JOIN penduduk ki ON ki.id = i.penduduk_id").
		Joins("JOIN pengguna p ON p.penduduk_id = ki.id").
		Where("a.id = ?", anakID).
		Where("p.id = ?", userID).
		Count(&count).Error
	if err != nil {
		return false, customerror.NewInternalServiceError("gagal memverifikasi kepemilikan data anak")
	}
	return count > 0, nil
}

func (r *lembarPemantauanRepository) FindAll() ([]models.LembarPemantauan, error) {
	var lembars []models.LembarPemantauan
	err := r.db.
		Where("deleted_at IS NULL").
		Preload("Anak").
		Preload("Anak.Penduduk").
		Preload("Anak.Kehamilan").
		Preload("Anak.Kehamilan.Ibu").
		Preload("Anak.Kehamilan.Ibu.Kependudukan").
		Preload("RentangUsia").
		Preload("DetailGejala", func(db *gorm.DB) *gorm.DB {
			return db.Where("deleted_at IS NULL")
		}).
		Preload("DetailGejala.KategoriTandaSakit").
		Order("tanggal_periksa DESC").
		Find(&lembars).Error

	if err != nil {
		return nil, customerror.NewInternalServiceError("gagal mengambil data lembar pemantauan")
	}
	return lembars, nil
}

func (r *lembarPemantauanRepository) FindRentangUsia() ([]models.RentangUsia, error) {
	var rentang []models.RentangUsia
	err := r.db.Order("id ASC").Find(&rentang).Error
	if err != nil {
		return nil, customerror.NewInternalServiceError("gagal mengambil data rentang usia")
	}
	return rentang, nil
}

func (r *lembarPemantauanRepository) FindKategoriTandaSakitByRentangUsiaID(rentangUsiaID uint) ([]models.KategoriTandaSakit, error) {
	var kategori []models.KategoriTandaSakit
	err := r.db.
		Where("rentang_usia_id = ? AND is_active = ?", rentangUsiaID, true).
		Order("id ASC").
		Find(&kategori).Error
	if err != nil {
		return nil, customerror.NewInternalServiceError("gagal mengambil data kategori tanda sakit")
	}
	return kategori, nil
}

func (r *lembarPemantauanRepository) Update(lembarPemantauan *models.LembarPemantauan) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.LembarPemantauan{}).Where("id = ?", lembarPemantauan.ID).Updates(lembarPemantauan).Error; err != nil {
			return customerror.NewInternalServiceError("gagal update lembar pemantauan")
		}

		if err := tx.Model(&models.DetailPemantauan{}).Where("lembar_pemantauan_id = ?", lembarPemantauan.ID).Update("deleted_at", time.Now()).Error; err != nil {
			return customerror.NewInternalServiceError("gagal menghapus detail gejala lama")
		}

		if len(lembarPemantauan.DetailGejala) > 0 {
			for _, detail := range lembarPemantauan.DetailGejala {
				detail.LembarPemantauanID = lembarPemantauan.ID
				detail.DeletedAt = gorm.DeletedAt{}
				if err := tx.Create(&detail).Error; err != nil {
					return customerror.NewInternalServiceError("gagal membuat detail gejala baru")
				}
			}
		}
		return nil
	})
}

func (r *lembarPemantauanRepository) Verify(id uint, status string, namaPemeriksa string) error {
	err := r.db.Model(&models.LembarPemantauan{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":         status,
			"nama_pemeriksa": namaPemeriksa,
			"updated_at":     time.Now(),
		}).Error

	if err != nil {
		return customerror.NewInternalServiceError("gagal memverifikasi lembar pemantauan")
	}
	return nil
}

func (r *lembarPemantauanRepository) Delete(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.LembarPemantauan{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error; err != nil {
			return customerror.NewInternalServiceError("gagal menghapus lembar pemantauan")
		}
		if err := tx.Model(&models.DetailPemantauan{}).Where("lembar_pemantauan_id = ?", id).Update("deleted_at", time.Now()).Error; err != nil {
			return customerror.NewInternalServiceError("gagal menghapus detail gejala")
		}
		return nil
	})
}
