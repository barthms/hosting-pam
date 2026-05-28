package models

import (
	"time"

	"gorm.io/gorm"
)

type ASI struct {
	ID                int32          `json:"id" gorm:"primaryKey;autoIncrement"`
	KunjunganGiziID   int32          `json:"kunjungan_gizi_id" gorm:"uniqueIndex;not null;index"`
	FrekuensiMenyusui int32          `json:"frekuensi_menyusui" gorm:"not null"` //jumlah mennyusui
	PosisiMenyusui    string         `json:"posisi_menyusui"`	//baik atau tidak
	ASIPerah          string         `json:"asi_perah" gorm:"column:asi_perah;not null"` //ya atau tidak
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`
}

func (ASI) TableName() string {
	return "Pelayanan_Asi"
}

// func (A *ASI) BeforeCreate(tx *gorm.DB) error {
// 	if A.ID == "" {
// 		A.ID = uuid.New().String()
// 	}
// 	return nil
// }
