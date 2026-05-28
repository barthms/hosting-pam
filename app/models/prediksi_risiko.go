package models

type PrediksiRisikoRequest struct {
	UsiaIbu           float64 `json:"usia_ibu"`
	UsiaKehamilan     int     `json:"usia_kehamilan"`
	TrimesterNum      int     `json:"trimester_num"`
	Gravida           int     `json:"gravida"`
	Para              int     `json:"para"`
	Abortus           int     `json:"abortus"`
	IMT               float64 `json:"imt"`
	LiLA              float64 `json:"lila"`
	TinggiFundusUteri float64 `json:"tinggi_fundus_uteri"`
	TDSistolik        float64 `json:"td_sistolik"`
	TDDiastolik       float64 `json:"td_diastolik"`
	Hemoglobin        float64 `json:"hemoglobin"`
	KunjunganANCKe    int     `json:"kunjungan_anc_ke"`
	ImunisasiEnc      int     `json:"imunisasi_enc"`
	RiwayatEnc        int     `json:"riwayat_enc"`
	RiwayatBerat      int     `json:"riwayat_berat"`
	HIVRek            int     `json:"hiv_rek"`
	SifRek            int     `json:"sif_rek"`
	HepBRek           int     `json:"hepb_rek"`
}

type PrediksiRisikoResponse struct {
	Prediction    int       `json:"prediction"`
	Label         string    `json:"label"`
	Probabilities []float64 `json:"probabilities"`
	RiskScore     int       `json:"risk_score"`
}