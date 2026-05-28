package repositories

import (
	"errors"
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type RencanaPersalinanRepository struct {
	db *gorm.DB
}

func NewRencanaPersalinanRepository(db *gorm.DB) *RencanaPersalinanRepository {
	return &RencanaPersalinanRepository{db: db}
}

func (r *RencanaPersalinanRepository) Create(rp *models.RencanaPersalinan) error {
	return r.db.Create(rp).Error
}

func (r *RencanaPersalinanRepository) FindByID(id int32) (*models.RencanaPersalinan, error) {
	var rp models.RencanaPersalinan
	err := r.db.Preload("Kehamilan.Ibu.Kependudukan").First(&rp, "id_rencana_persalinan = ?", id).Error
	return &rp, err
}

// FindByIDWithoutPreload untuk update (tanpa preload relasi)
func (r *RencanaPersalinanRepository) FindByIDWithoutPreload(id int32) (*models.RencanaPersalinan, error) {
	var rp models.RencanaPersalinan
	err := r.db.First(&rp, "id_rencana_persalinan = ?", id).Error
	return &rp, err
}

func (r *RencanaPersalinanRepository) FindByKehamilanID(kehamilanID int32) ([]models.RencanaPersalinan, error) {
	var list []models.RencanaPersalinan
	err := r.db.Where("kehamilan_id = ?", kehamilanID).Find(&list).Error
	return list, err
}

func (r *RencanaPersalinanRepository) Update(rp *models.RencanaPersalinan) error {
	// Hanya update field yang ada di tabel rencana_persalinan
	return r.db.Model(&models.RencanaPersalinan{}).
		Where("id_rencana_persalinan = ?", rp.IDRencanaPersalinan).
		Updates(map[string]interface{}{
			"nama_ibu_pernyataan":        rp.NamaIbuPernyataan,
			"alamat_ibu_pernyataan":      rp.AlamatIbuPernyataan,
			"perkiraan_bulan_persalinan": rp.PerkiraanBulanPersalinan,
			"perkiraan_tahun_persalinan": rp.PerkiraanTahunPersalinan,
			"fasyankes1_nama_tenaga":     rp.Fasyankes1NamaTenaga,
			"fasyankes1_nama_fasilitas":  rp.Fasyankes1NamaFasilitas,
			"fasyankes2_nama_tenaga":     rp.Fasyankes2NamaTenaga,
			"fasyankes2_nama_fasilitas":  rp.Fasyankes2NamaFasilitas,
			"sumber_dana_persalinan":     rp.SumberDanaPersalinan,
			"kendaraan1_nama":            rp.Kendaraan1Nama,
			"kendaraan1_hp":              rp.Kendaraan1HP,
			"metode_kontrasepsi_pilihan": rp.MetodeKontrasepsiPilihan,
			"donor_golongan_darah":       rp.DonorGolonganDarah,
			"donor_rhesus":               rp.DonorRhesus,
		}).Error
}

func (r *RencanaPersalinanRepository) Delete(id int32) error {
	result := r.db.Delete(&models.RencanaPersalinan{}, "id_rencana_persalinan = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("data rencana persalinan tidak ditemukan")
	}
	return nil
}
