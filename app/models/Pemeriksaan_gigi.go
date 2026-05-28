package models

import (
	"time"

	"gorm.io/gorm"
)

type PeriksaGigi struct {
	ID            int32     `json:"id" gorm:"primaryKey;autoIncrement"`
	AnakID        int32     `json:"anak_id" gorm:"not null;index"`
	Anak          *Anak     `json:"anak,omitempty" gorm:"foreignKey:AnakID"`
	Bulanke       int       `json:"bulan" gorm:"not null"`
	Tanggal       time.Time `json:"tanggal" gorm:"not null"`
	Jumlahgigi    int       `json:"jumlah_gigi" gorm:"not null;default:0"`
	GigiBerlubang int       `json:"gigi_berlubang" gorm:"not null"`
	// StatusPlak          string         `json:"status_plak" gorm:"type:enum('Bersih','Kotor');not null"`                     // bersih atau kotor
	// ResikoGigiBerlubang string         `json:"resiko_gigi_berlubang" gorm:"type:enum('Tinggi','Sedang','Rendah');not null"` // misal "rendah", "sedang", "tinggi"
	StatusPlak          string         `json:"status_plak" gorm:"type:varchar(10);not null"`
	ResikoGigiBerlubang string         `json:"resiko_gigi_berlubang" gorm:"type:varchar(10);not null"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `json:"-" gorm:"index"`
}

func (PeriksaGigi) TableName() string {
	return "periksa_gigi"
}

// func (P *PeriksaGigi) BeforeCreate(tx *gorm.DB) error {
// 	if P.ID == "" {
// 		P.ID = uuid.New().String()
// 	}
// 	return nil
// }
