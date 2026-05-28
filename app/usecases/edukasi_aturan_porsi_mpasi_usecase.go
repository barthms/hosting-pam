package usecases

import (
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type AturanPorsiMPASIUsecase interface {
	Create(data *models.AturanPorsiMPASI) error
	GetAll() ([]models.AturanPorsiMPASI, error)
	GetByID(id int32) (*models.AturanPorsiMPASI, error)
	GetByAgeRange(bulanMin, bulanMax int) ([]models.AturanPorsiMPASI, error)
	Update(id int32, data *models.AturanPorsiMPASI) error
	Delete(id int32) error
}

type aturanPorsiMPASIUsecase struct {
	repo repositories.AturanPorsiMPASIRepository
}

func NewAturanPorsiMPASIUsecase(repo repositories.AturanPorsiMPASIRepository) AturanPorsiMPASIUsecase {
	return &aturanPorsiMPASIUsecase{repo}
}

func (u *aturanPorsiMPASIUsecase) Create(data *models.AturanPorsiMPASI) error {
	return u.repo.Create(data)
}

func (u *aturanPorsiMPASIUsecase) GetAll() ([]models.AturanPorsiMPASI, error) {
	return u.repo.GetAll()
}

func (u *aturanPorsiMPASIUsecase) GetByID(id int32) (*models.AturanPorsiMPASI, error) {
	return u.repo.GetByID(id)
}

func (u *aturanPorsiMPASIUsecase) GetByAgeRange(bulanMin, bulanMax int) ([]models.AturanPorsiMPASI, error) {
	return u.repo.GetByAgeRange(bulanMin, bulanMax)
}

func (u *aturanPorsiMPASIUsecase) Update(id int32, data *models.AturanPorsiMPASI) error {
	return u.repo.Update(id, data)
}

func (u *aturanPorsiMPASIUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}
