// package usecases

// import (
// 	"errors"
// 	"fmt"
// 	"monitoring-service/app/models"
// 	"monitoring-service/app/repositories"
// 	"time"
// )

// type GrafikEvaluasiKehamilanUsecase interface {
// 	Create(g *models.GrafikEvaluasiKehamilan) error
// 	GetByID(id int32) (*models.GrafikEvaluasiKehamilan, error)
// 	GetByKehamilanID(kehamilanID int32) ([]models.GrafikEvaluasiKehamilan, error)
// 	Update(g *models.GrafikEvaluasiKehamilan) error
// 	Delete(id int32) error
// 	GetGrafik(kehamilanID int32) (*GrafikResponse, error)

// 	// Modul Ibu
// 	GetMine(userID int32) ([]models.GrafikEvaluasiKehamilan, error)
// 	GetByIDForOrangtua(id int32, userID int32) (*models.GrafikEvaluasiKehamilan, error)
// 	GetGrafikForOrangtua(userID int32) (*GrafikResponse, error)

// 	GetGrafikOnTheFlyForOrangtua(userID int32) (*GrafikOnTheFlyResponse, error)
// }

// type GrafikTFUPoint struct {
// 	Usia   int     `json:"usia"`
// 	TFU    float64 `json:"tfu"`
// 	Normal float64 `json:"normal"`
// 	Upper  float64 `json:"upper"`
// 	Lower  float64 `json:"lower"`
// }

// type GrafikDJJPoint struct {
// 	Usia  int `json:"usia"`
// 	DJJ   int `json:"djj"`
// 	Upper int `json:"upper"`
// 	Lower int `json:"lower"`
// }

// type GrafikResponse struct {
// 	GrafikTFU []GrafikTFUPoint `json:"grafik_tfu"`
// 	GrafikDJJ []GrafikDJJPoint `json:"grafik_djj"`
// }

// type grafikEvaluasiKehamilanUsecase struct {
// 	repo *repositories.GrafikEvaluasiKehamilanRepository
// }

// // ─── Response untuk grafik on-the-fly ───────────────────────────────────────

// type GrafikOnTheFlyResponse struct {
// 	GrafikTFU             []GrafikTFUPointOnTheFly `json:"grafik_tfu"`
// 	GrafikDJJ             []GrafikDJJPointOnTheFly `json:"grafik_djj"`
// 	PenjelasanHasilGrafik *string                  `json:"penjelasan_hasil_grafik"`
// }

// // GrafikTFUPointOnTheFly - 1 titik pada grafik TFU, lengkap dengan data klinis
// type GrafikTFUPointOnTheFly struct {
// 	Usia              int      `json:"usia"`
// 	TFU               *float64 `json:"tfu"`
// 	Normal            float64  `json:"normal"`
// 	Upper             float64  `json:"upper"`
// 	Lower             float64  `json:"lower"`
// 	TanggalPeriksa    string   `json:"tanggal_periksa"`
// 	TekananDarah      string   `json:"tekanan_darah"`
// 	Hemoglobin        *float64 `json:"hemoglobin"`
// 	UrinProtein       string   `json:"urin_protein"`
// 	TabletTambahDarah *int     `json:"tablet_tambah_darah"`
// 	GerakanBayi       *string  `json:"gerakan_bayi"` // "normal" / "kurang"
// 	StatusTFU         string   `json:"status_tfu"`   // "normal" / "tinggi" / "rendah" / "tidak_ada_data"
// }

// // GrafikDJJPointOnTheFly - 1 titik pada grafik DJJ
// type GrafikDJJPointOnTheFly struct {
// 	Usia           int    `json:"usia"`
// 	DJJ            int    `json:"djj"`
// 	Upper          int    `json:"upper"`
// 	Lower          int    `json:"lower"`
// 	TanggalPeriksa string `json:"tanggal_periksa"`
// 	StatusDJJ      string `json:"status_djj"` // "normal" / "bradikardia" / "takikardia"
// }

