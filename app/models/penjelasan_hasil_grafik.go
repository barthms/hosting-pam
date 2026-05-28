package models

import "time"

type PenjelasanHasilGrafik struct {
	IDPenjelasan            int32      `gorm:"primaryKey" json:"id_penjelasan"`
	KehamilanID             int32      `gorm:"not null;index" json:"kehamilan_id"`
	Kehamilan               *Kehamilan `gorm:"foreignKey:KehamilanID;references:ID" json:"kehamilan,omitempty"`
	CatatanPenjelasanGrafik string     `json:"catatan_penjelasan_grafik"`
	CreatedAt               time.Time  `json:"created_at"`
}

func (PenjelasanHasilGrafik) TableName() string {
	return "penjelasan_hasil_grafik"
}
