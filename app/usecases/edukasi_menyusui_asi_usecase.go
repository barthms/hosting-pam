package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type EdukasiMenyusuiASIUsecase interface {
	Create(data *models.EdukasiMenyusuiASI) error
	GetAll() ([]models.EdukasiMenyusuiASI, error)
	GetByID(id int32) (*models.EdukasiMenyusuiASI, error)
	Update(id int32, data *models.EdukasiMenyusuiASI) error
	Delete(id int32) error
}

type edukasiMenyusuiASIUsecase struct {
	repo repositories.EdukasiMenyusuiASIRepository
}

func NewEdukasiMenyusuiASIUsecase(repo repositories.EdukasiMenyusuiASIRepository) EdukasiMenyusuiASIUsecase {
	return &edukasiMenyusuiASIUsecase{repo}
}

func (u *edukasiMenyusuiASIUsecase) Create(data *models.EdukasiMenyusuiASI) error {
	return u.repo.Create(data)
}

func (u *edukasiMenyusuiASIUsecase) GetAll() ([]models.EdukasiMenyusuiASI, error) {
	return u.repo.FindAll()
}

func (u *edukasiMenyusuiASIUsecase) GetByID(id int32) (*models.EdukasiMenyusuiASI, error) {
	return u.repo.FindByID(id)
}

func (u *edukasiMenyusuiASIUsecase) Update(id int32, data *models.EdukasiMenyusuiASI) error {
	existing, err := u.repo.FindByID(id)
	if err != nil {
		return errors.New("data not found")
	}

	existing.Judul = data.Judul
	existing.GambarURL = data.GambarURL
	existing.Isi = data.Isi
	existing.ManfaatASI = data.ManfaatASI
	existing.Cara = data.Cara
	existing.Masalah = data.Masalah
	existing.Solusi = data.Solusi

	return u.repo.Update(existing)
}

func (u *edukasiMenyusuiASIUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}
