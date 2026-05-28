// // package usecases

// // import (
// // 	"errors"
// // 	"monitoring-service/app/models"
// // 	"monitoring-service/app/repositories"
// // )

// // type PemeriksaanDokterTrimester1Usecase interface {
// // 	Create(p *models.PemeriksaanDokterTrimester1) error
// // 	GetByID(id int32) (*models.PemeriksaanDokterTrimester1, error)
// // 	GetMine(userID int32) (*models.PemeriksaanDokterTrimester1, error) // MODUL IBU
// // 	GetByKehamilanID(kehamilanID int32) ([]models.PemeriksaanDokterTrimester1, error)
// // 	Update(p *models.PemeriksaanDokterTrimester1) error
// // 	Delete(id int32) error
// // }

// // type pemeriksaanDokterTrimester1Usecase struct {
// // 	repo *repositories.PemeriksaanDokterTrimester1Repository
// // }

// // func NewPemeriksaanDokterTrimester1Usecase(repo *repositories.PemeriksaanDokterTrimester1Repository) PemeriksaanDokterTrimester1Usecase {
// // 	return &pemeriksaanDokterTrimester1Usecase{repo: repo}
// // }

// // func (u *pemeriksaanDokterTrimester1Usecase) Create(p *models.PemeriksaanDokterTrimester1) error {
// // 	if p.KehamilanID == 0 {
// // 		return errors.New("kehamilan_id wajib diisi")
// // 	}
// // 	return u.repo.Create(p)
// // }

// // func (u *pemeriksaanDokterTrimester1Usecase) GetByID(id int32) (*models.PemeriksaanDokterTrimester1, error) {
// // 	return u.repo.FindByID(id)
// // }

// // func (u *pemeriksaanDokterTrimester1Usecase) GetByKehamilanID(kehamilanID int32) ([]models.PemeriksaanDokterTrimester1, error) {
// // 	return u.repo.FindByKehamilanID(kehamilanID)
// // }

// // func (u *pemeriksaanDokterTrimester1Usecase) Update(p *models.PemeriksaanDokterTrimester1) error {
// // 	_, err := u.repo.FindByID(p.ID)
// // 	if err != nil {
// // 		return errors.New("data pemeriksaan dokter trimester 1 tidak ditemukan")
// // 	}
// // 	return u.repo.Update(p)
// // }

// // func (u *pemeriksaanDokterTrimester1Usecase) Delete(id int32) error {
// // 	return u.repo.Delete(id)
// // }

// // // MODUL IBU
// // func (u *pemeriksaanDokterTrimester1Usecase) GetMine(userID int32) (*models.PemeriksaanDokterTrimester1, error) {
// // 	return u.repo.FindMineByUserID(userID)
// // }


// package usecases

// import (
// 	"errors"
// 	"monitoring-service/app/models"
// 	"monitoring-service/app/repositories"
// 	"time"
// )

// type PemeriksaanDokterTrimester1Usecase interface {
// 	Create(p *models.PemeriksaanDokterTrimester1) error
// 	GetByID(id int32) (*models.PemeriksaanDokterTrimester1, error)
// 	GetMine(userID int32) (*models.PemeriksaanDokterTrimester1, error)     // single (lama)
// 	GetAllMine(userID int32) ([]models.PemeriksaanDokterTrimester1, error) // list (BARU)
// 	GetByKehamilanID(kehamilanID int32) ([]models.PemeriksaanDokterTrimester1, error)
// 	Update(p *models.PemeriksaanDokterTrimester1) error
// 	Delete(id int32) error
// }

