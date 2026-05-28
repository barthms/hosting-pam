package models

import (
	"time"

	"gorm.io/gorm"
)

type KunjunganAnak struct {
	ID              int32             `json:"id" gorm:"primaryKey;autoIncrement"`
	AnakID          int32             `json:"anak_id" gorm:"not null;index"`
	Anak            *Anak             `json:"anak,omitempty" gorm:"foreignKey:AnakID"`
	Tanggal         time.Time         `json:"tanggal" gorm:"not null"`
	KategoriUmurID  int32             `json:"kategori_umur_id" gorm:"not null;index"`
	KategoriUmur    *KategoriUmur     `json:"kategori_umur,omitempty" gorm:"foreignKey:KategoriUmurID"`
	PeriodeID       int32             `json:"periode_id" gorm:"not null;index"`
	Periode         *PeriodeKunjungan `json:"periode,omitempty" gorm:"foreignKey:PeriodeID"`
	Lokasi          string            `json:"lokasi" gorm:"not null"`
	DetailPelayanan []DetailPelayanan `json:"detail_pelayanan" gorm:"foreignKey:KunjunganAnakID;constraint:OnDelete:CASCADE"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	DeletedAt       gorm.DeletedAt    `json:"-" gorm:"index"`
}

func (KunjunganAnak) TableName() string {
	return "kunjungan_anak"
}

// func (k *KunjunganAnak) BeforeCreate(tx *gorm.DB) error {
// 	if k.ID == "" {
// 		k.ID = uuid.New().String()
// 	}
// 	return nil
// }
