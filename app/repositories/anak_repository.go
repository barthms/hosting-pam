package repositories

import (
	"strings"

	"monitoring-service/app/constants"
	"monitoring-service/app/models"
	"monitoring-service/pkg/customerror"
)

func (m *Main) GetAllAnak() ([]models.Anak, error) {
	var data []models.Anak
	err := m.postgres.
		Preload("Penduduk").
		Preload("Kehamilan").
		Preload("Kehamilan.Ibu").
		Preload("Kehamilan.Ibu.Kependudukan").
		Find(&data).Error
	if err != nil {
		return nil, customerror.NewInternalServiceError("gagal mengambil data anak")
	}

	return data, nil
}

func (m *Main) SearchAnak(namaAnak, namaIbu, noKK string) ([]models.Anak, error) {
	var data []models.Anak

	query := m.postgres.Model(&models.Anak{}).
		Joins("LEFT JOIN penduduk p ON p.id = anak.penduduk_id").
		Joins("LEFT JOIN kehamilan k ON k.id = anak.kehamilan_id").
		Joins("LEFT JOIN ibu i ON i.id = k.ibu_id").
		Joins("LEFT JOIN penduduk pi ON pi.id = i.penduduk_id").
		Joins("LEFT JOIN kartu_keluarga kk ON kk.id = p.kartu_keluarga_id")

	namaAnak = strings.TrimSpace(namaAnak)
	namaIbu = strings.TrimSpace(namaIbu)
	noKK = strings.TrimSpace(noKK)

	if namaAnak != "" {
		query = query.Where("p.nama_lengkap ILIKE ?", "%"+namaAnak+"%")
	}

	if namaIbu != "" {
		query = query.Where("pi.nama_lengkap ILIKE ?", "%"+namaIbu+"%")
	}

	if noKK != "" {
		searchNoKK := "%" + noKK + "%"
		query = query.Where("kk.no_kartu_keluarga ILIKE ?", searchNoKK)
	}

	err := query.
		Preload("Penduduk").
		Preload("Kehamilan").
		Preload("Kehamilan.Ibu").
		Preload("Kehamilan.Ibu.Kependudukan").
		Order("anak.created_at DESC").
		Find(&data).Error

	if err != nil {
		return nil, customerror.NewInternalServiceError("gagal mencari data anak")
	}

	return data, nil
}

func (m *Main) GetAnakByID(anakID uint) (*models.Anak, error) {
	var data models.Anak
	err := m.postgres.
		Preload("Penduduk").
		Preload("Kehamilan").
		Preload("Kehamilan.Ibu").
		Preload("Kehamilan.Ibu.Kependudukan").
		Where("id = ?", anakID).
		First(&data).Error
	if err != nil {
		if err.Error() == constants.GORM_ERR_NOT_FOUND {
			return nil, customerror.NewNotFoundError("data anak tidak ditemukan")
		}
		// Fallback: coba load hanya dengan Preload Penduduk
		var simpleData models.Anak
		simpleErr := m.postgres.
			Preload("Penduduk").
			Where("id = ?", anakID).
			First(&simpleData).Error
		if simpleErr != nil {
			if simpleErr.Error() == constants.GORM_ERR_NOT_FOUND {
				return nil, customerror.NewNotFoundError("data anak tidak ditemukan")
			}
			return nil, customerror.NewInternalServiceError("gagal mengambil data anak")
		}
		return &simpleData, nil
	}

	return &data, nil
}
