package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// RentangUsia
type RentangUsia struct {
	ID          uint      `gorm:"primaryKey;column:id" db:"id" json:"id"`
	NamaRentang string    `gorm:"column:nama_rentang;type:varchar(50);not null" db:"nama_rentang" json:"nama_rentang"`
	SatuanWaktu string    `gorm:"column:satuan_waktu;type:varchar(20);not null" db:"satuan_waktu" json:"satuan_waktu"`
	MaxPeriode  int       `gorm:"column:max_periode;not null" db:"max_periode" json:"max_periode"`
	CreatedAt   time.Time `gorm:"column:created_at" db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" db:"updated_at" json:"updated_at"`
}

func (RentangUsia) TableName() string {
	return "rentang_usia"
}

// KategoriTandaSakit
type KategoriTandaSakit struct {
	ID            uint           `gorm:"primaryKey;column:id" db:"id" json:"id"`
	RentangUsiaID uint           `gorm:"column:rentang_usia_id;not null;index" db:"rentang_usia_id" json:"rentang_usia_id"`
	Gejala        string         `gorm:"column:gejala;type:text;not null" db:"gejala" json:"gejala"`
	Deskripsi     string         `gorm:"column:deskripsi;type:text" db:"deskripsi" json:"deskripsi"`
	IsActive      bool           `gorm:"column:is_active;default:true" db:"is_active" json:"is_active"`
	CreatedAt     time.Time      `gorm:"column:created_at" db:"created_at" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" db:"updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`

	RentangUsia *RentangUsia `gorm:"foreignKey:RentangUsiaID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"rentang_usia,omitempty"`
}

func (KategoriTandaSakit) TableName() string {
	return "kategori_tanda_sakit"
}

// LembarPemantauan (Header)
type LembarPemantauan struct {
	ID            uint `gorm:"primaryKey;column:id" db:"id" json:"id"`
	AnakID        uint `gorm:"column:anak_id;not null;uniqueIndex:idx_anak_periode" db:"anak_id" json:"anak_id"`
	RentangUsiaID uint `gorm:"column:rentang_usia_id;not null;uniqueIndex:idx_anak_periode" db:"rentang_usia_id" json:"rentang_usia_id"`
	PeriodeWaktu  int  `gorm:"column:periode_waktu;not null;uniqueIndex:idx_anak_periode" db:"periode_waktu" json:"periode_waktu"`

	TanggalPeriksa *time.Time `gorm:"column:tanggal_periksa;type:date;not null" db:"tanggal_periksa" json:"tanggal_periksa"`
	NamaPemeriksa  string     `gorm:"column:nama_pemeriksa;type:varchar(100)" db:"nama_pemeriksa" json:"nama_pemeriksa"`
	Status         string     `gorm:"column:status;type:varchar(50);not null" db:"status" json:"status"`

	CreatedAt time.Time      `gorm:"column:created_at" db:"created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" db:"updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`

	Anak         *Anak              `gorm:"foreignKey:AnakID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"anak,omitempty"`
	RentangUsia  *RentangUsia       `gorm:"foreignKey:RentangUsiaID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"rentang_usia,omitempty"`
	DetailGejala []DetailPemantauan `gorm:"foreignKey:LembarPemantauanID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"detail_gejala,omitempty"`
}

func (LembarPemantauan) TableName() string {
	return "lembar_pemantauan"
}

// DetailPemantauan (Detail)
type DetailPemantauan struct {
	ID                   uint           `gorm:"primaryKey;column:id" db:"id" json:"id"`
	LembarPemantauanID   uint           `gorm:"column:lembar_pemantauan_id;not null;index" db:"lembar_pemantauan_id" json:"lembar_pemantauan_id"`
	KategoriTandaSakitID uint           `gorm:"column:kategori_tanda_sakit_id;not null;index" db:"kategori_tanda_sakit_id" json:"kategori_tanda_sakit_id"`
	IsTerjadi            bool           `gorm:"column:is_terjadi;not null" db:"is_terjadi" json:"is_terjadi"`
	CreatedAt            time.Time      `gorm:"column:created_at" db:"created_at" json:"created_at"`
	UpdatedAt            time.Time      `gorm:"column:updated_at" db:"updated_at" json:"updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`

	LembarPemantauan   *LembarPemantauan   `gorm:"foreignKey:LembarPemantauanID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	KategoriTandaSakit *KategoriTandaSakit `gorm:"foreignKey:KategoriTandaSakitID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"kategori_tanda_sakit,omitempty"`
}

func (DetailPemantauan) TableName() string {
	return "detail_pemantauan"
}

// ==================== REQUEST MODELS ====================

type DetailPemantauanDetailRequest struct {
	KategoriTandaSakitID uint `json:"kategori_tanda_sakit_id" validate:"required"`
	IsTerjadi            bool `json:"is_terjadi"`
}

type LembarPemantauanRequest struct {
	AnakID         uint                            `json:"anak_id" validate:"required"`
	RentangUsiaID  uint                            `json:"rentang_usia_id" validate:"required"`
	PeriodeWaktu   int                             `json:"periode_waktu" validate:"required,min=1"`
	TanggalPeriksa string                          `json:"tanggal_periksa" validate:"required"`
	NamaPemeriksa  string                          `json:"nama_pemeriksa"`
	DetailGejala   []DetailPemantauanDetailRequest `json:"detail_gejala" validate:"required,dive"`
}

func (r *LembarPemantauanRequest) Validate() error {
	if r.AnakID <= 0 || r.RentangUsiaID <= 0 || r.PeriodeWaktu <= 0 {
		return errors.New("anak_id, rentang_usia_id, dan periode_waktu harus lebih dari 0")
	}
	if _, err := time.Parse("2006-01-02", r.TanggalPeriksa); err != nil {
		return errors.New("format tanggal_periksa harus YYYY-MM-DD")
	}
	if len(r.DetailGejala) == 0 {
		return errors.New("detail_gejala wajib diisi minimal 1")
	}
	return nil
}

type LembarPemantauanVerifikasiRequest struct {
	NamaPemeriksa string `json:"nama_pemeriksa" validate:"required"`
	Status        string `json:"status" validate:"required,oneof='Diterima' 'Ditolak'"`
}

func (r *LembarPemantauanVerifikasiRequest) Validate() error {
	if r.NamaPemeriksa == "" {
		return errors.New("nama_pemeriksa tidak boleh kosong saat melakukan verifikasi")
	}
	if r.Status != "Diterima" && r.Status != "Ditolak" {
		return errors.New("status hanya boleh diisi 'Diterima' atau 'Ditolak'")
	}
	return nil
}
