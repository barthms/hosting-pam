package models

import (
	"time"

	"gorm.io/gorm"
)

// KategoriLingkungan - Master data kategori: "Kesehatan Lingkungan", "Keselamatan Lingkungan"
type KategoriLingkungan struct {
	ID        int32          `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama      string         `gorm:"type:varchar(100);not null" json:"nama"`
	Deskripsi string         `gorm:"type:text" json:"deskripsi"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Indikator []IndikatorLingkungan `gorm:"foreignKey:KategoriLingkunganID" json:"indikator,omitempty"`
}

func (KategoriLingkungan) TableName() string {
	return "kategori_lingkungan"
}

// IndikatorLingkungan - Master data pertanyaan/indikator per kategori
type IndikatorLingkungan struct {
	ID                   int32          `gorm:"primaryKey;autoIncrement" json:"id"`
	KategoriLingkunganID int32          `gorm:"not null;index" json:"kategori_lingkungan_id"`
	Pertanyaan           string         `gorm:"type:text;not null" json:"pertanyaan"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"index" json:"-"`

	Kategori KategoriLingkungan `gorm:"foreignKey:KategoriLingkunganID;references:ID" json:"-"`
}

func (IndikatorLingkungan) TableName() string {
	return "indikator_lingkungan"
}

// LembarLingkungan - Header transaksi pengisian kesehatan lingkungan
type LembarLingkungan struct {
	ID             int32     `gorm:"primaryKey;autoIncrement" json:"id"`
	IbuID          int32     `gorm:"not null;index" json:"ibu_id"`
	TanggalPeriksa time.Time `gorm:"type:date;not null" json:"tanggal_periksa"`
	Pemeriksa      string    `gorm:"type:varchar(100)" json:"pemeriksa"` // Nama Bidan/Kader
	Catatan        string    `gorm:"type:text" json:"catatan"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	Ibu           Ibu                `gorm:"foreignKey:IbuID;references:IDIbu" json:"ibu,omitempty"`
	DetailJawaban []DetailLingkungan `gorm:"foreignKey:LembarLingkunganID" json:"detail_jawaban,omitempty"`
}

func (LembarLingkungan) TableName() string {
	return "lembar_lingkungan"
}

// DetailLingkungan - Jawaban per indikator
type DetailLingkungan struct {
	ID                   int32     `gorm:"primaryKey;autoIncrement" json:"id"`
	LembarLingkunganID   int32     `gorm:"not null;index" json:"lembar_lingkungan_id"`
	IndikatorLingkunganID int32     `gorm:"not null;index" json:"indikator_lingkungan_id"`
	IsOk                 bool      `gorm:"not null" json:"is_ok"` // Jawaban Ya/Tidak (Sehat/Tidak Sehat)
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`

	Lembar    LembarLingkungan    `gorm:"foreignKey:LembarLingkunganID;references:ID;constraint:OnDelete:CASCADE" json:"-"`
	Indikator IndikatorLingkungan `gorm:"foreignKey:IndikatorLingkunganID;references:ID" json:"indikator,omitempty"`
}

func (DetailLingkungan) TableName() string {
	return "detail_lingkungan"
}

// ==================== REQUEST MODELS ====================

type IndikatorLingkunganRequest struct {
	Pertanyaan string `json:"pertanyaan" validate:"required"`
}

type KategoriLingkunganRequest struct {
	Nama      string                       `json:"nama" validate:"required"`
	Deskripsi string                       `json:"deskripsi"`
	Indikator []IndikatorLingkunganRequest `json:"indikator"`
}

type DetailLingkunganRequest struct {
	IndikatorLingkunganID int32 `json:"indikator_lingkungan_id" validate:"required"`
	IsOk                 bool  `json:"is_ok"`
}

type LembarLingkunganRequest struct {
	IbuID          int32                     `json:"ibu_id" validate:"required"`
	TanggalPeriksa string                    `json:"tanggal_periksa" validate:"required"` // YYYY-MM-DD
	Pemeriksa      string                    `json:"pemeriksa"`
	Catatan        string                    `json:"catatan"`
	DetailJawaban  []DetailLingkunganRequest `json:"detail_jawaban" validate:"required,dive"`
}
