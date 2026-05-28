package usecases

import (
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"

	"gorm.io/gorm"
)

type ChecklistPemantauanIbuNifasUsecase struct {
	repo          repositories.ChecklistPemantauanIbuNifasRepository
	kehamilanRepo *repositories.KehamilanRepository
}

func NewChecklistPemantauanIbuNifasUsecase(
	repo repositories.ChecklistPemantauanIbuNifasRepository,
	kehamilanRepo *repositories.KehamilanRepository,
) *ChecklistPemantauanIbuNifasUsecase {
	return &ChecklistPemantauanIbuNifasUsecase{
		repo:          repo,
		kehamilanRepo: kehamilanRepo,
	}
}

func (u *ChecklistPemantauanIbuNifasUsecase) Save(data *models.ChecklistPemantauanIbuNifas) error {
	existing, err := u.repo.GetByKehamilanIDAndHariNifas(
		data.KehamilanID,
		data.HariNifas,
	)

	if err == nil && existing != nil {
		data.ID = existing.ID
		return u.repo.Update(data)
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	return u.repo.Create(data)
}

func (u *ChecklistPemantauanIbuNifasUsecase) GetByKehamilanIDAndHariNifas(
	kehamilanID int32,
	hariNifas int32,
) (*models.ChecklistPemantauanIbuNifas, error) {
	return u.repo.GetByKehamilanIDAndHariNifas(kehamilanID, hariNifas)
}

func (u *ChecklistPemantauanIbuNifasUsecase) GetByUserIDAndHariNifas(
	userID int32,
	hariNifas int32,
) (*models.ChecklistPemantauanIbuNifas, error) {
	kehamilan, err := u.kehamilanRepo.FindAktifByUserID(userID)
	if err != nil {
		return nil, err
	}

	return u.repo.GetByKehamilanIDAndHariNifas(kehamilan.ID, hariNifas)
}

func (u *ChecklistPemantauanIbuNifasUsecase) GetFilledDaysByUserID(userID int32) ([]int32, error) {
	kehamilan, err := u.kehamilanRepo.FindAktifByUserID(userID)
	if err != nil {
		return nil, err
	}

	return u.repo.GetFilledDaysByKehamilanID(kehamilan.ID)
}