package models

import "time"

type EdukasiNifas struct {
	ID          int32     `json:"id" gorm:"primaryKey;autoIncrement"`
	Judul       string    `json:"judul" gorm:"type:varchar(255);not null"`
	Isi         string    `json:"isi" gorm:"type:text"`
	Perawatan   string    `json:"perawatan" gorm:"type:text"`
	TandaBahaya string    `json:"tanda_bahaya" gorm:"type:text"`
	GambarURL   string    `json:"gambar_url" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at" gorm:"->"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"->"`
}

func (EdukasiNifas) TableName() string {
	return "edukasi_nifas"
}
