package models

import (
	"time"

	"gorm.io/gorm"
)

type JenisPelayanan struct {
	ID   int32  `json:"id" gorm:"primaryKey;autoIncrement"`
	Nama string `json:"nama" gorm:"not null"`
	// TipeInput  string 			`json:"tipe_input" gorm:"type:enum('number','checkbox','text');not null"`
	TipeInput string          `json:"tipe_input" gorm:"type:varchar(20);not null"`
	GroupName string          `json:"group_name" gorm:"type:varchar(100)"`
	Section   string          `json:"section" gorm:"type:varchar(100)"` // contoh: "pemeriksaan_fisik", "pemeriksaan_penunjang", "imunisasi", "skrining"
	Kategori  []KategoriUmur  `json:"kategori,omitempty" gorm:"many2many:jenis_pelayanan_kategori;"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt gorm.DeletedAt  `json:"-" gorm:"index"`
}

func (JenisPelayanan) TableName() string {
	return "jenis_pelayanan"
}

// func (k *JenisPelayanan) BeforeCreate(tx *gorm.DB) error {
// 	if k.ID == "" {
// 		k.ID = uuid.New().String()
// 	}
// 	return nil
// }
