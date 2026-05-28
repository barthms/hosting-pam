package seeders

// import (
// 	"log"
// 	"time"

// 	"monitoring-service/app/models" // Sesuaikan dengan nama modul Anda

// 	"gorm.io/gorm"
// )

// type KategoriTandaSakitSeeder struct {
// 	db *gorm.DB
// }

// func NewKategoriTandaSakitSeeder(db *gorm.DB) *KategoriTandaSakitSeeder {
// 	return &KategoriTandaSakitSeeder{db: db}
// }

// func (s *KategoriTandaSakitSeeder) Seed() error {
// 	log.Println("Memulai proses seeding Rentang Usia dan Kategori Tanda Sakit...")

// 	now := time.Now()

// 	// ==========================================
// 	// 1. SEEDING RENTANG USIA
// 	// ==========================================
// 	rentangUsiaData := []models.RentangUsia{
// 		{NamaRentang: "0 - 28 Hari", SatuanWaktu: "Hari"},
// 		{NamaRentang: "29 Hari - 3 Bulan", SatuanWaktu: "Minggu"},
// 		{NamaRentang: "3 - 6 Bulan", SatuanWaktu: "Bulan"},
// 		{NamaRentang: "6 - 12 Bulan", SatuanWaktu: "Bulan"},
// 		{NamaRentang: "12 - 24 Bulan", SatuanWaktu: "Bulan"},
// 		{NamaRentang: "2 - 6 Tahun", SatuanWaktu: "Bulan"},
// 	}

// 	// Simpan map ID untuk digunakan saat seeding KategoriTandaSakit
// 	rentangUsiaIDs := make(map[string]int32)

// 	for _, ru := range rentangUsiaData {
// 		var existing models.RentangUsia
// 		err := s.db.Where("nama_rentang = ?", ru.NamaRentang).First(&existing).Error

// 		if err == gorm.ErrRecordNotFound {
// 			ru.CreatedAt = now
// 			ru.UpdatedAt = now
// 			if createErr := s.db.Create(&ru).Error; createErr != nil {
// 				log.Printf("Gagal membuat RentangUsia (%s): %v\n", ru.NamaRentang, createErr)
// 				return createErr
// 			}
// 			rentangUsiaIDs[ru.NamaRentang] = ru.ID
// 		} else if err == nil {
// 			rentangUsiaIDs[ru.NamaRentang] = existing.ID
// 		} else {
// 			return err
// 		}
// 	}

// 	// ==========================================
// 	// 2. DATA MASTER TANDA SAKIT (Berdasarkan Buku KIA)
// 	// ==========================================
// 	tandaSakitMaster := []struct {
// 		Gejala    string
// 		Deskripsi string
// 	}{
// 		{
// 			Gejala:    "Suhu tubuh panas > 38,5°C",
// 			Deskripsi: "Atau ada tanda pendarahan seperti mimisan, gusi berdarah, muntah kopi, atau BAB hitam.",
// 		},
// 		{
// 			Gejala:    "Sesak napas",
// 			Deskripsi: "Napas sesak disertai cuping hidung kembang kempis atau dada terlihat tertarik ke dalam saat bernapas.",
// 		},
// 		{
// 			Gejala:    "Batuk",
// 			Deskripsi: "Batuk yang disertai dengan bunyi grok-grok atau mengi.",
// 		},
// 		{
// 			Gejala:    "BAB lebih sering / lebih encer",
// 			Deskripsi: "Diare disertai dengan mata cekung, anak haus minum dengan lahap, atau diare disertai darah.",
// 		},
// 		{
// 			Gejala:    "Jumlah air kencing sedikit / tidak kencing",
// 			Deskripsi: "Tidak kencing selama 6 jam, warna kencing kuning pekat, kecoklatan, atau warna pekat lainnya.",
// 		},
// 		{
// 			Gejala:    "Warna kulit tampak biru / memar",
// 			Deskripsi: "Terlihat warna kebiruan atau memar di sekitar daerah mulut, tangan, atau kaki.",
// 		},
// 		{
// 			Gejala:    "Aktivitas tampak lemah",
// 			Deskripsi: "Anak tampak lemah, tidak bergerak, terus menangis, atau merintih.",
// 		},
// 		{
// 			Gejala:    "Hisapan bayi lemah / tidak bergerak",
// 			Deskripsi: "Disertai muntah susu atau cairan hijau, frekuensi kencing kurang dari 6x sehari, atau warna kencing kurang pekat.",
// 		},
// 		{
// 			Gejala:    "Tidak mau makan / minum",
// 			Deskripsi: "Anak sama sekali menolak untuk diberikan makan atau minum.",
// 		},
// 		{
// 			Gejala:    "Berat badan tidak naik",
// 			Deskripsi: "Berat badan anak tidak naik sesuai dengan grafik pertumbuhan.",
// 		},
// 	}

// 	// ==========================================
// 	// 3. SEEDING KATEGORI TANDA SAKIT PER RENTANG USIA
// 	// ==========================================
// 	totalSeededSakit := 0

// 	// Kita melakukan iterasi ke seluruh rentang usia, lalu memasukkan 10 gejala standar untuk masing-masing usia.
// 	for namaRentang, idRentang := range rentangUsiaIDs {
// 		for _, ts := range tandaSakitMaster {
// 			item := models.KategoriTandaSakit{
// 				RentangUsiaID: idRentang,
// 				Gejala:        ts.Gejala,
// 				Deskripsi:     ts.Deskripsi,
// 				IsActive:      true,
// 				CreatedAt:     now,
// 				UpdatedAt:     now,
// 			}

// 			// Cek duplikasi
// 			var existing models.KategoriTandaSakit
// 			err := s.db.Where("rentang_usia_id = ? AND gejala = ?", item.RentangUsiaID, item.Gejala).First(&existing).Error

// 			if err == gorm.ErrRecordNotFound {
// 				if createErr := s.db.Create(&item).Error; createErr != nil {
// 					log.Printf("Gagal membuat KategoriTandaSakit (%s untuk %s): %v\n", item.Gejala, namaRentang, createErr)
// 					return createErr
// 				}
// 				totalSeededSakit++
// 			} else if err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	log.Printf("Seeding Kategori Tanda Sakit selesai! %d data berhasil ditambahkan ke tabel kategori_tanda_sakit.\n", totalSeededSakit)
// 	return nil
// }
