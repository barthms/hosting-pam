package main

// import (
// 	"errors"
// 	"log"
// 	"monitoring-service/app/models"
// 	"os"
// 	"time"

// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// // Struktur sementara untuk membaca data dari tabel lama
// type LanjutanTemp struct {
// 	ID                                    int32
// 	KehamilanID                           int32
// 	HasilUSGCatatan                       string
// 	TanggalLab                            *time.Time
// 	LabHemoglobinHasil                    *float64
// 	LabHemoglobinRencana                  string
// 	LabProteinUrinHasil                   *int
// 	LabProteinUrinRencana                 string
// 	LabUrinReduksiHasil                   string
// 	LabUrinReduksiRencana                 string
// 	TanggalSkriningJiwa                   *time.Time
// 	SkriningJiwaHasil                     string
// 	SkriningJiwaTindakLanjut              string
// 	SkriningJiwaPerluRujukan              string
// 	RencanaKonsultasiGizi                 bool
// 	RencanaKonsultasiKebidanan            bool
// 	RencanaKonsultasiAnak                 bool
// 	RencanaKonsultasiPenyakitDalam        bool
// 	RencanaKonsultasiNeurologi            bool
// 	RencanaKonsultasiTHT                  bool
// 	RencanaKonsultasiPsikiatri            bool
// 	RencanaKonsultasiLainLain             string
// 	RencanaProsesMelahirkan               string
// 	RencanaKontrasepsiAKDR                bool
// 	RencanaKontrasepsiPil                 bool
// 	RencanaKontrasepsiSuntik              bool
// 	RencanaKontrasepsiSteril              bool
// 	RencanaKontrasepsiMAL                 bool
// 	RencanaKontrasepsiImplan              bool
// 	RencanaKontrasepsiBelumMemilih        bool
// 	KebutuhanKonseling                    string
// 	Penjelasan                            string
// 	KesimpulanRekomendasiTempatMelahirkan string
// }

// func main() {
// 	// Load file .env
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Println("⚠️ Peringatan: tidak dapat memuat .env, gunakan variabel environment sistem")
// 	}

// 	// Ambil DSN dari environment
// 	dsn := os.Getenv("DB_POSTGRES_DSN")
// 	if dsn == "" {
// 		log.Fatal("❌ DB_POSTGRES_DSN tidak ditemukan di .env")
// 	}
// 	log.Println("🔌 Mencoba koneksi ke database...")

// 	// Koneksi ke database
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("❌ Gagal koneksi database: %v", err)
// 	}
// 	log.Println("✅ Koneksi database berhasil")

// 	// ---- Tambah kolom trimester ke tabel pemeriksaan_laboratorium_jiwa jika belum ada ----
// 	log.Println("🔄 Mengecek & menambah kolom 'trimester' pada tabel pemeriksaan_laboratorium_jiwa...")
// 	if !db.Migrator().HasColumn(&models.PemeriksaanLaboratoriumJiwa{}, "trimester") {
// 		if err := db.Migrator().AddColumn(&models.PemeriksaanLaboratoriumJiwa{}, "trimester"); err != nil {
// 			log.Fatalf("❌ Gagal menambah kolom trimester: %v", err)
// 		}
// 		log.Println("✅ Kolom trimester berhasil ditambahkan")
// 	} else {
// 		log.Println("✅ Kolom trimester sudah ada")
// 	}

// 	// Set default trimester = 1 untuk data yang null
// 	db.Exec("UPDATE pemeriksaan_laboratorium_jiwa SET trimester = 1 WHERE trimester IS NULL OR trimester = 0")

// 	// ---- Tambah kolom dari PemeriksaanLanjutanTrimester3 ke PemeriksaanDokterTrimester3 ----
// 	log.Println("🔄 Mengecek & menambah kolom lanjutan pada tabel pemeriksaan_dokter_trimester_3...")
// 	if err := db.AutoMigrate(&models.PemeriksaanDokterTrimester3{}); err != nil {
// 		log.Fatalf("❌ Gagal auto migrate: %v", err)
// 	}
// 	log.Println("✅ Kolom lanjutan berhasil ditambahkan")

