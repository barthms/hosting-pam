package models

import (
	"time"

	"gorm.io/gorm"
)

// MeasurementDataForPrediction - struktur data pengukuran untuk prediksi
type MeasurementDataForPrediction struct {
	AnakID            int32     `json:"anak_id" gorm:"column:anak_id"`
	BeratBadan        float64   `json:"berat_badan" gorm:"column:berat_badan"`          // kg
	TinggiBadan       float64   `json:"tinggi_badan" gorm:"column:tinggi_badan"`        // cm
	LingkarKepala     float64   `json:"lingkar_kepala" gorm:"column:lingkar_kepala"`    // cm
	HasilLila         float64   `json:"hasil_lila" gorm:"column:hasil_lila"`            // cm (Lingkar Lengan Atas)
	UsiaUkurBulan     int       `json:"usia_ukur_bulan" gorm:"column:usia_ukur_bulan"`  // bulan
	TglUkur           time.Time `json:"tgl_ukur" gorm:"column:tgl_ukur"`
	StatusTBU         string    `json:"status_tb_u" gorm:"column:status_tb_u"`          // Status TB/Umur
	ZScoreTBU         float64   `json:"z_score_tb_u" gorm:"column:z_score_tb_u"`        // Z-Score TB/Umur
}

// PrediksiStunting - hasil prediksi stunting
type PrediksiStunting struct {
	ID                uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AnakID            int32     `gorm:"not null;index" json:"anak_id"`
	Anak              *Anak     `json:"anak,omitempty" gorm:"foreignKey:AnakID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	
	// Data input pengukuran
	BeratBadan        float64   `gorm:"type:decimal(5,2);not null" json:"berat_badan"`
	TinggiBadan       float64   `gorm:"type:decimal(5,2);not null" json:"tinggi_badan"`
	LingkarKepala     float64   `gorm:"type:decimal(5,2);not null" json:"lingkar_kepala"`
	HasilLila         float64   `gorm:"type:decimal(5,2);not null" json:"hasil_lila"`
	UsiaUkurBulan     int       `gorm:"not null" json:"usia_ukur_bulan"`
	
	// Hasil prediksi
	RiskPercentage    float64   `gorm:"type:decimal(5,2);not null" json:"risk_percentage"`   // 0-100%
	Classification    string    `gorm:"type:varchar(20);not null" json:"classification"`     // STUNTING, AT_RISK, NORMAL
	Confidence        float64   `gorm:"type:decimal(5,2)" json:"confidence"`                 // 0-100%
	
	// Data tambahan dari WHO standards
	ZScoreTBU         float64   `gorm:"type:decimal(5,2)" json:"z_score_tb_u"`
	StatusTBU         string    `gorm:"type:varchar(20)" json:"status_tb_u"`
	
	// Rekomendasi untuk bidan
	Rekomendasi       string    `gorm:"type:text" json:"rekomendasi"`
	Catatan           string    `gorm:"type:text" json:"catatan"`
	
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`
}

func (PrediksiStunting) TableName() string {
	return "prediksi_stunting"
}

// PrediksiStuntingRequest - request dari frontend
type PrediksiStuntingRequest struct {
	AnakID        int32   `json:"anak_id" validate:"required"`
	BeratBadan    float64 `json:"berat_badan" validate:"required,min=0"`
	TinggiBadan   float64 `json:"tinggi_badan" validate:"required,min=0"`
	LingkarKepala float64 `json:"lingkar_kepala" validate:"required,min=0"`
	HasilLila     float64 `json:"hasil_lila" validate:"required,min=0"`
	UsiaUkurBulan int     `json:"usia_ukur_bulan" validate:"required,min=0,max=60"`
	JenisKelamin  string  `json:"jenis_kelamin" validate:"required,oneof=Laki-laki Perempuan"`
}

// PrediksiResponse - response dari Python service
// Field z_score_tb_u_estimated sesuai dengan output FastAPI ML service
type PrediksiResponse struct {
	StuntingRisk   float64 `json:"stunting_risk"`
	Classification string  `json:"classification"`
	Confidence     float64 `json:"confidence"`
	ZScoreTBU      float64 `json:"z_score_tb_u_estimated"` // nama field di ML service
	StatusTBU      string  `json:"status_tb_u"`
	Rekomendasi    string  `json:"rekomendasi"`
	Message        string  `json:"message"`
}
