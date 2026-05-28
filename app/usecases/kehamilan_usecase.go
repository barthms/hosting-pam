// package usecases

// import (
// 	"errors"
// 	"monitoring-service/app/models"
// 	"monitoring-service/app/repositories"
// )

// type KehamilanUsecase interface {
// 	Create(kehamilan *models.Kehamilan) error
// 	GetByID(id int32) (*models.Kehamilan, error)
// 	// GetAktifByPhoneNumber(phoneNumber string) (*models.Kehamilan, error) // MODUL IBU
// 	GetAktifByUserID(userID int32) (*models.Kehamilan, error) // MODUL IBU
// 	GetByIbuID(ibuID int32) ([]models.Kehamilan, error)
// 	GetAll() ([]models.Kehamilan, error)
// 	Update(kehamilan *models.Kehamilan) error
// 	Delete(id int32) error
// }

// type kehamilanUsecase struct {
// 	repo *repositories.KehamilanRepository
// }

// func NewKehamilanUsecase(repo *repositories.KehamilanRepository) KehamilanUsecase {
// 	return &kehamilanUsecase{repo: repo}
// }

// func calculateIMT(bb, tb float64) (float64, error) {
// 	if tb <= 0 {
// 		return 0, errors.New("tb tidak boleh 0")
// 	}
// 	return bb / (tb * tb), nil
// }

// func (u *kehamilanUsecase) Create(kehamilan *models.Kehamilan) error {
// 	if kehamilan.IbuID == 0 {
// 		return errors.New("ibu_id wajib diisi")
// 	}
// 	if kehamilan.BB_Awal <= 0 {
// 		return errors.New("bb_awal tidak valid")
// 	}

// 	imt, err := calculateIMT(kehamilan.BB_Awal, kehamilan.TB)
// 	if err != nil {
// 		return err
// 	}

// 	kehamilan.IMT_Awal = imt
// 	return u.repo.Create(kehamilan)
// }

// func (u *kehamilanUsecase) GetByID(id int32) (*models.Kehamilan, error) {
// 	return u.repo.FindByID(id)
// }

// func (u *kehamilanUsecase) GetByIbuID(ibuID int32) ([]models.Kehamilan, error) {
// 	return u.repo.FindByIbuID(ibuID)
// }

// func (u *kehamilanUsecase) GetAll() ([]models.Kehamilan, error) {
// 	return u.repo.GetAll()
// }

// func (u *kehamilanUsecase) Update(kehamilan *models.Kehamilan) error {
// 	existing, err := u.repo.FindByID(kehamilan.ID)
// 	if err != nil {
// 		return errors.New("data kehamilan tidak ditemukan")
// 	}

// 	// UPDATE FIELD NON-IMT
// 	existing.Gravida = kehamilan.Gravida
// 	existing.Paritas = kehamilan.Paritas
// 	existing.Abortus = kehamilan.Abortus
// 	existing.HPHT = kehamilan.HPHT
// 	existing.TaksiranPersalinan = kehamilan.TaksiranPersalinan
// 	existing.UKKehamilanSaatIni = kehamilan.UKKehamilanSaatIni
// 	existing.JarakKehamilanSebelumnya = kehamilan.JarakKehamilanSebelumnya
// 	existing.StatusKehamilan = kehamilan.StatusKehamilan

// 	// UPDATE BB & TB
// 	if kehamilan.BB_Awal > 0 {
// 		existing.BB_Awal = kehamilan.BB_Awal
// 	}
// 	if kehamilan.TB > 0 {
// 		existing.TB = kehamilan.TB
// 	}

// 	// RECALCULATE IMT
// 	imt, err := calculateIMT(existing.BB_Awal, existing.TB)
// 	if err != nil {
// 		return err
// 	}

// 	existing.IMT_Awal = imt

// 	return u.repo.Update(existing)
// }

// func (u *kehamilanUsecase) Delete(id int32) error {
// 	return u.repo.Delete(id)
// }

// // MODUL IBU
// // func (u *kehamilanUsecase) GetAktifByPhoneNumber(phoneNumber string) (*models.Kehamilan, error) {
// // 	if phoneNumber == "" {
// // 		return nil, errors.New("nomor telepon tidak ditemukan")
// // 	}
// // 	return u.repo.FindAktifByPhoneNumber(phoneNumber)
// // }
//
//	func (u *kehamilanUsecase) GetAktifByUserID(userID int32) (*models.Kehamilan, error) {
//		if userID == 0 {
//			return nil, errors.New("user_id tidak valid")
//		}
//		return u.repo.FindAktifByUserID(userID)
//	}
package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
	"time"
)

type KehamilanUsecase interface {
	Create(kehamilan *models.Kehamilan) error
	GetByID(id int32) (*models.Kehamilan, error)
	GetAktifByUserID(userID int32) (*models.Kehamilan, error)
	GetByIbuID(ibuID int32) ([]models.Kehamilan, error)
	GetAll() ([]models.Kehamilan, error)
	Update(kehamilan *models.Kehamilan) error
	Delete(id int32) error
	// STUB BARU
	UpdateAllActiveGestationalAge() error
	GetDashboardIbuHamil(ibuID int32) (interface{}, error)
	ExistsActiveByIbuID(ibuID int32) (bool, error)
}

