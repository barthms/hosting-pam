// // package usecases

// // import (
// // 	"errors"
// // 	"monitoring-service/app/models"
// // 	"monitoring-service/app/repositories"
// // )

// // type PemeriksaanDokterTrimester3Usecase interface {
// // 	Create(p *models.PemeriksaanDokterTrimester3) error
// // 	GetByID(id int32) (*models.PemeriksaanDokterTrimester3, error)
// // 	GetMine(userID int32) (*models.PemeriksaanDokterTrimester3, error) // MODUL IBU
// // 	GetByKehamilanID(kehamilanID int32) ([]models.PemeriksaanDokterTrimester3, error)
// // 	Update(p *models.PemeriksaanDokterTrimester3) error
// // 	Delete(id int32) error
// // }

// // type pemeriksaanDokterTrimester3Usecase struct {
// // 	repo *repositories.PemeriksaanDokterTrimester3Repository
// // }

// // func NewPemeriksaanDokterTrimester3Usecase(repo *repositories.PemeriksaanDokterTrimester3Repository) PemeriksaanDokterTrimester3Usecase {
// // 	return &pemeriksaanDokterTrimester3Usecase{repo: repo}
// // }

// // func (u *pemeriksaanDokterTrimester3Usecase) Create(p *models.PemeriksaanDokterTrimester3) error {
// // 	if p.KehamilanID == 0 {
// // 		return errors.New("kehamilan_id wajib diisi")
// // 	}
// // 	return u.repo.Create(p)
// // }

// // func (u *pemeriksaanDokterTrimester3Usecase) GetByID(id int32) (*models.PemeriksaanDokterTrimester3, error) {
// // 	return u.repo.FindByID(id)
// // }

// // func (u *pemeriksaanDokterTrimester3Usecase) GetByKehamilanID(kehamilanID int32) ([]models.PemeriksaanDokterTrimester3, error) {
// // 	return u.repo.FindByKehamilanID(kehamilanID)
// // }

// // func (u *pemeriksaanDokterTrimester3Usecase) Update(p *models.PemeriksaanDokterTrimester3) error {
// // 	_, err := u.repo.FindByID(p.ID)
// // 	if err != nil {
// // 		return errors.New("data pemeriksaan dokter trimester 3 tidak ditemukan")
// // 	}
// // 	return u.repo.Update(p)
// // }

// // func (u *pemeriksaanDokterTrimester3Usecase) Delete(id int32) error {
// // 	return u.repo.Delete(id)
// // }

// // // MODUL IBU
// // func (u *pemeriksaanDokterTrimester3Usecase) GetMine(userID int32) (*models.PemeriksaanDokterTrimester3, error) {
// // 	return u.repo.FindMineByUserID(userID)
// // }

// package usecases

// import (
// 	"errors"
// 	"monitoring-service/app/models"
// 	"monitoring-service/app/repositories"
// 	"time"
// )

// type PemeriksaanDokterTrimester3Usecase interface {
// 	Create(p *models.PemeriksaanDokterTrimester3) error
// 	GetByID(id int32) (*models.PemeriksaanDokterTrimester3, error)
// 	GetMine(userID int32) (*models.PemeriksaanDokterTrimester3, error)    // single (lama)
// 	GetAllMine(userID int32) ([]models.PemeriksaanDokterTrimester3, error) // list (BARU)
// 	GetByKehamilanID(kehamilanID int32) ([]models.PemeriksaanDokterTrimester3, error)
// 	Update(p *models.PemeriksaanDokterTrimester3) error
// 	Delete(id int32) error
// }

// type PemeriksaanDokterTrimester3Request struct {
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

