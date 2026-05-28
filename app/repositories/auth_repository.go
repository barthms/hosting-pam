package repositories

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/pkg/customerror"

	"gorm.io/gorm"
)

func (m *Main) GetRoleByName(roleName string) (*models.Role, error) {
	var role models.Role
	if err := m.postgres.Where("name = ?", roleName).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerror.NewNotFoundError("role tidak ditemukan")
		}
		return nil, err
	}
	return &role, nil
}

func (m *Main) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := m.postgres.Preload("Role").Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerror.NewNotFoundError("email belum terdaftar")
		}
		return nil, err
	}
	return &user, nil
}

func (m *Main) GetUserByPhoneNumber(phoneNumber string) (*models.User, error) {
	var user models.User
	if err := m.postgres.Preload("Role").Where("nomor_telepon = ?", phoneNumber).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customerror.NewNotFoundError("nomor hp belum terdaftar")
		}
		return nil, err
	}
	return &user, nil
}

func (m *Main) CreateUser(user *models.User) error {
	if err := m.postgres.Create(user).Error; err != nil {
		return err
	}
	return nil
}
