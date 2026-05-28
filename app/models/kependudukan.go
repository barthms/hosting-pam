package models

import "time"

type Kependudukan struct {
	IDKependudukan     int32      `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	KartuKeluargaID    *int64     `gorm:"column:kartu_keluarga_id" json:"kartu_keluarga_id,omitempty"`
	NIK                *string    `gorm:"column:nik;type:varchar(30);uniqueIndex" json:"nik,omitempty"`
	NamaLengkap        string     `gorm:"column:nama_lengkap;type:text;not null" json:"nama_lengkap"`
	JenisKelamin       string     `gorm:"column:jenis_kelamin;type:text" json:"jenis_kelamin"`
	TanggalLahir       time.Time  `gorm:"column:tanggal_lahir" json:"tanggal_lahir"`
	TempatLahir        string     `gorm:"column:tempat_lahir;type:text" json:"tempat_lahir"`
	GolonganDarah      string     `gorm:"column:golongan_darah;type:text" json:"golongan_darah"`
	Agama              string     `gorm:"column:agama;type:text" json:"agama"`
	StatusPerkawinan   string     `gorm:"column:status_perkawinan;type:text" json:"status_perkawinan"`
	PendidikanTerakhir string     `gorm:"column:pendidikan_terakhir;type:text" json:"pendidikan_terakhir"`
	Pekerjaan          string     `gorm:"column:pekerjaan;type:text" json:"pekerjaan"`
	BacaHuruf          string     `gorm:"column:baca_huruf;type:text" json:"baca_huruf"`
	KedudukanKeluarga  string     `gorm:"column:kedudukan_keluarga;type:text" json:"kedudukan_keluarga"`
	Dusun              string     `gorm:"column:dusun;type:text" json:"dusun"`
	Kecamatan          string     `gorm:"column:kecamatan;type:text" json:"kecamatan"`
	Desa               string     `gorm:"column:desa;type:text" json:"desa"`
	TanggalPenambahan  *time.Time `gorm:"column:tanggal_penambahan" json:"tanggal_penambahan,omitempty"`
	AsalPenduduk       string     `gorm:"column:asal_penduduk;type:text" json:"asal_penduduk"`
	TanggalPengurangan *time.Time `gorm:"column:tanggal_pengurangan" json:"tanggal_pengurangan,omitempty"`
	TujuanPindah       string     `gorm:"column:tujuan_pindah;type:text" json:"tujuan_pindah"`
	TempatMeninggal    string     `gorm:"column:tempat_meninggal;type:text" json:"tempat_meninggal"`
	Keterangan         string     `gorm:"column:keterangan;type:text" json:"keterangan"`
	CreatedAt          time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt          time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt          *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
}

func (Kependudukan) TableName() string { return "penduduk" }
