package usecases

import (
	"time"

	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type JadwalLayananUsecase interface {
	Create(data *models.JadwalLayanan) error
	GetAll() ([]models.JadwalLayanan, error)
	GetByID(id int32) (*models.JadwalLayanan, error)
	GetByPosyandu(posyanduID int32) ([]models.JadwalLayanan, error)
	GetByDateRange(posyanduID *int32, from, to *time.Time) ([]models.JadwalLayanan, error)
	GetUpcoming(limit int) ([]models.JadwalLayanan, error)
	Update(id int32, data *models.JadwalLayanan) error
	Delete(id int32) error
}

type jadwalLayananUsecase struct {
	repo repositories.JadwalLayananRepository
}

func NewJadwalLayananUsecase(r repositories.JadwalLayananRepository) JadwalLayananUsecase {
	return &jadwalLayananUsecase{repo: r}
}

func (u *jadwalLayananUsecase) Create(data *models.JadwalLayanan) error {
	return u.repo.Create(data)
}

func (u *jadwalLayananUsecase) GetAll() ([]models.JadwalLayanan, error) {
	return u.repo.GetAll()
}

func (u *jadwalLayananUsecase) GetByID(id int32) (*models.JadwalLayanan, error) {
	return u.repo.GetByID(id)
}

func (u *jadwalLayananUsecase) GetByPosyandu(posyanduID int32) ([]models.JadwalLayanan, error) {
	return u.repo.GetByPosyandu(posyanduID)
}

func (u *jadwalLayananUsecase) GetByDateRange(posyanduID *int32, from, to *time.Time) ([]models.JadwalLayanan, error) {
	return u.repo.GetByDateRange(posyanduID, from, to)
}

func (u *jadwalLayananUsecase) GetUpcoming(limit int) ([]models.JadwalLayanan, error) {
	return u.repo.GetUpcoming(limit)
}

func (u *jadwalLayananUsecase) Update(id int32, data *models.JadwalLayanan) error {
	return u.repo.Update(id, data)
}

func (u *jadwalLayananUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}
