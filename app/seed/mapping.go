package seed

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"monitoring-service/app/models"
)

type Mapping struct {
	Kategori string
	Periode  string
	Items    []string
}

func SeederMapping(db *gorm.DB, kategori map[string]int32, pelayanan map[string]int32, periode map[string]int32,aturanMap map[string]int32,) error {

	log.Println("🌱 Seed: mapping jenis pelayanan ke kategori...")

	mappings := []Mapping{
		{Kategori: "bayi_0_28_hari", Periode: "0-6 jam", Items: []string{
			"Perawatan Tali Pusat", "IMD", "Vitamin K1", "Imunisasi Hepatitis B",
			"Salep/Tetes Mata Antibiotik", "KIA", "Tripel Eliminasi",
			"Kondisi Umum", "BB", "PB", "LK", "Masalah", "Dirujuk",
		}},
		{Kategori: "bayi_0_28_hari", Periode: "6-48 jam", Items: []string{
			"BB", "PB", "LK", "Menyusu", "Perawatan Tali Pusat",
			"Imunisasi Hepatitis B", "Salep/Tetes Mata Antibiotik", "Vitamin K1",
			"Skrining Hipotiroid Kongenital", "Hasil Skrining Hipotiroid Kongenital",
			"Masalah", "Dirujuk", "Tripel Eliminasi",
		}},
		{Kategori: "bayi_0_28_hari", Periode: "3-7 hari", Items: []string{
			"Menyusu", "Perawatan Tali Pusat", "Imunisasi Hepatitis B",
			"Tanda Bahaya", "Identifikasi Kuning",
			"Skrining Hipotiroid Kongenital", "Masalah", "Dirujuk", "Tripel Eliminasi",
		}},
		{Kategori: "bayi_0_28_hari", Periode: "8-28 hari", Items: []string{
			"Perawatan Tali Pusat", "KIA", "Tripel Eliminasi",
		}},
		{Kategori: "bayi_0_1_tahun", Periode: "0-12 bulan", Items: []string{
			"BB", "PB", "LK", "Perkembangan", "KIE", "Imunisasi", "Vitamin A", "Tripel Eliminasi",
		}},
	}

	urutan := 1
	for _, m := range mappings {
		for _, item := range m.Items {
			// pastikan item ada di map pelayanan
			pID, ok := pelayanan[item]
			if !ok {
				log.Printf("⚠️ Pelayanan '%s' tidak ditemukan di map", item)
				continue
			}

			// key periode sesuai format SeederPeriode
			periodeKey := fmt.Sprintf("%s_%d", m.Periode, kategori[m.Kategori])
			pIDPeriode, ok := periode[periodeKey]
			if !ok || pIDPeriode == 0 {
				log.Printf("❌ Periode ID tidak ditemukan untuk key: %s", periodeKey)
				continue
			}

			obj := models.JenisPelayananKategori{
				JenisPelayananID: pID,
				KategoriUmurID:   kategori[m.Kategori],
				PeriodeID:        pIDPeriode,
				Urutan:           urutan,
				IsActive:         true,
			}

			if err := db.Where("jenis_pelayanan_id = ? AND kategori_umur_id = ? AND periode_id = ?",
				obj.JenisPelayananID, obj.KategoriUmurID, obj.PeriodeID).
				FirstOrCreate(&obj).Error; err != nil {
				return err
			}
			urutan++
		}
	}

	log.Println("✅ Seeder mapping selesai")
	return nil
}
