package models

import (
	"time"
	"gorm.io/gorm" 
)
type Vaksin struct {
	ID        		uint      	`gorm:"column:id;primaryKey" json:"id"`
	Name      		string    	`gorm:"column:nama;type:varchar(50);not null" json:"nama"`
	Deskripsi    	string 		`gorm:"column:deskripsi;type:text;not null" json:"deskripsi"`
	EfekSamping    	string 		`gorm:"column:efek_samping;type:text;not null" json:"efek_samping"`
	CreatedAt 		time.Time 	`gorm:"column:created_at" json:"created_at"`
	UpdatedAt 		time.Time 	`gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at"`
}

func (Vaksin) TableName() string {
	return "vaksin"
}