package repositories

import (
	"errors"
	"monitoring-service/app/models"
	"time"

	"gorm.io/gorm"
)

type RiwayatImunisasiResult struct {
	TanggalDiberikan time.Time
}

func (m *Main) GetAnakByUserID(userID int32) ([]models.Anak, error) {
	var anaks []models.Anak

	err := m.postgres.
		Table("anak a").
		Select("a.*").
		Joins("JOIN kehamilan k ON k.id = a.kehamilan_id").
		Joins("JOIN ibu i ON i.id = k.ibu_id").
		Joins("JOIN pengguna p ON p.penduduk_id = i.penduduk_id").
		Where("p.id = ?", userID).
		Find(&anaks).Error

	if err != nil {
		return nil, err
	}

	return anaks, nil
}

func (m *Main) GetAturanVaksinAnak() ([]models.AturanVaksinAnak, error) {
	var aturan []models.AturanVaksinAnak

	err := m.postgres.
		Where("deleted_at IS NULL").
		Order("id ASC").
		Find(&aturan).Error

	if err != nil {
		return nil, err
	}

	return aturan, nil
}

func (m *Main) IsJadwalExist(
	anakID int32,
	dosisID int64,
) (bool, error) {

	var count int64

	err := m.postgres.
		Model(&models.JadwalImunisasiAnak{}).
		Where("id_anak = ?", anakID).
		Where("id_dosis_vaksin = ?", dosisID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (m *Main) GetRiwayatImunisasi(
	anakID int32,
	dosisID int64,
) (*RiwayatImunisasiResult, error) {

	var result RiwayatImunisasiResult

	err := m.postgres.
		Table("catatan_imunisasi_anak cia").
		Select("cia.tanggal_diberikan").
		Joins("JOIN jadwal_imunisasi_anak jia ON jia.id = cia.jadwal_imunisasi_anak_id").
		Where("jia.anak_id = ?", anakID).
		Where("jia.id_dosis_vaksin = ?", dosisID).
		Order("cia.tanggal_diberikan DESC").
		First(&result).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	return &result, nil
}

func (m *Main) CreateJadwalImunisasiAnak(
	jadwal *models.JadwalImunisasiAnak,
) error {

	return m.postgres.Create(jadwal).Error
}

func (m *Main) UpdateJadwalStatus() error {
	var jadwals []models.JadwalImunisasiAnak

	if err := m.postgres.Find(&jadwals).Error; err != nil {
		return err
	}

	today := time.Now()

	for _, jadwal := range jadwals {

		diff := int(jadwal.TanggalEstimasi.Sub(today).Hours() / 24)

		var statusID int32

		switch {
		case diff >= -7 && diff < 0:
			statusID = 1 // mendekati
		case diff == 0:
			statusID = 2 // jatuh tempo
		case diff >= -3:
			statusID = 3 // terlewat
		case diff >= -13:
			statusID  = 4 // terlambat
		default:
			statusID = 5 // krisis
		}

		_ = m.postgres.
			Model(&models.JadwalImunisasiAnak{}).
			Where("id = ?", jadwal.ID).
			Update("id_status_jadwal", statusID).Error
	}

	return nil
}


// func (m *Main) CountAnakImunisasiTerlambat() (int64, error) {
// 	var count int64

// 	err := m.postgres.
// 		Table("jadwal_imunisasi_anak jia").
// 		Where("jia.tanggal_estimasi IS NOT NULL").
// 		Where("jia.tanggal_estimasi < ?", time.Now()).
// 		Where("jia.id_status_jadwal != ?", 3). 
// 		Distinct("jia.id_anak").
// 		Count(&count).Error

// 	if err != nil {
// 		return 0, err
// 	}

// 	return count, nil
// }