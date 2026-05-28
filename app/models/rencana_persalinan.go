package models

import "time"

type RencanaPersalinan struct {
	IDRencanaPersalinan int32      `gorm:"primaryKey" json:"id_rencana_persalinan"`
	KehamilanID         int32      `gorm:"not null;index" json:"kehamilan_id"`
	Kehamilan           *Kehamilan `gorm:"foreignKey:KehamilanID;references:ID" json:"kehamilan,omitempty"`

	NamaIbuPernyataan        string `gorm:"type:varchar(255)" json:"nama_ibu_pernyataan"`
	AlamatIbuPernyataan      string `json:"alamat_ibu_pernyataan"`
	PerkiraanBulanPersalinan string `gorm:"type:varchar(20)" json:"perkiraan_bulan_persalinan"`
	PerkiraanTahunPersalinan *int   `json:"perkiraan_tahun_persalinan"`

	Fasyankes1NamaTenaga    string `gorm:"type:varchar(255)" json:"fasyankes_1_nama_tenaga"`
	Fasyankes1NamaFasilitas string `gorm:"type:varchar(255)" json:"fasyankes_1_nama_fasilitas"`
	Fasyankes2NamaTenaga    string `gorm:"type:varchar(255)" json:"fasyankes_2_nama_tenaga"`
	Fasyankes2NamaFasilitas string `gorm:"type:varchar(255)" json:"fasyankes_2_nama_fasilitas"`

	SumberDanaPersalinan string `gorm:"type:varchar(255)" json:"sumber_dana_persalinan"`

	Kendaraan1Nama string `gorm:"type:varchar(255)" json:"kendaraan_1_nama"`
	Kendaraan1HP   string `gorm:"type:varchar(20)" json:"kendaraan_1_hp"`
	Kendaraan2Nama string `gorm:"type:varchar(255)" json:"kendaraan_2_nama"`
	Kendaraan2HP   string `gorm:"type:varchar(20)" json:"kendaraan_2_hp"`
	Kendaraan3Nama string `gorm:"type:varchar(255)" json:"kendaraan_3_nama"`
	Kendaraan3HP   string `gorm:"type:varchar(20)" json:"kendaraan_3_hp"`

	MetodeKontrasepsiPilihan string `gorm:"type:varchar(255)" json:"metode_kontrasepsi_pilihan"`

	DonorGolonganDarah string `gorm:"type:varchar(5)" json:"donor_golongan_darah"`
	DonorRhesus        string `gorm:"type:varchar(10)" json:"donor_rhesus"`
	Donor1Nama         string `gorm:"type:varchar(255)" json:"donor_1_nama"`
	Donor1HP           string `gorm:"type:varchar(20)" json:"donor_1_hp"`
	Donor2Nama         string `gorm:"type:varchar(255)" json:"donor_2_nama"`
	Donor2HP           string `gorm:"type:varchar(20)" json:"donor_2_hp"`
	Donor3Nama         string `gorm:"type:varchar(255)" json:"donor_3_nama"`
	Donor3HP           string `gorm:"type:varchar(20)" json:"donor_3_hp"`
	Donor4Nama         string `gorm:"type:varchar(255)" json:"donor_4_nama"`
	Donor4HP           string `gorm:"type:varchar(20)" json:"donor_4_hp"`

	TanggalPernyataan    *time.Time `gorm:"type:date" json:"tanggal_pernyataan"`
	NamaSuamiKeluargaTTD string     `json:"nama_suami_keluarga_ttd"`
	NamaibuTTD           string     `json:"nama_ibu_hamil_ttd"`
	NamaBidanDokterTTD   string     `json:"nama_bidan_dokter_ttd"`

	CreatedAt time.Time `json:"created_at"`
}

func (RencanaPersalinan) TableName() string {
	return "rencana_persalinan"
}
