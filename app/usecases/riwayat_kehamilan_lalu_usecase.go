package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type RiwayatKehamilanLaluUsecase interface {
	Create(rk *models.RiwayatKehamilanLalu) error
	GetByID(id int32) (*models.RiwayatKehamilanLalu, error)
	GetByEvaluasiID(evaluasiID int32) ([]models.RiwayatKehamilanLalu, error)
	Update(rk *models.RiwayatKehamilanLalu) error
	Delete(id int32) error
}

type riwayatKehamilanLaluUsecase struct {
	repo *repositories.RiwayatKehamilanLaluRepository
}

func NewRiwayatKehamilanLaluUsecase(repo *repositories.RiwayatKehamilanLaluRepository) RiwayatKehamilanLaluUsecase {
	return &riwayatKehamilanLaluUsecase{repo: repo}
}

func (u *riwayatKehamilanLaluUsecase) Create(rk *models.RiwayatKehamilanLalu) error {
	if rk.IDEvaluasi == 0 {
		return errors.New("id_evaluasi wajib diisi")
	}
	return u.repo.Create(rk)
}

func (u *riwayatKehamilanLaluUsecase) GetByID(id int32) (*models.RiwayatKehamilanLalu, error) {
	return u.repo.FindByID(id)
}

func (u *riwayatKehamilanLaluUsecase) GetByEvaluasiID(evaluasiID int32) ([]models.RiwayatKehamilanLalu, error) {
	return u.repo.FindByEvaluasiID(evaluasiID)
}

func (u *riwayatKehamilanLaluUsecase) Update(rk *models.RiwayatKehamilanLalu) error {
	_, err := u.repo.FindByID(rk.IDRiwayat)
	if err != nil {
		return errors.New("data riwayat kehamilan lalu tidak ditemukan")
	}
	return u.repo.Update(rk)
}

func (u *riwayatKehamilanLaluUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}