// 	USGTrimester3Dilakukan                   string   `json:"usg_trimester_3_dilakukan"`
// 	UKBerdasarkanUSGTrimester1Minggu         *int     `json:"uk_berdasarkan_usg_trimester_1_minggu"`
// 	UKBerdasarkanHPHTMinggu                  *int     `json:"uk_berdasarkan_hpht_minggu"`
// 	UKBerdasarkanBiometriUSGTrimester3Minggu *int     `json:"uk_berdasarkan_biometri_usg_trimester_3_minggu"`
// 	SelisihUK3MingguAtauLebih                string   `json:"selisih_uk_3_minggu_atau_lebih"`
// 	USGJumlahBayi                            string   `json:"usg_jumlah_bayi"`
// 	USGLetakBayi                             string   `json:"usg_letak_bayi"`
// 	USGPresentasiBayi                        string   `json:"usg_presentasi_bayi"`
// 	USGKeadaanBayi                           string   `json:"usg_keadaan_bayi"`
// 	USGDJNilai                               *int     `json:"usg_djj_nilai"`
// 	USGDJJStatus                             string   `json:"usg_djj_status"`
// 	USGLokasiPlasenta                        string   `json:"usg_lokasi_plasenta"`
// 	USGCairanKetubanSDPCm                    *float64 `json:"usg_cairan_ketuban_sdp_cm"`
// 	USGCairanKetubanStatus                   string   `json:"usg_cairan_ketuban_status"`

// 	BiometriBPDCm        *float64 `json:"biometri_bpd_cm"`
// 	BiometriBPDMinggu    *int     `json:"biometri_bpd_minggu"`
// 	BiometriHCCm         *float64 `json:"biometri_hc_cm"`
// 	BiometriHCMinggu     *int     `json:"biometri_hc_minggu"`
// 	BiometriACCm         *float64 `json:"biometri_ac_cm"`
// 	BiometriACMinggu     *int     `json:"biometri_ac_minggu"`
// 	BiometriFLCm         *float64 `json:"biometri_fl_cm"`
// 	BiometriFLMinggu     *int     `json:"biometri_fl_minggu"`
// 	BiometriEFWTBJGram   *int     `json:"biometri_efw_tbj_gram"`
// 	BiometriEFWTBJMinggu *int     `json:"biometri_efw_tbj_minggu"`

// 	USGKecurigaanTemuanAbnormal string `json:"usg_kecurigaan_temuan_abnormal"`
// 	USGKeteranganTemuanAbnormal string `json:"usg_keterangan_temuan_abnormal"`

// 	// LANJUTAN TRIMESTER 3 (digabung)
// 	HasilUSGCatatan string `json:"hasil_usg_catatan"`
// 	// TanggalLab dihapus - akan auto set
// 	LabHemoglobinHasil    *float64 `json:"lab_hemoglobin_hasil"`
// 	LabHemoglobinRencana  string   `json:"lab_hemoglobin_rencana_tindak_lanjut"`
// 	LabProteinUrinHasil   *int     `json:"lab_protein_urin_hasil"`
// 	LabProteinUrinRencana string   `json:"lab_protein_urin_rencana_tindak_lanjut"`
// 	LabUrinReduksiHasil   string   `json:"lab_urin_reduksi_hasil"`
// 	LabUrinReduksiRencana string   `json:"lab_urin_reduksi_rencana_tindak_lanjut"`
// 	// TanggalSkriningJiwa dihapus - akan auto set
// 	SkriningJiwaHasil                     string `json:"skrining_jiwa_hasil"`
// 	SkriningJiwaTindakLanjut              string `json:"skrining_jiwa_tindak_lanjut"`
// 	SkriningJiwaPerluRujukan              string `json:"skrining_jiwa_perlu_rujukan"`
// 	RencanaKonsultasiGizi                 bool   `json:"rencana_konsultasi_gizi"`
// 	RencanaKonsultasiKebidanan            bool   `json:"rencana_konsultasi_kebidanan"`
// 	RencanaKonsultasiAnak                 bool   `json:"rencana_konsultasi_anak"`
// 	RencanaKonsultasiPenyakitDalam        bool   `json:"rencana_konsultasi_penyakit_dalam"`
// 	RencanaKonsultasiNeurologi            bool   `json:"rencana_konsultasi_neurologi"`
// 	RencanaKonsultasiTHT                  bool   `json:"rencana_konsultasi_tht"`
// 	RencanaKonsultasiPsikiatri            bool   `json:"rencana_konsultasi_psikiatri"`
// 	RencanaKonsultasiLainLain             string `json:"rencana_konsultasi_lain_lain"`
// 	RencanaProsesMelahirkan               string `json:"rencana_proses_melahirkan"`
// 	RencanaKontrasepsiAKDR                bool   `json:"rencana_kontrasepsi_akdr"`
// 	RencanaKontrasepsiPil                 bool   `json:"rencana_kontrasepsi_pil"`
// 	RencanaKontrasepsiSuntik              bool   `json:"rencana_kontrasepsi_suntik"`
// 	RencanaKontrasepsiSteril              bool   `json:"rencana_kontrasepsi_steril"`
// 	RencanaKontrasepsiMAL                 bool   `json:"rencana_kontrasepsi_mal"`
// 	RencanaKontrasepsiImplan              bool   `json:"rencana_kontrasepsi_implan"`
// 	RencanaKontrasepsiBelumMemilih        bool   `json:"rencana_kontrasepsi_belum_memilih"`
// 	KebutuhanKonseling                    string `json:"kebutuhan_konseling"`
// 	Penjelasan                            string `json:"penjelasan"`
// 	KesimpulanRekomendasiTempatMelahirkan string `json:"kesimpulan_rekomendasi_tempat_melahirkan"`

