package models

import "time"

type EdukasiTandaMelahirkan struct {
	ID        int32     `json:"id" gorm:"primaryKey;autoIncrement"`
	Judul     string    `json:"judul" gorm:"type:varchar(255);not null"`
	Isi       string    `json:"isi" gorm:"type:text"`
	Tanda     string    `json:"tanda" gorm:"type:text"`
	Tindakan  string    `json:"tindakan" gorm:"type:text"`
	GambarURL string    `json:"gambar_url" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"->"`
	UpdatedAt time.Time `json:"updated_at" gorm:"->"`
}

func (EdukasiTandaMelahirkan) TableName() string {
	return "edukasi_tanda_melahirkan"
}
