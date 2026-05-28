package models

import (
	"time"

	"gorm.io/gorm"
)

type KategoriUmur struct {
	ID           int32  `json:"id" gorm:"primaryKey;autoIncrement"`
	KategoriUmur string `json:"kategori_umur" gorm:"not null;unique"`
	MinValue     int    `json:"min_value" gorm:"not null"`
	MaxValue     int    `json:"max_value" gorm:"not null"`
	Unit         string `json:"unit" gorm:"type:varchar(10);not null"`
	// Unit      		string     			`json:"unit" gorm:"type:enum('jam','hari','bulan','tahun');not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (KategoriUmur) TableName() string {
	return "Kategori_umur"
}

// func (k *KategoriUmur) BeforeCreeate(tx *gorm.DB) error{
// 	if k.ID == ""{
// 		k.ID = uuid.New().String()
// 	}
// 	return nil
// }
