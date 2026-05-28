package seeders

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type MasterStandarLKUSeeder struct {
	db *gorm.DB
}

type standarLKURow struct {
	UsiaReferensi string // format: "tahun:bulan" (contoh 1:4) atau bulan langsung (contoh 18)
	JenisKelamin  string
	SD3Neg        float64
	SD2Neg        float64
	SD1Neg        float64
	Median        float64
	SD1Pos        float64
	SD2Pos        float64
	SD3Pos        float64
}

func parseUsiaReferensiToBulan(raw string) (int, error) {
	v := strings.TrimSpace(raw)
	if v == "" {
		return 0, fmt.Errorf("usia referensi tidak boleh kosong")
	}

	if strings.Contains(v, ":") {
		parts := strings.Split(v, ":")
		if len(parts) != 2 {
			return 0, fmt.Errorf("format usia referensi tidak valid: %s", raw)
		}

		tahun, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil || tahun < 0 {
			return 0, fmt.Errorf("tahun tidak valid pada usia referensi: %s", raw)
		}

		bulan, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil || bulan < 0 || bulan > 11 {
			return 0, fmt.Errorf("bulan harus 0-11 pada usia referensi: %s", raw)
		}

		return (tahun * 12) + bulan, nil
	}

	totalBulan, err := strconv.Atoi(v)
	if err != nil || totalBulan < 0 {
		return 0, fmt.Errorf("usia referensi harus bulan >= 0 atau format tahun:bulan: %s", raw)
	}

	return totalBulan, nil
}

func NewMasterStandarLKUSeeder(db *gorm.DB) *MasterStandarLKUSeeder {
	return &MasterStandarLKUSeeder{db: db}
}

