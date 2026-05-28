package models

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    interface{} `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}
type BasicResponse struct {
	StatusCode int         `json:"status_code"`
	Message    interface{} `json:"message,omitempty"`
}

type ResponseWithPaginate struct {
	StatusCode int         `json:"status_code"`
	Message    interface{} `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	Page      int `json:"page,omitempty"`
	PageSize  int `json:"page_size,omitempty"`
	Total     int `json:"total,omitempty"`
	TotalPage int `json:"total_page,omitempty"`
}
type KehamilanSimple struct {
	ID  int32 `json:"id"`
	Ibu struct {
		NamaIbu string `json:"nama_ibu"`
	} `json:"ibu,omitempty"`
}

type AnakResponse struct {
	ID            int32    `json:"id"`
	KehamilanID   int32    `json:"kehamilan_id"`
	PendudukID    int32    `json:"penduduk_id"`
	BeratLahirKg  *float64 `json:"berat_lahir_kg"`
	TinggiLahirCm *float64 `json:"tinggi_lahir_cm"`

	AnakKe        int32    `json:"anak_ke,omitempty"`
	LingkarKepalaCm *float64 `json:"lingkar_kepala_cm,omitempty"`
	NamaIbu       string   `json:"nama_ibu,omitempty"`
	NamaAyah      string   `json:"nama_ayah,omitempty"`
	IbuID         int32    `json:"ibu_id,omitempty"`

	Nama          string `json:"nama,omitempty"`
	TanggalLahir  string `json:"tanggal_lahir,omitempty"`
	JenisKelamin  string `json:"jenis_kelamin,omitempty"`
	GolonganDarah string `json:"golongan_darah,omitempty"`

	// Usia dalam bulan dan representasi teksnya
	UsiaBulan int    `json:"usia_bulan,omitempty"`
	UsiaTeks  string `json:"usia_teks,omitempty"`

	// TAMBAHKAN INI
	Kehamilan *KehamilanSimple `json:"kehamilan,omitempty"`
}

type PendudukSimple struct {
	ID   int32  `json:"id"`
	Nama string `json:"nama"`
}
