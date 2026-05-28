package models

import (
	"time"

	"gorm.io/gorm"
)

type CatatanPelayanan struct {
	ID                int32          `json:"id" gorm:"primaryKey;autoIncrement"`
	AnakID            int32          `json:"anak_id" gorm:"not null;index"`
	Anak              *Anak          `json:"anak,omitempty" gorm:"foreignKey:AnakID"`
	TenagaKesehatanID int32          `json:"tenaga_kesehatan_id" gorm:"not null;index"`
	TenagaKesehatan   *User          `json:"tenaga_kesehatan,omitempty" gorm:"foreignKey:TenagaKesehatanID"`
	Tanggal_periksa   time.Time      `json:"tanggal_periksa" gorm:"not null"`
	Tanggal_kembali   time.Time      `json:"tanggal_kembali" gorm:"not null"`
	CatatanPelayanan  string         `json:"catatan_pelayanan" gorm:"type:text;not null"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`
}

func (CatatanPelayanan) TableName() string {
	return "catatan_pelayanan"
}

// func (C *CatatanPelayanan) BeforeCreate(tx *gorm.DB) error {
// 	if C.ID == "" {
// 		C.ID = uuid.New().String()
// 	}
// 	return nil
// }
