package models

import "time"

type EdukasiMenyusuiASI struct {
	ID         int32     `json:"id" gorm:"primaryKey;autoIncrement"`
	Judul      string    `json:"judul" gorm:"type:varchar(255);not null"`
	Isi        string    `json:"isi" gorm:"type:text"`
	ManfaatASI string    `json:"manfaat_asi" gorm:"type:text"`
	Cara       string    `json:"cara" gorm:"type:text"`
	Masalah    string    `json:"masalah" gorm:"type:text"`
	Solusi     string    `json:"solusi" gorm:"type:text"`
	GambarURL  string    `json:"gambar_url" gorm:"type:text"`
	CreatedAt  time.Time `json:"created_at" gorm:"->"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"->"`
}

func (EdukasiMenyusuiASI) TableName() string {
	return "edukasi_menyusui_asi"
}