// type PemeriksaanDokterTrimester1Request struct {
// 	KehamilanID int32  `json:"kehamilan_id"`
// 	NamaDokter  string `json:"nama_dokter"`
// 	// TanggalPeriksa dihapus - akan auto set ke hari ini di backend
// 	KonsepAnamnesaPemeriksaan string `json:"konsep_anamnesa_pemeriksaan"`
// 	GambarUSG                 string `json:"gambar_usg"` // Base64 encoded image
// 	FisikKonjungtiva          string `json:"fisik_konjungtiva"`
// 	FisikSklera               string `json:"fisik_sklera"`
// 	FisikKulit                string `json:"fisik_kulit"`
// 	FisikLeher                string `json:"fisik_leher"`
// 	FisikGigiMulut            string `json:"fisik_gigi_mulut"`
// 	FisikTHT                  string `json:"fisik_tht"`
// 	FisikDadaJantung          string `json:"fisik_dada_jantung"`
// 	FisikDadaParu             string `json:"fisik_dada_paru"`
// 	FisikPerut                string `json:"fisik_perut"`
// 	FisikTungkai              string `json:"fisik_tungkai"`
// 	HPHT                      string `json:"hpht"` // Kept as tidak redundant (dari riwayat)
// 	KeteraturanHaid           string `json:"keteraturan_haid"`
// 	UmurHamilHPHTMinggu       *int   `json:"umur_hamil_hpht_minggu"`
// 	// HPLBerdasarkanHPHT dihapus - calculated
// 	UmurHamilUSGMinggu *int `json:"umur_hamil_usg_minggu"`
// 	// HPLBerdasarkanUSG dihapus - calculated
// 	USGJumlahGS                 string   `json:"usg_jumlah_gs"`
// 	USGDiameterGS_cm            *float64 `json:"usg_diameter_gs_cm"`
// 	USGDiameterGSMinggu         *int     `json:"usg_diameter_gs_minggu"`
// 	USGDiameterGSHari           *int     `json:"usg_diameter_gs_hari"`
// 	USGJumlahBayi               string   `json:"usg_jumlah_bayi"`
// 	USGCRL_cm                   *float64 `json:"usg_crl_cm"`
// 	USGCRLMinggu                *int     `json:"usg_crl_minggu"`
// 	USGCRLHari                  *int     `json:"usg_crl_hari"`
// 	USGLetakProdukKehamilan     string   `json:"usg_letak_produk_kehamilan"`
// 	USGPulsasiJantung           string   `json:"usg_pulsasi_jantung"`
// 	USGKecurigaanTemuanAbnormal string   `json:"usg_kecurigaan_temuan_abnormal"`
// 	USGKeteranganTemuanAbnormal string   `json:"usg_keterangan_temuan_abnormal"`

// 	// Laboratorium & Skrining Jiwa (Trimester 1)
// 	// TanggalLabJiwa dihapus - akan auto set
// 	LabHemoglobinHasilJiwa        *float64 `json:"lab_hemoglobin_hasil_jiwa"`
// 	LabHemoglobinRencanaJiwa      string   `json:"lab_hemoglobin_rencana_tindak_lanjut_jiwa"`
// 	LabGolonganDarahRhesusHasil   string   `json:"lab_golongan_darah_rhesus_hasil"`
// 	LabGolonganDarahRhesusRencana string   `json:"lab_golongan_darah_rhesus_rencana_tindak_lanjut"`
// 	LabGulaDarahSewaktuHasil      *int     `json:"lab_gula_darah_sewaktu_hasil"`
// 	LabGulaDarahSewaktuRencana    string   `json:"lab_gula_darah_sewaktu_rencana_tindak_lanjut"`
// 	LabHIVHasil                   string   `json:"lab_hiv_hasil"`
// 	LabHIVRencana                 string   `json:"lab_hiv_rencana_tindak_lanjut"`
// 	LabSifilisHasil               string   `json:"lab_sifilis_hasil"`
// 	LabSifilisRencana             string   `json:"lab_sifilis_rencana_tindak_lanjut"`
// 	LabHepatitisBHasil            string   `json:"lab_hepatitis_b_hasil"`
// 	LabHepatitisBRencana          string   `json:"lab_hepatitis_b_rencana_tindak_lanjut"`
// 	// TanggalSkriningJiwa dihapus - akan auto set
// 	SkriningJiwaHasil        string `json:"skrining_jiwa_hasil"`
// 	SkriningJiwaTindakLanjut string `json:"skrining_jiwa_tindak_lanjut"`
// 	SkriningJiwaPerluRujukan string `json:"skrining_jiwa_perlu_rujukan"`
// 	KesimpulanJiwa           string `json:"kesimpulan"`
// 	RekomendasiJiwa          string `json:"rekomendasi"`
// }

