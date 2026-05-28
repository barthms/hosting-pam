package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type EdukasiPolaAsuhUsecase interface {
	Create(data *models.EdukasiPolaAsuh) error
	GetAll() ([]models.EdukasiPolaAsuh, error)
	GetByID(id int32) (*models.EdukasiPolaAsuh, error)
	Update(id int32, data *models.EdukasiPolaAsuh) error
	Delete(id int32) error
}

type edukasiPolaAsuhUsecase struct {
	repo repositories.EdukasiPolaAsuhRepository
}

func NewEdukasiPolaAsuhUsecase(repo repositories.EdukasiPolaAsuhRepository) EdukasiPolaAsuhUsecase {
	return &edukasiPolaAsuhUsecase{repo}
}

func (u *edukasiPolaAsuhUsecase) Create(data *models.EdukasiPolaAsuh) error {
	return u.repo.Create(data)
}

func (u *edukasiPolaAsuhUsecase) GetAll() ([]models.EdukasiPolaAsuh, error) {
	return u.repo.FindAll()
}

func (u *edukasiPolaAsuhUsecase) GetByID(id int32) (*models.EdukasiPolaAsuh, error) {
	return u.repo.FindByID(id)
}

func (u *edukasiPolaAsuhUsecase) Update(id int32, data *models.EdukasiPolaAsuh) error {
	existing, err := u.repo.FindByID(id)
	if err != nil {
		return errors.New("data not found")
	}

	existing.Judul = data.Judul
	existing.GambarURL = data.GambarURL
	existing.Isi = data.Isi

	return u.repo.Update(existing)
}

func (u *edukasiPolaAsuhUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}
