package usecases

import (
	"strings"

	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
	"monitoring-service/pkg/customerror"

	"gorm.io/gorm"
)

type AdminCreateBidanRequest struct {
	PendudukID int32  `json:"penduduk_id"`
	NoSTR      string `json:"no_str"`
	NoSIPB     string `json:"no_sipb"`
	Status     string `json:"status"`
}

type AdminUpdateBidanRequest struct {
	NoSTR  string `json:"no_str"`
	NoSIPB string `json:"no_sipb"`
	Status string `json:"status"`
}

type AdminCreateKaderRequest struct {
	PendudukID int32  `json:"penduduk_id"`
	PosyanduID *int64 `json:"posyandu_id"`
	Status     string `json:"status"`
}

type AdminUpdateKaderRequest struct {
	PosyanduID *int64 `json:"posyandu_id"`
	Status     string `json:"status"`
}

type AdminCreatePosyanduRequest struct {
	IDPuskesmas int32  `json:"id_puskesmas" form:"id_puskesmas" binding:"required"`
	Nama        string `json:"nama" form:"nama" binding:"required"`
	Alamat      string `json:"alamat" form:"alamat"`
}

type AdminUpdatePosyanduRequest struct {
	Nama   string `json:"nama" form:"nama"`
	Alamat string `json:"alamat" form:"alamat"`
}

type AdminTenagaKesehatanUsecase struct {
	bidanRepo        *repositories.BidanRepository
	kaderRepo        *repositories.KaderRepository
	kependudukanRepo *repositories.KependudukanRepository
	userRepo         *repositories.UserRepository
	roleRepo         *repositories.RoleRepository
}

func NewAdminTenagaKesehatanUsecase(
	bidanRepo *repositories.BidanRepository,
	kaderRepo *repositories.KaderRepository,
	kependudukanRepo *repositories.KependudukanRepository,
	userRepo *repositories.UserRepository,
	roleRepo *repositories.RoleRepository,
) *AdminTenagaKesehatanUsecase {
	return &AdminTenagaKesehatanUsecase{
		bidanRepo:        bidanRepo,
		kaderRepo:        kaderRepo,
		kependudukanRepo: kependudukanRepo,
		userRepo:         userRepo,
		roleRepo:         roleRepo,
	}
}

func normalizeStatus(status string) string {
	if strings.TrimSpace(status) == "" {
		return "aktif"
	}
	return strings.ToLower(strings.TrimSpace(status))
}

func validateStatus(status string) error {
	if status != "aktif" && status != "nonaktif" {
		return customerror.NewBadRequestError("status harus aktif atau nonaktif")
	}
	return nil
}

func isNotFound(err error) bool {
	return err != nil && err == gorm.ErrRecordNotFound
}

func (u *AdminTenagaKesehatanUsecase) CreateBidan(req *AdminCreateBidanRequest) (*models.Bidan, error) {
	if req == nil {
		return nil, customerror.NewBadRequestError("request tidak valid")
	}
	if req.PendudukID == 0 {
		return nil, customerror.NewBadRequestError("penduduk_id wajib diisi")
	}

	status := normalizeStatus(req.Status)
	if err := validateStatus(status); err != nil {
		return nil, err
	}

	if _, err := u.bidanRepo.FindByPendudukID(req.PendudukID); err == nil {
		return nil, customerror.NewConflictError("penduduk sudah terdaftar sebagai bidan")
	} else if !isNotFound(err) {
		return nil, customerror.NewInternalServiceError("gagal memvalidasi data bidan")
	}

	if _, err := u.kaderRepo.FindByPendudukID(req.PendudukID); err == nil {
		return nil, customerror.NewConflictError("penduduk sudah terdaftar sebagai kader")
	} else if !isNotFound(err) {
		return nil, customerror.NewInternalServiceError("gagal memvalidasi data kader")
	}

	if _, err := u.kependudukanRepo.FindByID(req.PendudukID); err != nil {
		return nil, customerror.NewNotFoundError("penduduk tidak ditemukan")
	}

	data := &models.Bidan{
		PendudukID: req.PendudukID,
		NoSTR:      strings.TrimSpace(req.NoSTR),
		NoSIPB:     strings.TrimSpace(req.NoSIPB),
		Status:     status,
	}

	if data.NoSIPB == "" {
		return nil, customerror.NewBadRequestError("no_sipb wajib diisi")
	}

	if softDeleted, err := u.bidanRepo.FindAnyByPendudukID(req.PendudukID); err == nil {
		if softDeleted.DeletedAt != nil {
			softDeleted.NoSTR = data.NoSTR
			softDeleted.NoSIPB = data.NoSIPB
			softDeleted.Status = data.Status
			softDeleted.DeletedAt = nil

			if err := u.bidanRepo.Update(softDeleted); err != nil {
				return nil, customerror.NewInternalServiceError("gagal memulihkan data bidan")
			}
			return softDeleted, nil
		}
	} else if !isNotFound(err) {
		return nil, customerror.NewInternalServiceError("gagal memvalidasi data bidan")
	}

	if err := u.bidanRepo.Create(data); err != nil {
		if _, findErr := u.bidanRepo.FindByPendudukID(req.PendudukID); findErr == nil {
			return nil, customerror.NewConflictError("penduduk sudah terdaftar sebagai bidan")
		}
		return nil, customerror.NewInternalServiceError("gagal menyimpan data bidan: " + err.Error())
	}
	return data, nil
}

