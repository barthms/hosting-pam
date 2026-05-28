package repositories

import (
	"errors"
	"monitoring-service/app/models"

	"time"

	"gorm.io/gorm"
)

type GrafikEvaluasiKehamilanRepository struct {
	db *gorm.DB
}

func NewGrafikEvaluasiKehamilanRepository(db *gorm.DB) *GrafikEvaluasiKehamilanRepository {
	return &GrafikEvaluasiKehamilanRepository{db: db}
}

// ====================== CREATE ======================
func (r *GrafikEvaluasiKehamilanRepository) Create(g *models.GrafikEvaluasiKehamilan) error {
	return r.db.Create(g).Error
}

// ====================== FIND BY ID ======================
func (r *GrafikEvaluasiKehamilanRepository) FindByID(id int32) (*models.GrafikEvaluasiKehamilan, error) {
	var g models.GrafikEvaluasiKehamilan
	err := r.db.
		Preload("Kehamilan.Ibu.Kependudukan").
		First(&g, id).Error

	if err != nil {
		return nil, err
	}
	return &g, nil
}

// ====================== FIND BY KEHAMILAN (ALL DATA) ======================
// Digunakan untuk Usecase yang butuh history lengkap (untuk GeneratePenjelasan)
func (r *GrafikEvaluasiKehamilanRepository) FindByKehamilanID(kehamilanID int32) ([]models.GrafikEvaluasiKehamilan, error) {
	var list []models.GrafikEvaluasiKehamilan
	err := r.db.
		Where("kehamilan_id = ?", kehamilanID).
		Order("usia_gestasi_minggu ASC, tanggal_bulan_tahun ASC").
		Find(&list).Error

	return list, err
}

// ====================== UPDATE ======================
func (r *GrafikEvaluasiKehamilanRepository) Update(g *models.GrafikEvaluasiKehamilan) error {
	result := r.db.Model(&models.GrafikEvaluasiKehamilan{}).
		Where("id = ?", g.ID).
		Updates(g)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("data grafik evaluasi kehamilan tidak ditemukan")
	}
	return nil
}

