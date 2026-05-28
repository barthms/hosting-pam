package models

import (
	"time"

	"gorm.io/gorm"
)

type JenisPelayananKategori struct {
	ID               int32             `json:"id" gorm:"primaryKey;autoIncrement"`
	JenisPelayananID int32             `json:"jenis_pelayanan_id" gorm:"not null;index"`
	JenisPelayanan   *JenisPelayanan   `json:"jenis_pelayanan,omitempty" gorm:"foreignKey:JenisPelayananID"`
	KategoriUmurID   int32             `json:"kategori_umur_id" gorm:"not null;index"`
	KategoriUmur     *KategoriUmur     `json:"kategori_umur,omitempty" gorm:"foreignKey:KategoriUmurID"`
	PeriodeID        int32             `json:"periode_id" gorm:"index"`
	Periode          *PeriodeKunjungan `json:"periode,omitempty" gorm:"foreignKey:PeriodeID"`
	Urutan           int               `json:"urutan" gorm:"not null"`
	IsActive         bool              `json:"is_active" gorm:"default:true"`
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
	DeletedAt        gorm.DeletedAt    `json:"-" gorm:"index"`
}

func (JenisPelayananKategori) TableName() string {
	return "jenis_pelayanan_kategori"
}

// func (k *JenisPelayananKategori) BeforeCreate(tx *gorm.DB) error {
// 	if k.ID == "" {
// 		k.ID = uuid.New().String()
// 	}
// 	return nil
// }
