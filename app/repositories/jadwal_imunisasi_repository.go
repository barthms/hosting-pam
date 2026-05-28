package repositories

import (
	"time"
	// "monitoring-service/app/models"
)

type JadwalImunisasiJoin struct {
	AnakID          int32
	NamaAnak        string
	TanggalLahir    *time.Time
	JadwalID        uint
	NamaDosis       string
	TanggalEstimasi *time.Time
	Deskripsi       string
	EfekSamping     string
	StatusID        uint
	Status          string
}

func (m *Main) GetJadwalImunisasiByUserID(
	userID int32,
) ([]JadwalImunisasiJoin, error) {

	var result []JadwalImunisasiJoin

	err := m.postgres.
		Table("pengguna p").
		Select(`
		a.id as anak_id,
		pd_anak.nama_lengkap as nama_anak,
		a.tanggal_lahir,

		j.id as jadwal_id,
		j.id_dosis_vaksin as dosis_vaksin_id,
		dv.nama_dosis,
		j.tanggal_estimasi,

		sj.id as status_id,
		sj.nama_status as status,

		v.deskripsi,
		v.efek_samping
	`).
		Joins(`
		JOIN penduduk pd_ibu
		ON pd_ibu.id = p.penduduk_id
	`).
		Joins(`
		JOIN ibu i
		ON i.penduduk_id = pd_ibu.id
	`).
		Joins(`
		JOIN kehamilan k
		ON k.ibu_id = i.id
	`).
		Joins(`
		JOIN anak a
		ON a.kehamilan_id = k.id
	`).
		Joins(`
		JOIN penduduk pd_anak
		ON pd_anak.id = a.penduduk_id
	`).
		Joins(`
		LEFT JOIN jadwal_imunisasi_anak j
		ON j.id_anak = a.id
	`).
		Joins(`
		LEFT JOIN dosis_vaksin dv
		ON dv.id = j.id_dosis_vaksin
	`).
		Joins(`
		LEFT JOIN status_jadwal sj
		ON sj.id = j.id_status_jadwal
	`).
		Joins(`
		INNER JOIN vaksin v
		ON dv.id_vaksin = v.id
	`).
		Where("p.id = ?", userID).
		Order("a.id ASC, j.tanggal_estimasi ASC").
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *Main) GetJadwalImunisasiByAnakID(
	userID int32,
	anakID int32,
) ([]JadwalImunisasiJoin, error) {

	var result []JadwalImunisasiJoin

	err := m.postgres.
		Table("pengguna p").
		Select(`
			a.id as anak_id,
			pd_anak.nama_lengkap as nama_anak,
			a.tanggal_lahir,

			j.id as jadwal_id,
			dv.nama_dosis,
			j.tanggal_estimasi,

			sj.id as status_id,
			sj.nama_status as status,

			v.deskripsi,
			v.efek_samping
		`).
		Joins(`
			JOIN penduduk pd_ibu
			ON pd_ibu.id = p.penduduk_id
		`).
		Joins(`
			JOIN ibu i
			ON i.penduduk_id = pd_ibu.id
		`).
		Joins(`
			JOIN kehamilan k
			ON k.ibu_id = i.id
		`).
		Joins(`
			JOIN anak a
			ON a.kehamilan_id = k.id
		`).
		Joins(`
			JOIN penduduk pd_anak
			ON pd_anak.id = a.penduduk_id
		`).
		Joins(`
			LEFT JOIN jadwal_imunisasi_anak j
			ON j.id_anak = a.id
		`).
		Joins(`
			LEFT JOIN dosis_vaksin dv
			ON dv.id = j.id_dosis_vaksin
		`).
		Joins(`
			LEFT JOIN status_jadwal sj
			ON sj.id = j.id_status_jadwal
		`).
		Joins(`
			LEFT JOIN vaksin v
			ON v.id = dv.id_vaksin
		`).
		Where("p.id = ?", int(userID)).
		Where("a.id = ?", int(anakID)). // 🔥 filter anak spesifik
		Order("j.tanggal_estimasi ASC").
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *Main) UpdateTanggalEstimasi(
	jadwalID uint,
	newDate time.Time,
) error {

	return m.postgres.
		Table("jadwal_imunisasi_anak").
		Where("id = ?", jadwalID).
		Update("tanggal_estimasi", newDate).Error
}

