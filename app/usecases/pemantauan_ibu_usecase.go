package usecases

import (
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
	"time"
)

type PemantauanIbuUseCase interface {
	SavePemantauan(req *models.LembarPemantauanIbuRequest) error
	GetAllByIbu(ibuID int32) ([]models.LembarPemantauanIbu, error)
	GetKategori() ([]models.KategoriPemantauanIbu, error)
	DeletePemantauan(id int32) error
}

type pemantauanIbuUseCase struct {
	repo repositories.PemantauanIbuRepository
}

func NewPemantauanIbuUseCase(repo repositories.PemantauanIbuRepository) PemantauanIbuUseCase {
	return &pemantauanIbuUseCase{repo}
}

func (u *pemantauanIbuUseCase) SavePemantauan(req *models.LembarPemantauanIbuRequest) error {
	tgl, err := time.Parse("2006-01-02", req.TanggalPeriksa)
	if err != nil {
		return err
	}

	details := make([]models.DetailPemantauanIbu, len(req.DetailGejala))
	for i, d := range req.DetailGejala {
		details[i] = models.DetailPemantauanIbu{
			KategoriPemantauanID: d.KategoriPemantauanID,
			IsTerjadi:            d.IsTerjadi,
		}
	}

	existing, _ := u.repo.GetByPeriode(req.IbuID, req.PeriodeWaktu)
	if existing != nil {
		existing.TanggalPeriksa = tgl
		existing.Pemeriksa = req.Pemeriksa
		existing.DetailGejala = details
		return u.repo.Update(existing)
	}

	record := &models.LembarPemantauanIbu{
		IbuID:          req.IbuID,
		PeriodeWaktu:   req.PeriodeWaktu,
		TanggalPeriksa: tgl,
		Pemeriksa:      req.Pemeriksa,
		DetailGejala:   details,
	}

	return u.repo.Create(record)
}

func (u *pemantauanIbuUseCase) GetAllByIbu(ibuID int32) ([]models.LembarPemantauanIbu, error) {
	return u.repo.FindByIbu(ibuID)
}

func (u *pemantauanIbuUseCase) GetKategori() ([]models.KategoriPemantauanIbu, error) {
	return u.repo.GetKategori()
}

func (u *pemantauanIbuUseCase) DeletePemantauan(id int32) error {
	return u.repo.Delete(id)
}
