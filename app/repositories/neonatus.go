package repositories

import (
	"fmt"
	"monitoring-service/app/models"
	"time"

	"gorm.io/gorm"
)

type PelayananNeonatusRepository interface {
	Create(data *models.Neonatus) error
	GetByAnakID(anakID int32) ([]models.Neonatus, error)
	GetByID(id int32) (*models.Neonatus, error)
	GetAll() ([]models.Neonatus, error)
	Update(id int32, req models.UpdatePelayananNeonatusRequest, tanggal time.Time, now time.Time) error
	Delete(id int32) error
}

type pelayananNeonatusRepository struct {
	db *gorm.DB
}

func NewPelayananNeonatusRepository(db *gorm.DB) PelayananNeonatusRepository {
	return &pelayananNeonatusRepository{db: db}
}

func (r *pelayananNeonatusRepository) Create(neonatus *models.Neonatus) error {
	tx := r.db.Begin()

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	// 🔥 cukup ini saja
	if err := tx.Create(neonatus).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Get All
func (r *pelayananNeonatusRepository) GetAll() ([]models.Neonatus, error) {
	var result []models.Neonatus

	err := r.db.
		Preload("DetailPelayanan"). // relasi
		Order("tanggal DESC").
		Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

// GET BY ANAK ID
func (r *pelayananNeonatusRepository) GetByAnakID(anakID int32) ([]models.Neonatus, error) {
	var result []models.Neonatus

	err := r.db.
		Preload("DetailPelayanan").
		Where("anak_id = ?", anakID).
		Find(&result).Error

	return result, err
}

// GET BY ID
func (r *pelayananNeonatusRepository) GetByID(id int32) (*models.Neonatus, error) {
	var data models.Neonatus

	err := r.db.
		Preload("DetailPelayanan").
		First(&data, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

// UPDATE (parent + detail)
func (r *pelayananNeonatusRepository) Update(id int32,req models.UpdatePelayananNeonatusRequest,tanggal time.Time,now time.Time,) error {

	tx := r.db.Begin()

	// UPDATE PARENT
	
	parent := map[string]interface{}{
		"updated_at": now,
	}

	if req.AnakID != 0 {
		parent["anak_id"] = req.AnakID
	}
	if req.Tanggal != "" {
		parent["tanggal"] = tanggal
	}
	if req.PeriodeID != 0 {
		parent["periode_id"] = req.PeriodeID
	}
	if req.KategoriUmurID != 0 {
		parent["kategori_umur_id"] = req.KategoriUmurID
	}

	if req.TenagaKesehatanID != 0 {
		parent["tenaga_kesehatan_id"] = req.TenagaKesehatanID
	}

	if err := tx.Model(&models.Neonatus{}).
		Where("id = ?", id).
		Updates(parent).Error; err != nil {
		tx.Rollback()
		return err
	}

	
	// AMBIL DATA EXISTING
	
	var existing []models.DetailPelayananNeonatus
	if err := tx.Where("neonatus_id = ?", id).Find(&existing).Error; err != nil {
		tx.Rollback()
		return err
	}

	// map by ID
	existingMap := make(map[int32]models.DetailPelayananNeonatus)
	// map by jenis_pelayanan
	existingByJenis := make(map[int32]models.DetailPelayananNeonatus)

	for _, d := range existing {
		existingMap[d.ID] = d
		existingByJenis[d.JenisPelayananID] = d
	}

	
	// LOOP REQUEST
	
	for _, d := range req.DetailPelayanan {

		
		// CASE 1: ID ADA → UPDATE BY ID
		
		if d.ID != 0 {

			existingDetail, ok := existingMap[d.ID]
			if !ok {
				tx.Rollback()
				return fmt.Errorf("detail id %d tidak ditemukan", d.ID)
			}

			if existingDetail.NeonatusID != id {
				tx.Rollback()
				return fmt.Errorf("detail id %d bukan milik neonatus ini", d.ID)
			}

			if err := tx.Model(&models.DetailPelayananNeonatus{}).
				Where("id = ?", d.ID).
				Updates(map[string]interface{}{
					"jenis_pelayanan_id": d.JenisPelayananID,
					"nilai":              d.Nilai,
					"keterangan":         d.Keterangan,
					"updated_at":         now,
				}).Error; err != nil {
				tx.Rollback()
				return err
			}

			continue
		}

		// CASE 2: ID TIDAK ADA → CEK JENIS
		if existingDetail, ok := existingByJenis[d.JenisPelayananID]; ok {

			// UPDATE existing (anti duplikat)
			if err := tx.Model(&models.DetailPelayananNeonatus{}).
				Where("id = ?", existingDetail.ID).
				Updates(map[string]interface{}{
					"nilai":      d.Nilai,
					"keterangan": d.Keterangan,
					"updated_at": now,
				}).Error; err != nil {
				tx.Rollback()
				return err
			}

			continue
		}

		// CASE 3: BENAR-BENAR BARU → INSERT

		newDetail := models.DetailPelayananNeonatus{
			NeonatusID:       id,
			JenisPelayananID: d.JenisPelayananID,
			Nilai:            d.Nilai,
			Keterangan:       d.Keterangan,
			CreatedAt:        now,
			UpdatedAt:        now,
		}

		if err := tx.Create(&newDetail).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// COMMIT

	return tx.Commit().Error
}

// DELETE
func (r *pelayananNeonatusRepository) Delete(id int32) error {
	return r.db.Delete(&models.Neonatus{}, "id = ?", id).Error
}
