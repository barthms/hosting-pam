package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type EdukasiTrimesterUsecase interface {
	Create(data *models.EdukasiTrimester) error
	GetAll() ([]models.EdukasiTrimester, error)
	GetByID(id int32) (*models.EdukasiTrimester, error)
	Update(id int32, data *models.EdukasiTrimester) error
	Delete(id int32) error
}

type edukasiTrimesterUsecase struct {
	repo repositories.EdukasiTrimesterRepository
}

func NewEdukasiTrimesterUsecase(repo repositories.EdukasiTrimesterRepository) EdukasiTrimesterUsecase {
	return &edukasiTrimesterUsecase{repo}
}

func (u *edukasiTrimesterUsecase) Create(data *models.EdukasiTrimester) error {
	return u.repo.Create(data)
}

func (u *edukasiTrimesterUsecase) GetAll() ([]models.EdukasiTrimester, error) {
	return u.repo.FindAll()
}

func (u *edukasiTrimesterUsecase) GetByID(id int32) (*models.EdukasiTrimester, error) {
	return u.repo.FindByID(id)
}

func (u *edukasiTrimesterUsecase) Update(id int32, data *models.EdukasiTrimester) error {
	existing, err := u.repo.FindByID(id)
	if err != nil {
		return errors.New("data not found")
	}
	existing.Judul = data.Judul
	existing.GambarURL = data.GambarURL
	existing.Isi = data.Isi
	return u.repo.Update(existing)
}

func (u *edukasiTrimesterUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}
