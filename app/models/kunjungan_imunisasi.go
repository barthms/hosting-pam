package models

import (
	"time"
)

type KunjunganImunisasi struct {
	ID               uint             `gorm:"primaryKey" json:"id"`
	StatusID         uint             `gorm:"column:status;"json:"status"`
	TanggalKunjungan *time.Time       `gorm:"column:tanggal_kunjungan;" json:"tanggal_kunjungan"`
	StatusKunjungan  *StatusKunjungan `json:"status_kunjungan,omitempty" gorm:"foreignKey:StatusID"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (KunjunganImunisasi) TableName() string {
	return "kunjungan_imunisasi"
}

type KunjunganImunisasiDetailResponse struct {
	KunjunganID      uint       `json:"kunjungan_id"`
	TanggalKunjungan *time.Time `json:"tanggal_kunjungan,omitempty"`
	StatusKunjungan  string     `json:"status_kunjungan"`

	NamaAnak     string     `json:"nama_anak"`
	TanggalLahir *time.Time `json:"tanggal_lahir,omitempty"`

	NamaIbu string `json:"nama_ibu"`
	NomorTeleponIbu string `json:"nomor_telepon_ibu"`
	NamaVaksin      string     `json:"nama_vaksin"`
	NamaDosis       string     `json:"nama_dosis"`
	JadwalImunisasi *time.Time `json:"jadwal_imunisasi,omitempty"`
}

type KunjunganImunisasiResponse struct {
	KunjunganID      uint       `json:"kunjungan_id"`
	TanggalKunjungan *time.Time `json:"tanggal_kunjungan,omitempty"`
	StatusKunjungan  string     `json:"status_kunjungan"`
	NamaAnak         string     `json:"nama_anak"`
}

type UpdateStatusKunjunganRequest struct {
	StatusKunjunganID uint `json:"status_kunjungan_id"`
}

type UpdateTanggalKunjunganRequest struct {
	TanggalKunjungan string `json:"tanggal_kunjungan"`
}