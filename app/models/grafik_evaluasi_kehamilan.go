package models

import "time"

// PemeriksaanKehamilan [V]
// Kehamilan [V]
// PemantauanIbuHamil [V]
// SkriningPreeklampsia [V]
// PenjelasanHasilGrafik [V]

type GrafikEvaluasiKehamilan struct {
	ID          int32      `gorm:"primaryKey" json:"id"`
	KehamilanID int32      `gorm:"not null;index" json:"kehamilan_id"`
	Kehamilan   *Kehamilan `gorm:"foreignKey:KehamilanID;references:ID" json:"kehamilan,omitempty"`

	TanggalBulanTahun *time.Time `gorm:"type:date" json:"tanggal_bulan_tahun"` // tanggal_periksa - PemeriksaanKehamilan
	UsiaGestasiMinggu *int       `json:"usia_gestasi_minggu"` // uk_kehamilan_saat_ini - kehamilan

	TinggiFundusUteriCm     *float64 `gorm:"type:decimal(4,1)" json:"tinggi_fundus_uteri_cm"` // tinnggi_rahim - PemeriksaanKehamilan
	DenyutJantungBayiXMenit *int     `json:"denyut_jantung_bayi_x_menit"` // usg_djj_nilai - PemeriksaanDokterTrimester3

	TekananDarahSistole  *int `json:"tekanan_darah_sistole"` // tekanan_darah - PemeriksaanKehamilan
	TekananDarahDiastole *int `json:"tekanan_darah_diastole"` // tekanan_darah - PemeriksaanKehamilan
	NadiPerMenit         *int `json:"nadi_per_menit"` 

	GerakanBayi *string `gorm:"type:varchar(50)" json:"gerakan_bayi"` // gerakan_bayi_kurang - PemantauanIbuHamil
	UrinProtein *string `gorm:"type:varchar(50)" json:"urin_protein"` // tes_lab_protein_urine - PemeriksaanKehamilan
	UrinReduksi *string `gorm:"type:varchar(50)" json:"urin_reduksi"`// lab_urin_reduksi_hasil - PemeriksaanLanjutanTrimester3 

	Hemoglobin        *float64 `gorm:"type:decimal(4,1)" json:"hemoglobin"` // tes_lab_hb - PemeriksaanKehamilan
	TabletTambahDarah *int     `json:"tablet_tambah_darah"` // tablet_tambah_darah - PemeriksaanKehamilan

	Kalsium *string `gorm:"type:varchar(50)" json:"kalsium"`
	Aspirin *string `gorm:"type:varchar(50)" json:"aspirin"` // kesimpulan_skrining_preeklampsia - SkriningPreeklampsia

	PenjelasanHasilGrafik *string `json:"penjelasan_hasil_grafik"` // catatan_penjelasan_grafik - PenjelasanHasilGrafik
	KategoriRisiko        *string `gorm:"type:varchar(20)" json:"kategori_risiko"`// kesimpulan_skrining_preeklampsia - SkriningPreeklampsia

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (GrafikEvaluasiKehamilan) TableName() string {
	return "grafik_evaluasi_kehamilan"
}
