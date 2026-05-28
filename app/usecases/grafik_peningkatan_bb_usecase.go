// package usecases

// import (
// 	"errors"
// 	"monitoring-service/app/models"
// 	"monitoring-service/app/repositories"
// )

// type GrafikPeningkatanBBUsecase interface {
// 	Create(g *models.GrafikPeningkatanBB) error
// 	GetByID(id int32) (*models.GrafikPeningkatanBB, error)
// 	GetByKehamilanID(kehamilanID int32) ([]models.GrafikPeningkatanBB, error)
// 	Update(g *models.GrafikPeningkatanBB) error
// 	Delete(id int32) error
// }

// type grafikPeningkatanBBUsecase struct {
// 	repo          *repositories.GrafikPeningkatanBBRepository
// 	kehamilanRepo *repositories.KehamilanRepository
// }

// func NewGrafikPeningkatanBBUsecase(
// 	repo *repositories.GrafikPeningkatanBBRepository,
// 	kehamilanRepo *repositories.KehamilanRepository,
// ) GrafikPeningkatanBBUsecase {
// 	return &grafikPeningkatanBBUsecase{
// 		repo:          repo,
// 		kehamilanRepo: kehamilanRepo,
// 	}
// }

// //
// // 🔥 CORE LOGIC (dipakai Create & Update)
// //
// func (u *grafikPeningkatanBBUsecase) processGrafik(g *models.GrafikPeningkatanBB) error {

// 	if g.KehamilanID == 0 {
// 		return errors.New("kehamilan_id wajib diisi")
// 	}
// 	if g.BeratBadan == nil || g.MingguKehamilan == nil {
// 		return errors.New("berat_badan dan minggu_kehamilan wajib diisi")
// 	}

// 	// 🔹 Ambil data kehamilan
// 	kehamilan, err := u.kehamilanRepo.FindByID(g.KehamilanID)
// 	if err != nil {
// 		return errors.New("data kehamilan tidak ditemukan")
// 	}

// 	bbAwal := kehamilan.BB_Awal
// 	imt := kehamilan.IMT_Awal

// 	// 🔹 Tentukan range berdasarkan IMT
// 	var minTotal, maxTotal float64

// 	switch {
// 	case imt < 18.5:
// 		minTotal, maxTotal = 12.5, 18
// 	case imt < 25:
// 		minTotal, maxTotal = 11.5, 16
// 	case imt < 30:
// 		minTotal, maxTotal = 7, 11.5
// 	default:
// 		minTotal, maxTotal = 5, 9
// 	}

// 	// 🔹 Hitung batas minggu
// 	minggu := float64(*g.MingguKehamilan)

// 	minMinggu := minggu * (minTotal / 40)
// 	maxMinggu := minggu * (maxTotal / 40)

// 	// 🔹 Hitung kenaikan aktual
// 	kenaikan := *g.BeratBadan - bbAwal

// 	// 🔹 Generate penjelasan
// 	if kenaikan < minMinggu {
// 		g.PenjelasanHasilGrafik = "Kenaikan berat badan kurang dari standar"
// 	} else if kenaikan > maxMinggu {
// 		g.PenjelasanHasilGrafik = "Kenaikan berat badan berlebih dari standar"
// 	} else {
// 		g.PenjelasanHasilGrafik = "Kenaikan berat badan sesuai standar"
// 	}

// 	return nil
// }

// //
// // CREATE
// //
// func (u *grafikPeningkatanBBUsecase) Create(g *models.GrafikPeningkatanBB) error {
// 	if err := u.processGrafik(g); err != nil {
// 		return err
// 	}
// 	return u.repo.Create(g)
// }

// //
// // GET
// //
// func (u *grafikPeningkatanBBUsecase) GetByID(id int32) (*models.GrafikPeningkatanBB, error) {
// 	return u.repo.FindByID(id)
// }

// func (u *grafikPeningkatanBBUsecase) GetByKehamilanID(kehamilanID int32) ([]models.GrafikPeningkatanBB, error) {
// 	return u.repo.FindByKehamilanID(kehamilanID)
// }

// //
// // UPDATE
// //
// func (u *grafikPeningkatanBBUsecase) Update(g *models.GrafikPeningkatanBB) error {
// 	existing, err := u.repo.FindByID(g.ID)
// 	if err != nil {
// 		return errors.New("data grafik peningkatan berat badan tidak ditemukan")
// 	}

