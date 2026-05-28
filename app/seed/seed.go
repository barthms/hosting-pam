package seed

import (
	"log"

	"gorm.io/gorm"
)

func RunAllSeed(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {

		log.Println("🚀 Start seeding...")

		// // 1. Role
		// if err := SeederRole(tx); err != nil {
		// 	log.Println("❌ SeederRole failed:", err)
		// 	return err
		// }

		// // 2. Kategori Umur
		// kategoriMap, err := SeederKategoriUmur(tx)
		// if err != nil {
		// 	log.Println("❌ SeederKategoriUmur failed:", err)
		// 	return err
		// }

		// // 3. Periode
		// periodeMap, err := SeederPeriode(tx, kategoriMap)
		// if err != nil {
		// 	log.Println("❌ SeederPeriode failed:", err)
		// 	return err
		// }

		// // 4. Jenis Pelayanan
		// pelayananMap, err := SeederJenisPelayanan(tx)
		// if err != nil {
		// 	log.Println("❌ SeederJenisPelayanan failed:", err)
		// 	return err
		// }

		// //.aturan_seed
		// aturanMap, err := SeederAturanPelayanan(tx)
		// if err != nil {
		// 	log.Println("❌ SeederAturanPelayanan failed:", err)
		// 	return err
		// }

		// // 5. Mapping (perbaiki urutan parameter)
		// if err := SeederMapping(tx, kategoriMap, pelayananMap, periodeMap, aturanMap); err != nil {
		// 	log.Println("❌ SeederMapping failed:", err)
		// 	return err
		// }

		// 6. Kategori Capaian Perkembangan Anak
		if err := SeederKategoriCapaian(tx); err != nil {
			log.Println("❌ SeederKategoriCapaian failed:", err)
			return err
		}

		// // 7. Kategori Tanda Bahaya
		// if err := seeders.SeedKategoriTandaBahaya(tx); err != nil {
		// 	log.Println("❌ SeedKategoriTandaBahaya failed:", err)
		// 	return err
		// }

		// 6. Edukasi ibu hamil & nifas
		// if err := SeederEdukasiIbuHamil(tx); err != nil {
		// 	log.Println("❌ SeederEdukasiIbuHamil failed:", err)
		// 	return err
		// }

		log.Println("✅ All seeding success")
		return nil
	})
}
