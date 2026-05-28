package usecases

import (
	"context"
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type RingkasanPelayananPersalinanUsecase interface {
	Create(rp *models.RingkasanPelayananPersalinan) error
	GetByID(id int32) (*models.RingkasanPelayananPersalinan, error)
	GetByKehamilanID(kehamilanID int32) ([]models.RingkasanPelayananPersalinan, error)
	Update(rp *models.RingkasanPelayananPersalinan) error
	Delete(id int32) error
	GetMine(ctx context.Context,userID int32,) ([]models.RingkasanPelayananPersalinan, error)
}

type RingkasanPelayananPersalinan struct {
	repo *repositories.RingkasanPelayananPersalinanRepository
}

func NewRingkasanPelayananPersalinanUsecase(repo *repositories.RingkasanPelayananPersalinanRepository) RingkasanPelayananPersalinanUsecase {
	return &RingkasanPelayananPersalinan{repo: repo}
}

func (u *RingkasanPelayananPersalinan) Create(rp *models.RingkasanPelayananPersalinan) error {
	if rp.KehamilanID == 0 {
		return errors.New("kehamilan_id wajib diisi")
	}
	return u.repo.Create(rp)
}

func (u *RingkasanPelayananPersalinan) GetByID(id int32) (*models.RingkasanPelayananPersalinan, error) {
	return u.repo.FindByID(id)
}

func (u *RingkasanPelayananPersalinan) GetByKehamilanID(kehamilanID int32) ([]models.RingkasanPelayananPersalinan, error) {
	return u.repo.FindByKehamilanID(kehamilanID)
}

func (u *RingkasanPelayananPersalinan) Update(rp *models.RingkasanPelayananPersalinan) error {
	_, err := u.repo.FindByID(rp.ID)
	if err != nil {
		return errors.New("data ringkasan pelayanan persalinan tidak ditemukan")
	}
	return u.repo.Update(rp)
}

func (u *RingkasanPelayananPersalinan) GetMine(
	ctx context.Context,
	userID int32,
) ([]models.RingkasanPelayananPersalinan, error) {

	return u.repo.GetMine(
		userID,
	)
}

func (u *RingkasanPelayananPersalinan) Delete(id int32) error {
	return u.repo.Delete(id)
}
