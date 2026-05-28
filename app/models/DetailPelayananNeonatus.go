package models

import (
	"time"

	"gorm.io/gorm"
)

type DetailPelayananNeonatus struct {
	ID               int32             `json:"id" gorm:"primaryKey;autoIncrement"`
	NeonatusID  int32             `json:"neonatus_id" gorm:"not null;index"`
	Neonatus    *Neonatus    `json:"neonatus,omitempty" gorm:"foreignKey:NeonatusID"`
	JenisPelayananID int32             `json:"jenis_pelayanan_id" gorm:"not null;index"`
	JenisPelayanan   *JenisPelayanan   `json:"jenis_pelayanan,omitempty" gorm:"foreignKey:JenisPelayananID"`
	Nilai            string            `json:"nilai" gorm:"not null"`
	Keterangan       string            `json:"keterangan"`
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
	DeletedAt        gorm.DeletedAt    `json:"-" gorm:"index"`
}

func (DetailPelayananNeonatus) TableName() string {
	return "detail_pelayanan_neonatus"
}

// func (k *DetailPelayanan) BeforeCreate(tx *gorm.DB) error {
// 	if k.ID == "" {
// 		k.ID = uuid.New().String()
// 	}
// 	return nil
// }
