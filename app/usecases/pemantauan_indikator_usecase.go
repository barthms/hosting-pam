package usecases

import (
	"errors"
	"strings"

	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type PemantauanIndikatorUsecase interface {
	Create(data *models.PemantauanIndikator) (*models.PemantauanIndikator, error)
	GetAll(kategoriUsia, q string) ([]models.PemantauanIndikator, error)
	Update(id int32, kategoriUsia, deskripsi string) (*models.PemantauanIndikator, error)
	Delete(id int32) error
}

type pemantauanIndikatorUsecase struct {
	repo *repositories.PemantauanIndikatorRepository
}

func NewPemantauanIndikatorUsecase(repo *repositories.PemantauanIndikatorRepository) PemantauanIndikatorUsecase {
	return &pemantauanIndikatorUsecase{repo: repo}
}

func (u *pemantauanIndikatorUsecase) Create(data *models.PemantauanIndikator) (*models.PemantauanIndikator, error) {
	data.KategoriUsia = strings.TrimSpace(data.KategoriUsia)
	data.Deskripsi = strings.TrimSpace(data.Deskripsi)

	if data.KategoriUsia == "" {
		return nil, errors.New("kategori_usia wajib diisi")
	}
	if data.Deskripsi == "" {
		return nil, errors.New("deskripsi wajib diisi")
	}

	if err := u.repo.Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (u *pemantauanIndikatorUsecase) GetAll(kategoriUsia, q string) ([]models.PemantauanIndikator, error) {
	return u.repo.GetAll(strings.TrimSpace(kategoriUsia), strings.TrimSpace(q))
}

func (u *pemantauanIndikatorUsecase) Update(id int32, kategoriUsia, deskripsi string) (*models.PemantauanIndikator, error) {
	existing, err := u.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("indikator pemantauan tidak ditemukan")
	}

	if trimmedKategori := strings.TrimSpace(kategoriUsia); trimmedKategori != "" {
		existing.KategoriUsia = trimmedKategori
	}
	if trimmedDeskripsi := strings.TrimSpace(deskripsi); trimmedDeskripsi != "" {
		existing.Deskripsi = trimmedDeskripsi
	}

	if strings.TrimSpace(existing.KategoriUsia) == "" {
		return nil, errors.New("kategori_usia wajib diisi")
	}
	if strings.TrimSpace(existing.Deskripsi) == "" {
		return nil, errors.New("deskripsi wajib diisi")
	}

	if err := u.repo.Update(existing); err != nil {
		return nil, err
	}
	return existing, nil
}

func (u *pemantauanIndikatorUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}
