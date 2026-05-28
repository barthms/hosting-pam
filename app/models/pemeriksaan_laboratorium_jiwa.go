package models

import "time"

type PemeriksaanLaboratoriumJiwa struct {
	IDLabJiwa   int32      `gorm:"primaryKey" json:"id_lab_jiwa"`
	KehamilanID int32      `gorm:"not null;index" json:"kehamilan_id"`
	Kehamilan   *Kehamilan `gorm:"foreignKey:KehamilanID;references:ID" json:"kehamilan,omitempty"`
	Trimester   int32      `gorm:"not null;default:1" json:"trimester"` // 1, 2, atau 3

	TanggalLab *time.Time `gorm:"type:date" json:"tanggal_lab"`

	// Pemeriksaan Darah
	LabHemoglobinHasil               *float64 `gorm:"type:decimal(4,1)" json:"lab_hemoglobin_hasil"`
	LabHemoglobinRencanaTindakLanjut string   `json:"lab_hemoglobin_rencana_tindak_lanjut"`
	LabGolonganDarahRhesusHasil      string   `gorm:"type:varchar(20)" json:"lab_golongan_darah_rhesus_hasil"`
	LabGolonganDarahRhesusRencana    string   `json:"lab_golongan_darah_rhesus_rencana_tindak_lanjut"`
	LabGulaDarahSewaktuHasil         *int     `json:"lab_gula_darah_sewaktu_hasil"`
	LabGulaDarahSewaktuRencana       string   `json:"lab_gula_darah_sewaktu_rencana_tindak_lanjut"`

	// Triple Eliminasi
	LabHIVHasil          string `gorm:"type:varchar(20)" json:"lab_hiv_hasil"`
	LabHIVRencana        string `json:"lab_hiv_rencana_tindak_lanjut"`
	LabSifilisHasil      string `gorm:"type:varchar(20)" json:"lab_sifilis_hasil"`
	LabSifilisRencana    string `json:"lab_sifilis_rencana_tindak_lanjut"`
	LabHepatitisBHasil   string `gorm:"type:varchar(20)" json:"lab_hepatitis_b_hasil"`
	LabHepatitisBRencana string `json:"lab_hepatitis_b_rencana_tindak_lanjut"`

	// Skrining Jiwa
	TanggalSkriningJiwa      *time.Time `gorm:"type:date" json:"tanggal_skrining_jiwa"`
	SkriningJiwaHasil        string     `gorm:"type:varchar(10)" json:"skrining_jiwa_hasil"`
	SkriningJiwaTindakLanjut string     `gorm:"type:varchar(20)" json:"skrining_jiwa_tindak_lanjut"`
	SkriningJiwaPerluRujukan string     `gorm:"type:varchar(10)" json:"skrining_jiwa_perlu_rujukan"`

	Kesimpulan  string `json:"kesimpulan"`
	Rekomendasi string `json:"rekomendasi"`

	CreatedAt time.Time `json:"created_at"`
}

func (PemeriksaanLaboratoriumJiwa) TableName() string {
	return "pemeriksaan_laboratorium_jiwa"
}
