package models

import (
	"time"

	"gorm.io/gorm"
)

type PersiapanMelahirkan struct {
	ID                    int32          `gorm:"primaryKey;autoIncrement" json:"id"`
	KehamilanID           int32          `gorm:"not null;index;uniqueIndex" json:"kehamilan_id"`
	Kehamilan             *Kehamilan     `gorm:"foreignKey:KehamilanID;references:ID;constraint:OnDelete:CASCADE" json:"kehamilan,omitempty"`

	PerkiraanPersalinan   bool `gorm:"not null;default:false" json:"perkiraan_persalinan"`
	PendampingPersalinan  bool `gorm:"not null;default:false" json:"pendamping_persalinan"`
	DanaPersalinan        bool `gorm:"not null;default:false" json:"dana_persalinan"`
	StatusJKN             bool `gorm:"not null;default:false" json:"status_jkn"`
	FaskesPersalinan      bool `gorm:"not null;default:false" json:"faskes_persalinan"`
	PendonorDarah         bool `gorm:"not null;default:false" json:"pendonor_darah"`
	Transportasi          bool `gorm:"not null;default:false" json:"transportasi"`
	MetodeKB              bool `gorm:"not null;default:false" json:"metode_kb"`
	ProgramP4K            bool `gorm:"not null;default:false" json:"program_p4k"`
	DokumenPenting        bool `gorm:"not null;default:false" json:"dokumen_penting"`

	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `json:"-" gorm:"index"`
}

func (PersiapanMelahirkan) TableName() string {
	return "persiapan_melahirkan"
}