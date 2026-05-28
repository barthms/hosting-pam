package repositories

import "monitoring-service/pkg/customerror"

func (m *Main) IsAnakMilikOrangtua(userID, anakID uint) (bool, error) {
	var count int64
	err := m.postgres.Table("anak a").
		Joins("LEFT JOIN ibu i ON i.id = a.ibu_id").
		Joins("LEFT JOIN kependudukan ki ON ki.id = i.kependudukan_id").
		Joins("LEFT JOIN kartu_keluarga kk ON kk.id = ki.no_kartu_keluarga_id").
		Where("a.id = ?", anakID).
		Where("kk.user_id = ?", userID).
		Count(&count).Error
	if err != nil {
		return false, customerror.NewInternalServiceError("gagal memverifikasi kepemilikan data anak")
	}

	return count > 0, nil
}

func (m *Main) IsCatatanMilikOrangtua(userID, catatanID uint) (bool, error) {
	var count int64
	err := m.postgres.Table("catatan_pertumbuhan cp").
		Joins("JOIN anak a ON a.id = cp.anak_id").
		Joins("LEFT JOIN ibu i ON i.id = a.ibu_id").
		Joins("LEFT JOIN kependudukan ki ON ki.id = i.kependudukan_id").
		Joins("LEFT JOIN kartu_keluarga kk ON kk.id = ki.no_kartu_keluarga_id").
		Where("cp.id = ?", catatanID).
		Where("kk.user_id = ?", userID).
		Count(&count).Error
	if err != nil {
		return false, customerror.NewInternalServiceError("gagal memverifikasi kepemilikan data catatan")
	}

	return count > 0, nil
}
