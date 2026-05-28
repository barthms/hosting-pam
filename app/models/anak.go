package models

import (
	"time"

	"gorm.io/gorm"
)

type Anak struct {
	ID              int32          `gorm:"primaryKey;autoIncrement" json:"id"`
	KehamilanID     int32          `json:"kehamilan_id" gorm:"not null;index;constraint:OnDelete:CASCADE"`
	Kehamilan       *Kehamilan     `json:"kehamilan,omitempty" gorm:"foreignKey:KehamilanID;constraint:OnDelete:CASCADE"`
	PendudukID      int32          `json:"penduduk_id" gorm:"not null;index;constraint:OnDelete:CASCADE"`
	Penduduk        *Kependudukan  `json:"penduduk,omitempty" gorm:"foreignKey:PendudukID;constraint:OnDelete:CASCADE"`
	BeratLahirKg    *float64       `json:"berat_lahir_kg,omitempty"`
	TinggiLahirCm   *float64       `json:"tinggi_lahir_cm,omitempty"`
	AnakKe          int32          `json:"anak_ke"`
	LingkarKepalaCm *float64       `json:"lingkar_kepala_cm,omitempty"`
	NamaIbu         string         `json:"nama_ibu"`
	NamaAyah        string         `json:"nama_ayah"`
	IbuID           int32          `json:"ibu_id"`
	TanggalLahir    *time.Time     `json:"tanggal_lahir,omitempty" gorm:"column:tanggal_lahir;type:date"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`
}

func (Anak) TableName() string { return "anak" }
