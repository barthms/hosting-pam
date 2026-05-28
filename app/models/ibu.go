package models

import (
	"time"

	"gorm.io/gorm"
)

type Ibu struct {
	IDIbu          int32         `gorm:"primaryKey;column:id;autoIncrement" json:"id_ibu"`
	IDKependudukan int32         `gorm:"column:penduduk_id;not null;index:idx_ibu_penduduk" json:"id_kependudukan"`
	Kependudukan   *Kependudukan `gorm:"foreignKey:IDKependudukan;references:IDKependudukan;constraint:-" json:"kependudukan"`

	IDSuami *int32        `gorm:"column:suami_id" json:"id_suami,omitempty"`
	Suami   *Kependudukan `gorm:"foreignKey:IDSuami;references:IDKependudukan;constraint:-" json:"suami,omitempty"`

	Gravida int32 `json:"gravida,omitempty"`
	Paritas int32 `json:"paritas,omitempty"`
	Abortus int32 `json:"abortus,omitempty"`

	// StatusKehamilan string        `json:"status_kehamilan"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	IsDeleted gorm.DeletedAt `gorm:"column:is_deleted" json:"is_deleted,omitempty"`
}

func (Ibu) TableName() string { return "ibu" }
