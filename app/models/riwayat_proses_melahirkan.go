package models

import "time"

type RiwayatProsesMelahirkan struct {
	ID          int32      `gorm:"primaryKey" json:"id"`
	KehamilanID int32      `gorm:"not null;index" json:"kehamilan_id"`
	Kehamilan   *Kehamilan `gorm:"foreignKey:KehamilanID;references:ID" json:"kehamilan,omitempty"`

	GGravida int `json:"g_gravida"`
	PPartus  int `json:"p_partus"`
	AAbortus int `json:"a_abortus"`

	HariMelahirkan    string     `gorm:"type:varchar(20)" json:"hari_melahirkan"`
	TanggalMelahirkan *time.Time `gorm:"type:date" json:"tanggal_melahirkan"`
	PukulMelahirkan   *time.Time `json:"pukul_melahirkan"`

	CaraMelahirkanSpontan   bool `json:"cara_melahirkan_spontan"`
	CaraMelahirkanSungsang  bool `json:"cara_melahirkan_sungsang"`
	TindakanEkstraksiVakum  bool `json:"tindakan_ekstraksi_vakum"`
	TindakanEkstraksiForsep bool `json:"tindakan_ekstraksi_forsep"`
	TindakanSC              bool `json:"tindakan_sc"`

	PenolongDokterSpesialis bool `json:"penolong_dokter_spesialis"`
	PenolongDokter          bool `json:"penolong_dokter"`
	PenolongBidan           bool `json:"penolong_bidan"`

	TaksiranMelahirkan            string `json:"taksiran_melahirkan"`
	FasyankesTempatMelahirkan     string `json:"fasyankes_tempat_melahirkan"`
	RujukanKeterangan             string `json:"rujukan_keterangan"`
	InisiasiMenyusuDiniKeterangan string `json:"inisiasi_menyusu_dini_keterangan"`

	CapKakiBayiImageURL []byte `json:"cap_kaki_bayi_image_url"`

	CreatedAt time.Time `json:"created_at"`
}

func (RiwayatProsesMelahirkan) TableName() string {
	return "riwayat_proses_melahirkan"
}
