package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type CatatanPelayananTrimester3Usecase interface {
	Create(c *models.CatatanPelayananTrimester3) (*models.CatatanPelayananTrimester3, error)
	GetByID(id int32) (*models.CatatanPelayananTrimester3, error)
	GetByKehamilanID(kehamilanID int32) ([]models.CatatanPelayananTrimester3, error)
	Update(c *models.CatatanPelayananTrimester3) (*models.CatatanPelayananTrimester3, error)
	Delete(id int32) error
}

type catatanPelayananTrimester3Usecase struct {
	repo *repositories.CatatanPelayananTrimester3Repository
}

func NewCatatanPelayananTrimester3Usecase(repo *repositories.CatatanPelayananTrimester3Repository) CatatanPelayananTrimester3Usecase {
	return &catatanPelayananTrimester3Usecase{repo: repo}
}

// Create menambah catatan baru dan mengembalikan data yang sudah tersimpan dengan ID dari database
func (u *catatanPelayananTrimester3Usecase) Create(c *models.CatatanPelayananTrimester3) (*models.CatatanPelayananTrimester3, error) {
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
func (u *catatanPelayananTrimester3Usecase) GetByID(id int32) (*models.CatatanPelayananTrimester3, error) {
	return u.repo.FindByID(id)
}

// GetByKehamilanID mengambil semua catatan untuk satu kehamilan
func (u *catatanPelayananTrimester3Usecase) GetByKehamilanID(kehamilanID int32) ([]models.CatatanPelayananTrimester3, error) {
	return u.repo.FindByKehamilanID(kehamilanID)
}

// Update memperbarui catatan dan mengembalikan data yang sudah diperbarui dari database
func (u *catatanPelayananTrimester3Usecase) Update(c *models.CatatanPelayananTrimester3) (*models.CatatanPelayananTrimester3, error) {
	_, err := u.repo.FindByID(c.IDCatatanT3)
	if err != nil {
		return nil, errors.New("data catatan pelayanan trimester 3 tidak ditemukan")
	}

	// Repository Update akan memperbarui dan mengembalikan data yang baru dari database
	updatedData, err := u.repo.Update(c)
	if err != nil {
		return nil, err
	}

	return updatedData, nil
}

// Delete menghapus catatan berdasarkan ID
func (u *catatanPelayananTrimester3Usecase) Delete(id int32) error {
	return u.repo.Delete(id)
}
