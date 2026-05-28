package usecases

import (
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type EdukasiMPASIUsecase interface {
	// Materi
	CreateMateri(data *models.MateriMPASI) error
	GetMateriByBulan(bulan int) ([]models.MateriMPASI, error)
	GetMateriAll() ([]models.MateriMPASI, error)
	GetMateriByID(id int32) (*models.MateriMPASI, error)
	UpdateMateri(id int32, data *models.MateriMPASI) error
	DeleteMateri(id int32) error

	// Aturan Porsi
	CreateAturanPorsi(data *models.AturanPorsiMPASI) error
	GetAturanPorsiByBulan(bulan int) (*models.AturanPorsiMPASI, error)
	GetAturanPorsiAll() ([]models.AturanPorsiMPASI, error)
	GetAturanPorsiByID(id int32) (*models.AturanPorsiMPASI, error)
	UpdateAturanPorsi(id int32, data *models.AturanPorsiMPASI) error
	DeleteAturanPorsi(id int32) error

	// Jadwal
	CreateJadwal(data *models.JadwalHarianMPASI) error
	GetJadwalByBulan(bulan int) ([]models.JadwalHarianMPASI, error)
	GetJadwalAll() ([]models.JadwalHarianMPASI, error)
	GetJadwalByID(id int32) (*models.JadwalHarianMPASI, error)
	UpdateJadwal(id int32, data *models.JadwalHarianMPASI) error
	DeleteJadwal(id int32) error

	// Resep
	CreateResep(data *models.ResepMPASI) error
	GetResepByBulan(bulan int) ([]models.ResepMPASI, error)
	GetResepAll() ([]models.ResepMPASI, error)
	GetResepByID(id int32) (*models.ResepMPASI, error)
	UpdateResep(id int32, data *models.ResepMPASI) error
	DeleteResep(id int32) error
}

type edukasiMPASIUsecase struct {
	repo repositories.EdukasiMPASIRepository
}

func NewEdukasiMPASIUsecase(repo repositories.EdukasiMPASIRepository) EdukasiMPASIUsecase {
	return &edukasiMPASIUsecase{repo}
}

// Implementations for Materi
func (u *edukasiMPASIUsecase) CreateMateri(data *models.MateriMPASI) error {
	return u.repo.CreateMateri(data)
}
func (u *edukasiMPASIUsecase) GetMateriByBulan(bulan int) ([]models.MateriMPASI, error) {
	return u.repo.GetMateriByBulan(bulan)
}
func (u *edukasiMPASIUsecase) GetMateriAll() ([]models.MateriMPASI, error) {
	return u.repo.GetMateriAll()
}
func (u *edukasiMPASIUsecase) GetMateriByID(id int32) (*models.MateriMPASI, error) {
	return u.repo.GetMateriByID(id)
}
func (u *edukasiMPASIUsecase) UpdateMateri(id int32, data *models.MateriMPASI) error {
	return u.repo.UpdateMateri(id, data)
}
func (u *edukasiMPASIUsecase) DeleteMateri(id int32) error {
	return u.repo.DeleteMateri(id)
}

// Implementations for Aturan Porsi
func (u *edukasiMPASIUsecase) CreateAturanPorsi(data *models.AturanPorsiMPASI) error {
	return u.repo.CreateAturanPorsi(data)
}
func (u *edukasiMPASIUsecase) GetAturanPorsiByBulan(bulan int) (*models.AturanPorsiMPASI, error) {
	return u.repo.GetAturanPorsiByBulan(bulan)
}
func (u *edukasiMPASIUsecase) GetAturanPorsiAll() ([]models.AturanPorsiMPASI, error) {
	return u.repo.GetAturanPorsiAll()
}
func (u *edukasiMPASIUsecase) GetAturanPorsiByID(id int32) (*models.AturanPorsiMPASI, error) {
	return u.repo.GetAturanPorsiByID(id)
}
func (u *edukasiMPASIUsecase) UpdateAturanPorsi(id int32, data *models.AturanPorsiMPASI) error {
	return u.repo.UpdateAturanPorsi(id, data)
}
func (u *edukasiMPASIUsecase) DeleteAturanPorsi(id int32) error {
	return u.repo.DeleteAturanPorsi(id)
}

// Implementations for Jadwal
func (u *edukasiMPASIUsecase) CreateJadwal(data *models.JadwalHarianMPASI) error {
	return u.repo.CreateJadwal(data)
}
func (u *edukasiMPASIUsecase) GetJadwalByBulan(bulan int) ([]models.JadwalHarianMPASI, error) {
	return u.repo.GetJadwalByBulan(bulan)
}
func (u *edukasiMPASIUsecase) GetJadwalAll() ([]models.JadwalHarianMPASI, error) {
	return u.repo.GetJadwalAll()
}
func (u *edukasiMPASIUsecase) GetJadwalByID(id int32) (*models.JadwalHarianMPASI, error) {
	return u.repo.GetJadwalByID(id)
}
func (u *edukasiMPASIUsecase) UpdateJadwal(id int32, data *models.JadwalHarianMPASI) error {
	return u.repo.UpdateJadwal(id, data)
}
func (u *edukasiMPASIUsecase) DeleteJadwal(id int32) error {
	return u.repo.DeleteJadwal(id)
}

// Implementations for Resep
func (u *edukasiMPASIUsecase) CreateResep(data *models.ResepMPASI) error {
	return u.repo.CreateResep(data)
}
func (u *edukasiMPASIUsecase) GetResepByBulan(bulan int) ([]models.ResepMPASI, error) {
	return u.repo.GetResepByBulan(bulan)
}
func (u *edukasiMPASIUsecase) GetResepAll() ([]models.ResepMPASI, error) {
	return u.repo.GetResepAll()
}
func (u *edukasiMPASIUsecase) GetResepByID(id int32) (*models.ResepMPASI, error) {
	return u.repo.GetResepByID(id)
}
func (u *edukasiMPASIUsecase) UpdateResep(id int32, data *models.ResepMPASI) error {
	return u.repo.UpdateResep(id, data)
}
func (u *edukasiMPASIUsecase) DeleteResep(id int32) error {
	return u.repo.DeleteResep(id)
}