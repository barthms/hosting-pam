package repositories

import (
	"monitoring-service/app/models"
	"time"

	"gorm.io/gorm"
)

type PemantauanPertumbuhanRepository interface {
	Create(k *models.DeteksiDiniPenyimpangan) error
	GetByAnakID(anakID int32) ([]models.DeteksiDiniPenyimpangan, error)
	GetByID(id int32) (*models.DeteksiDiniPenyimpangan, error)
	GetAll() ([]models.DeteksiDiniPenyimpangan, error)
	Update(id int32, req models.UpdatePemantauanPemeriksaanRequest, now time.Time) error
	Delete(id int32) error
}
type pemantauanPertumbuhanRepository struct {
	db *gorm.DB
}

func NewPemantauanPertumbuhanRepository(db *gorm.DB) PemantauanPertumbuhanRepository {
	return &pemantauanPertumbuhanRepository{db: db}
}
func (r *pemantauanPertumbuhanRepository) withTx(fn func(tx *gorm.DB) error) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
func (r *pemantauanPertumbuhanRepository) Create(k *models.DeteksiDiniPenyimpangan) error {
	return r.withTx(func(tx *gorm.DB) error {
		return tx.Create(k).Error
	})
}
func (r *pemantauanPertumbuhanRepository) GetByAnakID(anakID int32) ([]models.DeteksiDiniPenyimpangan, error) {
	var data []models.DeteksiDiniPenyimpangan

	err := r.db.
		Where("anak_id = ?", anakID).
		Order("tanggal DESC").
		Find(&data).Error

	return data, err
}
func (r *pemantauanPertumbuhanRepository) GetByID(id int32) (*models.DeteksiDiniPenyimpangan, error) {
	var data models.DeteksiDiniPenyimpangan

	err := r.db.
		First(&data, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // lebih aman untuk service layer
		}
		return nil, err
	}

	return &data, nil
}
func (r *pemantauanPertumbuhanRepository) GetAll() ([]models.DeteksiDiniPenyimpangan, error) {
	var data []models.DeteksiDiniPenyimpangan

	err := r.db.
		Order("created_at DESC").
		Find(&data).Error

	return data, err
}
func (r *pemantauanPertumbuhanRepository) Update(id int32, req models.UpdatePemantauanPemeriksaanRequest, now time.Time) error {
	return r.withTx(func(tx *gorm.DB) error {

		updates := map[string]interface{}{
			"updated_at": now,
		}

		if req.AnakID != 0 {
			updates["anak_id"] = req.AnakID
		}
		if req.Bulanke != 0 {
			updates["bulanke"] = req.Bulanke
		}
		if req.Tanggal != "" {
			updates["tanggal"] = req.Tanggal
		}
		if req.TenagaKesehatanID != 0 {
			updates["tenaga_kesehatan_id"] = req.TenagaKesehatanID
		}

		if req.BBperU != "" {
			updates["bb_per_u"] = req.BBperU
		}

		if req.BBperTB != "" {
			updates["bb_per_tb"] = req.BBperTB
		}

		if req.TBperU != "" {
			updates["tb_per_u"] = req.TBperU
		}

		if req.LKperU != "" {
			updates["lk_per_u"] = req.LKperU
		}

		if req.LILA != "" {
			updates["lila"] = req.LILA
		}

		if req.KPSP != "" {
			updates["kpsp"] = req.KPSP
		}

		if req.TDD != "" {
			updates["tdd"] = req.TDD
		}

		if req.TDL != "" {
			updates["tdl"] = req.TDL
		}

		if req.KMPE != "" {
			updates["kmpe"] = req.KMPE
		}

		if req.MCHATRevised != "" {
			updates["mchat_revised"] = req.MCHATRevised
		}

		if req.ACTRS != "" {
			updates["actrs"] = req.ACTRS
		}

		if req.HasilPKAT != "" {
			updates["hasil_pkat"] = req.HasilPKAT
		}

		if req.Tindakan != "" {
			updates["tindakan"] = req.Tindakan
		}
		if !req.KunjunganUlang.IsZero() {
			updates["kunjungan_ulang"] = req.KunjunganUlang
		}

		return tx.Model(&models.DeteksiDiniPenyimpangan{}).
			Where("id = ?", id).
			Updates(updates).Error
	})
}
func (r *pemantauanPertumbuhanRepository) Delete(id int32) error {
	return r.withTx(func(tx *gorm.DB) error {
		return tx.Delete(&models.DeteksiDiniPenyimpangan{}, id).Error
	})
}
