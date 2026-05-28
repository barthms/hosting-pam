package usecases

// import (
// 	"errors"
// 	"time"

// 	"monitoring-service/app/models"
// 	"monitoring-service/app/repositories"

// 	"golang.org/x/crypto/bcrypt"
// )

// type RegisterOrangTuaRequest struct {
// 	NoKK               string `json:"no_kk"`
// 	TanggalTerbit      string `json:"tanggal_terbit"` // YYYY-MM-DD
// 	NIK                string `json:"nik"`
// 	Dusun              string `json:"dusun"`
// 	NamaLengkap        string `json:"nama_lengkap"`
// 	GolonganDarah      string `json:"golongan_darah"`
// 	JenisKelamin       string `json:"jenis_kelamin"` // Laki-Laki / Perempuan
// 	TempatLahir        string `json:"tempat_lahir"`
// 	TanggalLahir       string `json:"tanggal_lahir"` // YYYY-MM-DD
// 	Pekerjaan          string `json:"pekerjaan"`
// 	PendidikanTerakhir string `json:"pendidikan_terakhir"`
// 	Username           string `json:"username"`
// 	Email              string `json:"email"`
// 	Password           string `json:"password"`
// 	PhoneNumber        string `json:"phone_number"`
// }

// type RegisterOrangTuaUsecase struct {
// 	userRepo         *repositories.UserRepository
// 	roleRepo         *repositories.RoleRepository // perlu dibuat
// 	kkRepo           *repositories.KartuKeluargaRepository
// 	kependudukanRepo *repositories.KependudukanRepository
// 	ibuRepo          *repositories.IbuRepository
// }

// func NewRegisterOrangTuaUsecase(
// 	userRepo *repositories.UserRepository,
// 	roleRepo *repositories.RoleRepository,
// 	kkRepo *repositories.KartuKeluargaRepository,
// 	kependudukanRepo *repositories.KependudukanRepository,
// 	ibuRepo *repositories.IbuRepository,
// ) *RegisterOrangTuaUsecase {
// 	return &RegisterOrangTuaUsecase{
// 		userRepo:         userRepo,
// 		roleRepo:         roleRepo,
// 		kkRepo:           kkRepo,
// 		kependudukanRepo: kependudukanRepo,
// 		ibuRepo:          ibuRepo,
// 	}
// }

// func (u *RegisterOrangTuaUsecase) Register(req *RegisterOrangTuaRequest) error {
// 	// Validasi input
// 	if req.NoKK == "" || req.NIK == "" || req.NamaLengkap == "" || req.Username == "" || req.Email == "" || req.Password == "" {
// 		return errors.New("semua field wajib diisi")
// 	}

// 	// Cek apakah username/email/phone sudah terdaftar
// 	if _, err := u.userRepo.FindByUsername(req.Username); err == nil {
// 		return errors.New("username sudah terdaftar")
// 	}
// 	if _, err := u.userRepo.FindByEmail(req.Email); err == nil {
// 		return errors.New("email sudah terdaftar")
// 	}
// 	if _, err := u.userRepo.FindByPhoneNumber(req.PhoneNumber); err == nil {
// 		return errors.New("nomor telepon sudah terdaftar")
// 	}

// 	// Cek apakah NoKK sudah terdaftar
// 	if _, err := u.kkRepo.FindByNoKK(req.NoKK); err == nil {
// 		return errors.New("nomor KK sudah terdaftar")
// 	}

// 	// Cek apakah NIK sudah terdaftar di kependudukan
// 	if _, err := u.kependudukanRepo.FindByNIK(req.NIK); err == nil {
// 		return errors.New("NIK sudah terdaftar")
// 	}

// 	// Ambil role "Orangtua"
// 	role, err := u.roleRepo.FindByName("Orangtua")
// 	if err != nil {
// 		return errors.New("role Orangtua tidak ditemukan")
// 	}

// 	// Hash password
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}

// 	// Buat user
// 	user := &models.User{
// 		RoleID:      role.ID,
// 		Name:        req.Username,
// 		Email:       req.Email,
// 		Password:    string(hashedPassword),
// 		PhoneNumber: req.PhoneNumber,
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 	}

// 	// Buat KK
// 	tanggalTerbit, _ := time.Parse("2006-01-02", req.TanggalTerbit)
// 	kk := &models.KartuKeluarga{
// 		NoKK:          req.NoKK,
// 		TanggalTerbit: &tanggalTerbit,
// 		CreatedAt:     time.Now(),
// 		UpdatedAt:     time.Now(),
// 	}
// 	if err := u.kkRepo.Create(kk); err != nil {
// 		return err
// 	}

// 	// Buat data kependudukan
// 	tanggalLahir, _ := time.Parse("2006-01-02", req.TanggalLahir)
// 	kependudukan := &models.Kependudukan{
// 		KartuKeluargaID:    &kk.ID,
// 		NIK:                req.NIK,
// 		Dusun:              req.Dusun,
// 		NamaLengkap:        req.NamaLengkap,
// 		GolonganDarah:      req.GolonganDarah,
// 		JenisKelamin:       req.JenisKelamin,
// 		TempatLahir:        req.TempatLahir,
// 		TanggalLahir:       &tanggalLahir,
// 		Pekerjaan:          req.Pekerjaan,
// 		PendidikanTerakhir: req.PendidikanTerakhir,
// 		NomorTelepon:       req.PhoneNumber,
// 		KedudukanKeluarga:  "Kepala Keluarga",
// 		CreatedAt:          time.Now(),
// 		UpdatedAt:          time.Now(),
// 	}
// 	if err := u.kependudukanRepo.Create(kependudukan); err != nil {
// 		return err
// 	}

// 	pendudukID := int64(kependudukan.IDKependudukan)
// 	user.PendudukID = &pendudukID
// 	if err := u.userRepo.Create(user); err != nil {
// 		return err
// 	}

// 	// Buat data ibu (karena orang tua yang register adalah ibu? Asumsikan orang tua adalah ibu)
// 	// Tapi bisa juga ayah, tapi kita sederhanakan: orang tua = ibu
// 	ibu := &models.Ibu{
// 		IDKependudukan:  kependudukan.IDKependudukan,
// 		StatusKehamilan: "TIDAK HAMIL",
// 		CreatedAt:       time.Now(),
// 		UpdatedAt:       time.Now(),
// 	}
// 	if err := u.ibuRepo.Create(ibu); err != nil {
// 		return err
// 	}

// 	return nil
// }
