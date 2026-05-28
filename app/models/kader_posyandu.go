package models

import "time"

type KaderPosyandu struct {
	ID             uint      `gorm:"primaryKey;column:id" db:"id" json:"id"`
	UserID         uint      `gorm:"column:user_id;uniqueIndex;not null" db:"user_id" json:"user_id"`
	NoSk           string    `gorm:"column:no_sk;uniqueIndex" db:"no_sk" json:"no_sk,omitempty"`
	WilayahBinaan  string    `gorm:"column:wilayah_binaan;not null" db:"wilayah_binaan" json:"wilayah_binaan"`
	PosyanduNama   string    `gorm:"column:posyandu_nama" db:"posyandu_nama" json:"posyandu_nama,omitempty"`
	TahunBergabung int       `gorm:"column:tahun_bergabung" db:"tahun_bergabung" json:"tahun_bergabung,omitempty"`
	StatusAktif    bool      `gorm:"column:status_aktif;default:true" db:"status_aktif" json:"status_aktif"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
	Isdeleted      time.Time `gorm:"column:is_deleted" json:"is_deleted,omitempty"`

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}

func (KaderPosyandu) TableName() string {
	return "kader_posyandu"
}
