package repositories

import (
	"errors"

	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type PelayananIbuNifasRepository struct {
	db *gorm.DB
}

func NewPelayananIbuNifasRepository(
	db *gorm.DB,
) *PelayananIbuNifasRepository {

	return &PelayananIbuNifasRepository{
		db: db,
	}
}

func (r *PelayananIbuNifasRepository) Create(
	p *models.PelayananIbuNifas,
) error {

	return r.db.Create(p).Error
}

func (r *PelayananIbuNifasRepository) FindByID(
	id int32,
) (*models.PelayananIbuNifas, error) {

	var p models.PelayananIbuNifas

	err := r.db.
		Preload("Kehamilan.Ibu.Kependudukan").
		First(&p, id).Error

	return &p, err
}

func (r *PelayananIbuNifasRepository) FindByKehamilanID(
	kehamilanID int32,
) ([]models.PelayananIbuNifas, error) {

	var list []models.PelayananIbuNifas

	err := r.db.
		Where("kehamilan_id = ?", kehamilanID).
		Order("created_at DESC").
		Find(&list).Error

	return list, err
}

func (r *PelayananIbuNifasRepository) Update(
	p *models.PelayananIbuNifas,
) error {

	return r.db.Save(p).Error
}

// ==================== MODUL IBU ====================

func (r *PelayananIbuNifasRepository) FindMineByUserID(
	userID int32,
) ([]models.PelayananIbuNifas, error) {

	var list []models.PelayananIbuNifas

	err := r.db.
		Table("pelayanan_ibu_nifas").
		Joins("JOIN kehamilan k ON k.id = pelayanan_ibu_nifas.kehamilan_id").
		Joins("JOIN ibu i ON i.id = k.ibu_id").
		Joins("JOIN penduduk pd ON pd.id = i.penduduk_id").
		Joins("JOIN pengguna u ON u.penduduk_id = pd.id").
		Where("u.id = ?", userID).
		Order("pelayanan_ibu_nifas.created_at DESC").
		Find(&list).Error

	return list, err
}

func (r *PelayananIbuNifasRepository) Delete(
	id int32,
) error {

	result := r.db.Delete(
		&models.PelayananIbuNifas{},
		id,
	)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New(
			"data pelayanan ibu nifas tidak ditemukan",
		)
	}

	return nil
}
