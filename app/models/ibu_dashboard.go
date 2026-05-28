package models

import "time"

type IbuDashboardDTO struct {
	IDIbu       int32  `json:"id_ibu"`
	NamaLengkap string `json:"nama_lengkap"`
	Dusun       string `json:"dusun"`

	StatusKehamilan string `json:"status_kehamilan"`
	UsiaKehamilan   int32  `json:"usia_kehamilan"`

	TanggalPeriksa *time.Time `json:"tanggal_periksa"`

	Hpht          *time.Time `json:"hpht"`
	TempatPeriksa string     `json:"tempat_periksa"`

	Trimester   string `json:"trimester"`
	KunjunganKe int32  `json:"kunjungan_ke"`

	SkorRisiko   int32    `json:"skor_risiko"`
	StatusRisiko string   `json:"status_risiko"`
	KehamilanID  int32    `json:"kehamilan_id"`
	Sistole      int32    `json:"sistole"`
	Diastole     int32    `json:"diastole"`
	Hb           *float64 `json:"hb"`
}
