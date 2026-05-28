package usecases

import (
	"fmt"
	"monitoring-service/app/models"
	"time"
)

func (m *Main) GenerateJadwalImunisasi(
	userID int32,
) error {

	fmt.Println("========== START GENERATE ==========")
	fmt.Println("USER ID:", userID)

	anaks, err := m.repository.GetAnakByUserID(userID)
	if err != nil {
		return err
	}
	fmt.Println("TOTAL ANAK:", len(anaks))

	for _, a := range anaks {
		fmt.Println(
			"ANAK:",
			a.ID,
			"TanggalLahir:",
			a.TanggalLahir,
		)
	}

	aturanList, err := m.repository.GetAturanVaksinAnak()
	if err != nil {
		return err
	}
	fmt.Println(
		"TOTAL RULE:",
		len(aturanList),
	)

	for _, r := range aturanList {
		fmt.Println(
			"RULE:",
			r.ID,
			r.DosisVaksinID,
			r.MinUsiaHari,
			r.MaxUsiaHari,
		)
	}

	today := time.Now()

	for _, anak := range anaks {

		if anak.TanggalLahir == nil {
			continue
		}

		umurHari := int(
			today.Sub(*anak.TanggalLahir).Hours() / 24,
		)

		fmt.Println(
			"ANAK ID:",
			anak.ID,
			"UMUR:",
			umurHari,
		)

		for _, rule := range aturanList {

			// if umurHari < int(rule.MinUsiaHari) {

			// 	fmt.Println(
			// 		"SKIP: umur kurang",
			// 		umurHari,
			// 		"<",
			// 		rule.MinUsiaHari,
			// 	)

			// 	continue
			// }

			// if rule.MaxUsiaHari > 0 &&
			// 	umurHari > int(rule.MaxUsiaHari) {
			// 	fmt.Println(
			// 		"SKIP: umur melebihi batas",
			// 	)
			// 	continue
			// }

			alreadyExist, err :=
				m.repository.IsJadwalExist(
					anak.ID,
					int64(rule.DosisVaksinID),
				)

			if err != nil {
				return err
			}

			if alreadyExist {
				fmt.Println(
					"SKIP: jadwal sudah ada",
				)
				continue
			}

			var tanggalEstimasi time.Time

			if rule.DosisSebelumnyaID != nil {
				riwayat, err :=
					m.repository.GetRiwayatImunisasi(
						anak.ID,
						int64(*rule.DosisSebelumnyaID),
					)

				if err != nil {
					continue
				}
				if rule.MinIntervalHari == 0 {
					continue
				}

				selisihHari := int(
					today.Sub(
						riwayat.TanggalDiberikan,
					).Hours() / 24,
				)
				if selisihHari < int(rule.MinIntervalHari) {
					continue
				}
				tanggalEstimasi =
					riwayat.TanggalDiberikan.AddDate(
						0,
						0,
						int(rule.MinIntervalHari),
					)
			} else {

				// if rule.MinUsiaHari == 0 {
				// 	continue
				// }

				tanggalEstimasi =
					anak.TanggalLahir.AddDate(
						0,
						0,
						int(rule.MinUsiaHari),
					)
			}
			statusID := calculateStatusID(
				tanggalEstimasi,
			)

			fmt.Println(
				"CREATE JADWAL",
				"Anak:", anak.ID,
				"Dosis:", rule.DosisVaksinID,
				"Tanggal:", tanggalEstimasi,
				"Status:", statusID,
			)

			jadwal := &models.JadwalImunisasiAnak{
				AnakID:          uint(anak.ID),
				DosisVaksinID:   rule.DosisVaksinID,
				TanggalEstimasi: &tanggalEstimasi,
				StatusJadwalID:  uint(statusID),
			}
			err =
				m.repository.
					CreateJadwalImunisasiAnak(
						jadwal,
					)

			if err != nil {
				fmt.Println(
					"ERROR INSERT:",
					err,
				)

				return err
			}
			fmt.Println(
				"SUCCESS INSERT",
			)
		}
	}

	_ = m.repository.UpdateJadwalStatus()

	return nil
}

func calculateStatusID(
	tanggalEstimasi time.Time,
) int32 {

	today := time.Now()

	today = time.Date(
		today.Year(),
		today.Month(),
		today.Day(),
		0, 0, 0, 0,
		today.Location(),
	)

	tanggalEstimasi = time.Date(
		tanggalEstimasi.Year(),
		tanggalEstimasi.Month(),
		tanggalEstimasi.Day(),
		0, 0, 0, 0,
		tanggalEstimasi.Location(),
	)

	diff := int(
		tanggalEstimasi.Sub(today).
			Hours() / 24,
	)

	switch {

	// H-7 sampai H-1
	case diff >= 1 && diff <= 7:
		return 1 // mendekati

	// Hari H
	case diff == 0:
		return 2 // jatuh tempo

	// H+1 sampai H+3
	case diff >= -3:
		return 3 // terlewat

	// H+4 sampai H+7
	case diff >= -7:
		return 4 // terlambat

	// > H+14
	default:
		return 5 // krisis
	}
}