// ====================== DELETE ======================
func (r *GrafikEvaluasiKehamilanRepository) Delete(id int32) error {
	result := r.db.Delete(&models.GrafikEvaluasiKehamilan{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("data grafik evaluasi kehamilan tidak ditemukan")
	}
	return nil
}

// ====================== 1. TFU GRAPH DATA ======================
func (r *GrafikEvaluasiKehamilanRepository) FindGrafikTFU(kehamilanID int32) ([]models.GrafikEvaluasiKehamilan, error) {
	var list []models.GrafikEvaluasiKehamilan
	err := r.db.
		Select("usia_gestasi_minggu, tinggi_fundus_uteri_cm").
		Where("kehamilan_id = ? AND usia_gestasi_minggu IS NOT NULL AND tinggi_fundus_uteri_cm IS NOT NULL", kehamilanID).
		Order("usia_gestasi_minggu ASC").
		Find(&list).Error

	return list, err
}

// ====================== 2. DJJ GRAPH DATA ======================
func (r *GrafikEvaluasiKehamilanRepository) FindGrafikDJJ(kehamilanID int32) ([]models.GrafikEvaluasiKehamilan, error) {
	var list []models.GrafikEvaluasiKehamilan
	err := r.db.
		Select("usia_gestasi_minggu, denyut_jantung_bayi_x_menit").
		Where("kehamilan_id = ? AND usia_gestasi_minggu IS NOT NULL AND denyut_jantung_bayi_x_menit IS NOT NULL", kehamilanID).
		Order("usia_gestasi_minggu ASC").
		Find(&list).Error

	return list, err
}

// ====================== 3. TEKANAN DARAH GRAPH DATA ======================
// Menambahkan method baru untuk grafik ketiga
func (r *GrafikEvaluasiKehamilanRepository) FindGrafikTekananDarah(kehamilanID int32) ([]models.GrafikEvaluasiKehamilan, error) {
	var list []models.GrafikEvaluasiKehamilan
	err := r.db.
		Select("usia_gestasi_minggu, tekanan_darah_sistole, tekanan_darah_diastole").
		Where("kehamilan_id = ? AND usia_gestasi_minggu IS NOT NULL AND tekanan_darah_sistole IS NOT NULL", kehamilanID).
		Order("usia_gestasi_minggu ASC").
		Find(&list).Error

	return list, err
}

// ====================== LATEST STATUS ======================
// Tambahan: Mengambil penjelasan dan kategori risiko terbaru
func (r *GrafikEvaluasiKehamilanRepository) FindLatestStatus(kehamilanID int32) (*models.GrafikEvaluasiKehamilan, error) {
	var g models.GrafikEvaluasiKehamilan
	err := r.db.
		Select("penjelasan_hasil_grafik, kategori_risiko").
		Where("kehamilan_id = ?", kehamilanID).
		Order("usia_gestasi_minggu DESC, tanggal_bulan_tahun DESC").
		First(&g).Error

	if err != nil {
		return nil, err
	}
	return &g, nil
}

//
// ====================== MODUL IBU ======================
//

func (r *GrafikEvaluasiKehamilanRepository) IsOwnedByUser(grafikID int32, userID int32) (bool, error) {
	var count int64

	err := r.db.
		Model(&models.GrafikEvaluasiKehamilan{}).
		Joins("JOIN kehamilan k ON k.id = grafik_evaluasi_kehamilan.kehamilan_id").
		Joins("JOIN ibu i ON i.id = k.ibu_id").
		Joins("JOIN penduduk pd ON pd.id = i.penduduk_id").
		Joins("JOIN pengguna u ON u.penduduk_id = pd.id").
		Where("grafik_evaluasi_kehamilan.id = ? AND u.id = ?", grafikID, userID).
		Count(&count).Error

	return count > 0, err
}

func (r *GrafikEvaluasiKehamilanRepository) FindMineByUserID(userID int32) ([]models.GrafikEvaluasiKehamilan, error) {
	var list []models.GrafikEvaluasiKehamilan

	err := r.db.
		Model(&models.GrafikEvaluasiKehamilan{}).
		Joins("JOIN kehamilan k ON k.id = grafik_evaluasi_kehamilan.kehamilan_id").
		Joins("JOIN ibu i ON i.id = k.ibu_id").
		Joins("JOIN penduduk pd ON pd.id = i.penduduk_id").
		Joins("JOIN pengguna u ON u.id = pd.id").
		Where("u.id = ?", userID).
		Order("grafik_evaluasi_kehamilan.usia_gestasi_minggu ASC, grafik_evaluasi_kehamilan.tanggal_bulan_tahun ASC").
		Find(&list).Error

	return list, err
}

func (r *GrafikEvaluasiKehamilanRepository) FindGrafikTFUByUserID(userID int32) ([]models.GrafikEvaluasiKehamilan, error) {
	var list []models.GrafikEvaluasiKehamilan

	err := r.db.
		Model(&models.GrafikEvaluasiKehamilan{}).
		Select("grafik_evaluasi_kehamilan.usia_gestasi_minggu, grafik_evaluasi_kehamilan.tinggi_fundus_uteri_cm").
		Joins("JOIN kehamilan k ON k.id = grafik_evaluasi_kehamilan.kehamilan_id").
		Joins("JOIN ibu i ON i.id = k.ibu_id").
		Joins("JOIN penduduk pd ON pd.id = i.penduduk_id").
		Joins("JOIN pengguna u ON u.id = pd.id").
		Where("u.id = ?", userID).
		Where("grafik_evaluasi_kehamilan.usia_gestasi_minggu IS NOT NULL").
		Where("grafik_evaluasi_kehamilan.tinggi_fundus_uteri_cm IS NOT NULL").
		Order("grafik_evaluasi_kehamilan.usia_gestasi_minggu ASC").
		Find(&list).Error

	return list, err
}

func (r *GrafikEvaluasiKehamilanRepository) FindGrafikDJJByUserID(userID int32) ([]models.GrafikEvaluasiKehamilan, error) {
	var list []models.GrafikEvaluasiKehamilan

	err := r.db.
		Model(&models.GrafikEvaluasiKehamilan{}).
		Select("grafik_evaluasi_kehamilan.usia_gestasi_minggu, grafik_evaluasi_kehamilan.denyut_jantung_bayi_x_menit").
		Joins("JOIN kehamilan k ON k.id = grafik_evaluasi_kehamilan.kehamilan_id").
		Joins("JOIN ibu i ON i.id = k.ibu_id").
		Joins("JOIN penduduk pd ON pd.id = i.penduduk_id").
		Joins("JOIN pengguna u ON u.id = pd.id").
		Where("u.id = ?", userID).
		Where("grafik_evaluasi_kehamilan.usia_gestasi_minggu IS NOT NULL").
		Where("grafik_evaluasi_kehamilan.denyut_jantung_bayi_x_menit IS NOT NULL").
		Order("grafik_evaluasi_kehamilan.usia_gestasi_minggu ASC").
		Find(&list).Error

	return list, err
}

//
// ====================== MODUL IBU ======================
//

// Struct internal untuk data mentah dari masing-masing tabel
type PemeriksaanGrafikRaw struct {
	KehamilanID        int32      `gorm:"column:kehamilan_id"`
	TanggalPeriksa     *time.Time `gorm:"column:tanggal_periksa"`
	TinggiRahim        *float64   `gorm:"column:tinggi_rahim"`
	TekananDarah       string     `gorm:"column:tekanan_darah"`
	TesLabProteinUrine string     `gorm:"column:tes_lab_protein_urine"`
	TesLabHb           *float64   `gorm:"column:tes_lab_hb"`
	TabletTambahDarah  *int       `gorm:"column:tablet_tambah_darah"`
}

type DJJGrafikRaw struct {
	KehamilanID    int32      `gorm:"column:kehamilan_id"`
	TanggalPeriksa *time.Time `gorm:"column:tanggal_periksa"`
	USGDJNilai     *int       `gorm:"column:usg_djj_nilai"`
}

type GerakanBayiGrafikRaw struct {
	KehamilanID       int32 `gorm:"column:kehamilan_id"`
	MingguKehamilan   int   `gorm:"column:minggu_kehamilan"`
	GerakanBayiKurang bool  `gorm:"column:gerakan_bayi_kurang"`
}

type UrinReduksiGrafikRaw struct {
	KehamilanID         int32      `gorm:"column:kehamilan_id"`
	TanggalLab          *time.Time `gorm:"column:tanggal_lab"`
	LabUrinReduksiHasil string     `gorm:"column:lab_urin_reduksi_hasil"`
}

// FindKehamilanAktifByUserID - ambil kehamilan aktif untuk user
// func (r *GrafikEvaluasiKehamilanRepository) FindKehamilanAktifByUserID(userID int32) (*models.Kehamilan, error) {
//     var kehamilan models.Kehamilan

//     err := r.db.
//         Model(&models.Kehamilan{}).
//         Joins("JOIN ibu i ON i.id = kehamilan.ibu_id").
//         Joins("JOIN penduduk pd ON pd.id = i.penduduk_id").
//         Joins("JOIN pengguna u ON u.penduduk_id = pd.id").
//         Where("u.id = ?", userID).
//         Where("kehamilan.status_kehamilan LIKE ?", "%TRIMESTER%").
//         Order("kehamilan.created_at DESC").
//         First(&kehamilan).Error

//     if err != nil {
//         return nil, err
//     }

//     return &kehamilan, nil
// }

// FindKehamilanAktifByUserID - ambil kehamilan aktif untuk user
func (r *GrafikEvaluasiKehamilanRepository) FindKehamilanAktifByUserID(userID int32) (*models.Kehamilan, error) {
	var kehamilan models.Kehamilan

	err := r.db.
		Table("kehamilan").
		Select("kehamilan.*").
		Joins("JOIN ibu i ON i.id = kehamilan.ibu_id").
		Joins("JOIN penduduk pd ON pd.id = i.penduduk_id").
		Joins("JOIN pengguna u ON u.penduduk_id = pd.id").
		Where("u.id = ?", userID).
		Where("kehamilan.status_kehamilan LIKE ?", "%TRIMESTER%").
		Where("kehamilan.deleted_at IS NULL").
		Order("kehamilan.created_at DESC").
		First(&kehamilan).Error

	if err != nil {
		return nil, err
	}

	return &kehamilan, nil
}

// FindPemeriksaanForGrafik - ambil data pemeriksaan kehamilan untuk grafik
func (r *GrafikEvaluasiKehamilanRepository) FindPemeriksaanForGrafik(kehamilanID int32) ([]PemeriksaanGrafikRaw, error) {
	var result []PemeriksaanGrafikRaw

	err := r.db.
		Table("pemeriksaan_kehamilan").
		Select("kehamilan_id, tanggal_periksa, tinggi_rahim, tekanan_darah, tes_lab_protein_urine, tes_lab_hb, tablet_tambah_darah").
		Where("kehamilan_id = ?", kehamilanID).
		Where("tanggal_periksa IS NOT NULL").
		Order("tanggal_periksa ASC").
		Find(&result).Error

	return result, err
}

func (r *GrafikEvaluasiKehamilanRepository) FindDJJForGrafik(kehamilanID int32) ([]DJJGrafikRaw, error) {
	var result []DJJGrafikRaw

	err := r.db.
		Table("pemeriksaan_dokter_trimester_3").
		Select("kehamilan_id, tanggal_periksa, usg_djj_nilai").
		Where("kehamilan_id = ?", kehamilanID).
		Where("tanggal_periksa IS NOT NULL").
		Where("usg_djj_nilai IS NOT NULL").
		Order("tanggal_periksa ASC").
		Find(&result).Error

	return result, err
}

// FindGerakanBayiForGrafik - ambil gerakan bayi dari pemantauan ibu hamil
func (r *GrafikEvaluasiKehamilanRepository) FindGerakanBayiForGrafik(kehamilanID int32) ([]GerakanBayiGrafikRaw, error) {
	var result []GerakanBayiGrafikRaw

	err := r.db.
		Table("pemantauan_ibu_hamil").
		Select("kehamilan_id, minggu_kehamilan, gerakan_bayi_kurang").
		Where("kehamilan_id = ?", kehamilanID).
		Order("minggu_kehamilan ASC").
		Find(&result).Error

	return result, err
}

// FindUrinReduksiForGrafik - ambil urin reduksi dari pemeriksaan lanjutan trimester 3
func (r *GrafikEvaluasiKehamilanRepository) FindUrinReduksiForGrafik(kehamilanID int32) ([]UrinReduksiGrafikRaw, error) {
	var result []UrinReduksiGrafikRaw

	err := r.db.
		Table("pemeriksaan_lanjutan_trimester_3").
		Select("kehamilan_id, tanggal_lab, lab_urin_reduksi_hasil").
		Where("kehamilan_id = ?", kehamilanID).
		Where("tanggal_lab IS NOT NULL").
		Where("lab_urin_reduksi_hasil IS NOT NULL").
		Where("lab_urin_reduksi_hasil != ''").
		Order("tanggal_lab ASC").
		Find(&result).Error

	return result, err
}

// FindPenjelasanForGrafik - ambil penjelasan grafik (1 record per kehamilan)
func (r *GrafikEvaluasiKehamilanRepository) FindPenjelasanForGrafik(kehamilanID int32) (*models.PenjelasanHasilGrafik, error) {
	var result models.PenjelasanHasilGrafik

	err := r.db.
		Where("kehamilan_id = ?", kehamilanID).
		First(&result).Error

	if err != nil {
		// Jika tidak ditemukan, return nil tanpa error (bukan error fatal)
		return nil, nil
	}

	return &result, nil
}
