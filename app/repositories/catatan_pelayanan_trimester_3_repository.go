package repositories

import (
	"errors"
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type CatatanPelayananTrimester3Repository struct {
	db *gorm.DB
}

func NewCatatanPelayananTrimester3Repository(db *gorm.DB) *CatatanPelayananTrimester3Repository {
	return &CatatanPelayananTrimester3Repository{db: db}
}

// Create menyimpan catatan baru dan mengembalikan data yang sudah tersimpan dengan ID dari database
func (r *CatatanPelayananTrimester3Repository) Create(c *models.CatatanPelayananTrimester3) (*models.CatatanPelayananTrimester3, error) {
	if err := r.db.Create(c).Error; err != nil {
		return nil, err
	}

	// Query kembali untuk mendapatkan data lengkap dari database dengan ID yang sudah di-generate
	var created models.CatatanPelayananTrimester3
	if err := r.db.Preload("Kehamilan.Ibu.Kependudukan").First(&created, c.IDCatatanT3).Error; err != nil {
		return nil, err
	}

	return &created, nil
}

// FindByID mengambil catatan berdasarkan ID
func (r *CatatanPelayananTrimester3Repository) FindByID(id int32) (*models.CatatanPelayananTrimester3, error) {
	var c models.CatatanPelayananTrimester3
	err := r.db.Preload("Kehamilan.Ibu.Kependudukan").First(&c, id).Error
	return &c, err
}

// FindByKehamilanID mengambil semua catatan untuk satu kehamilan, diurutkan dari terbaru
func (r *CatatanPelayananTrimester3Repository) FindByKehamilanID(kehamilanID int32) ([]models.CatatanPelayananTrimester3, error) {
	var list []models.CatatanPelayananTrimester3
	err := r.db.
		Where("kehamilan_id = ?", kehamilanID).
		Order("tanggal_periksa_stamp_paraf DESC").
		Find(&list).Error
	return list, err
}

// Update memperbarui catatan menggunakan Model().Updates() untuk update partial dan mengembalikan data terbaru
func (r *CatatanPelayananTrimester3Repository) Update(c *models.CatatanPelayananTrimester3) (*models.CatatanPelayananTrimester3, error) {
	// Gunakan .Model().Updates() untuk update partial, bukan .Save()
	// .Save() akan overwrite semua field dengan nilai zero value jika kosong
	// .Model().Updates() hanya update field yang dikirim
	if err := r.db.Model(c).Updates(c).Error; err != nil {
		return nil, err
	}

	// Query kembali untuk mendapatkan data terbaru dari database
	var updated models.CatatanPelayananTrimester3
	if err := r.db.Preload("Kehamilan.Ibu.Kependudukan").First(&updated, c.IDCatatanT3).Error; err != nil {
		return nil, err
	}

	return &updated, nil
}

// Delete menghapus catatan berdasarkan ID
func (r *CatatanPelayananTrimester3Repository) Delete(id int32) error {
	result := r.db.Delete(&models.CatatanPelayananTrimester3{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("data catatan pelayanan trimester 3 tidak ditemukan")
	}
	return nil
}
