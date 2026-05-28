package usecases

import (
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
	"time"
)

type PemantauanAnakUseCase interface {
	SavePemantauan(req *models.LembarPemantauanRequest) error
	GetHistory(anakID uint, rentangID uint) ([]models.LembarPemantauan, error)
	GetRentangUsia() ([]models.RentangUsia, error)
	GetKategoriByRentang(rentangID uint) ([]models.KategoriTandaSakit, error)
	DeletePemantauan(id uint) error
	CreateKategori(data *models.KategoriTandaSakit) error
	UpdateKategori(id uint, data *models.KategoriTandaSakit) error
	DeleteKategori(id uint) error
	VerifikasiPemantauan(id uint, req *models.LembarPemantauanVerifikasiRequest) error
}

type pemantauanAnakUseCase struct {
	repo repositories.PemantauanAnakRepository
}

func NewPemantauanAnakUseCase(repo repositories.PemantauanAnakRepository) PemantauanAnakUseCase {
	return &pemantauanAnakUseCase{repo}
}

func (u *pemantauanAnakUseCase) SavePemantauan(req *models.LembarPemantauanRequest) error {
	tgl, err := time.Parse("2006-01-02", req.TanggalPeriksa)
	if err != nil {
		return err
	}

	details := make([]models.DetailPemantauan, len(req.DetailGejala))
	for i, d := range req.DetailGejala {
		details[i] = models.DetailPemantauan{
			KategoriTandaSakitID: d.KategoriTandaSakitID,
			IsTerjadi:            d.IsTerjadi,
		}
	}

	existing, _ := u.repo.GetByPeriode(req.AnakID, req.RentangUsiaID, req.PeriodeWaktu)
	if existing != nil {
		existing.TanggalPeriksa = &tgl
		existing.NamaPemeriksa = req.NamaPemeriksa
		existing.DetailGejala = details
		// If it's being updated, maybe we reset the status or keep it?
		// We'll reset it to 'Menunggu Verifikasi' if it's updated by kader/bidan again.
		existing.Status = "Menunggu Verifikasi"
		return u.repo.Update(existing)
	}

	record := &models.LembarPemantauan{
		AnakID:         req.AnakID,
		RentangUsiaID:  req.RentangUsiaID,
		PeriodeWaktu:   req.PeriodeWaktu,
		TanggalPeriksa: &tgl,
		NamaPemeriksa:  req.NamaPemeriksa,
		Status:         "Menunggu Verifikasi",
		DetailGejala:   details,
	}

	return u.repo.Create(record)
}

func (u *pemantauanAnakUseCase) GetHistory(anakID uint, rentangID uint) ([]models.LembarPemantauan, error) {
	return u.repo.FindByChildAndRange(anakID, rentangID)
}

func (u *pemantauanAnakUseCase) GetRentangUsia() ([]models.RentangUsia, error) {
	return u.repo.GetRentangUsia()
}

func (u *pemantauanAnakUseCase) GetKategoriByRentang(rentangID uint) ([]models.KategoriTandaSakit, error) {
	return u.repo.GetKategoriByRentang(rentangID)
}

func (u *pemantauanAnakUseCase) DeletePemantauan(id uint) error {
	return u.repo.Delete(id)
}

func (u *pemantauanAnakUseCase) CreateKategori(data *models.KategoriTandaSakit) error {
	return u.repo.CreateKategori(data)
}

func (u *pemantauanAnakUseCase) UpdateKategori(id uint, data *models.KategoriTandaSakit) error {
	data.ID = id
	return u.repo.UpdateKategori(data)
}

func (u *pemantauanAnakUseCase) DeleteKategori(id uint) error {
	return u.repo.DeleteKategori(id)
}

func (u *pemantauanAnakUseCase) VerifikasiPemantauan(id uint, req *models.LembarPemantauanVerifikasiRequest) error {
	// Pengecekan ada tidaknya data mungkin dilakukan di controller/repo
	return u.repo.UpdateStatus(id, req.Status, req.NamaPemeriksa)
}
