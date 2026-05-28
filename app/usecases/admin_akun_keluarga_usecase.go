package usecases

import (
	"math"
	"strings"
	"time"

	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
	"monitoring-service/pkg/customerror"

	"gorm.io/gorm"
)

type AdminAnggotaKeluargaRequest struct {
	NIK                string `json:"nik"`
	NamaLengkap        string `json:"nama_lengkap"`
	JenisKelamin       string `json:"jenis_kelamin"`
	TanggalLahir       string `json:"tanggal_lahir"`
	TempatLahir        string `json:"tempat_lahir"`
	GolonganDarah      string `json:"golongan_darah"`
	Agama              string `json:"agama"`
	StatusPerkawinan   string `json:"status_perkawinan"`
	Pekerjaan          string `json:"pekerjaan"`
	PendidikanTerakhir string `json:"pendidikan_terakhir"`
	BacaHuruf          string `json:"baca_huruf"`
	KedudukanKeluarga  string `json:"kedudukan_keluarga"`
	Dusun              string `json:"dusun"`
	AsalPenduduk       string `json:"asal_penduduk"`
	TujuanPindah       string `json:"tujuan_pindah"`
	TempatMeninggal    string `json:"tempat_meninggal"`
	Keterangan         string `json:"keterangan"`
}

// AdminCreateKartuKeluargaRequest - Hanya membuat Kartu Keluarga + Anggota (Penduduk)
type AdminCreateKartuKeluargaRequest struct {
	NoKK            string                        `json:"no_kk"`
	TanggalTerbit   string                        `json:"tanggal_terbit"`
	AnggotaKeluarga []AdminAnggotaKeluargaRequest `json:"anggota_keluarga"`
}

type AdminListKartuKeluargaPagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type AdminListKartuKeluargaItem struct {
	KartuKeluargaID int64       `json:"kartu_keluarga_id"`
	NoKK            string      `json:"no_kk"`
	TanggalTerbit   *string     `json:"tanggal_terbit"`
	KepalaKeluarga  interface{} `json:"kepala_keluarga"`
	JumlahAnggota   int         `json:"jumlah_anggota"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

type AdminKepalaKeluarga struct {
	PendudukID  int32  `json:"penduduk_id"`
	NIK         string `json:"nik"`
	NamaLengkap string `json:"nama_lengkap"`
}

type AdminDetailKartuKeluargaAnggota struct {
	PendudukID         int32  `json:"penduduk_id"`
	NIK                string `json:"nik"`
	NamaLengkap        string `json:"nama_lengkap"`
	JenisKelamin       string `json:"jenis_kelamin"`
	TanggalLahir       string `json:"tanggal_lahir"`
	TempatLahir        string `json:"tempat_lahir"`
	GolonganDarah      string `json:"golongan_darah"`
	Agama              string `json:"agama"`
	StatusPerkawinan   string `json:"status_perkawinan"`
	Pekerjaan          string `json:"pekerjaan"`
	PendidikanTerakhir string `json:"pendidikan_terakhir"`
	BacaHuruf          string `json:"baca_huruf"`
	KedudukanKeluarga  string `json:"kedudukan_keluarga"`
	Dusun              string `json:"dusun"`
	AsalPenduduk       string `json:"asal_penduduk"`
	TujuanPindah       string `json:"tujuan_pindah"`
	TempatMeninggal    string `json:"tempat_meninggal"`
	Keterangan         string `json:"keterangan"`
}

type AdminDetailKartuKeluargaResponse struct {
	KartuKeluargaID int64                             `json:"kartu_keluarga_id"`
	NoKK            string                            `json:"no_kk"`
	TanggalTerbit   *string                           `json:"tanggal_terbit"`
	AnggotaKeluarga []AdminDetailKartuKeluargaAnggota `json:"anggota_keluarga"`
}

type AdminUpdateKartuKeluargaRequest struct {
	NoKK          string `json:"no_kk"`
	TanggalTerbit string `json:"tanggal_terbit"`
}

type AdminAkunKeluargaUsecase struct {
	kkRepo           *repositories.KartuKeluargaRepository
	kependudukanRepo *repositories.KependudukanRepository
}

func NewAdminAkunKeluargaUsecase(
	kkRepo *repositories.KartuKeluargaRepository,
	kependudukanRepo *repositories.KependudukanRepository,
) *AdminAkunKeluargaUsecase {
	return &AdminAkunKeluargaUsecase{
		kkRepo:           kkRepo,
		kependudukanRepo: kependudukanRepo,
	}
}

// CreateKartuKeluarga - Admin membuat Kartu Keluarga + Anggota (Penduduk) tanpa akun user
func (u *AdminAkunKeluargaUsecase) CreateKartuKeluarga(req *AdminCreateKartuKeluargaRequest) (map[string]interface{}, error) {
	if req == nil {
		return nil, customerror.NewBadRequestError("request tidak valid")
	}

	req.NoKK = strings.TrimSpace(req.NoKK)
	for i := range req.AnggotaKeluarga {
		req.AnggotaKeluarga[i].NIK = strings.TrimSpace(req.AnggotaKeluarga[i].NIK)
		req.AnggotaKeluarga[i].NamaLengkap = strings.TrimSpace(req.AnggotaKeluarga[i].NamaLengkap)
		req.AnggotaKeluarga[i].JenisKelamin = strings.TrimSpace(req.AnggotaKeluarga[i].JenisKelamin)
		req.AnggotaKeluarga[i].TanggalLahir = strings.TrimSpace(req.AnggotaKeluarga[i].TanggalLahir)
		req.AnggotaKeluarga[i].TempatLahir = strings.TrimSpace(req.AnggotaKeluarga[i].TempatLahir)
		req.AnggotaKeluarga[i].GolonganDarah = strings.TrimSpace(req.AnggotaKeluarga[i].GolonganDarah)
		req.AnggotaKeluarga[i].Agama = strings.TrimSpace(req.AnggotaKeluarga[i].Agama)
		req.AnggotaKeluarga[i].StatusPerkawinan = strings.TrimSpace(req.AnggotaKeluarga[i].StatusPerkawinan)
		req.AnggotaKeluarga[i].Pekerjaan = strings.TrimSpace(req.AnggotaKeluarga[i].Pekerjaan)
		req.AnggotaKeluarga[i].PendidikanTerakhir = strings.TrimSpace(req.AnggotaKeluarga[i].PendidikanTerakhir)
		req.AnggotaKeluarga[i].BacaHuruf = strings.TrimSpace(req.AnggotaKeluarga[i].BacaHuruf)
		req.AnggotaKeluarga[i].KedudukanKeluarga = strings.TrimSpace(req.AnggotaKeluarga[i].KedudukanKeluarga)
		req.AnggotaKeluarga[i].Dusun = strings.TrimSpace(req.AnggotaKeluarga[i].Dusun)
		req.AnggotaKeluarga[i].AsalPenduduk = strings.TrimSpace(req.AnggotaKeluarga[i].AsalPenduduk)
		req.AnggotaKeluarga[i].TujuanPindah = strings.TrimSpace(req.AnggotaKeluarga[i].TujuanPindah)
		req.AnggotaKeluarga[i].TempatMeninggal = strings.TrimSpace(req.AnggotaKeluarga[i].TempatMeninggal)
		req.AnggotaKeluarga[i].Keterangan = strings.TrimSpace(req.AnggotaKeluarga[i].Keterangan)
	}

	if req.NoKK == "" {
		return nil, customerror.NewBadRequestError("no_kk wajib diisi")
	}

	if len(req.AnggotaKeluarga) == 0 {
		return nil, customerror.NewBadRequestError("anggota_keluarga wajib diisi minimal 1 orang")
	}

	if _, err := u.kkRepo.FindByNoKK(req.NoKK); err == nil {
		return nil, customerror.NewBadRequestError("nomor KK sudah terdaftar")
	}

	nikSeen := map[string]struct{}{}
	for _, anggota := range req.AnggotaKeluarga {
		if anggota.NIK == "" || anggota.NamaLengkap == "" || anggota.TanggalLahir == "" {
			return nil, customerror.NewBadRequestError("setiap anggota wajib mengisi nik, nama_lengkap, dan tanggal_lahir")
		}

		if _, exists := nikSeen[anggota.NIK]; exists {
			return nil, customerror.NewBadRequestError("terdapat NIK duplikat pada anggota_keluarga")
		}
		nikSeen[anggota.NIK] = struct{}{}

		nikPtr := &anggota.NIK
		if _, err := u.kependudukanRepo.FindByNIK(nikPtr); err == nil {
			return nil, customerror.NewBadRequestError("NIK sudah terdaftar: " + anggota.NIK)
		}
	}

	var tanggalTerbit *time.Time
	if strings.TrimSpace(req.TanggalTerbit) != "" {
		parsed, err := time.Parse("2006-01-02", req.TanggalTerbit)
		if err != nil {
			return nil, customerror.NewBadRequestError("format tanggal_terbit harus YYYY-MM-DD")
		}
		tanggalTerbit = &parsed
	}

	kk := &models.KartuKeluarga{
		NoKK:          req.NoKK,
		TanggalTerbit: tanggalTerbit,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	if err := u.kkRepo.Create(kk); err != nil {
		return nil, err
	}

	createdPenduduk := make([]*models.Kependudukan, 0, len(req.AnggotaKeluarga))
	for _, anggota := range req.AnggotaKeluarga {
		tanggalLahir, err := time.Parse("2006-01-02", anggota.TanggalLahir)
		if err != nil {
			return nil, customerror.NewBadRequestError("format tanggal_lahir harus YYYY-MM-DD")
		}

		nikPtr := &anggota.NIK
		penduduk := &models.Kependudukan{
			KartuKeluargaID:    &kk.ID,
			NIK:                nikPtr,
			Dusun:              anggota.Dusun,
			NamaLengkap:        anggota.NamaLengkap,
			GolonganDarah:      anggota.GolonganDarah,
			JenisKelamin:       anggota.JenisKelamin,
			TempatLahir:        anggota.TempatLahir,
			TanggalLahir:       tanggalLahir,
			Agama:              anggota.Agama,
			StatusPerkawinan:   anggota.StatusPerkawinan,
			Pekerjaan:          anggota.Pekerjaan,
			PendidikanTerakhir: anggota.PendidikanTerakhir,
			BacaHuruf:          anggota.BacaHuruf,
			KedudukanKeluarga:  anggota.KedudukanKeluarga,
			AsalPenduduk:       anggota.AsalPenduduk,
			TujuanPindah:       anggota.TujuanPindah,
			TempatMeninggal:    anggota.TempatMeninggal,
			Keterangan:         anggota.Keterangan,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		}
		if err := u.kependudukanRepo.Create(penduduk); err != nil {
			return nil, err
		}
		createdPenduduk = append(createdPenduduk, penduduk)
	}

	// Mapping anggota untuk response
	resAnggota := make([]map[string]interface{}, 0, len(createdPenduduk))
	for _, p := range createdPenduduk {
		nik := ""
		if p.NIK != nil {
			nik = *p.NIK
		}
		resAnggota = append(resAnggota, map[string]interface{}{
			"penduduk_id":        p.IDKependudukan,
			"nik":                nik,
			"nama_lengkap":       p.NamaLengkap,
			"jenis_kelamin":      p.JenisKelamin,
			"tanggal_lahir":      p.TanggalLahir.Format("2006-01-02"),
			"kedudukan_keluarga": p.KedudukanKeluarga,
		})
	}

	return map[string]interface{}{
		"kartu_keluarga_id": kk.ID,
		"no_kk":             req.NoKK,
		"tanggal_terbit":    tanggalTerbit,
		"total_anggota":     len(createdPenduduk),
		"anggota_keluarga":  resAnggota,
	}, nil
}

func (u *AdminAkunKeluargaUsecase) ListKartuKeluarga(search string, page int, limit int, sortBy string, sortDir string) (map[string]interface{}, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	list, total, err := u.kkRepo.ListPaginated(search, page, limit, sortBy, sortDir)
	if err != nil {
		return nil, customerror.NewInternalServiceError("gagal mengambil data kartu keluarga")
	}

	items := make([]AdminListKartuKeluargaItem, 0, len(list))
	for _, kk := range list {
		anggota, err := u.kependudukanRepo.ListByKartuKeluargaID(kk.ID)
		if err != nil {
			return nil, customerror.NewInternalServiceError("gagal mengambil data anggota keluarga")
		}

		var kepala interface{}
		for _, a := range anggota {
			if strings.EqualFold(strings.TrimSpace(a.KedudukanKeluarga), "Kepala Keluarga") {
				nik := ""
				if a.NIK != nil {
					nik = *a.NIK
				}
				kepala = AdminKepalaKeluarga{
					PendudukID:  a.IDKependudukan,
					NIK:         nik,
					NamaLengkap: a.NamaLengkap,
				}
				break
			}
		}
		if kepala == nil && len(anggota) > 0 {
			nik := ""
			if anggota[0].NIK != nil {
				nik = *anggota[0].NIK
			}
			kepala = AdminKepalaKeluarga{
				PendudukID:  anggota[0].IDKependudukan,
				NIK:         nik,
				NamaLengkap: anggota[0].NamaLengkap,
			}
		}

		var tanggalTerbit *string
		if kk.TanggalTerbit != nil {
			formatted := kk.TanggalTerbit.Format("2006-01-02")
			tanggalTerbit = &formatted
		}

		items = append(items, AdminListKartuKeluargaItem{
			KartuKeluargaID: kk.ID,
			NoKK:            kk.NoKK,
			TanggalTerbit:   tanggalTerbit,
			KepalaKeluarga:  kepala,
			JumlahAnggota:   len(anggota),
			CreatedAt:       kk.CreatedAt,
			UpdatedAt:       kk.UpdatedAt,
		})
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	if total == 0 {
		totalPages = 0
	}

	return map[string]interface{}{
		"items": items,
		"pagination": AdminListKartuKeluargaPagination{
			Page:       page,
			Limit:      limit,
			Total:      int(total),
			TotalPages: totalPages,
		},
	}, nil
}

func (u *AdminAkunKeluargaUsecase) DetailKartuKeluarga(kartuKeluargaID int64) (*AdminDetailKartuKeluargaResponse, error) {
	kk, err := u.kkRepo.FindByID(kartuKeluargaID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, customerror.NewNotFoundError("kartu keluarga tidak ditemukan")
		}
		return nil, customerror.NewInternalServiceError("gagal mengambil data kartu keluarga")
	}

	anggota, err := u.kependudukanRepo.ListByKartuKeluargaID(kartuKeluargaID)
	if err != nil {
		return nil, customerror.NewInternalServiceError("gagal mengambil data anggota keluarga")
	}

	resAnggota := make([]AdminDetailKartuKeluargaAnggota, 0, len(anggota))
	for _, a := range anggota {
		resAnggota = append(resAnggota, mapPendudukToAnggota(a))
	}

	var tanggalTerbit *string
	if kk.TanggalTerbit != nil {
		formatted := kk.TanggalTerbit.Format("2006-01-02")
		tanggalTerbit = &formatted
	}

	return &AdminDetailKartuKeluargaResponse{
		KartuKeluargaID: kk.ID,
		NoKK:            kk.NoKK,
		TanggalTerbit:   tanggalTerbit,
		AnggotaKeluarga: resAnggota,
	}, nil
}

func (u *AdminAkunKeluargaUsecase) UpdateKartuKeluarga(kartuKeluargaID int64, req *AdminUpdateKartuKeluargaRequest) (*AdminDetailKartuKeluargaResponse, error) {
	if req == nil {
		return nil, customerror.NewBadRequestError("request tidak valid")
	}

	req.NoKK = strings.TrimSpace(req.NoKK)

	if req.NoKK == "" {
		return nil, customerror.NewBadRequestError("no_kk wajib diisi")
	}

	kk, err := u.kkRepo.FindByID(kartuKeluargaID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, customerror.NewNotFoundError("kartu keluarga tidak ditemukan")
		}
		return nil, customerror.NewInternalServiceError("gagal mengambil data kartu keluarga")
	}

	if _, err := u.kkRepo.FindByNoKKExceptID(req.NoKK, kartuKeluargaID); err == nil {
		return nil, customerror.NewConflictError("no_kk sudah dipakai")
	} else if err != gorm.ErrRecordNotFound {
		return nil, customerror.NewInternalServiceError("gagal validasi no_kk")
	}

	var tanggalTerbit *time.Time
	if strings.TrimSpace(req.TanggalTerbit) != "" {
		parsed, err := time.Parse("2006-01-02", req.TanggalTerbit)
		if err != nil {
			return nil, customerror.NewBadRequestError("format tanggal_terbit harus YYYY-MM-DD")
		}
		tanggalTerbit = &parsed
	}

	kk.NoKK = req.NoKK
	kk.TanggalTerbit = tanggalTerbit
	kk.UpdatedAt = time.Now()
	if err := u.kkRepo.Update(kk); err != nil {
		return nil, customerror.NewInternalServiceError("gagal memperbarui kartu keluarga")
	}

	return u.DetailKartuKeluarga(kartuKeluargaID)
}

func (u *AdminAkunKeluargaUsecase) UpdateAnggotaKeluarga(kartuKeluargaID int64, pendudukID int32, req *AdminAnggotaKeluargaRequest) (*AdminDetailKartuKeluargaAnggota, error) {
	if req == nil {
		return nil, customerror.NewBadRequestError("request tidak valid")
	}

	anggota, err := u.kependudukanRepo.FindByIDAndKartuKeluargaID(pendudukID, kartuKeluargaID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, customerror.NewNotFoundError("anggota keluarga tidak ditemukan")
		}
		return nil, customerror.NewInternalServiceError("gagal mengambil data anggota keluarga")
	}

	req.NIK = strings.TrimSpace(req.NIK)
	req.NamaLengkap = strings.TrimSpace(req.NamaLengkap)
	req.JenisKelamin = strings.TrimSpace(req.JenisKelamin)
	req.TanggalLahir = strings.TrimSpace(req.TanggalLahir)
	req.TempatLahir = strings.TrimSpace(req.TempatLahir)
	req.GolonganDarah = strings.TrimSpace(req.GolonganDarah)
	req.Agama = strings.TrimSpace(req.Agama)
	req.StatusPerkawinan = strings.TrimSpace(req.StatusPerkawinan)
	req.Pekerjaan = strings.TrimSpace(req.Pekerjaan)
	req.PendidikanTerakhir = strings.TrimSpace(req.PendidikanTerakhir)
	req.BacaHuruf = strings.TrimSpace(req.BacaHuruf)
	req.KedudukanKeluarga = strings.TrimSpace(req.KedudukanKeluarga)
	req.Dusun = strings.TrimSpace(req.Dusun)
	req.AsalPenduduk = strings.TrimSpace(req.AsalPenduduk)
	req.TujuanPindah = strings.TrimSpace(req.TujuanPindah)
	req.TempatMeninggal = strings.TrimSpace(req.TempatMeninggal)
	req.Keterangan = strings.TrimSpace(req.Keterangan)

	if req.NIK == "" || req.NamaLengkap == "" {
		return nil, customerror.NewBadRequestError("nik dan nama_lengkap wajib diisi")
	}

	if _, err := u.kependudukanRepo.FindByNIKExceptID(req.NIK, pendudukID); err == nil {
		return nil, customerror.NewConflictError("nik sudah dipakai")
	} else if err != gorm.ErrRecordNotFound {
		return nil, customerror.NewInternalServiceError("gagal validasi nik")
	}

	if req.TanggalLahir != "" {
		parsedTanggalLahir, err := time.Parse("2006-01-02", req.TanggalLahir)
		if err != nil {
			return nil, customerror.NewBadRequestError("format tanggal_lahir harus YYYY-MM-DD")
		}
		anggota.TanggalLahir = parsedTanggalLahir
	}

	nikPtr := &req.NIK
	anggota.NIK = nikPtr
	anggota.NamaLengkap = req.NamaLengkap
	anggota.JenisKelamin = req.JenisKelamin
	anggota.TempatLahir = req.TempatLahir
	anggota.GolonganDarah = req.GolonganDarah
	anggota.Agama = req.Agama
	anggota.StatusPerkawinan = req.StatusPerkawinan
	anggota.Pekerjaan = req.Pekerjaan
	anggota.PendidikanTerakhir = req.PendidikanTerakhir
	anggota.BacaHuruf = req.BacaHuruf
	anggota.KedudukanKeluarga = req.KedudukanKeluarga
	anggota.Dusun = req.Dusun
	anggota.AsalPenduduk = req.AsalPenduduk
	anggota.TujuanPindah = req.TujuanPindah
	anggota.TempatMeninggal = req.TempatMeninggal
	anggota.Keterangan = req.Keterangan
	anggota.UpdatedAt = time.Now()

	if err := u.kependudukanRepo.Update(anggota); err != nil {
		return nil, customerror.NewInternalServiceError("gagal memperbarui anggota keluarga")
	}

	res := mapPendudukToAnggota(*anggota)
	return &res, nil
}

func (u *AdminAkunKeluargaUsecase) AddAnggotaKeluarga(kartuKeluargaID int64, req *AdminAnggotaKeluargaRequest) (*AdminDetailKartuKeluargaAnggota, error) {
	if req == nil {
		return nil, customerror.NewBadRequestError("request tidak valid")
	}

	if _, err := u.kkRepo.FindByID(kartuKeluargaID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, customerror.NewNotFoundError("kartu keluarga tidak ditemukan")
		}
		return nil, customerror.NewInternalServiceError("gagal mengambil data kartu keluarga")
	}

	req.NIK = strings.TrimSpace(req.NIK)
	req.NamaLengkap = strings.TrimSpace(req.NamaLengkap)
	req.TanggalLahir = strings.TrimSpace(req.TanggalLahir)
	if req.NIK == "" || req.NamaLengkap == "" || req.TanggalLahir == "" {
		return nil, customerror.NewBadRequestError("nik, nama_lengkap, dan tanggal_lahir wajib diisi")
	}

	reqNIKPtr := &req.NIK
	if _, err := u.kependudukanRepo.FindByNIK(reqNIKPtr); err == nil {
		return nil, customerror.NewConflictError("nik sudah dipakai")
	} else if err != gorm.ErrRecordNotFound {
		return nil, customerror.NewInternalServiceError("gagal validasi nik")
	}

	tanggalLahir, err := time.Parse("2006-01-02", req.TanggalLahir)
	if err != nil {
		return nil, customerror.NewBadRequestError("format tanggal_lahir harus YYYY-MM-DD")
	}

	anggota := &models.Kependudukan{
		KartuKeluargaID:    &kartuKeluargaID,
		NIK:                reqNIKPtr,
		NamaLengkap:        req.NamaLengkap,
		JenisKelamin:       strings.TrimSpace(req.JenisKelamin),
		TanggalLahir:       tanggalLahir,
		TempatLahir:        strings.TrimSpace(req.TempatLahir),
		GolonganDarah:      strings.TrimSpace(req.GolonganDarah),
		Agama:              strings.TrimSpace(req.Agama),
		StatusPerkawinan:   strings.TrimSpace(req.StatusPerkawinan),
		Pekerjaan:          strings.TrimSpace(req.Pekerjaan),
		PendidikanTerakhir: strings.TrimSpace(req.PendidikanTerakhir),
		BacaHuruf:          strings.TrimSpace(req.BacaHuruf),
		KedudukanKeluarga:  strings.TrimSpace(req.KedudukanKeluarga),
		Dusun:              strings.TrimSpace(req.Dusun),
		AsalPenduduk:       strings.TrimSpace(req.AsalPenduduk),
		TujuanPindah:       strings.TrimSpace(req.TujuanPindah),
		TempatMeninggal:    strings.TrimSpace(req.TempatMeninggal),
		Keterangan:         strings.TrimSpace(req.Keterangan),
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	if err := u.kependudukanRepo.Create(anggota); err != nil {
		return nil, customerror.NewInternalServiceError("gagal menambahkan anggota keluarga")
	}

	res := mapPendudukToAnggota(*anggota)
	return &res, nil
}

func (u *AdminAkunKeluargaUsecase) DeleteAnggotaKeluarga(kartuKeluargaID int64, pendudukID int32) error {
	anggota, err := u.kependudukanRepo.FindByIDAndKartuKeluargaID(pendudukID, kartuKeluargaID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return customerror.NewNotFoundError("anggota keluarga tidak ditemukan")
		}
		return customerror.NewInternalServiceError("gagal mengambil data anggota keluarga")
	}

	if err := u.kependudukanRepo.SoftDeleteByID(anggota.IDKependudukan); err != nil {
		return customerror.NewInternalServiceError("gagal menghapus anggota keluarga")
	}

	return nil
}

func (u *AdminAkunKeluargaUsecase) DeleteKartuKeluarga(kartuKeluargaID int64) error {
	if _, err := u.kkRepo.FindByID(kartuKeluargaID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return customerror.NewNotFoundError("kartu keluarga tidak ditemukan")
		}
		return customerror.NewInternalServiceError("gagal mengambil data kartu keluarga")
	}

	if err := u.kependudukanRepo.SoftDeleteByKartuKeluargaID(kartuKeluargaID); err != nil {
		return customerror.NewInternalServiceError("gagal menghapus anggota keluarga")
	}

	if err := u.kkRepo.SoftDeleteByID(kartuKeluargaID); err != nil {
		return customerror.NewInternalServiceError("gagal menghapus kartu keluarga")
	}

	return nil
}

func mapPendudukToAnggota(a models.Kependudukan) AdminDetailKartuKeluargaAnggota {
	nik := ""
	if a.NIK != nil {
		nik = *a.NIK
	}
	return AdminDetailKartuKeluargaAnggota{
		PendudukID:         a.IDKependudukan,
		NIK:                nik,
		NamaLengkap:        a.NamaLengkap,
		JenisKelamin:       a.JenisKelamin,
		TanggalLahir:       a.TanggalLahir.Format("2006-01-02"),
		TempatLahir:        a.TempatLahir,
		GolonganDarah:      a.GolonganDarah,
		Agama:              a.Agama,
		StatusPerkawinan:   a.StatusPerkawinan,
		Pekerjaan:          a.Pekerjaan,
		PendidikanTerakhir: a.PendidikanTerakhir,
		BacaHuruf:          a.BacaHuruf,
		KedudukanKeluarga:  a.KedudukanKeluarga,
		Dusun:              a.Dusun,
		AsalPenduduk:       a.AsalPenduduk,
		TujuanPindah:       a.TujuanPindah,
		TempatMeninggal:    a.TempatMeninggal,
		Keterangan:         a.Keterangan,
	}
}