func (u *AdminTenagaKesehatanUsecase) ListBidan(desa string) ([]repositories.BidanListItem, error) {
	return u.bidanRepo.List(strings.TrimSpace(desa))
}

func (u *AdminTenagaKesehatanUsecase) GetBidanDetail(id int32) (*models.Bidan, error) {
	data, err := u.bidanRepo.FindByID(id)
	if err != nil {
		return nil, customerror.NewNotFoundError("bidan tidak ditemukan")
	}
	return data, nil
}

func (u *AdminTenagaKesehatanUsecase) UpdateBidan(id int32, req *AdminUpdateBidanRequest) (*models.Bidan, error) {
	if req == nil {
		return nil, customerror.NewBadRequestError("request tidak valid")
	}

	data, err := u.bidanRepo.FindByID(id)
	if err != nil {
		return nil, customerror.NewNotFoundError("bidan tidak ditemukan")
	}

	data.NoSTR = strings.TrimSpace(req.NoSTR)
	data.NoSIPB = strings.TrimSpace(req.NoSIPB)
	if data.NoSIPB == "" {
		return nil, customerror.NewBadRequestError("no_sipb wajib diisi")
	}

	if strings.TrimSpace(req.Status) != "" {
		status := normalizeStatus(req.Status)
		if err := validateStatus(status); err != nil {
			return nil, err
		}
		data.Status = status
	}

	if err := u.bidanRepo.Update(data); err != nil {
		return nil, customerror.NewInternalServiceError("gagal memperbarui data bidan")
	}
	return data, nil
}

func (u *AdminTenagaKesehatanUsecase) SetStatusBidan(id int32, status string) error {
	status = normalizeStatus(status)
	if err := validateStatus(status); err != nil {
		return err
	}

	if _, err := u.bidanRepo.FindByID(id); err != nil {
		return customerror.NewNotFoundError("bidan tidak ditemukan")
	}

	if err := u.bidanRepo.SetStatus(id, status); err != nil {
		return customerror.NewInternalServiceError("gagal memperbarui status bidan")
	}
	return nil
}

func (u *AdminTenagaKesehatanUsecase) CreateKader(req *AdminCreateKaderRequest) (*models.Kader, error) {
	if req == nil {
		return nil, customerror.NewBadRequestError("request tidak valid")
	}
	if req.PendudukID == 0 {
		return nil, customerror.NewBadRequestError("penduduk_id wajib diisi")
	}

	status := normalizeStatus(req.Status)
	if err := validateStatus(status); err != nil {
		return nil, err
	}

	if _, err := u.kaderRepo.FindByPendudukID(req.PendudukID); err == nil {
		return nil, customerror.NewConflictError("penduduk sudah terdaftar sebagai kader")
	} else if !isNotFound(err) {
		return nil, customerror.NewInternalServiceError("gagal memvalidasi data kader")
	}

	if _, err := u.bidanRepo.FindByPendudukID(req.PendudukID); err == nil {
		return nil, customerror.NewConflictError("penduduk sudah terdaftar sebagai bidan")
	} else if !isNotFound(err) {
		return nil, customerror.NewInternalServiceError("gagal memvalidasi data bidan")
	}

	if _, err := u.kependudukanRepo.FindByID(req.PendudukID); err != nil {
		return nil, customerror.NewNotFoundError("penduduk tidak ditemukan")
	}

	data := &models.Kader{
		PendudukID: req.PendudukID,
		PosyanduID: req.PosyanduID,
		Status:     status,
	}

	if softDeleted, err := u.kaderRepo.FindAnyByPendudukID(req.PendudukID); err == nil {
		if softDeleted.DeletedAt != nil {
			softDeleted.PosyanduID = data.PosyanduID
			softDeleted.Status = data.Status
			softDeleted.DeletedAt = nil

			if err := u.kaderRepo.Update(softDeleted); err != nil {
				return nil, customerror.NewInternalServiceError("gagal memulihkan data kader")
			}
			return softDeleted, nil
		}
	} else if !isNotFound(err) {
		return nil, customerror.NewInternalServiceError("gagal memvalidasi data kader")
	}

	if err := u.kaderRepo.Create(data); err != nil {
		if _, findErr := u.kaderRepo.FindByPendudukID(req.PendudukID); findErr == nil {
			return nil, customerror.NewConflictError("penduduk sudah terdaftar sebagai kader")
		}
		return nil, customerror.NewInternalServiceError("gagal menyimpan data kader: " + err.Error())
	}
	return data, nil
}

