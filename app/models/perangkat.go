package models

import (
	"time"
	"gorm.io/gorm" 
)
type Perangkat struct {
	ID        		uint      	`gorm:"column:id;primaryKey" json:"id"`
	PenggunaID   	uint       `json:"id_pengguna" gorm:"column:id_pengguna;not null;index;constraint:OnDelete:CASCADE"`
	Pengguna     	*User     	`json:"pengguna,omitempty" gorm:"foreignKey:PenggunaID;constraint:OnDelete:CASCADE"`
	FcmToken      	string    	`gorm:"column:fcm_token;type:text;not null" json:"fcm_token"`
	CreatedAt 		time.Time 	`gorm:"column:created_at" json:"created_at"`
	UpdatedAt 		time.Time 	`gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at"`
}

func (Perangkat) TableName() string {
	return "perangkat"
}

type TokenRequest struct {
    PenggunaID   	uint   `json:"id_pengguna"`
    FcmToken 		string `json:"fcm_token"`
}

// Perangkat
type CreatePerangkatRequest struct {
	PenggunaID uint `json:"id_pengguna"`
	FcmToken   string `json:"fcm_token"`
}

type UpdatePerangkatRequest struct {
	PenggunaID uint `json:"id_pengguna"`
	FcmToken   string `json:"fcm_token"`
}
