package models

import (
	"time"

	"gorm.io/gorm"
)

// KategoriPemantauanIbu - Master data gejala/tanda bahaya untuk Ibu Nifas
type KategoriPemantauanIbu struct {
	ID        int32          `gorm:"primaryKey;autoIncrement" json:"id"`
	Gejala    string         `gorm:"type:text;not null" json:"gejala"`
	Deskripsi string         `gorm:"type:text" json:"deskripsi"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (KategoriPemantauanIbu) TableName() string {
	return "kategori_pemantauan_ibu"
}

// LembarPemantauanIbu - Header transaksi pemantauan harian Ibu Nifas
type LembarPemantauanIbu struct {
	ID             int32          `gorm:"primaryKey;autoIncrement" json:"id"`
	IbuID          int32          `gorm:"not null;index" json:"ibu_id"`
	PeriodeWaktu   int            `gorm:"not null" json:"periode_waktu"` // Hari ke-X setelah melahirkan
	TanggalPeriksa time.Time      `gorm:"type:date;not null" json:"tanggal_periksa"`
	Pemeriksa      string         `gorm:"type:varchar(100)" json:"pemeriksa"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`

	Ibu          Ibu                   `gorm:"foreignKey:IbuID;references:IDIbu" json:"ibu,omitempty"`
	DetailGejala []DetailPemantauanIbu `gorm:"foreignKey:LembarPemantauanIbuID" json:"detail_gejala,omitempty"`
}

func (LembarPemantauanIbu) TableName() string {
	return "lembar_pemantauan_ibu"
}

// DetailPemantauanIbu - Status checkbox per gejala untuk Ibu Nifas
type DetailPemantauanIbu struct {
	ID                    int32          `gorm:"primaryKey;autoIncrement" json:"id"`
	LembarPemantauanIbuID int32          `gorm:"not null;index" json:"lembar_pemantauan_ibu_id"`
	KategoriPemantauanID  int32          `gorm:"not null;index" json:"kategori_pemantauan_id"`
	IsTerjadi             bool           `gorm:"not null" json:"is_terjadi"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `gorm:"index" json:"-"`

	LembarPemantauan   LembarPemantauanIbu   `gorm:"foreignKey:LembarPemantauanIbuID;references:ID;constraint:OnDelete:CASCADE" json:"-"`
	KategoriPemantauan KategoriPemantauanIbu `gorm:"foreignKey:KategoriPemantauanID;references:ID" json:"kategori_pemantauan,omitempty"`
}

func (DetailPemantauanIbu) TableName() string {
	return "detail_pemantauan_ibu"
}

// ==================== REQUEST MODELS ====================

type DetailPemantauanIbuRequest struct {
	KategoriPemantauanID int32 `json:"kategori_pemantauan_id" validate:"required"`
	IsTerjadi            bool  `json:"is_terjadi"`
}

type LembarPemantauanIbuRequest struct {
	IbuID          int32                        `json:"ibu_id" validate:"required"`
	PeriodeWaktu   int                          `json:"periode_waktu" validate:"required"`
	TanggalPeriksa string                       `json:"tanggal_periksa" validate:"required"` // YYYY-MM-DD
	Pemeriksa      string                       `json:"pemeriksa"`
	DetailGejala   []DetailPemantauanIbuRequest `json:"detail_gejala" validate:"required,dive"`
}
