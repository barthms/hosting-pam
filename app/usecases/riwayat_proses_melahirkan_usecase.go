// package usecases

// import (
// 	"errors"
// 	"monitoring-service/app/models"
// 	"monitoring-service/app/repositories"
// )

// type RiwayatProsesMelahirkanUsecase interface {
// 	Create(rp *models.RiwayatProsesMelahirkan) error
// 	GetByID(id int32) (*models.RiwayatProsesMelahirkan, error)
// 	GetByKehamilanID(kehamilanID int32) ([]models.RiwayatProsesMelahirkan, error)
// 	Update(rp *models.RiwayatProsesMelahirkan) error
// 	Delete(id int32) error
// }

// type riwayatProsesMelahirkanUsecase struct {
// 	repo *repositories.RiwayatProsesMelahirkanRepository
// }

// func NewRiwayatProsesMelahirkanUsecase(repo *repositories.RiwayatProsesMelahirkanRepository) RiwayatProsesMelahirkanUsecase {
// 	return &riwayatProsesMelahirkanUsecase{repo: repo}
// }

// func (u *riwayatProsesMelahirkanUsecase) Create(rp *models.RiwayatProsesMelahirkan) error {
// 	if rp.KehamilanID == 0 {
// 		return errors.New("kehamilan_id wajib diisi")
// 	}
// 	return u.repo.Create(rp)
// }

// func (u *riwayatProsesMelahirkanUsecase) GetByID(id int32) (*models.RiwayatProsesMelahirkan, error) {
// 	return u.repo.FindByID(id)
// }

// func (u *riwayatProsesMelahirkanUsecase) GetByKehamilanID(kehamilanID int32) ([]models.RiwayatProsesMelahirkan, error) {
// 	return u.repo.FindByKehamilanID(kehamilanID)
// }

// func (u *riwayatProsesMelahirkanUsecase) Update(rp *models.RiwayatProsesMelahirkan) error {
// 	_, err := u.repo.FindByID(rp.ID)
// 	if err != nil {
// 		return errors.New("data riwayat proses melahirkan tidak ditemukan")
// 	}
// 	return u.repo.Update(rp)
// }

// func (u *riwayatProsesMelahirkanUsecase) Delete(id int32) error {
// 	return u.repo.Delete(id)
// }


package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type RiwayatProsesMelahirkanUsecase interface {
	Create(rp *models.RiwayatProsesMelahirkan) error
	GetByID(id int32) (*models.RiwayatProsesMelahirkan, error)
	GetByKehamilanID(kehamilanID int32) ([]models.RiwayatProsesMelahirkan, error)
	// GetMineByUserID — dipakai endpoint /modul-ibu/riwayat-proses-melahirkan/me
	// Mengambil data riwayat milik kehamilan aktif user (ibu) yang sedang login.
	GetMineByUserID(userID int32) (*models.RiwayatProsesMelahirkan, error)
	Update(rp *models.RiwayatProsesMelahirkan) error
	Delete(id int32) error
}

type riwayatProsesMelahirkanUsecase struct {
	repo *repositories.RiwayatProsesMelahirkanRepository
}

func NewRiwayatProsesMelahirkanUsecase(repo *repositories.RiwayatProsesMelahirkanRepository) RiwayatProsesMelahirkanUsecase {
	return &riwayatProsesMelahirkanUsecase{repo: repo}
}

func (u *riwayatProsesMelahirkanUsecase) Create(rp *models.RiwayatProsesMelahirkan) error {
	if rp.KehamilanID == 0 {
		return errors.New("kehamilan_id wajib diisi")
	}
	return u.repo.Create(rp)
}

func (u *riwayatProsesMelahirkanUsecase) GetByID(id int32) (*models.RiwayatProsesMelahirkan, error) {
	return u.repo.FindByID(id)
}

func (u *riwayatProsesMelahirkanUsecase) GetByKehamilanID(kehamilanID int32) ([]models.RiwayatProsesMelahirkan, error) {
	return u.repo.FindByKehamilanID(kehamilanID)
}

// GetMineByUserID — cari kehamilan aktif user, lalu ambil riwayat proses melahirkan terakhirnya.
// Logika: kehamilan aktif bisa saja sudah selesai (pasca-lahir), sehingga kita ambil kehamilan
// paling terakhir milik ibu tanpa filter status — karena riwayat ini memang dicatat pasca bersalin.
func (u *riwayatProsesMelahirkanUsecase) GetMineByUserID(userID int32) (*models.RiwayatProsesMelahirkan, error) {
	if userID == 0 {
		return nil, errors.New("user_id tidak valid")
	}

	kehamilan, err := u.repo.FindLatestKehamilanByUserID(userID)
	if err != nil {
		return nil, errors.New("data kehamilan tidak ditemukan")
	}

	data, err := u.repo.FindLatestByKehamilanID(kehamilan.ID)
	if err != nil {
		return nil, errors.New("riwayat proses melahirkan belum tersedia")
	}

	return data, nil
}

func (u *riwayatProsesMelahirkanUsecase) Update(rp *models.RiwayatProsesMelahirkan) error {
	_, err := u.repo.FindByID(rp.ID)
	if err != nil {
		return errors.New("data riwayat proses melahirkan tidak ditemukan")
	}
	return u.repo.Update(rp)
}

func (u *riwayatProsesMelahirkanUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}