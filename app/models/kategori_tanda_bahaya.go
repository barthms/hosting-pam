package models

import (
	"time"

	"gorm.io/gorm"
)

type KategoriTandaBahaya struct {
	ID           int32          `json:"id" gorm:"primaryKey;autoIncrement"`
	TipeLembar   string         `json:"tipe_lembar" gorm:"type:varchar(10);check:tipe_lembar IN ('A','B','Umum');not null"`
	Gejala       string         `json:"gejala" gorm:"type:text;not null"`
	KategoriUsia string         `json:"kategori_usia" gorm:"type:varchar(50);not null"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

func (KategoriTandaBahaya) TableName() string {
	return "kategori_tanda_bahaya"
}
