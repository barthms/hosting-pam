package seeders

import (
	"log"
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type PemantauanAnakSeeder struct {
	db *gorm.DB
}

func NewPemantauanAnakSeeder(db *gorm.DB) *PemantauanAnakSeeder {
	return &PemantauanAnakSeeder{db: db}
}

func (s *PemantauanAnakSeeder) Seed() error {
	log.Println("Reseting Pemantauan Anak data (Truncating tables)...")

	// Pembersihan total untuk mereset data yang berantakan
	s.db.Exec("TRUNCATE TABLE detail_pemantauan RESTART IDENTITY CASCADE")
	s.db.Exec("TRUNCATE TABLE lembar_pemantauan RESTART IDENTITY CASCADE")
	s.db.Exec("TRUNCATE TABLE kategori_tanda_sakit RESTART IDENTITY CASCADE")
	s.db.Exec("TRUNCATE TABLE rentang_usia RESTART IDENTITY CASCADE")

	rentangData := []struct {
		Nama        string
		Satuan      string
		MaxPeriode  int
		Indikators  []string
	}{
		{
			Nama:       "29 Hari - 3 Bulan",
			Satuan:     "Minggu",
			MaxPeriode: 12,
			Indikators: []string{
				"Sesak napas / cuping hidung kembang kempis / dada tertarik ke dalam",
				"Batuk dengan bunyi grok-grok/mengi",
				"Suhu tubuh panas > 38.5C / ada tanda pendarahan (mimisan/gusi berdarah/muntah kopi/BAB hitam)",
				"BAB lebih sering/lebih encer, dengan mata cekung/haus minum dengan lahap/diare disertai darah",
				"Jumlah air kencing sedikit/tidak kencing selama 6 jam, warna kuning pekat, kecokelatan, atau warna lainnya",
				"Warna kulit tampak biru/memar di sekitar mulut/tangan/kaki",
				"Aktivitas tampak lemah/tidak bergerak/menangis merintih",
				"Hisapan bayi lemah/tidak bergerak. Muntah susu/cairan hijau. Kencing < 6x/hari. Warna kencing kurang pekat",
				"Tidak mau makan/minum. Berat badan tidak naik sesuai pertumbuhan",
			},
		},
		{
			Nama:       "3 - 6 Bulan",
			Satuan:     "Bulan",
			MaxPeriode: 6,
			Indikators: []string{
				"Tidak mau menyusu",
				"Kejang",
				"Lemah atau kesadaran menurun",
				"Sesak napas",
				"Merintih",
				"Demam",
				"Sangat kuning",
				"Mata cekung dan malas minum",
			},
		},
		{
			Nama:       "6 - 12 Bulan",
			Satuan:     "Bulan",
			MaxPeriode: 12,
			Indikators: []string{
				"Kenaikan berat badan tidak sesuai",
				"Diare",
				"Demam/Panas",
				"Batuk/Pilek",
				"Masalah Kulit",
			},
		},
		{
			Nama:       "12 - 24 Bulan",
			Satuan:     "Bulan",
			MaxPeriode: 24,
			Indikators: []string{
				"Kenaikan berat badan tidak sesuai",
				"Diare",
				"Demam/Panas",
				"Batuk/Pilek",
				"Masalah Kulit",
			},
		},
		{
			Nama:       "2 - 6 Tahun",
			Satuan:     "Tahun",
			MaxPeriode: 6,
			Indikators: []string{
				"Kenaikan berat badan tidak sesuai",
				"Diare",
				"Demam/Panas",
				"Batuk/Pilek",
				"Masalah Kulit",
			},
		},
	}

	for _, rd := range rentangData {
		var rentang models.RentangUsia
		err := s.db.Where(models.RentangUsia{NamaRentang: rd.Nama}).FirstOrCreate(&rentang, models.RentangUsia{
			NamaRentang: rd.Nama,
			SatuanWaktu: rd.Satuan,
			MaxPeriode:  rd.MaxPeriode,
		}).Error
		if err != nil {
			return err
		}

		for _, gej := range rd.Indikators {
			var kategori models.KategoriTandaSakit
			err := s.db.Where(models.KategoriTandaSakit{RentangUsiaID: rentang.ID, Gejala: gej}).FirstOrCreate(&kategori, models.KategoriTandaSakit{
				RentangUsiaID: rentang.ID,
				Gejala:        gej,
				IsActive:      true,
			}).Error
			if err != nil {
				return err
			}
		}
	}

	log.Println("Pemantauan Anak seeding completed!")
	return nil
}
