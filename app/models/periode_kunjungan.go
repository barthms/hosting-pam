package models

import (
	"time"

	"gorm.io/gorm"
)

type PeriodeKunjungan struct {
	ID             int32         `json:"id" gorm:"primaryKey;autoIncrement"`
	Nama           string        `json:"nama" gorm:"not null"`
	KategoriUmurID int32         `json:"kategori_umur_id" gorm:"not null; index"`
	KategoriUmur   *KategoriUmur `json:"kategori_umur,omitempty" gorm:"foreignKey:KategoriUmurID"`
	MinValue       int           `json:"min_value" gorm:"not null"`
	MaxValue       int           `json:"max_value" gorm:"not null"`
	// Unit      			string     			`json:"unit" gorm:"type:enum('jam','hari','bulan','tahun');not null"`
	Unit      string         `json:"unit" gorm:"type:varchar(10);not null"`
	Urutan    int            `json:"urutan" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (PeriodeKunjungan) TableName() string {
	return "periode_kunjungan"
}

// func (k *PeriodeKunjungan) BeforeCreate(tx *gorm.DB) error {
// 	if k.ID == "" {
// 		k.ID = uuid.New().String()
// 	}
// 	return nil
// }