// 	// Laboratorium & Skrining Jiwa untuk trimester 3 (opsional)
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
// 	// TanggalSkriningJiwaTr dihapus - akan auto set
// 	SkriningJiwaHasilTr        string `json:"skrining_jiwa_hasil_tr"`
// 	SkriningJiwaTindakLanjutTr string `json:"skrining_jiwa_tindak_lanjut_tr"`
// 	SkriningJiwaPerluRujukanTr string `json:"skrining_jiwa_perlu_rujukan_tr"`
// 	KesimpulanJiwaTr           string `json:"kesimpulan_tr"`
// 	RekomendasiJiwaTr          string `json:"rekomendasi_tr"`
// }

// type PemeriksaanDokterTrimester3Response struct {
// 	Dokter *models.PemeriksaanDokterTrimester3 `json:"dokter"`
// 	Lab    *models.PemeriksaanLaboratoriumJiwa `json:"lab_jiwa,omitempty"`
// }

// type pemeriksaanDokterTrimester3Usecase struct {
// 	repo    *repositories.PemeriksaanDokterTrimester3Repository
// 	labRepo *repositories.PemeriksaanLaboratoriumJiwaRepository
// }

// func NewPemeriksaanDokterTrimester3Usecase(
// 	repo *repositories.PemeriksaanDokterTrimester3Repository,
// 	labRepo *repositories.PemeriksaanLaboratoriumJiwaRepository,
// ) PemeriksaanDokterTrimester3Usecase {
// 	return &pemeriksaanDokterTrimester3Usecase{repo: repo, labRepo: labRepo}
// }

// func (u *pemeriksaanDokterTrimester3Usecase) mapRequestToDokter(req *PemeriksaanDokterTrimester3Request) *models.PemeriksaanDokterTrimester3 {
// 	today := time.Now()
// 	dokter := &models.PemeriksaanDokterTrimester3{
// 		KehamilanID:               req.KehamilanID,
// 		NamaDokter:                req.NamaDokter,
// 		TanggalPeriksa:            &today, // Set to today
// 		KonsepAnamnesaPemeriksaan: req.KonsepAnamnesaPemeriksaan,
// 		GambarUSG:                 req.GambarUSG,
// 		FisikKonjungtiva:          req.FisikKonjungtiva,
// 		FisikSklera:               req.FisikSklera,
// 		FisikKulit:                req.FisikKulit,
// 		FisikLeher:                req.FisikLeher,
// 		FisikGigiMulut:            req.FisikGigiMulut,
// 		FisikTHT:                  req.FisikTHT,
// 		FisikDadaJantung:          req.FisikDadaJantung,
// 		FisikDadaParu:             req.FisikDadaParu,
// 		FisikPerut:                req.FisikPerut,
// 		FisikTungkai:              req.FisikTungkai,

