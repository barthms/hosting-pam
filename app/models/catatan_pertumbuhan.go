package models

import (
	"time"

	"gorm.io/gorm"
)

// CatatanPertumbuhan menyimpan data pengukuran anak
type CatatanPertumbuhan struct {
	ID            int32     `gorm:"primaryKey;column:id" db:"id" json:"id"`
	AnakID        int32     `gorm:"column:anak_id;type:int;not null;index" db:"anak_id" json:"anak_id"`
	TglUkur       time.Time `gorm:"column:tgl_ukur;type:date;not null;index" db:"tgl_ukur" json:"tgl_ukur"`
	BeratBadan    float64   `gorm:"column:berat_badan;type:decimal(5,2)" db:"berat_badan" json:"berat_badan"`
	TinggiBadan   float64   `gorm:"column:tinggi_badan;type:decimal(5,2)" db:"tinggi_badan" json:"tinggi_badan"`
	LingkarKepala float64   `gorm:"column:lingkar_kepala;type:decimal(5,2)" db:"lingkar_kepala" json:"lingkar_kepala,omitempty"`
	HasilLila     float64   `gorm:"column:hasil_lila;type:decimal(5,2)" db:"hasil_lila" json:"hasil_lila,omitempty"`
	IMT           float64   `gorm:"column:imt;type:decimal(5,2)" db:"imt" json:"imt,omitempty"`

	// Status gizi (hasil perhitungan)
	StatusBBU  string `gorm:"column:status_bb_u;type:varchar" db:"status_bb_u" json:"status_bb_u,omitempty"`
	StatusTBU  string `gorm:"column:status_tb_u;type:varchar" db:"status_tb_u" json:"status_tb_u,omitempty"`
	StatusIMTU string `gorm:"column:status_imt_u;type:varchar" db:"status_imt_u" json:"status_imt_u,omitempty"`
	StatusBBTB string `gorm:"column:status_bb_tb;type:varchar" db:"status_bb_tb" json:"status_bb_tb,omitempty"`
	StatusLKU  string `gorm:"column:status_lk_u;type:varchar" db:"status_lk_u" json:"status_lk_u,omitempty"`

	// Z-Score (hasil perhitungan)
	ZScoreBBU  float64 `gorm:"column:z_score_bb_u;type:decimal(5,2)" db:"z_score_bb_u" json:"z_score_bb_u,omitempty"`
	ZScoreTBU  float64 `gorm:"column:z_score_tb_u;type:decimal(5,2)" db:"z_score_tb_u" json:"z_score_tb_u,omitempty"`
	ZScoreIMTU float64 `gorm:"column:z_score_imt_u;type:decimal(5,2)" db:"z_score_imt_u" json:"z_score_imt_u,omitempty"`
	ZScoreBBTB float64 `gorm:"column:z_score_bb_tb;type:decimal(5,2)" db:"z_score_bb_tb" json:"z_score_bb_tb,omitempty"`
	ZScoreLKU  float64 `gorm:"column:z_score_lk_u;type:decimal(5,2)" db:"z_score_lk_u" json:"z_score_lk_u,omitempty"`

	// simpan di DB
	UsiaUkurBulan int    `gorm:"column:usia_ukur_bulan;type:int;not null;index" db:"usia_ukur_bulan" json:"usia_ukur_bulan"`
	CatatanNakes  string `gorm:"column:catatan_nakes;type:text" db:"catatan_nakes" json:"catatan_nakes,omitempty"`
	Anak          *Anak  `gorm:"foreignKey:AnakID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"anak,omitempty"`

	CreatedAt time.Time      `gorm:"column:created_at" db:"created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" db:"updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}

func (CatatanPertumbuhan) TableName() string {
	return "catatan_pertumbuhan"
}

// Hitung usia dalam bulan pada saat pengukuran
func (c CatatanPertumbuhan) HitungUsiaBulan(tanggalLahir time.Time) int {
	if tanggalLahir.After(c.TglUkur) {
		return 0
	}
	years := c.TglUkur.Year() - tanggalLahir.Year()
	months := int(c.TglUkur.Month() - tanggalLahir.Month())
	if months < 0 {
		years--
		months += 12
	}
	return years*12 + months
}

// Hitung IMT
func (c *CatatanPertumbuhan) HitungIMT() float64 {
	if c.TinggiBadan <= 0 {
		return 0
	}
	tinggiMeter := c.TinggiBadan / 100
	return c.BeratBadan / (tinggiMeter * tinggiMeter)
}
