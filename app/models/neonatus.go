package models

import (
	"time"

	"gorm.io/gorm"
)

type Neonatus struct {
	ID                int32             `json:"id" gorm:"primaryKey;autoIncrement"`
	AnakID            int32             `json:"anak_id" gorm:"not null;index"`
	Anak              *Anak             `json:"anak,omitempty" gorm:"foreignKey:AnakID"`
	Tanggal           time.Time         `json:"tanggal" gorm:"not null"`
	KategoriUmurID    int32             `json:"kategori_umur_id" gorm:"not null;index"`
	KategoriUmur      *KategoriUmur     `json:"kategori_umur,omitempty" gorm:"foreignKey:KategoriUmurID"`
	PeriodeID         int32             `json:"periode_id" gorm:"not null;index"`
	Periode           *PeriodeKunjungan `json:"periode,omitempty" gorm:"foreignKey:PeriodeID"`
	TenagaKesehatanID int32             `json:"tenaga_kesehatan_id" gorm:"not null;index"`
	TenagaKesehatan   *User             `json:"tenaga_kesehatan,omitempty" gorm:"foreignKey:TenagaKesehatanID"`
	// Lokasi          string            `json:"lokasi" gorm:"not null"`
	DetailPelayanan []DetailPelayananNeonatus `json:"detail_pelayanan" gorm:"foreignKey:NeonatusID;constraint:OnDelete:CASCADE"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	DeletedAt       gorm.DeletedAt    `json:"-" gorm:"index"`
}

func (Neonatus) TableName() string {
	return "Neonatus"
}

// func (k *KunjunganAnak) BeforeCreate(tx *gorm.DB) error {
// 	if k.ID == "" {
// 		k.ID = uuid.New().String()
// 	}
// 	return nil
// }