// func NewGrafikEvaluasiKehamilanUsecase(repo *repositories.GrafikEvaluasiKehamilanRepository) GrafikEvaluasiKehamilanUsecase {
// 	return &grafikEvaluasiKehamilanUsecase{repo: repo}
// }

// func (u *grafikEvaluasiKehamilanUsecase) Create(g *models.GrafikEvaluasiKehamilan) error {
// 	if g.KehamilanID == 0 {
// 		return errors.New("kehamilan_id wajib diisi")
// 	}

// 	history, err := u.repo.FindByKehamilanID(g.KehamilanID)
// 	if err != nil {
// 		return err
// 	}

// 	penjelasan, risk := GeneratePenjelasanSingle(g, history)

// 	fmt.Println("=== DEBUG ===")
// 	fmt.Println("Penjelasan:", penjelasan)
// 	fmt.Println("Risk:", risk)
// 	fmt.Println("================")

// 	g.PenjelasanHasilGrafik = &penjelasan
// 	g.KategoriRisiko = &risk

// 	return u.repo.Create(g)
// }

// // func (u *grafikEvaluasiKehamilanUsecase) GetByID(id int32) (*models.GrafikEvaluasiKehamilan, error) {
// // 	return u.repo.FindByID(id)
// // }

// // func (u *grafikEvaluasiKehamilanUsecase) GetByKehamilanID(kehamilanID int32) ([]models.GrafikEvaluasiKehamilan, error) {
// // 	return u.repo.FindByKehamilanID(kehamilanID)
// // }

// // func (u *grafikEvaluasiKehamilanUsecase) Update(g *models.GrafikEvaluasiKehamilan) error {
// // 	_, err := u.repo.FindByID(g.ID)
// // 	if err != nil {
// // 		return errors.New("data grafik evaluasi kehamilan tidak ditemukan")
// // 	}
// // 	return u.repo.Update(g)
// // }

// func (u *grafikEvaluasiKehamilanUsecase) Delete(id int32) error {
// 	return u.repo.Delete(id)
// }

// // PERHITUNGAN
// func (u *grafikEvaluasiKehamilanUsecase) GetGrafik(kehamilanID int32) (*GrafikResponse, error) {
// 	tfuData, err := u.repo.FindGrafikTFU(kehamilanID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Sumber : https://www.alodokter.com/tinggi-fundus-uteri-dan-cara-menghitungnya
// 	// Sumber : https://www.haibunda.com/kehamilan/20231120160857-49-321521/tabel-tinggi-fundus-uteri-normal-sesuai-usia-kehamilan-2-cara-mengukurnya-tanpa-usg
// 	var tfuResult []GrafikTFUPoint
// 	for _, d := range tfuData {
// 		if d.UsiaGestasiMinggu != nil && d.TinggiFundusUteriCm != nil {
// 			usia := *d.UsiaGestasiMinggu
// 			tfu := *d.TinggiFundusUteriCm
// 			tfuResult = append(tfuResult, GrafikTFUPoint{
// 				Usia:   usia,
// 				TFU:    tfu,
// 				Normal: float64(usia),
// 				Upper:  float64(usia) + 2,
// 				Lower:  float64(usia) - 2,
// 			})
// 		}
// 	}

// 	djjData, err := u.repo.FindGrafikDJJ(kehamilanID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Sumber : https://pmc.ncbi.nlm.nih.gov/articles/PMC3678114/
// 	// Sumber : https://www.droracle.ai/articles/287588/what-is-the-normal-fetal-heart-rate-range-in
// 	var djjResult []GrafikDJJPoint
// 	for _, d := range djjData {
// 		if d.UsiaGestasiMinggu != nil && d.DenyutJantungBayiXMenit != nil {
// 			usia := *d.UsiaGestasiMinggu
// 			djj := *d.DenyutJantungBayiXMenit
// 			djjResult = append(djjResult, GrafikDJJPoint{
// 				Usia:  usia,
// 				DJJ:   djj,
// 				Upper: 160,
// 				Lower: 110,
// 			})
// 		}
// 	}

