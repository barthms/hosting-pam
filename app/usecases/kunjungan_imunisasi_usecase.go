package usecases

import (
	"context"
	"fmt"
	"log"
	"monitoring-service/app/models"
	"strconv"
	"strings"
	"time"
)

func (m *Main) GetKunjunganImunisasiByID(
	kunjunganID uint,
) (*models.KunjunganImunisasiDetailResponse, error) {

	row, err :=
		m.repository.
			GetKunjunganImunisasiByID(
				kunjunganID,
			)

	if err != nil {
		return nil, err
	}

	if row == nil || row.KunjunganID == 0 {
		return nil, nil
	}

	result :=
		&models.KunjunganImunisasiDetailResponse{
			KunjunganID:      row.KunjunganID,
			TanggalKunjungan: row.TanggalKunjungan,
			StatusKunjungan:  row.StatusKunjungan,

			NamaAnak:     row.NamaAnak,
			TanggalLahir: row.TanggalLahir,

			NamaIbu:         row.NamaIbu,
			NomorTeleponIbu: row.NomorTeleponIbu,

			NamaVaksin:      row.NamaVaksin,
			NamaDosis:       row.NamaDosis,
			JadwalImunisasi: row.JadwalImunisasi,
		}

	return result, nil
}

func (m *Main) GetAllKunjunganImunisasi() ([]models.KunjunganImunisasiResponse, error) {

	rows, err :=
		m.repository.
			GetAllKunjunganImunisasi()

	if err != nil {
		return nil, err
	}

	response :=
		[]models.KunjunganImunisasiResponse{}

	for _, row := range rows {

		response =
			append(
				response,
				models.KunjunganImunisasiResponse{
					KunjunganID:      row.KunjunganID,
					TanggalKunjungan: row.TanggalKunjungan,
					StatusKunjungan:  row.StatusKunjungan,
					NamaAnak:         row.NamaAnak,
				},
			)
	}

	return response, nil
}

func (m *Main) UpdateStatusKunjungan(
	kunjunganID uint,
	statusID uint,
) error {

	// cek data exist
	data, err :=
		m.repository.
			GetKunjunganImunisasiByID(
				kunjunganID,
			)

	if err != nil {
		return err
	}

	// kalau tidak ditemukan
	if data == nil ||
		data.KunjunganID == 0 {

		return fmt.Errorf(
			"kunjungan tidak ditemukan",
		)
	}

	// update status
	if err := m.repository.
		UpdateStatusKunjungan(
			kunjunganID,
			statusID,
		); err != nil {
		return err
	}

	if m.notifier != nil && data.IbuPenggunaID != 0 {
		statusName, sErr := m.repository.GetStatusKunjunganNameByID(statusID)
		statusName = strings.TrimSpace(statusName)
		if sErr != nil || statusName == "" {
			statusName = "diperbarui"
		}

		judul := "Status Kunjungan Imunisasi"
		anak := strings.TrimSpace(data.NamaAnak)
		body := fmt.Sprintf("Kunjungan %s untuk %s.", strings.ToLower(statusName), anak)

		if statusID == 3 { // Selesai / Dikunjungi
			judul = "Kunjungan Imunisasi Selesai"
			body = fmt.Sprintf("Kunjungan imunisasi untuk %s telah selesai dilaksanakan.", anak)
		} else if statusID == 4 { // Dibatalkan
			judul = "Kunjungan Imunisasi Dibatalkan"
			body = fmt.Sprintf("Kunjungan imunisasi untuk %s telah dibatalkan.", anak)
		}

		payload := map[string]string{
			"type":         "kunjungan_status",
			"kunjungan_id": strconv.FormatUint(uint64(kunjunganID), 10),
			"status_id":    strconv.FormatUint(uint64(statusID), 10),
			"status_name":  statusName,
			"anak_id":      strconv.FormatUint(uint64(data.AnakID), 10),
			"nama_anak":    anak,
		}

		m.sendKunjunganNotification(data.IbuPenggunaID, judul, body, payload)
	}

	return nil
}

func (m *Main) sendKunjunganNotification(ibuPenggunaID uint, judul string, body string, payload map[string]string) {
	if m.notifier == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. Send to Ibu
	if ibuPenggunaID != 0 {
		perangkat, err := m.repository.GetPerangkatByUserID(ibuPenggunaID)
		if err == nil && perangkat != nil && strings.TrimSpace(perangkat.FcmToken) != "" {
			if err := m.notifier.SendToToken(ctx, perangkat.FcmToken, judul, body, payload); err != nil {
				log.Printf("[FCM] gagal kirim notifikasi kunjungan ke ibu (user %d): %v", ibuPenggunaID, err)
			} else {
				log.Printf("[FCM] berhasil kirim notifikasi kunjungan ke ibu (user %d)", ibuPenggunaID)
			}
		}
	}

	// 2. Send to Kaders
	kaderTokens, err := m.repository.GetFcmTokensByRoleName("Kader")
	if err == nil && len(kaderTokens) > 0 {
		for _, token := range kaderTokens {
			trimmedToken := strings.TrimSpace(token)
			if trimmedToken != "" {
				if err := m.notifier.SendToToken(ctx, trimmedToken, judul, body, payload); err != nil {
					log.Printf("[FCM] gagal kirim notifikasi kunjungan ke kader token %s: %v", trimmedToken, err)
				} else {
					log.Printf("[FCM] berhasil kirim notifikasi kunjungan ke kader token %s", trimmedToken)
				}
			}
		}
	}
}

