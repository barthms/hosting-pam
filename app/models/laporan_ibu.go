package models

import "time"

type LaporanIbu struct {
	NIK              string    `json:"nik"`
	NamaIbu          string    `json:"nama_ibu"`
	NamaSuami        string    `json:"nama_suami"`
	TanggalLahir     time.Time `json:"tanggal_lahir"`

	HPHT             time.Time `json:"hpht"`
	HPL              time.Time `json:"hpl"`

	UsiaKehamilan    int32     `json:"usia_kehamilan"`
	Trimester        string    `json:"trimester"`

	Gravida          int32     `json:"gravida"`
	Paritas          int32     `json:"paritas"`
	Abortus          int32     `json:"abortus"`

	BBAwal           float64   `json:"bb_awal"`
	TinggiBadan      float64   `json:"tinggi_badan"`
	IMT              float64   `json:"imt"`

	LILA             float64   `json:"lila"`

	TekananDarah     string    `json:"tekanan_darah"`
	Sistole          int32     `json:"sistole"`
	Diastole         int32     `json:"diastole"`

	TinggiFundus     float64   `json:"tinggi_fundus"`

	Hb               float64   `json:"hb"`
	GolonganDarah    string    `json:"golongan_darah"`

	StatusImunisasi  string    `json:"status_imunisasi"`

	TripelEliminasi  string    `json:"tripel_eliminasi"`

	KunjunganANC     int32     `json:"kunjungan_anc"`

	Tindakan         string    `json:"tindakan"`

	Kecamatan        string    `json:"kecamatan"`
	Desa             string    `json:"desa"`
}