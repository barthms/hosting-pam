package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
	"time"
)

type PemantauanPertumbuhanAnakUseCase interface {
	Create(req models.CreatePemantauanPemeriksaanRequest) error
	Update(id int32, req models.UpdatePemantauanPemeriksaanRequest) error
	GetByAnakID(anakID int32) ([]models.DeteksiDiniPenyimpangan, error)
	GetByID(id int32) (*models.DeteksiDiniPenyimpangan, error)
	GetAll() ([]models.DeteksiDiniPenyimpangan, error)
	Delete(id int32) error
}

type pemantauanpertumbuhanUseCase struct {
	repo repositories.PemantauanPertumbuhanRepository
}

func NewPemantauanPertumbuhanUseCase(repo repositories.PemantauanPertumbuhanRepository) PemantauanPertumbuhanAnakUseCase {
	return &pemantauanpertumbuhanUseCase{repo: repo}
}

func (u *pemantauanpertumbuhanUseCase) Create(req models.CreatePemantauanPemeriksaanRequest) error {

	if req.AnakID == 0 {
		return errors.New("anak_id wajib diisi")
	}

	if req.TenagaKesehatanID == 0 {
		return errors.New("tenaga_kesehatan_id wajib diisi")
	}

	now := time.Now()

	// Parse tanggal
	tgl := now
	if req.Tanggal != "" {
		if t, err := time.Parse(time.RFC3339, req.Tanggal); err == nil {
			tgl = t
		} else if t, err := time.Parse("2006-01-02", req.Tanggal); err == nil {
			tgl = t
		}
	}

	// Parse kunjungan ulang
	var kunjunganUlang time.Time
	if req.KunjunganUlang != "" {
		if ku, err := time.Parse(time.RFC3339, req.KunjunganUlang); err == nil {
			kunjunganUlang = ku
		} else if ku, err := time.Parse("2006-01-02", req.KunjunganUlang); err == nil {
			kunjunganUlang = ku
		}
	}

	hasilPKAT := req.HasilPKAT
	if hasilPKAT == "" {
		hasilPKAT = "Normal"
	}
	tindakan := req.Tindakan
	if tindakan == "" {
		tindakan = "Pemberian stimulasi sesuai usia"
	}

	pemeriksaan := models.DeteksiDiniPenyimpangan{
		AnakID:            req.AnakID,
		BulanKe:           req.Bulanke,
		Tanggal:           tgl,
		TenagaKesehatanID: req.TenagaKesehatanID,
		BBperU:            req.BBperU,
		BBperTB:           req.BBperTB,
		TBperU:            req.TBperU,
		LKperU:            req.LKperU,
		LILA:              req.LILA,
		KPSP:              req.KPSP,
		TDD:               req.TDD,
		TDL:               req.TDL,
		KMPE:              req.KMPE,
		MCHATRevised:      req.MCHATRevised,
		ACTRS:             req.ACTRS,
		HasilPKAT:         hasilPKAT,
		Tindakan:          tindakan,
		KunjunganUlang:    kunjunganUlang,
		CreatedAt:         now,
		UpdatedAt:         now,
	}

	return u.repo.Create(&pemeriksaan)
}
func (u *pemantauanpertumbuhanUseCase) Update(id int32, req models.UpdatePemantauanPemeriksaanRequest) error {
	now := time.Now()
	return u.repo.Update(id, req, now)
}
func (u *pemantauanpertumbuhanUseCase) GetByAnakID(anakID int32) ([]models.DeteksiDiniPenyimpangan, error) {
	return u.repo.GetByAnakID(anakID)
}

func (u *pemantauanpertumbuhanUseCase) GetByID(id int32) (*models.DeteksiDiniPenyimpangan, error) {
	return u.repo.GetByID(id)
}

func (u *pemantauanpertumbuhanUseCase) GetAll() ([]models.DeteksiDiniPenyimpangan, error) {
	return u.repo.GetAll()
}
func (u *pemantauanpertumbuhanUseCase) Delete(id int32) error {
	return u.repo.Delete(id)
}
