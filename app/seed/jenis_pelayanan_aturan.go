package seed

import (
	"log"
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

func SeederAturanPelayanan(db *gorm.DB)(map[string]int32, error) {
	log.Println("🌱 Seed: aturan pelayanan...")

	result := make(map[string]int32)

	data := []models.AturanPelayanan{

		// =========================
		// VITAMIN A BIRU (6–11 bulan)
		// =========================
		{JenisPelayananID: 30, UmurMinBulan: 6, UmurMaxBulan: 11, Bulan: 2},
		{JenisPelayananID: 30, UmurMinBulan: 6, UmurMaxBulan: 11, Bulan: 8},

		// =========================
		// VITAMIN A MERAH (12–59 bulan)
		// =========================
		{JenisPelayananID: 31, UmurMinBulan: 12, UmurMaxBulan: 59, Bulan: 2},
		{JenisPelayananID: 31, UmurMinBulan: 12, UmurMaxBulan: 59, Bulan: 8},

		// =========================
		// OBAT CACING (12–59 bulan)
		// =========================
		{JenisPelayananID: 16, UmurMinBulan: 12, UmurMaxBulan: 59,},


		// imunisasi
		{JenisPelayananID: 32, UmurMinBulan: 0, UmurMaxBulan: 1, Bulan:0},
		{JenisPelayananID: 33, UmurMinBulan: 0, UmurMaxBulan: 11,Bulan:1},
		{JenisPelayananID: 34, UmurMinBulan: 0, UmurMaxBulan: 11, Bulan: 1},
		{JenisPelayananID: 35, UmurMinBulan: 2, UmurMaxBulan: 11, Bulan: 2},
		{JenisPelayananID: 36, UmurMinBulan: 2, UmurMaxBulan: 11, Bulan: 2},
		{JenisPelayananID: 37, UmurMinBulan: 2, UmurMaxBulan: 6,Bulan: 2},
		{JenisPelayananID: 38, UmurMinBulan: 2, UmurMaxBulan: 11,Bulan: 2},
		{JenisPelayananID: 39, UmurMinBulan: 3, UmurMaxBulan: 11, Bulan: 3},
		{JenisPelayananID: 40, UmurMinBulan: 3, UmurMaxBulan: 11, Bulan: 3},
		{JenisPelayananID: 41, UmurMinBulan: 3, UmurMaxBulan: 6,Bulan: 3},
		{JenisPelayananID: 42, UmurMinBulan: 3, UmurMaxBulan: 11,Bulan: 3},
		{JenisPelayananID: 43, UmurMinBulan: 4, UmurMaxBulan: 11, Bulan: 4},
		{JenisPelayananID: 44, UmurMinBulan: 4, UmurMaxBulan: 11, Bulan: 4},
		{JenisPelayananID: 45, UmurMinBulan: 4, UmurMaxBulan: 11,Bulan: 4},
		{JenisPelayananID: 46, UmurMinBulan: 4, UmurMaxBulan: 6,Bulan: 4},
		{JenisPelayananID: 47, UmurMinBulan: 9, UmurMaxBulan: 11, Bulan: 9},
		{JenisPelayananID: 49, UmurMinBulan: 9, UmurMaxBulan: 11, Bulan: 9},
		{JenisPelayananID: 48, UmurMinBulan: 10, UmurMaxBulan: 10, Bulan: 10},
		{JenisPelayananID: 50, UmurMinBulan: 12, UmurMaxBulan: 10,Bulan: 12},
		{JenisPelayananID: 51, UmurMinBulan: 18, UmurMaxBulan: 23, Bulan: 18},
		{JenisPelayananID: 52, UmurMinBulan: 18, UmurMaxBulan: 23, Bulan: 18},

	}

	for _, item := range data {

		// 🔥 biar tidak duplikat
		err := db.Where(
			"jenis_pelayanan_id = ? AND umur_min_bulan = ? AND umur_max_bulan = ? AND bulan = ?",
			item.JenisPelayananID,
			item.UmurMinBulan,
			item.UmurMaxBulan,
			item.Bulan,
		).FirstOrCreate(&item).Error

		if err != nil {
			return nil,err
		}
	}

	log.Println("✅ Seeder Aturan Pelayanan selesai")
	return result, nil
}