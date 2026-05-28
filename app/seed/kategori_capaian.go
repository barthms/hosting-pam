package seed

import (
	"log"

	"gorm.io/gorm"

	"monitoring-service/app/models"
)

func SeederKategoriCapaian(db *gorm.DB) error {
	log.Println("🌱 Seed: kategori capaian perkembangan anak...")

	data := []models.KategoriCapaian{
		// ===== 0-12 Bulan =====
		{RentangUsia: "0-12 Bulan", PertanyaaanCeklist: "Bayi dapat mengikuti benda dengan matanya", Aspek: "Perkembangan Motorik"},
		{RentangUsia: "0-12 Bulan", PertanyaaanCeklist: "Bayi dapat membuat suara (berceloteh)", Aspek: "Perkembangan Bahasa"},
		{RentangUsia: "0-12 Bulan", PertanyaaanCeklist: "Bayi dapat memegang benda (grasp reflex)", Aspek: "Perkembangan Motorik"},
		{RentangUsia: "0-12 Bulan", PertanyaaanCeklist: "Bayi dapat tersenyum", Aspek: "Perkembangan Sosial Emosional"},
		{RentangUsia: "0-12 Bulan", PertanyaaanCeklist: "Bayi dapat mengangkat kepala (head control)", Aspek: "Perkembangan Motorik"},
		{RentangUsia: "0-12 Bulan", PertanyaaanCeklist: "Bayi dapat merespons suara", Aspek: "Perkembangan Bahasa"},
		{RentangUsia: "0-12 Bulan", PertanyaaanCeklist: "Bayi dapat berguling dari belakang ke depan", Aspek: "Perkembangan Motorik"},
		{RentangUsia: "0-12 Bulan", PertanyaaanCeklist: "Bayi dapat duduk dengan sedikit bantuan", Aspek: "Perkembangan Motorik"},

		// ===== 1-2 Tahun =====
		{RentangUsia: "1-2 Tahun", PertanyaaanCeklist: "Anak dapat berjalan sendiri", Aspek: "Perkembangan Motorik"},
		{RentangUsia: "1-2 Tahun", PertanyaaanCeklist: "Anak dapat menunjuk dengan jari", Aspek: "Perkembangan Kognitif"},
		{RentangUsia: "1-2 Tahun", PertanyaaanCeklist: "Anak dapat mengatakan beberapa kata", Aspek: "Perkembangan Bahasa"},
		{RentangUsia: "1-2 Tahun", PertanyaaanCeklist: "Anak dapat makan sendiri dengan makanan lembut", Aspek: "Perkembangan Motorik Halus"},
		{RentangUsia: "1-2 Tahun", PertanyaaanCeklist: "Anak dapat bermain dengan mainan (misal: memasukkan bola ke lubang)", Aspek: "Perkembangan Kognitif"},
		{RentangUsia: "1-2 Tahun", PertanyaaanCeklist: "Anak dapat memahami perintah sederhana", Aspek: "Perkembangan Bahasa"},
		{RentangUsia: "1-2 Tahun", PertanyaaanCeklist: "Anak dapat menunjukkan kasih sayang", Aspek: "Perkembangan Sosial Emosional"},
		{RentangUsia: "1-2 Tahun", PertanyaaanCeklist: "Anak dapat mengikuti perawatan pribadi (misalnya cuci tangan)", Aspek: "Perkembangan Sosial Emosional"},

		// ===== 2-3 Tahun =====
		{RentangUsia: "2-3 Tahun", PertanyaaanCeklist: "Anak dapat berlari dan melompat", Aspek: "Perkembangan Motorik"},
		{RentangUsia: "2-3 Tahun", PertanyaaanCeklist: "Anak dapat berbicara dengan 2-3 kata", Aspek: "Perkembangan Bahasa"},
		{RentangUsia: "2-3 Tahun", PertanyaaanCeklist: "Anak dapat menggambar garis vertikal", Aspek: "Perkembangan Motorik Halus"},
		{RentangUsia: "2-3 Tahun", PertanyaaanCeklist: "Anak dapat mengganti pakaian dengan bantuan", Aspek: "Perkembangan Motorik Halus"},
		{RentangUsia: "2-3 Tahun", PertanyaaanCeklist: "Anak dapat bermain dengan anak lain (awal sosialisasi)", Aspek: "Perkembangan Sosial Emosional"},
		{RentangUsia: "2-3 Tahun", PertanyaaanCeklist: "Anak dapat mengikuti instruksi 2 langkah", Aspek: "Perkembangan Kognitif"},
		{RentangUsia: "2-3 Tahun", PertanyaaanCeklist: "Anak dapat menunjukkan emosi (senang, sedih, marah)", Aspek: "Perkembangan Sosial Emosional"},
		{RentangUsia: "2-3 Tahun", PertanyaaanCeklist: "Anak dapat meniru aktivitas sehari-hari (bermain rumahan)", Aspek: "Perkembangan Kognitif"},

		// ===== 3-4 Tahun =====
		{RentangUsia: "3-4 Tahun", PertanyaaanCeklist: "Anak dapat bermain dengan satu kaki", Aspek: "Perkembangan Motorik"},
		{RentangUsia: "3-4 Tahun", PertanyaaanCeklist: "Anak dapat berbicara dengan kalimat 3-4 kata", Aspek: "Perkembangan Bahasa"},
		{RentangUsia: "3-4 Tahun", PertanyaaanCeklist: "Anak dapat menggambar lingkaran", Aspek: "Perkembangan Motorik Halus"},
		{RentangUsia: "3-4 Tahun", PertanyaaanCeklist: "Anak dapat mengenakan pakaian dengan bantuan minimal", Aspek: "Perkembangan Motorik Halus"},
		{RentangUsia: "3-4 Tahun", PertanyaaanCeklist: "Anak dapat bermain dengan anak lain dalam permainan sederhana", Aspek: "Perkembangan Sosial Emosional"},
		{RentangUsia: "3-4 Tahun", PertanyaaanCeklist: "Anak dapat memahami cerita sederhana", Aspek: "Perkembangan Kognitif"},
		{RentangUsia: "3-4 Tahun", PertanyaaanCeklist: "Anak dapat menunjukkan kepribadian yang jelas", Aspek: "Perkembangan Sosial Emosional"},
		{RentangUsia: "3-4 Tahun", PertanyaaanCeklist: "Anak dapat mengikuti instruksi 3 langkah", Aspek: "Perkembangan Kognitif"},

		// ===== 4-5 Tahun =====a
		{RentangUsia: "4-5 Tahun", PertanyaaanCeklist: "Anak dapat berlari dengan perubahan arah dengan lancar", Aspek: "Perkembangan Motorik"},
		{RentangUsia: "4-5 Tahun", PertanyaaanCeklist: "Anak dapat berbicara dengan kalimat panjang dan jelas", Aspek: "Perkembangan Bahasa"},
		{RentangUsia: "4-5 Tahun", PertanyaaanCeklist: "Anak dapat menggambar bentuk (segitiga, persegi)", Aspek: "Perkembangan Motorik Halus"},
		{RentangUsia: "4-5 Tahun", PertanyaaanCeklist: "Anak dapat mengenakan pakaian sendiri", Aspek: "Perkembangan Motorik Halus"},
		{RentangUsia: "4-5 Tahun", PertanyaaanCeklist: "Anak dapat bermain permainan dengan aturan sederhana", Aspek: "Perkembangan Sosial Emosional"},
		{RentangUsia: "4-5 Tahun", PertanyaaanCeklist: "Anak dapat mengerti konsep angka dan warna", Aspek: "Perkembangan Kognitif"},
		{RentangUsia: "4-5 Tahun", PertanyaaanCeklist: "Anak dapat menunjukkan kemandirian dan kontrol diri", Aspek: "Perkembangan Sosial Emosional"},
		{RentangUsia: "4-5 Tahun", PertanyaaanCeklist: "Anak dapat menceritakan pengalaman dengan detail", Aspek: "Perkembangan Kognitif"},

		// ===== 5-6 Tahun =====a
		{RentangUsia: "5-6 Tahun", PertanyaaanCeklist: "Anak dapat melakukan aktivitas fisik kompleks (melompat tali, naik turun)", Aspek: "Perkembangan Motorik"},
		{RentangUsia: "5-6 Tahun", PertanyaaanCeklist: "Anak dapat berbicara dengan tata bahasa yang benar", Aspek: "Perkembangan Bahasa"},
		{RentangUsia: "5-6 Tahun", PertanyaaanCeklist: "Anak dapat menulis beberapa huruf atau nama", Aspek: "Perkembangan Motorik Halus"},
		{RentangUsia: "5-6 Tahun", PertanyaaanCeklist: "Anak dapat merawat diri dengan mandiri (mandi, makan, pakaian)", Aspek: "Perkembangan Motorik Halus"},
		{RentangUsia: "5-6 Tahun", PertanyaaanCeklist: "Anak dapat bermain dengan teman dan berbagi mainan", Aspek: "Perkembangan Sosial Emosional"},
		{RentangUsia: "5-6 Tahun", PertanyaaanCeklist: "Anak dapat membedakan antara baik dan buruk", Aspek: "Perkembangan Kognitif"},
		{RentangUsia: "5-6 Tahun", PertanyaaanCeklist: "Anak dapat menunjukkan tanggung jawab dan kepercayaan diri", Aspek: "Perkembangan Sosial Emosional"},
		{RentangUsia: "5-6 Tahun", PertanyaaanCeklist: "Anak dapat memecahkan masalah sederhana dengan pemikiran", Aspek: "Perkembangan Kognitif"},
	}

	// Check if data already exists
	var count int64
	if err := db.Model(&models.KategoriCapaian{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		log.Println("⏭️ Kategori capaian already seeded, skipping...")
		return nil
	}

	// Create data
	if err := db.CreateInBatches(data, 100).Error; err != nil {
		return err
	}

	log.Printf("✅ Seeded %d kategori capaian\n", len(data))
	return nil
}
