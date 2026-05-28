package usecases

import (
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type KategoriUmurUsecase interface {
	GetAll() ([]models.KategoriUmur, error)
}

type kategoriUmurUsecase struct {
	repo repositories.KategoriUmurRepository
}

func NewKategoriUmurUsecase(repo repositories.KategoriUmurRepository) KategoriUmurUsecase {
	return &kategoriUmurUsecase{repo}
}

func (u *kategoriUmurUsecase) GetAll() ([]models.KategoriUmur, error) {
	return u.repo.FindAll()
}