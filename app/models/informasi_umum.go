package models

import (
	"time"

	"gorm.io/gorm"
)

type InformasiUmum struct {
	ID               int32          `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Tipe             string         `gorm:"column:tipe;type:varchar(20);not null" json:"tipe"`
	Judul            string         `gorm:"column:judul;type:varchar(255);not null" json:"judul"`
	UmurTarget       string         `gorm:"column:umur_target;type:varchar(50)" json:"umur_target,omitempty"`
	DurasiBaca       string         `gorm:"column:durasi_baca;type:varchar(30)" json:"durasi_baca,omitempty"`
	Ringkasan        string         `gorm:"column:ringkasan;type:text" json:"ringkasan,omitempty"`
	Konten           string         `gorm:"column:konten;type:text;not null" json:"konten"`
	YangPerluDiingat string         `gorm:"column:yang_perlu_diingat;type:text" json:"yang_perlu_diingat,omitempty"`
	ThumbnailURL     string         `gorm:"column:thumbnail_url;type:text" json:"thumbnail_url,omitempty"`
	IsActive         bool           `gorm:"column:is_active;default:true" json:"is_active"`
	CreatedAt        time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}

// yang tidak ada dimiliki oleh pa 2
// materiinti
// halpenting

// yang tidak ada dimiliki oleh pa 3
// tipe, umurtarget, durasibaca,
// is active

// yang mirip tapi beda
// ringkasan -> deskripsi
// thumbnailurl -> gambarurl
// konten ->isikonten

func (InformasiUmum) TableName() string {
	return "informasi_umum"
}
