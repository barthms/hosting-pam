// package repositories

// import (
// 	"errors"
// 	"monitoring-service/app/models"

// 	"gorm.io/gorm"
// )

// type RiwayatProsesMelahirkanRepository struct {
// 	db *gorm.DB
// }

// func NewRiwayatProsesMelahirkanRepository(db *gorm.DB) *RiwayatProsesMelahirkanRepository {
// 	return &RiwayatProsesMelahirkanRepository{db: db}
// }

// func (r *RiwayatProsesMelahirkanRepository) Create(rp *models.RiwayatProsesMelahirkan) error {
// 	return r.db.Create(rp).Error
// }

// func (r *RiwayatProsesMelahirkanRepository) FindByID(id int32) (*models.RiwayatProsesMelahirkan, error) {
// 	var rp models.RiwayatProsesMelahirkan
// 	err := r.db.Preload("Kehamilan.Ibu.Kependudukan").First(&rp, id).Error
// 	return &rp, err
// }

// func (r *RiwayatProsesMelahirkanRepository) FindByKehamilanID(kehamilanID int32) ([]models.RiwayatProsesMelahirkan, error) {
// 	var list []models.RiwayatProsesMelahirkan
// 	err := r.db.Where("kehamilan_id = ?", kehamilanID).Find(&list).Error
// 	return list, err
// }

// func (r *RiwayatProsesMelahirkanRepository) Update(rp *models.RiwayatProsesMelahirkan) error {
// 	return r.db.Save(rp).Error
// }

// func (r *RiwayatProsesMelahirkanRepository) Delete(id int32) error {
// 	result := r.db.Delete(&models.RiwayatProsesMelahirkan{}, id)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	if result.RowsAffected == 0 {
// 		return errors.New("data riwayat proses melahirkan tidak ditemukan")
// 	}
// 	return nil
// }


package repositories

import (
	"errors"
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type RiwayatProsesMelahirkanRepository struct {
	db *gorm.DB
}

func NewRiwayatProsesMelahirkanRepository(db *gorm.DB) *RiwayatProsesMelahirkanRepository {
	return &RiwayatProsesMelahirkanRepository{db: db}
}

func (r *RiwayatProsesMelahirkanRepository) Create(rp *models.RiwayatProsesMelahirkan) error {
	return r.db.Create(rp).Error
}

func (r *RiwayatProsesMelahirkanRepository) FindByID(id int32) (*models.RiwayatProsesMelahirkan, error) {
	var rp models.RiwayatProsesMelahirkan
	err := r.db.Preload("Kehamilan.Ibu.Kependudukan").First(&rp, id).Error
	return &rp, err
}

func (r *RiwayatProsesMelahirkanRepository) FindByKehamilanID(kehamilanID int32) ([]models.RiwayatProsesMelahirkan, error) {
	var list []models.RiwayatProsesMelahirkan
	err := r.db.Where("kehamilan_id = ?", kehamilanID).Find(&list).Error
	return list, err
}

// FindLatestByKehamilanID — ambil record riwayat proses melahirkan terbaru untuk kehamilan tertentu.
func (r *RiwayatProsesMelahirkanRepository) FindLatestByKehamilanID(kehamilanID int32) (*models.RiwayatProsesMelahirkan, error) {
	var rp models.RiwayatProsesMelahirkan
	err := r.db.
		Where("kehamilan_id = ?", kehamilanID).
		Order("created_at DESC").
		First(&rp).Error
	return &rp, err
}

// FindLatestKehamilanByUserID — ambil kehamilan PALING TERAKHIR milik user (ibu),
// tanpa filter status. Ini penting karena riwayat proses melahirkan dicatat
// setelah kehamilan selesai (status mungkin bukan 'aktif' lagi).
func (r *RiwayatProsesMelahirkanRepository) FindLatestKehamilanByUserID(userID int32) (*models.Kehamilan, error) {
	var kehamilan models.Kehamilan

	err := r.db.
		Table("kehamilan AS k").
		Select("k.*").
		Joins("JOIN ibu AS i ON i.id = k.ibu_id").
		Joins("JOIN penduduk AS p ON p.id = i.penduduk_id").
		Joins("JOIN pengguna AS u ON u.penduduk_id = p.id").
		Where("u.id = ?", userID).
		Order("k.created_at DESC").
		First(&kehamilan).Error

	return &kehamilan, err
}

func (r *RiwayatProsesMelahirkanRepository) Update(rp *models.RiwayatProsesMelahirkan) error {
	return r.db.Save(rp).Error
}

func (r *RiwayatProsesMelahirkanRepository) Delete(id int32) error {
	result := r.db.Delete(&models.RiwayatProsesMelahirkan{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("data riwayat proses melahirkan tidak ditemukan")
	}
	return nil
}