// 	return &GrafikResponse{
// 		GrafikTFU: tfuResult,
// 		GrafikDJJ: djjResult,
// 	}, nil
// }

// // ─── Modul Ibu ───────────────────────────────────────────────────────────────

// func (u *grafikEvaluasiKehamilanUsecase) GetMine(userID int32) ([]models.GrafikEvaluasiKehamilan, error) {
// 	return u.repo.FindMineByUserID(userID)
// }

// func (u *grafikEvaluasiKehamilanUsecase) GetByIDForOrangtua(id int32, userID int32) (*models.GrafikEvaluasiKehamilan, error) {
// 	allowed, err := u.repo.IsOwnedByUser(id, userID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if !allowed {
// 		return nil, errors.New("anda tidak memiliki akses ke data ini")
// 	}
// 	return u.repo.FindByID(id)
// }

// func (u *grafikEvaluasiKehamilanUsecase) GetGrafikForOrangtua(userID int32) (*GrafikResponse, error) {
// 	tfuData, err := u.repo.FindGrafikTFUByUserID(userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var tfuResult []GrafikTFUPoint
// 	for _, d := range tfuData {
// 		if d.UsiaGestasiMinggu != nil && d.TinggiFundusUteriCm != nil {
// 			usia := *d.UsiaGestasiMinggu
// 			tfu := *d.TinggiFundusUteriCm
// 			tfuResult = append(tfuResult, GrafikTFUPoint{
// 				Usia:   usia,
// 				TFU:    tfu,
// 				Normal: float64(usia),
// 				Upper:  float64(usia) + 2,
// 				Lower:  float64(usia) - 2,
// 			})
// 		}
// 	}

// 	djjData, err := u.repo.FindGrafikDJJByUserID(userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var djjResult []GrafikDJJPoint
// 	for _, d := range djjData {
// 		if d.UsiaGestasiMinggu != nil && d.DenyutJantungBayiXMenit != nil {
// 			usia := *d.UsiaGestasiMinggu
// 			djj := *d.DenyutJantungBayiXMenit
// 			djjResult = append(djjResult, GrafikDJJPoint{
// 				Usia:  usia,
// 				DJJ:   djj,
// 				Upper: 160,
// 				Lower: 110,
// 			})
// 		}
// 	}

// 	return &GrafikResponse{
// 		GrafikTFU: tfuResult,
// 		GrafikDJJ: djjResult,
// 	}, nil
// }

// // ─── GetGrafikOnTheFlyForOrangtua - endpoint utama modul ibu ─────────────────

// func (u *grafikEvaluasiKehamilanUsecase) GetGrafikOnTheFlyForOrangtua(userID int32) (*GrafikOnTheFlyResponse, error) {
// 	// 1. Cari kehamilan aktif
// 	kehamilan, err := u.repo.FindKehamilanAktifByUserID(userID)
// 	if err != nil {
// 		return nil, errors.New("kehamilan aktif tidak ditemukan")
// 	}

// 	kehamilanID := kehamilan.ID
// 	hpht := kehamilan.HPHT

// 	// 2. Ambil semua data sumber
// 	pemeriksaanList, err := u.repo.FindPemeriksaanForGrafik(kehamilanID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	djjList, err := u.repo.FindDJJForGrafik(kehamilanID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	gerakanBayiList, err := u.repo.FindGerakanBayiForGrafik(kehamilanID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	_, err = u.repo.FindUrinReduksiForGrafik(kehamilanID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	penjelasan, _ := u.repo.FindPenjelasanForGrafik(kehamilanID)

// 	// 3. Build map gerakan bayi per minggu
// 	gerakanBayiByMinggu := make(map[int]repositories.GerakanBayiGrafikRaw)
// 	for _, gb := range gerakanBayiList {
// 		gerakanBayiByMinggu[gb.MingguKehamilan] = gb
// 	}

