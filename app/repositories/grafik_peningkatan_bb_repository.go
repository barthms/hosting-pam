// package repositories

// import (
// 	"errors"
// 	"monitoring-service/app/models"

// 	"gorm.io/gorm"
// )

// type GrafikPeningkatanBBRepository struct {
// 	db *gorm.DB
// }

// func NewGrafikPeningkatanBBRepository(db *gorm.DB) *GrafikPeningkatanBBRepository {
// 	return &GrafikPeningkatanBBRepository{db: db}
// }

// //
// // CREATE
// //
// func (r *GrafikPeningkatanBBRepository) Create(g *models.GrafikPeningkatanBB) error {
// 	return r.db.Create(g).Error
// }

// //
// // FIND BY ID
// //
// func (r *GrafikPeningkatanBBRepository) FindByID(id int32) (*models.GrafikPeningkatanBB, error) {
// 	var g models.GrafikPeningkatanBB

// 	err := r.db.
// 		Preload("Kehamilan").
// 		First(&g, id).Error

// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, errors.New("data grafik tidak ditemukan")
// 		}
// 		return nil, err
// 	}

// 	return &g, nil
// }

// //
// // FIND BY KEHAMILAN ID (UNTUK GRAFIK)
// //
// func (r *GrafikPeningkatanBBRepository) FindByKehamilanID(kehamilanID int32) ([]models.GrafikPeningkatanBB, error) {
// 	var list []models.GrafikPeningkatanBB

// 	err := r.db.
// 		Where("kehamilan_id = ?", kehamilanID).
// 		Order("minggu_kehamilan ASC").
// 		Find(&list).Error

// 	return list, err
// }

// //
// // 🔥 CEK DUPLIKAT MINGGU (PENTING)
// //
// func (r *GrafikPeningkatanBBRepository) FindByKehamilanIDAndMinggu(kehamilanID int32, minggu int) (*models.GrafikPeningkatanBB, error) {
// 	var g models.GrafikPeningkatanBB

// 	err := r.db.
// 		Where("kehamilan_id = ? AND minggu_kehamilan = ?", kehamilanID, minggu).
// 		First(&g).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &g, nil
// }

// //
// // UPDATE
// //
// func (r *GrafikPeningkatanBBRepository) Update(g *models.GrafikPeningkatanBB) error {
// 	return r.db.Save(g).Error
// }

// //
// // DELETE
// //
// func (r *GrafikPeningkatanBBRepository) Delete(id int32) error {
// 	result := r.db.Delete(&models.GrafikPeningkatanBB{}, id)

// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	if result.RowsAffected == 0 {
// 		return errors.New("data grafik peningkatan berat badan tidak ditemukan")
// 	}

// 	return nil
// }

package repositories

import (
	"errors"
	"monitoring-service/app/models"
	"strings"
	"time"

	"gorm.io/gorm"
)

type GrafikPeningkatanBBRepository struct {
	db *gorm.DB
}

func NewGrafikPeningkatanBBRepository(db *gorm.DB) *GrafikPeningkatanBBRepository {
	return &GrafikPeningkatanBBRepository{db: db}
}

//
// ====================== TABEL LAMA (TENAGA KESEHATAN) ======================
// Tetap dipertahankan agar endpoint /tenaga tidak rusak
//

func (r *GrafikPeningkatanBBRepository) Create(g *models.GrafikPeningkatanBB) error {
	err := r.db.Create(g).Error
	if err != nil {
		// 🔥 handle duplicate dari DB (unique index)
		if strings.Contains(err.Error(), "duplicate") {
			return errors.New("data minggu ini sudah ada")
		}
		return err
	}
	return nil
}

func (r *GrafikPeningkatanBBRepository) FindByID(id int32) (*models.GrafikPeningkatanBB, error) {
	var g models.GrafikPeningkatanBB

	err := r.db.
		Preload("Kehamilan").
		First(&g, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 🔥 ubah jadi nil
		}
		return nil, err
	}

	return &g, nil
}