type kehamilanUsecase struct {
	repo *repositories.KehamilanRepository
}

func NewKehamilanUsecase(repo *repositories.KehamilanRepository) KehamilanUsecase {
	return &kehamilanUsecase{repo: repo}
}

func calculateIMT(bb, tb float64) (float64, error) {
	if tb <= 0 {
		return 0, errors.New("tb tidak boleh 0")
	}
	return bb / (tb * tb), nil
}

func HitungUKDariHPHT(hpht time.Time) int32 {
	if hpht.IsZero() {
		return 0
	}
	now := time.Now()
	if hpht.After(now) {
		return 0
	}
	days := int(now.Sub(hpht).Hours() / 24)
	weeks := int32(days / 7)
	if weeks > 42 {
		weeks = 42
	}
	return weeks
}

func (u *kehamilanUsecase) Create(kehamilan *models.Kehamilan) error {
	if kehamilan.IbuID == 0 {
		return errors.New("ibu_id wajib diisi")
	}
	if kehamilan.BB_Awal <= 0 {
		return errors.New("bb_awal tidak valid")
	}
	imt, err := calculateIMT(kehamilan.BB_Awal, kehamilan.TB)
	if err != nil {
		return err
	}
	kehamilan.IMT_Awal = imt
	if !kehamilan.HPHT.IsZero() {
		kehamilan.UKKehamilanSaatIni = HitungUKDariHPHT(kehamilan.HPHT)
	}
	return u.repo.Create(kehamilan)
}

func (u *kehamilanUsecase) GetByID(id int32) (*models.Kehamilan, error) {
	kehamilan, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if !kehamilan.HPHT.IsZero() {
		kehamilan.UKKehamilanSaatIni = HitungUKDariHPHT(kehamilan.HPHT)
	}
	return kehamilan, nil
}

func (u *kehamilanUsecase) GetByIbuID(ibuID int32) ([]models.Kehamilan, error) {
	list, err := u.repo.FindByIbuID(ibuID)
	if err != nil {
		return nil, err
	}
	for i := range list {
		if !list[i].HPHT.IsZero() {
			list[i].UKKehamilanSaatIni = HitungUKDariHPHT(list[i].HPHT)
		}
	}
	return list, nil
}

func (u *kehamilanUsecase) GetAll() ([]models.Kehamilan, error) {
	list, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}
	for i := range list {
		if !list[i].HPHT.IsZero() {
			list[i].UKKehamilanSaatIni = HitungUKDariHPHT(list[i].HPHT)
		}
	}
	return list, nil
}

func (u *kehamilanUsecase) Update(kehamilan *models.Kehamilan) error {
	existing, err := u.repo.FindByID(kehamilan.ID)
	if err != nil {
		return errors.New("data kehamilan tidak ditemukan")
	}
	existing.Gravida = kehamilan.Gravida
	existing.Paritas = kehamilan.Paritas
	existing.Abortus = kehamilan.Abortus
	existing.HPHT = kehamilan.HPHT
	existing.TaksiranPersalinan = kehamilan.TaksiranPersalinan
	existing.JarakKehamilanSebelumnya = kehamilan.JarakKehamilanSebelumnya
	existing.StatusKehamilan = kehamilan.StatusKehamilan
	if kehamilan.BB_Awal > 0 {
		existing.BB_Awal = kehamilan.BB_Awal
	}
	if kehamilan.TB > 0 {
		existing.TB = kehamilan.TB
	}
	imt, err := calculateIMT(existing.BB_Awal, existing.TB)
	if err != nil {
		return err
	}
	existing.IMT_Awal = imt
	if !existing.HPHT.IsZero() {
		existing.UKKehamilanSaatIni = HitungUKDariHPHT(existing.HPHT)
	}
	return u.repo.Update(existing)
}

func (u *kehamilanUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}

func (u *kehamilanUsecase) GetAktifByUserID(userID int32) (*models.Kehamilan, error) {
	if userID == 0 {
		return nil, errors.New("user_id tidak valid")
	}
	kehamilan, err := u.repo.FindAktifByUserID(userID)
	if err != nil {
		return nil, err
	}
	if kehamilan == nil {
		return nil, errors.New("data kehamilan aktif tidak ditemukan")
	}
	if !kehamilan.HPHT.IsZero() {
		kehamilan.UKKehamilanSaatIni = HitungUKDariHPHT(kehamilan.HPHT)
	}
	return kehamilan, nil
}

// COBA DULU
func (u *kehamilanUsecase) UpdateAllActiveGestationalAge() error {
	// TODO: Implementasi cron job update usia kehamilan
	return u.repo.UpdateAllActiveGestationalAge()
}

func (u *kehamilanUsecase) GetDashboardIbuHamil(ibuID int32) (interface{}, error) {
	// TODO: Implementasi logika dashboard ibu hamil
	return nil, nil
}

func (u *kehamilanUsecase) ExistsActiveByIbuID(ibuID int32) (bool, error) {
	// TODO: Implementasi pengecekan kehamilan aktif by Ibu ID
	return u.repo.ExistsActiveByIbuID(ibuID)
}
