// package models

// import "time"

// type PemeriksaanDokterTrimester3 struct {
// 	ID                        int32      `gorm:"primaryKey" json:"id"`
// 	KehamilanID               int32      `gorm:"not null;index" json:"kehamilan_id"`
// 	Kehamilan                 *Kehamilan `gorm:"foreignKey:KehamilanID;references:ID" json:"kehamilan,omitempty"`
// 	NamaDokter                string     `gorm:"type:varchar(255)" json:"nama_dokter"`
// 	TanggalPeriksa            *time.Time `gorm:"type:date" json:"tanggal_periksa"`
// 	KonsepAnamnesaPemeriksaan string     `json:"konsep_anamnesa_pemeriksaan"`

// 	FisikKonjungtiva string `gorm:"type:varchar(20)" json:"fisik_konjungtiva"`
// 	FisikSklera      string `gorm:"type:varchar(20)" json:"fisik_sklera"`
// 	FisikKulit       string `gorm:"type:varchar(20)" json:"fisik_kulit"`
// 	FisikLeher       string `gorm:"type:varchar(20)" json:"fisik_leher"`
// 	FisikGigiMulut   string `gorm:"type:varchar(20)" json:"fisik_gigi_mulut"`
// 	FisikTHT         string `gorm:"type:varchar(20)" json:"fisik_tht"`
// 	FisikDadaJantung string `gorm:"type:varchar(20)" json:"fisik_dada_jantung"`
// 	FisikDadaParu    string `gorm:"type:varchar(20)" json:"fisik_dada_paru"`
// 	FisikPerut       string `gorm:"type:varchar(20)" json:"fisik_perut"`
// 	FisikTungkai     string `gorm:"type:varchar(20)" json:"fisik_tungkai"`

// 	USGTrimester3Dilakukan                   string `gorm:"type:varchar(10)" json:"usg_trimester_3_dilakukan"`
// 	UKBerdasarkanUSGTrimester1Minggu         *int   `json:"uk_berdasarkan_usg_trimester_1_minggu"`
// 	UKBerdasarkanHPHTMinggu                  *int   `json:"uk_berdasarkan_hpht_minggu"`
// 	UKBerdasarkanBiometriUSGTrimester3Minggu *int   `json:"uk_berdasarkan_biometri_usg_trimester_3_minggu"`
// 	SelisihUK3MingguAtauLebih                string `gorm:"type:varchar(10)" json:"selisih_uk_3_minggu_atau_lebih"`

// 	USGJumlahBayi     string `gorm:"type:varchar(20)" json:"usg_jumlah_bayi"`
// 	USGLetakBayi      string `gorm:"type:varchar(50)" json:"usg_letak_bayi"`
// 	USGPresentasiBayi string `gorm:"type:varchar(50)" json:"usg_presentasi_bayi"`
// 	USGKeadaanBayi    string `gorm:"type:varchar(20)" json:"usg_keadaan_bayi"`
// 	USGDJNilai        *int   `json:"usg_djj_nilai"`
// 	USGDJJStatus      string `gorm:"type:varchar(20)" json:"usg_djj_status"`
// 	USGLokasiPlasenta string `gorm:"type:varchar(50)" json:"usg_lokasi_plasenta"`

// 	USGCairanKetubanSDPCm  *float64 `gorm:"type:decimal(5,2)" json:"usg_cairan_ketuban_sdp_cm"`
// 	USGCairanKetubanStatus string   `gorm:"type:varchar(20)" json:"usg_cairan_ketuban_status"`

// 	BiometriBPDCm        *float64 `gorm:"type:decimal(5,2)" json:"biometri_bpd_cm"`
// 	BiometriBPDMinggu    *int     `json:"biometri_bpd_minggu"`
// 	BiometriHCCm         *float64 `gorm:"type:decimal(5,2)" json:"biometri_hc_cm"`
// 	BiometriHCMinggu     *int     `json:"biometri_hc_minggu"`
// 	BiometriACCm         *float64 `gorm:"type:decimal(5,2)" json:"biometri_ac_cm"`
// 	BiometriACMinggu     *int     `json:"biometri_ac_minggu"`
// 	BiometriFLCm         *float64 `gorm:"type:decimal(5,2)" json:"biometri_fl_cm"`
// 	BiometriFLMinggu     *int     `json:"biometri_fl_minggu"`
// 	BiometriEFWTBJGram   *int     `json:"biometri_efw_tbj_gram"`
// 	BiometriEFWTBJMinggu *int     `json:"biometri_efw_tbj_minggu"`

