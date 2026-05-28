package usecases

import (
	"errors"
	"fmt"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
	"time"
)

// interface
type PelayananKesehatanAnakUseCase interface {
	Create(req models.CreatePelayananKesehatanAnakRequest) error
	GetByAnakID(anakID int32) ([]models.KunjunganAnak, error)
	GetByID(id int32) (*models.KunjunganAnak, error)
	GetAll()([]models.KunjunganAnak, error)
	Update(id int32, req models.UpdatePelayananKesehatanAnakRequest) error
	Delete(id int32) error
}

// struct implementasi
type pelayananKesehatanAnakUseCase struct {
	pelayananRepo repositories.PelayananKesehatanAnakRepository
}

// constructor
func NewPelayananKesehatanAnakUseCase(repo repositories.PelayananKesehatanAnakRepository) PelayananKesehatanAnakUseCase {
	return &pelayananKesehatanAnakUseCase{pelayananRepo: repo}
}

func (uc *pelayananKesehatanAnakUseCase) Create(req models.CreatePelayananKesehatanAnakRequest) error {

	if req.AnakID == 0 {
		return errors.New("anak_id wajib diisi")
	}
	if req.PeriodeID == 0 {
		return errors.New("periode_id wajib diisi")
	}
	if req.Tanggal == "" {
		return errors.New("tanggal wajib diisi")
	}
	if len(req.DetailPelayanan) == 0 {
		return errors.New("detail pelayanan wajib diisi")
	}

	tanggal, err := time.Parse("2006-01-02", req.Tanggal)
	if err != nil {
		return fmt.Errorf("format tanggal harus YYYY-MM-DD")
	}

	now := time.Now()

	kunjungan := models.KunjunganAnak{
		AnakID:    req.AnakID,
		PeriodeID: req.PeriodeID,
		KategoriUmurID:  req.KategoriUmurID,
		Tanggal:   tanggal,
		Lokasi:    req.Lokasi,
		CreatedAt: now,
		UpdatedAt: now,
	}

	details := make([]models.DetailPelayanan, 0, len(req.DetailPelayanan))

	for _, d := range req.DetailPelayanan {
		details = append(details, models.DetailPelayanan{
			JenisPelayananID: d.JenisPelayananID,
			Nilai:            d.Nilai,
			Keterangan:       d.Keterangan,
			CreatedAt:        now,
			UpdatedAt:        now,
		})
	}
	fmt.Printf("DETAIL FINAL: %+v\n", details)

	kunjungan.DetailPelayanan = details

	return uc.pelayananRepo.Create(&kunjungan)
}
func (uc *pelayananKesehatanAnakUseCase) GetByAnakID(anakID int32) ([]models.KunjunganAnak, error) {
	if anakID == 0 {
		return nil, errors.New("anak_id wajib diisi")
	}

	return uc.pelayananRepo.GetByAnakID(anakID)
}

func (uc *pelayananKesehatanAnakUseCase) GetByID(id int32) (*models.KunjunganAnak, error) {
	if id == 0 {
		return nil, errors.New("id wajib diisi")
	}

	return uc.pelayananRepo.GetByID(id)
}

func (uc *pelayananKesehatanAnakUseCase) GetAll() ([]models.KunjunganAnak, error) {
	return uc.pelayananRepo.GetAll()
}
func (uc *pelayananKesehatanAnakUseCase) Update(id int32, req models.UpdatePelayananKesehatanAnakRequest) error {

	tanggal, err := time.Parse("2006-01-02", req.Tanggal)
	if err != nil && req.Tanggal != "" {
		return fmt.Errorf("format tanggal harus YYYY-MM-DD")
	}

	now := time.Now()

	return uc.pelayananRepo.Update(id, req, tanggal, now)
}

func (uc *pelayananKesehatanAnakUseCase) Delete(id int32) error {
	if id == 0 {
		return errors.New("id wajib di isi")
	}
	return uc.pelayananRepo.Delete(id)

}
