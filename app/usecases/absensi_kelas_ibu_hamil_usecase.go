package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type AbsensiKelasIbuHamilUsecase interface {
	GetMine(userID int32) ([]models.AbsensiKelasIbuHamil, error)
	SaveMine(userID int32, req models.AbsensiKelasIbuHamil) (*models.AbsensiKelasIbuHamil, error)
}

type absensiKelasIbuHamilUsecase struct {
	repo *repositories.AbsensiKelasIbuHamilRepository
}

func NewAbsensiKelasIbuHamilUsecase(
	repo *repositories.AbsensiKelasIbuHamilRepository,
) AbsensiKelasIbuHamilUsecase {
	return &absensiKelasIbuHamilUsecase{repo: repo}
}

func (u *absensiKelasIbuHamilUsecase) GetMine(
	userID int32,
) ([]models.AbsensiKelasIbuHamil, error) {
	if userID == 0 {
		return nil, errors.New("user_id tidak valid")
	}

	kehamilan, err := u.repo.FindActiveKehamilanByUserID(userID)
	if err != nil {
		return nil, errors.New("kehamilan aktif tidak ditemukan")
	}

	return u.repo.FindByKehamilanID(kehamilan.ID)
}

func (u *absensiKelasIbuHamilUsecase) SaveMine(
	userID int32,
	req models.AbsensiKelasIbuHamil,
) (*models.AbsensiKelasIbuHamil, error) {
	if userID == 0 {
		return nil, errors.New("user_id tidak valid")
	}

	if req.PertemuanKe < 1 || req.PertemuanKe > 9 {
		return nil, errors.New("pertemuan_ke harus antara 1 sampai 9")
	}

	kehamilan, err := u.repo.FindActiveKehamilanByUserID(userID)
	if err != nil {
		return nil, errors.New("kehamilan aktif tidak ditemukan")
	}

	data := &models.AbsensiKelasIbuHamil{
		KehamilanID:  kehamilan.ID,
		PertemuanKe:  req.PertemuanKe,
		Tanggal:      req.Tanggal,
		NamaKader:    req.NamaKader,
		TanggalParaf: req.TanggalParaf,
	}

	if err := u.repo.Upsert(data); err != nil {
		return nil, err
	}

	return data, nil
}
