package repositories

import (
	"monitoring-service/app/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Role").Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Role").Where("nama = ?", username).First(&user).Error
	return &user, err
}

func (r *UserRepository) FindByPhoneNumber(phone string) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Role").Where("nomor_telepon = ?", phone).First(&user).Error
	return &user, err
}

func (r *UserRepository) FindByID(id int32) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Role").First(&user, id).Error
	return &user, err
}

func (r *UserRepository) FindByIDExceptEmail(email string, exceptID int32) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Role").Where("email = ? AND id <> ?", email, exceptID).First(&user).Error
	return &user, err
}

func (r *UserRepository) FindByPhoneNumberExceptID(phone string, exceptID int32) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Role").Where("nomor_telepon = ? AND id <> ?", phone, exceptID).First(&user).Error
	return &user, err
}

func (r *UserRepository) FindByPendudukID(pendudukID int64) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Role").Where("penduduk_id = ?", pendudukID).First(&user).Error
	return &user, err
}

func (r *UserRepository) FindByKartuKeluargaID(kartuKeluargaID int64) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Role").
		Joins("JOIN penduduk p ON p.id = pengguna.penduduk_id").
		Where("p.kartu_keluarga_id = ? AND p.deleted_at IS NULL", kartuKeluargaID).
		Order("pengguna.id ASC").
		First(&user).Error
	return &user, err
}

func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}
