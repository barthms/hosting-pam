package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

// 1. MateriMPASI: Untuk informasi umum yang tidak selalu terikat usia spesifik
type MateriMPASI struct {
	ID        int32  `json:"id" gorm:"primaryKey;autoIncrement"`
	Judul     string `json:"judul" gorm:"type:varchar(255);not null"`
	Konten    string `json:"konten" gorm:"type:text;not null"`
	GambarURL string `json:"gambar_url" gorm:"type:text"`

	// Age range in months: NULL berarti berlaku untuk semua usia
	BulanMin *int `json:"bulan_min"` // NULL = applicable to all ages
	BulanMax *int `json:"bulan_max"` // NULL = applicable to all ages

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (MateriMPASI) TableName() string {
	return "materi_mpasi" // Tetap tanpa skema public
}

// Alias agar layer backend yang sudah memakai nama EdukasiMPASI tetap terintegrasi
type EdukasiMPASI = MateriMPASI

// 2. AturanPorsiMPASI: Tabel untuk Tekstur, Frekuensi, Porsi per age range
type AturanPorsiMPASI struct {
	ID        int32          `json:"id" gorm:"primaryKey;autoIncrement"`
	BulanMin  int            `json:"bulan_min" gorm:"not null"`
	BulanMax  int            `json:"bulan_max" gorm:"not null"`
	Tekstur   string         `json:"tekstur" gorm:"type:text;not null"`
	Frekuensi string         `json:"frekuensi" gorm:"type:text;not null"`
	Porsi     string         `json:"porsi" gorm:"type:text;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (AturanPorsiMPASI) TableName() string {
	return "aturan_porsi_mpasi"
}

// 3. JadwalHarianMPASI: Daftar waktu dan aktivitas
type JadwalHarianMPASI struct {
	ID        int32          `json:"id" gorm:"primaryKey;autoIncrement"`
	BulanMin  int            `json:"bulan_min" gorm:"not null"`
	BulanMax  int            `json:"bulan_max" gorm:"not null"`
	Waktu     string         `json:"waktu" gorm:"type:varchar(10);not null"`
	Aktivitas string         `json:"aktivitas" gorm:"type:varchar(100);not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (JadwalHarianMPASI) TableName() string {
	return "jadwal_harian_mpasi"
}

// 4. ResepMPASI: Tabel resep menggunakan Array PostgreSQL, organized by age range
type ResepMPASI struct {
	ID             int32          `json:"id" gorm:"primaryKey;autoIncrement"`
	BulanMin       int            `json:"bulan_min" gorm:"not null"`
	BulanMax       int            `json:"bulan_max" gorm:"not null"`
	Judul          string         `json:"judul" gorm:"type:varchar(255);not null"`
	Tipe           string         `json:"tipe" gorm:"type:varchar(50);not null"`
	GambarURL      string         `json:"gambar_url" gorm:"type:text"`
	WaktuPersiapan int32          `json:"waktu_persiapan" gorm:"not null"`
	Kalori         int32          `json:"kalori" gorm:"not null"`
	Porsi          string         `json:"porsi" gorm:"type:varchar(50);not null"`
	BahanBahan     pq.StringArray `json:"bahan_bahan" gorm:"type:text[];not null"`
	CaraMembuat    pq.StringArray `json:"cara_membuat" gorm:"type:text[];not null"`
	Manfaat        string         `json:"manfaat" gorm:"type:text"`
	Tips           string         `json:"tips" gorm:"type:text"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
}

func (ResepMPASI) TableName() string {
	return "resep_mpasi"
}
