package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type ChecklistPemantauanIbuNifasRepository interface {
	GetByKehamilanIDAndHariNifas(kehamilanID int32, hariNifas int32) (*models.ChecklistPemantauanIbuNifas, error)
	GetFilledDaysByKehamilanID(kehamilanID int32) ([]int32, error)
	Create(data *models.ChecklistPemantauanIbuNifas) error
	Update(data *models.ChecklistPemantauanIbuNifas) error
}

type checklistPemantauanIbuNifasRepository struct {
	db *gorm.DB
}

func NewChecklistPemantauanIbuNifasRepository(db *gorm.DB) ChecklistPemantauanIbuNifasRepository {
	return &checklistPemantauanIbuNifasRepository{
		db: db,
	}
}

func (r *checklistPemantauanIbuNifasRepository) GetByKehamilanIDAndHariNifas(
	kehamilanID int32,
	hariNifas int32,
) (*models.ChecklistPemantauanIbuNifas, error) {
	var data models.ChecklistPemantauanIbuNifas

	err := r.db.
		Where("kehamilan_id = ? AND hari_nifas = ? AND deleted_at IS NULL", kehamilanID, hariNifas).
		First(&data).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *checklistPemantauanIbuNifasRepository) GetFilledDaysByKehamilanID(kehamilanID int32) ([]int32, error) {
	var days []int32

	err := r.db.
		Model(&models.ChecklistPemantauanIbuNifas{}).
		Where("kehamilan_id = ? AND deleted_at IS NULL", kehamilanID).
		Order("hari_nifas ASC").
		Pluck("hari_nifas", &days).Error

	if err != nil {
		return nil, err
	}

	return days, nil
}

func (r *checklistPemantauanIbuNifasRepository) Create(data *models.ChecklistPemantauanIbuNifas) error {
	return r.db.Create(data).Error
}

func (r *checklistPemantauanIbuNifasRepository) Update(data *models.ChecklistPemantauanIbuNifas) error {
	return r.db.Save(data).Error
}