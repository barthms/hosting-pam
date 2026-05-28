package seed

import (
	"log"

	"gorm.io/gorm"

	"monitoring-service/app/models"
)

func SeederKategoriUmur(db *gorm.DB) (map[string]int32, error) {
	log.Println("🌱 Seed: kategori umur...")

	kategoriList := []models.KategoriUmur{
		{KategoriUmur: "bayi_0_28_hari", MinValue: 0, MaxValue: 28, Unit: "hari"},
		{KategoriUmur: "bayi_0_1_tahun", MinValue: 0, MaxValue: 12, Unit: "bulan"},
		{KategoriUmur: "bayi_6_11_bulan", MinValue: 6, MaxValue: 11, Unit: "bulan"},
		{KategoriUmur: "anak_1_2_tahun", MinValue: 1, MaxValue: 2, Unit: "tahun"},
		{KategoriUmur: "anak_2_3_tahun", MinValue: 2, MaxValue: 3, Unit: "tahun"},
		{KategoriUmur: "anak_3_4_tahun", MinValue: 3, MaxValue: 4, Unit: "tahun"},
		{KategoriUmur: "anak_4_5_tahun", MinValue: 4, MaxValue: 5, Unit: "tahun"},
		{KategoriUmur: "anak_5_6_tahun", MinValue: 5, MaxValue: 6, Unit: "tahun"},
	}

	result := map[string]int32{}

	for _, k := range kategoriList {
		obj := models.KategoriUmur{}

		err := db.Where("kategori_umur = ?", k.KategoriUmur).First(&obj).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// create baru
				obj = models.KategoriUmur{
					KategoriUmur: k.KategoriUmur,
					MinValue:     k.MinValue,
					MaxValue:     k.MaxValue,
					Unit:         k.Unit,
				}
				if err := db.Create(&obj).Error; err != nil {
					return nil, err
				}
			} else {
				return nil, err
			}
		} else {
			// update jika sudah ada
			db.Model(&obj).Updates(models.KategoriUmur{
				MinValue: k.MinValue,
				MaxValue: k.MaxValue,
				Unit:     k.Unit,
			})
		}

		result[k.KategoriUmur] = obj.ID
	}

	log.Println("✅ Seeder kategori umur selesai")
	return result, nil
}