// type PemeriksaanDokterTrimester1Response struct {
// 	Dokter *models.PemeriksaanDokterTrimester1 `json:"dokter"`
// 	Lab    *models.PemeriksaanLaboratoriumJiwa `json:"lab_jiwa,omitempty"`
// }

// type pemeriksaanDokterTrimester1Usecase struct {
// 	repo    *repositories.PemeriksaanDokterTrimester1Repository
// 	labRepo *repositories.PemeriksaanLaboratoriumJiwaRepository
// }

// func NewPemeriksaanDokterTrimester1Usecase(
// 	repo *repositories.PemeriksaanDokterTrimester1Repository,
// 	labRepo *repositories.PemeriksaanLaboratoriumJiwaRepository,
// ) PemeriksaanDokterTrimester1Usecase {
// 	return &pemeriksaanDokterTrimester1Usecase{repo: repo, labRepo: labRepo}
// }

// func (u *pemeriksaanDokterTrimester1Usecase) Create(req *PemeriksaanDokterTrimester1Request) error {
// 	if req.KehamilanID == 0 {
// 		return errors.New("kehamilan_id wajib diisi")
// 	}
// 	dokter := u.mapRequestToDokter(req)
// 	lab := u.mapRequestToLab(req, 1)
// 	return u.repo.CreateWithLab(dokter, lab)
// }

// func (u *pemeriksaanDokterTrimester1Usecase) GetByID(id int32) (*models.PemeriksaanDokterTrimester1, error) {
// 	return u.repo.GetByID(id) // FIX: was FindByID
// }

// func (u *pemeriksaanDokterTrimester1Usecase) GetByKehamilanID(kehamilanID int32) ([]models.PemeriksaanDokterTrimester1, error) {
// 	return u.repo.GetByKehamilanID(kehamilanID) // FIX: was FindByKehamilanID
// }

// func (u *pemeriksaanDokterTrimester1Usecase) Update(p *models.PemeriksaanDokterTrimester1) error {
// 	_, err := u.repo.GetByID(p.ID) // FIX: was FindByID
// 	if err != nil {
// 		return errors.New("data tidak ditemukan")
// 	}
// 	dokter := u.mapRequestToDokter(req)
// 	dokter.KehamilanID = existing.KehamilanID // biar konsisten
// 	lab := u.mapRequestToLab(req, 1)
// 	return u.repo.UpdateWithLab(id, dokter, lab)
// }

