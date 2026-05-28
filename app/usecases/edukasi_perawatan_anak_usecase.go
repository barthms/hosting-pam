package usecases

import (
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type EdukasiPerawatanAnakUseCase interface {
	Create(data *models.EdukasiPerawatanAnak) error
	Update(id uint, data *models.EdukasiPerawatanAnak) error
	Delete(id uint) error
	GetAll() ([]models.EdukasiPerawatanAnak, error)
	GetByID(id uint) (*models.EdukasiPerawatanAnak, error)
}

type edukasiPerawatanAnakUseCase struct {
	repo repositories.EdukasiPerawatanAnakRepository
}

func NewEdukasiPerawatanAnakUseCase(repo repositories.EdukasiPerawatanAnakRepository) EdukasiPerawatanAnakUseCase {
	return &edukasiPerawatanAnakUseCase{repo}
}

func (u *edukasiPerawatanAnakUseCase) Create(data *models.EdukasiPerawatanAnak) error {
	return u.repo.Create(data)
}

func (u *edukasiPerawatanAnakUseCase) Update(id uint, data *models.EdukasiPerawatanAnak) error {
	existing, err := u.repo.FindByID(id)
	if err != nil {
		return err
	}
	existing.Judul = data.Judul
	existing.GambarURL = data.GambarURL
	existing.IsiKonten = data.IsiKonten
	return u.repo.Update(existing)
}

func (u *edukasiPerawatanAnakUseCase) Delete(id uint) error {
	return u.repo.Delete(id)
}

func (u *edukasiPerawatanAnakUseCase) GetAll() ([]models.EdukasiPerawatanAnak, error) {
	return u.repo.FindAll()
}

func (u *edukasiPerawatanAnakUseCase) GetByID(id uint) (*models.EdukasiPerawatanAnak, error) {
	return u.repo.FindByID(id)
}
