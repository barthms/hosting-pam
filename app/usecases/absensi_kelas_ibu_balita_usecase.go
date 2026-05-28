package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type AbsensiKelasIbuBalitaUsecase interface {
	GetMine(userID int32) ([]models.AbsensiKelasIbuBalita, error)
	SaveMine(userID int32, req models.AbsensiKelasIbuBalita) (*models.AbsensiKelasIbuBalita, error)
}

type absensiKelasIbuBalitaUsecase struct {
	repo *repositories.AbsensiKelasIbuBalitaRepository
}

func NewAbsensiKelasIbuBalitaUsecase(
	repo *repositories.AbsensiKelasIbuBalitaRepository,
) AbsensiKelasIbuBalitaUsecase {
	return &absensiKelasIbuBalitaUsecase{repo: repo}
}

func (u *absensiKelasIbuBalitaUsecase) GetMine(userID int32) ([]models.AbsensiKelasIbuBalita, error) {
	if userID == 0 {
		return nil, errors.New("user_id tidak valid")
	}

	ibuID, err := u.repo.FindIbuIDByUserID(userID)
	if err != nil {
		return nil, errors.New("data ibu tidak ditemukan")
	}

	return u.repo.FindByIbuID(ibuID)
}

func (u *absensiKelasIbuBalitaUsecase) SaveMine(
	userID int32,
	req models.AbsensiKelasIbuBalita,
) (*models.AbsensiKelasIbuBalita, error) {
	if userID == 0 {
		return nil, errors.New("user_id tidak valid")
	}

	ibuID, err := u.repo.FindIbuIDByUserID(userID)
	if err != nil {
		return nil, errors.New("data ibu tidak ditemukan")
	}

	// Hitung nomor pertemuan otomatis dari jumlah data yang sudah ada
	existing, err := u.repo.FindByIbuID(ibuID)
	if err != nil {
		return nil, err
	}

	data := &models.AbsensiKelasIbuBalita{
		IbuID:        ibuID,
		PertemuanKe:  int32(len(existing) + 1),
		Tanggal:      req.Tanggal,
		NamaKader:    req.NamaKader,
		TanggalParaf: req.TanggalParaf,
	}

	if err := u.repo.Create(data); err != nil {
		return nil, err
	}

	return data, nil
}