func (u *AdminTenagaKesehatanUsecase) ListKader(desa string) ([]repositories.KaderListItem, error) {
	return u.kaderRepo.List(strings.TrimSpace(desa))
}

func (u *AdminTenagaKesehatanUsecase) GetKaderDetail(id int32) (*models.Kader, error) {
	data, err := u.kaderRepo.FindByID(id)
	if err != nil {
		return nil, customerror.NewNotFoundError("kader tidak ditemukan")
	}
	return data, nil
}

func (u *AdminTenagaKesehatanUsecase) UpdateKader(id int32, req *AdminUpdateKaderRequest) (*models.Kader, error) {
	if req == nil {
		return nil, customerror.NewBadRequestError("request tidak valid")
	}

	data, err := u.kaderRepo.FindByID(id)
	if err != nil {
		return nil, customerror.NewNotFoundError("kader tidak ditemukan")
	}

	data.PosyanduID = req.PosyanduID
	if strings.TrimSpace(req.Status) != "" {
		status := normalizeStatus(req.Status)
		if err := validateStatus(status); err != nil {
			return nil, err
		}
		data.Status = status
	}

	if err := u.kaderRepo.Update(data); err != nil {
		return nil, customerror.NewInternalServiceError("gagal memperbarui data kader")
	}
	return data, nil
}

func (u *AdminTenagaKesehatanUsecase) SetStatusKader(id int32, status string) error {
	status = normalizeStatus(status)
	if err := validateStatus(status); err != nil {
		return err
	}

	if _, err := u.kaderRepo.FindByID(id); err != nil {
		return customerror.NewNotFoundError("kader tidak ditemukan")
	}

	if err := u.kaderRepo.SetStatus(id, status); err != nil {
		return customerror.NewInternalServiceError("gagal memperbarui status kader")
	}
	return nil
}

func (u *AdminTenagaKesehatanUsecase) ListEligiblePenduduk(role, search, kecamatan, desa string) ([]repositories.EligiblePendudukItem, error) {
	return u.kependudukanRepo.ListEligibleForRole(role, search, kecamatan, desa)
}

func (u *AdminTenagaKesehatanUsecase) CreatePosyandu(req *AdminCreatePosyanduRequest) error {
	if req == nil {
		return customerror.NewBadRequestError("request tidak valid")
	}
	if req.IDPuskesmas == 0 {
		return customerror.NewBadRequestError("id_puskesmas wajib diisi")
	}
	if strings.TrimSpace(req.Nama) == "" {
		return customerror.NewBadRequestError("nama posyandu wajib diisi")
	}

	posyandu := &models.Posyandu{
		IDPuskesmas: req.IDPuskesmas,
		Nama:        strings.TrimSpace(req.Nama),
		Alamat:      strings.TrimSpace(req.Alamat),
	}

	if err := u.kependudukanRepo.CreatePosyandu(posyandu); err != nil {
		return customerror.NewInternalServiceError("gagal menyimpan data posyandu: " + err.Error())
	}

	return nil
}

func (u *AdminTenagaKesehatanUsecase) ListPosyandu(search string) ([]repositories.PosyanduItem, error) {
	return u.kependudukanRepo.ListPosyandu(search)
}

func (u *AdminTenagaKesehatanUsecase) ListPenduduk(search string) ([]repositories.EligiblePendudukItem, error) {
	return u.kependudukanRepo.ListEligibleForRole("bidan", search, "", "")
}

func (u *AdminTenagaKesehatanUsecase) GetPosyanduDetail(id int32) (*models.Posyandu, error) {
	if id == 0 {
		return nil, customerror.NewBadRequestError("id posyandu tidak valid")
	}

	posyandu, err := u.kependudukanRepo.FindPosyanduByID(id)
	if err != nil {
		return nil, customerror.NewNotFoundError("posyandu tidak ditemukan")
	}

	return posyandu, nil
}

func (u *AdminTenagaKesehatanUsecase) UpdatePosyandu(id int32, req *AdminUpdatePosyanduRequest) (*models.Posyandu, error) {
	if req == nil {
		return nil, customerror.NewBadRequestError("request tidak valid")
	}

	posyandu, err := u.kependudukanRepo.FindPosyanduByID(id)
	if err != nil {
		return nil, customerror.NewNotFoundError("posyandu tidak ditemukan")
	}

	if strings.TrimSpace(req.Nama) != "" {
		posyandu.Nama = strings.TrimSpace(req.Nama)
	}

	posyandu.Alamat = strings.TrimSpace(req.Alamat)

	if err := u.kependudukanRepo.UpdatePosyandu(posyandu); err != nil {
		return nil, customerror.NewInternalServiceError("gagal memperbarui data posyandu")
	}

	return posyandu, nil
}
