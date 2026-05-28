package usecases

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"
	"time"

	"monitoring-service/app/models"
	"monitoring-service/pkg/customerror"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type roleDestination struct {
	TargetApp     string
	RedirectRoute string
}

var roleDestinations = map[string]roleDestination{
	"Admin":            {TargetApp: "website", RedirectRoute: "/dashboard/admin"},
	"Dokter":           {TargetApp: "website", RedirectRoute: "/dashboard/dokter"},
	"Tenaga-kesehatan": {TargetApp: "website", RedirectRoute: "/dashboard/tenaga-kesehatan"},
	"Kader":            {TargetApp: "mobile", RedirectRoute: "/mobile/home-kader"},
	"Bidan":            {TargetApp: "mobile", RedirectRoute: "/mobile/home-bidan"},
	"Ibu":              {TargetApp: "mobile", RedirectRoute: "/mobile/home-orangtua"},
}

var roleAliases = map[string]string{
	"admin":            "Admin",
	"dokter":           "Dokter",
	"tenagakesehatan":  "Tenaga-kesehatan",
	"tenaga-kesehatan": "Tenaga-kesehatan",
	"tenaga kesehatan": "Tenaga-kesehatan",
	"kader":            "Kader",
	"bidan":            "Bidan",
	"orangtua":         "Orangtua",
	"orang tua":        "Orangtua",
	"orang-tua":        "Orangtua",
	"Ibu":              "Orangtua",
}

var phonePattern = regexp.MustCompile(`^\+62[0-9]{8,13}$`)

func normalizeKey(raw string) string {
	raw = strings.TrimSpace(strings.ToLower(raw))
	raw = strings.ReplaceAll(raw, "-", "")
	raw = strings.ReplaceAll(raw, "_", "")
	raw = strings.ReplaceAll(raw, " ", "")
	return raw
}

func normalizeRoleName(roleName string) string {
	roleName = strings.TrimSpace(roleName)
	if roleName == "" {
		return ""
	}

	if canonical, ok := roleAliases[normalizeKey(roleName)]; ok {
		return canonical
	}

	return roleName
}

func roleRedirect(roleName string) (roleDestination, bool) {
	destination, ok := roleDestinations[roleName]
	return destination, ok
}

func isEmail(input string) bool {
	_, err := mail.ParseAddress(input)
	return err == nil
}

func normalizePhoneNumber(input string) (string, error) {
	phone := strings.TrimSpace(input)
	if phone == "" {
		return "", customerror.NewBadRequestError("nomor hp wajib diisi")
	}

	phone = strings.ReplaceAll(phone, " ", "")
	phone = strings.ReplaceAll(phone, "-", "")
	phone = strings.ReplaceAll(phone, "(", "")
	phone = strings.ReplaceAll(phone, ")", "")

	switch {
	case strings.HasPrefix(phone, "+62"):
		// already normalized
	case strings.HasPrefix(phone, "62"):
		phone = "+" + phone
	case strings.HasPrefix(phone, "08"):
		phone = "+62" + phone[1:]
	case strings.HasPrefix(phone, "8"):
		phone = "+62" + phone
	default:
		return "", customerror.NewBadRequestError("format nomor hp tidak valid")
	}

	if !phonePattern.MatchString(phone) {
		return "", customerror.NewBadRequestError("format nomor hp tidak valid")
	}

	return phone, nil
}

func validateRegisterInput(req *models.RegisterRequest) error {
	if req.Name == "" || req.Email == "" || req.PhoneNumber == "" || req.Password == "" || req.RoleName == "" {
		return customerror.NewBadRequestError("name, email, phone_number, password, dan role_name wajib diisi")
	}

	if len(req.Name) < 3 {
		return customerror.NewBadRequestError("name minimal 3 karakter")
	}

	if _, err := mail.ParseAddress(req.Email); err != nil {
		return customerror.NewBadRequestError("format email tidak valid")
	}

	if len(req.Password) < 8 {
		return customerror.NewBadRequestError("password minimal 8 karakter")
	}

	return nil
}

func (m *Main) buildAccessToken(user *models.User, destination roleDestination) (tokenString string, expiresIn int64, err error) {
	now := time.Now()
	expiry := now.Add(time.Duration(m.config.JWTAccessTokenMins) * time.Minute)

	claims := models.AuthClaims{
		UserID:        user.ID,
		Email:         user.Email,
		PhoneNumber:   user.PhoneNumber,
		Role:          user.Role.Name,
		TargetApp:     destination.TargetApp,
		RedirectRoute: destination.RedirectRoute,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   fmt.Sprintf("%d", user.ID),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiry),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(m.config.JWTSecret))
	if err != nil {
		return "", 0, err
	}

	return tokenString, int64(time.Until(expiry).Seconds()), nil
}

