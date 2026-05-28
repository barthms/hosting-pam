package models

import (
	"time"

	"gorm.io/gorm"
)

type Pertumbuhan struct {
	ID           int32          `json:"id" gorm:"primaryKey;autoIncrement"`
	AnakID       int32          `json:"anak_id" gorm:"not null;index"`
	Anak         *Anak          `json:"anak,omitempty" gorm:"foreignKey:AnakID"`
	Bulanke      int            `json:"bulan" gorm:"not null"`
	Beratbadan   float64        `json:"berat_badan" gorm:"type:decimal(5,2);not null"`
	Panjangbadan float64        `json:"panjang_badan" gorm:"type:decimal(5,2);not null"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

func (Pertumbuhan) TableName() string {
	return "pertumbuhan_Anak"
}

// func (P *Pertumbuhan) BeforeCreate(tx *gorm.DB) error {
// 	if P.ID == "" {
// 		P.ID = uuid.New().String()
// 	}
// 	return nil
// }
