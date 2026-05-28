package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type EdukasiNifasUsecase interface {
	Create(data *models.EdukasiNifas) error
	GetAll() ([]models.EdukasiNifas, error)
	GetByID(id int32) (*models.EdukasiNifas, error)
	Update(id int32, data *models.EdukasiNifas) error
	Delete(id int32) error
}

type edukasiNifasUsecase struct {
	repo repositories.EdukasiNifasRepository
}

func NewEdukasiNifasUsecase(repo repositories.EdukasiNifasRepository) EdukasiNifasUsecase {
	return &edukasiNifasUsecase{repo}
}

func (u *edukasiNifasUsecase) Create(data *models.EdukasiNifas) error {
	return u.repo.Create(data)
}

func (u *edukasiNifasUsecase) GetAll() ([]models.EdukasiNifas, error) {
	return u.repo.FindAll()
}

func (u *edukasiNifasUsecase) GetByID(id int32) (*models.EdukasiNifas, error) {
	return u.repo.FindByID(id)
}

func (u *edukasiNifasUsecase) Update(id int32, data *models.EdukasiNifas) error {
	existing, err := u.repo.FindByID(id)
	if err != nil {
		return errors.New("data not found")
	}

	existing.Judul = data.Judul
	existing.GambarURL = data.GambarURL
	existing.Isi = data.Isi
	existing.Perawatan = data.Perawatan
	existing.TandaBahaya = data.TandaBahaya

	return u.repo.Update(existing)
}

func (u *edukasiNifasUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}
