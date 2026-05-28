package usecases

// import (
// 	"errors"
// 	"log"
// 	"monitoring-service/app/models"
// 	"monitoring-service/app/repositories"
// )

// type KartuKeluargaUsecase interface {
// 	Create(kk *models.KartuKeluarga) (*models.KartuKeluarga, error)
// 	GetByID(id int32) (*models.KartuKeluarga, error)
// 	GetByNoKartuKeluarga(noKK string) (*models.KartuKeluarga, error)
// 	GetAll() ([]models.KartuKeluarga, error)
// 	Update(kk *models.KartuKeluarga) error
// 	Delete(id int32) error
// 	GetByNoKartuKeluargaWithMembers(noKK string) (*models.KartuKeluarga, error)
// }

// type kartuKeluargaUsecase struct {
// 	repo *repositories.KartuKeluargaRepository
// }

// func NewKartuKeluargaUsecase(repo *repositories.KartuKeluargaRepository) KartuKeluargaUsecase {
// 	return &kartuKeluargaUsecase{repo: repo}
// }

// func (u *kartuKeluargaUsecase) Create(kk *models.KartuKeluarga) (*models.KartuKeluarga, error) {
// 	// Validasi no_kk
// 	if kk.NoKK == "" {
// 		return nil, errors.New("no_kk wajib diisi")
// 	}

// 	// Cek apakah no_kk sudah ada
// 	existing, err := u.repo.FindByNoKartuKeluarga(kk.NoKK)
// 	if err == nil && existing != nil {
// 		return nil, errors.New("no_kk sudah terdaftar")
// 	}
// 	// Jika error bukan record not found, return error
// 	if err != nil && err.Error() != "record not found" {
// 		log.Println("Error checking no_kk:", err)
// 		return nil, err
// 	}

// 	// Create
// 	err = u.repo.Create(kk)
// 	if err != nil {
// 		log.Println("Error creating kartu_keluarga:", err)
// 		return nil, err
// 	}
// 	return kk, nil
// }

// func (u *kartuKeluargaUsecase) GetByID(id int32) (*models.KartuKeluarga, error) {
// 	return u.repo.FindByID(id)
// }

// func (u *kartuKeluargaUsecase) GetByNoKartuKeluarga(noKK string) (*models.KartuKeluarga, error) {
// 	if noKK == "" {
// 		return nil, errors.New("no_kk wajib diisi")
// 	}
// 	return u.repo.FindByNoKartuKeluarga(noKK)
// }

// func (u *kartuKeluargaUsecase) GetAll() ([]models.KartuKeluarga, error) {
// 	return u.repo.GetAll()
// }

// func (u *kartuKeluargaUsecase) Update(kk *models.KartuKeluarga) error {
// 	_, err := u.repo.FindByID(kk.ID)
// 	if err != nil {
// 		return errors.New("kartu_keluarga tidak ditemukan")
// 	}
// 	return u.repo.Update(kk)
// }

// func (u *kartuKeluargaUsecase) Delete(id int32) error {
// 	return u.repo.Delete(id)
// }

// func (u *kartuKeluargaUsecase) GetByNoKartuKeluargaWithMembers(noKK string) (*models.KartuKeluarga, error) {
// 	if noKK == "" {
// 		return nil, errors.New("no_kk wajib diisi")
// 	}
// 	return u.repo.GetByNoKartuKeluargaWithMembers(noKK)
// }
