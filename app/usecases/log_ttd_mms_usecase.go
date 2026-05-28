package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type LogTTDMMSUsecase interface {
	GetMine(userID int32) ([]models.LogTTDMMS, error)
	SaveMine(userID int32, bulanKe int32, hariKe int32, sudahDiminum bool) (*models.LogTTDMMS, error)
}

type logTTDMMSUsecase struct {
	repo *repositories.LogTTDMMSRepository
}

func NewLogTTDMMSUsecase(repo *repositories.LogTTDMMSRepository) LogTTDMMSUsecase {
	return &logTTDMMSUsecase{repo: repo}
}

func (u *logTTDMMSUsecase) GetMine(userID int32) ([]models.LogTTDMMS, error) {
	if userID == 0 {
		return nil, errors.New("user_id tidak valid")
	}

	kehamilan, err := u.repo.FindActiveKehamilanByUserID(userID)
	if err != nil {
		return nil, errors.New("kehamilan aktif tidak ditemukan")
	}

	return u.repo.FindByKehamilanID(kehamilan.ID)
}

func (u *logTTDMMSUsecase) SaveMine(
	userID int32,
	bulanKe int32,
	hariKe int32,
	sudahDiminum bool,
) (*models.LogTTDMMS, error) {
	if userID == 0 {
		return nil, errors.New("user_id tidak valid")
	}
	if bulanKe < 1 || bulanKe > 10 {
		return nil, errors.New("bulan_ke harus antara 1 sampai 10")
	}
	if hariKe < 1 || hariKe > 31 {
		return nil, errors.New("hari_ke harus antara 1 sampai 31")
	}

	kehamilan, err := u.repo.FindActiveKehamilanByUserID(userID)
	if err != nil {
		return nil, errors.New("kehamilan aktif tidak ditemukan")
	}

	log := &models.LogTTDMMS{
		KehamilanID:   kehamilan.ID,
		BulanKe:       bulanKe,
		HariKe:        hariKe,
		SudahDiminum:  sudahDiminum,
	}

	if err := u.repo.Upsert(log); err != nil {
		return nil, err
	}

	return log, nil
}