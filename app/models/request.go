package models

import "time"

type CreateAnakRequest struct {
	KehamilanID   int32    `json:"kehamilan_id" validate:"required"`
	PendudukID    int32    `json:"penduduk_id"`
	Nama          string   `json:"nama"`
	JenisKelamin  string   `json:"jenis_kelamin"`
	TanggalLahir  string   `json:"tanggal_lahir"`
	BeratLahirKg  *float64 `json:"berat_lahir_kg,omitempty"`
	TinggiLahirCm *float64 `json:"tinggi_lahir_cm,omitempty"`
	AnakKe        int32    `json:"anak_ke"`
	LingkarKepalaCm *float64 `json:"lingkar_kepala_cm,omitempty"`
	NamaIbu       string   `json:"nama_ibu"`
	NamaAyah      string   `json:"nama_ayah"`
	IbuID         int32    `json:"ibu_id"`
}

// CreateAnakDenganPendudukRequest: create anak + bikin penduduk baru sekaligus
type CreateAnakDenganPendudukRequest struct {
	// Data kehamilan dan ibu
	KehamilanID int32  `json:"kehamilan_id" validate:"required"`
	IbuID       int32  `json:"ibu_id" validate:"required"`
	NamaIbu     string `json:"nama_ibu"`
	NamaAyah    string `json:"nama_ayah"`
	
	// Data anak (measurement)
	AnakKe          int32    `json:"anak_ke"`
	BeratLahirKg    *float64 `json:"berat_lahir_kg,omitempty"`
	TinggiLahirCm   *float64 `json:"tinggi_lahir_cm,omitempty"`
	LingkarKepalaCm *float64 `json:"lingkar_kepala_cm,omitempty"`
	
	// Data penduduk baru (anak)
	Nama          string `json:"nama" validate:"required"`
	TanggalLahir  string `json:"tanggal_lahir" validate:"required,datetime=2006-01-02"`
	JenisKelamin  string `json:"jenis_kelamin" validate:"required"`
	TempatLahir   string `json:"tempat_lahir,omitempty"`
	GolonganDarah string `json:"golongan_darah,omitempty"`
}

// UpdateAnakRequest adalah body request untuk PUT /anak/:id.
type UpdateAnakRequest struct {
	BeratLahirKg    *float64 `json:"berat_lahir_kg,omitempty"`
	TinggiLahirCm   *float64 `json:"tinggi_lahir_cm,omitempty"`
	AnakKe          *int32   `json:"anak_ke,omitempty"`
	LingkarKepalaCm *float64 `json:"lingkar_kepala_cm,omitempty"`
	NamaIbu         *string  `json:"nama_ibu,omitempty"`
	NamaAyah        *string  `json:"nama_ayah,omitempty"`
	// Data Penduduk (anak) — diperbarui ke tabel kependudukan
	Nama         string `json:"nama,omitempty"`
	JenisKelamin string `json:"jenis_kelamin,omitempty"`
	TanggalLahir string `json:"tanggal_lahir,omitempty"`
}

type DetailPelayananRequest struct {
	ID               int32  `json:"id,omitempty"`
	KunjunganAnakID  int32  `json:"kunjungan_anak_id,omitempty"`
	JenisPelayananID int32  `json:"jenis_pelayanan_id" validate:"required"`
	Nilai            string `json:"nilai" validate:"required"`
	Keterangan       string `json:"keterangan"`
}

type CreatePelayananKesehatanAnakRequest struct {
	AnakID          int32                    `json:"anak_id" validate:"required"`
	Tanggal         string                   `json:"tanggal" validate:"required,datetime=2006-01-02"`
	PeriodeID       int32                    `json:"periode_id" validate:"required"`
	KategoriUmurID  int32                    `json:"kategori_umur_id" validate:"required"`
	Lokasi          string                   `json:"lokasi" validate:"required"`
	DetailPelayanan []DetailPelayananRequest `json:"detail_pelayanan" validate:"required,dive"`
}

