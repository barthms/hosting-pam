// seeders/kategori_capaian_seeder.go
package seeders

import (
	"log"
	"time"

	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type KategoriCapaianSeeder struct {
	db *gorm.DB
}

type kategoriCapaianRow struct {
	RentangUsia string
	// TipeLembarCapaian  string
	PertanyaaanCeklist string
	Aspek              string
}

func NewKategoriCapaianSeeder(db *gorm.DB) *KategoriCapaianSeeder {
	return &KategoriCapaianSeeder{db: db}
}

func (s *KategoriCapaianSeeder) Seed() error {
	log.Println("Starting kategori capaian seeding...")

	if err := s.seedKategoriCapaian(); err != nil {
		return err
	}

	log.Println("Kategori capaian seeding completed!")
	return nil
}

func (s *KategoriCapaianSeeder) seedKategoriCapaian() error {
	now := time.Now()

	// Data kategori capaian berdasarkan Buku KIA (Penanda Perkembangan Anak)
	rows := []kategoriCapaianRow{
		// ===== USIA 0-3 BULAN =====
		{"0-3", "Apakah bayi bisa mengangkat kepala mandiri hingga setinggi 45 derajat?", "motorik"},
		{"0-3", "Apakah bayi bisa menggerakkan kepala dari kiri/kanan ke tengah?", "motorik"},
		{"0-3", "Apakah bayi bisa melihat dan menatap wajah anda?", "sosial"},
		{"0-3", "Apakah bayi bisa mengoceh spontan atau bereaksi dengan mengoceh?", "bahasa"},
		{"0-3", "Apakah bayi suka tertawa keras?", "sosial"},
		{"0-3", "Apakah bayi bereaksi terkejut terhadap suara keras?", "bahasa"},
		{"0-3", "Apakah bayi membalas tersenyum ketika diajak bicara/tersenyum?", "sosial"},
		{"0-3", "Apakah bayi mengenal ibu dengan penglihatan, penciuman, pendengaran, kontak?", "sosial"},

		// ===== USIA 3-6 BULAN =====
		{"3-6", "Apakah bayi bisa berbalik dari telungkup ke telentang?", "motorik"},
		{"3-6", "Apakah bayi bisa mengangkat kepala secara mandiri hingga tegak 90 derajat?", "motorik"},
		{"3-6", "Apakah bayi bisa mempertahankan posisi kepala tetap tegak dan stabil?", "motorik"},
		{"3-6", "Apakah bayi bisa menggenggam mainan kecil atau mainan bertangkai?", "motorik"},
		{"3-6", "Apakah bayi bisa meraih benda yang ada dalam jangkauannya?", "motorik"},
		{"3-6", "Apakah bayi bisa mengamati tangannya sendiri?", "sosial"},
		{"3-6", "Apakah bayi berusaha memperluas pandangan?", "sosial"},
		{"3-6", "Apakah bayi mengarahkan matanya pada benda-benda kecil?", "sosial"},
		{"3-6", "Apakah bayi mengeluarkan suara gembira bernada tinggi atau memekik?", "bahasa"},
		{"3-6", "Apakah bayi tersenyum ketika melihat mainan/gambar yang menarik saat bermain sendiri?", "sosial"},

		// ===== USIA 6-9 BULAN =====
		{"6-9", "Apakah bayi bisa mengangkat badannya ke posisi berdiri?", "motorik"},
		{"6-9", "Apakah bayi belajar berdiri selama 30 detik atau berpegangan di kursi?", "motorik"},
		{"6-9", "Apakah bayi dapat berjalan dengan dituntun?", "motorik"},
		{"6-9", "Apakah bayi mengulurkan lengan/badan untuk meraih mainan yang diinginkan?", "motorik"},
		{"6-9", "Apakah bayi bisa menggenggam erat pensil?", "motorik"},
		{"6-9", "Apakah bayi memasukkan benda ke mulut?", "sosial"},
		{"6-9", "Apakah bayi mengulang menirukan bunyi yang didengar?", "bahasa"},
		{"6-9", "Apakah bayi menyebut 2-3 suku kata yang sama tanpa arti?", "bahasa"},
		{"6-9", "Apakah bayi mengeksplorasi sekitar, ingin tahu, ingin menyentuh apa saja?", "sosial"},
		{"6-9", "Apakah bayi bereaksi terhadap suara yang perlahan atau bisikan?", "bahasa"},
		{"6-9", "Apakah bayi senang diajak bermain Cilukba?", "sosial"},
		{"6-9", "Apakah bayi mengenal anggota keluarga, takut pada orang yang belum dikenal?", "sosial"},

		// ===== USIA 9-12 BULAN =====
		{"9-12", "Apakah bayi bisa duduk secara mandiri?", "motorik"},
		{"9-12", "Apakah bayi belajar berdiri, kedua kakinya menyangga sebagian berat badan?", "motorik"},
		{"9-12", "Apakah bayi bisa merangkak meraih mainan atau mendekati seseorang?", "motorik"},
		{"9-12", "Apakah bayi bisa memindahkan benda dari satu tangan ke tangan lainnya?", "motorik"},
		{"9-12", "Apakah bayi bisa memungut 2 benda, kedua tangan pegang 2 benda pada saat bersamaan?", "motorik"},
		{"9-12", "Apakah bayi bisa memungut benda sebesar kacang dengan cara meraup?", "motorik"},
		{"9-12", "Apakah bayi bersuara tanpa arti, mamama, bababa, dadada, tatatata?", "bahasa"},
		{"9-12", "Apakah bayi mencari mainan/benda yang dijatuhkan?", "sosial"},
		{"9-12", "Apakah bayi bermain tepuk tangan/Cilukba?", "sosial"},
		{"9-12", "Apakah bayi bergembira dengan melempar benda?", "sosial"},

		// ===== USIA 12-18 BULAN =====
		{"12-18", "Apakah anak bisa berdiri sendiri tanpa berpegangan 30 detik?", "motorik"},
		{"12-18", "Apakah anak bisa berjalan tanpa terhuyung-huyung?", "motorik"},
		{"12-18", "Apakah anak bisa menumpuk 4 buah kubus?", "motorik"},
		{"12-18", "Apakah anak bisa memungut benda kecil dengan ibu jari dan jari telunjuk?", "motorik"},
		{"12-18", "Apakah anak bisa menggelindingkan bola ke arah sasaran?", "motorik"},
		{"12-18", "Apakah anak bisa menyebut 3-6 kata yang mempunyai arti?", "bahasa"},
		{"12-18", "Apakah anak bisa membantu/menirukan pekerjaan rumah tangga?", "sosial"},
		{"12-18", "Apakah anak bisa memegang cangkir sendiri, belajar makan-minum sendiri?", "sosial"},

		// ===== USIA 18-24 BULAN =====
		{"18-24", "Apakah anak bisa berdiri sendiri tanpa berpegangan?", "motorik"},
		{"18-24", "Apakah anak bisa membungkuk memungut mainan kemudian berdiri kembali?", "motorik"},
		{"18-24", "Apakah anak bisa berjalan mundur lima langkah?", "motorik"},
		{"18-24", "Apakah anak bisa memanggil ayah dengan kata \"papa\", memanggil ibu dengan kata \"mama\"?", "bahasa"},
		{"18-24", "Apakah anak bisa menumpuk dua kubus?", "motorik"},
		{"18-24", "Apakah anak bisa memasukkan kubus di kotak?", "motorik"},
		{"18-24", "Apakah anak bisa menunjuk apa yang diinginkan tanpa menangis/merengek, anak bisa mengeluarkan suara yang menyenangkan atau menarik tangan ibu?", "sosial"},
		{"18-24", "Apakah anak bisa memperlihatkan rasa cemburu/bersaing?", "sosial"},

		// ===== USIA 2-3 TAHUN (24-36 BULAN) =====
		{"24-36", "Apakah anak bisa jalan naik tangga sendiri?", "motorik"},
		{"24-36", "Apakah anak bisa bermain dan menendang bola kecil?", "motorik"},
		{"24-36", "Apakah anak bisa mencoret-coret pensil pada kertas?", "motorik"},
		{"24-36", "Apakah anak bisa bicara dengan baik, menggunakan 2 kata?", "bahasa"},
		{"24-36", "Apakah anak bisa menunjuk 1 atau lebih bagian tubuhnya ketika diminta?", "sosial"},
		{"24-36", "Apakah anak bisa melihat gambar dan dapat menyebut dengan benar nama 2 benda atau lebih?", "sosial"},
		{"24-36", "Apakah anak bisa membantu memungut mainannya sendiri atau membantu mengangkat piring jika diminta?", "sosial"},
		{"24-36", "Apakah anak bisa makan nasi sendiri tanpa banyak tumpah?", "sosial"},
		{"24-36", "Apakah anak bisa melepas pakaiannya sendiri?", "sosial"},

		// ===== USIA 3-4 TAHUN (36-48 BULAN) =====
		{"36-48", "Apakah anak bisa berdiri 1 kaki 2 detik?", "motorik"},
		{"36-48", "Apakah anak bisa melompat kedua kaki diangkat?", "motorik"},
		{"36-48", "Apakah anak bisa mengayuh sepeda roda tiga?", "motorik"},
		{"36-48", "Apakah anak bisa menggambar garis lurus?", "motorik"},
		{"36-48", "Apakah anak bisa menumpuk 8 buah kubus?", "motorik"},
		{"36-48", "Apakah anak bisa mengenal 2-4 warna?", "sosial"},
		{"36-48", "Apakah anak bisa menyebut nama, umur, tempat?", "bahasa"},
		{"36-48", "Apakah anak bisa mengerti arti kata di atas, di bawah, di depan?", "bahasa"},
		{"36-48", "Apakah anak bisa mendengarkan cerita?", "bahasa"},
		{"36-48", "Apakah anak bisa mencuci dan mengeringkan tangan sendiri?", "sosial"},
		{"36-48", "Apakah anak bermain bersama teman, mengikuti aturan permainan?", "sosial"},
		{"36-48", "Apakah anak bisa mengenakan sepatu sendiri?", "sosial"},
		{"36-48", "Apakah anak bisa mengenakan celana panjang, kemeja, baju?", "sosial"},

		// ===== USIA 4-5 TAHUN (48-60 BULAN) =====
		{"48-60", "Apakah anak bisa berdiri 1 kaki 6 detik?", "motorik"},
		{"48-60", "Apakah anak bisa melompat-lompat 1 kaki?", "motorik"},
		{"48-60", "Apakah anak bisa menari?", "motorik"},
		{"48-60", "Apakah anak bisa menggambar tanda silang?", "motorik"},
		{"48-60", "Apakah anak bisa menggambar lingkaran?", "motorik"},
		{"48-60", "Apakah anak bisa menggambar orang dengan 3 bagian tubuh?", "motorik"},
		{"48-60", "Apakah anak bisa mengancing baju atau pakaian boneka?", "motorik"},
		{"48-60", "Apakah anak bisa menyebut nama lengkap tanpa dibantu?", "bahasa"},
		{"48-60", "Apakah anak senang menyebut kata-kata baru?", "bahasa"},
		{"48-60", "Apakah anak senang bertanya tentang sesuatu?", "bahasa"},
		{"48-60", "Apakah anak bisa menjawab pertanyaan dengan kata-kata yang benar?", "bahasa"},
		{"48-60", "Apakah anak bisa bicara yang mudah dimengerti?", "bahasa"},
		{"48-60", "Apakah anak bisa membandingkan/membedakan sesuatu dari ukuran dan bentuknya?", "sosial"},
		{"48-60", "Apakah anak bisa menyebut angka, menghitung jari?", "sosial"},

		// ===== USIA 5-6 TAHUN (60-72 BULAN) =====
		{"60-72", "Apakah anak bisa berjalan lurus?", "motorik"},
		{"60-72", "Apakah anak bisa berdiri dengan 1 kaki selama 11 detik?", "motorik"},
		{"60-72", "Apakah anak bisa menggambar dengan 6 bagian, menggambar orang lengkap?", "motorik"},
		{"60-72", "Apakah anak bisa menangkap bola kecil dengan kedua tangan?", "motorik"},
		{"60-72", "Apakah anak bisa menggambar segi empat?", "motorik"},
		{"60-72", "Apakah anak bisa mengerti arti lawan kata?", "bahasa"},
		{"60-72", "Apakah anak bisa mengerti pembicaraan yang menggunakan 7 kata atau lebih?", "bahasa"},
		{"60-72", "Apakah anak bisa menjawab pertanyaan tentang benda terbuat dari apa dan kegunaannya?", "bahasa"},
		{"60-72", "Apakah anak bisa mengenal angka, bisa menghitung angka 5-10?", "sosial"},
		{"60-72", "Apakah anak bisa mengenal warna-warni?", "sosial"},
		{"60-72", "Apakah anak bisa mengungkapkan simpati?", "sosial"},
		{"60-72", "Apakah anak bisa mengikuti aturan permainan?", "sosial"},
		{"60-72", "Apakah anak bisa berpakaian sendiri tanpa dibantu?", "sosial"},
	}

	for _, row := range rows {
		item := models.KategoriCapaian{
			RentangUsia: row.RentangUsia,
			// TipeLembarCapaian:  row.TipeLembarCapaian,
			PertanyaaanCeklist: row.PertanyaaanCeklist,
			Aspek:              row.Aspek,
			CreatedAt:          now,
			UpdatedAt:          now,
		}

		if err := s.db.Where("pertanyaan_ceklist = ?", item.PertanyaaanCeklist).FirstOrCreate(&item).Error; err != nil {
			log.Printf("Error seeding KategoriCapaian (Pertanyaan: %s): %v\n", item.PertanyaaanCeklist, err)
			return err
		}
	}

	log.Printf("Successfully seeded %d kategori capaian records\n", len(rows))
	return nil
}
