package models

import (
	"time"

	"gorm.io/gorm"
)

type LogTTDMMS struct {
	ID           int32          `gorm:"primaryKey;autoIncrement" json:"id"`
	KehamilanID  int32          `gorm:"not null;index:idx_log_ttd_unique,unique" json:"kehamilan_id"`
	Kehamilan    *Kehamilan     `gorm:"foreignKey:KehamilanID;references:ID;constraint:OnDelete:CASCADE" json:"kehamilan,omitempty"`
	BulanKe      int32          `gorm:"not null;index:idx_log_ttd_unique,unique" json:"bulan_ke"`
	HariKe       int32          `gorm:"not null;index:idx_log_ttd_unique,unique" json:"hari_ke"`
	SudahDiminum bool           `gorm:"not null;default:false" json:"sudah_diminum"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

func (LogTTDMMS) TableName() string {
	return "log_ttd_mms"
}
