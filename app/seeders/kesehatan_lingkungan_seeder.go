package seeders

import (
	"monitoring-service/app/models"
	"gorm.io/gorm"
)

func KesehatanLingkunganSeeder(db *gorm.DB) error {
	data := []models.KategoriLingkungan{
		{
			Nama:      "Sarana Sanitasi",
			Deskripsi: "Pemantauan fasilitas sanitasi dasar keluarga",
			Indikator: []models.IndikatorLingkungan{
				{Pertanyaan: "Di mana ibu dan keluarga buang air besar?"},
				{Pertanyaan: "Bila jamban milik sendiri, bagian bawahnya/bak penampung tinja berupa apa?"},
				{Pertanyaan: "Bagaimana bentuk kloset jambannya?"},
			},
		},
		{
			Nama:      "Cuci Tangan Pakai Sabun",
			Deskripsi: "Kebiasaan dan sarana cuci tangan keluarga",
			Indikator: []models.IndikatorLingkungan{
				{Pertanyaan: "Seperti apa jenis sarana cuci tangan dirumah ibu?"},
				{Pertanyaan: "Apakah ibu melakukan cuci tangan pakai sabun?"},
				{Pertanyaan: "Apakah ibu mengetahui waktu-waktu kritis cuci tangan pakai sabun?"},
			},
		},
		{
			Nama:      "Pengelolaan Makanan dan Air Minum",
			Deskripsi: "Cara mengelola konsumsi air dan makanan yang sehat",
			Indikator: []models.IndikatorLingkungan{
				{Pertanyaan: "Apa sumber air minum di rumah ibu?"},
				{Pertanyaan: "Bagaimana ibu mengelola air minum di rumah tangga?"},
				{Pertanyaan: "Bagaimana ibu mengelola makanan di dalam keluarga?"},
			},
		},
		{
			Nama:      "Pengelolaan Sampah",
			Deskripsi: "Kebersihan lingkungan dari limbah padat",
			Indikator: []models.IndikatorLingkungan{
				{Pertanyaan: "Bagaimana ibu mengelola sampah di rumah?"},
			},
		},
		{
			Nama:      "Pengelolaan Limbah Cair",
			Deskripsi: "Pembuangan air bekas cuci baju, piring, dan mandi",
			Indikator: []models.IndikatorLingkungan{
				{Pertanyaan: "Bagaimana ibu mengelola limbah cair di rumah?"},
			},
		},
	}

	for _, k := range data {
		var count int64
		db.Model(&models.KategoriLingkungan{}).Where("nama = ?", k.Nama).Count(&count)
		if count == 0 {
			if err := db.Create(&k).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
