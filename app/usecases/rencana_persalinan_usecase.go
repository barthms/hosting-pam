package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type RencanaPersalinanUsecase interface {
	Create(rp *models.RencanaPersalinan) error
	GetByID(id int32) (*models.RencanaPersalinan, error)
	GetByKehamilanID(kehamilanID int32) ([]models.RencanaPersalinan, error)
	Update(rp *models.RencanaPersalinan) error
	Delete(id int32) error
	// Tambahan untuk kebutuhan controller
	FindByIDWithoutPreload(id int32) (*models.RencanaPersalinan, error)
}

type rencanaPersalinanUsecase struct {
	repo *repositories.RencanaPersalinanRepository
}

func NewRencanaPersalinanUsecase(repo *repositories.RencanaPersalinanRepository) RencanaPersalinanUsecase {
	return &rencanaPersalinanUsecase{repo: repo}
}

func (u *rencanaPersalinanUsecase) Create(rp *models.RencanaPersalinan) error {
	if rp.KehamilanID == 0 {
		return errors.New("kehamilan_id wajib diisi")
	}
	return u.repo.Create(rp)
}

func (u *rencanaPersalinanUsecase) GetByID(id int32) (*models.RencanaPersalinan, error) {
	return u.repo.FindByID(id)
}

func (u *rencanaPersalinanUsecase) GetByKehamilanID(kehamilanID int32) ([]models.RencanaPersalinan, error) {
	return u.repo.FindByKehamilanID(kehamilanID)
}

func (u *rencanaPersalinanUsecase) Update(rp *models.RencanaPersalinan) error {
	_, err := u.repo.FindByIDWithoutPreload(rp.IDRencanaPersalinan)
	if err != nil {
		return errors.New("data rencana persalinan tidak ditemukan")
	}
	return u.repo.Update(rp)
}

func (u *rencanaPersalinanUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}

// FindByIDWithoutPreload implementasi untuk usecase
func (u *rencanaPersalinanUsecase) FindByIDWithoutPreload(id int32) (*models.RencanaPersalinan, error) {
	return u.repo.FindByIDWithoutPreload(id)
}
