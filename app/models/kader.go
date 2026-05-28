package models

import "time"

type Kader struct {
	ID         int32      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PendudukID int32      `gorm:"column:id_penduduk;not null;uniqueIndex" json:"penduduk_id"`
	PosyanduID *int64     `gorm:"column:id_posyandu" json:"posyandu_id,omitempty"`
	Status     string     `gorm:"column:status;type:varchar(20);not null;default:aktif" json:"status"`
	CreatedAt  time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`

	Penduduk Kependudukan `gorm:"foreignKey:PendudukID;references:IDKependudukan" json:"penduduk,omitempty"`
}

func (Kader) TableName() string { return "kader" }