// 		USGTrimester3Dilakukan:                   req.USGTrimester3Dilakukan,
// 		UKBerdasarkanUSGTrimester1Minggu:         req.UKBerdasarkanUSGTrimester1Minggu,
// 		UKBerdasarkanHPHTMinggu:                  req.UKBerdasarkanHPHTMinggu,
// 		UKBerdasarkanBiometriUSGTrimester3Minggu: req.UKBerdasarkanBiometriUSGTrimester3Minggu,
// 		SelisihUK3MingguAtauLebih:                req.SelisihUK3MingguAtauLebih,
// 		USGJumlahBayi:                            req.USGJumlahBayi,
// 		USGLetakBayi:                             req.USGLetakBayi,
// 		USGPresentasiBayi:                        req.USGPresentasiBayi,
// 		USGKeadaanBayi:                           req.USGKeadaanBayi,
// 		USGDJNilai:                               req.USGDJNilai,
// 		USGDJJStatus:                             req.USGDJJStatus,
// 		USGLokasiPlasenta:                        req.USGLokasiPlasenta,
// 		USGCairanKetubanSDPCm:                    req.USGCairanKetubanSDPCm,
// 		USGCairanKetubanStatus:                   req.USGCairanKetubanStatus,

// 		BiometriBPDCm:        req.BiometriBPDCm,
// 		BiometriBPDMinggu:    req.BiometriBPDMinggu,
// 		BiometriHCCm:         req.BiometriHCCm,
// 		BiometriHCMinggu:     req.BiometriHCMinggu,
// 		BiometriACCm:         req.BiometriACCm,
// 		BiometriACMinggu:     req.BiometriACMinggu,
// 		BiometriFLCm:         req.BiometriFLCm,
// 		BiometriFLMinggu:     req.BiometriFLMinggu,
// 		BiometriEFWTBJGram:   req.BiometriEFWTBJGram,
// 		BiometriEFWTBJMinggu: req.BiometriEFWTBJMinggu,

// 		USGKecurigaanTemuanAbnormal: req.USGKecurigaanTemuanAbnormal,
// 		USGKeteranganTemuanAbnormal: req.USGKeteranganTemuanAbnormal,

// 		HasilUSGCatatan:                       req.HasilUSGCatatan,
// 		TanggalLab:                            &today, // Set to today
// 		LabHemoglobinHasil:                    req.LabHemoglobinHasil,
// 		LabHemoglobinRencana:                  req.LabHemoglobinRencana,
// 		LabProteinUrinHasil:                   req.LabProteinUrinHasil,
// 		LabProteinUrinRencana:                 req.LabProteinUrinRencana,
// 		LabUrinReduksiHasil:                   req.LabUrinReduksiHasil,
// 		LabUrinReduksiRencana:                 req.LabUrinReduksiRencana,
// 		SkriningJiwaHasil:                     req.SkriningJiwaHasil,
// 		SkriningJiwaTindakLanjut:              req.SkriningJiwaTindakLanjut,
// 		SkriningJiwaPerluRujukan:              req.SkriningJiwaPerluRujukan,
// 		RencanaKonsultasiGizi:                 req.RencanaKonsultasiGizi,
// 		RencanaKonsultasiKebidanan:            req.RencanaKonsultasiKebidanan,
// 		RencanaKonsultasiAnak:                 req.RencanaKonsultasiAnak,
// 		RencanaKonsultasiPenyakitDalam:        req.RencanaKonsultasiPenyakitDalam,
// 		RencanaKonsultasiNeurologi:            req.RencanaKonsultasiNeurologi,
// 		RencanaKonsultasiTHT:                  req.RencanaKonsultasiTHT,
// 		RencanaKonsultasiPsikiatri:            req.RencanaKonsultasiPsikiatri,
// 		RencanaKonsultasiLainLain:             req.RencanaKonsultasiLainLain,
// 		RencanaProsesMelahirkan:               req.RencanaProsesMelahirkan,
// 		RencanaKontrasepsiAKDR:                req.RencanaKontrasepsiAKDR,
// 		RencanaKontrasepsiPil:                 req.RencanaKontrasepsiPil,
// 		RencanaKontrasepsiSuntik:              req.RencanaKontrasepsiSuntik,
// 		RencanaKontrasepsiSteril:              req.RencanaKontrasepsiSteril,
// 		RencanaKontrasepsiMAL:                 req.RencanaKontrasepsiMAL,
// 		RencanaKontrasepsiImplan:              req.RencanaKontrasepsiImplan,
// 		RencanaKontrasepsiBelumMemilih:        req.RencanaKontrasepsiBelumMemilih,
// 		KebutuhanKonseling:                    req.KebutuhanKonseling,
// 		Penjelasan:                            req.Penjelasan,
// 		KesimpulanRekomendasiTempatMelahirkan: req.KesimpulanRekomendasiTempatMelahirkan,
// 	}
// 	return dokter
// }

