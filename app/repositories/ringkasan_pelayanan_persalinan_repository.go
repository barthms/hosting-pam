package repositories

import (
	"errors"
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type RingkasanPelayananPersalinanRepository struct {
	db *gorm.DB
}

func NewRingkasanPelayananPersalinanRepository(db *gorm.DB) *RingkasanPelayananPersalinanRepository {
	return &RingkasanPelayananPersalinanRepository{db: db}
}

func (r *RingkasanPelayananPersalinanRepository) Create(rp *models.RingkasanPelayananPersalinan) error {
	return r.db.Create(rp).Error
}

func (r *RingkasanPelayananPersalinanRepository) FindByID(id int32) (*models.RingkasanPelayananPersalinan, error) {
	var rp models.RingkasanPelayananPersalinan
	err := r.db.Preload("Kehamilan.Ibu.Kependudukan").First(&rp, id).Error
	return &rp, err
}

func (r *RingkasanPelayananPersalinanRepository) FindByKehamilanID(kehamilanID int32) ([]models.RingkasanPelayananPersalinan, error) {
	var list []models.RingkasanPelayananPersalinan
	err := r.db.Where("kehamilan_id = ?", kehamilanID).Find(&list).Error
	return list, err
}

func (r *RingkasanPelayananPersalinanRepository) Update(rp *models.RingkasanPelayananPersalinan) error {
	return r.db.Model(rp).Updates(rp).Error
}

func (r *RingkasanPelayananPersalinanRepository) GetMine(
	userID int32,
) ([]models.RingkasanPelayananPersalinan, error) {

	var list []models.RingkasanPelayananPersalinan

	err := r.db.
		Table("ringkasan_pelayanan_persalinan rp").
		Joins("JOIN kehamilan k ON k.id = rp.kehamilan_id").
		Joins("JOIN ibu i ON i.id = k.ibu_id").
		Joins("JOIN penduduk pd ON pd.id = i.penduduk_id").
		Joins("JOIN pengguna u ON u.penduduk_id = pd.id").
		Where("u.id = ?", userID).
		Find(&list).Error

	return list, err
}

func (r *RingkasanPelayananPersalinanRepository) Delete(id int32) error {
	result := r.db.Delete(&models.RingkasanPelayananPersalinan{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("data ringkasan pelayanan persalinan tidak ditemukan")
	}
	return nil
}
