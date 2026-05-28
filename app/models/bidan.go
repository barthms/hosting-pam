package models

import "time"

type Bidan struct {
	ID         int32      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PendudukID int32      `gorm:"column:penduduk_id;not null;uniqueIndex" json:"penduduk_id"`
	NoSTR      string     `gorm:"column:no_str;type:varchar(100)" json:"no_str"`
	NoSIPB     string     `gorm:"column:no_sipb;type:varchar(100)" json:"no_sipb"`
	Status     string     `gorm:"column:status;type:varchar(20);not null;default:aktif" json:"status"`
	CreatedAt  time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`

	Penduduk Kependudukan `gorm:"foreignKey:PendudukID;references:IDKependudukan" json:"penduduk,omitempty"`
}

func (Bidan) TableName() string { return "bidan" }
