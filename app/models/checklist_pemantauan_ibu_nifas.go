package models

import (
	"time"

	"gorm.io/gorm"
)

type ChecklistPemantauanIbuNifas struct {
	ID          int32      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	KehamilanID int32     `gorm:"column:kehamilan_id;not null;uniqueIndex:idx_nifas_unique" json:"kehamilan_id"`
	HariNifas 	int32 `gorm:"column:hari_nifas;not null;uniqueIndex:idx_nifas_unique" json:"hari_nifas"`
	Kehamilan   *Kehamilan `gorm:"foreignKey:KehamilanID;references:ID;constraint:OnDelete:CASCADE" json:"kehamilan,omitempty"`

	// NIFAS A
	PemeriksaanNifas     bool `gorm:"column:pemeriksaan_nifas;default:false" json:"pemeriksaan_nifas"`
	KonsumsiVitaminA     bool `gorm:"column:konsumsi_vitamin_a;default:false" json:"konsumsi_vitamin_a"`
	PemenuhanGizi        bool `gorm:"column:pemenuhan_gizi;default:false" json:"pemenuhan_gizi"`
	DemamLebih38         bool `gorm:"column:demam_lebih_38;default:false" json:"demam_lebih_38"`
	SakitKepala          bool `gorm:"column:sakit_kepala;default:false" json:"sakit_kepala"`
	PandanganKabur       bool `gorm:"column:pandangan_kabur;default:false" json:"pandangan_kabur"`
	NyeriUluHati         bool `gorm:"column:nyeri_ulu_hati;default:false" json:"nyeri_ulu_hati"`
	MasalahKesehatanJiwa bool `gorm:"column:masalah_kesehatan_jiwa;default:false" json:"masalah_kesehatan_jiwa"`

	// NIFAS B
	JantungBerdebar     bool `gorm:"column:jantung_berdebar;default:false" json:"jantung_berdebar"`
	CairanJalanLahir    bool `gorm:"column:cairan_jalan_lahir;default:false" json:"cairan_jalan_lahir"`
	NapasPendek         bool `gorm:"column:napas_pendek;default:false" json:"napas_pendek"`
	PayudaraBermasalah  bool `gorm:"column:payudara_bermasalah;default:false" json:"payudara_bermasalah"`
	GangguanBAK         bool `gorm:"column:gangguan_bak;default:false" json:"gangguan_bak"`
	KelaminBermasalah   bool `gorm:"column:kelamin_bermasalah;default:false" json:"kelamin_bermasalah"`
	DarahNifasBerbau    bool `gorm:"column:darah_nifas_berbau;default:false" json:"darah_nifas_berbau"`
	PendarahanBerat     bool `gorm:"column:pendarahan_berat;default:false" json:"pendarahan_berat"`
	Keputihan           bool `gorm:"column:keputihan;default:false" json:"keputihan"`

	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}

func (ChecklistPemantauanIbuNifas) TableName() string {
	return "checklist_pemantauan_ibu_nifas"
}