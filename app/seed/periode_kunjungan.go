package seed

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"monitoring-service/app/models"
)

func SeederPeriode(db *gorm.DB, kategori map[string]int32) (map[string]int32, error) {

	log.Println("🌱 Seed: periode kunjungan...")

	periodeList := []models.PeriodeKunjungan{
		{Nama: "0-6 jam", KategoriUmurID: kategori["bayi_0_28_hari"], MinValue: 0, MaxValue: 6, Unit: "jam", Urutan: 1},
		{Nama: "6-48 jam", KategoriUmurID: kategori["bayi_0_28_hari"], MinValue: 6, MaxValue: 48, Unit: "jam", Urutan: 2},
		{Nama: "3-7 hari", KategoriUmurID: kategori["bayi_0_28_hari"], MinValue: 3, MaxValue: 7, Unit: "hari", Urutan: 3},
		{Nama: "8-28 hari", KategoriUmurID: kategori["bayi_0_28_hari"], MinValue: 8, MaxValue: 28, Unit: "hari", Urutan: 4},

		{Nama: "0-12 bulan", KategoriUmurID: kategori["bayi_0_1_tahun"], MinValue: 0, MaxValue: 12, Unit: "bulan", Urutan: 5},
		{Nama: "13-24 bulan", KategoriUmurID: kategori["anak_1_2_tahun"], MinValue: 13, MaxValue: 24, Unit: "bulan", Urutan: 6},
		{Nama: "25-36 bulan", KategoriUmurID: kategori["anak_2_3_tahun"], MinValue: 25, MaxValue: 36, Unit: "bulan", Urutan: 7},
		{Nama: "37-48 bulan", KategoriUmurID: kategori["anak_3_4_tahun"], MinValue: 37, MaxValue: 48, Unit: "bulan", Urutan: 8},
		{Nama: "49-60 bulan", KategoriUmurID: kategori["anak_4_5_tahun"], MinValue: 49, MaxValue: 60, Unit: "bulan", Urutan: 9},
		{Nama: "61-72 bulan", KategoriUmurID: kategori["anak_5_6_tahun"], MinValue: 61, MaxValue: 72, Unit: "bulan", Urutan: 10},
	}

	result := map[string]int32{}

	for _, p := range periodeList {

		obj := models.PeriodeKunjungan{}

		err := db.Where("nama = ? AND kategori_umur_id = ?", p.Nama, p.KategoriUmurID).
			Assign(models.PeriodeKunjungan{
				MinValue: p.MinValue,
				MaxValue: p.MaxValue,
				Unit:     p.Unit,
				Urutan:   p.Urutan,
			}).
			FirstOrCreate(&obj, models.PeriodeKunjungan{
				Nama:           p.Nama,
				KategoriUmurID: p.KategoriUmurID,
			}).Error

		if err != nil {
			return nil, err
		}

		key := p.Nama + "_" + fmt.Sprint(p.KategoriUmurID)
		result[key] = obj.ID
	}

	log.Println("✅ Seeder periode selesai")
	return result, nil
}
