package seeders

// import (
// 	"log"
// 	"time"

// 	"monitoring-service/app/models"

// 	"gorm.io/gorm"
// )

// type KependudukanSeeder struct {
// 	db *gorm.DB
// }

// func NewKependudukanSeeder(db *gorm.DB) *KependudukanSeeder {
// 	return &KependudukanSeeder{db: db}
// }

// func (s *KependudukanSeeder) Seed() error {
// 	log.Println("Starting kependudukan and anak seeding...")

// 	if err := s.seedKartuKeluarga(); err != nil {
// 		return err
// 	}

// 	// Seed kependudukan data
// 	if err := s.seedKependudukan(); err != nil {
// 		return err
// 	}

// 	// Seed ibu data
// 	if err := s.seedIbu(); err != nil {
// 		return err
// 	}

// 	// Seed anak data
// 	if err := s.seedAnak(); err != nil {
// 		return err
// 	}

// 	log.Println("Kependudukan and anak seeding completed!")
// 	return nil
// }

// func (s *KependudukanSeeder) seedKartuKeluarga() error {
// 	now := time.Now()

// 	list := []models.KartuKeluarga{
// 		{
// 			NoKartuKeluarga: 3171234567890101,
// 			TanggalTerbit:   "2018-01-10",
// 			Alamat:          "Dusun Mawar",
// 			RTRW:            "001/002",
// 			CreatedAt:       now,
// 			UpdatedAt:       now,
// 		},
// 		{
// 			NoKartuKeluarga: 3171234567890102,
// 			TanggalTerbit:   "2019-03-21",
// 			Alamat:          "Dusun Melati",
// 			RTRW:            "003/004",
// 			CreatedAt:       now,
// 			UpdatedAt:       now,
// 		},
// 	}

