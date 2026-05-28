package models

import (
	"time"

	"gorm.io/gorm"
)

type KehadiranImunisasi struct {
	ID        int32                      `json:"id" gorm:"primaryKey;autoIncrement"`
	AnakID    int32                      `json:"anak_id" gorm:"not null;index"`
	Anak      *Anak                      `json:"anak,omitempty" gorm:"foreignKey:AnakID"`
	Detail    []DetailPelayananImunisasi `json:"detail" gorm:"foreignKey:KunjunganImunisasiID;constraint:OnDelete:CASCADE"`
	Bulanke   int                        `json:"bulan_ke" gorm:"column:bulan_ke;not null"`
	CreatedAt time.Time                  `json:"created_at"`
	UpdatedAt time.Time                  `json:"updated_at"`
	DeletedAt gorm.DeletedAt             `json:"-" gorm:"index"`
}

func (KehadiranImunisasi) TableName() string {
	return "kehadiran_imunisasi"
}

// func (k *KehadiranImunisasi) BeforeCreate(tx *gorm.DB) error {
// 	if k.ID == "" {
// 		k.ID = uuid.New().String()
// 	}
// 	return nil
// }