// 	// 4. Build grafik TFU
// 	var grafikTFU []GrafikTFUPointOnTheFly
// 	for _, p := range pemeriksaanList {
// 		if p.TanggalPeriksa == nil {
// 			continue
// 		}

// 		minggu := hitungUsiaGestasi(*p.TanggalPeriksa, hpht)
// 		if minggu <= 0 {
// 			continue
// 		}

// 		// Hitung status TFU
// 		statusTFU := "tidak_ada_data"
// 		if p.TinggiRahim != nil {
// 			normal := float64(minggu)
// 			if *p.TinggiRahim < normal-2 {
// 				statusTFU = "rendah"
// 			} else if *p.TinggiRahim > normal+2 {
// 				statusTFU = "tinggi"
// 			} else {
// 				statusTFU = "normal"
// 			}
// 		}

// 		// Cek gerakan bayi di minggu ini
// 		var gerakanBayiStr *string
// 		if gb, ok := gerakanBayiByMinggu[minggu]; ok {
// 			var gbVal string
// 			if gb.GerakanBayiKurang {
// 				gbVal = "kurang"
// 			} else {
// 				gbVal = "normal"
// 			}
// 			gerakanBayiStr = &gbVal
// 		}

// 		point := GrafikTFUPointOnTheFly{
// 			Usia:              minggu,
// 			Normal:            float64(minggu),
// 			Upper:             float64(minggu) + 2,
// 			Lower:             float64(minggu) - 2,
// 			TanggalPeriksa:    p.TanggalPeriksa.Format("2006-01-02"),
// 			TFU:               p.TinggiRahim,
// 			TekananDarah:      p.TekananDarah,
// 			Hemoglobin:        p.TesLabHb,
// 			UrinProtein:       p.TesLabProteinUrine,
// 			TabletTambahDarah: p.TabletTambahDarah,
// 			GerakanBayi:       gerakanBayiStr,
// 			StatusTFU:         statusTFU,
// 		}

// 		grafikTFU = append(grafikTFU, point)
// 	}

// 	// 5. Build grafik DJJ
// 	var grafikDJJ []GrafikDJJPointOnTheFly
// 	for _, djj := range djjList {
// 		if djj.TanggalPeriksa == nil || djj.USGDJNilai == nil {
// 			continue
// 		}

// 		minggu := hitungUsiaGestasi(*djj.TanggalPeriksa, hpht)
// 		if minggu <= 0 {
// 			continue
// 		}

// 		// Hitung status DJJ
// 		statusDJJ := "normal"
// 		if *djj.USGDJNilai < 110 {
// 			statusDJJ = "bradikardia"
// 		} else if *djj.USGDJNilai > 160 {
// 			statusDJJ = "takikardia"
// 		}

// 		point := GrafikDJJPointOnTheFly{
// 			Usia:           minggu,
// 			DJJ:            *djj.USGDJNilai,
// 			Upper:          160,
// 			Lower:          110,
// 			TanggalPeriksa: djj.TanggalPeriksa.Format("2006-01-02"),
// 			StatusDJJ:      statusDJJ,
// 		}

// 		grafikDJJ = append(grafikDJJ, point)
// 	}

// 	// 6. Penjelasan grafik
// 	var penjelasanStr *string
// 	if penjelasan != nil {
// 		penjelasanStr = &penjelasan.CatatanPenjelasanGrafik
// 	}

// 	// 7. Pastikan tidak null
// 	if grafikTFU == nil {
// 		grafikTFU = []GrafikTFUPointOnTheFly{}
// 	}
// 	if grafikDJJ == nil {
// 		grafikDJJ = []GrafikDJJPointOnTheFly{}
// 	}

// 	return &GrafikOnTheFlyResponse{
// 		GrafikTFU:             grafikTFU,
// 		GrafikDJJ:             grafikDJJ,
// 		PenjelasanHasilGrafik: penjelasanStr,
// 	}, nil
// }

