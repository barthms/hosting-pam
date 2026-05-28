package usecases

import (
	"fmt"
	"monitoring-service/app/models"
	"time"
)

func (m *Main) GetJadwalImunisasi(
	userID int32,
) ([]models.JadwalImunisasiResponse, error) {

	rows, err :=
		m.repository.
			GetJadwalImunisasiByUserID(
				userID,
			)

	if err != nil {
		return nil, err
	}

	anakMap :=
		make(
			map[int32]*models.JadwalImunisasiResponse,
		)

	for _, row := range rows {

		if _, exists :=
			anakMap[row.AnakID]; !exists {

			anakMap[row.AnakID] =
				&models.JadwalImunisasiResponse{
					AnakID:         row.AnakID,
					NamaAnak:       row.NamaAnak,
					TanggalLahir:   row.TanggalLahir,
					JumlahTerlewat: 0,
					Jadwal:         []models.JadwalImunisasiItem{},
				}
		}

		if row.JadwalID != 0 {

			// Count imunisasi terlewat
			switch row.StatusID {

			// Terlewat
			case 3:
				anakMap[row.AnakID].
					JumlahTerlewat++

			// Terlambat & Krisis ikut dihitung
			case 4, 5:
				anakMap[row.AnakID].
					JumlahTerlewat++
			}

			anakMap[row.AnakID].
				Jadwal =
				append(
					anakMap[row.AnakID].Jadwal,
					models.JadwalImunisasiItem{
						JadwalID:        row.JadwalID,
						NamaDosis:       row.NamaDosis,
						TanggalEstimasi: row.TanggalEstimasi,
						Deskripsi:       row.Deskripsi,
						EfekSamping:     row.EfekSamping,
						StatusID:        row.StatusID,
						Status:          row.Status,
					},
				)
		}
	}

	response :=
		[]models.JadwalImunisasiResponse{}

	for _, anak := range anakMap {
		response =
			append(
				response,
				*anak,
			)
	}

	return response, nil
}

func (m *Main) GetJadwalImunisasiByAnakID(
	userID int32,
	anakID int32,
) ([]models.JadwalImunisasiResponse, error) {

	rows, err :=
		m.repository.GetJadwalImunisasiByAnakID(userID, anakID)

	if err != nil {
		return nil, err
	}

	anakMap :=
		make(map[int32]*models.JadwalImunisasiResponse)

	for _, row := range rows {

		if _, exists := anakMap[row.AnakID]; !exists {
			anakMap[row.AnakID] = &models.JadwalImunisasiResponse{
				AnakID:         row.AnakID,
				NamaAnak:       row.NamaAnak,
				TanggalLahir:   row.TanggalLahir,
				JumlahTerlewat: 0,
				Jadwal:         []models.JadwalImunisasiItem{},
			}
		}

		if row.JadwalID != 0 {

			switch row.StatusID {
			case 3, 4, 5:
				anakMap[row.AnakID].JumlahTerlewat++
			}

			anakMap[row.AnakID].Jadwal = append(
				anakMap[row.AnakID].Jadwal,
				models.JadwalImunisasiItem{
					JadwalID:        row.JadwalID,
					NamaDosis:       row.NamaDosis,
					TanggalEstimasi: row.TanggalEstimasi,
					Deskripsi:       row.Deskripsi,
					EfekSamping:     row.EfekSamping,
					StatusID:        row.StatusID,
					Status:          row.Status,
				},
			)
		}
	}

	response := []models.JadwalImunisasiResponse{}
	for _, anak := range anakMap {
		response = append(response, *anak)
	}

	return response, nil
}

func (m *Main) UpdateTanggalEstimasi(
	userID int32,
	jadwalID uint,
	newDate time.Time,
) error {

	// cek data exist (harus pakai userID juga)
	data, err := m.repository.GetJadwalImunisasiByJadwalID(userID, jadwalID)
	if err != nil {
		return err
	}

	// kalau tidak ditemukan
	if data == nil || data.JadwalID == 0 {
		return fmt.Errorf("jadwal tidak ditemukan")
	}

	// update langsung
	return m.repository.UpdateTanggalEstimasi(jadwalID, newDate)
}

func (m *Main) GetJadwalImunisasiByJadwalID(
	userID int32,
	jadwalID uint,
) (*models.JadwalImunisasiResponse, error) {

	row, err := m.repository.GetJadwalImunisasiByJadwalID(userID, jadwalID)
	if err != nil {
		return nil, err
	}

	if row == nil || row.JadwalID == 0 {
		return nil, nil
	}

	result := &models.JadwalImunisasiResponse{
		AnakID:         row.AnakID,
		NamaAnak:       row.NamaAnak,
		TanggalLahir:   row.TanggalLahir,
		JumlahTerlewat: 0,
		Jadwal: []models.JadwalImunisasiItem{
			{
				JadwalID:        row.JadwalID,
				NamaDosis:       row.NamaDosis,
				TanggalEstimasi: row.TanggalEstimasi,
				Deskripsi:       row.Deskripsi,
				EfekSamping:     row.EfekSamping,
				StatusID:        row.StatusID,
				Status:          row.Status,
			},
		},
	}

	// hitung terlewat
	switch row.StatusID {
	case 3, 4, 5:
		result.JumlahTerlewat = 1
	}

	return result, nil
}
