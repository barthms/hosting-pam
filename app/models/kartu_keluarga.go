package models

import "time"

type KartuKeluarga struct {
	ID            int64      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NoKK          string     `gorm:"column:no_kk;type:varchar(20);not null;uniqueIndex" json:"no_kk"`
	CreatedAt     time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	TanggalTerbit *time.Time `gorm:"column:tanggal_terbit" json:"tanggal_terbit,omitempty"`
}

func (KartuKeluarga) TableName() string { return "kartu_keluarga" }
