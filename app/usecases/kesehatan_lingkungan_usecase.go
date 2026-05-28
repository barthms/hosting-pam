package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
	"time"
)

type KesehatanLingkunganUsecase interface {
	// Kategori & Indikator
	CreateKategori(req models.KategoriLingkunganRequest) error
	GetAllKategori() ([]models.KategoriLingkungan, error)
	DeleteKategori(id int32) error

	AddIndikator(kategoriID int32, req models.IndikatorLingkunganRequest) error
	DeleteIndikator(id int32) error

	// Lembar & Detail
	SubmitLembar(req models.LembarLingkunganRequest) error
	GetHistory(ibuID int32) ([]models.LembarLingkungan, error)
	GetDetail(id int32) (*models.LembarLingkungan, error)
	DeleteLembar(id int32) error
}

type kesehatanLingkunganUsecase struct {
	repo repositories.KesehatanLingkunganRepository
}

func NewKesehatanLingkunganUsecase(repo repositories.KesehatanLingkunganRepository) KesehatanLingkunganUsecase {
	return &kesehatanLingkunganUsecase{repo}
}

func (u *kesehatanLingkunganUsecase) CreateKategori(req models.KategoriLingkunganRequest) error {
	kategori := &models.KategoriLingkungan{
		Nama:      req.Nama,
		Deskripsi: req.Deskripsi,
	}

	for _, ind := range req.Indikator {
		kategori.Indikator = append(kategori.Indikator, models.IndikatorLingkungan{
			Pertanyaan: ind.Pertanyaan,
		})
	}

	return u.repo.CreateKategori(kategori)
}

func (u *kesehatanLingkunganUsecase) GetAllKategori() ([]models.KategoriLingkungan, error) {
	return u.repo.GetAllKategori()
}

func (u *kesehatanLingkunganUsecase) DeleteKategori(id int32) error {
	return u.repo.DeleteKategori(id)
}

func (u *kesehatanLingkunganUsecase) AddIndikator(kategoriID int32, req models.IndikatorLingkunganRequest) error {
	indikator := &models.IndikatorLingkungan{
		KategoriLingkunganID: kategoriID,
		Pertanyaan:           req.Pertanyaan,
	}
	return u.repo.CreateIndikator(indikator)
}

func (u *kesehatanLingkunganUsecase) DeleteIndikator(id int32) error {
	return u.repo.DeleteIndikator(id)
}

func (u *kesehatanLingkunganUsecase) SubmitLembar(req models.LembarLingkunganRequest) error {
	tanggal, err := time.Parse("2006-01-02", req.TanggalPeriksa)
	if err != nil {
		return errors.New("format tanggal tidak valid (YYYY-MM-DD)")
	}

	lembar := &models.LembarLingkungan{
		IbuID:          req.IbuID,
		TanggalPeriksa: tanggal,
		Pemeriksa:      req.Pemeriksa,
		Catatan:        req.Catatan,
	}

	for _, det := range req.DetailJawaban {
		lembar.DetailJawaban = append(lembar.DetailJawaban, models.DetailLingkungan{
			IndikatorLingkunganID: det.IndikatorLingkunganID,
			IsOk:                 det.IsOk,
		})
	}

	return u.repo.SaveLembar(lembar)
}

func (u *kesehatanLingkunganUsecase) GetHistory(ibuID int32) ([]models.LembarLingkungan, error) {
	return u.repo.GetLembarHistory(ibuID)
}

func (u *kesehatanLingkunganUsecase) GetDetail(id int32) (*models.LembarLingkungan, error) {
	return u.repo.GetLembarByID(id)
}

func (u *kesehatanLingkunganUsecase) DeleteLembar(id int32) error {
	return u.repo.DeleteLembar(id)
}
