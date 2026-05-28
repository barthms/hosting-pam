// cmd/migrate/add_penduduk_id_to_pengguna.go
package main

// import (
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func main() {
// 	// Load .env
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Println("⚠️ Tidak dapat memuat .env, gunakan variabel environment")
// 	}

// 	// Ambil DSN dari environment
// 	dsn := os.Getenv("DB_POSTGRES_DSN")
// 	if dsn == "" {
// 		log.Fatal("❌ DB_POSTGRES_DSN tidak ditemukan di .env")
// 	}
// 	log.Println("🔌 Mencoba koneksi ke database Supabase...")

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("❌ Gagal koneksi database: %v", err)
// 	}
// 	log.Println("✅ Koneksi database berhasil")

// 	// Cek apakah kolom penduduk_id sudah ada
// 	// Kita gunakan raw SQL karena gorm Migrator().HasColumn membutuhkan struct
// 	var exists bool
// 	err = db.Raw(`
// 		SELECT EXISTS (
// 			SELECT 1 
// 			FROM information_schema.columns 
// 			WHERE table_name = 'pengguna' 
// 			AND column_name = 'penduduk_id'
// 		)
// 	`).Scan(&exists).Error
// 	if err != nil {
// 		log.Fatalf("❌ Gagal mengecek kolom: %v", err)
// 	}

// 	if !exists {
// 		log.Println("➕ Menambahkan kolom penduduk_id ke tabel pengguna...")
// 		// Tambah kolom penduduk_id (nullable bigint)
// 		if err := db.Exec("ALTER TABLE pengguna ADD COLUMN IF NOT EXISTS penduduk_id BIGINT NULL").Error; err != nil {
// 			log.Fatalf("❌ Gagal menambah kolom penduduk_id: %v", err)
// 		}
// 		log.Println("✅ Kolom penduduk_id berhasil ditambahkan")
// 	} else {
// 		log.Println("✅ Kolom penduduk_id sudah ada")
// 	}

// 	log.Println("🎉 Migrasi selesai. Sekarang registrasi dapat berjalan.")
// }
