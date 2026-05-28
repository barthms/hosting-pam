package models

import (
	"time"

	"gorm.io/gorm"
)

type ProsesMelahirkan struct {
	ID                     int32          `gorm:"primaryKey;autoIncrement" json:"id"`
	KehamilanID            int32          `gorm:"not null;index" json:"kehamilan_id"`
	Kehamilan              *Kehamilan     `gorm:"foreignKey:KehamilanID;references:ID;constraint:OnDelete:CASCADE" json:"kehamilan,omitempty"`

	TandaPersalinan        bool `gorm:"not null;default:false" json:"tanda_persalinan"`
	ProsesMelahirkan       bool `gorm:"not null;default:false" json:"proses_melahirkan"`
	HakIbuPendamping       bool `gorm:"not null;default:false" json:"hak_ibu_pendamping"`
	HakIbuPosisiMelahirkan bool `gorm:"not null;default:false" json:"hak_ibu_posisi_melahirkan"`
	Mulas                  bool `gorm:"not null;default:false" json:"mulas"`
	TeknikMengurangiNyeri  bool `gorm:"not null;default:false" json:"teknik_mengurangi_nyeri"`
	IMDKontakKulit         bool `gorm:"not null;default:false" json:"imd_kontak_kulit"`

	CreatedAt              time.Time      `json:"created_at"`
	UpdatedAt              time.Time      `json:"updated_at"`
	DeletedAt              gorm.DeletedAt `json:"-" gorm:"index"`
}

func (ProsesMelahirkan) TableName() string {
	return "proses_melahirkan"
}