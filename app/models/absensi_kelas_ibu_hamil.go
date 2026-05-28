package models

import (
	"time"
)

type AbsensiKelasIbuHamil struct {
	ID          int32 `gorm:"primaryKey;autoIncrement" json:"id"`
	KehamilanID int32 `gorm:"not null;index;uniqueIndex:idx_absensi_unique" json:"kehamilan_id"`

	PertemuanKe int32      `gorm:"not null;uniqueIndex:idx_absensi_unique" json:"pertemuan_ke"`
	Tanggal     *time.Time `gorm:"type:date" json:"tanggal"`

	NamaKader    string     `gorm:"type:varchar(255)" json:"nama_kader"`
	TanggalParaf *time.Time `gorm:"type:date" json:"tanggal_paraf"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (AbsensiKelasIbuHamil) TableName() string {
	return "absensi_kelas_ibu_hamil"
}
