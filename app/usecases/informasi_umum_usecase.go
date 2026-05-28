package usecases

import (
	"errors"
	"strings"

	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type InformasiUmumUsecase interface {
	Create(data *models.InformasiUmum) error
	GetByID(id int32) (*models.InformasiUmum, error)
	GetAll() ([]models.InformasiUmum, error)
	Update(data *models.InformasiUmum) error
	Delete(id int32) error
}

type informasiUmumUsecase struct {
	repo *repositories.InformasiUmumRepository
}

func NewInformasiUmumUsecase(repo *repositories.InformasiUmumRepository) InformasiUmumUsecase {
	return &informasiUmumUsecase{repo: repo}
}

func (u *informasiUmumUsecase) Create(data *models.InformasiUmum) error {
	if data == nil {
		return errors.New("payload informasi umum tidak valid")
	}
	if data.Tipe == "" {
		return errors.New("tipe wajib diisi")
	}
	if data.Judul == "" {
		return errors.New("judul wajib diisi")
	}
	if data.Konten == "" {
		return errors.New("konten wajib diisi")
	}

	data.Tipe = strings.ToUpper(strings.TrimSpace(data.Tipe))
	if data.Tipe != "ARTIKEL" && data.Tipe != "VIDEO" {
		return errors.New("tipe hanya boleh ARTIKEL atau VIDEO")
	}

	if strings.TrimSpace(data.DurasiBaca) == "" {
		data.DurasiBaca = "-"
	}
	if strings.TrimSpace(data.UmurTarget) == "" {
		data.UmurTarget = "Semua Umur"
	}
	if strings.TrimSpace(data.Ringkasan) == "" {
		data.Ringkasan = data.Judul
	}
	if strings.TrimSpace(data.YangPerluDiingat) == "" {
		data.YangPerluDiingat = ""
	}

	return u.repo.Create(data)
}

func (u *informasiUmumUsecase) GetByID(id int32) (*models.InformasiUmum, error) {
	if id == 0 {
		return nil, errors.New("id tidak valid")
	}
	return u.repo.FindByID(id)
}

func (u *informasiUmumUsecase) GetAll() ([]models.InformasiUmum, error) {
	return u.repo.FindAll()
}

func (u *informasiUmumUsecase) Update(data *models.InformasiUmum) error {
	if data == nil {
		return errors.New("payload informasi umum tidak valid")
	}
	if data.ID == 0 {
		return errors.New("id tidak valid")
	}

	_, err := u.repo.FindByID(data.ID)
	if err != nil {
		return errors.New("data informasi umum tidak ditemukan")
	}

	data.Tipe = strings.ToUpper(strings.TrimSpace(data.Tipe))
	if data.Tipe != "ARTIKEL" && data.Tipe != "VIDEO" {
		return errors.New("tipe hanya boleh ARTIKEL atau VIDEO")
	}
	if data.Judul == "" {
		return errors.New("judul wajib diisi")
	}
	if data.Konten == "" {
		return errors.New("konten wajib diisi")
	}
	if strings.TrimSpace(data.DurasiBaca) == "" {
		data.DurasiBaca = "-"
	}
	if strings.TrimSpace(data.UmurTarget) == "" {
		data.UmurTarget = "Semua Umur"
	}
	if strings.TrimSpace(data.Ringkasan) == "" {
		data.Ringkasan = data.Judul
	}
	if strings.TrimSpace(data.YangPerluDiingat) == "" {
		data.YangPerluDiingat = ""
	}

	return u.repo.Update(data)
}

func (u *informasiUmumUsecase) Delete(id int32) error {
	if id == 0 {
		return errors.New("id tidak valid")
	}
	return u.repo.Delete(id)
}
