package models

import "time"

type PelayananIbuNifas struct {
	ID          int32      `gorm:"primaryKey" json:"id"`
	KehamilanID int32      `gorm:"not null;index" json:"kehamilan_id"`
	Kehamilan   *Kehamilan `gorm:"foreignKey:KehamilanID;references:ID" json:"kehamilan,omitempty"`

	KunjunganKe            string     `gorm:"type:varchar(10)" json:"kunjungan_ke"`
	TanggalPeriksa         *time.Time `gorm:"type:date" json:"tanggal_periksa"`
	TandaVitalTekananDarah string     `gorm:"type:varchar(20)" json:"tanda_vital_tekanan_darah"`
	TandaVitalSuhuTubuh    *float64   `gorm:"type:decimal(4,1)" json:"tanda_vital_suhu_tubuh"`

	PelayananInvolusiUteri     string `json:"pelayanan_involusi_uteri"`
	PelayananCairanPervaginam  string `json:"pelayanan_cairan_pervaginam"`
	PelayananPeriksaJalanLahir string `json:"pelayanan_periksa_jalan_lahir"`
	PelayananPeriksaPayudara   string `json:"pelayanan_periksa_payudara"`
	PelayananASIEkslusif       string `gorm:"type:varchar(50)" json:"pelayanan_asi_eksklusif"`

	PemberianKapsulVitaminA          bool `json:"pemberian_kapsul_vitamin_a"`
	PemberianTabletTambahDarahJumlah *int `json:"pemberian_tablet_tambah_darah_jumlah"`

	PelayananSkriningDepresiNifas       string `json:"pelayanan_skrining_depresi_nifas"`
	PelayananKontrasepsiPascaPersalinan string `json:"pelayanan_kontrasepsi_pasca_persalinan"`
	PelayananPenangananRisikoMalaria    string `json:"pelayanan_penanganan_risiko_malaria"`

	KomplikasiNifas string `json:"komplikasi_nifas"`
	TindakanSaran   string `json:"tindakan_saran"`

	NamaPemeriksaParaf string     `gorm:"type:varchar(255)" json:"nama_pemeriksa_paraf"`
	TanggalKembali     *time.Time `gorm:"type:date" json:"tanggal_kembali"`

	CreatedAt time.Time `json:"created_at"`
}

func (PelayananIbuNifas) TableName() string {
	return "pelayanan_ibu_nifas"
}
