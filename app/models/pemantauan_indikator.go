package models

import "time"

type PemantauanIndikator struct {
	ID           int32      `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	KategoriUsia string     `gorm:"column:kategori_usia;type:varchar(50);not null;index" json:"kategori_usia"`
	Deskripsi    string     `gorm:"column:deskripsi;type:text;not null" json:"deskripsi"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	IsDeleted    *time.Time `json:"is_deleted,omitempty"`
}

func (PemantauanIndikator) TableName() string { return "pemantauan_indikator" }
