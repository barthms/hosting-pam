package usecases

import (
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
	"monitoring-service/pkg/customerror"
	"strings"
	"time"
)

// KaderUsecase - Interface untuk usecase kader
type KaderUsecase interface {
	// Bidan operations
	GetMyKaderList(bidanPosyanduID *int64, searchKeyword string) ([]repositories.KaderListItem, error)
	GetKaderDetail(id int32) (*models.Kader, error)
	UpdateMyKaderProfile(id int32, req *UpdateKaderProfileRequest) (*models.Kader, error)

	// For admin only
	GetAllKader(desa string) ([]repositories.KaderListItem, error)
	CreateKader(req *CreateKaderRequest) (*models.Kader, error)
	UpdateKader(id int32, req *UpdateKaderRequest) (*models.Kader, error)
	UpdateKaderStatus(id int32, status string) (*models.Kader, error)
	DeleteKader(id int32) error
}

// CreateKaderRequest - Request untuk membuat kader baru
type CreateKaderRequest struct {
	PendudukID int32  `json:"penduduk_id"`
	PosyanduID *int64 `json:"posyandu_id,omitempty"`
	Status     string `json:"status"`
}

// UpdateKaderProfileRequest - Request untuk update profile kader
type UpdateKaderProfileRequest struct {
	PosyanduID *int64 `json:"posyandu_id,omitempty"`
	Status     string `json:"status,omitempty"`
}

// UpdateKaderRequest - Request untuk update kader oleh admin
type UpdateKaderRequest struct {
	PendudukID *int32 `json:"penduduk_id,omitempty"`
	PosyanduID *int64 `json:"posyandu_id,omitempty"`
	Status     string `json:"status,omitempty"`
}

type kaderUsecase struct {
	kaderRepo        *repositories.KaderRepository
	kependudukanRepo *repositories.KependudukanRepository
}

func NewKaderUsecase(
	kaderRepo *repositories.KaderRepository,
	kependudukanRepo *repositories.KependudukanRepository,
) KaderUsecase {
	return &kaderUsecase{
		kaderRepo:        kaderRepo,
		kependudukanRepo: kependudukanRepo,
	}
}

// GetMyKaderList - Bidan mendapatkan list kader di posyandu mereka
func (u *kaderUsecase) GetMyKaderList(bidanPosyanduID *int64, searchKeyword string) ([]repositories.KaderListItem, error) {
	if bidanPosyanduID == nil {
		return nil, customerror.NewBadRequestError("bidan tidak memiliki posyandu")
	}

	// Get all kaders di posyandu ini
	allKaders, err := u.kaderRepo.List("")
	if err != nil {
		return nil, customerror.NewInternalServiceError("gagal mengambil data kader: " + err.Error())
	}

	// Filter kaders yang ada di posyandu yang sama dengan bidan
	var result []repositories.KaderListItem
	for _, kader := range allKaders {
		if kader.PosyanduID == nil {
			continue
		}
		if *kader.PosyanduID == *bidanPosyanduID {
			// Jika ada keyword search, filter berdasarkan nama atau NIK
			if searchKeyword != "" {
				searchLower := strings.ToLower(searchKeyword)
				nameLower := strings.ToLower(kader.NamaLengkap)
				nikLower := strings.ToLower(kader.NIK)

				if strings.Contains(nameLower, searchLower) || strings.Contains(nikLower, searchLower) {
					result = append(result, kader)
				}
			} else {
				result = append(result, kader)
			}
		}
	}

	return result, nil
}

// GetKaderDetail - Mendapatkan detail kader
func (u *kaderUsecase) GetKaderDetail(id int32) (*models.Kader, error) {
	if id <= 0 {
		return nil, customerror.NewBadRequestError("id kader tidak valid")
	}

	kader, err := u.kaderRepo.FindByID(id)
	if err != nil {
		return nil, customerror.NewNotFoundError("kader tidak ditemukan")
	}

	return kader, nil
}

// UpdateMyKaderProfile - Update profile kader (untuk bidan update kader mereka)
func (u *kaderUsecase) UpdateMyKaderProfile(id int32, req *UpdateKaderProfileRequest) (*models.Kader, error) {
	if id <= 0 {
		return nil, customerror.NewBadRequestError("id kader tidak valid")
	}

	if req == nil {
		return nil, customerror.NewBadRequestError("request tidak valid")
	}

	// Get existing kader
	kader, err := u.kaderRepo.FindByID(id)
	if err != nil {
		return nil, customerror.NewNotFoundError("kader tidak ditemukan")
	}

	// Update field yang diisi
	if req.PosyanduID != nil {
		kader.PosyanduID = req.PosyanduID
	}

	if req.Status != "" {
		normalizedStatus := strings.ToLower(strings.TrimSpace(req.Status))
		if normalizedStatus != "aktif" && normalizedStatus != "nonaktif" {
			return nil, customerror.NewBadRequestError("status harus 'aktif' atau 'nonaktif'")
		}
		kader.Status = normalizedStatus
	}

	kader.UpdatedAt = time.Now()

	// Save to database
	if err := u.kaderRepo.Update(kader); err != nil {
		return nil, customerror.NewInternalServiceError("gagal update data kader: " + err.Error())
	}

	return kader, nil
}

// GetAllKader - Admin mendapatkan semua kader (dengan filter optional)
func (u *kaderUsecase) GetAllKader(desa string) ([]repositories.KaderListItem, error) {
	kaders, err := u.kaderRepo.List(desa)
	if err != nil {
		return nil, customerror.NewInternalServiceError("gagal mengambil data kader: " + err.Error())
	}

	return kaders, nil
}

