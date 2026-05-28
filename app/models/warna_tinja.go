package models

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

const (
	WarnaTinjaPeriode2Minggu  = "2_minggu"
	WarnaTinjaPeriode1Bulan   = "1_bulan"
	WarnaTinjaPeriode2_4Bulan = "2_4_bulan"
)

var warnaTinjaPeriodeLabel = map[string]string{
	WarnaTinjaPeriode2Minggu:  "2 Minggu",
	WarnaTinjaPeriode1Bulan:   "1 Bulan",
	WarnaTinjaPeriode2_4Bulan: "2 - 4 Bulan",
}

type WarnaTinjaAnak struct {
	ID           uint       `gorm:"primaryKey;column:id" json:"id"`
	AnakID       uint       `gorm:"column:anak_id;not null;uniqueIndex:idx_warna_tinja_anak_periode" json:"anak_id"`
	PeriodeKey   string     `gorm:"column:periode_key;type:varchar(20);not null;uniqueIndex:idx_warna_tinja_anak_periode" json:"periode_key"`
	PeriodeLabel string     `gorm:"column:periode_label;type:varchar(30);not null" json:"periode_label"`
	TanggalCatat *time.Time `gorm:"column:tanggal_catat;type:date;not null" json:"tanggal_catat"`
	NomorWarna   int        `gorm:"column:nomor_warna;not null" json:"nomor_warna"`

	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`

	Anak *Anak `gorm:"foreignKey:AnakID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"anak,omitempty"`
}

func (WarnaTinjaAnak) TableName() string {
	return "warna_tinja_anak"
}

type WarnaTinjaSaveRequest struct {
	AnakID       uint   `json:"anak_id"`
	PeriodeKey   string `json:"periode_key"`
	TanggalCatat string `json:"tanggal_catat"`
	NomorWarna   int    `json:"nomor_warna"`
}

func (r *WarnaTinjaSaveRequest) Validate() error {
	if r.AnakID <= 0 {
		return errors.New("anak_id wajib lebih dari 0")
	}

	r.PeriodeKey = strings.TrimSpace(strings.ToLower(r.PeriodeKey))
	if _, ok := warnaTinjaPeriodeLabel[r.PeriodeKey]; !ok {
		return errors.New("periode_key tidak valid (gunakan: 2_minggu, 1_bulan, 2_4_bulan)")
	}

	if _, err := time.Parse("2006-01-02", r.TanggalCatat); err != nil {
		return errors.New("format tanggal_catat harus YYYY-MM-DD")
	}

	if r.NomorWarna < 1 || r.NomorWarna > 7 {
		return errors.New("nomor_warna harus antara 1 sampai 7")
	}

	return nil
}

func WarnaTinjaLabelForPeriode(periodeKey string) string {
	if label, ok := warnaTinjaPeriodeLabel[strings.TrimSpace(strings.ToLower(periodeKey))]; ok {
		return label
	}
	return ""
}