// func (u *pemeriksaanDokterTrimester3Usecase) mapRequestToLab(req *PemeriksaanDokterTrimester3Request, trimester int32) *models.PemeriksaanLaboratoriumJiwa {
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
// 		SkriningJiwaHasil:                req.SkriningJiwaHasilTr,
// 		SkriningJiwaTindakLanjut:         req.SkriningJiwaTindakLanjutTr,
// 		SkriningJiwaPerluRujukan:         req.SkriningJiwaPerluRujukanTr,
// 		Kesimpulan:                       req.KesimpulanJiwaTr,
// 		Rekomendasi:                      req.RekomendasiJiwaTr,
// 	}
// 	if req.SkriningJiwaHasilTr != "" {
// 		tab := &today
// 		lab.TanggalSkriningJiwa = tab // Set to today if skrining data provided
// 	}
// 	return lab
// }

// func (u *pemeriksaanDokterTrimester3Usecase) Create(req *PemeriksaanDokterTrimester3Request) error {
// 	if req.KehamilanID == 0 {
// 		return errors.New("kehamilan_id wajib diisi")
// 	}
// 	dokter := u.mapRequestToDokter(req)
// 	lab := u.mapRequestToLab(req, 3)
// 	return u.repo.CreateWithLab(dokter, lab)
// }

// func (u *pemeriksaanDokterTrimester3Usecase) Update(id int32, req *PemeriksaanDokterTrimester3Request) error {
// 	existing, err := u.repo.FindByID(id)
// 	if err != nil {
// 		return errors.New("data tidak ditemukan")
// 	}
// 	dokter := u.mapRequestToDokter(req)
// 	dokter.KehamilanID = existing.KehamilanID
// 	lab := u.mapRequestToLab(req, 3)
// 	return u.repo.UpdateWithLab(id, dokter, lab)
// }

// func (u *pemeriksaanDokterTrimester3Usecase) GetByID(id int32) (*PemeriksaanDokterTrimester3Response, error) {
// 	dokter, err := u.repo.FindByID(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	lab, _ := u.labRepo.FindByKehamilanIDAndTrimester(dokter.KehamilanID, 3)
// 	return &PemeriksaanDokterTrimester3Response{Dokter: dokter, Lab: lab}, nil
// }

// func (u *pemeriksaanDokterTrimester3Usecase) GetByKehamilanID(kehamilanID int32) ([]PemeriksaanDokterTrimester3Response, error) {
// 	list, err := u.repo.FindByKehamilanID(kehamilanID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var result []PemeriksaanDokterTrimester3Response
// 	for _, d := range list {
// 		lab, _ := u.labRepo.FindByKehamilanIDAndTrimester(d.KehamilanID, 3)
// 		result = append(result, PemeriksaanDokterTrimester3Response{Dokter: &d, Lab: lab})
// 	}
// 	return result, nil
// }

// func (u *pemeriksaanDokterTrimester3Usecase) Delete(id int32) error {
// 	return u.repo.Delete(id)
// }

// // MODUL IBU - single (existing)
// func (u *pemeriksaanDokterTrimester3Usecase) GetMine(userID int32) (*models.PemeriksaanDokterTrimester3, error) {
// 	return u.repo.FindMineByUserID(userID)
// }

// // MODUL IBU - list (BARU)
// func (u *pemeriksaanDokterTrimester3Usecase) GetAllMine(userID int32) ([]models.PemeriksaanDokterTrimester3, error) {
// 	return u.repo.FindAllMineByUserID(userID)
// }

// package usecases

// import (
// 	"errors"
// 	"monitoring-service/app/models"
// 	"monitoring-service/app/repositories"
// )

// type PemeriksaanDokterTrimester3Usecase interface {
// 	Create(p *models.PemeriksaanDokterTrimester3) error
// 	GetByID(id int32) (*models.PemeriksaanDokterTrimester3, error)
// 	GetMine(userID int32) (*models.PemeriksaanDokterTrimester3, error)    // single (lama)
// 	GetAllMine(userID int32) ([]models.PemeriksaanDokterTrimester3, error) // list (BARU)
// 	GetByKehamilanID(kehamilanID int32) ([]models.PemeriksaanDokterTrimester3, error)
// 	Update(p *models.PemeriksaanDokterTrimester3) error
// 	Delete(id int32) error
// }

