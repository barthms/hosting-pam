package usecases

import (
	"errors"
	"strings"
	"time"

	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

// PerawatanUsecase defines the business logic interface for perawatan
type PerawatanUsecase interface {
	// Perawatan operations
	CreatePerawatan(req models.CreatePerawatanRequest) (*models.Perawatan, error)
	GetPerawatanByID(id uint) (*models.Perawatan, error)
	GetPerawatanByAnakID(anakID int32) ([]models.Perawatan, error)
	GetPerawatanByAnakIDAndRentangUsia(anakID int32, rentangUsia string) ([]models.Perawatan, error)
	UpdatePerawatan(id uint, req models.UpdatePerawatanRequest) (*models.Perawatan, error)
	DeletePerawatan(id uint) error

	// KategoriCapaian operations
	GetAllKategoriCapaian() ([]models.KategoriCapaian, error)
	GetKategoriCapaianByRentangUsia(rentangUsia string) ([]models.KategoriCapaian, error)

	// Access control operations for ibu
	CreatePerawatanForIbu(req models.CreatePerawatanRequest, userID int32) (*models.Perawatan, error)
	GetPerawatanByAnakIDForIbu(anakID int32, userID int32) ([]models.Perawatan, error)
	GetPerawatanByAnakIDAndRentangUsiaForIbu(anakID int32, rentangUsia string, userID int32) ([]models.Perawatan, error)
	UpdatePerawatanForIbu(id uint, req models.UpdatePerawatanRequest, userID int32) (*models.Perawatan, error)
	DeletePerawatanForIbu(id uint, userID int32) error
}

// perawatanUsecase is concrete implementation of PerawatanUsecase
type perawatanUsecase struct {
	repo *repositories.Main
}

// NewPerawatanUsecase creates and returns a new PerawatanUsecase
func NewPerawatanUsecase(repo *repositories.Main) PerawatanUsecase {
	return &perawatanUsecase{
		repo: repo,
	}
}

// ─────────────────────────────────────────────────────────
// PERAWATAN OPERATIONS
// ─────────────────────────────────────────────────────────

func parsePerawatanTanggalPeriksa(value string) (*time.Time, error) {
	if strings.TrimSpace(value) == "" {
		return nil, nil
	}

	dateValue := strings.TrimSpace(value)
	if len(dateValue) >= 10 {
		dateValue = dateValue[:10]
	}

	parsed, err := time.Parse("2006-01-02", dateValue)
	if err != nil {
		return nil, errors.New("format tanggal_periksa tidak valid")
	}

	return &parsed, nil
}

// CreatePerawatan creates a new perawatan record with validation
func (u *perawatanUsecase) CreatePerawatan(req models.CreatePerawatanRequest) (*models.Perawatan, error) {
	// Validate required fields
	if req.AnakID == 0 {
		return nil, errors.New("anak_id tidak boleh kosong")
	}
	if req.KategoriCapaianID == 0 {
		return nil, errors.New("kategori_capaian_id tidak boleh kosong")
	}

	// Verify anak exists
	exists, err := u.repo.Perawatan.IsAnakExist(req.AnakID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("anak tidak ditemukan")
	}

	tanggalPeriksa, err := parsePerawatanTanggalPeriksa(req.TanggalPeriksa)
	if err != nil {
		return nil, err
	}

	// Verify kategori capaian exists
	_, err = u.repo.Perawatan.GetKategoriCapaianByID(uint(req.KategoriCapaianID))
	if err != nil {
		return nil, errors.New("kategori capaian tidak ditemukan")
	}

	// Create perawatan record
	now := time.Now()
	perawatan := &models.Perawatan{
		AnakID:            req.AnakID,
		KategoriCapaianID: req.KategoriCapaianID,
		Jawaban:           req.Jawaban,
		TanggalPeriksa:    tanggalPeriksa,
		CreatedAt:         now,
		UpdatedAt:         now,
	}

	if err := u.repo.Perawatan.CreatePerawatan(perawatan); err != nil {
		return nil, err
	}

	// Reload with relations
	return u.repo.Perawatan.GetPerawatanByID(perawatan.ID)
}

// GetPerawatanByID retrieves a perawatan by ID
func (u *perawatanUsecase) GetPerawatanByID(id uint) (*models.Perawatan, error) {
	if id == 0 {
		return nil, errors.New("id tidak valid")
	}
	return u.repo.Perawatan.GetPerawatanByID(id)
}

// GetPerawatanByAnakID retrieves all perawatan for a child
func (u *perawatanUsecase) GetPerawatanByAnakID(anakID int32) ([]models.Perawatan, error) {
	if anakID <= 0 {
		return nil, errors.New("anak_id tidak valid")
	}

	// Verify anak exists
	exists, err := u.repo.Perawatan.IsAnakExist(anakID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("anak tidak ditemukan")
	}

	return u.repo.Perawatan.GetPerawatanByAnakID(anakID)
}

// GetPerawatanByAnakIDAndRentangUsia retrieves perawatan for a child with specific age range
func (u *perawatanUsecase) GetPerawatanByAnakIDAndRentangUsia(anakID int32, rentangUsia string) ([]models.Perawatan, error) {
	if anakID <= 0 {
		return nil, errors.New("anak_id tidak valid")
	}
	if rentangUsia == "" {
		return nil, errors.New("rentang_usia tidak boleh kosong")
	}

	// Verify anak exists
	exists, err := u.repo.Perawatan.IsAnakExist(anakID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("anak tidak ditemukan")
	}

	return u.repo.Perawatan.GetPerawatanByAnakIDAndRentangUsia(anakID, rentangUsia)
}

// UpdatePerawatan updates an existing perawatan record
func (u *perawatanUsecase) UpdatePerawatan(id uint, req models.UpdatePerawatanRequest) (*models.Perawatan, error) {
	if id == 0 {
		return nil, errors.New("id tidak valid")
	}

	// Get existing perawatan
	perawatan, err := u.repo.Perawatan.GetPerawatanByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Jawaban != nil {
		perawatan.Jawaban = req.Jawaban
	}
	if strings.TrimSpace(req.TanggalPeriksa) != "" {
		parsedTanggal, err := parsePerawatanTanggalPeriksa(req.TanggalPeriksa)
		if err != nil {
			return nil, err
		}
		perawatan.TanggalPeriksa = parsedTanggal
	}

	// Save updated perawatan
	if err := u.repo.Perawatan.UpdatePerawatan(perawatan); err != nil {
		return nil, err
	}

	// Reload with relations
	return u.repo.Perawatan.GetPerawatanByID(id)
}

// DeletePerawatan soft deletes a perawatan record
func (u *perawatanUsecase) DeletePerawatan(id uint) error {
	if id == 0 {
		return errors.New("id tidak valid")
	}

	// Verify perawatan exists
	_, err := u.repo.Perawatan.GetPerawatanByID(id)
	if err != nil {
		return err
	}

	return u.repo.Perawatan.DeletePerawatan(id)
}

// ─────────────────────────────────────────────────────────
// KATEGORI CAPAIAN OPERATIONS
// ─────────────────────────────────────────────────────────

// GetAllKategoriCapaian retrieves all kategori capaian
func (u *perawatanUsecase) GetAllKategoriCapaian() ([]models.KategoriCapaian, error) {
	return u.repo.Perawatan.GetAllKategoriCapaian()
}

// GetKategoriCapaianByRentangUsia retrieves kategori capaian for a specific age range
func (u *perawatanUsecase) GetKategoriCapaianByRentangUsia(rentangUsia string) ([]models.KategoriCapaian, error) {
	if rentangUsia == "" {
		return nil, errors.New("rentang_usia tidak boleh kosong")
	}
	return u.repo.Perawatan.GetKategoriCapaianByRentangUsia(rentangUsia)
}

// ─────────────────────────────────────────────────────────
// ACCESS CONTROL OPERATIONS FOR IBU
// ─────────────────────────────────────────────────────────

// CreatePerawatanForIbu creates perawatan with ownership validation
func (u *perawatanUsecase) CreatePerawatanForIbu(req models.CreatePerawatanRequest, userID int32) (*models.Perawatan, error) {
	// Validate required fields
	if req.AnakID == 0 {
		return nil, errors.New("anak_id tidak boleh kosong")
	}
	if req.KategoriCapaianID == 0 {
		return nil, errors.New("kategori_capaian_id tidak boleh kosong")
	}

	// Verify anak exists
	exists, err := u.repo.Perawatan.IsAnakExist(req.AnakID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("anak tidak ditemukan")
	}

	// Verify ibu owns this anak
	owned, err := u.repo.Perawatan.IsAnakOwnedByIbu(req.AnakID, userID)
	if err != nil {
		return nil, err
	}
	if !owned {
		return nil, errors.New("Anda tidak memiliki akses ke anak ini")
	}

	tanggalPeriksa, err := parsePerawatanTanggalPeriksa(req.TanggalPeriksa)
	if err != nil {
		return nil, err
	}

	// Verify kategori capaian exists
	_, err = u.repo.Perawatan.GetKategoriCapaianByID(uint(req.KategoriCapaianID))
	if err != nil {
		return nil, errors.New("kategori capaian tidak ditemukan")
	}

	// Create perawatan record
	now := time.Now()
	perawatan := &models.Perawatan{
		AnakID:            req.AnakID,
		KategoriCapaianID: req.KategoriCapaianID,
		Jawaban:           req.Jawaban,
		TanggalPeriksa:    tanggalPeriksa,
		CreatedAt:         now,
		UpdatedAt:         now,
	}

	if err := u.repo.Perawatan.CreatePerawatan(perawatan); err != nil {
		return nil, err
	}

	// Reload with relations
	return u.repo.Perawatan.GetPerawatanByID(perawatan.ID)
}

// GetPerawatanByAnakIDForIbu retrieves perawatan with ownership check
func (u *perawatanUsecase) GetPerawatanByAnakIDForIbu(anakID int32, userID int32) ([]models.Perawatan, error) {
	if anakID <= 0 {
		return nil, errors.New("anak_id tidak valid")
	}

	// Verify anak exists
	exists, err := u.repo.Perawatan.IsAnakExist(anakID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("anak tidak ditemukan")
	}

	// Verify ibu owns this anak
	owned, err := u.repo.Perawatan.IsAnakOwnedByIbu(anakID, userID)
	if err != nil {
		return nil, err
	}
	if !owned {
		return nil, errors.New("Anda tidak memiliki akses ke anak ini")
	}

	return u.repo.Perawatan.GetPerawatanByAnakIDForIbu(anakID, userID)
}

// GetPerawatanByAnakIDAndRentangUsiaForIbu retrieves perawatan with ownership check and age range filter
func (u *perawatanUsecase) GetPerawatanByAnakIDAndRentangUsiaForIbu(anakID int32, rentangUsia string, userID int32) ([]models.Perawatan, error) {
	if anakID <= 0 {
		return nil, errors.New("anak_id tidak valid")
	}
	if rentangUsia == "" {
		return nil, errors.New("rentang_usia tidak boleh kosong")
	}

	// Verify anak exists
	exists, err := u.repo.Perawatan.IsAnakExist(anakID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("anak tidak ditemukan")
	}

	// Verify ibu owns this anak
	owned, err := u.repo.Perawatan.IsAnakOwnedByIbu(anakID, userID)
	if err != nil {
		return nil, err
	}
	if !owned {
		return nil, errors.New("Anda tidak memiliki akses ke anak ini")
	}

	return u.repo.Perawatan.GetPerawatanByAnakIDAndRentangUsiaForIbu(anakID, rentangUsia, userID)
}

// UpdatePerawatanForIbu updates perawatan with ownership validation
func (u *perawatanUsecase) UpdatePerawatanForIbu(id uint, req models.UpdatePerawatanRequest, userID int32) (*models.Perawatan, error) {
	if id == 0 {
		return nil, errors.New("id tidak valid")
	}

	// Get existing perawatan
	perawatan, err := u.repo.Perawatan.GetPerawatanByID(id)
	if err != nil {
		return nil, err
	}

	// Verify ibu owns this anak
	owned, err := u.repo.Perawatan.IsAnakOwnedByIbu(perawatan.AnakID, userID)
	if err != nil {
		return nil, err
	}
	if !owned {
		return nil, errors.New("Anda tidak memiliki akses ke anak ini")
	}

	// Update fields if provided
	if req.Jawaban != nil {
		perawatan.Jawaban = req.Jawaban
	}
	if strings.TrimSpace(req.TanggalPeriksa) != "" {
		parsedTanggal, err := parsePerawatanTanggalPeriksa(req.TanggalPeriksa)
		if err != nil {
			return nil, err
		}
		perawatan.TanggalPeriksa = parsedTanggal
	}

	// Save updated perawatan
	if err := u.repo.Perawatan.UpdatePerawatan(perawatan); err != nil {
		return nil, err
	}

	// Reload with relations
	return u.repo.Perawatan.GetPerawatanByID(id)
}

// DeletePerawatanForIbu soft deletes perawatan with ownership validation
func (u *perawatanUsecase) DeletePerawatanForIbu(id uint, userID int32) error {
	if id == 0 {
		return errors.New("id tidak valid")
	}

	// Get existing perawatan
	perawatan, err := u.repo.Perawatan.GetPerawatanByID(id)
	if err != nil {
		return err
	}

	// Verify ibu owns this anak
	owned, err := u.repo.Perawatan.IsAnakOwnedByIbu(perawatan.AnakID, userID)
	if err != nil {
		return err
	}
	if !owned {
		return errors.New("Anda tidak memiliki akses ke anak ini")
	}

	return u.repo.Perawatan.DeletePerawatan(id)
}