// // hitungUsiaGestasi - helper hitung minggu kehamilan dari HPHT
// func hitungUsiaGestasi(tanggal time.Time, hpht time.Time) int {
// 	selisihHari := tanggal.Sub(hpht).Hours() / 24
// 	minggu := int(selisihHari / 7)
// 	if minggu < 0 {
// 		return 0
// 	}
// 	return minggu
// }

package usecases

import (
	"errors"
	"fmt"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
	"time"
)

type GrafikEvaluasiKehamilanUsecase interface {
	Create(g *models.GrafikEvaluasiKehamilan) error
	GetByID(id int32) (*models.GrafikEvaluasiKehamilan, error)
	GetByKehamilanID(kehamilanID int32) ([]models.GrafikEvaluasiKehamilan, error)
	Update(g *models.GrafikEvaluasiKehamilan) error
	Delete(id int32) error
	GetGrafik(kehamilanID int32) (*GrafikResponse, error)

	// Modul Ibu
	GetMine(userID int32) ([]models.GrafikEvaluasiKehamilan, error)
	GetByIDForOrangtua(id int32, userID int32) (*models.GrafikEvaluasiKehamilan, error)
	GetGrafikForOrangtua(userID int32) (*GrafikResponse, error)

	GetGrafikOnTheFlyForOrangtua(userID int32) (*GrafikOnTheFlyResponse, error)
}

type GrafikTFUPoint struct {
	Usia   int     `json:"usia"`
	TFU    float64 `json:"tfu"`
	Normal float64 `json:"normal"`
	Upper  float64 `json:"upper"`
	Lower  float64 `json:"lower"`
}

type GrafikDJJPoint struct {
	Usia  int `json:"usia"`
	DJJ   int `json:"djj"`
	Upper int `json:"upper"`
	Lower int `json:"lower"`
}

type GrafikResponse struct {
	GrafikTFU []GrafikTFUPoint `json:"grafik_tfu"`
	GrafikDJJ []GrafikDJJPoint `json:"grafik_djj"`
}

type grafikEvaluasiKehamilanUsecase struct {
	repo *repositories.GrafikEvaluasiKehamilanRepository
	kehamilanRepo *repositories.KehamilanRepository
}

// ─── Response untuk grafik on-the-fly ───────────────────────────────────────

type GrafikOnTheFlyResponse struct {
	GrafikTFU             []GrafikTFUPointOnTheFly `json:"grafik_tfu"`
	GrafikDJJ             []GrafikDJJPointOnTheFly `json:"grafik_djj"`
	PenjelasanHasilGrafik *string                  `json:"penjelasan_hasil_grafik"`
}

// GrafikTFUPointOnTheFly - 1 titik pada grafik TFU, lengkap dengan data klinis
type GrafikTFUPointOnTheFly struct {
	Usia              int      `json:"usia"`
	TFU               *float64 `json:"tfu"`
	Normal            float64  `json:"normal"`
	Upper             float64  `json:"upper"`
	Lower             float64  `json:"lower"`
	TanggalPeriksa    string   `json:"tanggal_periksa"`
	TekananDarah      string   `json:"tekanan_darah"`
	Hemoglobin        *float64 `json:"hemoglobin"`
	UrinProtein       string   `json:"urin_protein"`
	TabletTambahDarah *int     `json:"tablet_tambah_darah"`
	GerakanBayi       *string  `json:"gerakan_bayi"` // "normal" / "kurang"
	StatusTFU         string   `json:"status_tfu"`   // "normal" / "tinggi" / "rendah" / "tidak_ada_data"
}

// GrafikDJJPointOnTheFly - 1 titik pada grafik DJJ
type GrafikDJJPointOnTheFly struct {
	Usia           int    `json:"usia"`
	DJJ            int    `json:"djj"`
	Upper          int    `json:"upper"`
	Lower          int    `json:"lower"`
	TanggalPeriksa string `json:"tanggal_periksa"`
	StatusDJJ      string `json:"status_djj"` // "normal" / "bradikardia" / "takikardia"
}

