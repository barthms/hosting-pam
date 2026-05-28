package usecases

import (
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type JadwalHarianMPASIUsecase interface {
	Create(data *models.JadwalHarianMPASI) error
	GetAll() ([]models.JadwalHarianMPASI, error)
	GetByID(id int32) (*models.JadwalHarianMPASI, error)
	GetByAgeRange(bulanMin, bulanMax int) ([]models.JadwalHarianMPASI, error)
	Update(id int32, data *models.JadwalHarianMPASI) error
	Delete(id int32) error
}

type jadwalHarianMPASIUsecase struct {
	repo repositories.JadwalHarianMPASIRepository
}

func NewJadwalHarianMPASIUsecase(repo repositories.JadwalHarianMPASIRepository) JadwalHarianMPASIUsecase {
	return &jadwalHarianMPASIUsecase{repo}
}

func (u *jadwalHarianMPASIUsecase) Create(data *models.JadwalHarianMPASI) error {
	return u.repo.Create(data)
}

func (u *jadwalHarianMPASIUsecase) GetAll() ([]models.JadwalHarianMPASI, error) {
	return u.repo.GetAll()
}

func (u *jadwalHarianMPASIUsecase) GetByID(id int32) (*models.JadwalHarianMPASI, error) {
	return u.repo.GetByID(id)
}

func (u *jadwalHarianMPASIUsecase) GetByAgeRange(bulanMin, bulanMax int) ([]models.JadwalHarianMPASI, error) {
	return u.repo.GetByAgeRange(bulanMin, bulanMax)
}

func (u *jadwalHarianMPASIUsecase) Update(id int32, data *models.JadwalHarianMPASI) error {
	return u.repo.Update(id, data)
}

func (u *jadwalHarianMPASIUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}