// 	// 🔹 Pastikan kehamilan tidak berubah
// 	g.KehamilanID = existing.KehamilanID

// 	if err := u.processGrafik(g); err != nil {
// 		return err
// 	}

// 	return u.repo.Update(g)
// }

// //
// // DELETE
// //
// func (u *grafikPeningkatanBBUsecase) Delete(id int32) error {
// 	return u.repo.Delete(id)
// }

// package usecases

// import (
// 	"errors"
// 	"monitoring-service/app/models"
// 	"monitoring-service/app/repositories"
// )

// type GrafikPeningkatanBBUsecase interface {
// 	Create(g *models.GrafikPeningkatanBB) error
// 	GetByID(id int32) (*models.GrafikPeningkatanBB, error)
// 	GetByKehamilanID(kehamilanID int32) ([]models.GrafikPeningkatanBB, error)
// 	Update(g *models.GrafikPeningkatanBB) error
// 	Delete(id int32) error

// 	// Modul Ibu
// 	GetGrafikBBForOrangtua(userID int32) (*GrafikBBOnTheFlyResponse, error)
// }

// // ─── Response untuk grafik BB on-the-fly ─────────────────────────────────────

// type GrafikBBOnTheFlyResponse struct {
// 	BBAwalan               float64          `json:"bb_awal"`
// 	IMTKategori            string           `json:"imt_kategori"`
// 	TargetMin              float64          `json:"target_kenaikan_min"`
// 	TargetMax              float64          `json:"target_kenaikan_max"`
// 	GrafikBB               []GrafikBBPoint  `json:"grafik_bb"`
// 	PenjelasanHasilGrafik  string           `json:"penjelasan_hasil_grafik"`
// }

// // GrafikBBPoint - satu titik pada grafik BB
// type GrafikBBPoint struct {
// 	MingguKehamilan int     `json:"minggu_kehamilan"`
// 	BeratBadan      float64 `json:"berat_badan"`
// 	KenaikanDariAwal float64 `json:"kenaikan_dari_awal"`
// 	BatasMinMinggu  float64 `json:"batas_min"`
// 	BatasMaxMinggu  float64 `json:"batas_max"`
// 	Status          string  `json:"status"` // "kurang" | "normal" | "lebih"
// }

// type grafikPeningkatanBBUsecase struct {
// 	repo          *repositories.GrafikPeningkatanBBRepository
// 	kehamilanRepo *repositories.KehamilanRepository
// }

// func NewGrafikPeningkatanBBUsecase(
// 	repo *repositories.GrafikPeningkatanBBRepository,
// 	kehamilanRepo *repositories.KehamilanRepository,
// ) GrafikPeningkatanBBUsecase {
// 	return &grafikPeningkatanBBUsecase{
// 		repo:          repo,
// 		kehamilanRepo: kehamilanRepo,
// 	}
// }

// // ─── Helper: tentukan range kenaikan BB & label IMT ──────────────────────────

// func bbRangeFromIMT(imt float64) (minTotal, maxTotal float64, label string) {
// 	switch {
// 	case imt < 18.5:
// 		return 12.5, 18, "Kurus (IMT < 18.5)"
// 	case imt < 25:
// 		return 11.5, 16, "Normal (IMT 18.5–24.9)"
// 	case imt < 30:
// 		return 7, 11.5, "Gemuk (IMT 25–29.9)"
// 	default:
// 		return 5, 9, "Obesitas (IMT ≥ 30)"
// 	}
// }

// // ─── CORE LOGIC (dipakai Create & Update) ────────────────────────────────────

// func (u *grafikPeningkatanBBUsecase) processGrafik(g *models.GrafikPeningkatanBB) error {
// 	if g.KehamilanID == 0 {
// 		return errors.New("kehamilan_id wajib diisi")
// 	}
// 	if g.BeratBadan == nil || g.MingguKehamilan == nil {
// 		return errors.New("berat_badan dan minggu_kehamilan wajib diisi")
// 	}

// 	kehamilan, err := u.kehamilanRepo.FindByID(g.KehamilanID)
// 	if err != nil {
// 		return errors.New("data kehamilan tidak ditemukan")
// 	}

// 	bbAwal := kehamilan.BB_Awal
// 	imt := kehamilan.IMT_Awal

// 	minTotal, maxTotal, _ := bbRangeFromIMT(imt)

// 	minggu := float64(*g.MingguKehamilan)
// 	minMinggu := minggu * (minTotal / 40)
// 	maxMinggu := minggu * (maxTotal / 40)

