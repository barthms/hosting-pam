package models

import (
	"time"

	"gorm.io/gorm"
)

type AturanVaksinAnak struct {
	ID          uint           `gorm:"column:id;primaryKey" json:"id"`
	DosisVaksinID uint           `gorm:"column:id_dosis_vaksin;" json:"dosis_vaksin_id"`
	MinUsiaHari   uint         `gorm:"column:min_usia_hari;" json:"min_usia_hari"`
	MaxUsiaHari   uint         `gorm:"column:max_usia_hari;" json:"max_usia_hari"`
	MinIntervalHari   uint         `gorm:"column:min_interval_hari;" json:"min_interval_hari"`
	DosisSebelumnyaID   *uint         `gorm:"column:id_dosis_sebelumnya;" json:"dosis_sebelumnya_id,omitempty"`
	DosisVaksin   *DosisVaksin   `json:"dosis_vaksin,omitempty" gorm:"foreignKey:DosisVaksinID;constraint:OnDelete:CASCADE"`
	DosisSebelumnya *DosisVaksin   `json:"dosis_sebelumnya,omitempty" gorm:"foreignKey:DosisSebelumnyaID;constraint:OnDelete:SET NULL"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at"`
}

func (AturanVaksinAnak) TableName() string {
	return "aturan_vaksin_anak"
}
