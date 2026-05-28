package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type SkriningPemantauanUsecase interface {
	Create(s *models.SkriningPemantauan) error
	GetByID(id int32) (*models.SkriningPemantauan, error)
	GetByAnakID(anakID int32) ([]models.SkriningPemantauan, error)
	GetAll() ([]models.SkriningPemantauan, error)
	Update(s *models.SkriningPemantauan) error
	Delete(id int32) error
}

type skriningPemantauanUsecase struct {
	repo *repositories.SkriningPemantauanRepository
}

func NewSkriningPemantauanUsecase(repo *repositories.SkriningPemantauanRepository) SkriningPemantauanUsecase {
	return &skriningPemantauanUsecase{repo: repo}
}

func (u *skriningPemantauanUsecase) Create(s *models.SkriningPemantauan) error {
	// Validasi input
	if s.AnakID == 0 {
		return errors.New("anak_id wajib diisi")
	}
	if s.KategoriTandaBahayaID == 0 {
		return errors.New("kategori_tanda_bahaya_id wajib diisi")
	}
	if s.PeriodeWaktu == 0 {
		return errors.New("periode_waktu wajib diisi")
	}
	if s.SatuanWaktu == "" {
		return errors.New("satuan_waktu wajib diisi")
	}
	if s.TglSkrining == "" {
		return errors.New("tgl_skrining wajib diisi")
	}

	return u.repo.Create(s)
}

func (u *skriningPemantauanUsecase) GetByID(id int32) (*models.SkriningPemantauan, error) {
	if id == 0 {
		return nil, errors.New("id tidak valid")
	}
	return u.repo.FindByID(id)
}

func (u *skriningPemantauanUsecase) GetByAnakID(anakID int32) ([]models.SkriningPemantauan, error) {
	if anakID == 0 {
		return nil, errors.New("anak_id tidak valid")
	}
	return u.repo.FindByAnakID(anakID)
}

func (u *skriningPemantauanUsecase) GetAll() ([]models.SkriningPemantauan, error) {
	return u.repo.FindAll()
}

func (u *skriningPemantauanUsecase) Update(s *models.SkriningPemantauan) error {
	if s.ID == 0 {
		return errors.New("id tidak valid")
	}
	if s.AnakID == 0 {
		return errors.New("anak_id wajib diisi")
	}
	if s.KategoriTandaBahayaID == 0 {
		return errors.New("kategori_tanda_bahaya_id wajib diisi")
	}

	// Periksa apakah data sudah ada
	_, err := u.repo.FindByID(s.ID)
	if err != nil {
		return errors.New("data skrining pemantauan tidak ditemukan")
	}

	return u.repo.Update(s)
}

func (u *skriningPemantauanUsecase) Delete(id int32) error {
	if id == 0 {
		return errors.New("id tidak valid")
	}
	return u.repo.Delete(id)
}
