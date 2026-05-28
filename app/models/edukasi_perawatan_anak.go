package models

import (
	"time"

	"gorm.io/gorm"
)

type EdukasiPerawatanAnak struct {
	ID        int32          `json:"id" gorm:"primaryKey;autoIncrement"`
	Judul     string         `json:"judul" gorm:"type:varchar(255);not null"`
	GambarURL string         `json:"gambar_url" gorm:"type:text"`
	IsiKonten string         `json:"isi_konten" gorm:"type:text"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (EdukasiPerawatanAnak) TableName() string {
	return "edukasi_perawatan_anak"
}