func NewGrafikEvaluasiKehamilanUsecase(repo *repositories.GrafikEvaluasiKehamilanRepository, kehamilanRepo *repositories.KehamilanRepository) GrafikEvaluasiKehamilanUsecase {
	return &grafikEvaluasiKehamilanUsecase{repo: repo, kehamilanRepo: kehamilanRepo}
}

func (u *grafikEvaluasiKehamilanUsecase) Create(g *models.GrafikEvaluasiKehamilan) error {
	if g.KehamilanID == 0 {
		return errors.New("kehamilan_id wajib diisi")
	}

	history, err := u.repo.FindByKehamilanID(g.KehamilanID)
	if err != nil {
		return err
	}

	penjelasan, risk := GeneratePenjelasanSingle(g, history)

	fmt.Println("=== DEBUG ===")
	fmt.Println("Penjelasan:", penjelasan)
	fmt.Println("Risk:", risk)
	fmt.Println("================")

	g.PenjelasanHasilGrafik = &penjelasan
	g.KategoriRisiko = &risk

	return u.repo.Create(g)
}

func (u *grafikEvaluasiKehamilanUsecase) GetByID(id int32) (*models.GrafikEvaluasiKehamilan, error) {
	return u.repo.FindByID(id)
}

func (u *grafikEvaluasiKehamilanUsecase) GetByKehamilanID(kehamilanID int32) ([]models.GrafikEvaluasiKehamilan, error) {
	return u.repo.FindByKehamilanID(kehamilanID)
}

func (u *grafikEvaluasiKehamilanUsecase) Update(g *models.GrafikEvaluasiKehamilan) error {
	_, err := u.repo.FindByID(g.ID)
	if err != nil {
		return errors.New("data grafik evaluasi kehamilan tidak ditemukan")
	}
	return u.repo.Update(g)
}

func (u *grafikEvaluasiKehamilanUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}

// PERHITUNGAN
func (u *grafikEvaluasiKehamilanUsecase) GetGrafik(kehamilanID int32) (*GrafikResponse, error) {
	tfuData, err := u.repo.FindGrafikTFU(kehamilanID)
	if err != nil {
		return nil, err
	}

	// Sumber : https://www.alodokter.com/tinggi-fundus-uteri-dan-cara-menghitungnya
	// Sumber : https://www.haibunda.com/kehamilan/20231120160857-49-321521/tabel-tinggi-fundus-uteri-normal-sesuai-usia-kehamilan-2-cara-mengukurnya-tanpa-usg
	var tfuResult []GrafikTFUPoint
	for _, d := range tfuData {
		if d.UsiaGestasiMinggu != nil && d.TinggiFundusUteriCm != nil {
			usia := *d.UsiaGestasiMinggu
			tfu := *d.TinggiFundusUteriCm
			tfuResult = append(tfuResult, GrafikTFUPoint{
				Usia:   usia,
				TFU:    tfu,
				Normal: float64(usia),
				Upper:  float64(usia) + 2,
				Lower:  float64(usia) - 2,
			})
		}
	}

	djjData, err := u.repo.FindGrafikDJJ(kehamilanID)
	if err != nil {
		return nil, err
	}

	// Sumber : https://pmc.ncbi.nlm.nih.gov/articles/PMC3678114/
	// Sumber : https://www.droracle.ai/articles/287588/what-is-the-normal-fetal-heart-rate-range-in
	var djjResult []GrafikDJJPoint
	for _, d := range djjData {
		if d.UsiaGestasiMinggu != nil && d.DenyutJantungBayiXMenit != nil {
			usia := *d.UsiaGestasiMinggu
			djj := *d.DenyutJantungBayiXMenit
			djjResult = append(djjResult, GrafikDJJPoint{
				Usia:  usia,
				DJJ:   djj,
				Upper: 160,
				Lower: 110,
			})
		}
	}

	return &GrafikResponse{
		GrafikTFU: tfuResult,
		GrafikDJJ: djjResult,
	}, nil
}

