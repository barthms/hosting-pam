package routes

import (
	"monitoring-service/app/controllers"
	"monitoring-service/app/middlewares"

	"github.com/labstack/echo/v4"
)

func ConfigureRouter(e *echo.Echo, controller *controllers.Main) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})

	// Auth routes
	auth := e.Group("/auth")
	auth.POST("/register", controller.Register)
	auth.POST("/login", controller.Login)
	// auth.POST("/register/ortu", controller.RegisterOrangTua) // registrasi khusus orang tua
	secured := auth.Group("")
	secured.Use(middlewares.JWTAuth(controller.JWTSecret()))
	secured.GET("/me", controller.Me)

	// ==================== INFORMASI UMUM ====================
	informasiUmum := e.Group("/informasi-umum")
	informasiUmum.GET("", controller.InformasiUmum.GetAll)
	informasiUmum.GET("/:id", controller.InformasiUmum.Detail)
	// CRUD dibuka agar web dashboard bisa langsung mengelola Informasi Umum.
	informasiUmum.POST("", controller.InformasiUmum.Create)
	informasiUmum.PUT("/:id", controller.InformasiUmum.Update)
	informasiUmum.DELETE("/:id", controller.InformasiUmum.Delete)

	// ==================== EDUKASI IBU ====================
	e.GET("/edukasi-imd", controller.EdukasiIMD.GetAll)
	e.GET("/edukasi-kesehatan-mental", controller.EdukasiKesehatanMental.GetAll)
	e.GET("/edukasi-menyusui-asi", controller.EdukasiMenyusuiASI.GetAll)
	e.GET("/edukasi-nifas", controller.EdukasiNifas.GetAll)
	e.GET("/edukasi-tanda-melahirkan", controller.EdukasiTandaMelahirkan.GetAll)
	e.GET("/edukasi-trimester", controller.EdukasiTrimester.GetAll)
	// e.GET("/edukasi-trimester/:trimester", controller.EdukasiTrimester.GetByTrimester)
	// e.GET("/edukasi-trimester/:trimester/:kategori", controller.EdukasiTrimester.GetByKategori)
	// ==================== MODUL ADMIN  ====================

	admin := e.Group("/admin")
	admin.Use(middlewares.JWTAuth(controller.JWTSecret()))
	admin.Use(middlewares.AdminOnly())
	// NOTE: Admin hanya bisa membuat Kartu Keluarga + Anggota (Penduduk)
	// Tidak bisa membuat akun user lagi
	admin.POST("/kartu-keluarga", controller.AdminCreateKartuKeluarga)
	admin.GET("/kartu-keluarga", controller.AdminListKartuKeluarga)
	admin.GET("/kartu-keluarga/:kartu_keluarga_id", controller.AdminDetailKartuKeluarga)
	admin.PUT("/kartu-keluarga/:kartu_keluarga_id", controller.AdminUpdateKartuKeluarga)
	admin.PUT("/kartu-keluarga/:kartu_keluarga_id/anggota/:penduduk_id", controller.AdminUpdateAnggotaKeluarga)
	admin.POST("/kartu-keluarga/:kartu_keluarga_id/anggota", controller.AdminAddAnggotaKeluarga)
	admin.DELETE("/kartu-keluarga/:kartu_keluarga_id/anggota/:penduduk_id", controller.AdminDeleteAnggotaKeluarga)
	admin.DELETE("/kartu-keluarga/:kartu_keluarga_id", controller.AdminDeleteKartuKeluarga)

	// ==================== MODUL BIDAN ====================

	bidan := e.Group("/bidan")
	bidan.Use(middlewares.JWTAuth(controller.JWTSecret()))
	bidan.Use(middlewares.BidanOnly())

	// Posyandu Management (Bidan manage posyandu mereka)
	bidan.POST("/posyandu", controller.BidanCreatePosyandu)
	bidan.GET("/posyandu", controller.BidanListPosyandu)
	bidan.GET("/penduduk", controller.BidanListPenduduk)
	bidan.GET("/posyandu/:id", controller.BidanGetPosyanduDetail)
	bidan.PUT("/posyandu/:id", controller.BidanUpdatePosyandu)

	// Bidan Management (Bidan manage Bidan lain di posyandu mereka)
	bidan.POST("/bidan", controller.BidanCreateBidan)
	bidan.GET("/bidan", controller.BidanListBidan)
	bidan.GET("/bidan/:id", controller.BidanGetBidanDetail)
	bidan.PUT("/bidan/:id", controller.BidanUpdateBidan)
	bidan.PATCH("/bidan/:id/status", controller.BidanUpdateBidanStatus)

	// Jadwal layanan (Imunisasi) - dashboard bidan
	bidan.GET("/dashboard/jadwal-layanan", controller.JadwalLayanan.GetAll)
	bidan.POST("/dashboard/jadwal-layanan", controller.JadwalLayanan.Create)
	bidan.GET("/dashboard/jadwal-layanan/:id", controller.JadwalLayanan.GetByID)
	bidan.PUT("/dashboard/jadwal-layanan/:id", controller.JadwalLayanan.Update)
	bidan.DELETE("/dashboard/jadwal-layanan/:id", controller.JadwalLayanan.Delete)

	// Kader Management (Bidan manage Kader di posyandu mereka)
	bidan.POST("/kader", controller.BidanCreateKader)
	bidan.GET("/kader", controller.BidanListKader)
	bidan.GET("/kader/:id", controller.BidanGetKaderDetail)
	bidan.PUT("/kader/:id", controller.BidanUpdateKader)
	bidan.PATCH("/kader/:id/status", controller.BidanUpdateKaderStatus)

	// ==================== MODUL Anak ====================
	anak := e.Group("/anak")
	anak.Use(middlewares.JWTAuth(controller.JWTSecret()))
	_ = anak

	masterStandar := e.Group("/master-standar")
	masterStandar.Use(middlewares.JWTAuth(controller.JWTSecret()))
	_ = masterStandar

	// Kategori Capaian Routes
	// kategoriCapaian := e.Group("/kategori-capaian")
	// kategoriCapaian.Use(middlewares.JWTAuth(controller.JWTSecret()))
	kategoriCapaian := e.Group("/kategori-capaian")
	kategoriCapaian.Use(middlewares.JWTAuth(controller.JWTSecret()))
	// kategoriCapaian.GET("", controller.GetAllKategoriCapaian)
	// kategoriCapaian.GET("/:id", controller.GetKategoriCapaianById)
	// kategoriCapaian.GET("/rentang-usia", controller.GetKategoriCapaianByRentangUsia)
	// kategoriCapaian.POST("", controller.CreateKategoriCapaian)
	// kategoriCapaian.PUT("/:id", controller.UpdateKategoriCapaian)
	// kategoriCapaian.DELETE("/:id", controller.DeleteKategoriCapaian)

	// Perkembangan Routes
	// perkembangan := e.Group("/perkembangan")
	// perkembangan.Use(middlewares.JWTAuth(controller.JWTSecret()))
	perkembangan := e.Group("/perkembangan")
	perkembangan.Use(middlewares.JWTAuth(controller.JWTSecret()))
	// perkembangan.GET("", controller.GetAllPerkembangan)
	// perkembangan.GET("/:id", controller.GetPerkembanganById)
	// perkembangan.GET("/anak/:anak_id", controller.GetPerkembanganByAnakId)
	// perkembangan.GET("/anak/:anak_id/kategori/:kategori_capaian_id", controller.GetPerkembanganByAnakIdAndKategoriId)
	// perkembangan.POST("", controller.CreatePerkembangan)
	// perkembangan.PUT("/:id", controller.UpdatePerkembangan)
	// perkembangan.DELETE("/:id", controller.DeletePerkembangan)
	// perkembangan.GET("/search", controller.SearchPerkembangan)

	// Group untuk tenaga kesehatan (termasuk bidan, dokter, tenaga-kesehatan)
	tenaga := e.Group("/tenaga-kesehatan")
	tenaga.Use(middlewares.JWTAuth(controller.JWTSecret()))
	tenaga.Use(middlewares.TenagaKesehatan())

	// ==================== PERTUMBUHAN ANAK ====================
	tenaga.GET("/pertumbuhan/anak/:anak_id", controller.GetRiwayatPertumbuhan)
	tenaga.GET("/pertumbuhan/chart/:anak_id", controller.GetPertumbuhanChart)
	tenaga.POST("/pertumbuhan", controller.AddCatatanPertumbuhan)
	tenaga.PUT("/pertumbuhan/:id", controller.UpdateCatatanPertumbuhan)
	tenaga.DELETE("/pertumbuhan/:id", controller.DeleteCatatanPertumbuhan)

	// ==================== MODUL ANAK & PELAYANAN ANAK (yang sudah ada) ====================

	tenaga.GET("/anak", controller.Anak.AdminList)
	tenaga.POST("/anak", controller.Anak.Create)
	// tenaga.POST("/anak/dengan-penduduk", controller.Anak.CreateDenganPenduduk)
	tenaga.GET("/anak/:id", controller.Anak.Detail)
	tenaga.PUT("/anak/:id", controller.Anak.Update)
	tenaga.DELETE("/anak/:id", controller.Anak.Delete)

	tenaga.GET("/pelayanan-kesehatan-anak", controller.PelayananKesehatanAnak.GetByAnakID)
	tenaga.GET("/pelayanan-kesehatan-anak/:id", controller.PelayananKesehatanAnak.GetByID)
	tenaga.POST("/pelayanan-kesehatan-anak", controller.PelayananKesehatanAnak.CreatePelayananHandler)
	tenaga.PUT("/pelayanan-kesehatan-anak/:id", controller.PelayananKesehatanAnak.Update)
	tenaga.DELETE("/pelayanan-kesehatan-anak/:id", controller.PelayananKesehatanAnak.Delete)

	tenaga.GET("/Neonatus", controller.Neonatus.GetByAnakID)
	tenaga.GET("/Neonatus/:id", controller.Neonatus.GetByID)
	tenaga.POST("/Neonatus", controller.Neonatus.CreatePelayananHandler)
	tenaga.PUT("/Neonatus/:id", controller.Neonatus.Update)
	tenaga.DELETE("/Neonatus/:id", controller.Neonatus.Delete)

	tenaga.GET("/Pelayanan-Gizi-Anak", controller.PelayananGiziAnak.GetByAnakID)
	tenaga.GET("/Pelayanan-Gizi-Anak/:id", controller.PelayananGiziAnak.GetByID)
	tenaga.POST("/Pelayanan-Gizi-Anak", controller.PelayananGiziAnak.Create)
	tenaga.PUT("/Pelayanan-Gizi-Anak/:id", controller.PelayananGiziAnak.Update)
	tenaga.DELETE("/Pelayanan-Gizi-Anak/:id", controller.PelayananGiziAnak.Delete)
	tenaga.GET("/kategori-umur", controller.KategoriUmur.GetAll)

	tenaga.GET("/Pelayanan-Vitamin-ObatCacing", controller.KunjunganVitamin.GetByAnakID)
	tenaga.GET("/Pelayanan-Vitamin-ObatCacing/:id", controller.KunjunganVitamin.GetByID)
	tenaga.POST("/Pelayanan-Vitamin-ObatCacing", controller.KunjunganVitamin.Create)
	tenaga.PUT("/Pelayanan-Vitamin-ObatCacing/:id", controller.KunjunganVitamin.Update)
	tenaga.DELETE("/Pelayanan-Vitamin-ObatCacing/:id", controller.KunjunganVitamin.Delete)

	tenaga.GET("/Pelayanan-Imunisasi", controller.KunjunganImunisasi.GetByAnakID)
	tenaga.GET("/Pelayanan-Imunisasi/:id", controller.KunjunganImunisasi.GetByID)
	tenaga.POST("/Pelayanan-Imunisasi", controller.KunjunganImunisasi.Create)
	tenaga.PUT("/Pelayanan-Imunisasi/:id", controller.KunjunganImunisasi.Update)
	tenaga.DELETE("/Pelayanan-Imunisasi/:id", controller.KunjunganImunisasi.Delete)

	tenaga.GET("/Pemeriksaan-Gigi", controller.PemeriksaanGigi.GetByAnakID)
	tenaga.GET("/Pemeriksaan-Gigi/:id", controller.PemeriksaanGigi.GetByID)
	tenaga.POST("/Pemeriksaan-Gigi", controller.PemeriksaanGigi.Create)
	tenaga.PUT("/Pemeriksaan-Gigi/:id", controller.PemeriksaanGigi.Update)
	tenaga.DELETE("/Pemeriksaan-Gigi/:id", controller.PemeriksaanGigi.Delete)

	tenaga.GET("/Pemantauan-Pertumbuhan-Anak", controller.PemantauanPertumbuhan.GetByAnakID)
	tenaga.GET("/Pemantauan-Pertumbuhan-Anak/:id", controller.PemantauanPertumbuhan.GetByID)
	tenaga.POST("/Pemantauan-Pertumbuhan-Anak", controller.PemantauanPertumbuhan.Create)
	tenaga.PUT("/Pemantauan-Pertumbuhan-Anak/:id", controller.PemantauanPertumbuhan.Update)
	tenaga.DELETE("/Pemantauan-Pertumbuhan-Anak/:id", controller.PemantauanPertumbuhan.Delete)

	tenaga.GET("/Pengukuran-LilA", controller.PengukuranLilA.GetByAnakID)
	tenaga.GET("/Pengukuran-LilA/:id", controller.PengukuranLilA.GetByID)
	tenaga.POST("/Pengukuran-LilA", controller.PengukuranLilA.Create)
	tenaga.PUT("/Pengukuran-LilA/:id", controller.PengukuranLilA.Update)
	tenaga.DELETE("/Pengukuran-LilA/:id", controller.PengukuranLilA.Delete)

	tenaga.GET("/Catatan-Pelayanan", controller.CatatanPelayanan.GetByAnakID)
	tenaga.GET("/Catatan-Pelayanan/:id", controller.CatatanPelayanan.GetByID)
	tenaga.POST("/Catatan-Pelayanan", controller.CatatanPelayanan.Create)
	tenaga.PUT("/Catatan-Pelayanan/:id", controller.CatatanPelayanan.Update)
	tenaga.DELETE("/Catatan-Pelayanan/:id", controller.CatatanPelayanan.Delete)

	// ==================== KELUHAN ANAK ====================
	tenaga.GET("/keluhan-anak/:anak_id", controller.KeluhanAnak.GetByAnakID)
	tenaga.GET("/keluhan-anak/detail/:id", controller.KeluhanAnak.GetByID)
	tenaga.POST("/keluhan-anak", controller.KeluhanAnak.Create)
	tenaga.PUT("/keluhan-anak/:id", controller.KeluhanAnak.Update)
	tenaga.DELETE("/keluhan-anak/:id", controller.KeluhanAnak.Delete)

	// ==================== PEMANTAUAN ANAK ====================
	tenaga.GET("/pemantauan-anak/history", controller.PemantauanAnak.GetHistory)
	tenaga.GET("/pemantauan-anak/rentang-usia", controller.PemantauanAnak.GetRentangUsia)
	tenaga.GET("/pemantauan-anak/kategori/:rentang_id", controller.PemantauanAnak.GetKategori)
	tenaga.POST("/pemantauan-anak", controller.PemantauanAnak.Save)
	tenaga.DELETE("/pemantauan-anak/:id", controller.PemantauanAnak.Delete)
	tenaga.PUT("/pemantauan-anak/:id/verifikasi", controller.PemantauanAnak.Verifikasi)
	tenaga.POST("/pemantauan-anak/indikator", controller.PemantauanAnak.CreateKategori)
	tenaga.PUT("/pemantauan-anak/indikator/:id", controller.PemantauanAnak.UpdateKategori)
	tenaga.DELETE("/pemantauan-anak/indikator/:id", controller.PemantauanAnak.DeleteKategori)

	// ==================== KATEGORI CAPAIAN (Lembar Perawatan Anak) ====================
	// tenaga.GET("/kategori-capaian", controller.KategoriCapaian.GetAll)
	// tenaga.GET("/kategori-capaian/:id", controller.KategoriCapaian.GetByID)
	// tenaga.POST("/kategori-capaian", controller.KategoriCapaian.Create)
	// tenaga.PUT("/kategori-capaian/:id", controller.KategoriCapaian.Update)
	// tenaga.DELETE("/kategori-capaian/:id", controller.KategoriCapaian.Delete)

	// ==================== PERAWATAN ANAK (Lembar Capaian) ====================
	// tenaga.GET("/perawatan", controller.Perawatan.GetByAnakID)
	// tenaga.GET("/perawatan/:id", controller.Perawatan.GetByID)
	// tenaga.POST("/perawatan", controller.Perawatan.Create)
	// tenaga.PUT("/perawatan/:id", controller.Perawatan.Update)
	// tenaga.DELETE("/perawatan/:id", controller.Perawatan.Delete)

	// ==================== KESEHATAN LINGKUNGAN ====================
	// tenaga.GET("/kesehatan-lingkungan", controller.KesehatanLingkunganDanCatatanKader.GetAll)
	// tenaga.POST("/kesehatan-lingkungan", controller.KesehatanLingkunganDanCatatanKader.Create)
	// tenaga.GET("/kesehatan-lingkungan/:id", controller.KesehatanLingkunganDanCatatanKader.GetByID)
	// tenaga.PUT("/kesehatan-lingkungan/:id", controller.KesehatanLingkunganDanCatatanKader.Update)
	// tenaga.DELETE("/kesehatan-lingkungan/:id", controller.KesehatanLingkunganDanCatatanKader.Delete)
	// tenaga.GET("/kesehatan-lingkungan/:id/catatan-kader", controller.KesehatanLingkunganDanCatatanKader.GetCatatan)
	// tenaga.POST("/kesehatan-lingkungan/:id/catatan-kader", controller.KesehatanLingkunganDanCatatanKader.CreateCatatan)
	// tenaga.PUT("/kesehatan-lingkungan/:id/catatan-kader/:catatanId", controller.KesehatanLingkunganDanCatatanKader.UpdateCatatan)
	// tenaga.DELETE("/kesehatan-lingkungan/:id/catatan-kader/:catatanId", controller.KesehatanLingkunganDanCatatanKader.DeleteCatatan)
	// tenaga.PUT("/kesehatan-lingkungan/:id/catatan-kader/:catatanId/kirim-mobile", controller.KesehatanLingkunganDanCatatanKader.KirimCatatanKeMobile)

	// ==================== LINGKUNGAN (kategori & history) ====================
	lingkungan := e.Group("/lingkungan")
	lingkungan.Use(middlewares.JWTAuth(controller.JWTSecret()))
	lingkungan.GET("/kategori", controller.KesehatanLingkungan.GetAllKategori)
	lingkungan.GET("/history", controller.KesehatanLingkungan.GetHistory)
	lingkungan.GET("/detail/:id", controller.KesehatanLingkungan.GetDetail)
	lingkungan.POST("/submit", controller.KesehatanLingkungan.SubmitLembar)

	tenaga.POST("/lingkungan/kategori", controller.KesehatanLingkungan.CreateKategori)
	tenaga.DELETE("/lingkungan/kategori/:id", controller.KesehatanLingkungan.DeleteKategori)
	tenaga.POST("/lingkungan/kategori/:id/indikator", controller.KesehatanLingkungan.AddIndikator)
	tenaga.DELETE("/lingkungan/indikator/:id", controller.KesehatanLingkungan.DeleteIndikator)
	tenaga.DELETE("/lingkungan/:id", controller.KesehatanLingkungan.DeleteLembar)

	// ==================== EDUKASI DIGITAL ====================
	tenaga.GET("/edukasi-informasi-umum", controller.EdukasiInformasiUmum.GetAll)
	tenaga.POST("/edukasi-informasi-umum", controller.EdukasiInformasiUmum.Create)
	tenaga.GET("/edukasi-informasi-umum/:id", controller.EdukasiInformasiUmum.GetByID)
	tenaga.PUT("/edukasi-informasi-umum/:id", controller.EdukasiInformasiUmum.Update)
	tenaga.DELETE("/edukasi-informasi-umum/:id", controller.EdukasiInformasiUmum.Delete)

	tenaga.GET("/edukasi-trimester", controller.EdukasiTrimester.GetAll)
	tenaga.POST("/edukasi-trimester", controller.EdukasiTrimester.Create)
	tenaga.GET("/edukasi-trimester/:id", controller.EdukasiTrimester.GetByID)
	tenaga.PUT("/edukasi-trimester/:id", controller.EdukasiTrimester.Update)
	tenaga.DELETE("/edukasi-trimester/:id", controller.EdukasiTrimester.Delete)

	tenaga.GET("/edukasi-nifas", controller.EdukasiNifas.GetAll)
	tenaga.POST("/edukasi-nifas", controller.EdukasiNifas.Create)
	tenaga.GET("/edukasi-nifas/:id", controller.EdukasiNifas.GetByID)
	tenaga.PUT("/edukasi-nifas/:id", controller.EdukasiNifas.Update)
	tenaga.DELETE("/edukasi-nifas/:id", controller.EdukasiNifas.Delete)

	tenaga.GET("/edukasi-tanda-melahirkan", controller.EdukasiTandaMelahirkan.GetAll)
	tenaga.POST("/edukasi-tanda-melahirkan", controller.EdukasiTandaMelahirkan.Create)
	tenaga.GET("/edukasi-tanda-melahirkan/:id", controller.EdukasiTandaMelahirkan.GetByID)
	tenaga.PUT("/edukasi-tanda-melahirkan/:id", controller.EdukasiTandaMelahirkan.Update)
	tenaga.DELETE("/edukasi-tanda-melahirkan/:id", controller.EdukasiTandaMelahirkan.Delete)

	tenaga.GET("/edukasi-imd", controller.EdukasiImd.GetAll)
	tenaga.POST("/edukasi-imd", controller.EdukasiImd.Create)
	tenaga.GET("/edukasi-imd/:id", controller.EdukasiImd.GetByID)
	tenaga.PUT("/edukasi-imd/:id", controller.EdukasiImd.Update)
	tenaga.DELETE("/edukasi-imd/:id", controller.EdukasiImd.Delete)

	tenaga.GET("/edukasi-setelah-melahirkan", controller.EdukasiSetelahMelahirkan.GetAll)
	tenaga.POST("/edukasi-setelah-melahirkan", controller.EdukasiSetelahMelahirkan.Create)
	tenaga.GET("/edukasi-setelah-melahirkan/:id", controller.EdukasiSetelahMelahirkan.GetByID)
	tenaga.PUT("/edukasi-setelah-melahirkan/:id", controller.EdukasiSetelahMelahirkan.Update)
	tenaga.DELETE("/edukasi-setelah-melahirkan/:id", controller.EdukasiSetelahMelahirkan.Delete)

	tenaga.GET("/edukasi-menyusui-asi", controller.EdukasiMenyusuiAsi.GetAll)
	tenaga.POST("/edukasi-menyusui-asi", controller.EdukasiMenyusuiAsi.Create)
	tenaga.GET("/edukasi-menyusui-asi/:id", controller.EdukasiMenyusuiAsi.GetByID)
	tenaga.PUT("/edukasi-menyusui-asi/:id", controller.EdukasiMenyusuiAsi.Update)
	tenaga.DELETE("/edukasi-menyusui-asi/:id", controller.EdukasiMenyusuiAsi.Delete)

	tenaga.GET("/edukasi-pola-asuh", controller.EdukasiPolaAsuh.GetAll)
	tenaga.POST("/edukasi-pola-asuh", controller.EdukasiPolaAsuh.Create)
	tenaga.GET("/edukasi-pola-asuh/:id", controller.EdukasiPolaAsuh.GetByID)
	tenaga.PUT("/edukasi-pola-asuh/:id", controller.EdukasiPolaAsuh.Update)
	tenaga.DELETE("/edukasi-pola-asuh/:id", controller.EdukasiPolaAsuh.Delete)

	tenaga.GET("/edukasi-kesehatan-mental", controller.EdukasiKesehatanMental.GetAll)
	tenaga.POST("/edukasi-kesehatan-mental", controller.EdukasiKesehatanMental.Create)
	tenaga.GET("/edukasi-kesehatan-mental/:id", controller.EdukasiKesehatanMental.GetByID)
	tenaga.PUT("/edukasi-kesehatan-mental/:id", controller.EdukasiKesehatanMental.Update)
	tenaga.DELETE("/edukasi-kesehatan-mental/:id", controller.EdukasiKesehatanMental.Delete)

	tenaga.GET("/edukasi-perawatan-anak", controller.EdukasiPerawatanAnak.GetAll)
	tenaga.POST("/edukasi-perawatan-anak", controller.EdukasiPerawatanAnak.Create)
	tenaga.GET("/edukasi-perawatan-anak/:id", controller.EdukasiPerawatanAnak.GetByID)
	tenaga.PUT("/edukasi-perawatan-anak/:id", controller.EdukasiPerawatanAnak.Update)
	tenaga.DELETE("/edukasi-perawatan-anak/:id", controller.EdukasiPerawatanAnak.Delete)

	// tenaga.GET("/edukasi-mpasi", controller.EdukasiMPASI.GetAll)
	// tenaga.POST("/edukasi-mpasi", controller.EdukasiMPASI.Create)
	// tenaga.GET("/edukasi-mpasi/:id", controller.EdukasiMPASI.GetByID)
	// tenaga.PUT("/edukasi-mpasi/:id", controller.EdukasiMPASI.Update)
	// tenaga.DELETE("/edukasi-mpasi/:id", controller.EdukasiMPASI.Delete)

	// MPASI Sub-modules
	tenaga.GET("/edukasi-mpasi-aturan-porsi", controller.EdukasiAturanPorsiMPASI.GetAll)
	tenaga.POST("/edukasi-mpasi-aturan-porsi", controller.EdukasiAturanPorsiMPASI.Create)
	tenaga.GET("/edukasi-mpasi-aturan-porsi/:id", controller.EdukasiAturanPorsiMPASI.GetByID)
	tenaga.PUT("/edukasi-mpasi-aturan-porsi/:id", controller.EdukasiAturanPorsiMPASI.Update)
	tenaga.DELETE("/edukasi-mpasi-aturan-porsi/:id", controller.EdukasiAturanPorsiMPASI.Delete)

	tenaga.GET("/edukasi-mpasi-jadwal-harian", controller.EdukasiJadwalHarianMPASI.GetAll)
	tenaga.POST("/edukasi-mpasi-jadwal-harian", controller.EdukasiJadwalHarianMPASI.Create)
	tenaga.GET("/edukasi-mpasi-jadwal-harian/:id", controller.EdukasiJadwalHarianMPASI.GetByID)
	tenaga.PUT("/edukasi-mpasi-jadwal-harian/:id", controller.EdukasiJadwalHarianMPASI.Update)
	tenaga.DELETE("/edukasi-mpasi-jadwal-harian/:id", controller.EdukasiJadwalHarianMPASI.Delete)

	tenaga.GET("/edukasi-mpasi-resep", controller.EdukasiResepMPASI.GetAll)
	tenaga.POST("/edukasi-mpasi-resep", controller.EdukasiResepMPASI.Create)
	tenaga.GET("/edukasi-mpasi-resep/:id", controller.EdukasiResepMPASI.GetByID)
	tenaga.PUT("/edukasi-mpasi-resep/:id", controller.EdukasiResepMPASI.Update)
	tenaga.DELETE("/edukasi-mpasi-resep/:id", controller.EdukasiResepMPASI.Delete)

	// ==================== PEMANTAUAN INDIKATOR ====================
	tenaga.GET("/pemantauan-indikator", controller.PemantauanIndikator.GetAll)
	tenaga.POST("/pemantauan-indikator", controller.PemantauanIndikator.Create)
	tenaga.PUT("/pemantauan-indikator/:id", controller.PemantauanIndikator.Update)
	tenaga.DELETE("/pemantauan-indikator/:id", controller.PemantauanIndikator.Delete)

	// ==================== ADMIN AKUN BIDAN & KADER ====================
	// admin.POST("/bidan/:id/akun", controller.AdminCreateAkunBidan)  // belum diimplementasi
	// admin.POST("/kader/:id/akun", controller.AdminCreateAkunKader)  // belum diimplementasi
	// ==================== PERAWATAN ANAK (MILESTONE/PERKEMBANGAN) ====================
	tenaga.GET("/kategori-capaian", controller.GetAllKategoriCapaian)
	tenaga.GET("/kategori-capaian/rentang-usia/:rentang_usia", controller.GetKategoriCapaianByRentangUsia)
	tenaga.GET("/perawatan/:id", controller.GetPerawatanByID)
	tenaga.GET("/perawatan/anak/:anak_id", controller.GetPerawatanByAnakID)
	tenaga.GET("/perawatan/anak/:anak_id/rentang-usia/:rentang_usia", controller.GetPerawatanByAnakIDAndRentangUsia)
	tenaga.POST("/perawatan", controller.CreatePerawatan)
	tenaga.PUT("/perawatan/:id", controller.UpdatePerawatan)
	tenaga.DELETE("/perawatan/:id", controller.DeletePerawatan)

	// ==================== LEMBAR PEMANTAUAN ANAK ====================
	pemantauanLembar := e.Group("/tenaga-kesehatan")
	pemantauanLembar.Use(middlewares.JWTAuth(controller.JWTSecret()))
	pemantauanLembar.Use(middlewares.PemantauanLembarAccess())
	pemantauanLembar.GET("/lembar-pemantauan", controller.LembarPemantauan.GetByAnakID)
	pemantauanLembar.GET("/lembar-pemantauan/:id", controller.LembarPemantauan.GetByID)
	pemantauanLembar.POST("/lembar-pemantauan", controller.LembarPemantauan.Create)
	pemantauanLembar.PUT("/lembar-pemantauan/:id", controller.LembarPemantauan.Update)
	pemantauanLembar.DELETE("/lembar-pemantauan/:id", controller.LembarPemantauan.Delete)
	pemantauanLembar.PATCH("/lembar-pemantauan/:id/verifikasi", controller.LembarPemantauan.Verify)

	// ==================== MODUL IBU & KEHAMILAN ====================
	tenaga.POST("/ibu", controller.Ibu.Create)
	tenaga.GET("/ibu", controller.Ibu.GetAll)
	tenaga.GET("/ibuk", controller.Ibu.GetDashboard)
	tenaga.GET("/ibu/:id", controller.Ibu.GetByID)
	tenaga.PUT("/ibu/:id", controller.Ibu.Update)
	tenaga.DELETE("/ibu/:id", controller.Ibu.Delete)
	tenaga.GET("/ibu/by-penduduk/:pendudukId", controller.Ibu.GetByPendudukID)

	tenaga.POST("/kehamilan", controller.Kehamilan.Create)
	tenaga.GET("/kehamilan/all", controller.Kehamilan.GetAll)
	tenaga.GET("/kehamilan/:id", controller.Kehamilan.GetByID)
	tenaga.GET("/kehamilan", controller.Kehamilan.GetByIbuID)
	tenaga.PUT("/kehamilan/:id", controller.Kehamilan.Update)
	tenaga.DELETE("/kehamilan/:id", controller.Kehamilan.Delete)

	// ==================== PEMERIKSAAN KEHAMILAN (RUTIN) ====================
	tenaga.POST("/pemeriksaan-kehamilan", controller.PemeriksaanKehamilan.Create)
	tenaga.GET("/pemeriksaan-kehamilan/:id", controller.PemeriksaanKehamilan.GetByID)
	tenaga.GET("/pemeriksaan-kehamilan", controller.PemeriksaanKehamilan.GetByKehamilanID)
	tenaga.PUT("/pemeriksaan-kehamilan/:id", controller.PemeriksaanKehamilan.Update)
	tenaga.DELETE("/pemeriksaan-kehamilan/:id", controller.PemeriksaanKehamilan.Delete)
	tenaga.GET("/pemeriksaan-kehamilan/grafik-anc", controller.PemeriksaanKehamilan.GetGrafikANC)

	// ==================== EVALUASI KESEHATAN IBU ====================
	tenaga.POST("/evaluasi-kesehatan-ibu", controller.EvaluasiKesehatanIbu.Create)
	tenaga.GET("/evaluasi-kesehatan-ibu/:id", controller.EvaluasiKesehatanIbu.GetByID)
	tenaga.GET("/evaluasi-kesehatan-ibu", controller.EvaluasiKesehatanIbu.GetByKehamilanID)
	tenaga.PUT("/evaluasi-kesehatan-ibu/:id", controller.EvaluasiKesehatanIbu.Update)
	tenaga.DELETE("/evaluasi-kesehatan-ibu/:id", controller.EvaluasiKesehatanIbu.Delete)

	// ==================== RIWAYAT KEHAMILAN LALU (terhubung ke evaluasi) ====================
	tenaga.POST("/riwayat-kehamilan-lalu", controller.RiwayatKehamilanLalu.Create)
	tenaga.GET("/riwayat-kehamilan-lalu/:id", controller.RiwayatKehamilanLalu.GetByID)
	tenaga.GET("/riwayat-kehamilan-lalu", controller.RiwayatKehamilanLalu.GetByEvaluasiID)
	tenaga.PUT("/riwayat-kehamilan-lalu/:id", controller.RiwayatKehamilanLalu.Update)
	tenaga.DELETE("/riwayat-kehamilan-lalu/:id", controller.RiwayatKehamilanLalu.Delete)

	// ==================== PEMERIKSAAN DOKTER TRIMESTER 1 ====================
	// tenaga.POST("/pemeriksaan-dokter-t1", controller.PemeriksaanDokterTrimester1.Create)
	// tenaga.GET("/pemeriksaan-dokter-t1/:id", controller.PemeriksaanDokterTrimester1.GetByID)
	// tenaga.GET("/pemeriksaan-dokter-t1", controller.PemeriksaanDokterTrimester1.GetByKehamilanID)
	// tenaga.PUT("/pemeriksaan-dokter-t1/:id", controller.PemeriksaanDokterTrimester1.Update)
	// tenaga.DELETE("/pemeriksaan-dokter-t1/:id", controller.PemeriksaanDokterTrimester1.Delete)

	// ==================== PEMERIKSAAN LABORATORIUM & JIWA ====================
	tenaga.POST("/pemeriksaan-lab-jiwa", controller.PemeriksaanLaboratoriumJiwa.Create)
	tenaga.GET("/pemeriksaan-lab-jiwa/:id", controller.PemeriksaanLaboratoriumJiwa.GetByID)
	tenaga.GET("/pemeriksaan-lab-jiwa", controller.PemeriksaanLaboratoriumJiwa.GetByKehamilanID)
	tenaga.PUT("/pemeriksaan-lab-jiwa/:id", controller.PemeriksaanLaboratoriumJiwa.Update)
	tenaga.DELETE("/pemeriksaan-lab-jiwa/:id", controller.PemeriksaanLaboratoriumJiwa.Delete)

	// ==================== CATATAN PELAYANAN TRIMESTER 1 ====================
	tenaga.POST("/catatan-pelayanan-t1", controller.CatatanPelayananTrimester1.Create)
	tenaga.GET("/catatan-pelayanan-t1/:id", controller.CatatanPelayananTrimester1.GetByID)
	tenaga.GET("/catatan-pelayanan-t1", controller.CatatanPelayananTrimester1.GetByKehamilanID)
	tenaga.PUT("/catatan-pelayanan-t1/:id", controller.CatatanPelayananTrimester1.Update)
	tenaga.DELETE("/catatan-pelayanan-t1/:id", controller.CatatanPelayananTrimester1.Delete)

	// ==================== CATATAN PELAYANAN TRIMESTER 2 ====================
	tenaga.POST("/catatan-pelayanan-t2", controller.CatatanPelayananTrimester2.Create)
	tenaga.GET("/catatan-pelayanan-t2/:id", controller.CatatanPelayananTrimester2.GetByID)
	tenaga.GET("/catatan-pelayanan-t2", controller.CatatanPelayananTrimester2.GetByKehamilanID)
	tenaga.PUT("/catatan-pelayanan-t2/:id", controller.CatatanPelayananTrimester2.Update)
	tenaga.DELETE("/catatan-pelayanan-t2/:id", controller.CatatanPelayananTrimester2.Delete)

	// ==================== CATATAN PELAYANAN TRIMESTER 3 ====================
	tenaga.POST("/catatan-pelayanan-t3", controller.CatatanPelayananTrimester3.Create)
	tenaga.GET("/catatan-pelayanan-t3/:id", controller.CatatanPelayananTrimester3.GetByID)
	tenaga.GET("/catatan-pelayanan-t3", controller.CatatanPelayananTrimester3.GetByKehamilanID)
	tenaga.PUT("/catatan-pelayanan-t3/:id", controller.CatatanPelayananTrimester3.Update)
	tenaga.DELETE("/catatan-pelayanan-t3/:id", controller.CatatanPelayananTrimester3.Delete)

	// ==================== CATATAN PELAYANAN NIFAS ====================
	tenaga.POST("/catatan-pelayanan-nifas", controller.CatatanPelayananNifas.Create)
	tenaga.GET("/catatan-pelayanan-nifas/:id", controller.CatatanPelayananNifas.GetByID)
	tenaga.GET("/catatan-pelayanan-nifas", controller.CatatanPelayananNifas.GetByKehamilanID)
	tenaga.PUT("/catatan-pelayanan-nifas/:id", controller.CatatanPelayananNifas.Update)
	tenaga.DELETE("/catatan-pelayanan-nifas/:id", controller.CatatanPelayananNifas.Delete)

	// ==================== GRAFIK EVALUASI KEHAMILAN ====================
	// tenaga.POST("/grafik-evaluasi-kehamilan", controller.GrafikEvaluasiKehamilan.Create)
	// tenaga.GET("/grafik-evaluasi-kehamilan/grafik", controller.GrafikEvaluasiKehamilan.GetGrafik)
	// tenaga.GET("/grafik-evaluasi-kehamilan/:id", controller.GrafikEvaluasiKehamilan.GetByID)
	// tenaga.GET("/grafik-evaluasi-kehamilan", controller.GrafikEvaluasiKehamilan.GetByKehamilanID)
	// tenaga.PUT("/grafik-evaluasi-kehamilan/:id", controller.GrafikEvaluasiKehamilan.Update)
	// tenaga.DELETE("/grafik-evaluasi-kehamilan/:id", controller.GrafikEvaluasiKehamilan.Delete)

	// ==================== GRAFIK PENINGKATAN BERAT BADAN ====================
	tenaga.POST("/grafik-peningkatan-bb", controller.GrafikPeningkatanBB.Create)
	tenaga.GET("/grafik-peningkatan-bb/:id", controller.GrafikPeningkatanBB.GetByID)
	tenaga.GET("/grafik-peningkatan-bb", controller.GrafikPeningkatanBB.GetByKehamilanID)
	tenaga.PUT("/grafik-peningkatan-bb/:id", controller.GrafikPeningkatanBB.Update)
	tenaga.DELETE("/grafik-peningkatan-bb/:id", controller.GrafikPeningkatanBB.Delete)

	// ==================== PENJELASAN HASIL GRAFIK ====================
	tenaga.POST("/penjelasan-hasil-grafik", controller.PenjelasanHasilGrafik.Create)
	tenaga.GET("/penjelasan-hasil-grafik/:id", controller.PenjelasanHasilGrafik.GetByID)
	tenaga.GET("/penjelasan-hasil-grafik", controller.PenjelasanHasilGrafik.GetByKehamilanID)
	tenaga.PUT("/penjelasan-hasil-grafik/:id", controller.PenjelasanHasilGrafik.Update)
	tenaga.DELETE("/penjelasan-hasil-grafik/:id", controller.PenjelasanHasilGrafik.Delete)

	// ==================== RENCANA PERSALINAN ====================
	tenaga.POST("/rencana-persalinan", controller.RencanaPersalinan.Create)
	tenaga.GET("/rencana-persalinan/:id", controller.RencanaPersalinan.GetByID)
	tenaga.GET("/rencana-persalinan", controller.RencanaPersalinan.GetByKehamilanID)
	tenaga.PUT("/rencana-persalinan/:id", controller.RencanaPersalinan.Update)
	tenaga.DELETE("/rencana-persalinan/:id", controller.RencanaPersalinan.Delete)

	// ==================== RINGKASAN PELAYANAN PERSALINAN ====================
	tenaga.POST("/ringkasan-persalinan", controller.RingkasanPelayananPersalinan.Create)
	tenaga.GET("/ringkasan-persalinan/:id", controller.RingkasanPelayananPersalinan.GetByID)
	tenaga.GET("/ringkasan-persalinan", controller.RingkasanPelayananPersalinan.GetByKehamilanID)
	tenaga.PUT("/ringkasan-persalinan/:id", controller.RingkasanPelayananPersalinan.Update)
	tenaga.DELETE("/ringkasan-persalinan/:id", controller.RingkasanPelayananPersalinan.Delete)

	// ==================== RIWAYAT PROSES MELAHIRKAN ====================
	tenaga.POST("/riwayat-proses-melahirkan", controller.RiwayatProsesMelahirkan.Create)
	tenaga.GET("/riwayat-proses-melahirkan/:id", controller.RiwayatProsesMelahirkan.GetByID)
	tenaga.GET("/riwayat-proses-melahirkan", controller.RiwayatProsesMelahirkan.GetByKehamilanID)
	tenaga.PUT("/riwayat-proses-melahirkan/:id", controller.RiwayatProsesMelahirkan.Update)
	tenaga.DELETE("/riwayat-proses-melahirkan/:id", controller.RiwayatProsesMelahirkan.Delete)

	// ==================== KETERANGAN LAHIR ====================
	tenaga.POST("/keterangan-lahir", controller.KeteranganLahir.Create)
	tenaga.GET("/keterangan-lahir/:id", controller.KeteranganLahir.GetByID)
	tenaga.GET("/keterangan-lahir", controller.KeteranganLahir.GetByIbuID) // menggunakan ibu_id relasi
	tenaga.PUT("/keterangan-lahir/:id", controller.KeteranganLahir.Update)
	tenaga.DELETE("/keterangan-lahir/:id", controller.KeteranganLahir.Delete)

	// ==================== PELAYANAN IBU NIFAS ====================
	tenaga.POST("/pelayanan-ibu-nifas", controller.PelayananIbuNifas.Create)
	tenaga.GET("/pelayanan-ibu-nifas/:id", controller.PelayananIbuNifas.GetByID)
	tenaga.GET("/pelayanan-ibu-nifas", controller.PelayananIbuNifas.GetByKehamilanID)
	tenaga.PUT("/pelayanan-ibu-nifas/:id", controller.PelayananIbuNifas.Update)
	tenaga.DELETE("/pelayanan-ibu-nifas/:id", controller.PelayananIbuNifas.Delete)

	// ==================== SKRINING PRE-EKLAMPSIA ====================
	tenaga.POST("/skrining-preeklampsia", controller.SkriningPreeklampsia.Create)
	tenaga.GET("/skrining-preeklampsia/:id", controller.SkriningPreeklampsia.GetByID)
	tenaga.GET("/skrining-preeklampsia", controller.SkriningPreeklampsia.GetByKehamilanID)
	tenaga.PUT("/skrining-preeklampsia/:id", controller.SkriningPreeklampsia.Update)
	tenaga.DELETE("/skrining-preeklampsia/:id", controller.SkriningPreeklampsia.Delete)

	// ==================== SKRINING DIABETES MELITUS GESTASIONAL ====================
	tenaga.POST("/skrining-dm-gestasional", controller.SkriningDMGestasional.Create)
	tenaga.GET("/skrining-dm-gestasional/:id", controller.SkriningDMGestasional.GetByID)
	tenaga.GET("/skrining-dm-gestasional", controller.SkriningDMGestasional.GetByKehamilanID)
	tenaga.PUT("/skrining-dm-gestasional/:id", controller.SkriningDMGestasional.Update)
	tenaga.DELETE("/skrining-dm-gestasional/:id", controller.SkriningDMGestasional.Delete)

	// ==================== RUJUKAN ====================
	tenaga.POST("/rujukan", controller.Rujukan.Create)
	tenaga.GET("/rujukan/:id", controller.Rujukan.GetByID)
	tenaga.GET("/rujukan", controller.Rujukan.GetByKehamilanID)
	tenaga.PUT("/rujukan/:id", controller.Rujukan.Update)
	tenaga.DELETE("/rujukan/:id", controller.Rujukan.Delete)

	// ==================== PEMERIKSAAN DOKTER TRIMESTER 3 ====================
	// tenaga.POST("/pemeriksaan-dokter-t3", controller.PemeriksaanDokterTrimester3.Create)
	// tenaga.GET("/pemeriksaan-dokter-t3/:id", controller.PemeriksaanDokterTrimester3.GetByID)
	// tenaga.GET("/pemeriksaan-dokter-t3", controller.PemeriksaanDokterTrimester3.GetByKehamilanID)
	// tenaga.PUT("/pemeriksaan-dokter-t3/:id", controller.PemeriksaanDokterTrimester3.Update)
	// tenaga.DELETE("/pemeriksaan-dokter-t3/:id", controller.PemeriksaanDokterTrimester3.Delete)

	// ==================== PEMERIKSAAN LANJUTAN TRIMESTER 3 ====================
	// tenaga.POST("/pemeriksaan-lanjutan-t3", controller.PemeriksaanLanjutanTrimester3.Create)
	// tenaga.GET("/pemeriksaan-lanjutan-t3/:id", controller.PemeriksaanLanjutanTrimester3.GetByID)
	// tenaga.GET("/pemeriksaan-lanjutan-t3", controller.PemeriksaanLanjutanTrimester3.GetByKehamilanID)
	// tenaga.PUT("/pemeriksaan-lanjutan-t3/:id", controller.PemeriksaanLanjutanTrimester3.Update)
	// tenaga.DELETE("/pemeriksaan-lanjutan-t3/:id", controller.PemeriksaanLanjutanTrimester3.Delete)

	// ==================== KATEGORI TANDA BAHAYA ====================
	tenaga.GET("/kategori-tanda-bahaya", controller.KategoriTandaBahaya.GetAll)
	tenaga.POST("/kategori-tanda-bahaya", controller.KategoriTandaBahaya.Create)
	tenaga.GET("/kategori-tanda-bahaya/:id", controller.KategoriTandaBahaya.Detail)
	tenaga.GET("/kategori-tanda-bahaya/filter", controller.KategoriTandaBahaya.GetByTipeAndKategoriUsia)
	tenaga.PUT("/kategori-tanda-bahaya/:id", controller.KategoriTandaBahaya.Update)
	tenaga.DELETE("/kategori-tanda-bahaya/:id", controller.KategoriTandaBahaya.Delete)

	// ==================== SKRINING PEMANTAUAN TANDA BAHAYA ====================
	tenaga.GET("/skrining-pemantauan", controller.SkriningPemantauan.GetAll)
	tenaga.POST("/skrining-pemantauan", controller.SkriningPemantauan.Create)
	tenaga.GET("/skrining-pemantauan/:id", controller.SkriningPemantauan.Detail)
	tenaga.GET("/skrining-pemantauan/anak/:anak_id", controller.SkriningPemantauan.GetByAnakID)
	tenaga.PUT("/skrining-pemantauan/:id", controller.SkriningPemantauan.Update)
	tenaga.DELETE("/skrining-pemantauan/:id", controller.SkriningPemantauan.Delete)

	// ==================== KARTU KELUARGA ====================
	// tenaga.GET("/kartu-keluarga", controller.KartuKeluarga.GetAll)
	// tenaga.POST("/kartu-keluarga", controller.KartuKeluarga.Create)
	// tenaga.GET("/kartu-keluarga/:id", controller.KartuKeluarga.GetByID)
	// tenaga.GET("/kartu-keluarga/no-kk/:no_kk", controller.KartuKeluarga.GetByNoKartuKeluarga)
	// tenaga.PUT("/kartu-keluarga/:id", controller.KartuKeluarga.Update)
	// tenaga.DELETE("/kartu-keluarga/:id", controller.KartuKeluarga.Delete)

	// ==================== KEPENDUDUKAN ====================
	tenaga.GET("/kependudukan", controller.Kependudukan.GetAll)
	tenaga.POST("/kependudukan", controller.Kependudukan.Create)
	tenaga.GET("/kependudukan/:id", controller.Kependudukan.GetByID)
	// tenaga.GET("/kependudukan/kartu-keluarga/:kartu_keluarga_id", controller.Kependudukan.GetByKartuKeluargaID)
	tenaga.PUT("/kependudukan/:id", controller.Kependudukan.Update)
	tenaga.DELETE("/kependudukan/:id", controller.Kependudukan.Delete)
	tenaga.GET("/penduduk/rekap-dusun", controller.Kependudukan.GetRekapPerDusun)

	//jenis pelayanan neonatus
	tenaga.GET("/jenis-pelayanan", controller.JenisPelayanan.GetJenisPelayanan)
	// Di dalam group tenaga, tambahkan:
	tenaga.POST("/pemeriksaan-dokter-t1-complete", controller.PemeriksaanDokterCombined.CreateT1)
	tenaga.PUT("/pemeriksaan-dokter-t1-complete/:id", controller.PemeriksaanDokterCombined.UpdateT1)
	tenaga.GET("/pemeriksaan-dokter-t1-complete/:id", controller.PemeriksaanDokterCombined.GetT1ByID)
	tenaga.GET("/pemeriksaan-dokter-t1-complete", controller.PemeriksaanDokterCombined.GetT1ByKehamilan)
	tenaga.DELETE("/pemeriksaan-dokter-t1-complete/:id", controller.PemeriksaanDokterCombined.DeleteT1)

	// Sama untuk T3 complete...
	tenaga.POST("/pemeriksaan-dokter-t3-complete", controller.PemeriksaanDokterCombined.CreateT3)
	tenaga.PUT("/pemeriksaan-dokter-t3-complete/:id", controller.PemeriksaanDokterCombined.UpdateT3)
	tenaga.GET("/pemeriksaan-dokter-t3-complete/:id", controller.PemeriksaanDokterCombined.GetT3ByID)
	tenaga.GET("/pemeriksaan-dokter-t3-complete", controller.PemeriksaanDokterCombined.GetT3ByKehamilan)
	tenaga.DELETE("/pemeriksaan-dokter-t3-complete/:id", controller.PemeriksaanDokterCombined.DeleteT3)

	// untuk laporan ibu
	tenaga.GET("/laporan/ibu/preview", controller.LaporanIbu.Preview)
	tenaga.GET("/laporan/ibu/export/excel", controller.LaporanIbu.ExportExcel)
	//==== IBU ====
	ibu := e.Group("/ibu")
	ibu.Use(middlewares.JWTAuth(controller.JWTSecret()))
	ibu.Use(middlewares.IbuOnly())

	//untuk pencatatan kesehatan ANC
	ibu.GET("/dashboard", controller.Ibu.GetDashboard)
	ibu.GET("/pemeriksaan-kehamilan/:id", controller.PemeriksaanKehamilan.GetByID)
	ibu.GET("/pemeriksaan-kehamilan", controller.PemeriksaanKehamilan.GetByKehamilanID)
	ibu.GET("/pemeriksaan-kehamilan/grafik-anc", controller.PemeriksaanKehamilan.GetGrafikANC)
	//trimester
	ibu.GET("/catatan-pelayanan-t1", controller.CatatanPelayananTrimester1.GetByKehamilanID)
	ibu.GET("/catatan-pelayanan-t2", controller.CatatanPelayananTrimester2.GetByKehamilanID)
	ibu.GET("/catatan-pelayanan-t3", controller.CatatanPelayananTrimester3.GetByKehamilanID)
	ibu.GET("/pemeriksaan-dokter-t1-complete", controller.PemeriksaanDokterCombined.GetT1ByKehamilan)
	ibu.GET("/skrining-preeklampsia", controller.SkriningPreeklampsia.GetByKehamilanID)
	ibu.GET("/pemeriksaan-dokter-t3-complete", controller.PemeriksaanDokterCombined.GetT3ByKehamilan)
	ibu.GET("/evaluasi-kesehatan-ibu", controller.EvaluasiKesehatanIbu.GetByKehamilanID)
	ibu.GET("/rujukan", controller.Rujukan.GetByKehamilanID)
	ibu.GET("/pemeriksaan-kehamilan/grafik-anc", controller.PemeriksaanKehamilan.GetGrafikANC)
	ibu.GET("/ringkasan-persalinan", controller.RingkasanPelayananPersalinan.GetByKehamilanID)
	ibu.GET("/riwayat-proses-melahirkan", controller.RiwayatProsesMelahirkan.GetByKehamilanID)
	ibu.GET("/keterangan-lahir", controller.KeteranganLahir.GetByIbuID)
	ibu.GET("/pelayanan-ibu-nifas", controller.PelayananIbuNifas.GetByKehamilanID)
	ibu.GET("/catatan-pelayanan-nifas", controller.CatatanPelayananNifas.GetByKehamilanID)
	// tenaga.GET("/kependudukan", controller.Kependudukan.GetAll)
	// tenaga.POST("/kependudukan", controller.Kependudukan.Create)
	// tenaga.GET("/kependudukan/:id", controller.Kependudukan.GetByID)
	// // tenaga.GET("/kep// endudukan/kartu-keluarga/:kartu_keluarga_id", controller.Kependudukan.GetByKartuKeluargaID)
	// tenaga.PUT("/kependudukan/:id", controller.Kependudukan.Update)
	// tenaga.DELETE("/kependudukan/:id", controller.Kependudukan.Delete)

	//jenis pelayanan neonatus
	tenaga.GET("/jenis-pelayanan", controller.JenisPelayanan.GetJenisPelayanan)

	// ========================================= MODUL IBU - KIA- Kel-3 =========================================

	ibuk := e.Group("/modul-ibu")
	ibuk.Use(middlewares.JWTAuth(controller.JWTSecret()))
	ibuk.Use(middlewares.IbuOnly())

	// Data Ibu
	ibuk.GET("/kehamilan-aktif", controller.Kehamilan.GetAktifForOrangtua)

	// Evaluasi Kehamilan Ibu
	ibuk.GET("/evaluasi-kesehatan-ibu/me", controller.EvaluasiKesehatanIbu.GetMine)
	ibuk.GET("/evaluasi-kesehatan-ibu/:id", controller.EvaluasiKesehatanIbu.GetByIDForOrangtua)

	// Pemeriksaan Kehamilan
	ibuk.GET("/pemeriksaan-kehamilan/me", controller.PemeriksaanKehamilan.GetMine)
	ibuk.GET("/pemeriksaan-kehamilan/:id", controller.PemeriksaanKehamilan.GetByIDForOrangtua)

	// Skrining Preeklampsia
	ibuk.GET("/skrining-preeklampsia/me", controller.SkriningPreeklampsia.GetMine)
	ibuk.GET("/skrining-preeklampsia/:id", controller.SkriningPreeklampsia.GetByIDForOrangtua)

	ibuk.GET("/rujukan/:id", controller.Rujukan.GetByIDForOrangtua)

	// Pemeriksaan Dokter Trimester 1 & 3
	// ibuk.GET("/pemeriksaan-dokter-trimester-1/me", controller.PemeriksaanDokterTrimester1.GetMine)
	// ibuk.GET("/pemeriksaan-dokter-trimester-3/me", controller.PemeriksaanDokterTrimester3.GetMine)

	ibuk.GET("/pemeriksaan-dokter-trimester-1/me", controller.PemeriksaanDokterTrimester1.GetMine)
	ibuk.GET("/pemeriksaan-dokter-trimester-1/all", controller.PemeriksaanDokterTrimester1.GetAllMine) // BARU: list semua kunjungan
	ibuk.GET("/pemeriksaan-dokter-trimester-3/me", controller.PemeriksaanDokterTrimester3.GetMine)
	ibuk.GET("/pemeriksaan-dokter-trimester-3/all", controller.PemeriksaanDokterTrimester3.GetAllMine) // BARU: list semua kunjungan

	ibuk.GET("/riwayat-proses-melahirkan/me", controller.RiwayatProsesMelahirkan.GetMine)

	// Log TTD MMS
	ibuk.GET("/log-ttd-mms/me", controller.LogTTDMMS.GetMine)
	ibuk.POST("/log-ttd-mms", controller.LogTTDMMS.SaveMine)

	// Pemantauan Ibu Hamil
	ibuk.GET("/pemantauan-ibu-hamil/me", controller.PemantauanIbuHamil.GetMine)
	ibuk.POST("/pemantauan-ibu-hamil", controller.PemantauanIbuHamil.SaveMine)

	// Persiapan Melahirkan
	ibuk.GET("/persiapan-melahirkan/me", controller.PersiapanMelahirkan.GetMine)
	ibuk.POST("/persiapan-melahirkan", controller.PersiapanMelahirkan.SaveMine)

	// Proses Melahirkan
	ibuk.GET("/proses-melahirkan/me", controller.ProsesMelahirkan.GetMine)
	ibuk.POST("/proses-melahirkan", controller.ProsesMelahirkan.SaveMine)

	// Absensi Kelas Ibu Hamil
	ibuk.GET("/absensi-kelas-ibu-hamil/me", controller.AbsensiKelasIbuHamil.GetMine)
	ibuk.POST("/absensi-kelas-ibu-hamil", controller.AbsensiKelasIbuHamil.SaveMine)

	// Absensi Kelas Ibu Balita
	ibuk.GET("/absensi-kelas-ibu-balita/me", controller.AbsensiKelasIbuBalita.GetMine)
	ibuk.POST("/absensi-kelas-ibu-balita", controller.AbsensiKelasIbuBalita.SaveMine)

	// Checklist Pemantauan Ibu Nifas
	ibuk.GET("/checklist-pemantauan-ibu-nifas/me", controller.ChecklistPemantauanIbuNifas.GetMine)
	ibuk.POST("/checklist-pemantauan-ibu-nifas", controller.ChecklistPemantauanIbuNifas.SaveMine)
	ibuk.GET("/checklist-pemantauan-ibu-nifas/filled-days", controller.ChecklistPemantauanIbuNifas.GetFilledDays)

	// Catatan Pelayanan Ibu
	ibuk.GET("/catatan-pelayanan-t1", controller.CatatanPelayananTrimester1.GetByKehamilanID)
	ibuk.GET("/catatan-pelayanan-t2", controller.CatatanPelayananTrimester2.GetByKehamilanID)
	ibuk.GET("/catatan-pelayanan-t3", controller.CatatanPelayananTrimester3.GetByKehamilanID)
	ibuk.GET("/rujukan", controller.Rujukan.GetByKehamilanID)

	// Grafik Evaluasi Kehamilan
	ibuk.GET("/grafik-evaluasi-kehamilan/grafik", controller.GrafikEvaluasiKehamilan.GetGrafikForOrangtua)
	ibuk.GET("/grafik-evaluasi-kehamilan/me", controller.GrafikEvaluasiKehamilan.GetMine)
	// Baru
	ibuk.GET("/grafik-evaluasi-kehamilan/v2", controller.GrafikEvaluasiKehamilan.GetGrafikOnTheFlyForOrangtua)
	ibuk.GET("/grafik-evaluasi-kehamilan/:id", controller.GrafikEvaluasiKehamilan.GetByIDForOrangtua)
	ibuk.GET("/grafik-peningkatan-bb/v2", controller.GrafikPeningkatanBB.GetGrafikBBForOrangtua)

	// Keterangan Lahir
	ibuk.GET("/keterangan-lahir/me", controller.KeteranganLahir.GetMine)

	// Ringkasan Persalinan
	ibuk.GET("/ringkasan-persalinan/me", controller.RingkasanPelayananPersalinan.GetMine)
	// Pelayanan Ibu Nifas
	ibuk.GET("/pelayanan-ibu-nifas/me", controller.PelayananIbuNifas.GetMine)
	ibuk.GET("/catatan-pelayanan-nifas/me", controller.CatatanPelayananNifas.GetMine)

	// ibu := e.Group("/ibu")
	ibu.Use(middlewares.JWTAuth(controller.JWTSecret()))
	ibu.Use(middlewares.IbuOnly())
	ibu.GET("/anak", controller.Ibu.GetAnakSaya)

	ibu.GET("/lembar-pemantauan/rentang-usia", controller.LembarPemantauan.GetRentangUsiaForIbu)
	ibu.GET("/lembar-pemantauan/kategori-tanda-sakit", controller.LembarPemantauan.GetKategoriByRentangUsiaForIbu)
	ibu.POST("/lembar-pemantauan", controller.LembarPemantauan.CreateForIbu)
	ibu.GET("/lembar-pemantauan", controller.LembarPemantauan.GetByAnakIDForIbu)

	ibu.GET("/warna-tinja", controller.WarnaTinja.GetByAnakIDForIbu)
	ibu.POST("/warna-tinja", controller.WarnaTinja.SaveForIbu)

	ibu.GET("/pemeriksaan-gigi", controller.PemeriksaanGigi.GetByAnakIDForIbu)
	ibu.GET("/pengukuran-lila", controller.PengukuranLilA.GetByAnakIDForIbu)
	ibu.GET("/pengukuran-lila/:anak_id", controller.PengukuranLilA.GetByAnakIDForIbu)

	// ==================== PERAWATAN/MILESTONE (IBU) ====================
	ibu.GET("/kategori-capaian", controller.GetAllKategoriCapaian)
	ibu.GET("/kategori-capaian/rentang-usia/:rentang_usia", controller.GetKategoriCapaianByRentangUsia)
	ibu.GET("/perawatan/:id", controller.GetPerawatanByID)
	ibu.GET("/perawatan/anak/:anak_id", controller.GetPerawatanByAnakID)
	ibu.GET("/perawatan/anak/:anak_id/rentang-usia/:rentang_usia", controller.GetPerawatanByAnakIDAndRentangUsia)
	ibu.POST("/perawatan", controller.CreatePerawatan)
	ibu.PUT("/perawatan/:id", controller.UpdatePerawatan)
	ibu.DELETE("/perawatan/:id", controller.DeletePerawatan)
	// Catatan: Ibu tidak memiliki akses UPDATE/DELETE/VERIFY untuk menjaga integritas rekam medis

	// ==================== IMUNISASI ====================
	ibu.GET("/jadwal-imunisasi", controller.GetJadwalImunisasi)
	ibu.GET("/jadwal-imunisasi/anak/:anak_id", controller.GetJadwalImunisasiByAnakID)
	ibu.PUT("/jadwal-imunisasi/:id/tanggal-estimasi", controller.UpdateTanggalEstimasi)
	ibu.GET("/jadwal-imunisasi/:id", controller.GetJadwalByID)

	kader := e.Group("/kader")
	kader.Use(middlewares.JWTAuth(controller.JWTSecret()))
	kader.Use(middlewares.Kader())
	kader.GET("/kunjungan-imunisasi", controller.GetAllKunjunganImunisasi)
	kader.GET("/kunjungan-imunisasi/:id", controller.GetKunjunganImunisasiByID)
	kader.PUT("/kunjungan-imunisasi/:id/status", controller.UpdateStatusKunjungan)
	kader.PUT("/kunjungan-imunisasi/:id/tanggal-kunjungan", controller.UpdateTanggalKunjungan)
	kader.GET("/kunjungan-imunisasi/status/:status_id", controller.GetKunjunganImunisasiByStatus)
	kader.GET("/status-kunjungan/count", controller.GetJumlahKunjunganByStatus)

	// ==================== KELUHAN ANAK ====================
	ibu.GET("/keluhan-anak", controller.KeluhanAnak.GetByAnakIDForIbu)
	ibu.GET("/keluhan-anak/:id", controller.KeluhanAnak.GetByIDForIbu)
	tenaga.GET("/keluhan-anak/:anak_id", controller.KeluhanAnak.GetByAnakID)
	tenaga.GET("/keluhan-anak/detail/:id", controller.KeluhanAnak.GetByID)
	tenaga.POST("/keluhan-anak", controller.KeluhanAnak.Create)
	tenaga.PUT("/keluhan-anak/:id", controller.KeluhanAnak.Update)
	tenaga.DELETE("/keluhan-anak/:id", controller.KeluhanAnak.Delete)

	// ==================== EDUKASI MPASI (IBU) ====================
	ibu.GET("/edukasi-mpasi/materi/:bulan", controller.EdukasiMPASI.GetMateriByBulan)
	ibu.GET("/edukasi-mpasi/porsi-jadwal/:bulan", controller.EdukasiMPASI.GetPorsiJadwalByBulan)
	ibu.GET("/edukasi-mpasi/resep/:bulan", controller.EdukasiMPASI.GetResepByBulan)

	// ==================== EDUKASI MPASI (NAKES) ====================

	// Materi
	tenaga.GET("/edukasi-mpasi/materi", controller.EdukasiMPASI.GetMateriAll)
	tenaga.POST("/edukasi-mpasi/materi", controller.EdukasiMPASI.CreateMateri)
	tenaga.GET("/edukasi-mpasi/materi/:id", controller.EdukasiMPASI.GetMateriByID)
	tenaga.PUT("/edukasi-mpasi/materi/:id", controller.EdukasiMPASI.UpdateMateri)
	tenaga.DELETE("/edukasi-mpasi/materi/:id", controller.EdukasiMPASI.DeleteMateri)

	// Aturan Porsi
	tenaga.GET("/edukasi-mpasi/aturan-porsi", controller.EdukasiMPASI.GetAturanPorsiAll)
	tenaga.POST("/edukasi-mpasi/aturan-porsi", controller.EdukasiMPASI.CreateAturanPorsi)
	tenaga.GET("/edukasi-mpasi/aturan-porsi/:id", controller.EdukasiMPASI.GetAturanPorsiByID)
	tenaga.PUT("/edukasi-mpasi/aturan-porsi/:id", controller.EdukasiMPASI.UpdateAturanPorsi)
	tenaga.DELETE("/edukasi-mpasi/aturan-porsi/:id", controller.EdukasiMPASI.DeleteAturanPorsi)

	// Jadwal
	tenaga.GET("/edukasi-mpasi/jadwal", controller.EdukasiMPASI.GetJadwalAll)
	tenaga.POST("/edukasi-mpasi/jadwal", controller.EdukasiMPASI.CreateJadwal)
	tenaga.GET("/edukasi-mpasi/jadwal/:id", controller.EdukasiMPASI.GetJadwalByID)
	tenaga.PUT("/edukasi-mpasi/jadwal/:id", controller.EdukasiMPASI.UpdateJadwal)
	tenaga.DELETE("/edukasi-mpasi/jadwal/:id", controller.EdukasiMPASI.DeleteJadwal)

	// Resep
	tenaga.GET("/edukasi-mpasi/resep", controller.EdukasiMPASI.GetResepAll)
	tenaga.POST("/edukasi-mpasi/resep", controller.EdukasiMPASI.CreateResep)
	tenaga.GET("/edukasi-mpasi/resep/:id", controller.EdukasiMPASI.GetResepByID)
	tenaga.PUT("/edukasi-mpasi/resep/:id", controller.EdukasiMPASI.UpdateResep)
	tenaga.DELETE("/edukasi-mpasi/resep/:id", controller.EdukasiMPASI.DeleteResep)

}