// 	kenaikan := *g.BeratBadan - bbAwal

// 	if kenaikan < minMinggu {
// 		g.PenjelasanHasilGrafik = "Kenaikan berat badan kurang dari standar"
// 	} else if kenaikan > maxMinggu {
// 		g.PenjelasanHasilGrafik = "Kenaikan berat badan berlebih dari standar"
// 	} else {
// 		g.PenjelasanHasilGrafik = "Kenaikan berat badan sesuai standar"
// 	}

// 	return nil
// }

// // ─── CREATE ──────────────────────────────────────────────────────────────────

// func (u *grafikPeningkatanBBUsecase) Create(g *models.GrafikPeningkatanBB) error {
// 	if err := u.processGrafik(g); err != nil {
// 		return err
// 	}
// 	return u.repo.Create(g)
// }

// // ─── GET ──────────────────────────────────────────────────────────────────────

// func (u *grafikPeningkatanBBUsecase) GetByID(id int32) (*models.GrafikPeningkatanBB, error) {
// 	return u.repo.FindByID(id)
// }

// func (u *grafikPeningkatanBBUsecase) GetByKehamilanID(kehamilanID int32) ([]models.GrafikPeningkatanBB, error) {
// 	return u.repo.FindByKehamilanID(kehamilanID)
// }

// // ─── UPDATE ──────────────────────────────────────────────────────────────────

// func (u *grafikPeningkatanBBUsecase) Update(g *models.GrafikPeningkatanBB) error {
// 	existing, err := u.repo.FindByID(g.ID)
// 	if err != nil {
// 		return errors.New("data grafik peningkatan berat badan tidak ditemukan")
// 	}

// 	g.KehamilanID = existing.KehamilanID

// 	if err := u.processGrafik(g); err != nil {
// 		return err
// 	}

// 	return u.repo.Update(g)
// }

// // ─── DELETE ──────────────────────────────────────────────────────────────────

// func (u *grafikPeningkatanBBUsecase) Delete(id int32) error {
// 	return u.repo.Delete(id)
// }

// // ─── GetGrafikBBForOrangtua - endpoint modul ibu /v2 ─────────────────────────

// func (u *grafikPeningkatanBBUsecase) GetGrafikBBForOrangtua(userID int32) (*GrafikBBOnTheFlyResponse, error) {
// 	// 1. Cari kehamilan aktif
// 	kehamilan, err := u.kehamilanRepo.FindAktifByUserID(userID)
// 	if err != nil {
// 		return nil, errors.New("kehamilan aktif tidak ditemukan")
// 	}

// 	bbAwal := kehamilan.BB_Awal
// 	imt := kehamilan.IMT_Awal

// 	minTotal, maxTotal, imtLabel := bbRangeFromIMT(imt)

// 	// 2. Ambil semua data BB untuk kehamilan ini
// 	list, err := u.repo.FindByKehamilanID(kehamilan.ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// 3. Build titik grafik
// 	var grafikBB []GrafikBBPoint
// 	var latestPenjelasan string

// 	for _, item := range list {
// 		if item.BeratBadan == nil || item.MingguKehamilan == nil {
// 			continue
// 		}

// 		minggu := float64(*item.MingguKehamilan)
// 		kenaikan := *item.BeratBadan - bbAwal

// 		minMinggu := minggu * (minTotal / 40)
// 		maxMinggu := minggu * (maxTotal / 40)

// 		var status string
// 		if kenaikan < minMinggu {
// 			status = "kurang"
// 		} else if kenaikan > maxMinggu {
// 			status = "lebih"
// 		} else {
// 			status = "normal"
// 		}

// 		grafikBB = append(grafikBB, GrafikBBPoint{
// 			MingguKehamilan:  *item.MingguKehamilan,
// 			BeratBadan:       *item.BeratBadan,
// 			KenaikanDariAwal: kenaikan,
// 			BatasMinMinggu:   minMinggu,
// 			BatasMaxMinggu:   maxMinggu,
// 			Status:           status,
// 		})

// 		latestPenjelasan = item.PenjelasanHasilGrafik
// 	}

// 	if len(grafikBB) == 0 {
// 		latestPenjelasan = "Belum ada data berat badan yang dicatat"
// 	}

// 	return &GrafikBBOnTheFlyResponse{
// 		BBAwalan:              bbAwal,
// 		IMTKategori:           imtLabel,
// 		TargetMin:             minTotal,
// 		TargetMax:             maxTotal,
// 		GrafikBB:              grafikBB,
// 		PenjelasanHasilGrafik: latestPenjelasan,
// 	}, nil
// }

