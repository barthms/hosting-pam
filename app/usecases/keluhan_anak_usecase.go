package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

var ErrKeluhanAnakForbidden = errors.New("akses keluhan anak ditolak")

type KeluhanAnakUseCase interface {
	Create(data *models.KeluhanAnak) error
	Update(id uint, data *models.KeluhanAnak) error
	Delete(id uint) error
	GetByAnakID(anakID uint) ([]models.KeluhanAnak, error)
	GetByAnakIDForIbu(anakID uint, userID uint) ([]models.KeluhanAnak, error)
	GetByID(id uint) (*models.KeluhanAnak, error)
	GetByIDForIbu(id uint, userID uint) (*models.KeluhanAnak, error)
}

type keluhanAnakUseCase struct {
	repo repositories.KeluhanAnakRepository
}

func NewKeluhanAnakUseCase(repo repositories.KeluhanAnakRepository) KeluhanAnakUseCase {
	return &keluhanAnakUseCase{repo}
}

func (u *keluhanAnakUseCase) Create(data *models.KeluhanAnak) error {
	return u.repo.Create(data)
}

func (u *keluhanAnakUseCase) Update(id uint, data *models.KeluhanAnak) error {
	existing, err := u.repo.FindByID(id)
	if err != nil {
		return err
	}
	existing.Tanggal = data.Tanggal
	existing.TanggalKembali = data.TanggalKembali
	existing.Keluhan = data.Keluhan
	existing.Tindakan = data.Tindakan
	existing.Pemeriksa = data.Pemeriksa
	return u.repo.Update(existing)
}

func (u *keluhanAnakUseCase) Delete(id uint) error {
	return u.repo.Delete(id)
}

func (u *keluhanAnakUseCase) GetByAnakID(anakID uint) ([]models.KeluhanAnak, error) {
	return u.repo.FindAllByAnakID(anakID)
}

func (u *keluhanAnakUseCase) GetByAnakIDForIbu(anakID uint, userID uint) ([]models.KeluhanAnak, error) {
	ok, err := u.repo.IsAnakMilikIbu(userID, anakID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrKeluhanAnakForbidden
	}

	return u.repo.FindAllByAnakID(anakID)
}

func (u *keluhanAnakUseCase) GetByID(id uint) (*models.KeluhanAnak, error) {
	return u.repo.FindByID(id)
}

func (u *keluhanAnakUseCase) GetByIDForIbu(id uint, userID uint) (*models.KeluhanAnak, error) {
	data, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	ok, err := u.repo.IsAnakMilikIbu(userID, uint(data.AnakID))
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrKeluhanAnakForbidden
	}

	return data, nil
}
