package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type EdukasiMPASIRepository interface {
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

type edukasiMPASIRepository struct {
	db *gorm.DB
}

func NewEdukasiMPASIRepository(db *gorm.DB) EdukasiMPASIRepository {
	return &edukasiMPASIRepository{db}
}

// Implementations for Materi
func (r *edukasiMPASIRepository) CreateMateri(data *models.MateriMPASI) error {
	return r.db.Create(data).Error
}
func (r *edukasiMPASIRepository) GetMateriByBulan(bulan int) ([]models.MateriMPASI, error) {
	var data []models.MateriMPASI
	err := r.db.Where("(bulan_min <= ? AND bulan_max >= ?) OR (bulan_min IS NULL AND bulan_max IS NULL)", bulan, bulan).Find(&data).Error
	return data, err
}
func (r *edukasiMPASIRepository) GetMateriAll() ([]models.MateriMPASI, error) {
	var data []models.MateriMPASI
	err := r.db.Find(&data).Error
	return data, err
}

func (r *edukasiMPASIRepository) GetAll() ([]models.EdukasiMPASI, error) {
	var data []models.EdukasiMPASI
	err := r.db.Order("bulan_min asc, judul asc").Find(&data).Error
	return data, err
}
func (r *edukasiMPASIRepository) GetMateriByID(id int32) (*models.MateriMPASI, error) {
	var data models.MateriMPASI
	err := r.db.First(&data, id).Error
	return &data, err
}
func (r *edukasiMPASIRepository) UpdateMateri(id int32, data *models.MateriMPASI) error {
	return r.db.Model(&models.MateriMPASI{}).Where("id = ?", id).Updates(data).Error
}
func (r *edukasiMPASIRepository) DeleteMateri(id int32) error {
	return r.db.Delete(&models.MateriMPASI{}, id).Error
}

// Implementations for Aturan Porsi
func (r *edukasiMPASIRepository) CreateAturanPorsi(data *models.AturanPorsiMPASI) error {
	return r.db.Create(data).Error
}
func (r *edukasiMPASIRepository) GetAturanPorsiByBulan(bulan int) (*models.AturanPorsiMPASI, error) {
	var data models.AturanPorsiMPASI
	err := r.db.Where("bulan_min <= ? AND bulan_max >= ?", bulan, bulan).First(&data).Error
	return &data, err
}
func (r *edukasiMPASIRepository) GetAturanPorsiAll() ([]models.AturanPorsiMPASI, error) {
	var data []models.AturanPorsiMPASI
	err := r.db.Find(&data).Error
	return data, err
}
func (r *edukasiMPASIRepository) GetAturanPorsiByID(id int32) (*models.AturanPorsiMPASI, error) {
	var data models.AturanPorsiMPASI
	err := r.db.First(&data, id).Error
	return &data, err
}
func (r *edukasiMPASIRepository) UpdateAturanPorsi(id int32, data *models.AturanPorsiMPASI) error {
	return r.db.Model(&models.AturanPorsiMPASI{}).Where("id = ?", id).Updates(data).Error
}
func (r *edukasiMPASIRepository) DeleteAturanPorsi(id int32) error {
	return r.db.Delete(&models.AturanPorsiMPASI{}, id).Error
}

// Implementations for Jadwal
func (r *edukasiMPASIRepository) CreateJadwal(data *models.JadwalHarianMPASI) error {
	return r.db.Create(data).Error
}
func (r *edukasiMPASIRepository) GetJadwalByBulan(bulan int) ([]models.JadwalHarianMPASI, error) {
	var data []models.JadwalHarianMPASI
	err := r.db.Where("bulan_min <= ? AND bulan_max >= ?", bulan, bulan).Order("waktu ASC").Find(&data).Error
	return data, err
}
func (r *edukasiMPASIRepository) GetJadwalAll() ([]models.JadwalHarianMPASI, error) {
	var data []models.JadwalHarianMPASI
	err := r.db.Order("waktu ASC").Find(&data).Error
	return data, err
}
func (r *edukasiMPASIRepository) GetJadwalByID(id int32) (*models.JadwalHarianMPASI, error) {
	var data models.JadwalHarianMPASI
	err := r.db.First(&data, id).Error
	return &data, err
}
func (r *edukasiMPASIRepository) UpdateJadwal(id int32, data *models.JadwalHarianMPASI) error {
	return r.db.Model(&models.JadwalHarianMPASI{}).Where("id = ?", id).Updates(data).Error
}
func (r *edukasiMPASIRepository) DeleteJadwal(id int32) error {
	return r.db.Delete(&models.JadwalHarianMPASI{}, id).Error
}

// Implementations for Resep
func (r *edukasiMPASIRepository) CreateResep(data *models.ResepMPASI) error {
	return r.db.Create(data).Error
}
func (r *edukasiMPASIRepository) GetResepByBulan(bulan int) ([]models.ResepMPASI, error) {
	var data []models.ResepMPASI
	err := r.db.Where("bulan_min <= ? AND bulan_max >= ?", bulan, bulan).Find(&data).Error
	return data, err
}
func (r *edukasiMPASIRepository) GetResepAll() ([]models.ResepMPASI, error) {
	var data []models.ResepMPASI
	err := r.db.Find(&data).Error
	return data, err
}
func (r *edukasiMPASIRepository) GetResepByID(id int32) (*models.ResepMPASI, error) {
	var data models.ResepMPASI
	err := r.db.First(&data, id).Error
	return &data, err
}
func (r *edukasiMPASIRepository) UpdateResep(id int32, data *models.ResepMPASI) error {
	return r.db.Model(&models.ResepMPASI{}).Where("id = ?", id).Updates(data).Error
}
func (r *edukasiMPASIRepository) DeleteResep(id int32) error {
	return r.db.Delete(&models.ResepMPASI{}, id).Error
}