type UpdatePelayananKesehatanAnakRequest struct {
	AnakID          int32                    `json:"anak_id"`
	PeriodeID       int32                    `json:"periode_id"`
	Tanggal         string                   `json:"tanggal"` // "YYYY-MM-DD"
	Lokasi          string                   `json:"lokasi"`
	KategoriUmurID  int32                    `json:"kategori_umur_id"`
	DetailPelayanan []DetailPelayananRequest `json:"detail_pelayanan"`
}

type CreatePelayananNeonatusRequest struct {
	AnakID            int32                    `json:"anak_id" validate:"required"`
	Tanggal           string                   `json:"tanggal" validate:"required,datetime=2006-01-02"`
	PeriodeID         int32                    `json:"periode_id" validate:"required"`
	KategoriUmurID    int32                    `json:"kategori_umur_id" validate:"required"`
	TenagaKesehatanID int32                    `json:"tenaga_kesehatan_id" validate:"required"`
	DetailPelayanan   []DetailPelayananRequest `json:"detail_pelayanan" validate:"required,dive"`
}

type UpdatePelayananNeonatusRequest struct {
	AnakID            int32                    `json:"anak_id"`
	PeriodeID         int32                    `json:"periode_id"`
	Tanggal           string                   `json:"tanggal"` // "YYYY-MM-DD"
	TenagaKesehatanID int32                    `json:"tenaga_kesehatan_id"`
	KategoriUmurID    int32                    `json:"kategori_umur_id"`
	DetailPelayanan   []DetailPelayananRequest `json:"detail_pelayanan"`
}

// untuk asi
type ASIRequest struct {
	FrekuensiMenyusui int32  `json:"frekuensi_menyusui"`
	PosisiMenyusui    string `json:"posisi_menyusui"`
	ASIPerah          string `json:"asi_perah"`
}
type MPASIRequest struct {
	DiberikanMPASI      bool     `json:"diberikan_mp_asi"`
	VariasiMPASI        []string `json:"variasi_mpasi"`
	JumlahMakanPerporsi string   `json:"jumlah_makan_perporsi"`
	FrekuensiMakan      string   `json:"frekuensi_makan_perhari"`
}
type CreatePelayananGiziRequest struct {
	AnakID            int32  `json:"anak_id"`
	Tanggal           string `json:"tanggal"` // "YYYY-MM-DD"
	TenagaKesehatanID int32  `json:"tenaga_kesehatan_id"`
	Bulanke           int    `json:"bulan_ke"`
	Lokasi            string `json:"lokasi"`

	ASI   *ASIRequest   `json:"asi"`
	MPASI *MPASIRequest `json:"mpasi"`
}

type UpdatePelayananGiziRequest struct {
	AnakID            int32  `json:"anak_id"`
	Tanggal           string `json:"tanggal"` // "YYYY-MM-DD"
	TenagaKesehatanID int32  `json:"tenaga_kesehatan_id"`
	Bulanke           int    `json:"bulan_ke"`
	Lokasi            string `json:"lokasi"`

	ASI   *ASIRequest   `json:"asi"`
	MPASI *MPASIRequest `json:"mpasi"`
}

type CreateKunjunganVitaminRequest struct {
	AnakID  int32  `json:"anak_id"`
	Tanggal string `json:"tanggal"`

	Detail []CreateDetailPelayananVitaminRequest `json:"detail"`
}
type CreateDetailPelayananVitaminRequest struct {
	JenisPelayananID int32  `json:"jenis_pelayanan_id"`
	Keterangan       string `json:"keterangan"`
}
type UpdateKunjunganVitaminRequest struct {
	ID      int32  `json:"id"`
	AnakID  int32  `json:"anak_id"`
	Tanggal string `json:"tanggal"`

	Detail []UpdateDetailPelayananVitaminRequest `json:"detail"`
}
type UpdateDetailPelayananVitaminRequest struct {
	ID               *int32 `json:"id,omitempty"` // kalau null → insert baru
	JenisPelayananID int32  `json:"jenis_pelayanan_id"`
	Keterangan       string `json:"keterangan"`
}

