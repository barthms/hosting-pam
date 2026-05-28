package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
	"time"
)

type PengukuranLilAUseCase interface {
	Create(req models.CreatePengukuranLilARequest) error
	Update(id int32, req models.UpdatePengukuranLilARequest) error
	GetByAnakID(anakID int32) ([]models.PengukuranLila, error)
	GetByAnakIDForIbu(anakID int32, userID uint) ([]models.PengukuranLila, error)
	GetByID(id int32) (*models.PengukuranLila, error)
	GetAll() ([]models.PengukuranLila, error)
	Delete(id int32) error
}

type pengukuranlilaUseCase struct {
	repo repositories.PengukuranLilaRepository
}

func NewPengukuranLilAUseCase(repo repositories.PengukuranLilaRepository) PengukuranLilAUseCase {
	return &pengukuranlilaUseCase{repo: repo}
}

func (u *pengukuranlilaUseCase) Create(req models.CreatePengukuranLilARequest) error {

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

	kategori := req.KategoriRisiko
	if kategori == "" {
		// Auto-klasifikasi berdasarkan usia dan hasil LILA
		switch {
		case req.HasilLila <= 0:
			kategori = ""
		case req.Bulanke < 6:
			if req.HasilLila < 9.5 {
				kategori = "Gizi Buruk"
			} else if req.HasilLila < 11.5 {
				kategori = "Gizi Kurang"
			} else {
				kategori = "Normal"
			}
		default:
			if req.HasilLila < 11.5 {
				kategori = "Gizi Buruk"
			} else if req.HasilLila < 12.5 {
				kategori = "Gizi Kurang"
			} else {
				kategori = "Normal"
			}
		}
	}

	pemeriksaan := models.PengukuranLila{
		AnakID:         req.AnakID,
		Bulanke:        req.Bulanke,
		Tanggal:        tgl,
		HasilLila:      req.HasilLila,
		KategoriRisiko: kategori,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	return u.repo.Create(&pemeriksaan)
}
func (u *pengukuranlilaUseCase) Update(id int32, req models.UpdatePengukuranLilARequest) error {
	now := time.Now()
	return u.repo.Update(id, req, now)
}
func (u *pengukuranlilaUseCase) GetByAnakID(anakID int32) ([]models.PengukuranLila, error) {
	return u.repo.GetByAnakID(anakID)
}

func (u *pengukuranlilaUseCase) GetByAnakIDForIbu(anakID int32, userID uint) ([]models.PengukuranLila, error) {
	ok, err := u.repo.IsAnakMilikIbu(userID, anakID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("akses ditolak: anak tidak ditemukan atau bukan milik anda")
	}
	return u.repo.GetByAnakID(anakID)
}

func (u *pengukuranlilaUseCase) GetByID(id int32) (*models.PengukuranLila, error) {
	return u.repo.GetByID(id)
}

func (u *pengukuranlilaUseCase) GetAll() ([]models.PengukuranLila, error) {
	return u.repo.GetAll()
}
func (u *pengukuranlilaUseCase) Delete(id int32) error {
	return u.repo.Delete(id)
}