// package usecases

// import (
// 	"errors"
// 	"monitoring-service/app/models"
// 	"monitoring-service/app/repositories"
// )

// type GrafikPeningkatanBBUsecase interface {
// 	Create(g *models.GrafikPeningkatanBB) error
// 	GetByID(id int32) (*models.GrafikPeningkatanBB, error)
// 	GetByKehamilanID(kehamilanID int32) ([]models.GrafikPeningkatanBB, error)
// 	Update(g *models.GrafikPeningkatanBB) error
// 	Delete(id int32) error

// 	// Modul Ibu
// 	GetGrafikBBForOrangtua(userID int32) (*GrafikBBOnTheFlyResponse, error)
// }

// // ─── Response untuk grafik BB on-the-fly ─────────────────────────────────────

// type GrafikBBOnTheFlyResponse struct {
// 	BBAwalan              float64         `json:"bb_awal"`
// 	IMTKategori           string          `json:"imt_kategori"`
// 	TargetMin             float64         `json:"target_kenaikan_min"`
// 	TargetMax             float64         `json:"target_kenaikan_max"`
// 	GrafikBB              []GrafikBBPoint `json:"grafik_bb"`
// 	PenjelasanHasilGrafik string          `json:"penjelasan_hasil_grafik"`
// }

// // GrafikBBPoint - satu titik pada grafik BB
// type GrafikBBPoint struct {
// 	MingguKehamilan  int     `json:"minggu_kehamilan"`
// 	BeratBadan       float64 `json:"berat_badan"`
// 	KenaikanDariAwal float64 `json:"kenaikan_dari_awal"`
// 	BatasMinMinggu   float64 `json:"batas_min"`
// 	BatasMaxMinggu   float64 `json:"batas_max"`
// 	Status           string  `json:"status"` // "kurang" | "normal" | "lebih"
// }

// type grafikPeningkatanBBUsecase struct {
// 	repo          *repositories.GrafikPeningkatanBBRepository
// 	kehamilanRepo *repositories.KehamilanRepository
// }

// func NewGrafikPeningkatanBBUsecase(
// 	repo *repositories.GrafikPeningkatanBBRepository,
// 	kehamilanRepo *repositories.KehamilanRepository,
// ) GrafikPeningkatanBBUsecase {
// 	return &grafikPeningkatanBBUsecase{
// 		repo:          repo,
// 		kehamilanRepo: kehamilanRepo,
// 	}
// }

// // ─── Helper: tentukan range kenaikan BB & label IMT ──────────────────────────

// func bbRangeFromIMT(imt float64) (minTotal, maxTotal float64, label string) {
// 	switch {
// 	case imt < 18.5:
// 		return 12.5, 18, "Kurus (IMT < 18.5)"
// 	case imt < 25:
// 		return 11.5, 16, "Normal (IMT 18.5–24.9)"
// 	case imt < 30:
// 		return 7, 11.5, "Gemuk (IMT 25–29.9)"
// 	default:
// 		return 5, 9, "Obesitas (IMT ≥ 30)"
// 	}
// }

// // ─── CORE LOGIC (dipakai Create & Update — endpoint tenaga kesehatan) ────────

// func (u *grafikPeningkatanBBUsecase) processGrafik(g *models.GrafikPeningkatanBB) error {
// 	if g.KehamilanID == 0 {
// 		return errors.New("kehamilan_id wajib diisi")
// 	}
// 	if g.BeratBadan == nil || g.MingguKehamilan == nil {
// 		return errors.New("berat_badan dan minggu_kehamilan wajib diisi")
// 	}

// 	kehamilan, err := u.kehamilanRepo.FindByID(g.KehamilanID)
// 	if err != nil {
// 		return errors.New("data kehamilan tidak ditemukan")
// 	}

// 	bbAwal := kehamilan.BB_Awal
// 	imt := kehamilan.IMT_Awal

// 	minTotal, maxTotal, _ := bbRangeFromIMT(imt)

// 	minggu := float64(*g.MingguKehamilan)
// 	minMinggu := minggu * (minTotal / 40)
// 	maxMinggu := minggu * (maxTotal / 40)

// 	kenaikan := *g.BeratBadan - bbAwal

