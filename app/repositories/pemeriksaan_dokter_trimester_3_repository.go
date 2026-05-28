// package repositories

// import (
// 	"errors"
// 	"monitoring-service/app/models"

// 	"gorm.io/gorm"
// )

// type PemeriksaanDokterTrimester3Repository struct {
// 	db *gorm.DB
// }

// func NewPemeriksaanDokterTrimester3Repository(db *gorm.DB) *PemeriksaanDokterTrimester3Repository {
// 	return &PemeriksaanDokterTrimester3Repository{db: db}
// }

// func (r *PemeriksaanDokterTrimester3Repository) Create(p *models.PemeriksaanDokterTrimester3) error {
// 	return r.db.Create(p).Error
// }

// func (r *PemeriksaanDokterTrimester3Repository) FindByID(id int32) (*models.PemeriksaanDokterTrimester3, error) {
// 	var p models.PemeriksaanDokterTrimester3
// 	err := r.db.Preload("Kehamilan.Ibu.Kependudukan").First(&p, id).Error
// 	return &p, err
// }

// func (r *PemeriksaanDokterTrimester3Repository) FindByKehamilanID(kehamilanID int32) ([]models.PemeriksaanDokterTrimester3, error) {
// 	var list []models.PemeriksaanDokterTrimester3
// 	err := r.db.Where("kehamilan_id = ?", kehamilanID).Find(&list).Error
// 	return list, err
// }

// func (r *PemeriksaanDokterTrimester3Repository) Update(p *models.PemeriksaanDokterTrimester3) error {
// 	return r.db.Save(p).Error
// }

// func (r *PemeriksaanDokterTrimester3Repository) Delete(id int32) error {
// 	result := r.db.Delete(&models.PemeriksaanDokterTrimester3{}, id)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	if result.RowsAffected == 0 {
// 		return errors.New("data pemeriksaan dokter trimester 3 tidak ditemukan")
// 	}
// 	return nil
// }

// // MODUL IBU
// func (r *PemeriksaanDokterTrimester3Repository) FindMineByUserID(userID int32) (*models.PemeriksaanDokterTrimester3, error) {
// 	var data models.PemeriksaanDokterTrimester3

// 	err := r.db.
// 		Table("pemeriksaan_dokter_trimester_3 p").
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

type PemeriksaanDokterTrimester3Repository struct {
	db *gorm.DB
}

func NewPemeriksaanDokterTrimester3Repository(db *gorm.DB) *PemeriksaanDokterTrimester3Repository {
	return &PemeriksaanDokterTrimester3Repository{db: db}
}

func (r *PemeriksaanDokterTrimester3Repository) Create(dokter *models.PemeriksaanDokterTrimester3) error {
	return r.db.Create(dokter).Error
}

func (r *PemeriksaanDokterTrimester3Repository) CreateWithLab(dokter *models.PemeriksaanDokterTrimester3, lab *models.PemeriksaanLaboratoriumJiwa) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(dokter).Error; err != nil {
			return err
		}
		if lab != nil {
			lab.KehamilanID = dokter.KehamilanID
			lab.Trimester = 3
			if err := tx.Create(lab).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *PemeriksaanDokterTrimester3Repository) UpdateWithLab(dokterID int32, dokter *models.PemeriksaanDokterTrimester3, lab *models.PemeriksaanLaboratoriumJiwa) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.PemeriksaanDokterTrimester3{}).Where("id = ?", dokterID).Updates(dokter).Error; err != nil {
			return err
		}
		if lab != nil {
			var existingLab models.PemeriksaanLaboratoriumJiwa
			err := tx.Where("kehamilan_id = ? AND trimester = ?", dokter.KehamilanID, 3).First(&existingLab).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				lab.KehamilanID = dokter.KehamilanID
				lab.Trimester = 3
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

func (r *PemeriksaanDokterTrimester3Repository) FindByID(id int32) (*models.PemeriksaanDokterTrimester3, error) {
	var p models.PemeriksaanDokterTrimester3
	err := r.db.Preload("Kehamilan.Ibu.Kependudukan").First(&p, id).Error
	return &p, err
}

func (r *PemeriksaanDokterTrimester3Repository) FindByKehamilanID(kehamilanID int32) ([]models.PemeriksaanDokterTrimester3, error) {
	var list []models.PemeriksaanDokterTrimester3
	err := r.db.Where("kehamilan_id = ?", kehamilanID).Find(&list).Error
	return list, err
}

func (r *PemeriksaanDokterTrimester3Repository) Update(dokter *models.PemeriksaanDokterTrimester3) error {
	return r.db.Save(dokter).Error
}

func (r *PemeriksaanDokterTrimester3Repository) Delete(id int32) error {
	result := r.db.Delete(&models.PemeriksaanDokterTrimester3{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("data pemeriksaan dokter trimester 3 tidak ditemukan")
	}
	return nil
}

// MODUL IBU
func (r *PemeriksaanDokterTrimester3Repository) FindMineByUserID(userID int32) (*models.PemeriksaanDokterTrimester3, error) {
	var data models.PemeriksaanDokterTrimester3

	err := r.db.
		Table("pemeriksaan_dokter_trimester_3 p").
		Joins("JOIN kehamilan k ON k.id = p.kehamilan_id").
		Joins("JOIN ibu i ON i.id = k.ibu_id").
		Joins("JOIN penduduk pd ON pd.id = i.penduduk_id").
		Joins("JOIN pengguna u ON u.penduduk_id = pd.id").
		Where("u.id = ?", userID).
		Order("p.tanggal_periksa DESC").
		First(&data).Error

	return &data, err
}

func (r *PemeriksaanDokterTrimester3Repository) FindAllMineByUserID(userID int32) ([]models.PemeriksaanDokterTrimester3, error) {
	var list []models.PemeriksaanDokterTrimester3

	err := r.db.
		Table("pemeriksaan_dokter_trimester_3 p").
		Joins("JOIN kehamilan k ON k.id = p.kehamilan_id").
		Joins("JOIN ibu i ON i.id = k.ibu_id").
		Joins("JOIN penduduk pd ON pd.id = i.penduduk_id").
		Joins("JOIN pengguna u ON u.penduduk_id = pd.id").
		Where("u.id = ?", userID).
		Order("p.tanggal_periksa DESC").
		Find(&list).Error

	return list, err
}