package models

import (
	"time"

	"gorm.io/gorm"
)

type PemantauanIbuNifas struct {
	ID                   int32          `gorm:"primaryKey;autoIncrement" json:"id"`
	KehamilanID          int32          `gorm:"not null;index:idx_pemantauan_nifas_unique,unique" json:"kehamilan_id"`
	Kehamilan            *Kehamilan     `gorm:"foreignKey:KehamilanID;references:ID;constraint:OnDelete:CASCADE" json:"kehamilan,omitempty"`

	HariKe               int32 `gorm:"not null;index:idx_pemantauan_nifas_unique,unique" json:"hari_ke"`

	PayudaraBermasalah   bool `gorm:"not null;default:false" json:"payudara_bermasalah"`
	GangguanBAK          bool `gorm:"not null;default:false" json:"gangguan_bak"`
	KelaminBermasalah    bool `gorm:"not null;default:false" json:"kelamin_bermasalah"`
	DarahNifasAbnormal   bool `gorm:"not null;default:false" json:"darah_nifas_abnormal"`
	PendarahanHebat      bool `gorm:"not null;default:false" json:"pendarahan_hebat"`
	Keputihan            bool `gorm:"not null;default:false" json:"keputihan"`
	Demam                bool `gorm:"not null;default:false" json:"demam"`
	SakitKepala          bool `gorm:"not null;default:false" json:"sakit_kepala"`
	PandanganKabur       bool `gorm:"not null;default:false" json:"pandangan_kabur"`
	NyeriUluHati         bool `gorm:"not null;default:false" json:"nyeri_ulu_hati"`
	GiziTerpenuhi        bool `gorm:"not null;default:false" json:"gizi_terpenuhi"`
	MasalahJiwa          bool `gorm:"not null;default:false" json:"masalah_jiwa"`
	JantungBerdebar      bool `gorm:"not null;default:false" json:"jantung_berdebar"`
	CairanJalanLahir     bool `gorm:"not null;default:false" json:"cairan_jalan_lahir"`
	NafasPendek          bool `gorm:"not null;default:false" json:"nafas_pendek"`
	KonsumsiVitaminA     bool `gorm:"not null;default:false" json:"konsumsi_vitamin_a"`
	PemeriksaanNifas     bool `gorm:"not null;default:false" json:"pemeriksaan_nifas"`

	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
	DeletedAt            gorm.DeletedAt `json:"-" gorm:"index"`
}

func (PemantauanIbuNifas) TableName() string {
	return "pemantauan_ibu_nifas"
}