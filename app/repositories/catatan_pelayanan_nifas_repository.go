package repositories

import (
	"errors"
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type CatatanPelayananNifasRepository struct {
	db *gorm.DB
}

func NewCatatanPelayananNifasRepository(db *gorm.DB) *CatatanPelayananNifasRepository {
	return &CatatanPelayananNifasRepository{db: db}
}

func (r *CatatanPelayananNifasRepository) Create(c *models.CatatanPelayananNifas) error {
	return r.db.Create(c).Error
}

func (r *CatatanPelayananNifasRepository) FindByID(id int32) (*models.CatatanPelayananNifas, error) {
	var c models.CatatanPelayananNifas
	err := r.db.Preload("Kehamilan.Ibu.Kependudukan").First(&c, id).Error
	return &c, err
}

func (r *CatatanPelayananNifasRepository) FindByKehamilanID(kehamilanID int32) ([]models.CatatanPelayananNifas, error) {
	var list []models.CatatanPelayananNifas
	err := r.db.Where("kehamilan_id = ?", kehamilanID).Order("tanggal_periksa_stamp_paraf DESC").Find(&list).Error
	return list, err
}

func (r *CatatanPelayananNifasRepository) Update(c *models.CatatanPelayananNifas) error {
	return r.db.Save(c).Error
}

func (r *CatatanPelayananNifasRepository) FindMineByUserID(
	userID int32,
) ([]models.CatatanPelayananNifas, error) {

	var list []models.CatatanPelayananNifas

	err := r.db.
		Table("catatan_pelayanan_nifas").
		Joins("JOIN kehamilan k ON k.id = catatan_pelayanan_nifas.kehamilan_id").
		Joins("JOIN ibu i ON i.id = k.ibu_id").
		Joins("JOIN penduduk pd ON pd.id = i.penduduk_id").
		Joins("JOIN pengguna u ON u.penduduk_id = pd.id").
		Where("u.id = ?", userID).
		Order("catatan_pelayanan_nifas.tanggal_periksa_stamp_paraf DESC").
		Find(&list).Error

	return list, err
}

func (r *CatatanPelayananNifasRepository) Delete(id int32) error {
	result := r.db.Delete(&models.CatatanPelayananNifas{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("data catatan pelayanan nifas tidak ditemukan")
	}
	return nil
}
