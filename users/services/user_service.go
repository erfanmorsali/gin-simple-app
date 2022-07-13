package services

import (
	"github.com/erfanmorsali/gin-simple-app.git/users/interfaces"
	"github.com/erfanmorsali/gin-simple-app.git/users/models"
)

type UserService struct {
	UserDao interfaces.UserDao
}

func NewUserService(userDao interfaces.UserDao) *UserService {
	return &UserService{UserDao: userDao}
}

func (s UserService) GetAllUsers() []models.UserOut {
	var result []models.UserOut
	users := s.UserDao.GetAll()
	for _, user := range users {
		userOut := models.UserOut{ID: user.ID, Username: user.Username}
		result = append(result, userOut)
	}
	return result
}
