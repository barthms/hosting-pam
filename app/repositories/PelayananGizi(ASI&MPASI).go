package repositories

import (
	"encoding/json"
	"errors"
	"monitoring-service/app/models"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type KunjunganGiziRepository interface {
	Create(data *models.KunjunganGizi) error
	GetByAnakID(anakID int32) ([]models.KunjunganGizi, error)
	GetByID(id int32) (*models.KunjunganGizi, error)
	GetAll() ([]models.KunjunganGizi, error)
	Update(id int32, req models.UpdatePelayananGiziRequest, tanggal time.Time, now time.Time) error
	Delete(id int32) error
}

type kunjunganGiziRepository struct {
	db *gorm.DB
}

func NewKunjunganGiziRepository(db *gorm.DB) KunjunganGiziRepository {
	return &kunjunganGiziRepository{db: db}
}

func (r *kunjunganGiziRepository) Create(kunjungan *models.KunjunganGizi) error {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(kunjungan).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// ================= GET =================

func (r *kunjunganGiziRepository) GetAll() ([]models.KunjunganGizi, error) {
	var result []models.KunjunganGizi

	err := r.db.
		Preload("ASI").
		Preload("MPASI").
		Order("tanggal DESC").
		Find(&result).Error

	return result, err
}

func (r *kunjunganGiziRepository) GetByAnakID(anakID int32) ([]models.KunjunganGizi, error) {
	var result []models.KunjunganGizi

	err := r.db.
		Preload("ASI").
		Preload("MPASI").
		Where("anak_id = ?", anakID).
		Find(&result).Error

	return result, err
}

func (r *kunjunganGiziRepository) GetByID(id int32) (*models.KunjunganGizi, error) {
	var data models.KunjunganGizi

	err := r.db.
		Preload("ASI").
		Preload("MPASI").
		First(&data, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

// ================= UPDATE =================

func (r *kunjunganGiziRepository) Update(id int32, req models.UpdatePelayananGiziRequest, tanggal time.Time, now time.Time) error {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// ===== UPDATE PARENT =====
	parent := map[string]interface{}{
		"updated_at": now,
	}

	if req.AnakID != 0 {
		parent["anak_id"] = req.AnakID
	}
	if req.Tanggal != "" {
		parent["tanggal"] = tanggal
	}
	if req.TenagaKesehatanID != 0 {
		parent["tenaga_kesehatan_id"] = req.TenagaKesehatanID
	}
	if req.Bulanke != 0 {
		parent["bulanke"] = req.Bulanke 
	}
	if req.Lokasi != "" {
		parent["lokasi"] = req.Lokasi
	}

	if err := tx.Model(&models.KunjunganGizi{}).
		Where("id = ?", id).
		Updates(parent).Error; err != nil {
		tx.Rollback()
		return err
	}

	// ===== ASI (ONE TO ONE) =====
	if req.ASI != nil {
		var asi models.ASI
		err := tx.Where("kunjungan_gizi_id = ?", id).First(&asi).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			// CREATE
			newASI := models.ASI{
				KunjunganGiziID:   id,
				FrekuensiMenyusui: req.ASI.FrekuensiMenyusui,
				PosisiMenyusui:    req.ASI.PosisiMenyusui,
				ASIPerah:          req.ASI.ASIPerah,
				CreatedAt:         now,
				UpdatedAt:         now,
			}

			if err := tx.Create(&newASI).Error; err != nil {
				tx.Rollback()
				return err
			}

		} else if err != nil {
			tx.Rollback()
			return err

		} else {
			// UPDATE
			if err := tx.Model(&asi).Updates(map[string]interface{}{
				"frekuensi_menyusui": req.ASI.FrekuensiMenyusui,
				"posisi_menyusui":    req.ASI.PosisiMenyusui,
				"asi_perah":           req.ASI.ASIPerah,
				"updated_at":         now,
			}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// ===== MPASI (ONE TO ONE) =====
	if req.MPASI != nil {
		var mpasi models.MPASI
		err := tx.Where("kunjungan_gizi_id = ?", id).First(&mpasi).Error

		variasiJSON, errMarshal := json.Marshal(req.MPASI.VariasiMPASI)
		if errMarshal != nil {
			tx.Rollback()
			return errors.New("gagal encode variasi MPASI")
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			// CREATE
			newMPASI := models.MPASI{
				KunjunganGiziID:     id,
				DiberikanMPASI:      req.MPASI.DiberikanMPASI,
				VariasiMPASI:        datatypes.JSON(variasiJSON),
				JumlahmakanPerporsi: req.MPASI.JumlahMakanPerporsi,
				FrekuensiMakan:      req.MPASI.FrekuensiMakan,
				CreatedAt:           now,
				UpdatedAt:           now,
			}

			if err := tx.Create(&newMPASI).Error; err != nil {
				tx.Rollback()
				return err
			}

		} else if err != nil {
			tx.Rollback()
			return err

		} else {
			// UPDATE
			if err := tx.Model(&mpasi).Updates(map[string]interface{}{
				"diberikan_mpasi":      req.MPASI.DiberikanMPASI,
				"variasi_mpasi":        datatypes.JSON(variasiJSON),
				"jumlahmakan_perporsi": req.MPASI.JumlahMakanPerporsi,
				"frekuensi_makan":      req.MPASI.FrekuensiMakan,
				"updated_at":           now,
			}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

// ================= DELETE =================

func (r *kunjunganGiziRepository) Delete(id int32) error {
	return r.db.Delete(&models.KunjunganGizi{}, "id = ?", id).Error
}