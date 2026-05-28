package usecases

import (
	"errors"
	"fmt"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
	"strconv"
	"strings"
	"time"
)

var ErrLembarPemantauanForbidden = errors.New("akses lembar pemantauan ditolak")

type LembarPemantauanUsecase interface {
	Create(req models.LembarPemantauanRequest) (*models.LembarPemantauan, error)
	CreateForIbu(req models.LembarPemantauanRequest, userID uint) (*models.LembarPemantauan, error)
	GetRentangUsia() ([]models.RentangUsia, error)
	GetKategoriTandaSakitByRentangUsiaID(rentangUsiaID string) ([]models.KategoriTandaSakit, error)
	GetByID(id string) (*models.LembarPemantauan, error)
	GetByAnakID(anakID string) ([]models.LembarPemantauan, error)
	GetByAnakIDForIbu(anakID string, userID uint) ([]models.LembarPemantauan, error)
	GetAll() ([]models.LembarPemantauan, error)
	Update(id string, req models.LembarPemantauanRequest) (*models.LembarPemantauan, error)
	Verify(id string, req models.LembarPemantauanVerifikasiRequest) (*models.LembarPemantauan, error)
	Delete(id string) error
}

type lembarPemantauanUsecase struct {
	repository repositories.LembarPemantauanRepository
}

func NewLembarPemantauanUsecase(repo repositories.LembarPemantauanRepository) LembarPemantauanUsecase {
	return &lembarPemantauanUsecase{
		repository: repo,
	}
}

// Helper function - Convert string ke uint
func parseUint(id string, fieldName string) (uint, error) {
	id = strings.TrimSpace(id)
	if id == "" {
		return 0, fmt.Errorf("%s tidak boleh kosong", fieldName)
	}

	val, err := strconv.ParseUint(id, 10, 32)
	if err != nil || val <= 0 {
		return 0, fmt.Errorf("%s harus berupa angka lebih dari 0", fieldName)
	}

	return uint(val), nil
}

