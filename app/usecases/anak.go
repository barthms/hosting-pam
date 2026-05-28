package usecases

import (
	"errors"
	"fmt"
	"time"

	"monitoring-service/app/models"
	"monitoring-service/app/repositories"

	"gorm.io/gorm"
)

type AnakUseCase struct {
	anakRepo        *repositories.AnakRepository
	kependudukanRepo *repositories.KependudukanRepository
}

func NewAnakUseCase(anakRepo *repositories.AnakRepository, kependudukanRepo *repositories.KependudukanRepository) *AnakUseCase {
	return &AnakUseCase{
		anakRepo:        anakRepo,
		kependudukanRepo: kependudukanRepo,
	}
}

// ====================== GET ======================
func (u *AnakUseCase) GetAnak(id int32) (*models.AnakResponse, error) {
	anak, err := u.anakRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("data anak tidak ditemukan")
		}
		return nil, err
	}

	resp := u.toAnakResponse(anak)
	return &resp, nil
}

// ====================== CREATE ======================
func (u *AnakUseCase) CreateAnak(req models.CreateAnakRequest) (*models.AnakResponse, error) {
	// Validasi input
	if req.KehamilanID == 0 {
		return nil, errors.New("kehamilan_id wajib diisi")
	}

	pendudukID := req.PendudukID
	if pendudukID == 0 {
		// Jika penduduk_id tidak ada, buat penduduk baru dari data yang dikirim
		if req.Nama == "" {
			return nil, errors.New("nama anak wajib diisi jika penduduk_id tidak ada")
		}

		tglLahir, _ := time.Parse("2006-01-02", req.TanggalLahir)
		// Generate NIK sementara jika tidak ada (karena NOT NULL di DB)
		nik := fmt.Sprintf("A%d", time.Now().UnixNano())

		newPenduduk := &models.Kependudukan{
			NIK:          &nik,
			NamaLengkap:  req.Nama,
			JenisKelamin: req.JenisKelamin,
			TanggalLahir: tglLahir,
		}

		if err := u.kependudukanRepo.Create(newPenduduk); err != nil {
			return nil, fmt.Errorf("gagal membuat data penduduk: %v", err)
		}
		pendudukID = newPenduduk.IDKependudukan
	}

	anak := &models.Anak{
		KehamilanID:     req.KehamilanID,
		PendudukID:      pendudukID, // Gunakan pendudukID (bisa dari auto-create atau req.PendudukID)
		BeratLahirKg:    req.BeratLahirKg,
		TinggiLahirCm:   req.TinggiLahirCm,
		AnakKe:          req.AnakKe,
		LingkarKepalaCm: req.LingkarKepalaCm,
		NamaIbu:         req.NamaIbu,
		NamaAyah:        req.NamaAyah,
		IbuID:           req.IbuID,
	}

	if err := u.anakRepo.Create(anak); err != nil {
		return nil, err
	}

	// Fetch complete data with relations
	createdAnak, err := u.anakRepo.FindByID(anak.ID)
	if err != nil {
		return nil, err
	}

	resp := u.toAnakResponse(createdAnak)
	return &resp, nil
}

