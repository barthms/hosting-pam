package models

import "time"

type EdukasiTrimester struct {
	ID         int32     `json:"id" gorm:"primaryKey;autoIncrement"`
	Trimester  string    `json:"trimester" gorm:"type:varchar(10)"`
	Kategori   string    `json:"kategori" gorm:"type:varchar(100)"`
	Judul      string    `json:"judul" gorm:"type:varchar(255)"`
	Isi        string    `json:"isi" gorm:"type:text"`
	GambarURL  string    `json:"gambar_url" gorm:"type:text"`
	CreatedAt  time.Time `json:"created_at" gorm:"->"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"->"`
}

func (EdukasiTrimester) TableName() string {
	return "edukasi_trimester"
}