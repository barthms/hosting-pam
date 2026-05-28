package usecases

import (
	"context"
	"errors"

	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type CatatanPelayananNifasUsecase interface {
	Create(p *models.CatatanPelayananNifas) error
	GetByID(id int32) (*models.CatatanPelayananNifas, error)
	GetByKehamilanID(kehamilanID int32) ([]models.CatatanPelayananNifas, error)
	Update(p *models.CatatanPelayananNifas) error
	GetMine(ctx context.Context, userID int32) ([]models.CatatanPelayananNifas, error)
	Delete(id int32) error
}

type catatanPelayananNifasUsecase struct {
	repo *repositories.CatatanPelayananNifasRepository
}

func NewCatatanPelayananNifasUsecase(
	repo *repositories.CatatanPelayananNifasRepository,
) CatatanPelayananNifasUsecase {

	return &catatanPelayananNifasUsecase{
		repo: repo,
	}
}

func (u *catatanPelayananNifasUsecase) Create(
	p *models.CatatanPelayananNifas,
) error {

	if p.KehamilanID == 0 {
		return errors.New("kehamilan_id wajib diisi")
	}

	return u.repo.Create(p)
}

func (u *catatanPelayananNifasUsecase) GetByID(
	id int32,
) (*models.CatatanPelayananNifas, error) {

	return u.repo.FindByID(id)
}

func (u *catatanPelayananNifasUsecase) GetByKehamilanID(
	kehamilanID int32,
) ([]models.CatatanPelayananNifas, error) {

	return u.repo.FindByKehamilanID(kehamilanID)
}

func (u *catatanPelayananNifasUsecase) Update(
	p *models.CatatanPelayananNifas,
) error {

	_, err := u.repo.FindByID(p.IDCatatanNifas)

	if err != nil {
		return errors.New("data catatan pelayanan nifas tidak ditemukan")
	}

	return u.repo.Update(p)
}

func (u *catatanPelayananNifasUsecase) GetMine(
	ctx context.Context,
	userID int32,
) ([]models.CatatanPelayananNifas, error) {

	return u.repo.FindMineByUserID(userID)
}

func (u *catatanPelayananNifasUsecase) Delete(
	id int32,
) error {

	return u.repo.Delete(id)
}