// CreateAnakDenganPenduduk: create anak + auto-create kependudukan jika diperlukan
func (u *AnakUseCase) CreateAnakDenganPenduduk(req models.CreateAnakDenganPendudukRequest) (*models.AnakResponse, error) {
	// Validasi input
	if req.KehamilanID == 0 {
		return nil, errors.New("kehamilan_id wajib diisi")
	}
	if req.IbuID == 0 {
		return nil, errors.New("ibu_id wajib diisi")
	}
	if req.Nama == "" {
		return nil, errors.New("nama anak wajib diisi")
	}
	if req.TanggalLahir == "" {
		return nil, errors.New("tanggal_lahir anak wajib diisi")
	}
	if req.JenisKelamin == "" {
		return nil, errors.New("jenis_kelamin wajib diisi")
	}

	// Parse tanggal lahir
	tanggalLahir, err := time.Parse("2006-01-02", req.TanggalLahir)
	if err != nil {
		return nil, errors.New("format tanggal_lahir harus YYYY-MM-DD")
	}

	// Buat kependudukan baru untuk anak
	newPenduduk := &models.Kependudukan{
		NamaLengkap:  req.Nama,
		JenisKelamin: req.JenisKelamin,
		TanggalLahir: tanggalLahir,
		TempatLahir:  req.TempatLahir,
		GolonganDarah: req.GolonganDarah,
		// NIK tidak diisi untuk newborn, akan disimpan sebagai NULL
	}

	if err := u.kependudukanRepo.Create(newPenduduk); err != nil {
		return nil, fmt.Errorf("gagal membuat data penduduk anak: %w", err)
	}

	// Buat anak dengan penduduk_id dari kependudukan yang baru dibuat
	anak := &models.Anak{
		KehamilanID:     req.KehamilanID,
		PendudukID:      newPenduduk.IDKependudukan,
		BeratLahirKg:    req.BeratLahirKg,
		TinggiLahirCm:   req.TinggiLahirCm,
		AnakKe:          req.AnakKe,
		LingkarKepalaCm: req.LingkarKepalaCm,
		NamaIbu:         req.NamaIbu,
		NamaAyah:        req.NamaAyah,
		IbuID:           req.IbuID,
	}

	if err := u.anakRepo.Create(anak); err != nil {
		return nil, fmt.Errorf("gagal membuat data anak: %w", err)
	}

	// Fetch complete data with relations
	createdAnak, err := u.anakRepo.FindByID(anak.ID)
	if err != nil {
		return nil, err
	}

	resp := u.toAnakResponse(createdAnak)
	return &resp, nil
}

// ====================== UPDATE ======================
func (u *AnakUseCase) UpdateAnak(id int32, req models.UpdateAnakRequest) (*models.AnakResponse, error) {
	anak, err := u.anakRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("data anak tidak ditemukan")
	}

	if req.BeratLahirKg != nil {
		anak.BeratLahirKg = req.BeratLahirKg
	}

	if req.TinggiLahirCm != nil {
		anak.TinggiLahirCm = req.TinggiLahirCm
	}

	if req.AnakKe != nil {
		anak.AnakKe = *req.AnakKe
	}

	if req.LingkarKepalaCm != nil {
		anak.LingkarKepalaCm = req.LingkarKepalaCm
	}

	if req.NamaIbu != nil {
		anak.NamaIbu = *req.NamaIbu
	}

	if req.NamaAyah != nil {
		anak.NamaAyah = *req.NamaAyah
	}

	if err := u.anakRepo.Update(anak); err != nil {
		return nil, err
	}

	// ── Perbarui data Kependudukan (Penduduk) jika ada perubahan ──
	if anak.PendudukID > 0 && (req.Nama != "" || req.JenisKelamin != "" || req.TanggalLahir != "") {
		penduduk, pendudukErr := u.kependudukanRepo.FindByID(anak.PendudukID)
		if pendudukErr == nil && penduduk != nil {
			changed := false
			if req.Nama != "" && req.Nama != penduduk.NamaLengkap {
				penduduk.NamaLengkap = req.Nama
				changed = true
			}
			if req.JenisKelamin != "" && req.JenisKelamin != penduduk.JenisKelamin {
				penduduk.JenisKelamin = req.JenisKelamin
				changed = true
			}
			if req.TanggalLahir != "" {
				tgl, parseErr := time.Parse("2006-01-02", req.TanggalLahir)
				if parseErr == nil && !tgl.IsZero() {
					currentTgl := penduduk.TanggalLahir.Format("2006-01-02")
					if currentTgl != req.TanggalLahir {
						penduduk.TanggalLahir = tgl
						changed = true
					}
				}
			}
			if changed {
				_ = u.kependudukanRepo.Update(penduduk)
			}
		}
	}

	// Re-fetch untuk mendapatkan data Penduduk terbaru
	updatedAnak, err := u.anakRepo.FindByID(id)
	if err != nil {
		resp := u.toAnakResponse(anak)
		return &resp, nil
	}
	resp := u.toAnakResponse(updatedAnak)
	return &resp, nil
}


// ====================== DELETE ======================
func (u *AnakUseCase) DeleteAnak(id int32) error {
	return u.anakRepo.Delete(id)
}

// ====================== LIST ======================
func (u *AnakUseCase) AdminListAnak(kehamilanID int32) ([]models.AnakResponse, error) {
	var (
		list []models.Anak
		err  error
	)

	// FIX: harus dibandingkan dengan 0
	if kehamilanID != 0 {
		list, err = u.anakRepo.FindByKehamilanID(kehamilanID)
	} else {
		list, err = u.anakRepo.FindAll()
	}

	if err != nil {
		return nil, err
	}

	result := make([]models.AnakResponse, 0, len(list))

	for _, k := range list {
		resp := u.toAnakResponse(&k)
		result = append(result, resp)
	}

	return result, nil
}

