package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type PemantauanIbuHamilUsecase interface {
	GetMine(userID int32) ([]models.PemantauanIbuHamil, error)
	SaveMine(userID int32, req models.PemantauanIbuHamil) (*models.PemantauanIbuHamil, error)
}

type pemantauanIbuHamilUsecase struct {
	repo *repositories.PemantauanIbuHamilRepository
}

func NewPemantauanIbuHamilUsecase(repo *repositories.PemantauanIbuHamilRepository) PemantauanIbuHamilUsecase {
	return &pemantauanIbuHamilUsecase{repo: repo}
}

func (u *pemantauanIbuHamilUsecase) GetMine(userID int32) ([]models.PemantauanIbuHamil, error) {
	if userID == 0 {
		return nil, errors.New("user_id tidak valid")
	}

	kehamilan, err := u.repo.FindActiveKehamilanByUserID(userID)
	if err != nil {
		return nil, errors.New("kehamilan aktif tidak ditemukan")
	}

	return u.repo.FindByKehamilanID(kehamilan.ID)
}

func (u *pemantauanIbuHamilUsecase) SaveMine(
	userID int32,
	req models.PemantauanIbuHamil,
) (*models.PemantauanIbuHamil, error) {
	if userID == 0 {
		return nil, errors.New("user_id tidak valid")
	}

	if req.MingguKehamilan < 4 || req.MingguKehamilan > 42 {
		return nil, errors.New("minggu_kehamilan harus antara 4 sampai 42")
	}

	kehamilan, err := u.repo.FindActiveKehamilanByUserID(userID)
	if err != nil {
		return nil, errors.New("kehamilan aktif tidak ditemukan")
	}

	data := &models.PemantauanIbuHamil{
		KehamilanID:       kehamilan.ID,
		MingguKehamilan:   req.MingguKehamilan,
		DemamLebih2Hari:   req.DemamLebih2Hari,
		SakitKepala:       req.SakitKepala,
		CemasBerlebih:     req.CemasBerlebih,
		ResikoTB:          req.ResikoTB,
		GerakanBayiKurang: req.GerakanBayiKurang,
		NyeriPerut:        req.NyeriPerut,
		CairanJalanLahir:  req.CairanJalanLahir,
		MasalahKemaluan:   req.MasalahKemaluan,
		DiareBerulang:     req.DiareBerulang,
	}

	if err := u.repo.Upsert(data); err != nil {
		return nil, err
	}

	return data, nil
}