// func (u *pemeriksaanDokterTrimester1Usecase) mapRequestToDokter(req *PemeriksaanDokterTrimester1Request) *models.PemeriksaanDokterTrimester1 {
// 	today := time.Now()
// 	dokter := &models.PemeriksaanDokterTrimester1{
// 		KehamilanID:                 req.KehamilanID,
// 		NamaDokter:                  req.NamaDokter,
// 		TanggalPeriksa:              &today, // Set to today
// 		KonsepAnamnesaPemeriksaan:   req.KonsepAnamnesaPemeriksaan,
// 		GambarUSG:                   req.GambarUSG,
// 		FisikKonjungtiva:            req.FisikKonjungtiva,
// 		FisikSklera:                 req.FisikSklera,
// 		FisikKulit:                  req.FisikKulit,
// 		FisikLeher:                  req.FisikLeher,
// 		FisikGigiMulut:              req.FisikGigiMulut,
// 		FisikTHT:                    req.FisikTHT,
// 		FisikDadaJantung:            req.FisikDadaJantung,
// 		FisikDadaParu:               req.FisikDadaParu,
// 		FisikPerut:                  req.FisikPerut,
// 		FisikTungkai:                req.FisikTungkai,
// 		KeteraturanHaid:             req.KeteraturanHaid,
// 		UmurHamilHPHTMinggu:         req.UmurHamilHPHTMinggu,
// 		UmurHamilUSGMinggu:          req.UmurHamilUSGMinggu,
// 		USGJumlahGS:                 req.USGJumlahGS,
// 		USGDiameterGS_cm:            req.USGDiameterGS_cm,
// 		USGDiameterGSMinggu:         req.USGDiameterGSMinggu,
// 		USGDiameterGSHari:           req.USGDiameterGSHari,
// 		USGJumlahBayi:               req.USGJumlahBayi,
// 		USGCRL_cm:                   req.USGCRL_cm,
// 		USGCRLMinggu:                req.USGCRLMinggu,
// 		USGCRLHari:                  req.USGCRLHari,
// 		USGLetakProdukKehamilan:     req.USGLetakProdukKehamilan,
// 		USGPulsasiJantung:           req.USGPulsasiJantung,
// 		USGKecurigaanTemuanAbnormal: req.USGKecurigaanTemuanAbnormal,
// 		USGKeteranganTemuanAbnormal: req.USGKeteranganTemuanAbnormal,
// 	}
// 	if req.HPHT != "" {
// 		t, _ := time.Parse("2006-01-02", req.HPHT)
// 		dokter.HPHT = &t
// 	}
// 	// HPL fields are auto-calculated, no need to parse from request
// 	return dokter
// }

// func (u *pemeriksaanDokterTrimester1Usecase) mapRequestToLab(req *PemeriksaanDokterTrimester1Request, trimester int32) *models.PemeriksaanLaboratoriumJiwa {
// 	if req.LabHemoglobinHasilJiwa == nil && req.LabGulaDarahSewaktuHasil == nil {
// 		return nil
// 	}
// 	today := time.Now()
// 	lab := &models.PemeriksaanLaboratoriumJiwa{
// 		Trimester:                        trimester,
// 		TanggalLab:                       &today, // Set to today
// 		LabHemoglobinHasil:               req.LabHemoglobinHasilJiwa,
// 		LabHemoglobinRencanaTindakLanjut: req.LabHemoglobinRencanaJiwa,
// 		LabGolonganDarahRhesusHasil:      req.LabGolonganDarahRhesusHasil,
// 		LabGolonganDarahRhesusRencana:    req.LabGolonganDarahRhesusRencana,
// 		LabGulaDarahSewaktuHasil:         req.LabGulaDarahSewaktuHasil,
// 		LabGulaDarahSewaktuRencana:       req.LabGulaDarahSewaktuRencana,
// 		LabHIVHasil:                      req.LabHIVHasil,
// 		LabHIVRencana:                    req.LabHIVRencana,
// 		LabSifilisHasil:                  req.LabSifilisHasil,
// 		LabSifilisRencana:                req.LabSifilisRencana,
// 		LabHepatitisBHasil:               req.LabHepatitisBHasil,
// 		LabHepatitisBRencana:             req.LabHepatitisBRencana,
// 		SkriningJiwaHasil:                req.SkriningJiwaHasil,
// 		SkriningJiwaTindakLanjut:         req.SkriningJiwaTindakLanjut,
// 		SkriningJiwaPerluRujukan:         req.SkriningJiwaPerluRujukan,
// 		Kesimpulan:                       req.KesimpulanJiwa,
// 		Rekomendasi:                      req.RekomendasiJiwa,
// 	}
// 	if req.SkriningJiwaHasil != "" {
// 		tab := &today
// 		lab.TanggalSkriningJiwa = tab // Set to today if skrining data provided
// 	}
// 	return lab
// }

