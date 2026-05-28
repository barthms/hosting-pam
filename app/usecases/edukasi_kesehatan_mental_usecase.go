package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type EdukasiKesehatanMentalUsecase interface {
	Create(data *models.EdukasiKesehatanMental) error
	GetAll() ([]models.EdukasiKesehatanMental, error)
	GetByID(id int32) (*models.EdukasiKesehatanMental, error)
	Update(id int32, data *models.EdukasiKesehatanMental) error
	Delete(id int32) error
}

type edukasiKesehatanMentalUsecase struct {
	repo repositories.EdukasiKesehatanMentalRepository
}

func NewEdukasiKesehatanMentalUsecase(repo repositories.EdukasiKesehatanMentalRepository) EdukasiKesehatanMentalUsecase {
	return &edukasiKesehatanMentalUsecase{repo}
}

func (u *edukasiKesehatanMentalUsecase) Create(data *models.EdukasiKesehatanMental) error {
	return u.repo.Create(data)
}

func (u *edukasiKesehatanMentalUsecase) GetAll() ([]models.EdukasiKesehatanMental, error) {
	return u.repo.FindAll()
}

func (u *edukasiKesehatanMentalUsecase) GetByID(id int32) (*models.EdukasiKesehatanMental, error) {
	return u.repo.FindByID(id)
}

func (u *edukasiKesehatanMentalUsecase) Update(id int32, data *models.EdukasiKesehatanMental) error {
	existing, err := u.repo.FindByID(id)
	if err != nil {
		return errors.New("data not found")
	}

	existing.Judul = data.Judul
	existing.GambarURL = data.GambarURL
	existing.Isi = data.Isi
	existing.TandaGejala = data.TandaGejala
	existing.Solusi = data.Solusi

	return u.repo.Update(existing)
}

func (u *edukasiKesehatanMentalUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}
