package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type EdukasiSetelahMelahirkanUsecase interface {
	Create(data *models.EdukasiSetelahMelahirkan) error
	GetAll() ([]models.EdukasiSetelahMelahirkan, error)
	GetByID(id int32) (*models.EdukasiSetelahMelahirkan, error)
	Update(id int32, data *models.EdukasiSetelahMelahirkan) error
	Delete(id int32) error
}

type edukasiSetelahMelahirkanUsecase struct {
	repo repositories.EdukasiSetelahMelahirkanRepository
}

func NewEdukasiSetelahMelahirkanUsecase(repo repositories.EdukasiSetelahMelahirkanRepository) EdukasiSetelahMelahirkanUsecase {
	return &edukasiSetelahMelahirkanUsecase{repo}
}

func (u *edukasiSetelahMelahirkanUsecase) Create(data *models.EdukasiSetelahMelahirkan) error {
	return u.repo.Create(data)
}

func (u *edukasiSetelahMelahirkanUsecase) GetAll() ([]models.EdukasiSetelahMelahirkan, error) {
	return u.repo.FindAll()
}

func (u *edukasiSetelahMelahirkanUsecase) GetByID(id int32) (*models.EdukasiSetelahMelahirkan, error) {
	return u.repo.FindByID(id)
}

func (u *edukasiSetelahMelahirkanUsecase) Update(id int32, data *models.EdukasiSetelahMelahirkan) error {
	existing, err := u.repo.FindByID(id)
	if err != nil {
		return errors.New("data not found")
	}

	existing.Judul = data.Judul
	existing.GambarURL = data.GambarURL
	existing.Isi = data.Isi

	return u.repo.Update(existing)
}

func (u *edukasiSetelahMelahirkanUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}
