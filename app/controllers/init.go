package controllers

//AbsensiKelasIbuBalita//
import (
	"monitoring-service/app/usecases"
	"monitoring-service/pkg/config"
)

type Main struct {
	usecases *usecases.Main
	config   *config.Config

	// Controller yang sudah ada (untuk modul lain)
	KategoriTandaBahaya       *KategoriTandaBahayaController
	PemeriksaanDokterCombined *PemeriksaanDokterCombinedController
	// KartuKeluarga                      *KartuKeluargaController
	Anak                   *AnakController
	PelayananKesehatanAnak *PelayananKesehatanAnakController
	Neonatus               *NeonatusController
	PelayananGiziAnak      *kunjunganGiziController
	KunjunganVitamin       *KunjunganVitaminController
	KunjunganImunisasi     *KunjunganImunisasiController
	PemeriksaanGigi        *PemeriksaanGigiController
	PemantauanPertumbuhan  *PemantauanPertumbuhanController
	PengukuranLilA         *PengukuranLilAController
	CatatanPelayanan       *CatatanPelayananController
	InformasiUmum          *InformasiUmumController

	// Controller baru untuk struktur kehamilan
	Ibu                           *IbuController
	Kehamilan                     *KehamilanController
	PemeriksaanKehamilan          *PemeriksaanKehamilanController
	EvaluasiKesehatanIbu          *EvaluasiKesehatanIbuController
	PemeriksaanDokterTrimester1   *PemeriksaanDokterTrimester1Controller
	PemeriksaanDokterTrimester3   *PemeriksaanDokterTrimester3Controller
	PemeriksaanLaboratoriumJiwa   *PemeriksaanLaboratoriumJiwaController
	PemeriksaanLanjutanTrimester3 *PemeriksaanLanjutanTrimester3Controller
	CatatanPelayananTrimester1    *CatatanPelayananTrimester1Controller
	CatatanPelayananTrimester2    *CatatanPelayananTrimester2Controller
	CatatanPelayananTrimester3    *CatatanPelayananTrimester3Controller
	CatatanPelayananNifas         *CatatanPelayananNifasController
	GrafikPeningkatanBB           *GrafikPeningkatanBBController
	PenjelasanHasilGrafik         *PenjelasanHasilGrafikController
	RencanaPersalinan             *RencanaPersalinanController
	RingkasanPelayananPersalinan  *RingkasanPelayananPersalinanController
	RiwayatProsesMelahirkan       *RiwayatProsesMelahirkanController
	Rujukan                       *RujukanController
	SkriningDMGestasional         *SkriningDMGestasionalController
	SkriningPreeklampsia          *SkriningPreeklampsiaController
	SkriningPemantauan            *SkriningPemantauanController
	PelayananIbuNifas             *PelayananIbuNifasController
	RiwayatKehamilanLalu          *RiwayatKehamilanLaluController

	// Controller tambahan (sebelumnya hilang dari routes)
	KategoriUmur *KategoriUmurController
	Kader        *KaderController

	// Controller tambahan
	KesehatanLingkungan *KesehatanLingkunganController
	// KesehatanLingkunganDanCatatanKader *KesehatanLingkunganDanCatatanKaderController
	PemantauanAnak      *PemantauanAnakController
	PemantauanIndikator *PemantauanIndikatorController

	// Perawatan Anak (Lembar Capaian)
	// KategoriCapaian *KategoriCapaianController
	// Perawatan       *PerawatanController

	// Edukasi Digital
	// EdukasiTandaBahayaTrimester *EdukasiTandaBahayaTrimesterController
	LembarPemantauan *LembarPemantauanController
	KeteranganLahir  *KeteranganLahirController // <-- TAMBAHKAN INI
	Kependudukan     *KependudukanController
	JenisPelayanan   *JenisPelayananController

	// MODUL IBU
	PemantauanIbuHamil          *PemantauanIbuHamilController
	LogTTDMMS                   *LogTTDMMSController
	PersiapanMelahirkan         *PersiapanMelahirkanController
	ProsesMelahirkan            *ProsesMelahirkanController
	AbsensiKelasIbuHamil        *AbsensiKelasIbuHamilController
	AbsensiKelasIbuBalita       *AbsensiKelasIbuBalitaController
	ChecklistPemantauanIbuNifas *ChecklistPemantauanIbuNifasController
	WarnaTinja                  *WarnaTinjaController
	KeluhanAnak                 *KeluhanAnakController
	EdukasiIMD                  *EdukasiIMDController
	EdukasiMenyusuiASI          *EdukasiMenyusuiASIController
	GrafikEvaluasiKehamilan     *GrafikEvaluasiKehamilanController
	EdukasiInformasiUmum        *EdukasiInformasiUmumController
	EdukasiNifas                *EdukasiNifasController
	EdukasiTrimester            *EdukasiTrimesterController
	EdukasiTandaMelahirkan      *EdukasiTandaMelahirkanController
	EdukasiImd                  *EdukasiIMDController
	EdukasiSetelahMelahirkan    *EdukasiSetelahMelahirkanController
	EdukasiMenyusuiAsi          *EdukasiMenyusuiASIController
	EdukasiPolaAsuh             *EdukasiPolaAsuhController
	EdukasiKesehatanMental      *EdukasiKesehatanMentalController
	EdukasiPerawatanAnak        *EdukasiPerawatanAnakController
	EdukasiMPASI                *EdukasiMPASIController
	EdukasiAturanPorsiMPASI     *AturanPorsiMPASIController
	EdukasiJadwalHarianMPASI    *JadwalHarianMPASIController
	EdukasiResepMPASI           *ResepMPASIController
	JadwalLayanan               *JadwalLayananController
	// EdukasiTandaBahayaTrimester *EdukasiTandaBahayaTrimesterController
	LaporanIbu *LaporanIbuController
}

