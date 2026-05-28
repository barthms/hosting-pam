package seed

import (
	"log"

	"gorm.io/gorm"

	"monitoring-service/app/models"
)

func SeederJenisPelayanan(db *gorm.DB) (map[string]int32, error) {
	log.Println("🌱 Seed: jenis pelayanan KIA 2024...")

	result := make(map[string]int32)

	data := []models.JenisPelayanan{
		// ===== BayiBaruLahir =====
		{Nama: "Perawatan Tali Pusat", TipeInput: "checkbox", GroupName: "BBL", Section: "pemeriksaan"},
		{Nama: "IMD", TipeInput: "checkbox", GroupName: "BBL", Section: "pemeriksaan"},
		{Nama: "Vitamin K1", TipeInput: "checkbox", GroupName: "BBL", Section: "pemeriksaan"},
		{Nama: "Salep/Tetes Mata Antibiotik", TipeInput: "checkbox", GroupName: "BBL", Section: "pemeriksaan"},
		{Nama: "Skrining BBL/SHK", TipeInput: "checkbox", GroupName: "BBL", Section: "pemeriksaan"},
		{Nama: "KIA", TipeInput: "checkbox", GroupName: "BBL", Section: "pemeriksaan"},
		{Nama: "Imunisasi Hepatitis B", TipeInput: "checkbox", GroupName: "BBL", Section: "pemeriksaan"},
		{Nama: "Tripel Eliminasi", TipeInput: "checkbox", GroupName: "BBL", Section: "pemeriksaan"},

		// ===== ANTROPOMETRI =====
		{Nama: "BB", TipeInput: "number", GroupName: "antropometri", Section: "pemeriksaan"},
		{Nama: "PB", TipeInput: "number", GroupName: "antropometri", Section: "pemeriksaan"},
		{Nama: "TB", TipeInput: "number", GroupName: "antropometri", Section: "pemeriksaan"},
		{Nama: "LK", TipeInput: "number", GroupName: "antropometri", Section: "pemeriksaan"},

		// ===== PERKEMBANGAN =====
		{Nama: "Perkembangan", TipeInput: "text", GroupName: "skrining", Section: "pemeriksaan"},
		{Nama: "KIE", TipeInput: "text", GroupName: "edukasi", Section: "pemeriksaan"},

		// ===== TERAPI =====
		{Nama: "Vitamin A", TipeInput: "checkbox", GroupName: "terapi", Section: "pemeriksaan"},
		{Nama: "Obat Cacing", TipeInput: "checkbox", GroupName: "terapi", Section: "pemeriksaan"},

		// ===== OBSERVASI =====
		{Nama: "Menyusu", TipeInput: "checkbox", GroupName: "observasi", Section: "kondisi"},
		{Nama: "Kondisi Umum", TipeInput: "checkbox", GroupName: "observasi", Section: "kondisi"},
		{Nama: "Tanda Bahaya", TipeInput: "checkbox", GroupName: "observasi", Section: "kondisi"},
		{Nama: "Vitamin K1", TipeInput: "checkbox", GroupName: "observasi", Section: "kondisi"},
		{Nama: "Dini(IMD)", TipeInput: "checkbox", GroupName: "observasi", Section: "kondisi"},
		{Nama: "Identifikasi Kuning", TipeInput: "text", GroupName: "observasi", Section: "kondisi"},
		{Nama: "Masalah", TipeInput: "text", GroupName: "observasi", Section: "kondisi"},
		{Nama: "Dirujuk", TipeInput: "checkbox", GroupName: "observasi", Section: "kondisi"},
		{Nama: "Imunisasi HB*", TipeInput: "text", GroupName: "observasi", Section: "kondisi"},
		{Nama: "Skrining Hipotiroid Kongenital", TipeInput: "text", GroupName: "observasi", Section: "kondisi"},
		{Nama: "Bagian Kuning", TipeInput: "text", GroupName: "observasi", Section: "kondisi"},
		{Nama: "Hasil Skrining Hipotiroid Kongenital", TipeInput: "text", GroupName: "observasi", Section: "kondisi"},
		{Nama: "Skrining Penyakit jantung bawaan kritis 24-48 jam.", TipeInput: "text", GroupName: "observasi", Section: "kondisi"},

		// ===== VITAMIN A & OBAT CACING =====
		{Nama: "VIT A KAPSUL BIRU(100.000 IU)", TipeInput: "checkbox", GroupName: "Vitamin A dan Obat Cacing", Section: "Februari/Agustus"},
		{Nama: "VIT A KAPSUL MERAH(200.000 IU)", TipeInput: "checkbox", GroupName: "Vitamin A dan Obat Cacing", Section: "Februari/Agustus"},

		// ===== IMUNISASI =====
		{Nama: "Hepatitis B (<24 Jam)", TipeInput: "checkbox", GroupName: "Imunisasi Dasar", Section: "0-24 Jam"},
		{Nama: "BCG", TipeInput: "checkbox", GroupName: "Imunisasi Dasar", Section: "0-1 Bulan"},
		{Nama: "Polio tetes 1", TipeInput: "checkbox", GroupName: "Imunisasi Dasar", Section: "0-1 Bulan"},
		{Nama: "DPT-HB-Hib 1", TipeInput: "checkbox", GroupName: "Imunisasi Dasar", Section: "2 Bulan"},
		{Nama: "Polio tetes 2", TipeInput: "checkbox", GroupName: "Imunisasi Dasar", Section: "2 Bulan"},
		{Nama: "Rotavirus (RV)1", TipeInput: "checkbox", GroupName: "Imunisasi Dasar", Section: "2 Bulan"},
		{Nama: "PCV 1", TipeInput: "checkbox", GroupName: "Imunisasi Dasar", Section: "2 Bulan"},
		{Nama: "DPT-HB-Hib 2", TipeInput: "checkbox", GroupName: "Imunisasi Dasar", Section: "3 Bulan"},
		{Nama: "Polio tetes 3", TipeInput: "checkbox", GroupName: "Imunisasi Dasar", Section: "3 Bulan"},
		{Nama: "Rotavirus (RV)2", TipeInput: "checkbox", GroupName: "Imunisasi Dasar", Section: "3 Bulan"},
		{Nama: "PCV 2", TipeInput: "checkbox", GroupName: "Imunisasi Dasar", Section: "3 Bulan"},
		{Nama: "DPT-HB-Hib 3", TipeInput: "checkbox", GroupName: "Imunisasi Dasar", Section: "4 Bulan"},
		{Nama: "Polio tetes 4", TipeInput: "checkbox", GroupName: "Imunisasi Dasar", Section: "4 Bulan"},
		{Nama: "Polio Suntik (IPV) 1", TipeInput: "checkbox", GroupName: "Imunisasi Dasar", Section: "4 Bulan"},
		{Nama: "Rotavirus (RV) 3", TipeInput: "checkbox", GroupName: "Imunisasi Dasar", Section: "4 Bulan"},
		{Nama: "Campak -Rubella (MR)", TipeInput: "checkbox", GroupName: "Imunisasi Dasar", Section: "9 Bulan"},
		{Nama: "Japanese Encephalitis (JE)", TipeInput: "checkbox", GroupName: "Imunisasi Dasar", Section: "9 Bulan"},
		{Nama: "Polio Suntik (IPV) 2", TipeInput: "checkbox", GroupName: "Imunisasi Lanjutan", Section: "18 Bulan"},
		{Nama: "PCV3", TipeInput: "checkbox", GroupName: "Imunisasi Lanjutan", Section: "12 Bulan"},
		{Nama: "DPT-HB-Hib Lanjutan", TipeInput: "checkbox", GroupName: "Imunisasi Lanjutan", Section: "18 Bulan"},
		{Nama: "Campak Rubella (MR) Lanjutan", TipeInput: "checkbox", GroupName: "Imunisasi Lanjutan", Section: "18 Bulan"},
	}

	for _, item := range data {
		obj := models.JenisPelayanan{}

		err := db.Where("nama = ?", item.Nama).
			FirstOrCreate(&obj, item).Error

		if err != nil {
			return nil, err
		}

		result[item.Nama] = obj.ID
	}

	log.Println("✅ Seeder Jenis Pelayanan Selesai")
	return result, nil
}