func (r *GrafikPeningkatanBBRepository) FindByKehamilanID(kehamilanID int32) ([]models.GrafikPeningkatanBB, error) {
	var list []models.GrafikPeningkatanBB

	err := r.db.
		Where("kehamilan_id = ?", kehamilanID).
		Order("minggu_kehamilan ASC").
		Find(&list).Error

	if err != nil {
		return nil, err
	}

	return list, nil
}

// CEK DUPLIKAT MINGGU (PENTING)
func (r *GrafikPeningkatanBBRepository) FindByKehamilanIDAndMinggu(kehamilanID int32, minggu int) (*models.GrafikPeningkatanBB, error) {
	var g models.GrafikPeningkatanBB

	err := r.db.
		Where("kehamilan_id = ? AND minggu_kehamilan = ?", kehamilanID, minggu).
		First(&g).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &g, nil
}

func (r *GrafikPeningkatanBBRepository) Update(g *models.GrafikPeningkatanBB) error {
	err := r.db.
		Model(&models.GrafikPeningkatanBB{}).
		Where("id = ?", g.ID).
		Updates(map[string]interface{}{
			"berat_badan":             g.BeratBadan,
			"minggu_kehamilan":        g.MingguKehamilan,
			"penjelasan_hasil_grafik": g.PenjelasanHasilGrafik,
		}).Error

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return errors.New("data minggu ini sudah ada")
		}
		return err
	}

	return nil
}

// DELETE
func (r *GrafikPeningkatanBBRepository) Delete(id int32) error {
	result := r.db.Delete(&models.GrafikPeningkatanBB{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("data tidak ditemukan")
	}

	return nil
}

//
// ====================== MODUL IBU — ON-THE-FLY dari sumber langsung ======================
//

// BBPemeriksaanRaw - raw data BB dari tabel pemeriksaan_kehamilan
type BBPemeriksaanRaw struct {
	KehamilanID    int32      `gorm:"column:kehamilan_id"`
	TanggalPeriksa *time.Time `gorm:"column:tanggal_periksa"`
	BeratBadan     *float64   `gorm:"column:berat_badan"`
}

// FindKehamilanAktifForBB - ambil kehamilan aktif beserta BB_Awal, IMT_Awal, dan HPHT
func (r *GrafikPeningkatanBBRepository) FindKehamilanAktifForBB(userID int32) (*models.Kehamilan, error) {
	var kehamilan models.Kehamilan

	err := r.db.
		Table("kehamilan k").
		Select("k.*").
		Joins("JOIN ibu i ON i.id = k.ibu_id").
		Joins("JOIN penduduk p ON p.id = i.penduduk_id").
		Joins("JOIN pengguna u ON u.penduduk_id = p.id").
		Where("u.id = ?", userID).
		Where("k.status_kehamilan IN ?", []string{"aktif", "TRIMESTER 1", "TRIMESTER 2", "TRIMESTER 3"}).
		Where("k.deleted_at IS NULL").
		Order("k.created_at DESC").
		First(&kehamilan).Error

	if err != nil {
		return nil, err
	}

	return &kehamilan, nil
}

// FindBBPemeriksaanForGrafik - ambil data BB dari pemeriksaan_kehamilan
// Hanya baris yang memiliki tanggal_periksa dan berat_badan terisi
func (r *GrafikPeningkatanBBRepository) FindBBPemeriksaanForGrafik(kehamilanID int32) ([]BBPemeriksaanRaw, error) {
	var result []BBPemeriksaanRaw

	err := r.db.
		Table("pemeriksaan_kehamilan").
		Select("kehamilan_id, tanggal_periksa, berat_badan").
		Where("kehamilan_id = ?", kehamilanID).
		Where("tanggal_periksa IS NOT NULL").
		Where("berat_badan IS NOT NULL").
		Order("tanggal_periksa ASC").
		Find(&result).Error

	return result, err
}
