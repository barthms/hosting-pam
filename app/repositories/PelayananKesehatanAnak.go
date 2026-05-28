package repositories

import (
	"fmt"
	"monitoring-service/app/models"
	"time"

	"gorm.io/gorm"
)

// interface
type PelayananKesehatanAnakRepository interface {
	Create(data *models.KunjunganAnak) error
	GetByAnakID(anakID int32) ([]models.KunjunganAnak, error)
	GetByID(id int32) (*models.KunjunganAnak, error)
	GetAll()([]models.KunjunganAnak, error)
	Update(id int32, req models.UpdatePelayananKesehatanAnakRequest, tanggal time.Time, now time.Time) error
	Delete(id int32) error
}

// struct implementasi
type pelayananKesehatanAnakRepository struct {
	db *gorm.DB
}

// constructor
func NewPelayananKesehatanAnakRepository(db *gorm.DB) PelayananKesehatanAnakRepository {
	return &pelayananKesehatanAnakRepository{db: db}
}

// CREATE (parent + child, transaction)
func (r *pelayananKesehatanAnakRepository) Create(kunjungan *models.KunjunganAnak) error {
	tx := r.db.Begin()

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	// 🔥 cukup ini saja
	if err := tx.Create(kunjungan).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
	//Get All
	func (r *pelayananKesehatanAnakRepository) GetAll() ([]models.KunjunganAnak, error) {
		var result []models.KunjunganAnak

		err := r.db.
			Preload("DetailPelayanan").// relasi
			Order("tanggal DESC"). 
			Find(&result).Error

		if err != nil {
			return nil, err
		}

		return result, nil
	}

	// GET BY ANAK ID
	func (r *pelayananKesehatanAnakRepository) GetByAnakID(anakID int32) ([]models.KunjunganAnak, error) {
		var result []models.KunjunganAnak

		err := r.db.
			Preload("DetailPelayanan").
			Where("anak_id = ?", anakID).
			Find(&result).Error

		return result, err
	}

// GET BY ID
func (r *pelayananKesehatanAnakRepository) GetByID(id int32) (*models.KunjunganAnak, error) {
	var data models.KunjunganAnak

	err := r.db.
		Preload("DetailPelayanan").
		First(&data, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *pelayananKesehatanAnakRepository) Update(id int32,req models.UpdatePelayananKesehatanAnakRequest,tanggal time.Time,now time.Time,) error {

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
	if req.Lokasi != "" {
		parent["lokasi"] = req.Lokasi
	}

	if err := tx.Model(&models.KunjunganAnak{}).
		Where("id = ?", id).
		Updates(parent).Error; err != nil {
		tx.Rollback()
		return err
	}

	// AMBIL EXISTING
	var existing []models.DetailPelayanan
	if err := tx.Where("kunjungan_anak_id = ?", id).Find(&existing).Error; err != nil {
		tx.Rollback()
		return err
	}

	existingMap := make(map[int32]models.DetailPelayanan)
	existingByJenis := make(map[int32]models.DetailPelayanan)

	for _, d := range existing {
		existingMap[d.ID] = d
		existingByJenis[d.JenisPelayananID] = d
	}

	// LOOP REQUEST
	for _, d := range req.DetailPelayanan {

		// CASE 1: UPDATE by ID
		if d.ID != 0 {

			existingDetail, ok := existingMap[d.ID]
			if !ok {
				tx.Rollback()
				return fmt.Errorf("detail id %d tidak ditemukan", d.ID)
			}

			if existingDetail.KunjunganAnakID != id {
				tx.Rollback()
				return fmt.Errorf("detail id %d bukan milik kunjungan ini", d.ID)
			}

			if err := tx.Model(&models.DetailPelayanan{}).
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

		// CASE 2: UPDATE by jenis (anti duplikat)
		if existingDetail, ok := existingByJenis[d.JenisPelayananID]; ok {

			if err := tx.Model(&models.DetailPelayanan{}).
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

		// CASE 3: INSERT
		newDetail := models.DetailPelayanan{
			KunjunganAnakID:  id,
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

	return tx.Commit().Error
}
// DELETE
func (r *pelayananKesehatanAnakRepository) Delete(id int32) error {
	return r.db.Delete(&models.KunjunganAnak{}, "id = ?", id).Error
}
