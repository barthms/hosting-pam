package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type PemeriksaanLaboratoriumJiwaUsecase interface {
	Create(p *models.PemeriksaanLaboratoriumJiwa) error
	GetByID(id int32) (*models.PemeriksaanLaboratoriumJiwa, error)
	GetByKehamilanID(kehamilanID int32) ([]models.PemeriksaanLaboratoriumJiwa, error)
	Update(p *models.PemeriksaanLaboratoriumJiwa) error
	Delete(id int32) error
}

type pemeriksaanLaboratoriumJiwaUsecase struct {
	repo *repositories.PemeriksaanLaboratoriumJiwaRepository
}

func NewPemeriksaanLaboratoriumJiwaUsecase(repo *repositories.PemeriksaanLaboratoriumJiwaRepository) PemeriksaanLaboratoriumJiwaUsecase {
	return &pemeriksaanLaboratoriumJiwaUsecase{repo: repo}
}

func (u *pemeriksaanLaboratoriumJiwaUsecase) Create(p *models.PemeriksaanLaboratoriumJiwa) error {
	if p.KehamilanID == 0 {
		return errors.New("kehamilan_id wajib diisi")
	}
	return u.repo.Create(p)
}

func (u *pemeriksaanLaboratoriumJiwaUsecase) GetByID(id int32) (*models.PemeriksaanLaboratoriumJiwa, error) {
	return u.repo.FindByID(id)
}

func (u *pemeriksaanLaboratoriumJiwaUsecase) GetByKehamilanID(kehamilanID int32) ([]models.PemeriksaanLaboratoriumJiwa, error) {
	return u.repo.FindByKehamilanID(kehamilanID)
}

func (u *pemeriksaanLaboratoriumJiwaUsecase) Update(p *models.PemeriksaanLaboratoriumJiwa) error {
	_, err := u.repo.FindByID(p.IDLabJiwa)
	if err != nil {
		return errors.New("data pemeriksaan laboratorium jiwa tidak ditemukan")
	}
	return u.repo.Update(p)
}

func (u *pemeriksaanLaboratoriumJiwaUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}
