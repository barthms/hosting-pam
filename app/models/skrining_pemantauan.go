package models

import (
	"time"

	"gorm.io/gorm"
)

type SkriningPemantauan struct {
	ID                    int32                `gorm:"primaryKey" json:"id"`
	AnakID                int32                `gorm:"not null;index" json:"anak_id"`
	KategoriTandaBahayaID int32                `gorm:"not null;index" json:"kategori_tanda_bahaya_id"`
	Anak                  *Anak                `gorm:"foreignKey:AnakID;references:ID" json:"anak,omitempty"`
	KategoriTandaBahaya   *KategoriTandaBahaya `gorm:"foreignKey:KategoriTandaBahayaID;references:ID" json:"kategori_tanda_bahaya,omitempty"`
	PeriodeWaktu          int                  `gorm:"not null" json:"periode_waktu"`                                           // Angkanya: 5, 2, atau 3
	SatuanWaktu           string               `gorm:"not null" json:"satuan_waktu"` // Enum/String: "HARI", "MINGGU", "BULAN", "TAHUN"
	DitemukanGejala       bool                 `gorm:"not null" json:"ditemukan_gejala"`
	TglSkrining           string               `gorm:"type:date;not null" json:"tgl_skrining"`
	CreatedAt             time.Time            `json:"created_at"`
	UpdatedAt             time.Time            `json:"updated_at"`
	DeletedAt             gorm.DeletedAt       `json:"-" gorm:"index"`
}

func (SkriningPemantauan) TableName() string {
	return "skrining_pemantauan"
}