type CreateKunjunganImunisasiRequest struct {
	AnakID  int32  `json:"anak_id"`
	Tanggal string `json:"tanggal"`
	Bulanke int    `json:"bulan_ke"`

	Detail []CreateDetailPelayananImunisasiRequest `json:"detail"`
}
type CreateDetailPelayananImunisasiRequest struct {
	JenisPelayananID int32  `json:"jenis_pelayanan_id"`
	Keterangan       string `json:"keterangan"`
}

type UpdateKunjunganImunisasiRequest struct {
	ID      int32  `json:"id"`
	AnakID  int32  `json:"anak_id"`
	Tanggal string `json:"tanggal"`
	Bulanke int    `json:"bulan_ke"`

	Detail []UpdateDetailPelayananImunisasiRequest `json:"detail"`
}
type UpdateDetailPelayananImunisasiRequest struct {
	ID               *int32 `json:"id,omitempty"` // kalau null → insert baru
	JenisPelayananID int32  `json:"jenis_pelayanan_id"`
	Keterangan       string `json:"keterangan"`
}

type CreatePemeriksaanGigiRequest struct {
	AnakID              int32  `json:"anak_id"`
	Bulanke             int    `json:"bulan_ke"`
	Tanggal             string `json:"tanggal"`
	Jumlahgigi          int    `json:"jumlah_gigi"`
	GigiBerlubang       int    `json:"gigi_berlubang"`
	StatusPlak          string `json:"status_plak"`
	ResikoGigiBerlubang string `json:"resiko_gigi_berlubang"`
}

type UpdatePemeriksaanGigiRequest struct {
	ID                  int32  `json:"id"`
	AnakID              int32  `json:"anak_id"`
	Bulanke             int    `json:"bulan_ke"`
	Tanggal             string `json:"tanggal"`
	Jumlahgigi          int    `json:"jumlah_gigi"`
	GigiBerlubang       int    `json:"gigi_berlubang"`
	StatusPlak          string `json:"status_plak"`
	ResikoGigiBerlubang string `json:"resiko_gigi_berlubang"`
}
type CreatePemantauanPemeriksaanRequest struct {
	AnakID            int32  `json:"anak_id"`
	Bulanke           int    `json:"bulan_ke"`
	Tanggal           string `json:"tanggal"`
	TenagaKesehatanID int32  `json:"tenaga_kesehatan_id"`
	BBperU            string `json:"bb_u"`
	BBperTB           string `json:"bb_tb"`
	TBperU            string `json:"tb_u"`
	LKperU            string `json:"lk_u"`
	LILA              string `json:"lila"`
	KPSP              string `json:"kpsp"`
	TDD               string `json:"tdd"`
	TDL               string `json:"tdl"`
	KMPE              string `json:"kmpe"`
	MCHATRevised      string `json:"m_chat_revised"`
	ACTRS             string `json:"actrs"`
	HasilPKAT         string `json:"hasil_pkat"`
	Tindakan          string `json:"tindakan"`
	KunjunganUlang    string `json:"kunjungan_ulang"`
}

type UpdatePemantauanPemeriksaanRequest struct {
	ID                int32  `json:"id"`
	AnakID            int32  `json:"anak_id"`
	Bulanke           int    `json:"bulan_ke"`
	Tanggal           string `json:"tanggal"`
	TenagaKesehatanID int32  `json:"tenaga_kesehatan_id"`
	BBperU            string `json:"bb_u"`
	BBperTB           string `json:"bb_tb"`
	TBperU            string `json:"tb_u"`
	LKperU            string `json:"lk_u"`
	LILA              string `json:"lila"`
	KPSP              string `json:"kpsp"`
	TDD               string `json:"tdd"`
	TDL               string `json:"tdl"`
	KMPE              string `json:"kmpe"`
	MCHATRevised      string    `json:"m_chat_revised"`
	ACTRS             string    `json:"actrs"`
	HasilPKAT         string    `json:"hasil_pkat"`
	Tindakan          string    `json:"tindakan"`
	KunjunganUlang    time.Time `json:"kunjungan_ulang"`
}

