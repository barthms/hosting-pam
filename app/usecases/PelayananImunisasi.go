package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
	"time"
)

type KunjunganImunisasiUseCase interface {
	Create(req models.CreateKunjunganImunisasiRequest) error
	Update(id int32, req models.UpdateKunjunganImunisasiRequest) error
	GetByAnakID(anakID int32) ([]models.KehadiranImunisasi, error)
	GetByID(id int32) (*models.KehadiranImunisasi, error)
	GetAll() ([]models.KehadiranImunisasi, error)
	Delete(id int32) error
}

type kunjunganImunisasiUseCase struct {
	repo repositories.KunjunganImunisasiRepository
}

func NewKunjunganImunisasiUseCase(repo repositories.KunjunganImunisasiRepository) KunjunganImunisasiUseCase {
	return &kunjunganImunisasiUseCase{repo: repo}
}
func (u *kunjunganImunisasiUseCase) Create(req models.CreateKunjunganImunisasiRequest) error {

	if req.AnakID == 0 {
		return errors.New("anak_id wajib diisi")
	}

	now := time.Now()

	kunjungan := models.KehadiranImunisasi{
		AnakID:    req.AnakID,
		Bulanke:   req.Bulanke,
		CreatedAt: now,
		UpdatedAt: now,
	}

	for _, d := range req.Detail {
		kunjungan.Detail = append(kunjungan.Detail, models.DetailPelayananImunisasi{
			JenisPelayananID: d.JenisPelayananID,
			Keterangan:       d.Keterangan,
			CreatedAt:        now,
			UpdatedAt:        now,
		})
	}

	return u.repo.Create(&kunjungan)
}
func (u *kunjunganImunisasiUseCase) Update(id int32, req models.UpdateKunjunganImunisasiRequest) error {
	now := time.Now()
	return u.repo.Update(id, req, now)
}
func (u *kunjunganImunisasiUseCase) GetByAnakID(anakID int32) ([]models.KehadiranImunisasi, error) {
	return u.repo.GetByAnakID(anakID)
}

func (u *kunjunganImunisasiUseCase) GetByID(id int32) (*models.KehadiranImunisasi, error) {
	return u.repo.GetByID(id)
}

func (u *kunjunganImunisasiUseCase) GetAll() ([]models.KehadiranImunisasi, error) {
	return u.repo.GetAll()
}
func (u *kunjunganImunisasiUseCase) Delete(id int32) error {
	return u.repo.Delete(id)
}
