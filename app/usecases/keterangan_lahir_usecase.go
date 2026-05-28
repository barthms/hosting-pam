// // app/usecases/keterangan_lahir_usecase.go
// package usecases

// import (
// 	"errors"
// 	"monitoring-service/app/models"
// 	"monitoring-service/app/repositories"
// )

// type KeteranganLahirUsecase interface {
// 	Create(k *models.KeteranganLahir) error
// 	GetByID(id int32) (*models.KeteranganLahir, error)
// 	GetByIbuID(ibuID int32) ([]models.KeteranganLahir, error)
// 	Update(k *models.KeteranganLahir) error
// 	Delete(id int32) error
// }

// type keteranganLahirUsecase struct {
// 	repo *repositories.KeteranganLahirRepository
// }

// func NewKeteranganLahirUsecase(repo *repositories.KeteranganLahirRepository) KeteranganLahirUsecase {
// 	return &keteranganLahirUsecase{repo: repo}
// }

// func (u *keteranganLahirUsecase) Create(k *models.KeteranganLahir) error {
// 	if k.IDIbuRelasi == 0 {
// 		return errors.New("id_ibu_relasi wajib diisi")
// 	}
// 	return u.repo.Create(k)
// }

// func (u *keteranganLahirUsecase) GetByID(id int32) (*models.KeteranganLahir, error) {
// 	return u.repo.FindByID(id)
// }

// func (u *keteranganLahirUsecase) GetByIbuID(ibuID int32) ([]models.KeteranganLahir, error) {
// 	if ibuID == 0 {
// 		return nil, errors.New("ibu_id wajib diisi")
// 	}
// 	return u.repo.FindByIbuID(ibuID)
// }

// func (u *keteranganLahirUsecase) Update(k *models.KeteranganLahir) error {
// 	_, err := u.repo.FindByID(k.ID)
// 	if err != nil {
// 		return errors.New("data keterangan lahir tidak ditemukan")
// 	}
// 	return u.repo.Update(k)
// }

// func (u *keteranganLahirUsecase) Delete(id int32) error {
// 	return u.repo.Delete(id)
// }


// app/usecases/keterangan_lahir_usecase.go
package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type KeteranganLahirUsecase interface {
	Create(k *models.KeteranganLahir) error
	GetByID(id int32) (*models.KeteranganLahir, error)
	GetByIbuID(ibuID int32) ([]models.KeteranganLahir, error)
	GetMine(userID int32) (*models.KeteranganLahir, error)
	Update(k *models.KeteranganLahir) error
	Delete(id int32) error
}

type keteranganLahirUsecase struct {
	repo *repositories.KeteranganLahirRepository
}

func NewKeteranganLahirUsecase(repo *repositories.KeteranganLahirRepository) KeteranganLahirUsecase {
	return &keteranganLahirUsecase{repo: repo}
}

func (u *keteranganLahirUsecase) Create(k *models.KeteranganLahir) error {
	if k.IDIbuRelasi == 0 {
		return errors.New("id_ibu_relasi wajib diisi")
	}
	return u.repo.Create(k)
}

func (u *keteranganLahirUsecase) GetByID(id int32) (*models.KeteranganLahir, error) {
	return u.repo.FindByID(id)
}

func (u *keteranganLahirUsecase) GetByIbuID(ibuID int32) ([]models.KeteranganLahir, error) {
	if ibuID == 0 {
		return nil, errors.New("ibu_id wajib diisi")
	}
	return u.repo.FindByIbuID(ibuID)
}

func (u *keteranganLahirUsecase) GetMine(userID int32) (*models.KeteranganLahir, error) {
	if userID == 0 {
		return nil, errors.New("user_id tidak valid")
	}
	return u.repo.FindMineByUserID(userID)
}

func (u *keteranganLahirUsecase) Update(k *models.KeteranganLahir) error {
	_, err := u.repo.FindByID(k.ID)
	if err != nil {
		return errors.New("data keterangan lahir tidak ditemukan")
	}
	return u.repo.Update(k)
}

func (u *keteranganLahirUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}