package repositories

import (
	"errors"
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type CatatanPelayananTrimester1Repository struct {
	db *gorm.DB
}

func NewCatatanPelayananTrimester1Repository(db *gorm.DB) *CatatanPelayananTrimester1Repository {
	return &CatatanPelayananTrimester1Repository{db: db}
}

// Create menyimpan catatan baru dan mengembalikan data yang sudah tersimpan dengan ID dari database
func (r *CatatanPelayananTrimester1Repository) Create(c *models.CatatanPelayananTrimester1) (*models.CatatanPelayananTrimester1, error) {
	if err := r.db.Create(c).Error; err != nil {
		return nil, err
	}

	// Query kembali untuk mendapatkan data lengkap dari database dengan ID yang sudah di-generate
	var created models.CatatanPelayananTrimester1
	if err := r.db.Preload("Kehamilan.Ibu.Kependudukan").First(&created, c.IDCatatan).Error; err != nil {
		return nil, err
	}

	return &created, nil
}

// FindByID mengambil catatan berdasarkan ID
func (r *CatatanPelayananTrimester1Repository) FindByID(id int32) (*models.CatatanPelayananTrimester1, error) {
	var c models.CatatanPelayananTrimester1
	err := r.db.Preload("Kehamilan.Ibu.Kependudukan").First(&c, id).Error
	return &c, err
}

// FindByKehamilanID mengambil semua catatan untuk satu kehamilan, diurutkan dari terbaru
func (r *CatatanPelayananTrimester1Repository) FindByKehamilanID(kehamilanID int32) ([]models.CatatanPelayananTrimester1, error) {
	var list []models.CatatanPelayananTrimester1
	err := r.db.
		Where("kehamilan_id = ?", kehamilanID).
		Order("tanggal_periksa_stamp_paraf DESC").
		Find(&list).Error
	return list, err
}

// Update memperbarui catatan menggunakan Model().Updates() untuk update partial dan mengembalikan data terbaru
func (r *CatatanPelayananTrimester1Repository) Update(c *models.CatatanPelayananTrimester1) (*models.CatatanPelayananTrimester1, error) {
	// Gunakan .Model().Updates() untuk update partial, bukan .Save()
	// .Save() akan overwrite semua field dengan nilai zero value jika kosong
	// .Model().Updates() hanya update field yang dikirim
	if err := r.db.Model(c).Updates(c).Error; err != nil {
		return nil, err
	}

	// Query kembali untuk mendapatkan data terbaru dari database
	var updated models.CatatanPelayananTrimester1
	if err := r.db.Preload("Kehamilan.Ibu.Kependudukan").First(&updated, c.IDCatatan).Error; err != nil {
		return nil, err
	}

	return &updated, nil
}

// Delete menghapus catatan berdasarkan ID
func (r *CatatanPelayananTrimester1Repository) Delete(id int32) error {
	result := r.db.Delete(&models.CatatanPelayananTrimester1{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("data catatan pelayanan trimester 1 tidak ditemukan")
	}
	return nil
}
