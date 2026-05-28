package models

import (
	"time"

	"gorm.io/gorm"
)

type KunjunganVitamin struct {
	ID        int32                    `json:"id" gorm:"primaryKey;autoIncrement"`
	AnakID    int32                    `json:"anak_id" gorm:"not null;index"`
	Anak      *Anak                    `json:"anak,omitempty" gorm:"foreignKey:AnakID"`
	Tanggal   time.Time                `json:"tanggal" gorm:"not null"`
	Detail    []DetailPelayananVitamin `json:"detail" gorm:"foreignKey:KunjunganVitaminID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time                `json:"created_at"`
	UpdatedAt time.Time                `json:"updated_at"`
	DeletedAt gorm.DeletedAt           `json:"-" gorm:"index"`
}

func (KunjunganVitamin) TableName() string {
	return "kunjungan_vitamin"
}
