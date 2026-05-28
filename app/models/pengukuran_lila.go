package models

import (
	"time"

	"gorm.io/gorm"
)

type PengukuranLila struct {
	ID             int32          `json:"id" gorm:"primaryKey;autoIncrement"`
	AnakID         int32          `json:"anak_id" gorm:"not null;index"`
	Anak           *Anak          `json:"anak,omitempty" gorm:"foreignKey:AnakID"`
	Bulanke        int            `json:"bulan" gorm:"not null"`
	Tanggal        time.Time      `json:"tanggal" gorm:"not null;"`
	HasilLila      float64        `json:"hasil_lila" gorm:"type:decimal(5,2);not null"`
	KategoriRisiko string         `json:"kategori_risiko" gorm:"type:varchar(20);not null"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
}

func (PengukuranLila) TableName() string {
	return "pengukuran_lila"
}

// func (P *PengukuranLila) BeforeCreate(tx *gorm.DB) error {
// 	if P.ID == "" {
// 		P.ID = uuid.New().String()
// 	}
// 	return nil
// }