// CreateKader - Admin membuat kader baru
func (u *kaderUsecase) CreateKader(req *CreateKaderRequest) (*models.Kader, error) {
	if req == nil {
		return nil, customerror.NewBadRequestError("request tidak valid")
	}

	if req.PendudukID == 0 {
		return nil, customerror.NewBadRequestError("penduduk_id wajib diisi")
	}

	// Validasi status
	normalizedStatus := strings.ToLower(strings.TrimSpace(req.Status))
	if normalizedStatus == "" {
		normalizedStatus = "aktif"
	}
	if normalizedStatus != "aktif" && normalizedStatus != "nonaktif" {
		return nil, customerror.NewBadRequestError("status harus 'aktif' atau 'nonaktif'")
	}

	// Cek apakah penduduk sudah terdaftar sebagai kader
	existing, err := u.kaderRepo.FindByPendudukID(req.PendudukID)
	if err == nil && existing != nil {
		return nil, customerror.NewConflictError("penduduk sudah terdaftar sebagai kader")
	}

	// Cek apakah penduduk ada
	penduduk, err := u.kependudukanRepo.FindByID(req.PendudukID)
	if err != nil {
		return nil, customerror.NewNotFoundError("penduduk tidak ditemukan")
	}

	newKader := &models.Kader{
		PendudukID: req.PendudukID,
		PosyanduID: req.PosyanduID,
		Status:     normalizedStatus,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := u.kaderRepo.Create(newKader); err != nil {
		return nil, customerror.NewInternalServiceError("gagal membuat kader baru: " + err.Error())
	}

	// Load penduduk relation
	newKader.Penduduk = *penduduk

	return newKader, nil
}

// DeleteKader - Admin menghapus kader (soft delete)
func (u *kaderUsecase) DeleteKader(id int32) error {
	if id <= 0 {
		return customerror.NewBadRequestError("id kader tidak valid")
	}

	// Cek apakah kader ada
	kader, err := u.kaderRepo.FindByID(id)
	if err != nil {
		return customerror.NewNotFoundError("kader tidak ditemukan")
	}

	// Soft delete
	now := time.Now()
	kader.DeletedAt = &now

	if err := u.kaderRepo.Update(kader); err != nil {
		return customerror.NewInternalServiceError("gagal menghapus kader: " + err.Error())
	}

	return nil
}

// UpdateKader - Admin mengupdate data kader
func (u *kaderUsecase) UpdateKader(id int32, req *UpdateKaderRequest) (*models.Kader, error) {
	if id <= 0 {
		return nil, customerror.NewBadRequestError("id kader tidak valid")
	}

	if req == nil {
		return nil, customerror.NewBadRequestError("request tidak valid")
	}

	// Get existing kader
	kader, err := u.kaderRepo.FindByID(id)
	if err != nil {
		return nil, customerror.NewNotFoundError("kader tidak ditemukan")
	}

	// Update penduduk_id jika ada (dengan validasi)
	if req.PendudukID != nil && *req.PendudukID > 0 {
		// Cek jika penduduk_id berbeda
		if *req.PendudukID != kader.PendudukID {
			// Cek apakah penduduk baru sudah menjadi kader
			existing, err := u.kaderRepo.FindByPendudukID(*req.PendudukID)
			if err == nil && existing != nil {
				return nil, customerror.NewConflictError("penduduk sudah terdaftar sebagai kader")
			}

			// Cek apakah penduduk baru ada
			penduduk, err := u.kependudukanRepo.FindByID(*req.PendudukID)
			if err != nil {
				return nil, customerror.NewNotFoundError("penduduk tidak ditemukan")
			}
			kader.PendudukID = *req.PendudukID
			kader.Penduduk = *penduduk
		}
	}

	// Update posyandu_id jika ada
	if req.PosyanduID != nil {
		kader.PosyanduID = req.PosyanduID
	}

	// Update status jika ada
	if req.Status != "" {
		normalizedStatus := strings.ToLower(strings.TrimSpace(req.Status))
		if normalizedStatus != "aktif" && normalizedStatus != "nonaktif" {
			return nil, customerror.NewBadRequestError("status harus 'aktif' atau 'nonaktif'")
		}
		kader.Status = normalizedStatus
	}

	kader.UpdatedAt = time.Now()

	// Save to database
	if err := u.kaderRepo.Update(kader); err != nil {
		return nil, customerror.NewInternalServiceError("gagal update data kader: " + err.Error())
	}

	return kader, nil
}

// UpdateKaderStatus - Update hanya status kader
func (u *kaderUsecase) UpdateKaderStatus(id int32, status string) (*models.Kader, error) {
	if id <= 0 {
		return nil, customerror.NewBadRequestError("id kader tidak valid")
	}

	if status == "" {
		return nil, customerror.NewBadRequestError("status tidak boleh kosong")
	}

	// Validasi status
	normalizedStatus := strings.ToLower(strings.TrimSpace(status))
	if normalizedStatus != "aktif" && normalizedStatus != "nonaktif" {
		return nil, customerror.NewBadRequestError("status harus 'aktif' atau 'nonaktif'")
	}

	// Get existing kader
	kader, err := u.kaderRepo.FindByID(id)
	if err != nil {
		return nil, customerror.NewNotFoundError("kader tidak ditemukan")
	}

	// Update status
	kader.Status = normalizedStatus
	kader.UpdatedAt = time.Now()

	// Save to database
	if err := u.kaderRepo.Update(kader); err != nil {
		return nil, customerror.NewInternalServiceError("gagal update status kader: " + err.Error())
	}

	return kader, nil
}
