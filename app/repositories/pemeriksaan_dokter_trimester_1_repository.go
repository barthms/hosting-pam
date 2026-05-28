// package repositories

// import (
// 	"errors"
// 	"monitoring-service/app/models"

// 	"gorm.io/gorm"
// )

// type PemeriksaanDokterTrimester1Repository struct {
// 	db *gorm.DB
// }

// func NewPemeriksaanDokterTrimester1Repository(db *gorm.DB) *PemeriksaanDokterTrimester1Repository {
// 	return &PemeriksaanDokterTrimester1Repository{db: db}
// }

// func (r *PemeriksaanDokterTrimester1Repository) Create(p *models.PemeriksaanDokterTrimester1) error {
// 	return r.db.Create(p).Error
// }

// func (r *PemeriksaanDokterTrimester1Repository) FindByID(id int32) (*models.PemeriksaanDokterTrimester1, error) {
// 	var p models.PemeriksaanDokterTrimester1
// 	err := r.db.Preload("Kehamilan.Ibu.Kependudukan").First(&p, id).Error
// 	return &p, err
// }

// func (r *PemeriksaanDokterTrimester1Repository) FindByKehamilanID(kehamilanID int32) ([]models.PemeriksaanDokterTrimester1, error) {
// 	var list []models.PemeriksaanDokterTrimester1
// 	err := r.db.Where("kehamilan_id = ?", kehamilanID).Find(&list).Error
// 	return list, err
// }

// func (r *PemeriksaanDokterTrimester1Repository) Update(p *models.PemeriksaanDokterTrimester1) error {
// 	return r.db.Save(p).Error
// }

// func (r *PemeriksaanDokterTrimester1Repository) Delete(id int32) error {
// 	result := r.db.Delete(&models.PemeriksaanDokterTrimester1{}, id)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	if result.RowsAffected == 0 {
// 		return errors.New("data pemeriksaan dokter trimester 1 tidak ditemukan")
// 	}
// 	return nil
// }

// // MODUL IBU
// func (r *PemeriksaanDokterTrimester1Repository) FindMineByUserID(userID int32) (*models.PemeriksaanDokterTrimester1, error) {
// 	var data models.PemeriksaanDokterTrimester1

// 	err := r.db.
// 		Table("pemeriksaan_dokter_trimester_1 p").
// 		Joins("JOIN kehamilan k ON k.id = p.kehamilan_id").
// 		Joins("JOIN ibu i ON i.id = k.ibu_id").
// 		Joins("JOIN penduduk pd ON pd.id = i.penduduk_id").
// 		Joins("JOIN pengguna u ON u.penduduk_id = pd.id").
// 		Where("u.id = ?", userID).
// 		Order("p.tanggal_periksa DESC").
// 		First(&data).Error

// 	return &data, err
// }

package repositories

import (
	"errors"
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type PemeriksaanDokterTrimester1Repository struct {
	db *gorm.DB
}

func NewPemeriksaanDokterTrimester1Repository(db *gorm.DB) *PemeriksaanDokterTrimester1Repository {
	return &PemeriksaanDokterTrimester1Repository{db: db}
}

func (r *PemeriksaanDokterTrimester1Repository) Create(dokter *models.PemeriksaanDokterTrimester1) error {
	return r.db.Create(dokter).Error
}

func (r *PemeriksaanDokterTrimester1Repository) CreateWithLab(dokter *models.PemeriksaanDokterTrimester1, lab *models.PemeriksaanLaboratoriumJiwa) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(dokter).Error; err != nil {
			return err
		}
		if lab != nil {
			lab.KehamilanID = dokter.KehamilanID
			lab.Trimester = 1
			if err := tx.Create(lab).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *PemeriksaanDokterTrimester1Repository) UpdateWithLab(dokterID int32, dokter *models.PemeriksaanDokterTrimester1, lab *models.PemeriksaanLaboratoriumJiwa) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.PemeriksaanDokterTrimester1{}).Where("id = ?", dokterID).Updates(dokter).Error; err != nil {
			return err
		}
		if lab != nil {
			var existingLab models.PemeriksaanLaboratoriumJiwa
			err := tx.Where("kehamilan_id = ? AND trimester = ?", dokter.KehamilanID, 1).First(&existingLab).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				lab.KehamilanID = dokter.KehamilanID
				lab.Trimester = 1
				if err := tx.Create(lab).Error; err != nil {
					return err
				}
			} else if err != nil {
				return err
			} else {
				if err := tx.Model(&existingLab).Updates(lab).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (r *PemeriksaanDokterTrimester1Repository) GetByID(id int32) (*models.PemeriksaanDokterTrimester1, error) {
	var data models.PemeriksaanDokterTrimester1
	err := r.db.First(&data, id).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *PemeriksaanDokterTrimester1Repository) GetByKehamilanID(kehamilanID int32) ([]models.PemeriksaanDokterTrimester1, error) {
	var list []models.PemeriksaanDokterTrimester1
	err := r.db.Where("kehamilan_id = ?", kehamilanID).Order("tanggal_periksa DESC").Find(&list).Error
	return list, err
}

func (r *PemeriksaanDokterTrimester1Repository) Update(dokter *models.PemeriksaanDokterTrimester1) error {
	return r.db.Save(dokter).Error
}

func (r *PemeriksaanDokterTrimester1Repository) Delete(id int32) error {
	result := r.db.Delete(&models.PemeriksaanDokterTrimester1{}, id)
	if result.RowsAffected == 0 {
		return errors.New("data tidak ditemukan")
	}
	return result.Error
}

// MODUL IBU - return single (existing, tetap dipakai GetMine lama)
func (r *PemeriksaanDokterTrimester1Repository) FindMineByUserID(userID int32) (*models.PemeriksaanDokterTrimester1, error) {
	var data models.PemeriksaanDokterTrimester1

	err := r.db.
		Table("pemeriksaan_dokter_trimester_1 p").
		Joins("JOIN kehamilan k ON k.id = p.kehamilan_id").
		Joins("JOIN ibu i ON i.id = k.ibu_id").
		Joins("JOIN penduduk pd ON pd.id = i.penduduk_id").
		Joins("JOIN pengguna u ON u.id = pd.id").
		Where("u.id = ?", userID).
		Order("p.tanggal_periksa DESC").
		First(&data).Error

	return &data, err
}

// MODUL IBU - return list (BARU: digunakan untuk tampilkan semua kunjungan)
func (r *PemeriksaanDokterTrimester1Repository) FindAllMineByUserID(userID int32) ([]models.PemeriksaanDokterTrimester1, error) {
	var list []models.PemeriksaanDokterTrimester1

	err := r.db.
		Table("pemeriksaan_dokter_trimester_1 p").
		Joins("JOIN kehamilan k ON k.id = p.kehamilan_id").
		Joins("JOIN ibu i ON i.id = k.ibu_id").
		Joins("JOIN penduduk pd ON pd.id = i.penduduk_id").
		Joins("JOIN pengguna u ON u.id = pd.id").
		Where("u.id = ?", userID).
		Order("p.tanggal_periksa DESC").
		Find(&list).Error

	return list, err
}

// COBA DULU
func (r *PemeriksaanDokterTrimester1Repository) FindByID(id int32) (*models.PemeriksaanDokterTrimester1, error) {
	// TODO: Implementasi DB Find By ID
	return nil, nil
}

func (r *PemeriksaanDokterTrimester1Repository) FindByKehamilanID(kehamilanID int32) ([]models.PemeriksaanDokterTrimester1, error) {
	// TODO: Implementasi DB Find By Kehamilan ID
	return nil, nil
}
