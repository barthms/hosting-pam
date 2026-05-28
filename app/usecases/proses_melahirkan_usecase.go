package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type ProsesMelahirkanUsecase interface {
	GetMine(userID int32) (*models.ProsesMelahirkan, error)
	SaveMine(userID int32, req models.ProsesMelahirkan) (*models.ProsesMelahirkan, error)
}

type prosesMelahirkanUsecase struct {
	repo *repositories.ProsesMelahirkanRepository
}

func NewProsesMelahirkanUsecase(
	repo *repositories.ProsesMelahirkanRepository,
) ProsesMelahirkanUsecase {
	return &prosesMelahirkanUsecase{repo: repo}
}

func (u *prosesMelahirkanUsecase) GetMine(userID int32) (*models.ProsesMelahirkan, error) {
	if userID == 0 {
		return nil, errors.New("user_id tidak valid")
	}

	kehamilan, err := u.repo.FindActiveKehamilanByUserID(userID)
	if err != nil {
		return nil, errors.New("kehamilan aktif tidak ditemukan")
	}

	return u.repo.FindByKehamilanID(kehamilan.ID)
}

func (u *prosesMelahirkanUsecase) SaveMine(
	userID int32,
	req models.ProsesMelahirkan,
) (*models.ProsesMelahirkan, error) {
	if userID == 0 {
		return nil, errors.New("user_id tidak valid")
	}

	kehamilan, err := u.repo.FindActiveKehamilanByUserID(userID)
	if err != nil {
		return nil, errors.New("kehamilan aktif tidak ditemukan")
	}

	data := &models.ProsesMelahirkan{
		KehamilanID:            kehamilan.ID,
		TandaPersalinan:        req.TandaPersalinan,
		ProsesMelahirkan:       req.ProsesMelahirkan,
		HakIbuPendamping:       req.HakIbuPendamping,
		HakIbuPosisiMelahirkan: req.HakIbuPosisiMelahirkan,
		Mulas:                  req.Mulas,
		TeknikMengurangiNyeri:  req.TeknikMengurangiNyeri,
		IMDKontakKulit:         req.IMDKontakKulit,
	}

	if err := u.repo.Upsert(data); err != nil {
		return nil, err
	}

	return data, nil
}