type CreatePengukuranLilARequest struct {
	AnakID         int32   `json:"anak_id"`
	Bulanke        int     `json:"bulan_ke"`
	Tanggal        string  `json:"tanggal"`
	HasilLila      float64 `json:"hasil_lila"`
	KategoriRisiko string  `json:"kategori_risiko"`
}

type UpdatePengukuranLilARequest struct {
	ID             int32   `json:"id"`
	AnakID         int32   `json:"anak_id"`
	Bulanke        int     `json:"bulan_ke"`
	Tanggal        string  `json:"tanggal"`
	HasilLila      float64 `json:"hasil_lila"`
	KategoriRisiko string  `json:"kategori_risiko"`
}

type CreateCatatanPelayananRequest struct {
	AnakID            int32     `json:"anak_id"`
	TenagaKesehatanID int32     `json:"tenaga_kesehatan_id"`
	TanggalPeriksa    time.Time `json:"tanggal_periksa"` // "YYYY-MM-DD"
	TanggalKembali    time.Time `json:"tanggal_kembali"` // "YYYY-MM-DD"
	CatatanPelayanan  string    `json:"catatan_pelayanan"`
}

type UpdateCatatanPelayananRequest struct {
	ID                int32     `json:"id"`
	AnakID            int32     `json:"anak_id"`
	TenagaKesehatanID int32     `json:"tenaga_kesehatan_id"`
	TanggalPeriksa    time.Time `json:"tanggal_periksa"` // "YYYY-MM-DD"
	TanggalKembali    time.Time `json:"tanggal_kembali"` // "YYYY-MM-DD"
	CatatanPelayanan  string    `json:"catatan_pelayanan"`
}

type CreatePertumbuhanRequest struct {
	AnakID        uint    `json:"anak_id" binding:"required"`
	TglUkur       string  `json:"tgl_ukur" binding:"required"` // Format: YYYY-MM-DD
	BeratBadan    float64 `json:"berat_badan" binding:"required"`
	TinggiBadan   float64 `json:"tinggi_badan" binding:"required"`
	LingkarKepala float64 `json:"lingkar_kepala,omitempty"`
	HasilLila     float64 `json:"hasil_lila,omitempty"`
	CatatanNakes  string  `json:"catatan_nakes,omitempty"`
}

type UpdatePertumbuhanRequest struct {
	TglUkur       string  `json:"tgl_ukur"`
	BeratBadan    float64 `json:"berat_badan"`
	TinggiBadan   float64 `json:"tinggi_badan"`
	LingkarKepala float64 `json:"lingkar_kepala,omitempty"`
	HasilLila     float64 `json:"hasil_lila,omitempty"`
	CatatanNakes  string  `json:"catatan_nakes,omitempty"`
}
type CatatanPertumbuhanResponse struct {
	ID            uint          `json:"id"`
	AnakID        uint          `json:"anak_id"`
	Anak          *AnakResponse `json:"anak,omitempty"`
	TglUkur       string        `json:"tgl_ukur"`
	UsiaUkurBulan int           `json:"usia_ukur_bulan"`
	UsiaUkurYM    string        `json:"usia_ukur_tahun_bulan,omitempty"` // Format: "1:4" = 1 tahun 4 bulan
	KurvaKMS      string        `json:"kurva_kms,omitempty"`             // "KMS 0-2 Tahun" atau "KMS 2-5 Tahun"
	BeratBadan    float64       `json:"berat_badan"`
	TinggiBadan   float64       `json:"tinggi_badan"`
	LingkarKepala float64       `json:"lingkar_kepala"`
	HasilLila     float64       `json:"hasil_lila"`
	IMT           float64       `json:"imt"`
	StatusBBU     string        `json:"status_bb_u"`
	StatusTBU     string        `json:"status_tb_u"`
	StatusIMTU    string        `json:"status_imt_u"`
	StatusBBTB    string        `json:"status_bb_tb"`
	StatusLKU     string        `json:"status_lk_u"`
	StatusKMSNaik string        `json:"status_kms_naik,omitempty"`  // Naik (N), Tidak Naik (T), Data Awal
	StatusKMSBGM  string        `json:"status_kms_bgm,omitempty"`   // Di bawah garis merah / Tidak
	KBMMinGram    int           `json:"kbm_min_gram,omitempty"`     // Kenaikan BB minimum (gram)
	KenaikanGram  float64       `json:"kenaikan_bb_gram,omitempty"` // Selisih dari penimbangan sebelumnya (gram)
	StatusKMSInfo string        `json:"status_kms_info,omitempty"`  // Ringkasan interpretasi KMS
	CatatanNakes  string        `json:"catatan_nakes"`
}

