package services

import (
	"github.com/erfanmorsali/gin-simple-app.git/auth/dtos"
	"github.com/erfanmorsali/gin-simple-app.git/database/models"
	"github.com/erfanmorsali/gin-simple-app.git/users/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userDao interfaces.UserDao
}

func NewAuthService(userDao interfaces.UserDao) *AuthService {
	return &AuthService{
		userDao: userDao,
	}
}

func (s AuthService) Login(email string, password string) bool {
	user, err := s.userDao.GetUserByEmail(email)
	if err != nil {
		return false
	}

	if !s.CheckPasswordHash(password, user.Password) {
		return false
	}

	return true
}

func (s AuthService) Register(registerInput dtos.RegisterIn) (*dtos.RegisterOut, error) {
	password, err := s.HashPassword(registerInput.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{Username: registerInput.Username, Email: registerInput.Email, Password: password}
	createdUser, err := s.userDao.Create(&user)
	if err != nil {
		return nil, err
	}

	return &dtos.RegisterOut{Email: createdUser.Email, Username: createdUser.Username, ID: createdUser.ID}, nil
}

func (s AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s AuthService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
