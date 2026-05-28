package models

type SkriningPreeklampsia struct {
	ID          int32      `gorm:"primaryKey" json:"id"`
	KehamilanID int32      `gorm:"not null;index" json:"kehamilan_id"`
	Kehamilan   *Kehamilan `gorm:"foreignKey:KehamilanID;references:ID" json:"kehamilan,omitempty"`

	AnamnesisMultiparaPasanganBaruSedang       bool `json:"anamnesis_multipara_pasangan_baru_sedang"`
	AnamnesisTeknologiReproduksiBerbantuSedang bool `json:"anamnesis_teknologi_reproduksi_berbantu_sedang"`
	AnamnesisUmurDiatas35TahunSedang           bool `json:"anamnesis_umur_diatas_35_tahun_sedang"`
	AnamnesisNuliparaSedang                    bool `json:"anamnesis_nulipara_sedang"`
	AnamnesisJarakKehamilanDiatas10TahunSedang bool `json:"anamnesis_jarak_kehamilan_diatas_10_tahun_sedang"`
	AnamnesisRiwayatPreeklampsiaKeluargaSedang bool `json:"anamnesis_riwayat_preeklampsia_keluarga_sedang"`
	AnamnesisObesitasIMTDiatas30Sedang         bool `json:"anamnesis_obesitas_imt_diatas_30_sedang"`

	AnamnesisRiwayatPreeklampsiaSebelumnyaTinggi bool `json:"anamnesis_riwayat_preeklampsia_sebelumnya_tinggi"`
	AnamnesisKehamilanMultipelTinggi             bool `json:"anamnesis_kehamilan_multipel_tinggi"`
	AnamnesisDiabetesDalamKehamilanTinggi        bool `json:"anamnesis_diabetes_dalam_kehamilan_tinggi"`
	AnamnesisHipertensiKronikTinggi              bool `json:"anamnesis_hipertensi_kronik_tinggi"`
	AnamnesisPenyakitGinjalTinggi                bool `json:"anamnesis_penyakit_ginjal_tinggi"`
	AnamnesisPenyakitAutoimunSLETinggi           bool `json:"anamnesis_penyakit_autoimun_sle_tinggi"`
	AnamnesisAntiPhospholipidSyndromeTinggi      bool `json:"anamnesis_anti_phospholipid_syndrome_tinggi"`

	FisikMAPDiatas90mmHg      bool `json:"fisik_map_diatas_90_mmhg"`
	FisikProteinuriaUrinCelup bool `json:"fisik_proteinuria_urin_celup"`

	KesimpulanSkriningPreeklampsia string `json:"kesimpulan_skrining_preeklampsia"`
}

func (SkriningPreeklampsia) TableName() string {
	return "skrining_preeklampsia"
}