// 	if kenaikan < minMinggu {
// 		g.PenjelasanHasilGrafik = "Kenaikan berat badan kurang dari standar"
// 	} else if kenaikan > maxMinggu {
// 		g.PenjelasanHasilGrafik = "Kenaikan berat badan berlebih dari standar"
// 	} else {
// 		g.PenjelasanHasilGrafik = "Kenaikan berat badan sesuai standar"
// 	}

// 	return nil
// }

// // ─── CREATE ──────────────────────────────────────────────────────────────────

// func (u *grafikPeningkatanBBUsecase) Create(g *models.GrafikPeningkatanBB) error {
// 	if err := u.processGrafik(g); err != nil {
// 		return err
// 	}
// 	return u.repo.Create(g)
// }

// // ─── GET ──────────────────────────────────────────────────────────────────────

// func (u *grafikPeningkatanBBUsecase) GetByID(id int32) (*models.GrafikPeningkatanBB, error) {
// 	return u.repo.FindByID(id)
// }

// func (u *grafikPeningkatanBBUsecase) GetByKehamilanID(kehamilanID int32) ([]models.GrafikPeningkatanBB, error) {
// 	return u.repo.FindByKehamilanID(kehamilanID)
// }

// // ─── UPDATE ──────────────────────────────────────────────────────────────────

// func (u *grafikPeningkatanBBUsecase) Update(g *models.GrafikPeningkatanBB) error {
// 	existing, err := u.repo.FindByID(g.ID)
// 	if err != nil {
// 		return errors.New("data grafik peningkatan berat badan tidak ditemukan")
// 	}

// 	g.KehamilanID = existing.KehamilanID

// 	if err := u.processGrafik(g); err != nil {
// 		return err
// 	}

// 	return u.repo.Update(g)
// }

// // ─── DELETE ──────────────────────────────────────────────────────────────────

// func (u *grafikPeningkatanBBUsecase) Delete(id int32) error {
// 	return u.repo.Delete(id)
// }

// // ─── GetGrafikBBForOrangtua - ON-THE-FLY dari pemeriksaan_kehamilan ──────────
// //
// // Pendekatan baru: data BB diambil langsung dari tabel pemeriksaan_kehamilan
// // (kolom berat_badan + tanggal_periksa), bukan dari tabel grafik_peningkatan_bb.
// // Minggu kehamilan dihitung dari selisih tanggal_periksa terhadap HPHT.

// func (u *grafikPeningkatanBBUsecase) GetGrafikBBForOrangtua(userID int32) (*GrafikBBOnTheFlyResponse, error) {
// 	// 1. Cari kehamilan aktif beserta BB_Awal, IMT_Awal, HPHT
// 	kehamilan, err := u.repo.FindKehamilanAktifForBB(userID)
// 	if err != nil {
// 		return nil, errors.New("kehamilan aktif tidak ditemukan")
// 	}

// 	bbAwal := kehamilan.BB_Awal
// 	imt := kehamilan.IMT_Awal
// 	hpht := kehamilan.HPHT

// 	minTotal, maxTotal, imtLabel := bbRangeFromIMT(imt)

// 	// 2. Ambil data BB dari pemeriksaan_kehamilan (sumber langsung)
// 	pemeriksaanList, err := u.repo.FindBBPemeriksaanForGrafik(kehamilan.ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// 3. Build titik grafik on-the-fly
// 	var grafikBB []GrafikBBPoint
// 	var latestStatus string
// 	var latestKenaikan float64

// 	for _, item := range pemeriksaanList {
// 		if item.TanggalPeriksa == nil || item.BeratBadan == nil {
// 			continue
// 		}

// 		// Hitung minggu kehamilan dari HPHT
// 		selisihHari := item.TanggalPeriksa.Sub(hpht).Hours() / 24
// 		minggu := int(selisihHari / 7)
// 		if minggu <= 0 {
// 			continue
// 		}

// 		mingguF := float64(minggu)
// 		kenaikan := *item.BeratBadan - bbAwal

// 		minMinggu := mingguF * (minTotal / 40)
// 		maxMinggu := mingguF * (maxTotal / 40)

// 		var status string
// 		if kenaikan < minMinggu {
// 			status = "kurang"
// 		} else if kenaikan > maxMinggu {
// 			status = "lebih"
// 		} else {
// 			status = "normal"
// 		}