// ─── Modul Ibu ───────────────────────────────────────────────────────────────

func (u *grafikEvaluasiKehamilanUsecase) GetMine(userID int32) ([]models.GrafikEvaluasiKehamilan, error) {
	return u.repo.FindMineByUserID(userID)
}

func (u *grafikEvaluasiKehamilanUsecase) GetByIDForOrangtua(id int32, userID int32) (*models.GrafikEvaluasiKehamilan, error) {
	allowed, err := u.repo.IsOwnedByUser(id, userID)
	if err != nil {
		return nil, err
	}
	if !allowed {
		return nil, errors.New("anda tidak memiliki akses ke data ini")
	}
	return u.repo.FindByID(id)
}

func (u *grafikEvaluasiKehamilanUsecase) GetGrafikForOrangtua(userID int32) (*GrafikResponse, error) {
	tfuData, err := u.repo.FindGrafikTFUByUserID(userID)
	if err != nil {
		return nil, err
	}

	var tfuResult []GrafikTFUPoint
	for _, d := range tfuData {
		if d.UsiaGestasiMinggu != nil && d.TinggiFundusUteriCm != nil {
			usia := *d.UsiaGestasiMinggu
			tfu := *d.TinggiFundusUteriCm
			tfuResult = append(tfuResult, GrafikTFUPoint{
				Usia:   usia,
				TFU:    tfu,
				Normal: float64(usia),
				Upper:  float64(usia) + 2,
				Lower:  float64(usia) - 2,
			})
		}
	}

	djjData, err := u.repo.FindGrafikDJJByUserID(userID)
	if err != nil {
		return nil, err
	}

	var djjResult []GrafikDJJPoint
	for _, d := range djjData {
		if d.UsiaGestasiMinggu != nil && d.DenyutJantungBayiXMenit != nil {
			usia := *d.UsiaGestasiMinggu
			djj := *d.DenyutJantungBayiXMenit
			djjResult = append(djjResult, GrafikDJJPoint{
				Usia:  usia,
				DJJ:   djj,
				Upper: 160,
				Lower: 110,
			})
		}
	}

	return &GrafikResponse{
		GrafikTFU: tfuResult,
		GrafikDJJ: djjResult,
	}, nil
}

// ─── GetGrafikOnTheFlyForOrangtua - endpoint utama modul ibu ─────────────────

