package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type CatatanPelayananTrimester2Usecase interface {
	Create(c *models.CatatanPelayananTrimester2) error
	GetByID(id int32) (*models.CatatanPelayananTrimester2, error)
	GetByKehamilanID(kehamilanID int32) ([]models.CatatanPelayananTrimester2, error)
	Update(c *models.CatatanPelayananTrimester2) error
	Delete(id int32) error
}

type catatanPelayananTrimester2Usecase struct {
	repo *repositories.CatatanPelayananTrimester2Repository
}

func NewCatatanPelayananTrimester2Usecase(repo *repositories.CatatanPelayananTrimester2Repository) CatatanPelayananTrimester2Usecase {
	return &catatanPelayananTrimester2Usecase{repo: repo}
}

func (u *catatanPelayananTrimester2Usecase) Create(c *models.CatatanPelayananTrimester2) error {
	if c.KehamilanID == 0 {
		return errors.New("kehamilan_id wajib diisi")
	}
	return u.repo.Create(c)
}

func (u *catatanPelayananTrimester2Usecase) GetByID(id int32) (*models.CatatanPelayananTrimester2, error) {
	return u.repo.FindByID(id)
}

func (u *catatanPelayananTrimester2Usecase) GetByKehamilanID(kehamilanID int32) ([]models.CatatanPelayananTrimester2, error) {
	return u.repo.FindByKehamilanID(kehamilanID)
}

func (u *catatanPelayananTrimester2Usecase) Update(c *models.CatatanPelayananTrimester2) error {
	_, err := u.repo.FindByID(c.IDCatatanT2)
	if err != nil {
		return errors.New("data catatan pelayanan trimester 2 tidak ditemukan")
	}
	return u.repo.Update(c)
}

func (u *catatanPelayananTrimester2Usecase) Delete(id int32) error {
	return u.repo.Delete(id)
}
