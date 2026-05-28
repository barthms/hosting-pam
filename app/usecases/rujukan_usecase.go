package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type RujukanUsecase interface {
	Create(rj *models.Rujukan) error
	GetByID(id int32) (*models.Rujukan, error)
	GetByIDForOrangtua(id int32, userID int32) (*models.Rujukan, error) // MODUL IBU
	GetByKehamilanID(kehamilanID int32) ([]models.Rujukan, error)
	Update(rj *models.Rujukan) error
	Delete(id int32) error
}

type rujukanUsecase struct {
	repo *repositories.RujukanRepository
}

func NewRujukanUsecase(repo *repositories.RujukanRepository) RujukanUsecase {
	return &rujukanUsecase{repo: repo}
}

func (u *rujukanUsecase) Create(rj *models.Rujukan) error {
	if rj.KehamilanID == 0 {
		return errors.New("kehamilan_id wajib diisi")
	}
	return u.repo.Create(rj)
}

func (u *rujukanUsecase) GetByID(id int32) (*models.Rujukan, error) {
	return u.repo.FindByID(id)
}

func (u *rujukanUsecase) GetByKehamilanID(kehamilanID int32) ([]models.Rujukan, error) {
	return u.repo.FindByKehamilanID(kehamilanID)
}

func (u *rujukanUsecase) Update(rj *models.Rujukan) error {
	_, err := u.repo.FindByID(rj.ID)
	if err != nil {
		return errors.New("data rujukan tidak ditemukan")
	}
	return u.repo.Update(rj)
}

func (u *rujukanUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}


// MODUL IBU
func (u *rujukanUsecase) GetByIDForOrangtua(id int32, userID int32) (*models.Rujukan, error) {
	allowed, err := u.repo.IsOwnedByUser(id, userID)
	if err != nil {
		return nil, err
	}

	if !allowed {
		return nil, errors.New("Anda tidak dapat mengakses data ini")
	}

	return u.repo.FindByID(id)
}