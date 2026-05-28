package models

import (
	"time"

	"gorm.io/gorm"
)

type AturanPelayanan struct {
	ID               int32           `json:"id" gorm:"primaryKey;autoIncrement"`
	JenisPelayananID int32           `json:"jenis_pelayanan_id" gorm:"not null;index"`
	JenisPelayanan   *JenisPelayanan `json:"jenis_pelayanan,omitempty" gorm:"foreignKey:JenisPelayananID"`
	UmurMinBulan     int             `json:"umur_min_bulan" gorm:"not null"`
	UmurMaxBulan     int             `json:"umur_max_bulan" gorm:"not null"`
	Bulan            int             `json:"bulan" gorm:"not null"` // 2=Feb, 8=Agustus
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
	DeletedAt        gorm.DeletedAt  `json:"-" gorm:"index"`
}
