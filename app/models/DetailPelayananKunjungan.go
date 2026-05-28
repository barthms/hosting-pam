package models

import (
	"time"

	"gorm.io/gorm"
)

type DetailPelayanan struct {
	ID               int32             `json:"id" gorm:"primaryKey;autoIncrement"`
	KunjunganAnakID  int32             `json:"kunjungan_anak_id" gorm:"not null;index"`
	KunjunganAnak    *KunjunganAnak    `json:"kunjungan_anak,omitempty" gorm:"foreignKey:KunjunganAnakID"`
	JenisPelayananID int32             `json:"jenis_pelayanan_id" gorm:"not null;index"`
	JenisPelayanan   *JenisPelayanan   `json:"jenis_pelayanan,omitempty" gorm:"foreignKey:JenisPelayananID"`
	Nilai            string            `json:"nilai" gorm:"not null"`
	Keterangan       string            `json:"keterangan"`
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
	DeletedAt        gorm.DeletedAt    `json:"-" gorm:"index"`
}

func (DetailPelayanan) TableName() string {
	return "detail_pelayanan_anak"
}

// func (k *DetailPelayanan) BeforeCreate(tx *gorm.DB) error {
// 	if k.ID == "" {
// 		k.ID = uuid.New().String()
// 	}
// 	return nil
// }
