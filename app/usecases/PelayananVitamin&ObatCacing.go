package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
	"time"
)
type KunjunganVitaminUseCase interface {
	Create(req models.CreateKunjunganVitaminRequest) error
	Update(id int32, req models.UpdateKunjunganVitaminRequest) error
	GetByAnakID(anakID int32) ([]models.KunjunganVitamin, error)
	GetByID(id int32) (*models.KunjunganVitamin, error)
	GetAll() ([]models.KunjunganVitamin, error)
	Delete(id int32) error
}

type kunjunganVitaminUseCase struct {
	repo repositories.KunjunganVitaminRepository
}

func NewKunjunganVitaminUseCase(repo repositories.KunjunganVitaminRepository) KunjunganVitaminUseCase {
	return &kunjunganVitaminUseCase{repo: repo}
}
func (u *kunjunganVitaminUseCase) Create(req models.CreateKunjunganVitaminRequest) error {

	if req.AnakID == 0 {
		return errors.New("anak_id wajib diisi")
	}

	now := time.Now()
	tgl := now
	if req.Tanggal != "" {
		if t, err := time.Parse("2006-01-02", req.Tanggal); err == nil {
			tgl = t
		} else if t, err := time.Parse(time.RFC3339, req.Tanggal); err == nil {
			tgl = t
		}
	}

	kunjungan := models.KunjunganVitamin{
		AnakID:    req.AnakID,
		Tanggal:   tgl,
		CreatedAt: now,
		UpdatedAt: now,
	}

	for _, d := range req.Detail {
		kunjungan.Detail = append(kunjungan.Detail, models.DetailPelayananVitamin{
			JenisPelayananID: d.JenisPelayananID,
			Keterangan:       d.Keterangan,
			CreatedAt:        now,
			UpdatedAt:        now,
		})
	}

	return u.repo.Create(&kunjungan)
}
func (u *kunjunganVitaminUseCase) Update(id int32, req models.UpdateKunjunganVitaminRequest) error {
	now := time.Now()
	return u.repo.Update(id, req, now)
}
func (u *kunjunganVitaminUseCase) GetByAnakID(anakID int32) ([]models.KunjunganVitamin, error) {
	return u.repo.GetByAnakID(anakID)
}

func (u *kunjunganVitaminUseCase) GetByID(id int32) (*models.KunjunganVitamin, error) {
	return u.repo.GetByID(id)
}

func (u *kunjunganVitaminUseCase) GetAll() ([]models.KunjunganVitamin, error) {
	return u.repo.GetAll()
}
func (u *kunjunganVitaminUseCase) Delete(id int32) error {
	return u.repo.Delete(id)
}
