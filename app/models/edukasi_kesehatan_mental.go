package models

import "time"

type EdukasiKesehatanMental struct {
	ID          int32     `json:"id" gorm:"primaryKey;autoIncrement"`
	Judul       string    `json:"judul" gorm:"type:varchar(255);not null"`
	Isi         string    `json:"isi" gorm:"type:text"`
	TandaGejala string    `json:"tanda_gejala" gorm:"type:text"`
	Solusi      string    `json:"solusi" gorm:"type:text"`
	GambarURL   string    `json:"gambar_url" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at" gorm:"->"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"->"`
}

func (EdukasiKesehatanMental) TableName() string {
	return "edukasi_kesehatan_mental"
}