// 	USGKecurigaanTemuanAbnormal string `gorm:"type:varchar(10)" json:"usg_kecurigaan_temuan_abnormal"`
// 	USGKeteranganTemuanAbnormal string `json:"usg_keterangan_temuan_abnormal"`

// 	CreatedAt time.Time `json:"created_at"`
// }

// func (PemeriksaanDokterTrimester3) TableName() string {
// 	return "pemeriksaan_dokter_trimester_3"
// }


package models

import "time"

type PemeriksaanDokterTrimester3 struct {
	// FIX: tambahkan column:id_trimester3 agar GORM tahu nama kolom PK yang benar di DB
	ID                        int32      `gorm:"primaryKey;column:id_trimester3" json:"id"`
	KehamilanID               int32      `gorm:"not null;index" json:"kehamilan_id"`
	Kehamilan                 *Kehamilan `gorm:"foreignKey:KehamilanID;references:ID;onDelete:CASCADE" json:"kehamilan,omitempty"`
	NamaDokter                string     `gorm:"type:varchar(255)" json:"nama_dokter"`
	TanggalPeriksa            *time.Time `gorm:"type:date" json:"tanggal_periksa"`
	KonsepAnamnesaPemeriksaan string     `json:"konsep_anamnesa_pemeriksaan"`
	GambarUSG                 string     `gorm:"column:gambar_usg;type:text" json:"gambar_usg"` // Base64 image data
	FisikKonjungtiva          string     `gorm:"type:varchar(20)" json:"fisik_konjungtiva"`
	FisikSklera               string     `gorm:"type:varchar(20)" json:"fisik_sklera"`
	FisikKulit                string     `gorm:"type:varchar(20)" json:"fisik_kulit"`
	FisikLeher                string     `gorm:"type:varchar(20)" json:"fisik_leher"`
	FisikGigiMulut            string     `gorm:"type:varchar(20)" json:"fisik_gigi_mulut"`
	FisikTHT                  string     `gorm:"type:varchar(20)" json:"fisik_tht"`
	FisikDadaJantung          string     `gorm:"type:varchar(20)" json:"fisik_dada_jantung"`
	FisikDadaParu             string     `gorm:"type:varchar(20)" json:"fisik_dada_paru"`
	FisikPerut                string     `gorm:"type:varchar(20)" json:"fisik_perut"`
	FisikTungkai              string     `gorm:"type:varchar(20)" json:"fisik_tungkai"`

	// USG Trimester 3
	USGTrimester3Dilakukan                   string   `gorm:"type:varchar(10)" json:"usg_trimester_3_dilakukan"`
	UKBerdasarkanUSGTrimester1Minggu         *int     `json:"uk_berdasarkan_usg_trimester_1_minggu"`
	UKBerdasarkanHPHTMinggu                  *int     `json:"uk_berdasarkan_hpht_minggu"`
	UKBerdasarkanBiometriUSGTrimester3Minggu *int     `json:"uk_berdasarkan_biometri_usg_trimester_3_minggu"`
	SelisihUK3MingguAtauLebih                string   `gorm:"type:varchar(10)" json:"selisih_uk_3_minggu_atau_lebih"`
	USGJumlahBayi                            string   `gorm:"type:varchar(20)" json:"usg_jumlah_bayi"`
	USGLetakBayi                             string   `gorm:"type:varchar(50)" json:"usg_letak_bayi"`
	USGPresentasiBayi                        string   `gorm:"type:varchar(50)" json:"usg_presentasi_bayi"`
	USGKeadaanBayi                           string   `gorm:"type:varchar(20)" json:"usg_keadaan_bayi"`
	USGDJNilai                               *int     `json:"usg_djj_nilai"`
	USGDJJStatus                             string   `gorm:"type:varchar(20)" json:"usg_djj_status"`
	USGLokasiPlasenta                        string   `gorm:"type:varchar(50)" json:"usg_lokasi_plasenta"`
	USGCairanKetubanSDPCm                    *float64 `gorm:"type:decimal(5,2)" json:"usg_cairan_ketuban_sdp_cm"`
	USGCairanKetubanStatus                   string   `gorm:"type:varchar(20)" json:"usg_cairan_ketuban_status"`

	BiometriBPDCm        *float64 `gorm:"type:decimal(5,2)" json:"biometri_bpd_cm"`
	BiometriBPDMinggu    *int     `json:"biometri_bpd_minggu"`
	BiometriHCCm         *float64 `gorm:"type:decimal(5,2)" json:"biometri_hc_cm"`
	BiometriHCMinggu     *int     `json:"biometri_hc_minggu"`
	BiometriACCm         *float64 `gorm:"type:decimal(5,2)" json:"biometri_ac_cm"`
	BiometriACMinggu     *int     `json:"biometri_ac_minggu"`
	BiometriFLCm         *float64 `gorm:"type:decimal(5,2)" json:"biometri_fl_cm"`
	BiometriFLMinggu     *int     `json:"biometri_fl_minggu"`
	BiometriEFWTBJGram   *int     `json:"biometri_efw_tbj_gram"`
	BiometriEFWTBJMinggu *int     `json:"biometri_efw_tbj_minggu"`

	USGKecurigaanTemuanAbnormal string `gorm:"type:varchar(10)" json:"usg_kecurigaan_temuan_abnormal"`
	USGKeteranganTemuanAbnormal string `json:"usg_keterangan_temuan_abnormal"`

	// ========== LANJUTAN TRIMESTER 3 (digabung di sini) ==========
	HasilUSGCatatan                       string     `gorm:"type:text" json:"hasil_usg_catatan"`
	TanggalLab                            *time.Time `gorm:"type:date" json:"tanggal_lab"`
	LabHemoglobinHasil                    *float64   `gorm:"type:decimal(4,1)" json:"lab_hemoglobin_hasil"`
	LabHemoglobinRencana                  string     `json:"lab_hemoglobin_rencana_tindak_lanjut"`
	LabProteinUrinHasil                   *int       `json:"lab_protein_urin_hasil"`
	LabProteinUrinRencana                 string     `json:"lab_protein_urin_rencana_tindak_lanjut"`
	LabUrinReduksiHasil                   string     `gorm:"type:varchar(20)" json:"lab_urin_reduksi_hasil"`
	LabUrinReduksiRencana                 string     `json:"lab_urin_reduksi_rencana_tindak_lanjut"`
	TanggalSkriningJiwa                   *time.Time `gorm:"type:date" json:"tanggal_skrining_jiwa"`
	SkriningJiwaHasil                     string     `gorm:"type:varchar(10)" json:"skrining_jiwa_hasil"`
	SkriningJiwaTindakLanjut              string     `gorm:"type:varchar(20)" json:"skrining_jiwa_tindak_lanjut"`
	SkriningJiwaPerluRujukan              string     `gorm:"type:varchar(10)" json:"skrining_jiwa_perlu_rujukan"`
	RencanaKonsultasiGizi                 bool       `json:"rencana_konsultasi_gizi"`
	RencanaKonsultasiKebidanan            bool       `json:"rencana_konsultasi_kebidanan"`
	RencanaKonsultasiAnak                 bool       `json:"rencana_konsultasi_anak"`
	RencanaKonsultasiPenyakitDalam        bool       `json:"rencana_konsultasi_penyakit_dalam"`
	RencanaKonsultasiNeurologi            bool       `json:"rencana_konsultasi_neurologi"`
	RencanaKonsultasiTHT                  bool       `json:"rencana_konsultasi_tht"`
	RencanaKonsultasiPsikiatri            bool       `json:"rencana_konsultasi_psikiatri"`
	RencanaKonsultasiLainLain             string     `json:"rencana_konsultasi_lain_lain"`
	RencanaProsesMelahirkan               string     `gorm:"type:varchar(50)" json:"rencana_proses_melahirkan"`
	RencanaKontrasepsiAKDR                bool       `json:"rencana_kontrasepsi_akdr"`
	RencanaKontrasepsiPil                 bool       `json:"rencana_kontrasepsi_pil"`
	RencanaKontrasepsiSuntik              bool       `json:"rencana_kontrasepsi_suntik"`
	RencanaKontrasepsiSteril              bool       `json:"rencana_kontrasepsi_steril"`
	RencanaKontrasepsiMAL                 bool       `json:"rencana_kontrasepsi_mal"`
	RencanaKontrasepsiImplan              bool       `json:"rencana_kontrasepsi_implan"`
	RencanaKontrasepsiBelumMemilih        bool       `json:"rencana_kontrasepsi_belum_memilih"`
	KebutuhanKonseling                    string     `gorm:"type:varchar(10)" json:"kebutuhan_konseling"`
	Penjelasan                            string     `json:"penjelasan"`
	KesimpulanRekomendasiTempatMelahirkan string     `gorm:"type:varchar(20)" json:"kesimpulan_rekomendasi_tempat_melahirkan"`

	CreatedAt time.Time `json:"created_at"`
}

func (PemeriksaanDokterTrimester3) TableName() string {
	return "pemeriksaan_dokter_trimester_3"
}