// 	for _, item := range list {
// 		if err := s.db.Where("no_kartu_keluarga = ?", item.NoKartuKeluarga).FirstOrCreate(&item).Error; err != nil {
// 			log.Printf("Error seeding KartuKeluarga (NoKK: %d): %v\n", item.NoKartuKeluarga, err)
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (s *KependudukanSeeder) seedKependudukan() error {
// 	now := time.Now()

// 	var kartuKeluargaList []models.KartuKeluarga
// 	if err := s.db.Where("no_kartu_keluarga IN ?", []int64{3171234567890101, 3171234567890102}).Find(&kartuKeluargaList).Error; err != nil {
// 		return err
// 	}

// 	kkIDByNo := map[int64]uint{}
// 	for _, kk := range kartuKeluargaList {
// 		kkIDByNo[kk.NoKartuKeluarga] = kk.ID
// 	}

// 	kk01ID, ok := kkIDByNo[3171234567890101]
// 	if !ok {
// 		return gorm.ErrRecordNotFound
// 	}
// 	kk02ID, ok := kkIDByNo[3171234567890102]
// 	if !ok {
// 		return gorm.ErrRecordNotFound
// 	}

// 	// Data contoh kependudukan (2 keluarga)
// 	kependudukanList := []models.Kependudukan{
// 		// Keluarga 1
// 		{
// 			NoKartuKeluargaID: &kk01ID,
// 			Nik:               3171234567890001,
// 			Nama:              "Budi Santoso",
// 			JenisKelamin:      "Laki-laki",
// 			TanggalLahir:      "1990-05-10",
// 			TempatLahir:       "Jakarta",
// 			GolonganDarah:     "O",
// 			// HubunganKeluarga:   "Kepala Keluarga",
// 			Dusun:              "Mawar",
// 			Pekerjaan:          "Karyawan Swasta",
// 			PendidikanTerakhir: "S1",
// 			CreatedAt:          now,
// 			UpdatedAt:          now,
// 		},
// 		{
// 			NoKartuKeluargaID: &kk01ID,
// 			Nik:               3171234567890002,
// 			Nama:              "Siti Aminah",
// 			JenisKelamin:      "Perempuan",
// 			TanggalLahir:      "1992-08-20",
// 			TempatLahir:       "Bandung",
// 			GolonganDarah:     "A",
// 			// HubunganKeluarga:   "Istri",
// 			Dusun:              "Mawar",
// 			Pekerjaan:          "Ibu Rumah Tangga",
// 			PendidikanTerakhir: "SMA",
// 			CreatedAt:          now,
// 			UpdatedAt:          now,
// 		},
// 		{
// 			NoKartuKeluargaID: &kk01ID,
// 			Nik:               3171234567890003, // NIK untuk disemai ke tabel Anak
// 			Nama:              "Ahmad Santoso",
// 			JenisKelamin:      "Laki-laki",
// 			TanggalLahir:      "2020-10-15",
// 			TempatLahir:       "Jakarta",
// 			GolonganDarah:     "O",
// 			// HubunganKeluarga:  "Anak",
// 			Dusun:     "Mawar",
// 			CreatedAt: now,
// 			UpdatedAt: now,
// 		},
// 		{
// 			NoKartuKeluargaID: &kk01ID,
// 			Nik:               3171234567890004, // NIK untuk disemai ke tabel Anak
// 			Nama:              "Rina Santoso",
// 			JenisKelamin:      "Perempuan",
// 			TanggalLahir:      "2023-01-25",
// 			TempatLahir:       "Jakarta",
// 			GolonganDarah:     "A",
// 			// HubunganKeluarga:  "Anak",
// 			Dusun:     "Mawar",
// 			CreatedAt: now,
// 			UpdatedAt: now,
// 		},
// 		{
// 			NoKartuKeluargaID: &kk01ID,
// 			Nik:               3171234567890005, // NIK untuk disemai ke tabel Anak
// 			Nama:              "Nahes Purba",
// 			JenisKelamin:      "Laki-laki",
// 			TanggalLahir:      "2025-12-09",
// 			TempatLahir:       "Jakarta",
// 			GolonganDarah:     "A",
// 			// HubunganKeluarga:  "Anak",
// 			Dusun:     "Luman Batu",
// 			CreatedAt: now,
// 			UpdatedAt: now,
// 		},
// 		// Keluarga 2
// 		{
// 			NoKartuKeluargaID: &kk02ID,
// 			Nik:               3171234567890011,
// 			Nama:              "Andi Pratama",
// 			JenisKelamin:      "Laki-laki",
// 			TanggalLahir:      "1989-01-17",
// 			TempatLahir:       "Jakarta",
// 			GolonganDarah:     "B",
// 			// HubunganKeluarga:   "Kepala Keluarga",
// 			Dusun:              "Melati",
// 			Pekerjaan:          "Wiraswasta",
// 			PendidikanTerakhir: "S1",
// 			CreatedAt:          now,
// 			UpdatedAt:          now,
// 		},
// 		{
// 			NoKartuKeluargaID: &kk02ID,
// 			Nik:               3171234567890012,
// 			Nama:              "Dewi Lestari",
// 			JenisKelamin:      "Perempuan",
// 			TanggalLahir:      "1991-04-03",
// 			TempatLahir:       "Bogor",
// 			GolonganDarah:     "AB",
// 			// HubunganKeluarga:   "Istri",
// 			Dusun:              "Melati",
// 			Pekerjaan:          "Ibu Rumah Tangga",
// 			PendidikanTerakhir: "SMA",
// 			CreatedAt:          now,
// 			UpdatedAt:          now,
// 		},
// 		{
// 			NoKartuKeluargaID: &kk02ID,
// 			Nik:               3171234567890013,
// 			Nama:              "Kania Pratama",
// 			JenisKelamin:      "Perempuan",
// 			TanggalLahir:      "2022-07-19",
// 			TempatLahir:       "Jakarta",
// 			GolonganDarah:     "AB",
// 			// HubunganKeluarga:  "Anak",
// 			Dusun:     "Melati",
// 			CreatedAt: now,
// 			UpdatedAt: now,
// 		},
// 		{
// 			NoKartuKeluargaID: &kk02ID,
// 			Nik:               3171234567890014,
// 			Nama:              "Rafi Pratama",
// 			JenisKelamin:      "Laki-laki",
// 			TanggalLahir:      "2024-11-01",
// 			TempatLahir:       "Jakarta",
// 			GolonganDarah:     "B",
// 			// HubunganKeluarga:  "Anak",
// 			Dusun:     "Melati",
// 			CreatedAt: now,
// 			UpdatedAt: now,
// 		},
// 	}

// 	for _, k := range kependudukanList {
// 		if err := s.db.Where("nik = ?", k.Nik).FirstOrCreate(&k).Error; err != nil {
// 			log.Printf("Error seeding Kependudukan (NIK: %d): %v\n", k.Nik, err)
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (s *KependudukanSeeder) seedIbu() error {
// 	ibuNikList := []int64{3171234567890002, 3171234567890012}

// 	for _, ibuNik := range ibuNikList {
// 		var kepIbu models.Kependudukan
// 		if err := s.db.Where("nik = ?", ibuNik).First(&kepIbu).Error; err != nil {
// 			if err == gorm.ErrRecordNotFound {
// 				log.Printf("Kependudukan ibu dengan NIK %d tidak ditemukan, melewati seed Ibu\n", ibuNik)
// 				continue
// 			}
// 			return err
// 		}

// 		kepID := kepIbu.ID
// 		ibu := models.Ibu{KependudukanID: &kepID}

// 		if err := s.db.Where("kependudukan_id = ?", kepID).FirstOrCreate(&ibu).Error; err != nil {
// 			log.Printf("Error seeding Ibu (KependudukanID: %d): %v\n", kepID, err)
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (s *KependudukanSeeder) seedAnak() error {
// 	now := time.Now()

// 	type anakSeedItem struct {
// 		BeratLahir   float64
// 		TinggiLahir  float64
// 		IbuNik       int64
// 		NamaAnak     string
// 		JenisKelamin string
// 		TanggalLahir string
// 		NoKK         int64
// 	}

// 	// Mapping berat dan tinggi anak berdasarkan NIK
// 	dataAnakMap := map[int64]anakSeedItem{
// 		3171234567890003: {BeratLahir: 3.2, TinggiLahir: 50.0, IbuNik: 3171234567890002, NamaAnak: "Ahmad Santoso", JenisKelamin: "Laki-laki", TanggalLahir: "2020-10-15", NoKK: 3171234567890101},
// 		3171234567890004: {BeratLahir: 2.9, TinggiLahir: 48.5, IbuNik: 3171234567890002, NamaAnak: "Rina Santoso", JenisKelamin: "Perempuan", TanggalLahir: "2023-01-25", NoKK: 3171234567890101},
// 		3171234567890005: {BeratLahir: 3.1, TinggiLahir: 49.0, IbuNik: 3171234567890002, NamaAnak: "Nahes Purba", JenisKelamin: "Laki-laki", TanggalLahir: "2025-12-09", NoKK: 3171234567890101},
// 		3171234567890013: {BeratLahir: 3.0, TinggiLahir: 49.5, IbuNik: 3171234567890012, NamaAnak: "Kania Pratama", JenisKelamin: "Perempuan", TanggalLahir: "2022-07-19", NoKK: 3171234567890102},
// 		3171234567890014: {BeratLahir: 3.3, TinggiLahir: 50.2, IbuNik: 3171234567890012, NamaAnak: "Rafi Pratama", JenisKelamin: "Laki-laki", TanggalLahir: "2024-11-01", NoKK: 3171234567890102},
// 	}

// 	for nik, detailAnak := range dataAnakMap {
// 		var kartuKeluarga models.KartuKeluarga
// 		if detailAnak.NoKK > 0 {
// 			if err := s.db.Where("no_kartu_keluarga = ?", detailAnak.NoKK).First(&kartuKeluarga).Error; err != nil {
// 				if err != gorm.ErrRecordNotFound {
// 					return err
// 				}
// 			}
// 		}

// 		var kependudukan models.Kependudukan
// 		var kependudukanID *uint

// 		// Cari ID Kependudukan berdasarkan NIK
// 		if err := s.db.Where("nik = ?", nik).First(&kependudukan).Error; err != nil {
// 			if err != gorm.ErrRecordNotFound {
// 				return err
// 			}
// 			log.Printf("Kependudukan dengan NIK %d tidak ditemukan, lanjut pakai data anak fallback\n", nik)
// 		} else {
// 			kID := kependudukan.ID
// 			kependudukanID = &kID
// 		}

// 		var ibu models.Ibu
// 		if detailAnak.IbuNik > 0 {
// 			var kepIbu models.Kependudukan
// 			if err := s.db.Where("nik = ?", detailAnak.IbuNik).First(&kepIbu).Error; err != nil {
// 				log.Printf("Kependudukan ibu dengan NIK %d tidak ditemukan, melewati seed Anak NIK %d\n", detailAnak.IbuNik, nik)
// 				continue
// 			}

// 			if err := s.db.Where("kependudukan_id = ?", kepIbu.ID).First(&ibu).Error; err != nil {
// 				log.Printf("Ibu dengan NIK %d tidak ditemukan, melewati seed Anak NIK %d\n", detailAnak.IbuNik, nik)
// 				continue
// 			}
// 		}

// 		var existing models.Anak
// 		var err error
// 		// Periksa apakah data anak untuk KependudukanID ini sudah ada
// 		if kependudukanID != nil {
// 			err = s.db.Where("kependudukan_id = ?", *kependudukanID).First(&existing).Error
// 		} else {
// 			err = s.db.Where("nama_anak = ? AND tanggal_lahir = ?", detailAnak.NamaAnak, detailAnak.TanggalLahir).First(&existing).Error
// 		}

// 		if err != nil {
// 			if err == gorm.ErrRecordNotFound {
// 				var ibuID *uint
// 				if ibu.ID > 0 {
// 					ibuID = &ibu.ID
// 				}

// 				anak := models.Anak{
// 					IbuID:           ibuID,
// 					KependudukanID:  kependudukanID,
// 					NoKartuKeluarga: kartuKeluarga.NoKartuKeluarga,
// 					NamaAnak:        detailAnak.NamaAnak,
// 					JenisKelamin:    detailAnak.JenisKelamin,
// 					TanggalLahir:    detailAnak.TanggalLahir,
// 					BeratLahir:      detailAnak.BeratLahir,
// 					TinggiLahir:     detailAnak.TinggiLahir,
// 					CreatedAt:       now,
// 					UpdatedAt:       now,
// 				}

// 				if err := s.db.Create(&anak).Error; err != nil {
// 					log.Printf("Error creating Anak (Nama: %s): %v\n", anak.NamaAnak, err)
// 					return err
// 				}
// 			} else {
// 				return err
// 			}
// 		}
// 	}

// 	return nil
// }
