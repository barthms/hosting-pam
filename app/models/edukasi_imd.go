package models

import "time"

type EdukasiIMD struct {
	ID        int32     `json:"id" gorm:"primaryKey;autoIncrement"`
	Judul     string    `json:"judul" gorm:"type:varchar(255);not null"`
	Isi       string    `json:"isi" gorm:"type:text"`
	Manfaat   string    `json:"manfaat" gorm:"type:text"`
	Langkah   string    `json:"langkah" gorm:"type:text"`
	GambarURL string    `json:"gambar_url" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"->"`
	UpdatedAt time.Time `json:"updated_at" gorm:"->"`
}

func (EdukasiIMD) TableName() string {
	return "edukasi_imd"
}