func (u *grafikEvaluasiKehamilanUsecase) GetGrafikOnTheFlyForOrangtua(userID int32) (*GrafikOnTheFlyResponse, error) {
	// 1. Cari kehamilan aktif
	kehamilan, err := u.repo.FindKehamilanAktifByUserID(userID)
	if err != nil {
		return nil, errors.New("kehamilan aktif tidak ditemukan")
	}

	kehamilanID := kehamilan.ID
	hpht := kehamilan.HPHT

	// 2. Ambil semua data sumber
	pemeriksaanList, err := u.repo.FindPemeriksaanForGrafik(kehamilanID)
	if err != nil {
		return nil, err
	}

	djjList, err := u.repo.FindDJJForGrafik(kehamilanID)
	if err != nil {
		return nil, err
	}

	gerakanBayiList, err := u.repo.FindGerakanBayiForGrafik(kehamilanID)
	if err != nil {
		return nil, err
	}

	_, err = u.repo.FindUrinReduksiForGrafik(kehamilanID)
	if err != nil {
		return nil, err
	}

	penjelasan, _ := u.repo.FindPenjelasanForGrafik(kehamilanID)

	// 3. Build map gerakan bayi per minggu
	gerakanBayiByMinggu := make(map[int]repositories.GerakanBayiGrafikRaw)
	for _, gb := range gerakanBayiList {
		gerakanBayiByMinggu[gb.MingguKehamilan] = gb
	}

	// 4. Build grafik TFU
	var grafikTFU []GrafikTFUPointOnTheFly
	for _, p := range pemeriksaanList {
		if p.TanggalPeriksa == nil {
			continue
		}

		minggu := hitungUsiaGestasi(*p.TanggalPeriksa, hpht)
		if minggu <= 0 {
			continue
		}

		// Hitung status TFU
		statusTFU := "tidak_ada_data"
		if p.TinggiRahim != nil {
			normal := float64(minggu)
			if *p.TinggiRahim < normal-2 {
				statusTFU = "rendah"
			} else if *p.TinggiRahim > normal+2 {
				statusTFU = "tinggi"
			} else {
				statusTFU = "normal"
			}
		}

		// Cek gerakan bayi di minggu ini
		var gerakanBayiStr *string
		if gb, ok := gerakanBayiByMinggu[minggu]; ok {
			var gbVal string
			if gb.GerakanBayiKurang {
				gbVal = "kurang"
			} else {
				gbVal = "normal"
			}
			gerakanBayiStr = &gbVal
		}

		point := GrafikTFUPointOnTheFly{
			Usia:              minggu,
			Normal:            float64(minggu),
			Upper:             float64(minggu) + 2,
			Lower:             float64(minggu) - 2,
			TanggalPeriksa:    p.TanggalPeriksa.Format("2006-01-02"),
			TFU:               p.TinggiRahim,
			TekananDarah:      p.TekananDarah,
			Hemoglobin:        p.TesLabHb,
			UrinProtein:       p.TesLabProteinUrine,
			TabletTambahDarah: p.TabletTambahDarah,
			GerakanBayi:       gerakanBayiStr,
			StatusTFU:         statusTFU,
		}

		grafikTFU = append(grafikTFU, point)
	}

	// 5. Build grafik DJJ
	var grafikDJJ []GrafikDJJPointOnTheFly
	for _, djj := range djjList {
		if djj.TanggalPeriksa == nil || djj.USGDJNilai == nil {
			continue
		}

		minggu := hitungUsiaGestasi(*djj.TanggalPeriksa, hpht)
		if minggu <= 0 {
			continue
		}

		// Hitung status DJJ
		statusDJJ := "normal"
		if *djj.USGDJNilai < 110 {
			statusDJJ = "bradikardia"
		} else if *djj.USGDJNilai > 160 {
			statusDJJ = "takikardia"
		}

		point := GrafikDJJPointOnTheFly{
			Usia:           minggu,
			DJJ:            *djj.USGDJNilai,
			Upper:          160,
			Lower:          110,
			TanggalPeriksa: djj.TanggalPeriksa.Format("2006-01-02"),
			StatusDJJ:      statusDJJ,
		}

		grafikDJJ = append(grafikDJJ, point)
	}

	// 6. Penjelasan grafik
	var penjelasanStr *string
	if penjelasan != nil {
		penjelasanStr = &penjelasan.CatatanPenjelasanGrafik
	}

	// 7. Pastikan tidak null
	if grafikTFU == nil {
		grafikTFU = []GrafikTFUPointOnTheFly{}
	}
	if grafikDJJ == nil {
		grafikDJJ = []GrafikDJJPointOnTheFly{}
	}

	return &GrafikOnTheFlyResponse{
		GrafikTFU:             grafikTFU,
		GrafikDJJ:             grafikDJJ,
		PenjelasanHasilGrafik: penjelasanStr,
	}, nil
}

// hitungUsiaGestasi - helper hitung minggu kehamilan dari HPHT
func hitungUsiaGestasi(tanggal time.Time, hpht time.Time) int {
	selisihHari := tanggal.Sub(hpht).Hours() / 24
	minggu := int(selisihHari / 7)
	if minggu < 0 {
		return 0
	}
	return minggu
}
