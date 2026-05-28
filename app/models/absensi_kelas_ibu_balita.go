package models

import "time"

type AbsensiKelasIbuBalita struct {
	ID    int32 `gorm:"primaryKey;autoIncrement" json:"id"`
	IbuID int32 `gorm:"not null;index" json:"ibu_id"`

	PertemuanKe  int32      `gorm:"not null" json:"pertemuan_ke"`
	Tanggal      *time.Time `gorm:"type:date" json:"tanggal"`
	NamaKader    string     `gorm:"type:varchar(255)" json:"nama_kader"`
	TanggalParaf *time.Time `gorm:"type:date" json:"tanggal_paraf"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (AbsensiKelasIbuBalita) TableName() string {
	return "absensi_kelas_ibu_balita"
}