// func (u *pemeriksaanDokterTrimester1Usecase) GetByID(id int32) (*PemeriksaanDokterTrimester1Response, error) {
// 	dokter, err := u.repo.FindByID(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	lab, _ := u.labRepo.FindByKehamilanIDAndTrimester(dokter.KehamilanID, 1)
// 	return &PemeriksaanDokterTrimester1Response{Dokter: dokter, Lab: lab}, nil
// }

// func (u *pemeriksaanDokterTrimester1Usecase) GetByKehamilanID(kehamilanID int32) ([]PemeriksaanDokterTrimester1Response, error) {
// 	list, err := u.repo.FindByKehamilanID(kehamilanID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var result []PemeriksaanDokterTrimester1Response
// 	for _, d := range list {
// 		lab, _ := u.labRepo.FindByKehamilanIDAndTrimester(d.KehamilanID, 1)
// 		result = append(result, PemeriksaanDokterTrimester1Response{Dokter: &d, Lab: lab})
// 	}
// 	return result, nil
// }

// func (u *pemeriksaanDokterTrimester1Usecase) Delete(id int32) error {
// 	return u.repo.Delete(id)
// }

// // MODUL IBU - single (existing)
// func (u *pemeriksaanDokterTrimester1Usecase) GetMine(userID int32) (*models.PemeriksaanDokterTrimester1, error) {
// 	return u.repo.FindMineByUserID(userID)
// }

// // MODUL IBU - list (BARU)
// func (u *pemeriksaanDokterTrimester1Usecase) GetAllMine(userID int32) ([]models.PemeriksaanDokterTrimester1, error) {
// 	return u.repo.FindAllMineByUserID(userID)
// }

// package usecases

// import (
// 	"errors"
// 	"monitoring-service/app/models"
// 	"monitoring-service/app/repositories"
// )

// type PemeriksaanDokterTrimester1Usecase interface {
// 	Create(p *models.PemeriksaanDokterTrimester1) error
// 	GetByID(id int32) (*models.PemeriksaanDokterTrimester1, error)
// 	GetMine(userID int32) (*models.PemeriksaanDokterTrimester1, error)     // single (lama)
// 	GetAllMine(userID int32) ([]models.PemeriksaanDokterTrimester1, error) // list (BARU)
// 	GetByKehamilanID(kehamilanID int32) ([]models.PemeriksaanDokterTrimester1, error)
// 	Update(p *models.PemeriksaanDokterTrimester1) error
// 	Delete(id int32) error
// }

// type pemeriksaanDokterTrimester1Usecase struct {
// 	repo *repositories.PemeriksaanDokterTrimester1Repository
// }

// func NewPemeriksaanDokterTrimester1Usecase(repo *repositories.PemeriksaanDokterTrimester1Repository) PemeriksaanDokterTrimester1Usecase {
// 	return &pemeriksaanDokterTrimester1Usecase{repo: repo}
// }

// func (u *pemeriksaanDokterTrimester1Usecase) Create(p *models.PemeriksaanDokterTrimester1) error {
// 	if p.KehamilanID == 0 {
// 		return errors.New("kehamilan_id wajib diisi")
// 	}
// 	return u.repo.Create(p)
// }

// func (u *pemeriksaanDokterTrimester1Usecase) GetByID(id int32) (*models.PemeriksaanDokterTrimester1, error) {
// 	return u.repo.GetByID(id) // FIX: was FindByID
// }

// func (u *pemeriksaanDokterTrimester1Usecase) GetByKehamilanID(kehamilanID int32) ([]models.PemeriksaanDokterTrimester1, error) {
// 	return u.repo.GetByKehamilanID(kehamilanID) // FIX: was FindByKehamilanID
// }

// func (u *pemeriksaanDokterTrimester1Usecase) Update(id int32, p *models.PemeriksaanDokterTrimester1) error {
// 	_, err := u.repo.GetByID(p.ID) // FIX: was FindByID
// 	if err != nil {
// 		return errors.New("data pemeriksaan dokter trimester 1 tidak ditemukan")
// 	}
// 	p.ID = id
// 	return u.repo.Update(p)
// }

