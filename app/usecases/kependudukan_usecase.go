package usecases

import (
	"errors"
	"log"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type KependudukanUsecase interface {
	Create(k *models.Kependudukan) (*models.Kependudukan, error)
	GetByID(id int32) (*models.Kependudukan, error)
	GetByNIK(nik *string) (*models.Kependudukan, error)
	GetAll() ([]models.Kependudukan, error)
	ListEligibleForRole(role, search, kecamatan, desa string) ([]repositories.EligiblePendudukItem, error)
	Update(k *models.Kependudukan) error
	Delete(id int32) error
	GetRekapPerDusun(kecamatan, desa string) ([]repositories.RekapDusun, error)
}

type kependudukanUsecase struct {
	repo *repositories.KependudukanRepository
}

func NewKependudukanUsecase(repo *repositories.KependudukanRepository) KependudukanUsecase {
	return &kependudukanUsecase{repo: repo}
}

func (u *kependudukanUsecase) Create(k *models.Kependudukan) (*models.Kependudukan, error) {
	// Cek apakah NIK sudah ada (jika NIK diisi)
	if k.NIK != nil && *k.NIK != "" {
		existing, err := u.repo.FindByNIK(k.NIK)
		if err == nil && existing != nil {
			return nil, errors.New("NIK sudah terdaftar")
		}
		// Jika error bukan record not found, return error
		if err != nil && err.Error() != "record not found" {
			log.Println("Error checking NIK:", err)
			return nil, err
		}
	}
	// Create
	err := u.repo.Create(k)
	if err != nil {
		log.Println("Error creating kependudukan:", err)
		return nil, err
	}
	return k, nil
}

func (u *kependudukanUsecase) GetByID(id int32) (*models.Kependudukan, error) {
	return u.repo.FindByID(id)
}

func (u *kependudukanUsecase) GetByNIK(nik *string) (*models.Kependudukan, error) {
	return u.repo.FindByNIK(nik)
}

func (u *kependudukanUsecase) GetAll() ([]models.Kependudukan, error) {
	return u.repo.GetAll()
}

func (u *kependudukanUsecase) ListEligibleForRole(role, search, kecamatan, desa string) ([]repositories.EligiblePendudukItem, error) {
	return u.repo.ListEligibleForRole(role, search, kecamatan, desa)
}

func (u *kependudukanUsecase) Update(k *models.Kependudukan) error {
	_, err := u.repo.FindByID(k.IDKependudukan)
	if err != nil {
		return errors.New("data kependudukan tidak ditemukan")
	}
	return u.repo.Update(k)
}

func (u *kependudukanUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}
func (u *kependudukanUsecase) GetRekapPerDusun(kecamatan, desa string) ([]repositories.RekapDusun, error) {
	return u.repo.GetRekapPerDusun(kecamatan, desa)
}