// 		grafikBB = append(grafikBB, GrafikBBPoint{
// 			MingguKehamilan:  minggu,
// 			BeratBadan:       *item.BeratBadan,
// 			KenaikanDariAwal: kenaikan,
// 			BatasMinMinggu:   minMinggu,
// 			BatasMaxMinggu:   maxMinggu,
// 			Status:           status,
// 		})

// 		latestStatus = status
// 		latestKenaikan = kenaikan
// 	}

// 	// 4. Build penjelasan dari data terakhir
// 	var penjelasan string
// 	if len(grafikBB) == 0 {
// 		penjelasan = "Belum ada data berat badan yang dicatat"
// 	} else {
// 		switch latestStatus {
// 		case "kurang":
// 			penjelasan = "Kenaikan berat badan kurang dari standar"
// 		case "lebih":
// 			penjelasan = "Kenaikan berat badan berlebih dari standar"
// 		default:
// 			penjelasan = "Kenaikan berat badan sesuai standar"
// 		}
// 		_ = latestKenaikan
// 	}

// 	// 5. Pastikan grafik tidak nil
// 	if grafikBB == nil {
// 		grafikBB = []GrafikBBPoint{}
// 	}

// 	return &GrafikBBOnTheFlyResponse{
// 		BBAwalan:              bbAwal,
// 		IMTKategori:           imtLabel,
// 		TargetMin:             minTotal,
// 		TargetMax:             maxTotal,
// 		GrafikBB:              grafikBB,
// 		PenjelasanHasilGrafik: penjelasan,
// 	}, nil
// }

package usecases

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
)

type GrafikPeningkatanBBUsecase interface {
	Create(g *models.GrafikPeningkatanBB) error
	GetByID(id int32) (*models.GrafikPeningkatanBB, error)
	GetByKehamilanID(kehamilanID int32) ([]models.GrafikPeningkatanBB, error)
	Update(g *models.GrafikPeningkatanBB) error
	Delete(id int32) error

	// Modul Ibu
	GetGrafikBBForOrangtua(userID int32) (*GrafikBBOnTheFlyResponse, error)
}

// ─── Response untuk grafik BB on-the-fly ─────────────────────────────────────

type GrafikBBOnTheFlyResponse struct {
	BBAwalan              float64         `json:"bb_awal"`
	IMTKategori           string          `json:"imt_kategori"`
	TargetMin             float64         `json:"target_kenaikan_min"`
	TargetMax             float64         `json:"target_kenaikan_max"`
	GrafikBB              []GrafikBBPoint `json:"grafik_bb"`
	PenjelasanHasilGrafik string          `json:"penjelasan_hasil_grafik"`
}

// GrafikBBPoint - satu titik pada grafik BB
type GrafikBBPoint struct {
	MingguKehamilan  int     `json:"minggu_kehamilan"`
	BeratBadan       float64 `json:"berat_badan"`
	KenaikanDariAwal float64 `json:"kenaikan_dari_awal"`
	BatasMinMinggu   float64 `json:"batas_min"`
	BatasMaxMinggu   float64 `json:"batas_max"`
	Status           string  `json:"status"` // "kurang" | "normal" | "lebih"
}

type grafikPeningkatanBBUsecase struct {
	repo          *repositories.GrafikPeningkatanBBRepository
	kehamilanRepo *repositories.KehamilanRepository
}

func NewGrafikPeningkatanBBUsecase(
	repo *repositories.GrafikPeningkatanBBRepository,
	kehamilanRepo *repositories.KehamilanRepository,
) GrafikPeningkatanBBUsecase {
	return &grafikPeningkatanBBUsecase{
		repo:          repo,
		kehamilanRepo: kehamilanRepo,
	}
}

// ─── Parameter IOM 2009 per kategori IMT ─────────────────────────────────────

type iomParams struct {
	// Trimester 1 (minggu 1–13): rentang total kenaikan
	t1Min float64 // kg
	t1Max float64 // kg

	// Trimester 2–3 (minggu 14–40): kenaikan per minggu
	perMingguMin float64 // kg/minggu
	perMingguMax float64 // kg/minggu

	// Total target seluruh kehamilan (untuk tampilan info)
	totalMin float64
	totalMax float64

	label string
}