func (m *Main) Register(req *models.RegisterRequest) error {
	if req == nil {
		return customerror.NewBadRequestError("request tidak valid")
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	req.PhoneNumber = strings.TrimSpace(req.PhoneNumber)
	req.RoleName = normalizeRoleName(req.RoleName)

	if err := validateRegisterInput(req); err != nil {
		return err
	}

	normalizedPhoneNumber, err := normalizePhoneNumber(req.PhoneNumber)
	if err != nil {
		return err
	}
	req.PhoneNumber = normalizedPhoneNumber

	if _, err := m.repository.GetUserByEmail(req.Email); err == nil {
		return customerror.NewBadRequestError("email sudah terdaftar")
	} else {
		if _, ok := err.(customerror.NotFoundError); !ok {
			return err
		}
	}

	if _, err := m.repository.GetUserByPhoneNumber(req.PhoneNumber); err == nil {
		return customerror.NewBadRequestError("nomor hp sudah terdaftar")
	} else {
		if _, ok := err.(customerror.NotFoundError); !ok {
			return err
		}
	}

	role, err := m.repository.GetRoleByName(req.RoleName)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return customerror.NewInternalServiceError("gagal memproses password")
	}

	user := &models.User{
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    string(hashedPassword),
		RoleID:      role.ID,
	}

	if err := m.repository.CreateUser(user); err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") || strings.Contains(strings.ToLower(err.Error()), "unique") {
			if strings.Contains(strings.ToLower(err.Error()), "phone") {
				return customerror.NewBadRequestError("nomor hp sudah terdaftar")
			}
			return customerror.NewBadRequestError("email sudah terdaftar")
		}
		return err
	}

	return nil
}

func (m *Main) Login(req *models.LoginRequest) (*models.LoginResponse, error) {
	if req == nil {
		return nil, customerror.NewBadRequestError("request tidak valid")
	}

	identifier := strings.TrimSpace(req.Identifier)
	if identifier == "" {
		identifier = strings.TrimSpace(req.Email)
	}

	if identifier == "" || req.Password == "" {
		return nil, customerror.NewBadRequestError("identifier/email dan password wajib diisi")
	}

	var user *models.User
	var err error
	if isEmail(identifier) {
		user, err = m.repository.GetUserByEmail(strings.ToLower(identifier))
	} else {
		normalizedPhoneNumber, nErr := normalizePhoneNumber(identifier)
		if nErr != nil {
			return nil, customerror.NewBadRequestError("identifier harus email atau nomor hp valid")
		}
		user, err = m.repository.GetUserByPhoneNumber(normalizedPhoneNumber)
	}

	if err != nil {
		if _, ok := err.(customerror.NotFoundError); ok {
			return nil, customerror.NewBadRequestError("email/nomor hp atau password salah")
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, customerror.NewBadRequestError("email/nomor hp atau password salah")
	}

	canonicalRoleName := normalizeRoleName(user.Role.Name)
	if user.PendudukID != nil {
		switch canonicalRoleName {
		case "Bidan":
			bidan, bErr := m.repository.Bidan.FindByPendudukID(int32(*user.PendudukID))
			if bErr != nil || strings.ToLower(strings.TrimSpace(bidan.Status)) != "aktif" {
				return nil, customerror.NewBadRequestError("akun bidan nonaktif")
			}
		case "Kader":
			fmt.Println("PendudukID:", *user.PendudukID)

			kader, kErr := m.repository.Kader.FindByPendudukID(int32(*user.PendudukID))

			fmt.Println("Kader:", kader)
			fmt.Println("Error:", kErr)

			if kErr != nil {
				fmt.Println("Find kader error:", kErr)
				return nil, customerror.NewBadRequestError("akun kader nonaktif")
			}

			fmt.Println("Status kader:", kader.Status)

			if strings.ToLower(strings.TrimSpace(kader.Status)) != "aktif" {
				return nil, customerror.NewBadRequestError("akun kader nonaktif")
			}
		}
	}

	destination, ok := roleRedirect(canonicalRoleName)
	if !ok {
		return nil, customerror.NewInternalServiceError("role belum memiliki mapping target aplikasi")
	}
	user.Role.Name = canonicalRoleName

	accessToken, expiresIn, err := m.buildAccessToken(user, destination)
	if err != nil {
		return nil, customerror.NewInternalServiceError("gagal membuat access token")
	}

	res := &models.LoginResponse{
		AccessToken:   accessToken,
		TokenType:     "Bearer",
		ExpiresIn:     expiresIn,
		UserID:        user.ID,
		Name:          user.Name,
		Email:         user.Email,
		PhoneNumber:   user.PhoneNumber,
		Role:          user.Role.Name,
		TargetApp:     destination.TargetApp,
		RedirectRoute: destination.RedirectRoute,
	}
	// if canonicalRoleName == "Ibu" {

	// 	_ = m.GenerateJadwalImunisasi(
	// 		user.ID,
	// 	)
	// }
	return res, nil
}