// type pemeriksaanDokterTrimester3Usecase struct {
// 	repo *repositories.PemeriksaanDokterTrimester3Repository
// }

// func NewPemeriksaanDokterTrimester3Usecase(repo *repositories.PemeriksaanDokterTrimester3Repository) PemeriksaanDokterTrimester3Usecase {
// 	return &pemeriksaanDokterTrimester3Usecase{repo: repo}
// }

// func (u *pemeriksaanDokterTrimester3Usecase) Create(p *models.PemeriksaanDokterTrimester3) error {
// 	if p.KehamilanID == 0 {
// 		return errors.New("kehamilan_id wajib diisi")
// 	}
// 	return u.repo.Create(p)
// }

// func (u *pemeriksaanDokterTrimester3Usecase) GetByID(id int32) (*models.PemeriksaanDokterTrimester3, error) {
// 	return u.repo.FindByID(id)
// }

// func (u *pemeriksaanDokterTrimester3Usecase) GetByKehamilanID(kehamilanID int32) ([]models.PemeriksaanDokterTrimester3, error) {
// 	return u.repo.FindByKehamilanID(kehamilanID)
// }

// func (u *pemeriksaanDokterTrimester3Usecase) Update(p *models.PemeriksaanDokterTrimester3) error {
// 	_, err := u.repo.FindByID(p.ID)
// 	if err != nil {
// 		return errors.New("data pemeriksaan dokter trimester 3 tidak ditemukan")
// 	}
// 	return u.repo.Update(p)
// }

// func (u *pemeriksaanDokterTrimester3Usecase) Delete(id int32) error {
// 	return u.repo.Delete(id)
// }

// // MODUL IBU - single (existing)
// func (u *pemeriksaanDokterTrimester3Usecase) GetMine(userID int32) (*models.PemeriksaanDokterTrimester3, error) {
// 	return u.repo.FindMineByUserID(userID)
// }

// // MODUL IBU - list (BARU)
// func (u *pemeriksaanDokterTrimester3Usecase) GetAllMine(userID int32) ([]models.PemeriksaanDokterTrimester3, error) {
// 	return u.repo.FindAllMineByUserID(userID)
// }

// package usecases

// import (
//     "errors"
//     "monitoring-service/app/models"
//     "monitoring-service/app/repositories"
// )

// type PemeriksaanDokterTrimester3Usecase interface {
//     Create(p *models.PemeriksaanDokterTrimester3) error
//     GetByID(id int32) (*models.PemeriksaanDokterTrimester3, error)
//     GetMine(userID int32) (*models.PemeriksaanDokterTrimester3, error)    // single (lama)
//     GetAllMine(userID int32) ([]models.PemeriksaanDokterTrimester3, error) // list (BARU)
//     GetByKehamilanID(kehamilanID int32) ([]models.PemeriksaanDokterTrimester3, error)
//     Update(id int32, p *models.PemeriksaanDokterTrimester3) error // DIUBAH: menambahkan id int32
//     Delete(id int32) error
// }

// type pemeriksaanDokterTrimester3Usecase struct {
//     repo *repositories.PemeriksaanDokterTrimester3Repository
// }

// func NewPemeriksaanDokterTrimester3Usecase(repo *repositories.PemeriksaanDokterTrimester3Repository) PemeriksaanDokterTrimester3Usecase {
//     return &pemeriksaanDokterTrimester3Usecase{repo: repo}
// }

// func (u *pemeriksaanDokterTrimester3Usecase) Create(p *models.PemeriksaanDokterTrimester3) error {
//     if p.KehamilanID == 0 {
//         return errors.New("kehamilan_id wajib diisi")
//     }
//     return u.repo.Create(p)
// }

// func (u *pemeriksaanDokterTrimester3Usecase) GetByID(id int32) (*models.PemeriksaanDokterTrimester3, error) {
//     return u.repo.FindByID(id)
// }

// func (u *pemeriksaanDokterTrimester3Usecase) GetByKehamilanID(kehamilanID int32) ([]models.PemeriksaanDokterTrimester3, error) {
//     return u.repo.FindByKehamilanID(kehamilanID)
// }

