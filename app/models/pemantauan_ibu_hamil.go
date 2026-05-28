package models

import (
	"time"

	"gorm.io/gorm"
)

type PemantauanIbuHamil struct {
	ID                 int32          `gorm:"primaryKey;autoIncrement" json:"id"`
	KehamilanID        int32          `gorm:"not null;index:idx_pemantauan_hamil_unique,unique" json:"kehamilan_id"`
	Kehamilan          *Kehamilan     `gorm:"foreignKey:KehamilanID;references:ID;constraint:OnDelete:CASCADE" json:"kehamilan,omitempty"`

	MingguKehamilan    int32 `gorm:"not null;index:idx_pemantauan_hamil_unique,unique" json:"minggu_kehamilan"`

	DemamLebih2Hari    bool `gorm:"not null;default:false" json:"demam_lebih_2_hari"`
	SakitKepala        bool `gorm:"not null;default:false" json:"sakit_kepala"`
	CemasBerlebih      bool `gorm:"not null;default:false" json:"cemas_berlebih"`
	ResikoTB           bool `gorm:"not null;default:false" json:"resiko_tb"`
	GerakanBayiKurang  bool `gorm:"not null;default:false" json:"gerakan_bayi_kurang"`
	NyeriPerut         bool `gorm:"not null;default:false" json:"nyeri_perut"`
	CairanJalanLahir   bool `gorm:"not null;default:false" json:"cairan_jalan_lahir"`
	MasalahKemaluan    bool `gorm:"not null;default:false" json:"masalah_kemaluan"`
	DiareBerulang      bool `gorm:"not null;default:false" json:"diare_berulang"`

	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `json:"-" gorm:"index"`
}

func (PemantauanIbuHamil) TableName() string {
	return "pemantauan_ibu_hamil"
}