func (m *Main) UpdateTanggalKunjungan(
	kunjunganID uint,
	tanggalKunjungan string,
) error {

	// cek data exist
	data, err :=
		m.repository.
			GetKunjunganImunisasiByID(
				kunjunganID,
			)

	if err != nil {
		return err
	}

	// kalau tidak ditemukan
	if data == nil ||
		data.KunjunganID == 0 {

		return fmt.Errorf(
			"kunjungan tidak ditemukan",
		)
	}

	// update tanggal kunjungan
	if err := m.repository.
		UpdateTanggalKunjungan(
			kunjunganID,
			tanggalKunjungan,
		); err != nil {
		return err
	}

	if m.notifier != nil && data.IbuPenggunaID != 0 {
		judul := "Kunjungan Imunisasi Dijadwalkan Ulang"
		anak := strings.TrimSpace(data.NamaAnak)

		tglTeks := tanggalKunjungan
		if parsedTime, pErr := time.Parse("2006-01-02", tanggalKunjungan); pErr == nil {
			tglTeks = parsedTime.Format("02-01-2006")
		}

		body := fmt.Sprintf("Kunjungan imunisasi untuk %s telah dijadwalkan ulang menjadi tanggal %s.", anak, tglTeks)

		payload := map[string]string{
			"type":         "kunjungan_reschedule",
			"kunjungan_id": strconv.FormatUint(uint64(kunjunganID), 10),
			"tanggal":      tanggalKunjungan,
			"anak_id":      strconv.FormatUint(uint64(data.AnakID), 10),
			"nama_anak":    anak,
		}

		m.sendKunjunganNotification(data.IbuPenggunaID, judul, body, payload)
	}

	return nil
}

func (m *Main) GetKunjunganImunisasiByStatus(
	statusID uint,
) (
	[]models.KunjunganImunisasiResponse,
	error,
) {

	rows, err :=
		m.repository.
			GetKunjunganImunisasiByStatus(
				statusID,
			)

	if err != nil {
		return nil, err
	}

	response :=
		[]models.KunjunganImunisasiResponse{}

	for _, row := range rows {

		response =
			append(
				response,
				models.KunjunganImunisasiResponse{
					KunjunganID:      row.KunjunganID,
					TanggalKunjungan: row.TanggalKunjungan,
					StatusKunjungan:  row.StatusKunjungan,
					NamaAnak:         row.NamaAnak,
				},
			)
	}

	return response, nil
}

func (m *Main) SendScheduledImmunizationReminders() error {
	if m.notifier == nil {
		log.Println("[Reminders] Notifier is nil, skipping scheduled reminders.")
		return nil
	}

	today := time.Now()
	dates := []struct {
		Label    string
		DateStr  string
		Title    string
		BodyFunc func(tgl, dosis, anak string) string
		Payload  string
	}{
		{
			Label:   "Hari-H",
			DateStr: today.Format("2006-01-02"),
			Title:   "Jadwal Imunisasi Hari Ini",
			BodyFunc: func(tgl, dosis, anak string) string {
				return fmt.Sprintf("Halo Ibu, hari ini adalah jadwal imunisasi %s untuk %s. Silakan kunjungi Posyandu terdekat!", dosis, anak)
			},
			Payload: "h_0",
		},
		{
			Label:   "H-1",
			DateStr: today.AddDate(0, 0, 1).Format("2006-01-02"),
			Title:   "Pengingat Imunisasi (H-1)",
			BodyFunc: func(tgl, dosis, anak string) string {
				return fmt.Sprintf("Halo Ibu, besok (tanggal %s) adalah jadwal imunisasi %s untuk %s. Mohon persiapkan diri ya!", tgl, dosis, anak)
			},
			Payload: "h_1",
		},
		{
			Label:   "H-3",
			DateStr: today.AddDate(0, 0, 3).Format("2006-01-02"),
			Title:   "Pengingat Imunisasi (H-3)",
			BodyFunc: func(tgl, dosis, anak string) string {
				return fmt.Sprintf("Halo Ibu, 3 hari lagi (tanggal %s) adalah jadwal imunisasi %s untuk %s. Jangan lupa ya!", tgl, dosis, anak)
			},
			Payload: "h_3",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	for _, d := range dates {
		schedules, err := m.repository.GetUpcomingJadwalImunisasiByDate(d.DateStr)
		if err != nil {
			log.Printf("[Reminders] Gagal mengambil jadwal untuk %s (%s): %v", d.Label, d.DateStr, err)
			continue
		}

		log.Printf("[Reminders] Ditemukan %d jadwal untuk %s (%s)", len(schedules), d.Label, d.DateStr)

		for _, s := range schedules {
			if s.IbuPenggunaID == 0 {
				continue
			}

			perangkat, err := m.repository.GetPerangkatByUserID(s.IbuPenggunaID)
			if err != nil || perangkat == nil || strings.TrimSpace(perangkat.FcmToken) == "" {
				continue
			}

			tglTeks := d.DateStr
			if s.TanggalEstimasi != nil {
				tglTeks = s.TanggalEstimasi.Format("02-01-2006")
			}

			dosis := strings.TrimSpace(s.NamaDosis)
			anak := strings.TrimSpace(s.NamaAnak)
			body := d.BodyFunc(tglTeks, dosis, anak)

			payload := map[string]string{
				"type":      "imunisasi_reminder",
				"jadwal_id": strconv.FormatUint(uint64(s.JadwalID), 10),
				"label":     d.Payload,
			}

			if err := m.notifier.SendToToken(ctx, perangkat.FcmToken, d.Title, body, payload); err != nil {
				log.Printf("[Reminders] Gagal kirim notifikasi ke Ibu (user %d): %v", s.IbuPenggunaID, err)
			} else {
				log.Printf("[Reminders] Berhasil kirim notifikasi %s ke Ibu (user %d)", d.Label, s.IbuPenggunaID)
			}
		}
	}

	return nil
}