func (s *MasterStandarLKUSeeder) Seed() error {
	log.Println("Starting master standar LKU seeding...")

	now := time.Now()
	// Gunakan UsiaReferensi dengan format "tahun:bulan" atau bulan langsung.
	rows := []standarLKURow{
		// laki-laki
		{"0:0", "Laki-laki", 30.7, 31.9, 33.2, 34.5, 35.7, 37.0, 38.3},
		{"0:1", "Laki-laki", 33.8, 34.9, 36.1, 37.3, 38.4, 39.6, 40.8},
		{"0:2", "Laki-laki", 35.6, 36.8, 38.0, 39.1, 40.3, 41.5, 42.6},
		{"0:3", "Laki-laki", 37.0, 38.1, 39.3, 40.5, 41.7, 42.9, 44.1},
		{"0:4", "Laki-laki", 38.0, 39.2, 40.4, 41.6, 42.8, 44.0, 45.2},
		{"0:5", "Laki-laki", 38.9, 40.1, 41.4, 42.6, 43.8, 45.0, 46.2},
		{"0:6", "Laki-laki", 39.7, 40.9, 42.1, 43.3, 44.6, 45.8, 47.0},
		{"0:7", "Laki-laki", 40.3, 41.5, 42.7, 44.0, 45.2, 46.4, 47.7},
		{"0:8", "Laki-laki", 40.8, 42.0, 43.3, 44.5, 45.8, 47.0, 48.3},
		{"0:9", "Laki-laki", 41.2, 42.5, 43.7, 45.0, 46.3, 47.5, 48.8},
		{"0:10", "Laki-laki", 41.6, 42.9, 44.1, 45.4, 46.7, 47.9, 49.2},
		{"0:11", "Laki-laki", 41.9, 43.2, 44.5, 45.8, 47.0, 48.3, 49.6},
		{"1:0", "Laki-laki", 42.2, 43.5, 44.8, 46.1, 47.4, 48.6, 49.9},
		{"1:1", "Laki-laki", 42.5, 43.8, 45.0, 46.3, 47.6, 48.9, 50.2},
		{"1:2", "Laki-laki", 42.7, 44.0, 45.3, 46.6, 47.9, 49.2, 50.5},
		{"1:3", "Laki-laki", 42.9, 44.2, 45.5, 46.8, 48.1, 49.4, 50.7},
		{"1:4", "Laki-laki", 43.1, 44.4, 45.7, 47.0, 48.3, 49.6, 51.0},
		{"1:5", "Laki-laki", 43.2, 44.6, 45.9, 47.2, 48.5, 49.8, 51.2},
		{"1:6", "Laki-laki", 43.4, 44.7, 46.0, 47.4, 48.7, 50.0, 51.4},
		{"1:7", "Laki-laki", 43.5, 44.9, 46.2, 47.5, 48.9, 50.2, 51.5},
		{"1:8", "Laki-laki", 43.7, 45.0, 46.4, 47.7, 49.0, 50.4, 51.7},
		{"1:9", "Laki-laki", 43.8, 45.2, 46.5, 47.8, 49.2, 50.5, 51.9},
		{"1:10", "Laki-laki", 43.9, 45.3, 46.6, 48.0, 49.3, 50.7, 52.0},
		{"1:11", "Laki-laki", 44.1, 45.4, 46.8, 48.1, 49.5, 50.8, 52.2},
		{"2:0", "Laki-laki", 44.2, 45.5, 46.9, 48.3, 49.6, 51.0, 52.3},
		{"2:1", "Laki-laki", 44.3, 45.6, 47.0, 48.4, 49.7, 51.1, 52.5},
		{"2:2", "Laki-laki", 44.4, 45.8, 47.1, 48.5, 49.9, 51.2, 52.6},
		{"2:3", "Laki-laki", 44.5, 45.9, 47.2, 48.6, 50.0, 51.4, 52.7},
		{"2:4", "Laki-laki", 44.6, 46.0, 47.3, 48.7, 50.1, 51.5, 52.9},
		{"2:5", "Laki-laki", 44.7, 46.1, 47.4, 48.8, 50.2, 51.6, 53.0},
		{"2:6", "Laki-laki", 44.8, 46.1, 47.5, 48.9, 50.3, 51.7, 53.1},
		{"2:7", "Laki-laki", 44.8, 46.2, 47.6, 49.0, 50.4, 51.8, 53.2},
		{"2:8", "Laki-laki", 44.9, 46.3, 47.7, 49.1, 50.5, 51.9, 53.3},
		{"2:9", "Laki-laki", 45.0, 46.4, 47.8, 49.2, 50.6, 52.0, 53.4},
		{"2:10", "Laki-laki", 45.1, 46.5, 47.9, 49.3, 50.7, 52.1, 53.5},
		{"2:11", "Laki-laki", 45.1, 46.6, 48.0, 49.4, 50.8, 52.2, 53.6},
		{"3:0", "Laki-laki", 45.2, 46.6, 48.0, 49.5, 50.9, 52.3, 53.7},
		{"3:1", "Laki-laki", 45.3, 46.7, 48.1, 49.5, 51.0, 52.4, 53.8},
		{"3:2", "Laki-laki", 45.3, 46.8, 48.2, 49.6, 51.0, 52.5, 53.9},
		{"3:3", "Laki-laki", 45.4, 46.8, 48.2, 49.7, 51.1, 52.5, 54.0},
		{"3:4", "Laki-laki", 45.4, 46.9, 48.3, 49.7, 51.2, 52.6, 54.1},
		{"3:5", "Laki-laki", 45.5, 46.9, 48.4, 49.8, 51.3, 52.7, 54.1},
		{"3:6", "Laki-laki", 45.5, 47.0, 48.4, 49.9, 51.3, 52.8, 54.2},
		{"3:7", "Laki-laki", 45.6, 47.0, 48.5, 49.9, 51.4, 52.8, 54.3},
		{"3:8", "Laki-laki", 45.6, 47.1, 48.5, 50.0, 51.4, 52.9, 54.3},
		{"3:9", "Laki-laki", 45.7, 47.1, 48.6, 50.1, 51.5, 53.0, 54.4},
		{"3:10", "Laki-laki", 45.7, 47.2, 48.7, 50.1, 51.6, 53.0, 54.5},
		{"3:11", "Laki-laki", 45.8, 47.2, 48.7, 50.2, 51.6, 53.1, 54.5},
		{"4:0", "Laki-laki", 45.8, 47.3, 48.7, 50.2, 51.7, 53.1, 54.6},

		// perempuan
		{"0:0", "Perempuan", 30.3, 31.5, 32.7, 33.9, 35.1, 36.2, 37.4},
		{"0:1", "Perempuan", 33.0, 34.2, 35.4, 36.5, 37.7, 38.9, 40.1},
		{"0:2", "Perempuan", 34.6, 35.8, 37.0, 38.3, 39.5, 40.7, 41.9},
		{"0:3", "Perempuan", 35.8, 37.1, 38.3, 39.5, 40.8, 42.0, 43.3},
		{"0:4", "Perempuan", 36.8, 38.1, 39.3, 40.6, 41.8, 43.1, 44.4},
		{"0:5", "Perempuan", 37.6, 38.9, 40.2, 41.5, 42.7, 44.0, 45.3},
		{"0:6", "Perempuan", 38.3, 39.6, 40.9, 42.2, 43.5, 44.8, 46.1},
		{"0:7", "Perempuan", 38.9, 40.2, 41.5, 42.8, 44.1, 45.5, 46.8},
		{"0:8", "Perempuan", 39.4, 40.7, 42.0, 43.4, 44.7, 46.0, 47.4},
		{"0:9", "Perempuan", 39.8, 41.2, 42.5, 43.8, 45.2, 46.5, 47.8},
		{"0:10", "Perempuan", 40.2, 41.5, 42.9, 44.2, 45.6, 46.9, 48.3},
		{"0:11", "Perempuan", 40.5, 41.9, 43.2, 44.6, 45.9, 47.3, 48.6},
		{"1:0", "Perempuan", 40.8, 42.2, 43.5, 44.9, 46.3, 47.6, 49.0},
		{"1:1", "Perempuan", 41.1, 42.4, 43.8, 45.2, 46.5, 47.9, 49.3},
		{"1:2", "Perempuan", 41.3, 42.7, 44.1, 45.4, 46.8, 48.2, 49.5},
		{"1:3", "Perempuan", 41.5, 42.9, 44.3, 45.7, 47.0, 48.4, 49.8},
		{"1:4", "Perempuan", 41.7, 43.1, 44.5, 45.9, 47.2, 48.6, 50.0},
		{"1:5", "Perempuan", 41.9, 43.3, 44.7, 46.1, 47.4, 48.8, 50.2},
		{"1:6", "Perempuan", 42.1, 43.5, 44.9, 46.2, 47.6, 49.0, 50.4},
		{"1:7", "Perempuan", 42.3, 43.6, 45.0, 46.4, 47.8, 49.2, 50.6},
		{"1:8", "Perempuan", 42.4, 43.8, 45.2, 46.6, 48.0, 49.4, 50.7},
		{"1:9", "Perempuan", 42.6, 44.0, 45.3, 46.7, 48.1, 49.5, 50.9},
		{"1:10", "Perempuan", 42.7, 44.1, 45.5, 46.9, 48.3, 49.7, 51.1},
		{"1:11", "Perempuan", 42.9, 44.3, 45.6, 47.0, 48.4, 49.8, 51.2},
		{"2:0", "Perempuan", 43.0, 44.4, 45.8, 47.2, 48.6, 50.0, 51.4},
		{"2:1", "Perempuan", 43.1, 44.5, 45.9, 47.3, 48.7, 50.1, 51.5},
		{"2:2", "Perempuan", 43.3, 44.7, 46.1, 47.5, 48.9, 50.3, 51.7},
		{"2:3", "Perempuan", 43.4, 44.8, 46.2, 47.6, 49.0, 50.4, 51.8},
		{"2:4", "Perempuan", 43.5, 44.9, 46.3, 47.7, 49.1, 50.5, 51.9},
		{"2:5", "Perempuan", 43.6, 45.0, 46.4, 47.8, 49.2, 50.6, 52.0},
		{"2:6", "Perempuan", 43.7, 45.1, 46.5, 47.9, 49.3, 50.7, 52.2},
		{"2:7", "Perempuan", 43.8, 45.2, 46.6, 48.0, 49.4, 50.9, 52.3},
		{"2:8", "Perempuan", 43.9, 45.3, 46.7, 48.1, 49.6, 51.0, 52.4},
		{"2:9", "Perempuan", 44.0, 45.4, 46.8, 48.2, 49.7, 51.1, 52.5},
		{"2:10", "Perempuan", 44.1, 45.5, 46.9, 48.3, 49.7, 51.2, 52.6},
		{"2:11", "Perempuan", 44.2, 45.6, 47.0, 48.4, 49.8, 51.2, 52.7},
		{"3:0", "Perempuan", 44.3, 45.7, 47.1, 48.5, 49.9, 51.3, 52.7},
		{"3:1", "Perempuan", 44.4, 45.8, 47.2, 48.6, 50.0, 51.4, 52.8},
		{"3:2", "Perempuan", 44.4, 45.8, 47.3, 48.7, 50.1, 51.5, 52.9},
		{"3:3", "Perempuan", 44.5, 45.9, 47.3, 48.7, 50.2, 51.6, 53.0},
		{"3:4", "Perempuan", 44.6, 46.0, 47.4, 48.8, 50.2, 51.7, 53.1},
		{"3:5", "Perempuan", 44.6, 46.1, 47.5, 48.9, 50.3, 51.7, 53.1},
		{"3:6", "Perempuan", 44.7, 46.1, 47.5, 49.0, 50.4, 51.8, 53.2},
		{"3:7", "Perempuan", 44.8, 46.2, 47.6, 49.0, 50.4, 51.9, 53.3},
		{"3:8", "Perempuan", 44.8, 46.3, 47.7, 49.1, 50.5, 51.9, 53.3},
		{"3:9", "Perempuan", 44.9, 46.3, 47.7, 49.2, 50.6, 52.0, 53.4},
		{"3:10", "Perempuan", 45.0, 46.4, 47.8, 49.2, 50.6, 52.1, 53.5},
		{"3:11", "Perempuan", 45.0, 46.4, 47.9, 49.3, 50.7, 52.1, 53.5},
		{"4:0", "Perempuan", 45.1, 46.5, 47.9, 49.3, 50.8, 52.2, 53.6},
		{"4:1", "Perempuan", 45.1, 46.5, 48.0, 49.4, 50.8, 52.2, 53.6},
		{"4:2", "Perempuan", 45.2, 46.6, 48.0, 49.4, 50.9, 52.3, 53.7},
		{"4:3", "Perempuan", 45.2, 46.7, 48.1, 49.5, 50.9, 52.3, 53.8},
		{"4:4", "Perempuan", 45.3, 46.7, 48.1, 49.5, 51.0, 52.4, 53.8},
		{"4:5", "Perempuan", 45.3, 46.8, 48.2, 49.6, 51.0, 52.4, 53.9},
		{"4:6", "Perempuan", 45.4, 46.8, 48.2, 49.6, 51.1, 52.5, 53.9},
		{"4:7", "Perempuan", 45.4, 46.9, 48.3, 49.7, 51.1, 52.5, 54.0},
		{"4:8", "Perempuan", 45.5, 46.9, 48.3, 49.7, 51.2, 52.6, 54.0},
		{"4:9", "Perempuan", 45.5, 46.9, 48.4, 49.8, 51.2, 52.6, 54.1},
		{"4:10", "Perempuan", 45.6, 47.0, 48.4, 49.8, 51.3, 52.7, 54.1},
		{"4:11", "Perempuan", 45.6, 47.0, 48.5, 49.9, 51.3, 52.7, 54.1},
		{"5:0", "Perempuan", 45.7, 47.1, 48.5, 49.9, 51.3, 52.8, 54.2},
	}

	for _, row := range rows {
		usiaBulan, err := parseUsiaReferensiToBulan(row.UsiaReferensi)
		if err != nil {
			return err
		}

		item := models.MasterStandarAntropometri{
			Parameter:    "lk_u",
			JenisKelamin: models.GenderType(row.JenisKelamin),
			NilaiSumbuX:  float64(usiaBulan),
			SD3Neg:       row.SD3Neg,
			SD2Neg:       row.SD2Neg,
			SD1Neg:       row.SD1Neg,
			Median:       row.Median,
			SD1Pos:       row.SD1Pos,
			SD2Pos:       row.SD2Pos,
			SD3Pos:       row.SD3Pos,
			CreatedAt:    now,
			UpdatedAt:    now,
		}

		if err = s.db.Where(
			"parameter = ? AND jenis_kelamin = ? AND nilai_sumbu_x = ?",
			item.Parameter,
			item.JenisKelamin,
			item.NilaiSumbuX,
		).FirstOrCreate(&item).Error; err != nil {
			log.Printf("Error seeding MasterStandarAntropometri (%s-%s-%s): %v\n", item.Parameter, item.JenisKelamin, row.UsiaReferensi, err)
			return err
		}
	}

	log.Println("Master standar LKU seeding completed!")
	return nil
}
