package repositories

import (
	"strings"
	"time"

	"monitoring-service/app/constants"
	"monitoring-service/app/models"
	"monitoring-service/pkg/customerror"
)

type PertumbuhanRepository interface {
	Create(data *models.CatatanPertumbuhan) error
	GetLastByAnakID(anakID uint) (*models.CatatanPertumbuhan, error)
	GetLastTwoByAnakID(anakID uint) ([]models.CatatanPertumbuhan, error)
}

func normalizeGender(gender string) string {
	v := strings.TrimSpace(strings.ToLower(gender))
	if v == "m" || v == "male" || v == "l" || strings.Contains(v, "laki") {
		return "Laki-laki"
	}
	if v == "f" || v == "female" || v == "p" || strings.Contains(v, "perem") {
		return "Perempuan"
	}
	return gender
}

func (m *Main) CreateCatatanPertumbuhan(data *models.CatatanPertumbuhan) error {
	if err := m.postgres.Create(data).Error; err != nil {
		return customerror.NewInternalServiceError("gagal membuat catatan pertumbuhan")
	}
	return nil
}

func (m *Main) GetRiwayatPertumbuhanByAnakID(anakID uint) ([]models.CatatanPertumbuhan, error) {
	var result []models.CatatanPertumbuhan
	if err := m.postgres.Where("anak_id = ?", anakID).Order("tgl_ukur ASC").Find(&result).Error; err != nil {
		return nil, customerror.NewInternalServiceError("gagal mengambil riwayat pertumbuhan")
	}
	return result, nil
}

func (m *Main) GetCatatanPertumbuhanByID(id uint) (*models.CatatanPertumbuhan, error) {
	var result models.CatatanPertumbuhan
	err := m.postgres.Where("id = ?", id).First(&result).Error
	if err != nil {
		if err.Error() == constants.GORM_ERR_NOT_FOUND {
			return nil, customerror.NewNotFoundError("catatan pertumbuhan tidak ditemukan")
		}
		return nil, customerror.NewInternalServiceError("gagal mengambil catatan pertumbuhan")
	}
	return &result, nil
}

func (m *Main) UpdateCatatanPertumbuhan(data *models.CatatanPertumbuhan) error {
	if err := m.postgres.Save(data).Error; err != nil {
		return customerror.NewInternalServiceError("gagal mengubah catatan pertumbuhan")
	}
	return nil
}

func (m *Main) DeleteCatatanPertumbuhan(id uint) error {
	if err := m.postgres.Model(&models.CatatanPertumbuhan{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error; err != nil {
		return customerror.NewInternalServiceError("gagal menghapus catatan pertumbuhan")
	}
	return nil
}
