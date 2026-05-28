package models

import "time"

type RingkasanPelayananPersalinan struct {
	ID          int32      `gorm:"primaryKey" json:"id"`
	KehamilanID int32      `gorm:"not null;index" json:"kehamilan_id"`
	Kehamilan   *Kehamilan `gorm:"foreignKey:KehamilanID;references:ID" json:"kehamilan,omitempty"`

	TanggalMelahirkan        *time.Time `gorm:"type:date" json:"tanggal_melahirkan"`
	PukulMelahirkan          *time.Time `json:"pukul_melahirkan"`
	UmurKehamilanMinggu      *int       `json:"umur_kehamilan_minggu"`
	PenolongProsesMelahirkan string     `gorm:"type:varchar(100)" json:"penolong_proses_melahirkan"`
	CaraMelahirkan           string     `gorm:"type:varchar(50)" json:"cara_melahirkan"`
	KeadaanIbu               string     `gorm:"type:varchar(100)" json:"keadaan_ibu"`
	KeadaanIbuDetailSakit    string     `json:"keadaan_ibu_detail_sakit"`
	KBPascaMelahirkan        string     `gorm:"type:varchar(100)" json:"kb_pasca_melahirkan"`
	KeteranganTambahanIbu    string     `json:"keterangan_tambahan_ibu"`

	BayiAnakKe          int       `json:"bayi_anak_ke"`
	BayiBeratLahirGram  *float64  `json:"bayi_berat_lahir_gram"`
	BayiPanjangBadanCm  *float64  `json:"bayi_panjang_badan_cm"`
	BayiLingkarKepalaCm *float64  `json:"bayi_lingkar_kepala_cm"`
	BayiJenisKelamin    string    `gorm:"type:varchar(50)" json:"bayi_jenis_kelamin"`

	KondisiBayiSegeraMenangis        bool   `json:"kondisi_bayi_segera_menangis"`
	KondisiBayiMenangisBeberapaSaat  bool   `json:"kondisi_bayi_menangis_beberapa_saat"`
	KondisiBayiTidakMenangis         bool   `json:"kondisi_bayi_tidak_menangis"`
	KondisiBayiSeluruhTubuhKemerahan bool   `json:"kondisi_bayi_seluruh_tubuh_kemerahan"`
	KondisiBayiAnggotaGerakKebiruan  bool   `json:"kondisi_bayi_anggota_gerak_kebiruan"`
	KondisiBayiSeluruhTubuhBiru      bool   `json:"kondisi_bayi_seluruh_tubuh_biru"`
	KondisiBayiKelainanBawaan        bool   `json:"kondisi_bayi_kelainan_bawaan"`
	KondisiBayiKelainanBawaanDetail  string `json:"kondisi_bayi_kelainan_bawaan_detail"`
	KondisiBayiMeninggal             bool   `json:"kondisi_bayi_meninggal"`

	AsuhanIMD1JamPertama       bool `json:"asuhan_imd_1_jam_pertama"`
	AsuhanSuntikanVitaminK1    bool `json:"asuhan_suntikan_vitamin_k1"`
	AsuhanSalepMataAntibiotika bool `json:"asuhan_salep_mata_antibiotika"`
	AsuhanImunisasiHB0         bool `json:"asuhan_imunisasi_hb0"`

	KeteranganTambahanBayi string `json:"keterangan_tambahan_bayi"`

	CreatedAt time.Time `json:"created_at"`
}

func (RingkasanPelayananPersalinan) TableName() string {
	return "ringkasan_pelayanan_persalinan"
}
