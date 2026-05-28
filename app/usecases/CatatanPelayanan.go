package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
	"time"
)

type CatatanPelayananUseCase interface {
	Create(req models.CreateCatatanPelayananRequest) error
	Update(id int32, req models.UpdateCatatanPelayananRequest) error
	GetByAnakID(anakID int32) ([]models.CatatanPelayanan, error)
	GetByID(id int32) (*models.CatatanPelayanan, error)
	GetAll() ([]models.CatatanPelayanan, error)
	Delete(id int32) error
}

type catatanpelayananUseCase struct {
	repo repositories.CatatanPelayananRepository
}

func NewCatatanPelayananUseCase(repo repositories.CatatanPelayananRepository) CatatanPelayananUseCase {
	return &catatanpelayananUseCase{repo: repo}
}

func (u *catatanpelayananUseCase) Create(req models.CreateCatatanPelayananRequest) error {

	if req.AnakID == 0 {
		return errors.New("anak_id wajib diisi")
	}

	now := time.Now()

	pemeriksaan := models.CatatanPelayanan{
		AnakID:    req.AnakID,
		TenagaKesehatanID: req.TenagaKesehatanID,
		Tanggal_periksa: req.TanggalPeriksa,
		Tanggal_kembali: req.TanggalKembali,
		CatatanPelayanan: req.CatatanPelayanan,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return u.repo.Create(&pemeriksaan)
}
func (u *catatanpelayananUseCase) Update(id int32, req models.UpdateCatatanPelayananRequest) error {
	now := time.Now()
	return u.repo.Update(id, req, now)
}
func (u *catatanpelayananUseCase) GetByAnakID(anakID int32) ([]models.CatatanPelayanan, error) {
	return u.repo.GetByAnakID(anakID)
}

func (u *catatanpelayananUseCase) GetByID(id int32) (*models.CatatanPelayanan, error) {
	return u.repo.GetByID(id)
}

func (u *catatanpelayananUseCase) GetAll() ([]models.CatatanPelayanan, error) {
	return u.repo.GetAll()
}
func (u *catatanpelayananUseCase) Delete(id int32) error {
	return u.repo.Delete(id)
}
