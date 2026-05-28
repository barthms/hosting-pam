package models

import "time"

type MasterImunisasi struct {
	ID                   uint      `gorm:"primaryKey;column:id" db:"id" json:"id"`
	NamaVaksin           string    `gorm:"column:nama_vaksin" db:"nama_vaksin" json:"nama_vaksin"`
	UsiaRekomendasiBulan int       `gorm:"column:usia_rekomendasi_bulan" db:"usia_rekomendasi_bulan" json:"usia_rekomendasi_bulan"`
	DeskripsiManfaat     string    `gorm:"column:deskripsi_manfaat" db:"deskripsi_manfaat" json:"deskripsi_manfaat,omitempty"`
	CreatedAt            time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt            time.Time `gorm:"column:updated_at" json:"updated_at"`
	Isdeleted            time.Time `gorm:"column:is_deleted" json:"is_deleted,omitempty"`
}

func (MasterImunisasi) TableName() string {
	return "master_imunisasi"
}
