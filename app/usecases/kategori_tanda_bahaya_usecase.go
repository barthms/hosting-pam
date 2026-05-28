package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type KategoriTandaBahayaUsecase interface {
	Create(k *models.KategoriTandaBahaya) error
	GetByID(id int32) (*models.KategoriTandaBahaya, error)
	GetAll() ([]models.KategoriTandaBahaya, error)
	GetByTipeAndKategoriUsia(tipe, kategoriUsia string) ([]models.KategoriTandaBahaya, error)
	Update(k *models.KategoriTandaBahaya) error
	Delete(id int32) error
}

type kategoriTandaBahayaUsecase struct {
	repo *repositories.KategoriTandaBahayaRepository
}

func NewKategoriTandaBahayaUsecase(repo *repositories.KategoriTandaBahayaRepository) KategoriTandaBahayaUsecase {
	return &kategoriTandaBahayaUsecase{repo: repo}
}

func (u *kategoriTandaBahayaUsecase) Create(k *models.KategoriTandaBahaya) error {
	// Validasi input
	if k.TipeLembar == "" {
		return errors.New("tipe_lembar wajib diisi")
	}
	if k.Gejala == "" {
		return errors.New("gejala wajib diisi")
	}
	if k.KategoriUsia == "" {
		return errors.New("kategori_usia wajib diisi")
	}

	return u.repo.Create(k)
}

func (u *kategoriTandaBahayaUsecase) GetByID(id int32) (*models.KategoriTandaBahaya, error) {
	if id == 0 {
		return nil, errors.New("id tidak valid")
	}
	return u.repo.FindByID(id)
}

func (u *kategoriTandaBahayaUsecase) GetAll() ([]models.KategoriTandaBahaya, error) {
	return u.repo.FindAll()
}

func (u *kategoriTandaBahayaUsecase) GetByTipeAndKategoriUsia(tipe, kategoriUsia string) ([]models.KategoriTandaBahaya, error) {
	if tipe == "" || kategoriUsia == "" {
		return nil, errors.New("tipe dan kategori_usia wajib diisi")
	}
	return u.repo.FindByTipeAndKategoriUsia(tipe, kategoriUsia)
}

func (u *kategoriTandaBahayaUsecase) Update(k *models.KategoriTandaBahaya) error {
	if k.ID == 0 {
		return errors.New("id tidak valid")
	}

	// Periksa apakah data sudah ada
	_, err := u.repo.FindByID(k.ID)
	if err != nil {
		return errors.New("data kategori tanda bahaya tidak ditemukan")
	}

	return u.repo.Update(k)
}

func (u *kategoriTandaBahayaUsecase) Delete(id int32) error {
	if id == 0 {
		return errors.New("id tidak valid")
	}
	return u.repo.Delete(id)
}
