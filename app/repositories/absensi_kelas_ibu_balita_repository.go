package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type AbsensiKelasIbuBalitaRepository struct {
	db *gorm.DB
}

func NewAbsensiKelasIbuBalitaRepository(db *gorm.DB) *AbsensiKelasIbuBalitaRepository {
	return &AbsensiKelasIbuBalitaRepository{db: db}
}

// FindIbuIDByUserID mencari ibu_id berdasarkan user (pengguna) yang sedang login.
// Relasi: pengguna → penduduk → ibu
func (r *AbsensiKelasIbuBalitaRepository) FindIbuIDByUserID(userID int32) (int32, error) {
	var ibu models.Ibu

	err := r.db.
		Table("ibu AS i").
		Select("i.*").
		Joins("JOIN penduduk AS p ON p.id = i.penduduk_id").
		Joins("JOIN pengguna AS u ON u.penduduk_id = p.id").
		Where("u.id = ?", userID).
		First(&ibu).Error

	if err != nil {
		return 0, err
	}

	return ibu.IDIbu, nil
}

// FindByIbuID mengambil semua data absensi milik ibu tertentu, urut pertemuan.
func (r *AbsensiKelasIbuBalitaRepository) FindByIbuID(ibuID int32) ([]models.AbsensiKelasIbuBalita, error) {
	var list []models.AbsensiKelasIbuBalita

	err := r.db.
		Where("ibu_id = ?", ibuID).
		Order("pertemuan_ke ASC").
		Find(&list).Error

	return list, err
}

// Create menyimpan entri absensi baru.
func (r *AbsensiKelasIbuBalitaRepository) Create(data *models.AbsensiKelasIbuBalita) error {
	return r.db.Create(data).Error
}
