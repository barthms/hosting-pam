package repositories

import (
	"errors"
	"monitoring-service/app/models"
	"time"

	"gorm.io/gorm"
)

type KehamilanRepository struct {
	db *gorm.DB
}

func NewKehamilanRepository(db *gorm.DB) *KehamilanRepository {
	return &KehamilanRepository{db: db}
}

func (r *KehamilanRepository) Create(kehamilan *models.Kehamilan) error {
	return r.db.Create(kehamilan).Error
}

func (r *KehamilanRepository) FindByID(id int32) (*models.Kehamilan, error) {
	var kehamilan models.Kehamilan
	err := r.db.Preload("Ibu").Preload("Ibu.Kependudukan").Preload("Anak").First(&kehamilan, id).Error
	if err != nil {
		return nil, err
	}
	return &kehamilan, nil
}

func (r *KehamilanRepository) FindByIbuID(ibuID int32) ([]models.Kehamilan, error) {
	var list []models.Kehamilan
	err := r.db.Where("ibu_id = ?", ibuID).Order("created_at DESC").Find(&list).Error
	return list, err
}

func (r *KehamilanRepository) GetAll() ([]models.Kehamilan, error) {
	var list []models.Kehamilan
	err := r.db.Preload("Ibu.Kependudukan").Preload("Anak").Find(&list).Error
	return list, err
}

func (r *KehamilanRepository) Update(kehamilan *models.Kehamilan) error {
	// Gunakan Save untuk update semua field (termasuk yang kosong)
	return r.db.Save(kehamilan).Error
}

func (r *KehamilanRepository) Delete(id int32) error {
	result := r.db.Delete(&models.Kehamilan{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("data kehamilan tidak ditemukan")
	}
	return nil
}
func (r *KehamilanRepository) UpdateStatusKehamilan(id int32, status string) error {
	return r.db.Model(&models.Kehamilan{}).
		Where("id = ?", id).
		Update("status_kehamilan", status).Error
}

type KehamilanDashboardDTO struct {
	IDIbu           int32     `json:"id_ibu"`
	NamaLengkap     string    `json:"nama_lengkap"`
	Dusun           string    `json:"dusun"`
	Desa            string    `json:"desa"`
	KehamilanID     int32     `json:"kehamilan_id"`
	HPHT            time.Time `json:"hpht"`
	StatusKehamilan string    `json:"status_kehamilan"`

	SkorRisiko   *int32  `json:"skor_risiko"`
	StatusRisiko *string `json:"status_risiko"`
}

func (r *KehamilanRepository) GetDashboardIbuHamil() ([]KehamilanDashboardDTO, error) {
	var result []KehamilanDashboardDTO

	err := r.db.
		Table("kehamilan k").
		Select(`
			k.id as kehamilan_id,
		 	i.id as id_ibu,
			p.nama_lengkap,
			p.dusun,
			p.desa,
			k.hpht,
			k.status_kehamilan,
			anc.skor_risiko,
			anc.status_risiko
		`).
		Joins("JOIN ibu i ON k.ibu_id = i.id").
		Joins("JOIN penduduk p ON i.penduduk_id = p.id").
		Joins(`
			LEFT JOIN pemeriksaan_kehamilan anc 
			ON anc.id_periksa = (
				SELECT id_periksa FROM pemeriksaan_kehamilan
				WHERE kehamilan_id = k.id
				ORDER BY created_at DESC
				LIMIT 1
			)
		`).
		Where("k.status_kehamilan IN ?", []string{
			"TRIMESTER 1",
			"TRIMESTER 2",
			"TRIMESTER 3",
		}).
		Scan(&result).Error

	return result, err
}
func (r *KehamilanRepository) ExistsActiveByIbuID(ibuID int32) (bool, error) {
	var count int64

	err := r.db.Model(&models.Kehamilan{}).
		Where("ibu_id = ? AND status_kehamilan != ?", ibuID, "NON-AKTIF").
		Count(&count).Error

	return count > 0, err
}

// GetActiveKehamilanList mengambil semua kehamilan dengan status TRIMESTER (1,2,3) dan HPHT tidak kosong
func (r *KehamilanRepository) GetActiveKehamilanList() ([]models.Kehamilan, error) {
	var list []models.Kehamilan
	err := r.db.
		Where("status_kehamilan IN ?", []string{"TRIMESTER 1", "TRIMESTER 2", "TRIMESTER 3"}).
		Where("hpht IS NOT NULL").
		Find(&list).Error
	return list, err
}

// UpdateUsiaDanStatusKehamilan mengupdate uk_kehamilan_saat_ini dan status_kehamilan berdasarkan ID
func (r *KehamilanRepository) UpdateUsiaDanStatusKehamilan(id int32, usia int32, status string) error {
	return r.db.Model(&models.Kehamilan{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"uk_kehamilan_saat_ini": usia,
			"status_kehamilan":      status,
		}).Error
}

func (r *KehamilanRepository) FindAktifByUserID(userID int32) (*models.Kehamilan, error) {
    var kehamilan models.Kehamilan

    err := r.db.
        Table("kehamilan AS k").
        Select("k.*").
        Joins("JOIN ibu AS i ON i.id = k.ibu_id").
        Joins("JOIN penduduk AS p ON p.id = i.penduduk_id").
        Joins("JOIN pengguna AS u ON u.penduduk_id = p.id").
        Where("u.id = ?", userID).
        Where("k.status_kehamilan IN ?", []string{"aktif", "TRIMESTER 1", "TRIMESTER 2", "TRIMESTER 3"}).
        Order("k.created_at DESC").
        First(&kehamilan).Error

    if err != nil {
        return nil, err
    }
    return &kehamilan, nil
}

// func (r *KehamilanRepository) ExistsActiveByIbuID(ibuID int32) (bool, error) {
//     return false, nil
// }

func (r *KehamilanRepository) UpdateAllActiveGestationalAge() error {
    return nil
}