func (u *lembarPemantauanUsecase) Create(req models.LembarPemantauanRequest) (*models.LembarPemantauan, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Validasi Max Periode agar tidak melewati batas buku KIA
	rentang, err := u.repository.FindRentangUsiaByID(req.RentangUsiaID)
	if err != nil {
		return nil, err
	}
	if req.PeriodeWaktu > rentang.MaxPeriode {
		return nil, fmt.Errorf("periode waktu melampaui batas maksimal (%d) untuk rentang usia ini", rentang.MaxPeriode)
	}

	tanggalPeriksa, err := time.Parse("2006-01-02", req.TanggalPeriksa)
	if err != nil {
		return nil, errors.New("format tanggal_periksa harus YYYY-MM-DD")
	}

	lembar := &models.LembarPemantauan{
		AnakID:         req.AnakID,
		RentangUsiaID:  req.RentangUsiaID,
		PeriodeWaktu:   req.PeriodeWaktu,
		TanggalPeriksa: &tanggalPeriksa,
		NamaPemeriksa:  req.NamaPemeriksa,
		Status:         "Menunggu verifikasi",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	detailGejala := make([]models.DetailPemantauan, 0, len(req.DetailGejala))
	for _, detail := range req.DetailGejala {
		if detail.KategoriTandaSakitID <= 0 {
			return nil, errors.New("kategori_tanda_sakit_id harus lebih dari 0")
		}
		detailGejala = append(detailGejala, models.DetailPemantauan{
			KategoriTandaSakitID: detail.KategoriTandaSakitID,
			IsTerjadi:            detail.IsTerjadi,
			CreatedAt:            time.Now(),
			UpdatedAt:            time.Now(),
		})
	}

	lembar.DetailGejala = detailGejala

	if err := u.repository.Create(lembar); err != nil {
		return nil, err
	}

	return lembar, nil
}

func (u *lembarPemantauanUsecase) CreateForIbu(req models.LembarPemantauanRequest, userID uint) (*models.LembarPemantauan, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	ok, err := u.repository.IsAnakMilikIbu(userID, req.AnakID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrLembarPemantauanForbidden
	}

	// req.NamaPemeriksa = "Orang Tua" // Hardcode jika yang isi ibu
	return u.Create(req)
}

func (u *lembarPemantauanUsecase) GetRentangUsia() ([]models.RentangUsia, error) {
	return u.repository.FindRentangUsia()
}

func (u *lembarPemantauanUsecase) GetKategoriTandaSakitByRentangUsiaID(rentangUsiaID string) ([]models.KategoriTandaSakit, error) {
	id, err := parseUint(rentangUsiaID, "rentang_usia_id")
	if err != nil {
		return nil, err
	}
	return u.repository.FindKategoriTandaSakitByRentangUsiaID(id)
}

func (u *lembarPemantauanUsecase) GetByID(id string) (*models.LembarPemantauan, error) {
	lembarID, err := parseUint(id, "lembar_pemantauan_id")
	if err != nil {
		return nil, err
	}
	return u.repository.FindByID(lembarID)
}

func (u *lembarPemantauanUsecase) GetByAnakID(anakID string) ([]models.LembarPemantauan, error) {
	id, err := parseUint(anakID, "anak_id")
	if err != nil {
		return nil, err
	}
	return u.repository.FindByAnakID(id)
}

func (u *lembarPemantauanUsecase) GetByAnakIDForIbu(anakID string, userID uint) ([]models.LembarPemantauan, error) {
	id, err := parseUint(anakID, "anak_id")
	if err != nil {
		return nil, err
	}

	ok, err := u.repository.IsAnakMilikIbu(userID, id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrLembarPemantauanForbidden
	}

	return u.repository.FindByAnakID(id)
}

func (u *lembarPemantauanUsecase) GetAll() ([]models.LembarPemantauan, error) {
	return u.repository.FindAll()
}

func (u *lembarPemantauanUsecase) Update(id string, req models.LembarPemantauanRequest) (*models.LembarPemantauan, error) {
	lembarID, err := parseUint(id, "lembar_pemantauan_id")
	if err != nil {
		return nil, err
	}

	if err := req.Validate(); err != nil {
		return nil, err
	}

	lembar, err := u.repository.FindByID(lembarID)
	if err != nil {
		return nil, err
	}

	tanggalPeriksa, err := time.Parse("2006-01-02", req.TanggalPeriksa)
	if err != nil {
		return nil, errors.New("format tanggal_periksa harus YYYY-MM-DD")
	}

	lembar.RentangUsiaID = req.RentangUsiaID
	lembar.PeriodeWaktu = req.PeriodeWaktu
	lembar.TanggalPeriksa = &tanggalPeriksa
	lembar.NamaPemeriksa = req.NamaPemeriksa
	lembar.UpdatedAt = time.Now()

	detailGejala := make([]models.DetailPemantauan, 0, len(req.DetailGejala))
	for _, detail := range req.DetailGejala {
		if detail.KategoriTandaSakitID <= 0 {
			return nil, errors.New("kategori_tanda_sakit_id harus lebih dari 0")
		}
		detailGejala = append(detailGejala, models.DetailPemantauan{
			KategoriTandaSakitID: detail.KategoriTandaSakitID,
			IsTerjadi:            detail.IsTerjadi,
			CreatedAt:            time.Now(),
			UpdatedAt:            time.Now(),
		})
	}
	lembar.DetailGejala = detailGejala

	if err := u.repository.Update(lembar); err != nil {
		return nil, err
	}
	return lembar, nil
}

func (u *lembarPemantauanUsecase) Verify(id string, req models.LembarPemantauanVerifikasiRequest) (*models.LembarPemantauan, error) {
	lembarID, err := parseUint(id, "lembar_pemantauan_id")
	if err != nil {
		return nil, err
	}

	if err := req.Validate(); err != nil {
		return nil, err
	}

	lembar, err := u.repository.FindByID(lembarID)
	if err != nil {
		return nil, err
	}

	if err := u.repository.Verify(lembarID, req.Status, req.NamaPemeriksa); err != nil {
		return nil, err
	}

	lembar.Status = req.Status
	lembar.NamaPemeriksa = req.NamaPemeriksa

	return lembar, nil
}

func (u *lembarPemantauanUsecase) Delete(id string) error {
	lembarID, err := parseUint(id, "lembar_pemantauan_id")
	if err != nil {
		return err
	}
	return u.repository.Delete(lembarID)
}