type Options struct {
	Config   *config.Config
	UseCases *usecases.Main
}

func Init(opts Options) *Main {
	m := &Main{
		usecases: opts.UseCases,
		config:   opts.Config,
	}

	// Controller yang sudah ada (tidak diubah)
	m.Anak = NewAnakController(opts.UseCases.Anak)
	m.PelayananKesehatanAnak = NewPelayananKesehatanAnakController(opts.UseCases.PelayananKesehatanAnak)
	m.Neonatus = NewPelayananNeonatusController(opts.UseCases.Neonatus)
	m.PelayananGiziAnak = NewKunjunganGiziController(opts.UseCases.KunjunganGizi)
	m.KunjunganVitamin = NewKunjunganVitaminController(opts.UseCases.KunjunganVitamin)
	m.KunjunganImunisasi = NewKunjunganImunisasiController(opts.UseCases.KunjunganImunisasi)
	m.PemeriksaanGigi = NewPemeriksaanGigiController(opts.UseCases.PemeriksaanGigi)
	m.PemantauanPertumbuhan = NewPemantauanPertumbuhanController(opts.UseCases.PemantauanPertumbuhan)
	m.PengukuranLilA = NewPengukuranLilAController(opts.UseCases.PengukuranLilA)
	m.CatatanPelayanan = NewCatatanPelayananController(opts.UseCases.CatatanPelayanan)
	m.KategoriTandaBahaya = NewKategoriTandaBahayaController(opts.UseCases.KategoriTandaBahaya)
	m.PemeriksaanDokterTrimester1 = NewPemeriksaanDokterTrimester1Controller(opts.UseCases.PemeriksaanDokterTrimester1)
	m.PemeriksaanDokterTrimester3 = NewPemeriksaanDokterTrimester3Controller(opts.UseCases.PemeriksaanDokterTrimester3)

	m.InformasiUmum = NewInformasiUmumController(opts.UseCases.InformasiUmum)
	m.EdukasiMPASI = NewEdukasiMPASIController(opts.UseCases.EdukasiMPASI)
	// Controller baru
	m.Ibu = NewIbuController(opts.UseCases.Ibu)
	// m.KartuKeluarga = NewKartuKeluargaController(opts.UseCases.KartuKeluarga)
	m.Kehamilan = NewKehamilanController(opts.UseCases.Kehamilan)
	m.PemeriksaanKehamilan = NewPemeriksaanKehamilanController(opts.UseCases.PemeriksaanKehamilan)
	m.EvaluasiKesehatanIbu = NewEvaluasiKesehatanIbuController(opts.UseCases.EvaluasiKesehatanIbu)
	m.PemeriksaanDokterTrimester1 = NewPemeriksaanDokterTrimester1Controller(opts.UseCases.PemeriksaanDokterTrimester1)
	m.PemeriksaanDokterTrimester3 = NewPemeriksaanDokterTrimester3Controller(opts.UseCases.PemeriksaanDokterTrimester3)
	m.PemeriksaanLaboratoriumJiwa = NewPemeriksaanLaboratoriumJiwaController(opts.UseCases.PemeriksaanLaboratoriumJiwa)
	m.PemeriksaanLanjutanTrimester3 = NewPemeriksaanLanjutanTrimester3Controller(opts.UseCases.PemeriksaanLanjutanTrimester3)
	m.CatatanPelayananTrimester1 = NewCatatanPelayananTrimester1Controller(opts.UseCases.CatatanPelayananTrimester1)
	m.CatatanPelayananTrimester2 = NewCatatanPelayananTrimester2Controller(opts.UseCases.CatatanPelayananTrimester2)
	m.CatatanPelayananTrimester3 = NewCatatanPelayananTrimester3Controller(opts.UseCases.CatatanPelayananTrimester3)
	m.CatatanPelayananNifas = NewCatatanPelayananNifasController(opts.UseCases.CatatanPelayananNifas)
	m.GrafikEvaluasiKehamilan = NewGrafikEvaluasiKehamilanController(opts.UseCases.GrafikEvaluasiKehamilan)
	m.GrafikPeningkatanBB = NewGrafikPeningkatanBBController(opts.UseCases.GrafikPeningkatanBB)
	m.PenjelasanHasilGrafik = NewPenjelasanHasilGrafikController(opts.UseCases.PenjelasanHasilGrafik)
	m.RencanaPersalinan = NewRencanaPersalinanController(opts.UseCases.RencanaPersalinan)
	m.RingkasanPelayananPersalinan = NewRingkasanPelayananPersalinanController(opts.UseCases.RingkasanPelayananPersalinan)
	m.RiwayatProsesMelahirkan = NewRiwayatProsesMelahirkanController(opts.UseCases.RiwayatProsesMelahirkan)
	m.Rujukan = NewRujukanController(opts.UseCases.Rujukan)
	m.SkriningDMGestasional = NewSkriningDMGestasionalController(opts.UseCases.SkriningDMGestasional)
	m.SkriningPreeklampsia = NewSkriningPreeklampsiaController(opts.UseCases.SkriningPreeklampsia)
	// m.SkriningPemantauan = NewSkriningPemantauanController(opts.UseCases.SkriningPemantauan)
	m.PelayananIbuNifas = NewPelayananIbuNifasController(opts.UseCases.PelayananIbuNifas)
	m.RiwayatKehamilanLalu = NewRiwayatKehamilanLaluController(opts.UseCases.RiwayatKehamilanLalu)
	m.KeteranganLahir = NewKeteranganLahirController(opts.UseCases.KeteranganLahir)
	m.Kependudukan = NewKependudukanController(opts.UseCases.Kependudukan)
	m.JenisPelayanan = NewJenisPelayananController(opts.UseCases.JenisPelayanan)
	m.KategoriUmur = NewKategoriUmurController(opts.UseCases.KategoriUmur)
	m.Kader = NewKaderController(opts.UseCases.Kader)
	m.PemeriksaanDokterCombined = NewPemeriksaanDokterCombinedController(
		opts.UseCases.PemeriksaanDokterTrimester1,
		opts.UseCases.PemeriksaanDokterTrimester3,
	)

	// Controller tambahan
	m.KeluhanAnak = NewKeluhanAnakController(opts.UseCases.KeluhanAnak)
	m.KesehatanLingkungan = NewKesehatanLingkunganController(opts.UseCases.KesehatanLingkungan)
	// m.KesehatanLingkunganDanCatatanKader = NewKesehatanLingkunganDanCatatanKaderController(
	// 	opts.UseCases.KesehatanLingkunganDanCatatanKader,
	// 	opts.UseCases.Ibu,
	// )
	m.PemantauanAnak = NewPemantauanAnakController(opts.UseCases.PemantauanAnak)
	m.PemantauanIndikator = NewPemantauanIndikatorController(opts.UseCases.PemantauanIndikator)

	// Perawatan Anak (Lembar Capaian)
	// m.KategoriCapaian = NewKategoriCapaianController(opts.UseCases.KategoriCapaian)
	// m.Perawatan = NewPerawatanController(opts.UseCases.Perawatan)

	// Edukasi Digital
	m.EdukasiInformasiUmum = NewEdukasiInformasiUmumController(opts.UseCases.EdukasiInformasiUmum)
	m.EdukasiNifas = NewEdukasiNifasController(opts.UseCases.EdukasiNifas)
	m.EdukasiTrimester = NewEdukasiTrimesterController(opts.UseCases.EdukasiTrimester)
	// m.EdukasiNifas = NewEdukasiNifasController(opts.UseCases.EdukasiNifas)
	// m.EdukasiTandaBahayaTrimester = NewEdukasiTandaBahayaTrimesterController(opts.UseCases.EdukasiTandaBahayaTrimester)
	m.EdukasiTandaMelahirkan = NewEdukasiTandaMelahirkanController(opts.UseCases.EdukasiTandaMelahirkan)
	m.EdukasiImd = NewEdukasiIMDController(opts.UseCases.EdukasiImd)
	m.EdukasiSetelahMelahirkan = NewEdukasiSetelahMelahirkanController(opts.UseCases.EdukasiSetelahMelahirkan)
	m.EdukasiMenyusuiAsi = NewEdukasiMenyusuiASIController(opts.UseCases.EdukasiMenyusuiAsi)
	m.EdukasiPolaAsuh = NewEdukasiPolaAsuhController(opts.UseCases.EdukasiPolaAsuh)
	m.EdukasiKesehatanMental = NewEdukasiKesehatanMentalController(opts.UseCases.EdukasiKesehatanMental)
	m.EdukasiPerawatanAnak = NewEdukasiPerawatanAnakController(opts.UseCases.EdukasiPerawatanAnak)
	// m.EdukasiMPASI = NewEdukasiMPASIController(opts.UseCases.EdukasiMPASI)
	m.LembarPemantauan = NewLembarPemantauanController(opts.UseCases.LembarPemantauan)
	m.KeteranganLahir = NewKeteranganLahirController(opts.UseCases.KeteranganLahir) // <-- TAMBAHKAN INI
	m.Kependudukan = NewKependudukanController(opts.UseCases.Kependudukan)
	m.JenisPelayanan = NewJenisPelayananController(opts.UseCases.JenisPelayanan)
	m.KeluhanAnak = NewKeluhanAnakController(opts.UseCases.KeluhanAnak)
	m.EdukasiAturanPorsiMPASI = NewAturanPorsiMPASIController(opts.UseCases.EdukasiAturanPorsiMPASI)
	m.EdukasiJadwalHarianMPASI = NewJadwalHarianMPASIController(opts.UseCases.EdukasiJadwalHarianMPASI)
	m.EdukasiResepMPASI = NewResepMPASIController(opts.UseCases.EdukasiResepMPASI)
	m.LaporanIbu = NewLaporanIbuController(opts.UseCases.LaporanIbu)

	// Jadwal layanan (imunisasi)
	m.JadwalLayanan = NewJadwalLayananController(opts.UseCases.JadwalLayanan)

	// MODEL IBU
	m.LogTTDMMS = NewLogTTDMMSController(opts.UseCases.LogTTDMMS)
	m.PemantauanIbuHamil = NewPemantauanIbuHamilController(opts.UseCases.PemantauanIbuHamil)
	m.PersiapanMelahirkan = NewPersiapanMelahirkanController(opts.UseCases.PersiapanMelahirkan)
	m.ProsesMelahirkan = NewProsesMelahirkanController(opts.UseCases.ProsesMelahirkan)
	m.AbsensiKelasIbuHamil = NewAbsensiKelasIbuHamilController(opts.UseCases.AbsensiKelasIbuHamil)
	m.AbsensiKelasIbuBalita = NewAbsensiKelasIbuBalitaController(opts.UseCases.AbsensiKelasIbuBalita)
	m.ChecklistPemantauanIbuNifas = NewChecklistPemantauanIbuNifasController(
		opts.UseCases.ChecklistPemantauanIbuNifas)
	m.WarnaTinja = NewWarnaTinjaController(opts.UseCases.WarnaTinja)

	m.EdukasiIMD = NewEdukasiIMDController(opts.UseCases.EdukasiIMD)
	m.EdukasiKesehatanMental = NewEdukasiKesehatanMentalController(opts.UseCases.EdukasiKesehatanMental)
	m.EdukasiMenyusuiASI = NewEdukasiMenyusuiASIController(opts.UseCases.EdukasiMenyusuiASI)
	m.EdukasiNifas = NewEdukasiNifasController(opts.UseCases.EdukasiNifas)
	m.EdukasiTandaMelahirkan = NewEdukasiTandaMelahirkanController(opts.UseCases.EdukasiTandaMelahirkan)
	m.EdukasiTrimester = NewEdukasiTrimesterController(opts.UseCases.EdukasiTrimester)
	m.GrafikEvaluasiKehamilan = NewGrafikEvaluasiKehamilanController(opts.UseCases.GrafikEvaluasiKehamilan)
	return m
}

func (m *Main) JWTSecret() string {
	return m.config.JWTSecret
}