// PERHITUNGAN
func bbIOMParamsFromIMT(imt float64) iomParams {
	switch {
	case imt < 18.5:
		// Kurus: total 12.5–18 kg | T2-T3: 0.44–0.58 kg/minggu
		return iomParams{
			t1Min: 0.5, t1Max: 2.0,
			perMingguMin: 0.44, perMingguMax: 0.58,
			totalMin: 12.5, totalMax: 18.0,
			label: "Kurus (IMT < 18.5)",
		}
	case imt < 25.0:
		// Normal: total 11.5–16 kg | T2-T3: 0.35–0.50 kg/minggu
		return iomParams{
			t1Min: 0.5, t1Max: 2.0,
			perMingguMin: 0.35, perMingguMax: 0.50,
			totalMin: 11.5, totalMax: 16.0,
			label: "Normal (IMT 18.5–24.9)",
		}
	case imt < 30.0:
		// Gemuk: total 7–11.5 kg | T2-T3: 0.23–0.33 kg/minggu
		return iomParams{
			t1Min: 0.5, t1Max: 2.0,
			perMingguMin: 0.23, perMingguMax: 0.33,
			totalMin: 7.0, totalMax: 11.5,
			label: "Gemuk (IMT 25–29.9)",
		}
	default:
		// Obesitas: total 5–9 kg | T2-T3: 0.17–0.27 kg/minggu
		return iomParams{
			t1Min: 0.5, t1Max: 2.0,
			perMingguMin: 0.17, perMingguMax: 0.27,
			totalMin: 5.0, totalMax: 9.0,
			label: "Obesitas (IMT ≥ 30)",
		}
	}
}

// PERHITUNGAN
// Sumber : https://www.acog.org/clinical/clinical-guidance/committee-opinion/articles/2013/01/weight-gain-during-pregnancy
// Sumber : https://pmc.ncbi.nlm.nih.gov/articles/PMC2847829/
func hitungBatasIOM(minggu int, p iomParams) (batasMin, batasMax float64) {
	if minggu <= 0 {
		return 0, 0
	}

	if minggu <= 13 {
		// Trimester 1: linear dari 0 ke t1Min/t1Max
		fraksi := float64(minggu) / 13.0
		batasMin = fraksi * p.t1Min
		batasMax = fraksi * p.t1Max
	} else {
		// Trimester 2–3: mulai dari akhir T1 + slope per minggu
		mingguT23 := float64(minggu - 13)
		batasMin = p.t1Min + mingguT23*p.perMingguMin
		batasMax = p.t1Max + mingguT23*p.perMingguMax
	}

	return batasMin, batasMax
}

// ─── CORE LOGIC (endpoint tenaga kesehatan) ───────────────────────────────────

func (u *grafikPeningkatanBBUsecase) processGrafik(g *models.GrafikPeningkatanBB) error {
	if g.KehamilanID == 0 {
		return errors.New("kehamilan_id wajib diisi")
	}
	if g.BeratBadan == nil || g.MingguKehamilan == nil {
		return errors.New("berat_badan dan minggu_kehamilan wajib diisi")
	}

	kehamilan, err := u.kehamilanRepo.FindByID(g.KehamilanID)
	if err != nil {
		return errors.New("data kehamilan tidak ditemukan")
	}

	p := bbIOMParamsFromIMT(kehamilan.IMT_Awal)
	kenaikan := *g.BeratBadan - kehamilan.BB_Awal
	batasMin, batasMax := hitungBatasIOM(*g.MingguKehamilan, p)

	if kenaikan < batasMin {
		g.PenjelasanHasilGrafik = "Kenaikan berat badan kurang dari standar"
	} else if kenaikan > batasMax {
		g.PenjelasanHasilGrafik = "Kenaikan berat badan berlebih dari standar"
	} else {
		g.Deviasi = 0
		g.Status = "NORMAL"

		g.PenjelasanHasilGrafik = "Kenaikan berat badan sesuai standar"
	}

	return nil
}

// ─── CREATE ──────────────────────────────────────────────────────────────────

func (u *grafikPeningkatanBBUsecase) Create(g *models.GrafikPeningkatanBB) error {

	// VALIDASI AWAL (hindari panic)
	if g.MingguKehamilan == nil {
		return errors.New("minggu_kehamilan wajib diisi")
	}

	// CEK DUPLIKAT
	existing, err := u.repo.FindByKehamilanIDAndMinggu(g.KehamilanID, *g.MingguKehamilan)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("data minggu ini sudah ada")
	}

	// PROCESS
	if err := u.processGrafik(g); err != nil {
		return err
	}

	return u.repo.Create(g)
}

// ─── GET ──────────────────────────────────────────────────────────────────────

func (u *grafikPeningkatanBBUsecase) GetByID(id int32) (*models.GrafikPeningkatanBB, error) {
	return u.repo.FindByID(id)
}

