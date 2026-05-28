package usecases

import (
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type ResepMPASIUsecase interface {
	Create(data *models.ResepMPASI) error
	GetAll() ([]models.ResepMPASI, error)
	GetByID(id int32) (*models.ResepMPASI, error)
	GetByAgeRange(bulanMin, bulanMax int) ([]models.ResepMPASI, error)
	GetByType(tipe string) ([]models.ResepMPASI, error)
	Update(id int32, data *models.ResepMPASI) error
	Delete(id int32) error
}

type resepMPASIUsecase struct {
	repo repositories.ResepMPASIRepository
}

func NewResepMPASIUsecase(repo repositories.ResepMPASIRepository) ResepMPASIUsecase {
	return &resepMPASIUsecase{repo}
}

func (u *resepMPASIUsecase) Create(data *models.ResepMPASI) error {
	return u.repo.Create(data)
}

func (u *resepMPASIUsecase) GetAll() ([]models.ResepMPASI, error) {
	return u.repo.GetAll()
}

func (u *resepMPASIUsecase) GetByID(id int32) (*models.ResepMPASI, error) {
	return u.repo.GetByID(id)
}

func (u *resepMPASIUsecase) GetByAgeRange(bulanMin, bulanMax int) ([]models.ResepMPASI, error) {
	return u.repo.GetByAgeRange(bulanMin, bulanMax)
}

func (u *resepMPASIUsecase) GetByType(tipe string) ([]models.ResepMPASI, error) {
	return u.repo.GetByType(tipe)
}

func (u *resepMPASIUsecase) Update(id int32, data *models.ResepMPASI) error {
	return u.repo.Update(id, data)
}

func (u *resepMPASIUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}
