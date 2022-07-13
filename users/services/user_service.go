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

func (s UserService) CreateUser(userIn models.UserIn) *models.UserOut {
	var user models.User = models.User{Username: userIn.Username, Password: userIn.Password}

	createdUser := s.UserDao.CreateUser(&user)
	return &models.UserOut{ID: createdUser.ID, Username: createdUser.Username}
}

func (s UserService) GetUserById(id uint) (*models.UserOut, error) {
	user, err := s.UserDao.GetById(id)
	if err != nil {
		return nil, err
	}
	userOut := models.UserOut{ID: user.ID, Username: user.Username}
	return &userOut, nil
}

func (s UserService) DeleteUserById(id uint) error {
	user, err := s.UserDao.GetById(id)
	if err != nil {
		return err
	}

	return s.UserDao.DeleteUser(user)
}
