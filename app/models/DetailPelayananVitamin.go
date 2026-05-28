package models

import (
	"time"

	"gorm.io/gorm"
)

type DetailPelayananVitamin struct {
	ID                 int32             `json:"id" gorm:"primaryKey;autoIncrement"`
	KunjunganVitaminID int32             `json:"kunjungan_vitamin_id" gorm:"not null;index"`
	KunjunganVitamin   *KunjunganVitamin `json:"kunjungan_vitamin,omitempty" gorm:"foreignKey:KunjunganVitaminID"`
	JenisPelayananID   int32             `json:"jenis_pelayanan_id" gorm:"not null;index"`
	JenisPelayanan     *JenisPelayanan   `json:"jenis_pelayanan,omitempty" gorm:"foreignKey:JenisPelayananID"`
	Keterangan         string            `json:"keterangan"`
	CreatedAt          time.Time         `json:"created_at"`
	UpdatedAt          time.Time         `json:"updated_at"`
	DeletedAt          gorm.DeletedAt    `json:"-" gorm:"index"`
}

func (DetailPelayananVitamin) TableName() string {
	return "detail_pelayanan_vitamin"
}

// func (k *DetailPelayanan) BeforeCreate(tx *gorm.DB) error {
// 	if k.ID == "" {
// 		k.ID = uuid.New().String()
// 	}
// 	return nil
// }
