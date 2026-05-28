package usecases

import (
	"errors"
	"fmt"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
	"time"
)

type NeonatusUsecase interface {
	Create(req models.CreatePelayananNeonatusRequest) error
	GetByAnakID(anakID int32) ([]models.Neonatus, error)
	GetByID(id int32)(*models.Neonatus, error)
	GetAll()([]models.Neonatus, error)
	Update(id int32, req models.UpdatePelayananNeonatusRequest) error 
	Delete(id int32) error
}

type pelayananNeonatusUseCase struct{
	NeonatusRepo repositories.PelayananNeonatusRepository
}


func NewPelayananNeonatusUseCase(repo repositories.PelayananNeonatusRepository) NeonatusUsecase{
	return &pelayananNeonatusUseCase{NeonatusRepo: repo}
}

func (uc *pelayananNeonatusUseCase)Create(req models.CreatePelayananNeonatusRequest) error{

	if req.AnakID ==0 {
		return errors.New("anak_id wajib di isi")
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


	neonatus := models.Neonatus{
		AnakID:    req.AnakID,
		PeriodeID: req.PeriodeID,
		KategoriUmurID:  req.KategoriUmurID,
		Tanggal:   tanggal,
		TenagaKesehatanID:    req.TenagaKesehatanID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	details := make([]models.DetailPelayananNeonatus, 0 , len(req.DetailPelayanan))

	for _, d:= range req.DetailPelayanan{
		details = append(details, models.DetailPelayananNeonatus{
			JenisPelayananID: d.JenisPelayananID,
			Nilai:            d.Nilai,
			Keterangan:       d.Keterangan,
			CreatedAt:        now,
			UpdatedAt:        now,
		})
	}
	fmt.Printf("Detail Final : %+v\n", details)

	neonatus.DetailPelayanan = details


	return uc.NeonatusRepo.Create(&neonatus)

}

func (uc *pelayananNeonatusUseCase) GetByAnakID(anakID int32) ([]models.Neonatus, error) {
	if anakID == 0 {
		return nil, errors.New("anak_id wajib diisi")
	}

	return uc.NeonatusRepo.GetByAnakID(anakID)
}

func (uc *pelayananNeonatusUseCase) GetByID(id int32) (*models.Neonatus, error) {
	if id == 0 {
		return nil, errors.New("id wajib diisi")
	}

	return uc.NeonatusRepo.GetByID(id)
}

func (uc *pelayananNeonatusUseCase) GetAll() ([]models.Neonatus, error) {
	return uc.NeonatusRepo.GetAll()
}

func (uc *pelayananNeonatusUseCase) Update(id int32, req models.UpdatePelayananNeonatusRequest) error {

	tanggal, err := time.Parse("2006-01-02", req.Tanggal)
	if err != nil && req.Tanggal != "" {
		return fmt.Errorf("format tanggal harus YYYY-MM-DD")
	}

	now := time.Now()

	return uc.NeonatusRepo.Update(id, req, tanggal, now)
}


func (uc *pelayananNeonatusUseCase) Delete(id int32) error {
	if id == 0 {
		return errors.New("id wajib di isi")
	}
	return uc.NeonatusRepo.Delete(id)

}
