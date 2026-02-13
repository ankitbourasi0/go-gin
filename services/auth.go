package services

import (
	"errors"
	internal "gin-tutorial/internal/models"

	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func (a *AuthService) InitAuthService(db *gorm.DB) *AuthService {
	a.db = db
	a.db.AutoMigrate(&internal.User{}) //auto register table in db
	return a
}

func (a *AuthService) Login(email *string, password *string) (*internal.User, error) {
	if email == nil || password == nil {
		return nil, errors.New("email or password cant be null")
	}
	var user internal.User
	if err := a.db.Where("email = ?", email).Where("password = ?", password).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil

}

func (a *AuthService) Register(email *string, password *string) (*internal.User, error) {
	if email == nil || password == nil {
		return nil, errors.New("email or password cant be null")
	}
	var user internal.User
	user.Email = *email
	user.Password = *password
	if err := a.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil

}
