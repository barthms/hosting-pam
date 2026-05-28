package models

import "time"

type PemeriksaanLanjutanTrimester3 struct {
	IDLanjutanT3 int32      `gorm:"primaryKey" json:"id_lanjutan_t3"`
	KehamilanID  int32      `gorm:"not null;index" json:"kehamilan_id"`
	Kehamilan    *Kehamilan `gorm:"foreignKey:KehamilanID;references:ID" json:"kehamilan,omitempty"`

	HasilUSGCatatan string `json:"hasil_usg_catatan"`

	TanggalLab            *time.Time `gorm:"type:date" json:"tanggal_lab"`
	LabHemoglobinHasil    *float64   `gorm:"type:decimal(4,1)" json:"lab_hemoglobin_hasil"`
	LabHemoglobinRencana  string     `json:"lab_hemoglobin_rencana_tindak_lanjut"`
	LabProteinUrinHasil   *int       `json:"lab_protein_urin_hasil"`
	LabProteinUrinRencana string     `json:"lab_protein_urin_rencana_tindak_lanjut"`
	LabUrinReduksiHasil   string     `gorm:"type:varchar(20)" json:"lab_urin_reduksi_hasil"`
	LabUrinReduksiRencana string     `json:"lab_urin_reduksi_rencana_tindak_lanjut"`

	TanggalSkriningJiwa      *time.Time `gorm:"type:date" json:"tanggal_skrining_jiwa"`
	SkriningJiwaHasil        string     `gorm:"type:varchar(10)" json:"skrining_jiwa_hasil"`
	SkriningJiwaTindakLanjut string     `gorm:"type:varchar(20)" json:"skrining_jiwa_tindak_lanjut"`
	SkriningJiwaPerluRujukan string     `gorm:"type:varchar(10)" json:"skrining_jiwa_perlu_rujukan"`

	RencanaKonsultasiGizi          bool   `json:"rencana_konsultasi_gizi"`
	RencanaKonsultasiKebidanan     bool   `json:"rencana_konsultasi_kebidanan"`
	RencanaKonsultasiAnak          bool   `json:"rencana_konsultasi_anak"`
	RencanaKonsultasiPenyakitDalam bool   `json:"rencana_konsultasi_penyakit_dalam"`
	RencanaKonsultasiNeurologi     bool   `json:"rencana_konsultasi_neurologi"`
	RencanaKonsultasiTHT           bool   `json:"rencana_konsultasi_tht"`
	RencanaKonsultasiPsikiatri     bool   `json:"rencana_konsultasi_psikiatri"`
	RencanaKonsultasiLainLain      string `json:"rencana_konsultasi_lain_lain"`

	RencanaProsesMelahirkan string `gorm:"type:varchar(50)" json:"rencana_proses_melahirkan"`

	RencanaKontrasepsiAKDR         bool `json:"rencana_kontrasepsi_akdr"`
	RencanaKontrasepsiPil          bool `json:"rencana_kontrasepsi_pil"`
	RencanaKontrasepsiSuntik       bool `json:"rencana_kontrasepsi_suntik"`
	RencanaKontrasepsiSteril       bool `json:"rencana_kontrasepsi_steril"`
	RencanaKontrasepsiMAL          bool `json:"rencana_kontrasepsi_mal"`
	RencanaKontrasepsiImplan       bool `json:"rencana_kontrasepsi_implan"`
	RencanaKontrasepsiBelumMemilih bool `json:"rencana_kontrasepsi_belum_memilih"`

	KebutuhanKonseling string `gorm:"type:varchar(10)" json:"kebutuhan_konseling"`
	Penjelasan         string `json:"penjelasan"`

	KesimpulanRekomendasiTempatMelahirkan string `gorm:"type:varchar(20)" json:"kesimpulan_rekomendasi_tempat_melahirkan"`

	CreatedAt time.Time `json:"created_at"`
}

func (PemeriksaanLanjutanTrimester3) TableName() string {
	return "pemeriksaan_lanjutan_trimester_3"
}