// ====================== UTIL ======================
func HitungUsiaBulan(tanggalLahir time.Time) int {
	now := time.Now()

	if tanggalLahir.IsZero() || tanggalLahir.Year() < 1900 || tanggalLahir.After(now) {
		return 0
	}

	years := now.Year() - tanggalLahir.Year()
	months := int(now.Month()) - int(tanggalLahir.Month())
	total := years*12 + months

	if now.Day() < tanggalLahir.Day() {
		total--
	}

	if total < 0 {
		return 0
	}

	return total
}

func FormatUsiaTeks(bulan int) string {
	if bulan == 0 {
		return "0 bulan"
	}
	if bulan < 12 {
		return fmt.Sprintf("%d bulan", bulan)
	}

	tahun := bulan / 12
	sisa := bulan % 12

	if sisa == 0 {
		return fmt.Sprintf("%d tahun", tahun)
	}

	return fmt.Sprintf("%d tahun %d bulan", tahun, sisa)
}

func FormatLabelUsia(bulan int) string {
	if bulan == 0 {
		return "Baru Lahir (0 Bulan)"
	}
	if bulan < 12 {
		return fmt.Sprintf("%d Bulan", bulan)
	}

	tahun := bulan / 12
	sisa := bulan % 12

	if sisa == 0 {
		return fmt.Sprintf("%d Tahun", tahun)
	}

	return fmt.Sprintf("%d Tahun %d Bulan", tahun, sisa)
}

// ====================== MAPPER ======================
func (u *AnakUseCase) toAnakResponse(anak *models.Anak) models.AnakResponse {
	resp := models.AnakResponse{
		ID:              anak.ID,
		KehamilanID:     anak.KehamilanID,
		PendudukID:      anak.PendudukID,
		BeratLahirKg:    anak.BeratLahirKg,
		TinggiLahirCm:   anak.TinggiLahirCm,
		AnakKe:          anak.AnakKe,
		LingkarKepalaCm: anak.LingkarKepalaCm,
		NamaIbu:         anak.NamaIbu,
		NamaAyah:        anak.NamaAyah,
		IbuID:           anak.IbuID,
	}

	// Ambil data dari Penduduk (Kependudukan)
	if anak.Penduduk != nil {
		resp.Nama = anak.Penduduk.NamaLengkap
		if !anak.Penduduk.TanggalLahir.IsZero() && anak.Penduduk.TanggalLahir.Year() >= 1900 {
			resp.TanggalLahir = anak.Penduduk.TanggalLahir.Format("2006-01-02")
			// Hitung usia dari tanggal lahir penduduk
			usiaBulan := HitungUsiaBulan(anak.Penduduk.TanggalLahir)
			resp.UsiaBulan = usiaBulan
			resp.UsiaTeks = FormatUsiaTeks(usiaBulan)
		} else {
			resp.TanggalLahir = ""
			resp.UsiaBulan = 0
			resp.UsiaTeks = "-"
		}
		resp.JenisKelamin = anak.Penduduk.JenisKelamin
		resp.GolonganDarah = anak.Penduduk.GolonganDarah
	}

	// Ambil data Kehamilan
	if anak.Kehamilan != nil {
		resp.Kehamilan = &models.KehamilanSimple{
			ID: anak.Kehamilan.ID,
		}

		if anak.Kehamilan.Ibu != nil && anak.Kehamilan.Ibu.Kependudukan != nil {
			resp.Kehamilan.Ibu.NamaIbu = anak.Kehamilan.Ibu.Kependudukan.NamaLengkap
		}
	}

	// Map data Pertumbuhan
	// if len(anak.Pertumbuhan) > 0 {
	// 	resp.Pertumbuhan = make([]models.PertumbuhanSimple, 0, len(anak.Pertumbuhan))
	// 	for _, p := range anak.Pertumbuhan {
	// 		resp.Pertumbuhan = append(resp.Pertumbuhan, models.PertumbuhanSimple{
	// 			Bulan:       p.UsiaUkurBulan,
	// 			BeratBadan:  p.BeratBadan,
	// 			TinggiBadan: p.TinggiBadan,
	// 			HasilLila:   p.HasilLila,
	// 		})
	// 	}
	// }

	return resp
}
