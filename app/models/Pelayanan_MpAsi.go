package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type MPASI struct {
	ID                  int32          `json:"id" gorm:"primaryKey;autoIncrement"`
	KunjunganGiziID     int32          `json:"kunjungan_gizi_id" gorm:"uniqueIndex;not null;index"`
	DiberikanMPASI      bool           `json:"diberikan_mp_asi" gorm:"type:bool;not null"` //ya atau tidak
	VariasiMPASI        datatypes.JSON `json:"variasi_mpasi"`
	JumlahmakanPerporsi string         `json:"jumlah_makan_perporsi" `   //misal "1 mangkuk / 2 sdm"
	FrekuensiMakan      string         `json:"frekuensi_makan_perhari" ` // misal "3x utama, 2x selingan"
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `json:"-" gorm:"index"`
}

func (MPASI) TableName() string {
	return "mp_asi"
}

// func (M *MPASI) BeforeCreate(tx *gorm.DB) error {
// 	if M.ID == "" {
// 		M.ID = uuid.New().String()
// 	}
// 	return nil
// }
