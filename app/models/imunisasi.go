package models

import "time"

type ImunisasiRequest struct {
	AnakID       int        `json:"anak_id"`
	ImunisasiID  int        `json:"imunisasi_id"`
	TglRencana   *time.Time `json:"tgl_rencana"`
	TglPemberian *time.Time `json:"tgl_pemberian"`
	Status       string     `json:"status"` // contoh: "Belum", "Sudah"
	Lokasi       string     `json:"lokasi"`
	Petugas      string     `json:"petugas"`
}

type Imunisasi struct {
	ID           uint       `gorm:"primaryKey;column:id" db:"id" json:"id"`
	AnakID       int        `gorm:"column:anak_id" db:"anak_id" json:"anak_id"`
	ImunisasiID  int        `gorm:"column:imunisasi_id" db:"imunisasi_id" json:"imunisasi_id"`
	TglRencana   *time.Time `gorm:"column:tgl_rencana" db:"tgl_rencana" json:"tgl_rencana,omitempty"`
	TglPemberian *time.Time `gorm:"column:tgl_pemberian" db:"tgl_pemberian" json:"tgl_pemberian,omitempty"`
	Status       string     `gorm:"column:status" db:"status" json:"status"`
	Lokasi       string     `gorm:"column:lokasi" db:"lokasi" json:"lokasi,omitempty"`
	Petugas      string     `gorm:"column:petugas" db:"petugas" json:"petugas,omitempty"`
	CreatedAt    time.Time  `gorm:"column:created_at" db:"created_at" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at" db:"updated_at" json:"updated_at"`
	Isdeleted    time.Time  `gorm:"column:is_deleted" db:"is_deleted" json:"is_deleted,omitempty"`

	// Foreign key constraints
	Anak      *Anak            `gorm:"foreignKey:AnakID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"-"`
	Imunisasi *MasterImunisasi `gorm:"foreignKey:ImunisasiID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"-"`
}

func (Imunisasi) TableName() string {
	return "imunisasi"
}
