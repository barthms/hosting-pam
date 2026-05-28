package models

import (
	"time"

	"gorm.io/gorm"
)

type KategoriCapaian struct {
	ID          uint   `gorm:"primaryKey;column:id" db:"id" json:"id"`
	RentangUsia string `gorm:"column:rentang_usia" db:"rentang_usia" json:"rentang_usia,omitempty"`
	// TipeLembarCapaian  string         `gorm:"column:tipe_lembar_capaian" db:"tipe_lembar_capaian" json:"tipe_lembar_capaian,omitempty"`
	PertanyaaanCeklist string         `gorm:"column:pertanyaan_ceklist" db:"pertanyaan_ceklist" json:"pertanyaan_ceklist,omitempty"`
	Aspek              string         `gorm:"column:aspek" db:"aspek" json:"aspek,omitempty"`
	CreatedAt          time.Time      `gorm:"column:created_at" db:"created_at" json:"created_at"`
	UpdatedAt          time.Time      `gorm:"column:updated_at" db:"updated_at" json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}

func (KategoriCapaian) TableName() string {
	return "kategori_capaian"
}

type Perawatan struct {
	ID                uint           `gorm:"primaryKey;column:id" db:"id" json:"id"`
	AnakID            int32          `gorm:"column:anak_id" db:"anak_id" json:"anak_id"`
	KategoriCapaianID uint           `gorm:"column:kategori_capaian_id" db:"kategori_capaian_id" json:"kategori_capaian_id"`
	Jawaban           *bool          `gorm:"column:jawaban" db:"jawaban" json:"jawaban,omitempty"`
	TanggalPeriksa    *time.Time     `gorm:"column:tanggal_periksa" db:"tanggal_periksa" json:"tanggal_periksa,omitempty"`
	CreatedAt         time.Time      `gorm:"column:created_at" db:"created_at" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"column:updated_at" db:"updated_at" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`

	// Foreign key constraints
	Anak            *Anak            `gorm:"foreignKey:AnakID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"-"`
	KategoriCapaian *KategoriCapaian `gorm:"foreignKey:KategoriCapaianID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"-"`
}

func (Perawatan) TableName() string {
	return "perawatan"
}

// ─────────────────────────────────────────────────────────
// PERAWATAN REQUEST
// ─────────────────────────────────────────────────────────

// CreatePerawatanRequest is body request for POST /perawatan
type CreatePerawatanRequest struct {
	AnakID            int32  `json:"anak_id" validate:"required"`
	KategoriCapaianID uint   `json:"kategori_capaian_id" validate:"required"`
	Jawaban           *bool  `json:"jawaban,omitempty"`
	TanggalPeriksa    string `json:"tanggal_periksa,omitempty"`
}

// UpdatePerawatanRequest is body request for PUT /perawatan/:id
type UpdatePerawatanRequest struct {
	Jawaban        *bool  `json:"jawaban,omitempty"`
	TanggalPeriksa string `json:"tanggal_periksa,omitempty"`
}