// 	// ---- Migrasi data dari tabel lama ke tabel dokter T3 ----
// 	log.Println("📦 Memindahkan data dari pemeriksaan_lanjutan_trimester_3 ke pemeriksaan_dokter_trimester_3...")

// 	var lanjutanList []LanjutanTemp
// 	if err := db.Table("pemeriksaan_lanjutan_trimester_3").Find(&lanjutanList).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			log.Println("ℹ️ Tidak ada data di tabel pemeriksaan_lanjutan_trimester_3")
// 		} else {
// 			log.Printf("⚠️ Gagal membaca tabel lanjutan: %v (mungkin tabel belum ada)", err)
// 		}
// 	} else {
// 		log.Printf("📊 Ditemukan %d data lanjutan T3", len(lanjutanList))
// 		for _, l := range lanjutanList {
// 			var dokter models.PemeriksaanDokterTrimester3
// 			err := db.Where("kehamilan_id = ?", l.KehamilanID).First(&dokter).Error
// 			if errors.Is(err, gorm.ErrRecordNotFound) {
// 				// Buat record baru di tabel dokter T3
// 				dokter = models.PemeriksaanDokterTrimester3{
// 					KehamilanID: l.KehamilanID,
// 					// field lanjutan
// 					HasilUSGCatatan:                       l.HasilUSGCatatan,
// 					TanggalLab:                            l.TanggalLab,
// 					LabHemoglobinHasil:                    l.LabHemoglobinHasil,
// 					LabHemoglobinRencana:                  l.LabHemoglobinRencana,
// 					LabProteinUrinHasil:                   l.LabProteinUrinHasil,
// 					LabProteinUrinRencana:                 l.LabProteinUrinRencana,
// 					LabUrinReduksiHasil:                   l.LabUrinReduksiHasil,
// 					LabUrinReduksiRencana:                 l.LabUrinReduksiRencana,
// 					TanggalSkriningJiwa:                   l.TanggalSkriningJiwa,
// 					SkriningJiwaHasil:                     l.SkriningJiwaHasil,
// 					SkriningJiwaTindakLanjut:              l.SkriningJiwaTindakLanjut,
// 					SkriningJiwaPerluRujukan:              l.SkriningJiwaPerluRujukan,
// 					RencanaKonsultasiGizi:                 l.RencanaKonsultasiGizi,
// 					RencanaKonsultasiKebidanan:            l.RencanaKonsultasiKebidanan,
// 					RencanaKonsultasiAnak:                 l.RencanaKonsultasiAnak,
// 					RencanaKonsultasiPenyakitDalam:        l.RencanaKonsultasiPenyakitDalam,
// 					RencanaKonsultasiNeurologi:            l.RencanaKonsultasiNeurologi,
// 					RencanaKonsultasiTHT:                  l.RencanaKonsultasiTHT,
// 					RencanaKonsultasiPsikiatri:            l.RencanaKonsultasiPsikiatri,
// 					RencanaKonsultasiLainLain:             l.RencanaKonsultasiLainLain,
// 					RencanaProsesMelahirkan:               l.RencanaProsesMelahirkan,
// 					RencanaKontrasepsiAKDR:                l.RencanaKontrasepsiAKDR,
// 					RencanaKontrasepsiPil:                 l.RencanaKontrasepsiPil,
// 					RencanaKontrasepsiSuntik:              l.RencanaKontrasepsiSuntik,
// 					RencanaKontrasepsiSteril:              l.RencanaKontrasepsiSteril,
// 					RencanaKontrasepsiMAL:                 l.RencanaKontrasepsiMAL,
// 					RencanaKontrasepsiImplan:              l.RencanaKontrasepsiImplan,
// 					RencanaKontrasepsiBelumMemilih:        l.RencanaKontrasepsiBelumMemilih,
// 					KebutuhanKonseling:                    l.KebutuhanKonseling,
// 					Penjelasan:                            l.Penjelasan,
// 					KesimpulanRekomendasiTempatMelahirkan: l.KesimpulanRekomendasiTempatMelahirkan,
// 				}
// 				if err := db.Create(&dokter).Error; err != nil {
// 					log.Printf("❌ Gagal membuat record baru untuk kehamilan %d: %v", l.KehamilanID, err)
// 				} else {
// 					log.Printf("✅ Record baru untuk kehamilan %d", l.KehamilanID)
// 				}
// 			} else if err != nil {
// 				log.Printf("⚠️ Error query kehamilan %d: %v", l.KehamilanID, err)
// 			} else {
// 				// Update record yang sudah ada
// 				updates := map[string]interface{}{
// 					"hasil_usg_catatan":                        l.HasilUSGCatatan,
// 					"tanggal_lab":                              l.TanggalLab,
// 					"lab_hemoglobin_hasil":                     l.LabHemoglobinHasil,
// 					"lab_hemoglobin_rencana_tindak_lanjut":     l.LabHemoglobinRencana,
// 					"lab_protein_urin_hasil":                   l.LabProteinUrinHasil,
// 					"lab_protein_urin_rencana_tindak_lanjut":   l.LabProteinUrinRencana,
// 					"lab_urin_reduksi_hasil":                   l.LabUrinReduksiHasil,
// 					"lab_urin_reduksi_rencana_tindak_lanjut":   l.LabUrinReduksiRencana,
// 					"tanggal_skrining_jiwa":                    l.TanggalSkriningJiwa,
// 					"skrining_jiwa_hasil":                      l.SkriningJiwaHasil,
// 					"skrining_jiwa_tindak_lanjut":              l.SkriningJiwaTindakLanjut,
// 					"skrining_jiwa_perlu_rujukan":              l.SkriningJiwaPerluRujukan,
// 					"rencana_konsultasi_gizi":                  l.RencanaKonsultasiGizi,
// 					"rencana_konsultasi_kebidanan":             l.RencanaKonsultasiKebidanan,
// 					"rencana_konsultasi_anak":                  l.RencanaKonsultasiAnak,
// 					"rencana_konsultasi_penyakit_dalam":        l.RencanaKonsultasiPenyakitDalam,
// 					"rencana_konsultasi_neurologi":             l.RencanaKonsultasiNeurologi,
// 					"rencana_konsultasi_tht":                   l.RencanaKonsultasiTHT,
// 					"rencana_konsultasi_psikiatri":             l.RencanaKonsultasiPsikiatri,
// 					"rencana_konsultasi_lain_lain":             l.RencanaKonsultasiLainLain,
// 					"rencana_proses_melahirkan":                l.RencanaProsesMelahirkan,
// 					"rencana_kontrasepsi_akdr":                 l.RencanaKontrasepsiAKDR,
// 					"rencana_kontrasepsi_pil":                  l.RencanaKontrasepsiPil,
// 					"rencana_kontrasepsi_suntik":               l.RencanaKontrasepsiSuntik,
// 					"rencana_kontrasepsi_steril":               l.RencanaKontrasepsiSteril,
// 					"rencana_kontrasepsi_mal":                  l.RencanaKontrasepsiMAL,
// 					"rencana_kontrasepsi_implan":               l.RencanaKontrasepsiImplan,
// 					"rencana_kontrasepsi_belum_memilih":        l.RencanaKontrasepsiBelumMemilih,
// 					"kebutuhan_konseling":                      l.KebutuhanKonseling,
// 					"penjelasan":                               l.Penjelasan,
// 					"kesimpulan_rekomendasi_tempat_melahirkan": l.KesimpulanRekomendasiTempatMelahirkan,
// 				}
// 				if err := db.Model(&dokter).Updates(updates).Error; err != nil {
// 					log.Printf("❌ Gagal update dokter T3 kehamilan %d: %v", l.KehamilanID, err)
// 				} else {
// 					log.Printf("✅ Update dokter T3 kehamilan %d", l.KehamilanID)
// 				}
// 			}
// 		}
// 	}

// 	log.Println("🎉 Migrasi selesai. Database sudah siap untuk perubahan kode.")
// }
