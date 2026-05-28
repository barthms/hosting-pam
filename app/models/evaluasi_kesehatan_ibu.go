package models

import "time"

type EvaluasiKesehatanIbu struct {
	ID                 int32      `gorm:"primaryKey" json:"id"`
	KehamilanID        int32      `gorm:"not null;index" json:"kehamilan_id"`
	Kehamilan          *Kehamilan `gorm:"foreignKey:KehamilanID;references:ID" json:"kehamilan,omitempty"`
	NamaDokter         string     `gorm:"type:varchar(255)" json:"nama_dokter"`
	TanggalPeriksa     *time.Time `gorm:"type:date" json:"tanggal_periksa"`
	FasilitasKesehatan string     `gorm:"type:varchar(255)" json:"fasilitas_kesehatan"`

	TbCm        *float64 `gorm:"type:decimal(5,2)" json:"tb_cm"`
	BbKg        *float64 `gorm:"type:decimal(5,2)" json:"bb_kg"`
	IMTKategori string   `gorm:"type:varchar(20)" json:"imt_kategori"`
	LilaCm      *float64 `gorm:"type:decimal(5,2)" json:"lila_cm"`

	StatusTT1               bool   `json:"status_tt_1"`
	StatusTT2               bool   `json:"status_tt_2"`
	StatusTT3               bool   `json:"status_tt_3"`
	StatusTT4               bool   `json:"status_tt_4"`
	StatusTT5               bool   `json:"status_tt_5"`
	ImunisasiLainnyaCovid19 string `json:"imunisasi_lainnya_covid19"`

	RiwayatAlergi           bool   `json:"riwayat_alergi"`
	RiwayatAsma             bool   `json:"riwayat_asma"`
	RiwayatAutoimun         bool   `json:"riwayat_autoimun"`
	RiwayatDiabetes         bool   `json:"riwayat_diabetes"`
	RiwayatHepatitisB       bool   `json:"riwayat_hepatitis_b"`
	RiwayatHipertensi       bool   `json:"riwayat_hipertensi"`
	RiwayatJantung          bool   `json:"riwayat_jantung"`
	RiwayatJiwa             bool   `json:"riwayat_jiwa"`
	RiwayatSifilis          bool   `json:"riwayat_sifilis"`
	RiwayatTb               bool   `json:"riwayat_tb"`
	RiwayatKesehatanLainnya string `json:"riwayat_kesehatan_lainnya"`

	PerilakuAktivitasFisikKurang bool   `json:"perilaku_aktivitas_fisik_kurang"`
	PerilakuAlkohol              bool   `json:"perilaku_alkohol"`
	PerilakuKosmetikBerbahaya    bool   `json:"perilaku_kosmetik_berbahaya"`
	PerilakuMerokok              bool   `json:"perilaku_merokok"`
	PerilakuObatTeratogenik      bool   `json:"perilaku_obat_teratogenik"`
	PerilakuPolaMakanBerisiko    bool   `json:"perilaku_pola_makan_berisiko"`
	PerilakuLainnya              string `json:"perilaku_lainnya"`

	KeluargaAlergi     bool   `json:"keluarga_alergi"`
	KeluargaAsma       bool   `json:"keluarga_asma"`
	KeluargaAutoimun   bool   `json:"keluarga_autoimun"`
	KeluargaDiabetes   bool   `json:"keluarga_diabetes"`
	KeluargaHepatitisB bool   `json:"keluarga_hepatitis_b"`
	KeluargaHipertensi bool   `json:"keluarga_hipertensi"`
	KeluargaJantung    bool   `json:"keluarga_jantung"`
	KeluargaJiwa       bool   `json:"keluarga_jiwa"`
	KeluargaSifilis    bool   `json:"keluarga_sifilis"`
	KeluargaTb         bool   `json:"keluarga_tb"`
	KeluargaLainnya    string `json:"keluarga_lainnya"`

	InspeksiPorsio  string `gorm:"type:varchar(20)" json:"inspeksi_porsio"`
	InspeksiUretra  string `gorm:"type:varchar(20)" json:"inspeksi_uretra"`
	InspeksiVagina  string `gorm:"type:varchar(20)" json:"inspeksi_vagina"`
	InspeksiVulva   string `gorm:"type:varchar(20)" json:"inspeksi_vulva"`
	InspeksiFluksus string `gorm:"type:varchar(10)" json:"inspeksi_fluksus"`
	InspeksiFluor   string `gorm:"type:varchar(10)" json:"inspeksi_fluor"`

	CreatedAt time.Time `json:"created_at"`
}

func (EvaluasiKesehatanIbu) TableName() string {
	return "evaluasi_kesehatan_ibu"
}