func (u *grafikPeningkatanBBUsecase) GetByKehamilanID(kehamilanID int32) ([]models.GrafikPeningkatanBB, error) {
	return u.repo.FindByKehamilanID(kehamilanID)
}

// ─── UPDATE ──────────────────────────────────────────────────────────────────

func (u *grafikPeningkatanBBUsecase) Update(g *models.GrafikPeningkatanBB) error {

	existing, err := u.repo.FindByID(g.ID)
	if err != nil {
		return errors.New("data grafik peningkatan berat badan tidak ditemukan")
	}

	g.KehamilanID = existing.KehamilanID

	//  CEK DUPLIKAT MINGGU (PENTING)
	if g.MingguKehamilan != nil {
		existingMinggu, err := u.repo.FindByKehamilanIDAndMinggu(g.KehamilanID, *g.MingguKehamilan)
		if err != nil {
			return err
		}

		// pastikan bukan data yang sama
		if existingMinggu != nil && existingMinggu.ID != g.ID {
			return errors.New("data minggu ini sudah ada")
		}
	}

	if err := u.processGrafik(g); err != nil {
		return err
	}

	return u.repo.Update(g)
}

// ─── DELETE ──────────────────────────────────────────────────────────────────

func (u *grafikPeningkatanBBUsecase) Delete(id int32) error {
	return u.repo.Delete(id)
}

// ─── GetGrafikBBForOrangtua - ON-THE-FLY dari pemeriksaan_kehamilan ──────────

func (u *grafikPeningkatanBBUsecase) GetGrafikBBForOrangtua(userID int32) (*GrafikBBOnTheFlyResponse, error) {
	// 1. Cari kehamilan aktif
	kehamilan, err := u.repo.FindKehamilanAktifForBB(userID)
	if err != nil {
		return nil, errors.New("kehamilan aktif tidak ditemukan")
	}

	bbAwal := kehamilan.BB_Awal
	hpht := kehamilan.HPHT
	p := bbIOMParamsFromIMT(kehamilan.IMT_Awal)

	// 2. Ambil data BB dari pemeriksaan_kehamilan (sumber langsung)
	pemeriksaanList, err := u.repo.FindBBPemeriksaanForGrafik(kehamilan.ID)
	if err != nil {
		return nil, err
	}

	// 3. Build titik grafik menggunakan IOM 2009
	var grafikBB []GrafikBBPoint
	var latestStatus string

	for _, item := range pemeriksaanList {
		if item.TanggalPeriksa == nil || item.BeratBadan == nil {
			continue
		}

		// Hitung minggu kehamilan dari HPHT
		selisihHari := item.TanggalPeriksa.Sub(hpht).Hours() / 24
		minggu := int(selisihHari / 7)
		if minggu <= 0 {
			continue
		}

		kenaikan := *item.BeratBadan - bbAwal
		batasMin, batasMax := hitungBatasIOM(minggu, p)

		var status string
		if kenaikan < batasMin {
			status = "kurang"
		} else if kenaikan > batasMax {
			status = "lebih"
		} else {
			status = "normal"
		}

		grafikBB = append(grafikBB, GrafikBBPoint{
			MingguKehamilan:  minggu,
			BeratBadan:       *item.BeratBadan,
			KenaikanDariAwal: kenaikan,
			BatasMinMinggu:   batasMin,
			BatasMaxMinggu:   batasMax,
			Status:           status,
		})

		latestStatus = status
	}

	// 4. Penjelasan dari titik terakhir
	var penjelasan string
	if len(grafikBB) == 0 {
		penjelasan = "Belum ada data berat badan yang dicatat"
	} else {
		switch latestStatus {
		case "kurang":
			penjelasan = "Kenaikan berat badan kurang dari standar IOM 2009"
		case "lebih":
			penjelasan = "Kenaikan berat badan berlebih dari standar IOM 2009"
		default:
			penjelasan = "Kenaikan berat badan sesuai standar IOM 2009"
		}
	}

	// 5. Pastikan tidak nil
	if grafikBB == nil {
		grafikBB = []GrafikBBPoint{}
	}

	return &GrafikBBOnTheFlyResponse{
		BBAwalan:              bbAwal,
		IMTKategori:           p.label,
		TargetMin:             p.totalMin,
		TargetMax:             p.totalMax,
		GrafikBB:              grafikBB,
		PenjelasanHasilGrafik: penjelasan,
	}, nil
}
