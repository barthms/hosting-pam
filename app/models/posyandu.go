package models

import "time"

type Posyandu struct {
	ID          int32      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	IDPuskesmas int32      `gorm:"column:id_puskesmas;not null" json:"id_puskesmas"`
	Nama        string     `gorm:"column:nama;type:varchar(255);not null" json:"nama"`
	Alamat      string     `gorm:"column:alamat;type:text" json:"alamat,omitempty"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
}

func (Posyandu) TableName() string { return "posyandu" }
