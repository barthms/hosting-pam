// app/repositories/laporan_ibu_repository.go
package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type LaporanIbuRepository interface {
	GetLaporanIbu() ([]models.LaporanIbu, error)
}

type laporanIbuRepository struct {
	db *gorm.DB
}

func NewLaporanIbuRepository(
	db *gorm.DB,
) LaporanIbuRepository {
	return &laporanIbuRepository{db}
}

func (r *laporanIbuRepository) GetLaporanIbu() ([]models.LaporanIbu, error) {

	var result []models.LaporanIbu

	err := r.db.Table("pemeriksaan_kehamilan pk").
		Select(`
			p.nik,
			p.nama_lengkap as nama_ibu,
			COALESCE(s.nama_lengkap, '') as nama_suami,
			p.tanggal_lahir,

			k.hpht,
			k.taksiran_persalinan as hpl,

			k.uk_kehamilan_saat_ini as usia_kehamilan,

			pk.trimester,

			i.gravida,
			i.paritas,
			i.abortus,

			k.bb_awal as bb_awal,
			k.tb as tinggi_badan,
			k.imt_awal as imt,

			COALESCE(pk.lingkar_lengan_atas, 0) as lila,

			CONCAT(pk.sistole, '/', pk.diastole) as tekanan_darah,
			pk.sistole,
			pk.diastole,

			COALESCE(pk.tinggi_rahim, 0) as tinggi_fundus,

			COALESCE(pk.tes_lab_hb, 0) as hb,
			pk.tes_golongan_darah as golongan_darah,

			pk.status_imunisasi_tetanus as status_imunisasi,

			pk.tripel_eliminasi,

			pk.kunjungan_ke as kunjungan_anc,

			pk.tata_laksana_kasus as tindakan,

			p.kecamatan,
			p.desa
		`).
		Joins("JOIN kehamilan k ON pk.kehamilan_id = k.id").
		Joins("JOIN ibu i ON k.ibu_id = i.id").
		Joins("JOIN penduduk p ON i.penduduk_id = p.id").
		Joins("LEFT JOIN penduduk s ON i.suami_id = s.id").
		Scan(&result).Error

	return result, err
}