package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type EdukasiIMDUsecase interface {
	Create(data *models.EdukasiIMD) error
	GetAll() ([]models.EdukasiIMD, error)
	GetByID(id int32) (*models.EdukasiIMD, error)
	Update(id int32, data *models.EdukasiIMD) error
	Delete(id int32) error
}

type edukasiIMDUsecase struct {
	repo repositories.EdukasiIMDRepository
}

func NewEdukasiIMDUsecase(repo repositories.EdukasiIMDRepository) EdukasiIMDUsecase {
	return &edukasiIMDUsecase{repo}
}

func (u *edukasiIMDUsecase) Create(data *models.EdukasiIMD) error {
	return u.repo.Create(data)
}

func (u *edukasiIMDUsecase) GetAll() ([]models.EdukasiIMD, error) {
	return u.repo.FindAll()
}

func (u *edukasiIMDUsecase) GetByID(id int32) (*models.EdukasiIMD, error) {
	return u.repo.FindByID(id)
}

func (u *edukasiIMDUsecase) Update(id int32, data *models.EdukasiIMD) error {
	existing, err := u.repo.FindByID(id)
	if err != nil {
		return errors.New("data not found")
	}

	existing.Judul = data.Judul
	existing.GambarURL = data.GambarURL
	existing.Isi = data.Isi
	existing.Manfaat = data.Manfaat
	existing.Langkah = data.Langkah

	return u.repo.Update(existing)
}

func (u *edukasiIMDUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}