func (m *Main) GetJadwalImunisasiByJadwalID(
	userID int32,
	jadwalID uint,
) (*JadwalImunisasiJoin, error) {

	var result JadwalImunisasiJoin

	err := m.postgres.
		Table("pengguna p").
		Select(`
			a.id as anak_id,
			pd_anak.nama_lengkap as nama_anak,
			a.tanggal_lahir,

			j.id as jadwal_id,
			dv.nama_dosis,
			j.tanggal_estimasi,

			sj.id as status_id,
			sj.nama_status as status,

			v.deskripsi,
			v.efek_samping
		`).
		Joins(`
			JOIN penduduk pd_ibu
			ON pd_ibu.id = p.penduduk_id
		`).
		Joins(`
			JOIN ibu i
			ON i.penduduk_id = pd_ibu.id
		`).
		Joins(`
			JOIN kehamilan k
			ON k.ibu_id = i.id
		`).
		Joins(`
			JOIN anak a
			ON a.kehamilan_id = k.id
		`).
		Joins(`
			JOIN penduduk pd_anak
			ON pd_anak.id = a.penduduk_id
		`).
		Joins(`
			LEFT JOIN jadwal_imunisasi_anak j
			ON j.id_anak = a.id
		`).
		Joins(`
			LEFT JOIN dosis_vaksin dv
			ON dv.id = j.id_dosis_vaksin
		`).
		Joins(`
			LEFT JOIN status_jadwal sj
			ON sj.id = j.id_status_jadwal
		`).
		Joins(`
			LEFT JOIN vaksin v
			ON v.id = dv.id_vaksin
		`).
		Where("p.id = ?", userID).
		Where("j.id = ?", jadwalID). // 🔥 ini kuncinya
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return &result, nil
}

type UpcomingJadwalImunisasi struct {
	JadwalID        uint       `json:"jadwal_id"`
	TanggalEstimasi *time.Time `json:"tanggal_estimasi"`
	NamaAnak        string     `json:"nama_anak"`
	NamaDosis       string     `json:"nama_dosis"`
	IbuPenggunaID    uint       `json:"ibu_pengguna_id"`
}

func (m *Main) GetUpcomingJadwalImunisasiByDate(dateStr string) ([]UpcomingJadwalImunisasi, error) {
	var result []UpcomingJadwalImunisasi

	err := m.postgres.
		Table("jadwal_imunisasi_anak jia").
		Select(`
			jia.id AS jadwal_id,
			jia.tanggal_estimasi,
			p_anak.nama_lengkap AS nama_anak,
			dv.nama_dosis AS nama_dosis,
			u.id AS ibu_pengguna_id
		`).
		Joins("INNER JOIN anak a ON a.id = jia.id_anak").
		Joins("INNER JOIN penduduk p_anak ON p_anak.id = a.penduduk_id").
		Joins("INNER JOIN kehamilan kh ON kh.id = a.kehamilan_id").
		Joins("INNER JOIN ibu i ON i.id = kh.ibu_id").
		Joins("INNER JOIN penduduk p_ibu ON p_ibu.id = i.penduduk_id").
		Joins("INNER JOIN pengguna u ON u.penduduk_id = p_ibu.id").
		Joins("LEFT JOIN dosis_vaksin dv ON dv.id = jia.id_dosis_vaksin").
		Joins("LEFT JOIN catatan_imunisasi_anak cia ON cia.jadwal_imunisasi_anak_id = jia.id").
		Where("cia.id IS NULL").
		Where("DATE(jia.tanggal_estimasi) = DATE(?)", dateStr).
		Scan(&result).Error

	return result, err
}
