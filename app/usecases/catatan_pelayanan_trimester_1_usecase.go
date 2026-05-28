package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type CatatanPelayananTrimester1Usecase interface {
	Create(c *models.CatatanPelayananTrimester1) (*models.CatatanPelayananTrimester1, error)
	GetByID(id int32) (*models.CatatanPelayananTrimester1, error)
	GetByKehamilanID(kehamilanID int32) ([]models.CatatanPelayananTrimester1, error)
	Update(c *models.CatatanPelayananTrimester1) (*models.CatatanPelayananTrimester1, error)
	Delete(id int32) error
}

type catatanPelayananTrimester1Usecase struct {
	repo *repositories.CatatanPelayananTrimester1Repository
}

func NewCatatanPelayananTrimester1Usecase(repo *repositories.CatatanPelayananTrimester1Repository) CatatanPelayananTrimester1Usecase {
	return &catatanPelayananTrimester1Usecase{repo: repo}
}

// Create menambah catatan baru dan mengembalikan data yang sudah tersimpan dengan ID dari database
func (u *catatanPelayananTrimester1Usecase) Create(c *models.CatatanPelayananTrimester1) (*models.CatatanPelayananTrimester1, error) {
	if c.KehamilanID == 0 {
		return nil, errors.New("kehamilan_id wajib diisi")
	}

	// Repository Create akan menyimpan dan mengembalikan data dengan ID
	createdData, err := u.repo.Create(c)
	if err != nil {
		return nil, err
	}

	return createdData, nil
}

// GetByID mengambil catatan berdasarkan ID
func (u *catatanPelayananTrimester1Usecase) GetByID(id int32) (*models.CatatanPelayananTrimester1, error) {
	return u.repo.FindByID(id)
}

// GetByKehamilanID mengambil semua catatan untuk satu kehamilan
func (u *catatanPelayananTrimester1Usecase) GetByKehamilanID(kehamilanID int32) ([]models.CatatanPelayananTrimester1, error) {
	return u.repo.FindByKehamilanID(kehamilanID)
}

// Update memperbarui catatan dan mengembalikan data yang sudah diperbarui dari database
func (u *catatanPelayananTrimester1Usecase) Update(c *models.CatatanPelayananTrimester1) (*models.CatatanPelayananTrimester1, error) {
	_, err := u.repo.FindByID(c.IDCatatan)
	if err != nil {
		return nil, errors.New("data catatan pelayanan trimester 1 tidak ditemukan")
	}

	// Repository Update akan memperbarui dan mengembalikan data yang baru dari database
	updatedData, err := u.repo.Update(c)
	if err != nil {
		return nil, err
	}

	return updatedData, nil
}

// Delete menghapus catatan berdasarkan ID
func (u *catatanPelayananTrimester1Usecase) Delete(id int32) error {
	return u.repo.Delete(id)
}
