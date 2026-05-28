// package models

// import "time"

// type PemeriksaanDokterTrimester1 struct {
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

// 	HPHT                *time.Time `gorm:"type:date" json:"hpht"`
// 	KeteraturanHaid     string     `gorm:"type:varchar(20)" json:"keteraturan_haid"`
// 	UmurHamilHPHTMinggu *int       `json:"umur_hamil_hpht_minggu"`
// 	HPLBerdasarkanHPHT  *time.Time `gorm:"type:date" json:"hpl_berdasarkan_hpht"`
// 	UmurHamilUSGMinggu  *int       `json:"umur_hamil_usg_minggu"`
// 	HPLBerdasarkanUSG   *time.Time `gorm:"type:date" json:"hpl_berdasarkan_usg"`

// 	USGJumlahGS                 string   `gorm:"type:varchar(20)" json:"usg_jumlah_gs"`
// 	USGDiameterGS_cm            *float64 `gorm:"type:decimal(5,2)" json:"usg_diameter_gs_cm"`
// 	USGDiameterGSMinggu         *int     `json:"usg_diameter_gs_minggu"`
// 	USGDiameterGSHari           *int     `json:"usg_diameter_gs_hari"`
// 	USGJumlahBayi               string   `gorm:"type:varchar(20)" json:"usg_jumlah_bayi"`
// 	USGCRL_cm                   *float64 `gorm:"type:decimal(5,2)" json:"usg_crl_cm"`
// 	USGCRLMinggu                *int     `json:"usg_crl_minggu"`
// 	USGCRLHari                  *int     `json:"usg_crl_hari"`
// 	USGLetakProdukKehamilan     string   `gorm:"type:varchar(50)" json:"usg_letak_produk_kehamilan"`
// 	USGPulsasiJantung           string   `gorm:"type:varchar(20)" json:"usg_pulsasi_jantung"`
// 	USGKecurigaanTemuanAbnormal string   `gorm:"type:varchar(10)" json:"usg_kecurigaan_temuan_abnormal"`
// 	USGKeteranganTemuanAbnormal string   `json:"usg_keterangan_temuan_abnormal"`

// 	CreatedAt time.Time `json:"created_at"`
// }

// func (PemeriksaanDokterTrimester1) TableName() string {
// 	return "pemeriksaan_dokter_trimester_1"
// }


package models

import "time"

type PemeriksaanDokterTrimester1 struct {
	// FIX: tambahkan column:id_trimester1 agar GORM tahu nama kolom PK yang benar di DB
	ID                        int32      `gorm:"primaryKey;column:id_trimester1" json:"id"`
	KehamilanID               int32      `gorm:"not null;index" json:"kehamilan_id"`
	Kehamilan                 *Kehamilan `gorm:"foreignKey:KehamilanID;references:ID;onDelete:CASCADE" json:"kehamilan,omitempty"`
	NamaDokter                string     `gorm:"type:varchar(255)" json:"nama_dokter"`
	TanggalPeriksa            *time.Time `gorm:"type:date" json:"tanggal_periksa"`
	KonsepAnamnesaPemeriksaan string     `json:"konsep_anamnesa_pemeriksaan"`
	GambarUSG                 string     `gorm:"column:gambar_usg;type:text" json:"gambar_usg"` // Base64 image data

	FisikKonjungtiva string `gorm:"type:varchar(20)" json:"fisik_konjungtiva"`
	FisikSklera      string `gorm:"type:varchar(20)" json:"fisik_sklera"`
	FisikKulit       string `gorm:"type:varchar(20)" json:"fisik_kulit"`
	FisikLeher       string `gorm:"type:varchar(20)" json:"fisik_leher"`
	FisikGigiMulut   string `gorm:"type:varchar(20)" json:"fisik_gigi_mulut"`
	FisikTHT         string `gorm:"type:varchar(20)" json:"fisik_tht"`
	FisikDadaJantung string `gorm:"type:varchar(20)" json:"fisik_dada_jantung"`
	FisikDadaParu    string `gorm:"type:varchar(20)" json:"fisik_dada_paru"`
	FisikPerut       string `gorm:"type:varchar(20)" json:"fisik_perut"`
	FisikTungkai     string `gorm:"type:varchar(20)" json:"fisik_tungkai"`

	HPHT                *time.Time `gorm:"type:date" json:"hpht"`
	KeteraturanHaid     string     `gorm:"type:varchar(20)" json:"keteraturan_haid"`
	UmurHamilHPHTMinggu *int       `json:"umur_hamil_hpht_minggu"`
	HPLBerdasarkanHPHT  *time.Time `gorm:"type:date" json:"hpl_berdasarkan_hpht"`
	UmurHamilUSGMinggu  *int       `json:"umur_hamil_usg_minggu"`
	HPLBerdasarkanUSG   *time.Time `gorm:"type:date" json:"hpl_berdasarkan_usg"`

	USGJumlahGS                 string   `gorm:"type:varchar(20)" json:"usg_jumlah_gs"`
	USGDiameterGS_cm            *float64 `gorm:"type:decimal(5,2)" json:"usg_diameter_gs_cm"`
	USGDiameterGSMinggu         *int     `json:"usg_diameter_gs_minggu"`
	USGDiameterGSHari           *int     `json:"usg_diameter_gs_hari"`
	USGJumlahBayi               string   `gorm:"type:varchar(20)" json:"usg_jumlah_bayi"`
	USGCRL_cm                   *float64 `gorm:"type:decimal(5,2)" json:"usg_crl_cm"`
	USGCRLMinggu                *int     `json:"usg_crl_minggu"`
	USGCRLHari                  *int     `json:"usg_crl_hari"`
	USGLetakProdukKehamilan     string   `gorm:"type:varchar(50)" json:"usg_letak_produk_kehamilan"`
	USGPulsasiJantung           string   `gorm:"type:varchar(20)" json:"usg_pulsasi_jantung"`
	USGKecurigaanTemuanAbnormal string   `gorm:"type:varchar(10)" json:"usg_kecurigaan_temuan_abnormal"`
	USGKeteranganTemuanAbnormal string   `json:"usg_keterangan_temuan_abnormal"`

	CreatedAt time.Time `json:"created_at"`
}

func (PemeriksaanDokterTrimester1) TableName() string {
	return "pemeriksaan_dokter_trimester_1"
}