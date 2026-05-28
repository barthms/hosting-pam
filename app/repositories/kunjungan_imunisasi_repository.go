package repositories

import "time"

type KunjunganImunisasiDetailJoin struct {
	KunjunganID      uint
	TanggalKunjungan *time.Time
	StatusKunjungan  string
	IbuPenggunaID    uint
	AnakID           uint

	NamaAnak     string
	TanggalLahir *time.Time

	NamaIbu         string
	NomorTeleponIbu string

	NamaVaksin      string
	NamaDosis       string
	JadwalImunisasi *time.Time
}

type KunjunganImunisasiJoin struct {
	KunjunganID      uint
	TanggalKunjungan *time.Time
	StatusKunjungan  string
	NamaAnak         string
}

func (m *Main) GetKunjunganImunisasiByID(
	kunjunganID uint,
) (*KunjunganImunisasiDetailJoin, error) {

	var result KunjunganImunisasiDetailJoin

	err := m.postgres.
		Table("kunjungan_imunisasi ki").
		Select(`
		ki.id AS kunjungan_id,
		ki.tanggal_kunjungan,
		sk.status_kunjungan,

		a.id AS anak_id,
		p_anak.nama_lengkap AS nama_anak,
		a.tanggal_lahir,

		p_ibu.nama_lengkap AS nama_ibu,
		p_ibu.telepon AS nomor_telepon_ibu,
		u.id AS ibu_pengguna_id,

		v.nama AS nama_vaksin,
		dv.nama_dosis,
		jia.tanggal_estimasi AS jadwal_imunisasi
	`).
		Joins(`
		INNER JOIN status_kunjungan sk
		ON sk.id = ki.id_status_kunjungan
	`).
		Joins(`
		INNER JOIN jadwal_imunisasi_anak jia
		ON jia.id = ki.id_jadwal_imunisasi
	`).
		Joins(`
		INNER JOIN anak a
		ON a.id = jia.id_anak
	`).
		Joins(`
		INNER JOIN penduduk p_anak
		ON p_anak.id = a.penduduk_id
	`).
		Joins(`
		INNER JOIN kehamilan kh
		ON kh.id = a.kehamilan_id
	`).
		Joins(`
		INNER JOIN ibu i
		ON i.id = kh.ibu_id
	`).
		Joins(`
		INNER JOIN penduduk p_ibu
		ON p_ibu.id = i.penduduk_id
	`).
		Joins(`
		LEFT JOIN pengguna u
		ON u.penduduk_id = p_ibu.id
	`).
		Joins(`
		LEFT JOIN dosis_vaksin dv
		ON dv.id = jia.id_dosis_vaksin
	`).
		Joins(`
		LEFT JOIN vaksin v
		ON v.id = dv.id_vaksin
	`).
		Where("ki.id = ?", kunjunganID).
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (m *Main) GetAllKunjunganImunisasi() ([]KunjunganImunisasiJoin, error) {

	var result []KunjunganImunisasiJoin

	err := m.postgres.
		Table("kunjungan_imunisasi ki").
		Select(`
			ki.id AS kunjungan_id,
			ki.tanggal_kunjungan,
			sk.status_kunjungan,

			p_anak.nama_lengkap AS nama_anak
		`).
		Joins(`
			INNER JOIN status_kunjungan sk
			ON sk.id = ki.id_status_kunjungan
		`).
		Joins(`
			INNER JOIN jadwal_imunisasi_anak jia
			ON jia.id = ki.id_jadwal_imunisasi
		`).
		Joins(`
			INNER JOIN anak a
			ON a.id = jia.id_anak
		`).
		Joins(`
			INNER JOIN penduduk p_anak
			ON p_anak.id = a.penduduk_id
		`).
		Order("ki.tanggal_kunjungan DESC").
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *Main) UpdateStatusKunjungan(
	kunjunganID uint,
	statusID uint,
) error {

	return m.postgres.
		Table("kunjungan_imunisasi").
		Where("id = ?", kunjunganID).
		Update(
			"id_status_kunjungan",
			statusID,
		).Error
}

func (m *Main) UpdateTanggalKunjungan(
	kunjunganID uint,
	tanggalKunjungan string,
) error {

	return m.postgres.
		Table("kunjungan_imunisasi").
		Where("id = ?", kunjunganID).
		Update(
			"tanggal_kunjungan",
			tanggalKunjungan,
		).Error
}

func (m *Main) GetKunjunganImunisasiByStatus(
	statusID uint,
) (
	[]KunjunganImunisasiJoin,
	error,
) {

	var result []KunjunganImunisasiJoin

	err := m.postgres.
		Table("kunjungan_imunisasi ki").
		Select(`
			ki.id AS kunjungan_id,
			ki.tanggal_kunjungan,
			sk.status_kunjungan,

			p_anak.nama_lengkap AS nama_anak
		`).
		Joins(`
			INNER JOIN status_kunjungan sk
			ON sk.id = ki.id_status_kunjungan
		`).
		Joins(`
			INNER JOIN jadwal_imunisasi_anak jia
			ON jia.id = ki.id_jadwal_imunisasi
		`).
		Joins(`
			INNER JOIN anak a
			ON a.id = jia.id_anak
		`).
		Joins(`
			INNER JOIN penduduk p_anak
			ON p_anak.id = a.penduduk_id
		`).
		Where(
			"ki.id_status_kunjungan = ?",
			statusID,
		).
		Order("ki.tanggal_kunjungan DESC").
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}
