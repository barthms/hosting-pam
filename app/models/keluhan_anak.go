package models

import (
	"time"

	"gorm.io/gorm"
)

type KeluhanAnak struct {
	ID             int32          `json:"id" gorm:"primaryKey;autoIncrement"`
	AnakID         int32          `json:"anak_id" gorm:"not null;index"`
	Anak           *Anak          `json:"anak,omitempty" gorm:"foreignKey:AnakID"`
	Tanggal        time.Time      `json:"tanggal" gorm:"type:date;not null"`
	TanggalKembali *time.Time     `json:"tanggal_kembali,omitempty" gorm:"type:date"`
	Keluhan        string         `json:"keluhan" gorm:"type:text;not null"`
	Tindakan       string         `json:"tindakan" gorm:"type:text"`
	Pemeriksa      string         `json:"pemeriksa" gorm:"type:varchar(100)"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
}

func (KeluhanAnak) TableName() string {
	return "keluhan_anak"
}
