package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type PersiapanMelahirkanUsecase interface {
	GetMine(userID int32) (*models.PersiapanMelahirkan, error)
	SaveMine(userID int32, req models.PersiapanMelahirkan) (*models.PersiapanMelahirkan, error)
}

type persiapanMelahirkanUsecase struct {
	repo *repositories.PersiapanMelahirkanRepository
}

func NewPersiapanMelahirkanUsecase(
	repo *repositories.PersiapanMelahirkanRepository,
) PersiapanMelahirkanUsecase {
	return &persiapanMelahirkanUsecase{repo: repo}
}

func (u *persiapanMelahirkanUsecase) GetMine(userID int32) (*models.PersiapanMelahirkan, error) {
	if userID == 0 {
		return nil, errors.New("user_id tidak valid")
	}

	kehamilan, err := u.repo.FindActiveKehamilanByUserID(userID)
	if err != nil {
		return nil, errors.New("kehamilan aktif tidak ditemukan")
	}

	return u.repo.FindByKehamilanID(kehamilan.ID)
}

func (u *persiapanMelahirkanUsecase) SaveMine(
	userID int32,
	req models.PersiapanMelahirkan,
) (*models.PersiapanMelahirkan, error) {
	if userID == 0 {
		return nil, errors.New("user_id tidak valid")
	}

	kehamilan, err := u.repo.FindActiveKehamilanByUserID(userID)
	if err != nil {
		return nil, errors.New("kehamilan aktif tidak ditemukan")
	}

	data := &models.PersiapanMelahirkan{
		KehamilanID:          kehamilan.ID,
		PerkiraanPersalinan:  req.PerkiraanPersalinan,
		PendampingPersalinan: req.PendampingPersalinan,
		DanaPersalinan:       req.DanaPersalinan,
		StatusJKN:            req.StatusJKN,
		FaskesPersalinan:     req.FaskesPersalinan,
		PendonorDarah:        req.PendonorDarah,
		Transportasi:         req.Transportasi,
		MetodeKB:             req.MetodeKB,
		ProgramP4K:           req.ProgramP4K,
		DokumenPenting:       req.DokumenPenting,
	}

	if err := u.repo.Upsert(data); err != nil {
		return nil, err
	}

	return data, nil
}