type CreateMasterStandarRequest struct {
	Parameter     string  `json:"parameter" binding:"required"` // bb_u, tb_u, imt_u, bb_tb, lk_u
	JenisKelamin  string  `json:"jenis_kelamin" binding:"required"`
	UsiaReferensi string  `json:"usia_referensi,omitempty"` // Khusus lk_u: bisa "1:4" atau "18"
	NilaiSumbuX   float64 `json:"nilai_sumbu_x" binding:"required"`
	SD3Neg        float64 `json:"sd_3_neg"`
	SD2Neg        float64 `json:"sd_2_neg" binding:"required"`
	SD1Neg        float64 `json:"sd_1_neg" binding:"required"`
	Median        float64 `json:"median" binding:"required"`
	SD1Pos        float64 `json:"sd_1_pos" binding:"required"`
	SD2Pos        float64 `json:"sd_2_pos" binding:"required"`
	SD3Pos        float64 `json:"sd_3_pos"`
}

type UpdateMasterStandarRequest struct {
	NilaiSumbuX float64 `json:"nilai_sumbu_x"`
	SD3Neg      float64 `json:"sd_3_neg"`
	SD2Neg      float64 `json:"sd_2_neg"`
	SD1Neg      float64 `json:"sd_1_neg"`
	Median      float64 `json:"median"`
	SD1Pos      float64 `json:"sd_1_pos"`
	SD2Pos      float64 `json:"sd_2_pos"`
	SD3Pos      float64 `json:"sd_3_pos"`
}
type MasterStandarResponse struct {
	ID           uint    `json:"id"`
	Parameter    string  `json:"parameter"`
	JenisKelamin string  `json:"jenis_kelamin"`
	NilaiSumbuX  float64 `json:"nilai_sumbu_x"`
	UsiaBulan    int     `json:"usia_bulan,omitempty"`
	UsiaYM       string  `json:"usia_tahun_bulan,omitempty"` // Format: "1:4" = 1 tahun 4 bulan
	SD3Neg       float64 `json:"sd_3_neg"`
	SD2Neg       float64 `json:"sd_2_neg"`
	SD1Neg       float64 `json:"sd_1_neg"`
	Median       float64 `json:"median"`
	SD1Pos       float64 `json:"sd_1_pos"`
	SD2Pos       float64 `json:"sd_2_pos"`
	SD3Pos       float64 `json:"sd_3_pos"`
}

// ANAK
type AnakRequest struct {
	ID              uint    `json:"id"`
	IbuID           *uint   `json:"ibu_id,omitempty"`
	KependudukanID  *uint   `json:"kependudukan_id,omitempty"`
	NoKartuKeluarga int64   `json:"no_kk,omitempty"`
	NamaAnak        string  `json:"nama_anak,omitempty"`
	JenisKelamin    string  `json:"jenis_kelamin,omitempty"`
	TanggalLahir    string  `json:"tanggal_lahir,omitempty"`
	BeratLahir      float64 `json:"berat_lahir"`
	TinggiLahir     float64 `json:"tinggi_lahir"`
}

type GenderType string

const (
	GenderMale   GenderType = "M"
	GenderFemale GenderType = "F"
)

type RekapDusun struct {
	Kecamatan string `json:"kecamatan"`
	Desa      string `json:"desa"`
	Dusun     string `json:"dusun"`
	Total     int64  `json:"total"`
}

