package models

import "time"

type PemeriksaanKehamilan struct {
	IDPeriksa       int32      `gorm:"primaryKey" json:"id_periksa"`
	KehamilanID     int32      `gorm:"not null;index" json:"kehamilan_id"`
	Kehamilan       *Kehamilan `gorm:"foreignKey:KehamilanID;references:ID" json:"kehamilan,omitempty"`
	Trimester       string     `gorm:"type:varchar(10)" json:"trimester"`
	KunjunganKe     int32      `json:"kunjungan_ke"`
	MingguKehamilan int32      `json:"minggu_kehamilan"`

	TanggalPeriksa *time.Time `gorm:"type:date" json:"tanggal_periksa"`
	TempatPeriksa  string     `gorm:"type:varchar(255)" json:"tempat_periksa"`

	BeratBadan        *float64 `gorm:"type:decimal(5,2)" json:"berat_badan"`
	TinggiBadan       *float64 `gorm:"type:decimal(5,2)" json:"tinggi_badan"`
	LingkarLenganAtas *float64 `gorm:"type:decimal(5,2)" json:"lingkar_lengan_atas"`

	// Tekanan Darah
	Sistole  int32 `json:"sistole"`
	Diastole int32 `json:"diastole"`

	TinggiRahim *float64 `gorm:"type:decimal(5,2)" json:"tinggi_rahim"`

	// DJJ (TAMBAHKAN FIELD NUMERIK INI)
	LetakDenyutJantungBayi string `gorm:"type:text" json:"letak_denyut_jantung_bayi"`
	DenyutJantungJanin     int32  `json:"denyut_jantung_janin"` // <--- WAJIB UNTUK GRAFIK

	StatusImunisasiTetanus string `gorm:"type:varchar(50)" json:"status_imunisasi_tetanus"`
	Konseling              string `gorm:"type:text" json:"konseling"`
	SkriningDokter         string `gorm:"type:text" json:"skrining_dokter"`
	TabletTambahDarah      *int32 `json:"tablet_tambah_darah"`

	TesLabHb           *float64 `gorm:"type:decimal(4,1)" json:"tes_lab_hb"`
	TesGolonganDarah   string   `gorm:"type:varchar(5)" json:"tes_golongan_darah"`
	TesLabProteinUrine string   `gorm:"type:varchar(50)" json:"tes_lab_protein_urine"`
	TesLabGulaDarah    *int32   `json:"tes_lab_gula_darah"`

	USG              string `gorm:"type:text" json:"usg"`
	TripelEliminasi  string `gorm:"type:varchar(100)" json:"tripel_eliminasi"`
	TataLaksanaKasus string `gorm:"type:text" json:"tata_laksana_kasus"`

	// TAMBAHKAN FIELD HASIL SKORING (OPSIONAL TAPI SANGAT MEMBANTU)
	SkorRisiko   int32  `json:"skor_risiko"`
	StatusRisiko string `gorm:"type:varchar(20)" json:"status_risiko"`
	DetailRisiko string `gorm:"type:text" json:"detail_risiko"`

	CreatedAt time.Time `json:"created_at"`
}

func (PemeriksaanKehamilan) TableName() string {
	return "pemeriksaan_kehamilan"
}
