package usecases

import (
	"errors"
	"fmt"
	"strings"

	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
	"time"
)

type IbuUsecase interface {
	Create(ibu *models.Ibu) (*models.Ibu, error)
	GetByID(id int32) (*models.Ibu, error)
	GetAll() ([]models.Ibu, error)
	GetAnakSaya(userID int32) ([]models.AnakResponse, error)
	Update(ibu *models.Ibu) error
	Delete(id int32) error
	GetDashboard() ([]models.IbuDashboardDTO, error)
	GetByPendudukID(pendudukID int32) (*models.Ibu, error)
}

type ibuUsecase struct {
	repo         *repositories.IbuRepository
	pendudukRepo *repositories.KependudukanRepository
}

func NewIbuUsecase(
	repo *repositories.IbuRepository,
	pendudukRepo *repositories.KependudukanRepository,
) IbuUsecase {
	return &ibuUsecase{
		repo:         repo,
		pendudukRepo: pendudukRepo,
	}
}

func (u *ibuUsecase) Create(ibu *models.Ibu) (*models.Ibu, error) {
	// ================= VALIDASI IBU =================
	if ibu.IDKependudukan == 0 {
		return nil, errors.New("id_kependudukan wajib diisi")
	}

	dataIbu, err := u.pendudukRepo.FindByID(ibu.IDKependudukan)
	if err != nil {
		return nil, errors.New("data penduduk ibu tidak ditemukan")
	}

	if strings.ToLower(dataIbu.JenisKelamin) != "perempuan" {
		return nil, errors.New("penduduk yang dipilih bukan perempuan")
	}

	// ================= VALIDASI SUAMI =================
	if ibu.IDSuami != nil {
		if *ibu.IDSuami == ibu.IDKependudukan {
			return nil, errors.New("suami tidak boleh sama dengan ibu")
		}
		dataSuami, err := u.pendudukRepo.FindByID(*ibu.IDSuami)
		if err != nil {
			return nil, errors.New("data suami tidak ditemukan")
		}
		if strings.ToLower(dataSuami.JenisKelamin) != "laki-laki" {
			return nil, errors.New("penduduk yang dipilih sebagai suami bukan laki-laki")
		}
	}

	// ================= CEK DUPLIKAT =================
	existing, err := u.repo.FindByPendudukID(ibu.IDKependudukan)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return existing, nil //
	}

	// ================= BUAT BARU =================
	if err := u.repo.Create(ibu); err != nil {
		return nil, err
	}
	// Ambil data yang baru dibuat (agar lengkap dengan ID)
	return u.repo.FindByPendudukID(ibu.IDKependudukan)
}

func (u *ibuUsecase) GetByID(id int32) (*models.Ibu, error) {
	return u.repo.FindByID(id)
}

func (u *ibuUsecase) GetAll() ([]models.Ibu, error) {
	return u.repo.FindAll()
}

func mapAnakToResponse(data models.Anak) models.AnakResponse {
	res := models.AnakResponse{
		ID:            data.ID,
		KehamilanID:   data.KehamilanID,
		PendudukID:    data.PendudukID,
		BeratLahirKg:  data.BeratLahirKg,
		TinggiLahirCm: data.TinggiLahirCm,
	}

	if data.Penduduk != nil {
		res.Nama = data.Penduduk.NamaLengkap
		res.TanggalLahir = data.Penduduk.TanggalLahir.Format("2006-01-02")
		res.JenisKelamin = data.Penduduk.JenisKelamin
		res.GolonganDarah = data.Penduduk.GolonganDarah

		now := time.Now()
		ageMonths := (now.Year()-data.Penduduk.TanggalLahir.Year())*12 + int(now.Month()-data.Penduduk.TanggalLahir.Month())
		if now.Day() < data.Penduduk.TanggalLahir.Day() {
			ageMonths--
		}
		if ageMonths < 0 {
			ageMonths = 0
		}
		res.UsiaBulan = ageMonths
		res.UsiaTeks = fmt.Sprintf("%d Tahun %d Bulan", ageMonths/12, ageMonths%12)
	}

	if data.Kehamilan != nil {
		res.Kehamilan = &models.KehamilanSimple{ID: data.Kehamilan.ID}
	}

	return res
}

func (u *ibuUsecase) GetAnakSaya(userID int32) ([]models.AnakResponse, error) {
	data, err := u.repo.FindAnakByUserID(userID)
	if err != nil {
		return nil, err
	}

	result := make([]models.AnakResponse, 0, len(data))
	for _, anak := range data {
		result = append(result, mapAnakToResponse(anak))
	}

	return result, nil
}

func (u *ibuUsecase) Update(ibu *models.Ibu) error {

	_, err := u.repo.FindByID(ibu.IDIbu)
	if err != nil {
		return errors.New("data ibu tidak ditemukan")
	}

	return u.repo.Update(ibu)
}

func (u *ibuUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}

func (u *ibuUsecase) GetDashboard() ([]models.IbuDashboardDTO, error) {
	return u.repo.GetDashboard()
}
func (u *ibuUsecase) GetByPendudukID(pendudukID int32) (*models.Ibu, error) {
	return u.repo.FindByPendudukID(pendudukID)
}
