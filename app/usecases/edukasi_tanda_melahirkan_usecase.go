package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type EdukasiTandaMelahirkanUsecase interface {
	Create(data *models.EdukasiTandaMelahirkan) error
	GetAll() ([]models.EdukasiTandaMelahirkan, error)
	GetByID(id int32) (*models.EdukasiTandaMelahirkan, error)
	Update(id int32, data *models.EdukasiTandaMelahirkan) error
	Delete(id int32) error
}

type edukasiTandaMelahirkanUsecase struct {
	repo repositories.EdukasiTandaMelahirkanRepository
}

func NewEdukasiTandaMelahirkanUsecase(repo repositories.EdukasiTandaMelahirkanRepository) EdukasiTandaMelahirkanUsecase {
	return &edukasiTandaMelahirkanUsecase{repo}
}

func (u *edukasiTandaMelahirkanUsecase) Create(data *models.EdukasiTandaMelahirkan) error {
	return u.repo.Create(data)
}

func (u *edukasiTandaMelahirkanUsecase) GetAll() ([]models.EdukasiTandaMelahirkan, error) {
	return u.repo.FindAll()
}

func (u *edukasiTandaMelahirkanUsecase) GetByID(id int32) (*models.EdukasiTandaMelahirkan, error) {
	return u.repo.FindByID(id)
}

func (u *edukasiTandaMelahirkanUsecase) Update(id int32, data *models.EdukasiTandaMelahirkan) error {
	existing, err := u.repo.FindByID(id)
	if err != nil {
		return errors.New("data not found")
	}

	existing.Judul = data.Judul
	existing.GambarURL = data.GambarURL
	existing.Isi = data.Isi
	existing.Tanda = data.Tanda
	existing.Tindakan = data.Tindakan

	return u.repo.Update(existing)
}

func (u *edukasiTandaMelahirkanUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}
