// package repositories

// import (
// 	"errors"
// 	"monitoring-service/app/models"

// 	"gorm.io/gorm"
// )

// type KeteranganLahirRepository struct {
// 	db *gorm.DB
// }

// func NewKeteranganLahirRepository(db *gorm.DB) *KeteranganLahirRepository {
// 	return &KeteranganLahirRepository{db: db}
// }

// func (r *KeteranganLahirRepository) Create(k *models.KeteranganLahir) error {
// 	return r.db.Create(k).Error
// }

// func (r *KeteranganLahirRepository) FindByID(id int32) (*models.KeteranganLahir, error) {
// 	var k models.KeteranganLahir
// 	err := r.db.Preload("Ibu.Kependudukan").First(&k, id).Error
// 	return &k, err
// }

// func (r *KeteranganLahirRepository) FindByIbuID(ibuID int32) ([]models.KeteranganLahir, error) {
// 	var list []models.KeteranganLahir
// 	err := r.db.Where("id_ibu_relasi = ?", ibuID).
// 		Preload("Ibu.Kependudukan").
// 		Order("tanggal_lahir DESC").
// 		Find(&list).Error
// 	return list, err
// }

// func (r *KeteranganLahirRepository) Update(k *models.KeteranganLahir) error {
// 	return r.db.Save(k).Error
// }

// func (r *KeteranganLahirRepository) Delete(id int32) error {
// 	result := r.db.Delete(&models.KeteranganLahir{}, id)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	if result.RowsAffected == 0 {
// 		return errors.New("data keterangan lahir tidak ditemukan")
// 	}
// 	return nil
// }

package repositories

import (
	"errors"
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type KeteranganLahirRepository struct {
	db *gorm.DB
}

func NewKeteranganLahirRepository(db *gorm.DB) *KeteranganLahirRepository {
	return &KeteranganLahirRepository{db: db}
}

func (r *KeteranganLahirRepository) Create(k *models.KeteranganLahir) error {
	return r.db.Create(k).Error
}

func (r *KeteranganLahirRepository) FindByID(id int32) (*models.KeteranganLahir, error) {
	var k models.KeteranganLahir
	err := r.db.Preload("Ibu.Kependudukan").First(&k, id).Error
	return &k, err
}

func (r *KeteranganLahirRepository) FindByIbuID(ibuID int32) ([]models.KeteranganLahir, error) {
	var list []models.KeteranganLahir
	err := r.db.Where("id_ibu_relasi = ?", ibuID).
		Preload("Ibu.Kependudukan").
		Order("tanggal_lahir DESC").
		Find(&list).Error
	return list, err
}

// FindMineByUserID mengambil keterangan lahir milik ibu yang sedang login.
// Pola JOIN sesuai repo lain yang sudah terbukti bekerja:
// keterangan_lahir -> ibu -> penduduk -> pengguna
func (r *KeteranganLahirRepository) FindMineByUserID(userID int32) (*models.KeteranganLahir, error) {
	var data models.KeteranganLahir
	err := r.db.
		Table("keterangan_lahir kl").
		Select("kl.*").
		Joins("JOIN ibu i ON i.id = kl.id_ibu_relasi").
		Joins("JOIN penduduk p ON p.id = i.penduduk_id").
		Joins("JOIN pengguna u ON u.penduduk_id = p.id").
		Where("u.id = ?", userID).
		Order("kl.id DESC").
		First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *KeteranganLahirRepository) Update(k *models.KeteranganLahir) error {
	return r.db.Save(k).Error
}

func (r *KeteranganLahirRepository) Delete(id int32) error {
	result := r.db.Delete(&models.KeteranganLahir{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("data keterangan lahir tidak ditemukan")
	}
	return nil
}