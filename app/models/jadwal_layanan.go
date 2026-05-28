package models

import (
	"time"

	"gorm.io/gorm"
)

// JadwalLayanan: minimal model untuk jadwal layanan/kunjungan posyandu
type JadwalLayanan struct {
	ID           int32          `json:"id" gorm:"primaryKey;autoIncrement"`
	PosyanduID   *int32         `json:"posyandu_id" gorm:"index"`
	Posyandu     *Posyandu      `json:"posyandu,omitempty" gorm:"foreignKey:PosyanduID"` // ← tambah ini
	Layanan      string         `json:"layanan" gorm:"type:varchar(255);not null"`
	Tanggal      *time.Time     `json:"tanggal"`
	WaktuMulai   *string        `json:"waktu_mulai" gorm:"column:waktu_mulai;type:time"`
	WaktuSelesai *string        `json:"waktu_selesai" gorm:"column:waktu_selesai;type:time"`
	Keterangan   string         `json:"keterangan" gorm:"type:text"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

func (JadwalLayanan) TableName() string {
	return "jadwal_layanan"
}
