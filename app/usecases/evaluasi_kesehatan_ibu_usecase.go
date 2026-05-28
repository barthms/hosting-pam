package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type EvaluasiKesehatanIbuUsecase interface {
	Create(e *models.EvaluasiKesehatanIbu) error
	GetByID(id int32) (*models.EvaluasiKesehatanIbu, error)
	GetByIDForOrangtua(id int32, userID int32) (*models.EvaluasiKesehatanIbu, error) // MODUL IBU
	GetMine(userID int32) (*models.EvaluasiKesehatanIbu, error) // MODUL IBU
	GetByKehamilanID(kehamilanID int32) ([]models.EvaluasiKesehatanIbu, error)
	Update(e *models.EvaluasiKesehatanIbu) error
	Delete(id int32) error
}

type evaluasiKesehatanIbuUsecase struct {
	repo *repositories.EvaluasiKesehatanIbuRepository
}

func NewEvaluasiKesehatanIbuUsecase(repo *repositories.EvaluasiKesehatanIbuRepository) EvaluasiKesehatanIbuUsecase {
	return &evaluasiKesehatanIbuUsecase{repo: repo}
}

func (u *evaluasiKesehatanIbuUsecase) Create(e *models.EvaluasiKesehatanIbu) error {
	if e.KehamilanID == 0 {
		return errors.New("kehamilan_id wajib diisi")
	}
	return u.repo.Create(e)
}

func (u *evaluasiKesehatanIbuUsecase) GetByID(id int32) (*models.EvaluasiKesehatanIbu, error) {
	return u.repo.FindByID(id)	
}

func (u *evaluasiKesehatanIbuUsecase) GetByKehamilanID(kehamilanID int32) ([]models.EvaluasiKesehatanIbu, error) {
	return u.repo.FindByKehamilanID(kehamilanID)
}

func (u *evaluasiKesehatanIbuUsecase) Update(e *models.EvaluasiKesehatanIbu) error {
	_, err := u.repo.FindByID(e.ID)
	if err != nil {
		return errors.New("data evaluasi kesehatan ibu tidak ditemukan")
	}
	return u.repo.Update(e)
}

func (u *evaluasiKesehatanIbuUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}

// MODUL IBU
func (u *evaluasiKesehatanIbuUsecase) GetByIDForOrangtua(id int32, userID int32) (*models.EvaluasiKesehatanIbu, error) {
	allowed, err := u.repo.IsOwnedByUser(id, userID)
	if err != nil {
		return nil, err
	}

	if !allowed {
		return nil, errors.New("Anda tidak memiliki akses ke data ini")
	}

	return u.repo.FindByID(id)
}

// MODUL IBU
func (u *evaluasiKesehatanIbuUsecase) GetMine(userID int32) (*models.EvaluasiKesehatanIbu, error) {
	return u.repo.FindMineByUserID(userID)
}