// func (u *pemeriksaanDokterTrimester1Usecase) Delete(id int32) error {
// 	return u.repo.Delete(id)
// }

// // MODUL IBU - single (existing)
// func (u *pemeriksaanDokterTrimester1Usecase) GetMine(userID int32) (*models.PemeriksaanDokterTrimester1, error) {
// 	return u.repo.FindMineByUserID(userID)
// }

// // MODUL IBU - list (BARU)
// func (u *pemeriksaanDokterTrimester1Usecase) GetAllMine(userID int32) ([]models.PemeriksaanDokterTrimester1, error) {
// 	return u.repo.FindAllMineByUserID(userID)
// }

package usecases

import (
    "errors"
    "monitoring-service/app/models"
    "monitoring-service/app/repositories"
)

type PemeriksaanDokterTrimester1Usecase interface {
    Create(p *models.PemeriksaanDokterTrimester1) error
    GetByID(id int32) (*models.PemeriksaanDokterTrimester1, error)
    GetMine(userID int32) (*models.PemeriksaanDokterTrimester1, error)
    GetAllMine(userID int32) ([]models.PemeriksaanDokterTrimester1, error)
    GetByKehamilanID(kehamilanID int32) ([]models.PemeriksaanDokterTrimester1, error)
    Update(id int32, p *models.PemeriksaanDokterTrimester1) error // DIUBAH
    Delete(id int32) error
}

type pemeriksaanDokterTrimester1Usecase struct {
    repo    *repositories.PemeriksaanDokterTrimester1Repository
    labRepo *repositories.PemeriksaanLaboratoriumJiwaRepository // DITAMBAHKAN
}

// DIUBAH: Menambahkan parameter labRepo agar cocok dengan init.go
func NewPemeriksaanDokterTrimester1Usecase(repo *repositories.PemeriksaanDokterTrimester1Repository, labRepo *repositories.PemeriksaanLaboratoriumJiwaRepository) PemeriksaanDokterTrimester1Usecase {
    return &pemeriksaanDokterTrimester1Usecase{repo: repo, labRepo: labRepo}
}

func (u *pemeriksaanDokterTrimester1Usecase) Create(p *models.PemeriksaanDokterTrimester1) error {
    if p.KehamilanID == 0 {
        return errors.New("kehamilan_id wajib diisi")
    }
    return u.repo.Create(p)
}

func (u *pemeriksaanDokterTrimester1Usecase) GetByID(id int32) (*models.PemeriksaanDokterTrimester1, error) {
    return u.repo.FindByID(id)
}

func (u *pemeriksaanDokterTrimester1Usecase) GetByKehamilanID(kehamilanID int32) ([]models.PemeriksaanDokterTrimester1, error) {
    return u.repo.FindByKehamilanID(kehamilanID)
}

// DIUBAH: Logic disesuaikan dengan parameter id
func (u *pemeriksaanDokterTrimester1Usecase) Update(id int32, p *models.PemeriksaanDokterTrimester1) error {
    _, err := u.repo.FindByID(id)
    if err != nil {
        return errors.New("data pemeriksaan dokter trimester 1 tidak ditemukan")
    }
    p.ID = id
    return u.repo.Update(p)
}

func (u *pemeriksaanDokterTrimester1Usecase) Delete(id int32) error {
    return u.repo.Delete(id)
}

func (u *pemeriksaanDokterTrimester1Usecase) GetMine(userID int32) (*models.PemeriksaanDokterTrimester1, error) {
    return u.repo.FindMineByUserID(userID)
}

func (u *pemeriksaanDokterTrimester1Usecase) GetAllMine(userID int32) ([]models.PemeriksaanDokterTrimester1, error) {
    return u.repo.FindAllMineByUserID(userID)
}