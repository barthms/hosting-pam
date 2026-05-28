package models

import (
	"time"

	"gorm.io/gorm"
)

type DetailPelayananImunisasi struct {
	ID                   int32               `json:"id" gorm:"primaryKey;autoIncrement"`
	KunjunganImunisasiID int32               `json:"kunjungan_imunisasi_id" gorm:"not null;index"`
	KunjunganImunisasi   *KehadiranImunisasi `json:"kunjungan_imunisasi,omitempty" gorm:"foreignKey:KunjunganImunisasiID"`
	JenisPelayananID     int32               `json:"jenis_pelayanan_id" gorm:"not null;index"`
	JenisPelayanan       *JenisPelayanan     `json:"jenis_pelayanan,omitempty" gorm:"foreignKey:JenisPelayananID"`
	Keterangan           string              `json:"keterangan"`
	CreatedAt            time.Time           `json:"created_at"`
	UpdatedAt            time.Time           `json:"updated_at"`
	DeletedAt            gorm.DeletedAt      `json:"-" gorm:"index"`
}

func (DetailPelayananImunisasi) TableName() string {
	return "detail_pelayanan_imunisasi"
}
