package models

import "time"

type KeteranganLahir struct {
	ID          int32  `gorm:"primaryKey" json:"id"`
	IDIbuRelasi int32  `gorm:"not null;index" json:"id_ibu_relasi"`
	Ibu         *Ibu   `gorm:"foreignKey:IDIbuRelasi;references:IDIbu" json:"ibu,omitempty"`
	NomorSurat  string `gorm:"type:varchar(100)" json:"nomor_surat"`

	HariLahir    string     `gorm:"type:varchar(20)" json:"hari_lahir"`
	TanggalLahir *time.Time `gorm:"type:date" json:"tanggal_lahir"`
	PukulLahir   *time.Time `json:"pukul_lahir"`

	JenisKelamin      string `gorm:"type:varchar(50)" json:"jenis_kelamin"`
	JenisKelahiran    string `gorm:"type:varchar(100)" json:"jenis_kelahiran"`
	AnakKe            int    `json:"anak_ke"`
	UsiaGestasiMinggu int    `json:"usia_gestasi_minggu"`
	BeratLahirGram    int    `json:"berat_lahir_gram"`
	PanjangBadanCm    int    `json:"panjang_badan_cm"`
	LingkarKepalaCm   int    `json:"lingkar_kepala_cm"`

	LokasiPersalinan       string `json:"lokasi_persalinan"`
	AlamatLokasiPersalinan string `json:"alamat_lokasi_persalinan"`
	NamaBayiDiberiNama     string `json:"nama_bayi_diberi_nama"`

	NamaIbu           string `gorm:"type:varchar(255)" json:"nama_ibu"`
	UmurIbu           int    `json:"umur_ibu"`
	NIKIbu            string `gorm:"type:varchar(20)" json:"nik_ibu"`
	NamaAyah          string `gorm:"type:varchar(255)" json:"nama_ayah"`
	NIKAyah           string `gorm:"type:varchar(20)" json:"nik_ayah"`
	PekerjaanOrangTua string `gorm:"type:varchar(255)" json:"pekerjaan_orang_tua"`
	AlamatOrangTua    string `json:"alamat_orang_tua"`
	RWRTOrangTua      string `gorm:"type:varchar(20)" json:"rw_rt_orang_tua"`
	KecamatanOrangTua string `gorm:"type:varchar(100)" json:"kecamatan_orang_tua"`
	KabKotaOrangTua   string `gorm:"type:varchar(100)" json:"kab_kota_orang_tua"`

	TanggalSurat          *time.Time `gorm:"type:date" json:"tanggal_surat"`
	NamaSaksi1            string     `json:"nama_saksi_1"`
	NamaSaksi2            string     `json:"nama_saksi_2"`
	NamaPenolongKelahiran string     `json:"nama_penolong_kelahiran"`

	CreatedAt time.Time `json:"created_at"`
}

func (KeteranganLahir) TableName() string {
	return "keterangan_lahir"
}
