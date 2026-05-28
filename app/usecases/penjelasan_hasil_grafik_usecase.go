package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type PenjelasanHasilGrafikUsecase interface {
	Create(p *models.PenjelasanHasilGrafik) error
	GetByID(id int32) (*models.PenjelasanHasilGrafik, error)
	GetByKehamilanID(kehamilanID int32) ([]models.PenjelasanHasilGrafik, error)
	Update(p *models.PenjelasanHasilGrafik) error
	Delete(id int32) error
}

type penjelasanHasilGrafikUsecase struct {
	repo *repositories.PenjelasanHasilGrafikRepository
}

func NewPenjelasanHasilGrafikUsecase(repo *repositories.PenjelasanHasilGrafikRepository) PenjelasanHasilGrafikUsecase {
	return &penjelasanHasilGrafikUsecase{repo: repo}
}

func (u *penjelasanHasilGrafikUsecase) Create(p *models.PenjelasanHasilGrafik) error {
	if p.KehamilanID == 0 {
		return errors.New("kehamilan_id wajib diisi")
	}
	return u.repo.Create(p)
}

func (u *penjelasanHasilGrafikUsecase) GetByID(id int32) (*models.PenjelasanHasilGrafik, error) {
	return u.repo.FindByID(id)
}

func (u *penjelasanHasilGrafikUsecase) GetByKehamilanID(kehamilanID int32) ([]models.PenjelasanHasilGrafik, error) {
	return u.repo.FindByKehamilanID(kehamilanID)
}

func (u *penjelasanHasilGrafikUsecase) Update(p *models.PenjelasanHasilGrafik) error {
	_, err := u.repo.FindByID(p.IDPenjelasan)
	if err != nil {
		return errors.New("data penjelasan hasil grafik tidak ditemukan")
	}
	return u.repo.Update(p)
}

func (u *penjelasanHasilGrafikUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}
