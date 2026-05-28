package models

import (
	"time"

	"gorm.io/gorm"
)

type DosisVaksin struct {
	ID          uint           `gorm:"column:id;primaryKey" json:"id"`
	NamaDosis   string         `gorm:"column:nama_dosis;type:text" json:"nama_dosis"`
	VaksinID    uint         `gorm:"column:id_vaksin;" json:"vaksin_id"`
	Vaksin      *Vaksin      `json:"vaksin,omitempty" gorm:"foreignKey:VaksinID;constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at"`
}

func (DosisVaksin) TableName() string {
	return "dosis_vaksin"
}