// // DIUBAH: Logic disesuaikan dengan parameter id
// func (u *pemeriksaanDokterTrimester3Usecase) Update(id int32, p *models.PemeriksaanDokterTrimester3) error {
//     _, err := u.repo.FindByID(id)
//     if err != nil {
//         return errors.New("data pemeriksaan dokter trimester 3 tidak ditemukan")
//     }
//     p.ID = id // Pastikan ID yang diupdate sesuai parameter
//     return u.repo.Update(p)
// }

// func (u *pemeriksaanDokterTrimester3Usecase) Delete(id int32) error {
//     return u.repo.Delete(id)
// }

// // MODUL IBU - single (existing)
// func (u *pemeriksaanDokterTrimester3Usecase) GetMine(userID int32) (*models.PemeriksaanDokterTrimester3, error) {
//     return u.repo.FindMineByUserID(userID)
// }

// // MODUL IBU - list (BARU)
// func (u *pemeriksaanDokterTrimester3Usecase) GetAllMine(userID int32) ([]models.PemeriksaanDokterTrimester3, error) {
//     return u.repo.FindAllMineByUserID(userID)
// }

package usecases

import (
    "errors"
    "monitoring-service/app/models"
    "monitoring-service/app/repositories"
)

type PemeriksaanDokterTrimester3Usecase interface {
    Create(p *models.PemeriksaanDokterTrimester3) error
    GetByID(id int32) (*models.PemeriksaanDokterTrimester3, error)
    GetMine(userID int32) (*models.PemeriksaanDokterTrimester3, error)
    GetAllMine(userID int32) ([]models.PemeriksaanDokterTrimester3, error)
    GetByKehamilanID(kehamilanID int32) ([]models.PemeriksaanDokterTrimester3, error)
    Update(id int32, p *models.PemeriksaanDokterTrimester3) error
    Delete(id int32) error
}

type pemeriksaanDokterTrimester3Usecase struct {
    repo    *repositories.PemeriksaanDokterTrimester3Repository
    labRepo *repositories.PemeriksaanLaboratoriumJiwaRepository // DIUBAH: Ditambahkan
}

// DIUBAH: Menambahkan parameter labRepo
func NewPemeriksaanDokterTrimester3Usecase(repo *repositories.PemeriksaanDokterTrimester3Repository, labRepo *repositories.PemeriksaanLaboratoriumJiwaRepository) PemeriksaanDokterTrimester3Usecase {
    return &pemeriksaanDokterTrimester3Usecase{repo: repo, labRepo: labRepo}
}

func (u *pemeriksaanDokterTrimester3Usecase) Create(p *models.PemeriksaanDokterTrimester3) error {
    if p.KehamilanID == 0 {
        return errors.New("kehamilan_id wajib diisi")
    }
    return u.repo.Create(p)
}

func (u *pemeriksaanDokterTrimester3Usecase) GetByID(id int32) (*models.PemeriksaanDokterTrimester3, error) {
    return u.repo.FindByID(id)
}

func (u *pemeriksaanDokterTrimester3Usecase) GetByKehamilanID(kehamilanID int32) ([]models.PemeriksaanDokterTrimester3, error) {
    return u.repo.FindByKehamilanID(kehamilanID)
}

func (u *pemeriksaanDokterTrimester3Usecase) Update(id int32, p *models.PemeriksaanDokterTrimester3) error {
    _, err := u.repo.FindByID(id)
    if err != nil {
        return errors.New("data pemeriksaan dokter trimester 3 tidak ditemukan")
    }
    p.ID = id
    return u.repo.Update(p)
}

func (u *pemeriksaanDokterTrimester3Usecase) Delete(id int32) error {
    return u.repo.Delete(id)
}

func (u *pemeriksaanDokterTrimester3Usecase) GetMine(userID int32) (*models.PemeriksaanDokterTrimester3, error) {
    return u.repo.FindMineByUserID(userID)
}

func (u *pemeriksaanDokterTrimester3Usecase) GetAllMine(userID int32) ([]models.PemeriksaanDokterTrimester3, error) {
    return u.repo.FindAllMineByUserID(userID)
}